package iris

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

const (
	// KeySubnet corresponds to the associated resource schema key
	KeySubnet = "subnet"
	// KeyMAC corresponds to the associated resource schema key
	KeyMAC = "mac"
	// KeyIP corresponds to the associated resource schema key
	KeyIP = "ipaddr"
	// KeyName corresponds to the associated resource schema key
	KeyName = "hostname"
)

func resourceIrisDHCPReservation() *schema.Resource {
	return &schema.Resource{
		Create: resourceDhcpReservationCreate,
		Read:   resourceDhcpReservationRead,
		Update: resourceDhcpReservationUpdate,
		Delete: resourceDhcpReservationDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			KeySubnet: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.All(validation.StringIsNotEmpty, validation.StringIsNotWhiteSpace),
			},
			KeyMAC: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.IsMACAddress,
			},
			KeyIP: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.IsIPAddress,
			},
			KeyName: {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.All(validation.StringIsNotEmpty, validation.StringIsNotWhiteSpace),
			},
		},
	}
}

func resourceDhcpReservationCreate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceDhcpReservationRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceDhcpReservationDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceDhcpReservationUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}
