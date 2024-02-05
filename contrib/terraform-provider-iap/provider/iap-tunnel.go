package provider

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	provider_diag "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-google/google/transport"
	"github.com/phayes/freeport"
	"github.com/synapsecns/sanguine/contrib/tfcore/generated/tunnel"
	"github.com/synapsecns/sanguine/contrib/tfcore/utils"
	"log"
	"net/http"
	"net/url"
	"time"
)

// dataSourceProxyURL generates a proxy over an iap bastion host.
func dataSourceProxyURL() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceProxy,
		CreateContext: func(ctx context.Context, data *schema.ResourceData, i interface{}) provider_diag.Diagnostics {
			err := dataSourceProxy(data, i)
			if err != nil {
				return provider_diag.FromErr(err)
			}
			return provider_diag.Diagnostics{}
		},
		Delete: dataSourceProxyDelete,

		Schema: map[string]*schema.Schema{
			// project of the bastion host
			"project": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			// zone of the bastion host
			"zone": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			// zone of the bastion host
			"instance": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			// network interface to use
			"interface": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			// port of the host to connect to
			"remote_port": {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validatePort,
				ForceNew:     true,
			},
			// output proxy url
			"proxy_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceProxyDelete(d *schema.ResourceData, meta interface{}) error {
	// Delete the proxy URL.
	// This could involve making a call to an API to delete the proxy, or just
	// cleaning up any resources created on your end.
	// ...

	// Remove all fields in the dataSourceProxyURL resource
	d.SetId("")
	for k := range dataSourceProxyURL().Schema {
		_ = d.Set(k, nil)
	}
	return nil
}

// nolint: cyclop
func dataSourceProxy(d *schema.ResourceData, meta interface{}) error {
	config, ok := meta.(*transport.Config)
	if !ok {
		return fmt.Errorf("could not cast config of type %T to %T", meta, config)
	}

	project, ok := d.Get("project").(string)
	if !ok {
		return fmt.Errorf("could not cast project of type %T to %T", d.Get("project"), project)
	}
	zone, ok := d.Get("zone").(string)
	if !ok {
		return fmt.Errorf("could not cast zone of type %T to %T", d.Get("zone"), zone)
	}
	instance, ok := d.Get("instance").(string)
	if !ok {
		return fmt.Errorf("could not cast instance of type %T to %T", d.Get("instance"), instance)
	}
	iface, ok := d.Get("interface").(string)
	if !ok {
		return fmt.Errorf("could not cast interface of type %T to %T", d.Get("interface"), iface)
	}
	remotePort, ok := d.Get("remote_port").(int)
	if !ok {
		return fmt.Errorf("could not cast remote_port of type %T to %T", d.Get("remote_port"), remotePort)
	}

	localPort, err := freeport.GetFreePort()
	if err != nil {
		return fmt.Errorf("could not get a free port: %w", err)
	}

	tm := tunnel.TunnelManager{
		Project:    project,
		RemotePort: remotePort,
		LocalPort:  localPort,
		Zone:       zone,
		Instance:   instance,
		Interface:  iface,
	}

	tm.SetTokenSource(utils.GetTokenSource(config))

	errChan := make(chan error)

	log.Printf("[INFO] creating tunnel")
	go func() {
		startTime := time.Now()
		err := tm.StartProxy(context.Background())
		if err != nil {
			fmt.Println(err)
			log.Printf("[DEBUG] Proxy Error %v", err)
			errChan <- err
		}

		log.Printf("[DEBUG] Proxy closed after %s", time.Since(startTime))
	}()

	select {
	// wait 5 seconds for an error, otherwise just log since this will run in the background for the course of the apply
	case <-time.NewTimer(time.Second * 5).C:
		break
	case err := <-errChan:
		log.Printf("[ERROR] Received error while booting provider: %v", err)
		return err
	}

	log.Printf("[DEBUG] Finished creating proxy on port %d", localPort)

	id := uuid.New().String()
	log.Printf("[DEBUG] setting proxy id to %s", id)
	d.SetId(id)

	proxyURL := fmt.Sprintf("http://localhost:%d", localPort)
	log.Printf("[DEBUG] setting proxy url to %s", proxyURL)
	err = d.Set("proxy_url", proxyURL)
	if err != nil {
		return fmt.Errorf("could not set proxy_url: %w", err)
	}

	// test the tunnel
	parsedURL, err := url.Parse(proxyURL)
	if err != nil {
		log.Printf("[ERROR] could not parse proxy url %s: %v", proxyURL, err)
		return fmt.Errorf("could not parse url: %w", err)
	}
	testClient := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(parsedURL)}}
	//nolint: noctx
	resp, err := testClient.Get("https://www.google.com/")
	if err != nil {
		log.Printf("[ERROR] could not connect through proxy %s: %v", proxyURL, err)
	}

	_ = resp.Body.Close()

	return nil
}
