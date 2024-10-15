// main.go
package main

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return &schema.Provider{
				Schema: map[string]*schema.Schema{},
				ResourcesMap: map[string]*schema.Resource{
					"october": resourceOctober(),
				},
				DataSourcesMap: map[string]*schema.Resource{},
			}
		},
	})
}
func resourceOctober() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceOctoberCreate,
		ReadContext:   resourceOctoberRead,
		UpdateContext: resourceOctoberUpdate,
		DeleteContext: resourceOctoberDelete,
		Schema: map[string]*schema.Schema{
			"value": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceOctoberCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	value := d.Get("value").(string)
	d.SetId("october_id") // set a unique ID for this resource instance
	d.Set("value", value)
	return nil
}

func resourceOctoberRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Placeholder for the read functionality
	return nil
}

func resourceOctoberUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	value := d.Get("value").(string)
	d.Set("value", value)
	return nil
}

func resourceOctoberDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	d.SetId("") // unset the ID to signify deletion
	return nil
}
