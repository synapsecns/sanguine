package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/phayes/freeport"
	"github.com/synapsecns/sanguine/contrib/terraform-provider-iap/generated/google"
	"github.com/synapsecns/sanguine/contrib/terraform-provider-iap/generated/tunnel"
	"log"
	"time"
)

// dataSourceProxyURL generates a proxy over an iap bastion host.
func dataSourceProxyURL() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceProxy,

		Schema: map[string]*schema.Schema{
			// project of the bastion host
			"project": {
				Type:     schema.TypeString,
				Required: true,
			},
			// zone of the bastion host
			"zone": {
				Type:     schema.TypeString,
				Required: true,
			},
			// zone of the bastion host
			"instance": {
				Type:     schema.TypeString,
				Required: true,
			},
			// network interface to use
			"interface": {
				Type:     schema.TypeString,
				Required: true,
			},
			// port of the host to connect to
			"remote_port": {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validatePort,
			},
			// output proxy url
			"proxy_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceProxy(d *schema.ResourceData, meta interface{}) error {
	config, ok := meta.(*google.Config)
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

	tm.SetTokenSource(config.GetTokenSource())

	errChan := make(chan error)

	go func() {
		err := tm.StartProxy(context.Background())
		if err != nil {
			fmt.Println(err)
			log.Printf("[DEBUG] Proxy Error %v", err)
			errChan <- err
		}
	}()

	select {
	// wait 5 seconds for an error, otherwise just log since this will run in the background for the course of the apply
	case <-time.NewTimer(time.Second * 5).C:
		break
	case err := <-errChan:
		return err
	}

	err = d.Set("proxy_url", fmt.Sprintf("http://localhost:%d", localPort))
	if err != nil {
		return fmt.Errorf("could not set proxy_url: %w", err)
	}

	return nil
}
