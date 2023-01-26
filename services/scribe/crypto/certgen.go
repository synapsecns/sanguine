package crypto

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"fmt"
	"math/big"
	"os"
	"time"
)

func publicKey(priv interface{}) interface{} {
	switch k := priv.(type) {
	case *rsa.PrivateKey:
		return &k.PublicKey
	case *ecdsa.PrivateKey:
		return &k.PublicKey
	default:
		return nil
	}
}

func pemBlockForKey(priv interface{}) (*pem.Block, error) {
	switch k := priv.(type) {
	case *rsa.PrivateKey:
		return &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(k)}, nil
	case *ecdsa.PrivateKey:
		b, err := x509.MarshalECPrivateKey(k)
		if err != nil {
			return nil, fmt.Errorf("unable to marshal ECDSA private key: %w", err)
		}
		return &pem.Block{Type: "EC PRIVATE KEY", Bytes: b}, nil
	default:
		return nil, errors.New("unknown key type")
	}
}

// SelfSignedCertProvider is a self signed cert provider.
type SelfSignedCertProvider struct {
	// TODO: should be a getter
	Pool              *x509.CertPool
	pubKey, privKey   string
	CertFile, KeyFile string
}

func GetSelfSignedCert() (*SelfSignedCertProvider, error) {
	// priv, err := rsa.GenerateKey(rand.Reader, *rsaBits)
	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}
	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			Organization: []string{"Self Signed"},
		},
		NotBefore: time.Now(),
		NotAfter:  time.Now().Add(time.Hour * 24 * 180),

		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, publicKey(priv), priv)
	if err != nil {
		return nil, err
	}

	pubKey := &bytes.Buffer{}
	err = pem.Encode(pubKey, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	if err != nil {
		return nil, fmt.Errorf("could not encode certificate: %v", err)
	}

	pubKeyFile, err := createTempFile(pubKey)
	if err != nil {
		return nil, fmt.Errorf("could not create pub key file")
	}

	privKey := &bytes.Buffer{}
	pbKey, err := pemBlockForKey(priv)
	if err != nil {
		return nil, fmt.Errorf("could not create pem block for key: %v", err)
	}

	err = pem.Encode(privKey, pbKey)
	if err != nil {
		return nil, fmt.Errorf("could not encode private key: %v", err)
	}

	privKeyFile, err := createTempFile(privKey)
	if err != nil {
		return nil, fmt.Errorf("could not create pub key file")
	}

	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(derBytes)

	return &SelfSignedCertProvider{
		Pool:     pool,
		pubKey:   pubKey.String(),
		privKey:  privKey.String(),
		CertFile: pubKeyFile,
		KeyFile:  privKeyFile,
	}, nil
}

func createTempFile(toWrite *bytes.Buffer) (path string, err error) {
	keyFile, err := os.CreateTemp(os.TempDir(), "key")
	if err != nil {
		return "", fmt.Errorf("could not store cert: %w", err)
	}

	_, err = keyFile.Write(toWrite.Bytes())
	if err != nil {
		return "", fmt.Errorf("could not write cert: %w", err)
	}

	defer func() {
		_ = keyFile.Close()
	}()

	return keyFile.Name(), nil
}
