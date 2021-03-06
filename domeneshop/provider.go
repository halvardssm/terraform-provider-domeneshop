package domeneshop

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/halvardssm/go-domeneshop-client/client DomeneshopClient"
)

func Provider() terraform.ResourceProvider {
	p := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("DOMENESHOP_TOKEN", ""),
				Description: "The DomeneShop Token",
			},
			"secret": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("DOMENESHOP_SECRET", ""),
				Description: "The DomeneShop Secret",
			},
		},
		DataSourcesMap: map[string]*schema.Resource{},
		ResourcesMap:   map[string]*schema.Resource{},
	}

	p.ConfigureFunc = providerConfigure

	return p
}

func providerConfigure(d *schema.ResourceData) (DomeneshopClient{}, error) {
	token := d.Get("token").(string)
	secret := d.Get("secret").(string)

	return DomeneshopClient(token, secret), nil
}
