package spectrocloud

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/spectrocloud/terraform-provider-spectrocloud/pkg/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceRegistry() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceRegistryRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func dataSourceRegistryRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.V1Client)
	var diags diag.Diagnostics
	if v, ok := d.GetOk("name"); ok {
		registry, err := c.GetPackRegistryCommonByName(v.(string))
		if err != nil {
			return diag.FromErr(err)
		}
		d.SetId(registry.UID)
		d.Set("name", registry.Name)
	}
	return diags
}