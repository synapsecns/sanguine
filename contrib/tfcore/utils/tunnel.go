package utils

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-google/google/transport"
	"github.com/phayes/freeport"
	"github.com/synapsecns/sanguine/contrib/tfcore/generated/tunnel"
	"log"
	"net/http"
	"net/url"
	"time"
)

// StartTunnel and returns the proxy url
// nolint: cyclop
func StartTunnel(ctx context.Context, d *schema.ResourceData, config *transport.Config) (proxyURL string, err error) {
	project, ok := d.Get("project").(string)
	if !ok {
		return proxyURL, fmt.Errorf("could not cast project of type %T to %T for project", d.Get("project"), project)
	}
	zone, ok := d.Get("zone").(string)
	if !ok {
		return proxyURL, fmt.Errorf("could not cast zone of type %T to %T for zone", d.Get("zone"), zone)
	}
	instance, ok := d.Get("instance").(string)
	if !ok {
		return proxyURL, fmt.Errorf("could not cast instance of type %T to %T for instance", d.Get("instance"), instance)
	}
	iface, ok := d.Get("interface").(string)
	if !ok {
		return proxyURL, fmt.Errorf("could not cast interface of type %T to %T for interface", d.Get("interface"), iface)
	}
	remotePort, ok := d.Get("remote_port").(int)
	if !ok {
		return proxyURL, fmt.Errorf("could not cast remote_port of type %T to %T for remote port", d.Get("remote_port"), remotePort)
	}

	localPort, err := freeport.GetFreePort()
	if err != nil {
		return proxyURL, fmt.Errorf("could not get a free port: %w", err)
	}

	tm := tunnel.TunnelManager{
		Project:    project,
		RemotePort: remotePort,
		LocalPort:  localPort,
		Zone:       zone,
		Instance:   instance,
		Interface:  iface,
	}

	tm.SetTokenSource(GetTokenSource(config))

	errChan := make(chan error)

	log.Printf("[INFO] creating tunnel")
	go func() {
		startTime := time.Now()
		err := tm.StartProxy(ctx)
		if err != nil {
			fmt.Println(err)
			log.Printf("[DEBUG] Proxy Error %v", err)
			errChan <- err
		}

		log.Printf("[DEBUG] Proxy closed after %s", time.Since(startTime))
	}()

	select {
	// wait 5 seconds for an error, otherwise just log since this will run in the background for the course of the apply
	case <-time.NewTimer(time.Second * 1).C:
		break
	case err := <-errChan:
		log.Printf("[ERROR] Received error while booting provider: %v", err)
		return proxyURL, fmt.Errorf("could not boot provider: %w", err)
	}

	log.Printf("[DEBUG] Finished creating proxy on port %d", localPort)

	// test the tunnel
	log.Printf("testing the tunnel")

	proxyURL = fmt.Sprintf("http://localhost:%d", localPort)
	log.Printf("[DEBUG] setting proxy url to %s", proxyURL)

	parsedURL, err := url.Parse(proxyURL)
	if err != nil {
		log.Printf("[ERROR] could not parse proxy url %s: %v", proxyURL, err)
		return proxyURL, fmt.Errorf("could not parse url: %w", err)
	}
	testClient := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(parsedURL)}}
	//nolint:noctx
	resp, err := testClient.Get("https://www.google.com/")
	if err != nil {
		log.Printf("[ERROR] could not connect through proxy %s: %v", proxyURL, err)
	} else {
		log.Printf("[INFO] proxy tunnel connected %s: %v", proxyURL, err)
		_ = resp.Body.Close()
	}

	return proxyURL, nil
}
