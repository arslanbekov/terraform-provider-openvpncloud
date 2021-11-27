package openvpncloud

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/patoarvizu/terraform-provider-openvpn-cloud/client"
)

func resourceDNSRecord() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceDNSRecordCreate,
		ReadContext:   resourceDNSRecordRead,
		DeleteContext: resourceDNSRecordDelete,
		UpdateContext: resourceDNSRecordUpdate,
		Schema: map[string]*schema.Schema{
			"domain": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ip_v4_addresses": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"ip_v6_addresses": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func resourceDNSRecordCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)
	var diags diag.Diagnostics
	domain := d.Get("domain").(string)
	ipV4Addresses := d.Get("ip_v4_addresses").([]interface{})
	ipV4AddressesSlice := make([]string, 0)
	for _, a := range ipV4Addresses {
		ipV4AddressesSlice = append(ipV4AddressesSlice, a.(string))
	}
	ipV6Addresses := d.Get("ip_v6_addresses").([]interface{})
	ipV6AddressesSlice := make([]string, 0)
	for _, a := range ipV6Addresses {
		ipV6AddressesSlice = append(ipV6AddressesSlice, a.(string))
	}
	dr := client.DNSRecord{
		Domain:        domain,
		IPV4Addresses: ipV4AddressesSlice,
		IPV6Addresses: ipV6AddressesSlice,
	}
	dnsRecord, err := c.CreateDNSRecord(dr)
	if err != nil {
		return append(diags, diag.FromErr(err)...)
	}
	d.SetId(dnsRecord.Id)
	return diags
}

func resourceDNSRecordRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)
	var diags diag.Diagnostics
	recordId := d.Id()
	r, err := c.GetDNSRecord(recordId)
	if err != nil {
		return append(diags, diag.FromErr(err)...)
	}
	d.Set("domain", r.Domain)
	d.Set("ip_v4_addresses", r.IPV4Addresses)
	d.Set("ip_v6_addresses", r.IPV6Addresses)
	return diags
}

func resourceDNSRecordUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)
	var diags diag.Diagnostics
	_, domain := d.GetChange("domain")
	_, ipV4Addresses := d.GetChange("ip_v4_addresses")
	ipV4AddressesSlice := make([]string, 0)
	for _, a := range ipV4Addresses.([]interface{}) {
		ipV4AddressesSlice = append(ipV4AddressesSlice, a.(string))
	}
	_, ipV6Addresses := d.GetChange("ip_v6_addresses")
	ipV6AddressesSlice := make([]string, 0)
	for _, a := range ipV6Addresses.([]interface{}) {
		ipV6AddressesSlice = append(ipV6AddressesSlice, a.(string))
	}
	dr := client.DNSRecord{
		Id:            d.Id(),
		Domain:        domain.(string),
		IPV4Addresses: ipV4AddressesSlice,
		IPV6Addresses: ipV6AddressesSlice,
	}
	err := c.UpdateDNSRecord(dr)
	if err != nil {
		return append(diags, diag.FromErr(err)...)
	}
	return diags
}

func resourceDNSRecordDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)
	var diags diag.Diagnostics
	routeId := d.Id()
	err := c.DeleteDNSRecord(routeId)
	if err != nil {
		return append(diags, diag.FromErr(err)...)
	}
	return diags
}
