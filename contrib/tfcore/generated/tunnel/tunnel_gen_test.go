// Code copied from github.com/gartnera/gcloud/compute/iap:/tunnel_test.go for testing by synapse modulecopier DO NOT EDIT."

package tunnel

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/augustoroman/hexdump"
	"github.com/stretchr/testify/require"
)

func TestStartTunnel(t *testing.T) {
	if _, ok := os.LookupEnv("MANUAL_TEST"); !ok {
		t.Skip()
	}

	ctx := context.Background()

	m := TunnelManager{
		Project:    os.Getenv("GOOGLE_PROJECT_ID"),
		Zone:       "us-west1-c",
		Instance:   os.Getenv("TEST_INSTANCE_NAME"),
		Interface:  "nic0",
		RemotePort: 22,
	}
	tunnel, err := m.StartTunnel(ctx)
	require.NoError(t, err)
	buf := make([]byte, SUBPROTOCOL_MAX_DATA_FRAME_SIZE)
	i, err := tunnel.Read(buf)
	require.NoError(t, err)
	fmt.Println(i)
	fmt.Println(hexdump.Dump(buf[:i]))
	_, err = tunnel.Write([]byte("SSH-2.0-jsssh.0.1\n"))
	require.NoError(t, err)
	i, err = tunnel.Read(buf)
	require.NoError(t, err)
	fmt.Println(i)
	fmt.Println(hexdump.Dump(buf[:i]))
	time.Sleep(time.Second * 1)
}

func TestStartProxy(t *testing.T) {
	if _, ok := os.LookupEnv("MANUAL_TEST"); !ok {
		t.Skip()
	}

	ctx := context.Background()

	m := TunnelManager{
		Project:    os.Getenv("GOOGLE_PROJECT_ID"),
		Zone:       "us-west1-c",
		Instance:   os.Getenv("TEST_INSTANCE_NAME"),
		Interface:  "nic0",
		RemotePort: 22,
		LocalPort:  2020,
	}
	err := m.StartProxy(ctx)
	require.NoError(t, err)
}
