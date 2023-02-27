package iris

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	"github.com/dalet-oss/terraform-provider-iris/models"
	"github.com/dalet-oss/terraform-provider-iris/sdk/dhcp"
)

const (
	// KeyCIDR corresponds to the associated resource schema key
	KeyCIDR = "cidr"
	// KeyPool corresponds to the associated resource schema key
	KeyPool = "pool"
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
			KeyCIDR: {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.IsCIDR,
			},
			KeyPool: {
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
			},
		},
	}
}

func newSubnet(d *schema.ResourceData) models.Subnet {
	cidr := d.Get(KeyCIDR).(string)
	subnet := models.Subnet{
		Subnet: &cidr,
	}

	for _, i := range d.Get(KeyPool).([]interface{}) {
		p := i.(map[string]interface{})
		pool := models.SubnetPoolRange{
			First: p[KeyFirst].(string),
			Last:  p[KeyLast].(string),
		}
		subnet.Pools = append(subnet.Pools, &pool)
	}

	return subnet
}

func subnetToResource(s *models.Subnet, d *schema.ResourceData) {
	// set object params
	d.Set(KeySubnet, s.Subnet)

	var pools []interface{}
	for _, i := range s.Pools {
		var p schema.ResourceData
		p.Set(KeyFirst, i.First)
		p.Set(KeyLast, i.Last)
		pools = append(pools, &p)
	}
	d.Set(KeyPool, pools)
}

func resourceDhcpSubnetCreate(d *schema.ResourceData, meta interface{}) error {
	pconf := meta.(*ProviderConfiguration)

	pconf.Mutex.Lock()
	defer pconf.Mutex.Unlock()

	// create a new subnet
	subnet := newSubnet(d)
	params := dhcp.NewCreateDHCPSubnetParams().WithBody(&subnet)
	s, err := pconf.Iris.Dhcp.CreateDHCPSubnet(params, nil)
	if err != nil {
		return err
	}

	// set resource ID accordingly
	d.SetId(s.Payload.ID)

	return err
}

func resourceDhcpSubnetRead(d *schema.ResourceData, meta interface{}) error {
	pconf := meta.(*ProviderConfiguration)

	pconf.Mutex.Lock()
	defer pconf.Mutex.Unlock()

	params := dhcp.NewGetDHCPSubnetParams().WithSubnetID(d.Id())
	s, err := pconf.Iris.Dhcp.GetDHCPSubnet(params, nil)
	if err != nil {
		return err
	}

	// set object params
	subnetToResource(s.Payload, d)

	return nil
}

func resourceDhcpSubnetDelete(d *schema.ResourceData, meta interface{}) error {
	pconf := meta.(*ProviderConfiguration)

	pconf.Mutex.Lock()
	defer pconf.Mutex.Unlock()

	params := dhcp.NewDeleteDHCPSubnetParams().WithSubnetID(d.Id())
	_, err := pconf.Iris.Dhcp.DeleteDHCPSubnet(params, nil)
	if err != nil {
		return err
	}

	return nil
}

func resourceDhcpSubnetUpdate(d *schema.ResourceData, meta interface{}) error {
	pconf := meta.(*ProviderConfiguration)

	pconf.Mutex.Lock()
	defer pconf.Mutex.Unlock()

	// update an existing subnet
	subnet := newSubnet(d)
	params := dhcp.NewUpdateDHCPSubnetParams().WithSubnetID(d.Id()).WithBody(&subnet)
	_, err := pconf.Iris.Dhcp.UpdateDHCPSubnet(params, nil)
	if err != nil {
		return err
	}

	return nil
}
