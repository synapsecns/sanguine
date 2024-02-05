package provider

import (
	context2 "context"
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-google/v4/google"
	"github.com/synapsecns/sanguine/contrib/tfcore/utils"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"time"
)

// keepAlive is a resource that keeps a tunnel alive
// by delaying the read of the datasource until the timeout is finished.
func keepAlive() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceKeepAlive,
		Schema: map[string]*schema.Schema{
			// timeout in seconds
			"timeout": {
				Type:     schema.TypeInt,
				Required: true,
			},
			// port of the host to connect to
			"proxy_url": {
				Type:     schema.TypeString,
				Required: true,
			},
			// wether or not the keep alive has timed out
			"timed_out": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

// nolint: cyclop
func dataSourceKeepAlive(d *schema.ResourceData, meta interface{}) error {
	config, ok := meta.(*google.Config)
	if !ok {
		return fmt.Errorf("could not cast config of type %T to %T", meta, config)
	}

	timeout, ok := d.Get("timeout").(int)
	if !ok {
		return fmt.Errorf("could not cast timeout of type %T to %T", d.Get("timeout"), timeout)
	}

	proxyURL, ok := d.Get("proxy_url").(string)
	if !ok {
		return fmt.Errorf("could not cast remote_port of type %T to %T", d.Get("proxy_url"), proxyURL)
	}

	// test the tunnel
	parsedURL, err := url.Parse(proxyURL)
	if err != nil {
		log.Printf("[ERROR] could not parse proxy url %s: %v", proxyURL, err)
		return fmt.Errorf("could not parse proxy url %s: %w", proxyURL, err)
	}

	id := uuid.New().String()
	log.Printf("[DEBUG] setting proxy id to %s", id)
	d.SetId(id)

	log.Printf("[INFO] waiting for %d seconds", timeout)

	timer := time.After(time.Duration(timeout) * time.Second)

	cfgReflection := reflect.ValueOf(config)
	context := utils.GetUnexportedField(cfgReflection.FieldByName("context")).(context2.Context)

	for {
		select {
		case <-timer:
			log.Printf("[INFO] finished waiting %d seconds", timeout)
			err := d.Set("timed_out", true)
			if err != nil {
				return fmt.Errorf("could not set timed_out to true: %w", err)
			}
			return nil
		case <-time.After(time.Second * 5):
			log.Printf("[INFO] testing proxy %s", proxyURL)
			testClient := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(parsedURL)}}
			//nolint: noctx
			resp, err := testClient.Get("https://www.google.com/")
			if err != nil {
				log.Printf("[ERROR] could not connect through proxy %s: %v", proxyURL, err)
			}
			_ = resp.Body.Close()
			log.Printf("[INFO] successfully connected through proxy %s", proxyURL)
			continue
		case <-context.Done():
			log.Printf("[ERROR] contet canceled before timeout (%d seconds)", timeout)
			return fmt.Errorf("context was canceled")
		}
	}
}
