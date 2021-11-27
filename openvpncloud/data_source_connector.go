package openvpncloud

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/patoarvizu/terraform-provider-openvpn-cloud/client"
)

func dataSourceConnector() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceConnectorRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"network_item_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"network_item_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vpn_region_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip_v4_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip_v6_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceConnectorRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)
	var diags diag.Diagnostics
	connector, err := c.GetConnector(d.Get("name").(string))
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Error requesting connector",
			Detail:   fmt.Sprintf("Error requesting connector %v", err),
		})
		return diags
	}
	d.Set("name", connector.Name)
	d.Set("network_item_id", connector.NetworkItemId)
	d.Set("network_item_type", connector.NetworkItemType)
	d.Set("vpn_region_id", connector.VpnRegionId)
	d.Set("ip_v4_address", connector.IPv4Address)
	d.Set("ip_v6_address", connector.IPv6Address)
	d.SetId(connector.Id)
	return diags
}
