package domeneshop

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func resourceRecord() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"domain_id": {
				Type:         schema.TypeInt,
				Required:     true,
				Description:  "ID of the domain",
				ForceNew:     true,
			},
			"host": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "A description of an item",
				
			},
			"ttl": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "An optional list of tags, represented as a key, value pair",
			},
		},
		Create: resourceCreateItem,
		Read:   resourceReadItem,
		Update: resourceUpdateItem,
		Delete: resourceDeleteItem,
		Exists: resourceExistsItem,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}