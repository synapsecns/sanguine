package moe

import (
	"bufio"
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"golang.org/x/crypto/ssh"
	"io"
	"net"
	"os"
	"regexp"
	"strings"
	"time"
)

// Get private key for ssh authentication.
func generatePrivateKey() (ssh.Signer, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, fmt.Errorf("failed to generate private key: %w", err)
	}

	privateKeyBytes := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})

	sshPrivateKey, err := ssh.ParseRawPrivateKey(privateKeyBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %w", err)
	}
	signer, err := ssh.NewSignerFromKey(sshPrivateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to create signer: %w", err)
	}
	return signer, nil
}

// Get ssh client config for our connection
// SSH config will use 2 authentication strategies: by key and by password.
func makeSSHConfig() (*ssh.ClientConfig, error) {
	key, err := generatePrivateKey()
	if err != nil {
		return nil, err
	}

	config := ssh.ClientConfig{
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(key),
		},
		//nolint: gosec
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	return &config, nil
}

// From https://sosedoff.com/2015/05/25/ssh-port-forwarding-with-go.html
// Handle local client connections and tunnel data to the remote server
// Will use io.Copy - http://golang.org/pkg/io/#Copy
func handleClient(client net.Conn, remote net.Conn) {
	defer func() {
		_ = client.Close()
	}()
	chDone := make(chan bool)

	// Start remote -> local data transfer
	go func() {
		_, err := io.Copy(client, remote)
		if err != nil {
			logger.Debugf("error while copy remote->local: %w", err)
		}
		chDone <- true
	}()

	// Start local -> remote data transfer
	go func() {
		_, err := io.Copy(remote, client)
		if err != nil {
			logger.Debugf("error while copy local->remote: %w", err)
		}
		chDone <- true
	}()

	<-chDone
}

func createTunnel(ctx context.Context, sshAddr string, localPort, remotePort int) (host string, err error) {
	errChan := make(chan error, 1)
	hostChan := make(chan string, 1)

	go func() {
		startReverseSSHTunnel(ctx, sshAddr, localPort, remotePort, errChan, hostChan)
	}()

	select {
	case <-ctx.Done():
		return "", fmt.Errorf("context canceled: %w", ctx.Err())
	case host := <-hostChan:
		// after this point, we'll need to consume and log the errors since we can't tell the client anything
		go func() {
			select {
			case <-ctx.Done():
				return
			case err := <-errChan:
				logger.Warn(err)
			}
		}()
		return host, nil
	case err := <-errChan:
		return "", fmt.Errorf("could not start proxy: %w", err)
	}
}

// appendToErrChan appends the error to the error channel if the context is not done.
func appendToErrChan(ctx context.Context, err error, errChan chan<- error) {
	select {
	case <-ctx.Done():
	case errChan <- err:
	}
}

// appendToHostChan adds the host to the host chan.
func appendToHostChan(ctx context.Context, host string, hostChan chan<- string) {
	select {
	case <-ctx.Done():
	case hostChan <- host:
	}
}

// nolint: cyclop
func startReverseSSHTunnel(ctx context.Context, sshAddr string, localPort, remotePort int, errChan chan<- error, hostChan chan<- string) {
	// Build SSH client configuration
	cfg, err := makeSSHConfig()
	if err != nil {
		appendToErrChan(ctx, fmt.Errorf("failed to make ssh config: %w", err), errChan)
		return
	}

	// Connect to SSH remote server using serverEndpoint
	serverConn, err := ssh.Dial("tcp", net.JoinHostPort(sshAddr, "22"), cfg)
	if err != nil {
		appendToErrChan(ctx, fmt.Errorf("failed to make ssh config: %w", err), errChan)
		return
	}

	// Listen on remote server port
	listener, err := serverConn.Listen("tcp", fmt.Sprintf("127.0.01:%d", remotePort))
	if err != nil {
		appendToErrChan(ctx, fmt.Errorf("listen open port ON remote server error: %w", err), errChan)
		return
	}

	defer func() {
		_ = listener.Close()
	}()

	// Create a new session
	session, err := serverConn.NewSession()
	if err != nil {
		appendToErrChan(ctx, fmt.Errorf("listen open port ON remote server error: %w", err), errChan)
		return
	}

	defer func() {
		_ = session.Close()
	}()

	session.Stderr = os.Stderr
	stdOut, err := session.StdoutPipe()
	if err != nil {
		appendToErrChan(ctx, fmt.Errorf("failed to make ssh config: %w", err), errChan)
		return
	}

	err = session.RequestPty("xterm", 80, 40, ssh.TerminalModes{
		ssh.ECHO:          1, // enable echoing
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	})
	if err != nil {
		appendToErrChan(ctx, fmt.Errorf("could not request pty: %w", err), errChan)
		return
	}

	err = session.Shell()
	if err != nil {
		appendToErrChan(ctx, fmt.Errorf("failed to start: %w", err), errChan)
		return
	}

	// wait for the banner
	consumeCtx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	var host string
	host, err = consumeBufferUntilURL(ctx, stdOut)
	if err != nil {
		appendToErrChan(consumeCtx, fmt.Errorf("could not request pty: %w", err), errChan)
		return
	}
	appendToHostChan(ctx, host, hostChan)

	// handle incoming connections on reverse forwarded tunnel
	for {
		// Open a (local) connection to localEndpoint whose content will be forwarded to serverEndpoint
		local, err := net.Dial("tcp", fmt.Sprintf(":%d", localPort))
		if err != nil {
			appendToErrChan(ctx, fmt.Errorf("dial INTO local service error: %w", err), errChan)
			return
		}

		client, err := listener.Accept()
		if err != nil {
			appendToErrChan(ctx, fmt.Errorf("could not accept connection: %w", err), errChan)
			return
		}

		handleClient(client, local)
	}
}

var re = regexp.MustCompile(`https?://[^\s]+`)

// nolint: nestif
func consumeBufferUntilURL(ctx context.Context, reader io.Reader) (host string, err error) {
	scanner := bufio.NewScanner(reader)
	var buffer string
	for {
		select {
		case <-ctx.Done():
			return "", fmt.Errorf("context canceled before string was found")
		default:
			if scanner.Scan() {
				buffer += scanner.Text()
				if strings.Contains(buffer, moeServer) {
					matches := re.FindStringSubmatch(buffer)
					if len(matches) == 0 {
						return "", fmt.Errorf("could not parse %s from %s", re.String(), buffer)
					}
					return matches[0], nil
				}
			} else {
				if err := scanner.Err(); err != nil {
					return "", fmt.Errorf("could not scan: %w", err)
				}
				return "", fmt.Errorf("EOF before string was found")
			}
		}
	}
}
