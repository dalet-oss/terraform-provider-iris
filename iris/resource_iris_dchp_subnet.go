package iris

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

const (
	// KeyID corresponds to the associated resource schema key
	KeyID = "id"
	// KeyCIDR corresponds to the associated resource schema key
	KeyCIDR = "cidr"
	// KeyPools corresponds to the associated resource schema key
	KeyPools = "pools"
	// KeyFirst corresponds to the associated resource schema key
	KeyFirst = "first"
	// KeyLast corresponds to the associated resource schema key
	KeyLast = "last"
)

func resourceIrisDHCPSubnet() *schema.Resource {
	return &schema.Resource{
		Create: resourceDhcpSubnetCreate,
		Read:   resourceDhcpSubnetRead,
		Update: resourceDhcpSubnetUpdate,
		Delete: resourceDhcpSubnetDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			KeyID: {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.All(validation.StringIsNotEmpty, validation.StringIsNotWhiteSpace),
			},
			KeyCIDR: {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.IsCIDR,
			},
			KeyPools: {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						KeyFirst: {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.IsIPAddress,
						},
						KeyLast: {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.IsIPAddress,
						},
					},
				},
				ValidateFunc: validation.All(validation.StringIsNotEmpty, validation.StringIsNotWhiteSpace),
			},
		},
	}
}

func resourceDhcpSubnetCreate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceDhcpSubnetRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceDhcpSubnetDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceDhcpSubnetUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}
