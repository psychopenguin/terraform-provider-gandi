package gandi

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceDomain() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
		Read: dataSourceDomainRead,
	}
}

func dataSourceDomainRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*GandiClients).Domain
	name := d.Get("name").(string)
	found, err := client.GetDomain(name)
	if err != nil {
		return fmt.Errorf("Unknown domain with name: '#{name}'")
	}
	d.SetId(found.ID)
	d.Set("name", found.FQDN)
	return nil
}
