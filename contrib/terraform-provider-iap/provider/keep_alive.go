package provider

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/synapsecns/sanguine/contrib/terraform-provider-iap/generated/google"
	"log"
	"time"
)

// keepAlive is a resource that keeps a tunnel alive
// by delaying the read of the datasource until the timeout is finished
func keepAlive() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceKeepAlive,
		Schema: map[string]*schema.Schema{
			// timeout in seconds
			"timeout": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"timed_out": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func dataSourceKeepAlive(d *schema.ResourceData, meta interface{}) error {
	config, ok := meta.(*google.Config)
	if !ok {
		return fmt.Errorf("could not cast config of type %T to %T", meta, config)
	}

	timeout, ok := d.Get("timeout").(int)
	if !ok {
		return fmt.Errorf("could not cast timeout of type %T to %T", d.Get("timeout"), timeout)
	}

	id := uuid.New().String()
	log.Printf("[DEBUG] setting proxy id to %s", id)
	d.SetId(id)

	log.Printf("[INFO] waiting for %d seconds", timeout)
	select {
	case <-time.After(time.Duration(timeout) * time.Second):
		log.Printf("[INFO] finished waiting %d seconds", timeout)
		err := d.Set("timed_out", true)
		if err != nil {
			return fmt.Errorf("could not set timed_out to true: %w", err)
		}
		break
	case <-config.GetContext().Done():
		log.Printf("[ERROR] contet cancelled before timeout (%d seconds)", timeout)
		return fmt.Errorf("context was cancelled")
	}
	return nil
}
