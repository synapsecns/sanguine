// Code copied from github.com/gartnera/gcloud/compute/iap:/tunnel.go for testing by synapse modulecopier DO NOT EDIT."

package tunnel

import (
	"bytes"
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"sync"

	"github.com/gartnera/gcloud/auth"
	"github.com/gorilla/websocket"
	"golang.org/x/oauth2"
	"golang.org/x/sync/errgroup"
)

const URL_SCHEME = "wss"
const URL_HOST = "tunnel.cloudproxy.app"
const MTLS_URL_HOST = "mtls.tunnel.cloudproxy.app"
const URL_PATH_ROOT = "/v4"
const CONNECT_ENDPOINT = "connect"
const RECONNECT_ENDPOINT = "reconnect"
const SEC_PROTOCOL_SUFFIX = "bearer.relay.tunnel.cloudproxy.app"
const TUNNEL_CLOUDPROXY_ORIGIN = "bot:iap-tunneler"

const SUBPROTOCOL_NAME = "relay.tunnel.cloudproxy.app"
const SUBPROTOCOL_TAG_LEN = 2
const SUBPROTOCOL_HEADER_LEN = SUBPROTOCOL_TAG_LEN + 4
const SUBPROTOCOL_MAX_DATA_FRAME_SIZE = 16384
const SUBPROTOCOL_TAG_CONNECT_SUCCESS_SID uint16 = 0x0001
const SUBPROTOCOL_TAG_RECONNECT_SUCCESS_ACK uint16 = 0x0002
const SUBPROTOCOL_TAG_DATA uint16 = 0x0004
const SUBPROTOCOL_TAG_ACK uint16 = 0x0007

// tunnelAdapter abstracts the iap websocket tunnel to an io.ReadWriteCloser
type tunnelAdapter struct {
	conn    *websocket.Conn
	inbound chan []byte
	acks    chan uint64

	outboundLock sync.Mutex

	totalInboundLen uint64
}

func newTunnelAdapter(conn *websocket.Conn) *tunnelAdapter {
	a := &tunnelAdapter{
		inbound: make(chan []byte),
		acks:    make(chan uint64),
		conn:    conn,
	}
	return a
}

func (a *tunnelAdapter) inboundAck(len uint64) error {
	buf := new(bytes.Buffer)
	_ = binary.Write(buf, binary.BigEndian, SUBPROTOCOL_TAG_ACK)
	_ = binary.Write(buf, binary.BigEndian, len)

	a.outboundLock.Lock()
	defer a.outboundLock.Unlock()
	err := a.conn.WriteMessage(websocket.BinaryMessage, buf.Bytes())
	if err != nil {
		return fmt.Errorf("unable to write inbound ack msg: %w", err)
	}
	return nil
}

func (a *tunnelAdapter) inboundHandler(ctx context.Context) error {
	for {
		_, msg, err := a.conn.ReadMessage()
		if err != nil {
			return fmt.Errorf("error while reading message: %w", err)
		}
		subprotocolTag := binary.BigEndian.Uint16(msg[:SUBPROTOCOL_TAG_LEN])
		msg = msg[SUBPROTOCOL_TAG_LEN:]
		if subprotocolTag == SUBPROTOCOL_TAG_CONNECT_SUCCESS_SID {
			continue
		} else if subprotocolTag == SUBPROTOCOL_TAG_ACK {
			// ack := binary.BigEndian.Uint64(msg)
			continue
		} else if subprotocolTag == SUBPROTOCOL_TAG_DATA {
			dataLen := binary.BigEndian.Uint32(msg[:4])
			msg = msg[4 : dataLen+4]
			a.inbound <- msg
			a.totalInboundLen += uint64(len(msg))
			err = a.inboundAck(a.totalInboundLen)
			if err != nil {
				fmt.Println("inbound ack err: %w", err)
			}
		} else {
			return errors.New("unknown tag")
		}
	}
}

func (a *tunnelAdapter) Read(p []byte) (int, error) {
	msg := <-a.inbound
	len := copy(p, msg)
	return len, nil
}

func (a *tunnelAdapter) Write(p []byte) (int, error) {
	for i := 0; i < len(p); i += SUBPROTOCOL_MAX_DATA_FRAME_SIZE {
		maxOrEnd := len(p)
		if maxOrEnd-i > SUBPROTOCOL_MAX_DATA_FRAME_SIZE {
			maxOrEnd = i + SUBPROTOCOL_MAX_DATA_FRAME_SIZE
		}
		currentLen := maxOrEnd - i
		buf := new(bytes.Buffer)
		_ = binary.Write(buf, binary.BigEndian, SUBPROTOCOL_TAG_DATA)
		_ = binary.Write(buf, binary.BigEndian, uint32(currentLen))
		_, _ = buf.Write(p[i:maxOrEnd])

		a.outboundLock.Lock()
		err := a.conn.WriteMessage(websocket.BinaryMessage, buf.Bytes())
		a.outboundLock.Unlock()
		if err != nil {
			return 0, fmt.Errorf("unable to write to websocket: %w", err)
		}
	}
	return len(p), nil
}

func (a *tunnelAdapter) Close() error {
	return nil
}

func (a *tunnelAdapter) Start(ctx context.Context) {
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		err := a.inboundHandler(ctx)
		if err != nil {
			fmt.Println("error in inboundHandler: %w", err)
		}
		return err
	})
	err := eg.Wait()
	if err != nil {
		fmt.Println("error in tunnel adapter: %w", err)
	}
}

type TunnelManager struct {
	Project    string
	RemotePort int
	LocalPort  int

	Zone      string
	Instance  string
	Interface string

	ts oauth2.TokenSource
}

func (m *TunnelManager) getHeaders() (http.Header, error) {
	tok, err := m.ts.Token()
	if err != nil {
		return nil, fmt.Errorf("unable to get token: %w", err)
	}
	return http.Header{
		"Origin":        []string{TUNNEL_CLOUDPROXY_ORIGIN},
		"Authorization": []string{fmt.Sprintf("Bearer %s", tok.AccessToken)},
	}, nil
}

func (m *TunnelManager) StartTunnel(ctx context.Context) (io.ReadWriteCloser, error) {
	var err error
	if m.ts == nil {
		m.ts, err = auth.TokenSource()
		if err != nil {
			return nil, fmt.Errorf("unable to get tokensource: %w", err)
		}
	}

	tunnelUrl := &url.URL{
		Scheme: URL_SCHEME,
		Host:   URL_HOST,
		Path:   fmt.Sprintf("%s/%s", URL_PATH_ROOT, CONNECT_ENDPOINT),
	}
	query := tunnelUrl.Query()
	query.Add("project", m.Project)
	query.Add("zone", m.Zone)
	query.Add("instance", m.Instance)
	query.Add("interface", "nic0")
	query.Add("port", fmt.Sprint(m.RemotePort))
	tunnelUrl.RawQuery = query.Encode()

	headers, err := m.getHeaders()
	if err != nil {
		return nil, fmt.Errorf("unable to get connect headers: %w", err)
	}

	urlStr := tunnelUrl.String()

	conn, _, err := websocket.DefaultDialer.Dial(urlStr, headers)
	if err != nil {
		return nil, fmt.Errorf("unable to connect: %w", err)
	}
	adapter := newTunnelAdapter(conn)
	go adapter.Start(ctx)
	return adapter, nil
}

func (m *TunnelManager) StartProxy(ctx context.Context) error {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", m.LocalPort))
	if err != nil {
		return fmt.Errorf("unable to listen: %w", err)
	}
	for {
		conn, err := lis.Accept()
		if err != nil {
			return fmt.Errorf("unable to accept: %w", err)
		}
		tunnel, err := m.StartTunnel(ctx)
		if err != nil {
			return fmt.Errorf("unable to start tunnel: %w", err)
		}
		go func() {
			_, err := io.Copy(conn, tunnel)
			if err != nil {
				fmt.Println("unable to copy from tunnel to conn", err)
			}
		}()
		go func() {
			_, err := io.Copy(tunnel, conn)
			if err != nil {
				fmt.Println("unable to copy from conn to tunnel", err)
			}
		}()
	}
}
