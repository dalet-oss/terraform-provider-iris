package iris

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	"github.com/dalet-oss/iris-api/client/dhcp"
	"github.com/dalet-oss/iris-api/models"
)

const (
	// KeyDomain corresponds to the associated resource schema key
	KeyDomain = "domain"
	// KeyHostname corresponds to the associated resource schema key
	KeyHostname = "hostname"
	// KeyIP corresponds to the associated resource schema key
	KeyIP = "ipaddr"
	// KeyMAC corresponds to the associated resource schema key
	KeyMAC = "mac"
	// KeySubnet corresponds to the associated resource schema key
	KeySubnet = "subnet"
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
			KeyDomain: {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.All(validation.StringIsNotEmpty, validation.StringIsNotWhiteSpace),
			},
			KeyHostname: {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.All(validation.StringIsNotEmpty, validation.StringIsNotWhiteSpace),
			},
			KeyIP: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.IsIPAddress,
			},
			KeyMAC: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.IsMACAddress,
			},
			KeySubnet: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.All(validation.StringIsNotEmpty, validation.StringIsNotWhiteSpace),
			},
		},
	}
}

func newReservation(d *schema.ResourceData) models.Reservation {
	res := models.Reservation{
		Domain:   d.Get(KeyDomain).(string),
		Hostname: d.Get(KeyHostname).(string),
		IP:       d.Get(KeyIP).(string),
		Mac:      d.Get(KeyMAC).(string),
	}

	return res
}

func reservationToResource(r *models.Reservation, d *schema.ResourceData) {
	// set object params
	d.Set(KeyDomain, r.Domain)
	d.Set(KeyHostname, r.Hostname)
	d.Set(KeyIP, r.IP)
	d.Set(KeyMAC, r.Mac)
}

func resourceDhcpReservationCreate(d *schema.ResourceData, meta interface{}) error {
	pconf := meta.(*ProviderConfiguration)

	pconf.Mutex.Lock()
	defer pconf.Mutex.Unlock()

	// create a new reservation
	res := newReservation(d)
	subnet := d.Get(KeySubnet).(string)
	params := dhcp.NewCreateDHCPSubnetReservationParams().WithSubnetID(subnet).WithBody(&res)
	s, err := pconf.Iris.Dhcp.CreateDHCPSubnetReservation(params, nil)
	if err != nil {
		return err
	}

	// set resource ID accordingly
	d.SetId(s.Payload.Mac)

	return err
}

func resourceDhcpReservationRead(d *schema.ResourceData, meta interface{}) error {
	pconf := meta.(*ProviderConfiguration)

	pconf.Mutex.Lock()
	defer pconf.Mutex.Unlock()

	subnet := d.Get(KeySubnet).(string)
	params := dhcp.NewGetDHCPSubnetReservationParams().WithSubnetID(subnet).WithMacID(d.Id())
	r, err := pconf.Iris.Dhcp.GetDHCPSubnetReservation(params, nil)
	if err != nil {
		return err
	}

	// set object params
	reservationToResource(r.Payload, d)

	return nil
}

func resourceDhcpReservationDelete(d *schema.ResourceData, meta interface{}) error {
	pconf := meta.(*ProviderConfiguration)

	pconf.Mutex.Lock()
	defer pconf.Mutex.Unlock()

	subnet := d.Get(KeySubnet).(string)
	params := dhcp.NewDeleteDHCPSubnetReservationParams().WithSubnetID(subnet).WithMacID(d.Id())
	_, err := pconf.Iris.Dhcp.DeleteDHCPSubnetReservation(params, nil)
	if err != nil {
		return err
	}

	return nil
}

func resourceDhcpReservationUpdate(d *schema.ResourceData, meta interface{}) error {
	pconf := meta.(*ProviderConfiguration)

	pconf.Mutex.Lock()
	defer pconf.Mutex.Unlock()

	// update an existing subnet
	res := newReservation(d)
	subnet := d.Get(KeySubnet).(string)
	params := dhcp.NewUpdateDHCPSubnetReservationParams().WithSubnetID(subnet).WithMacID(d.Id()).WithBody(&res)
	r, err := pconf.Iris.Dhcp.UpdateDHCPSubnetReservation(params, nil)
	if err != nil {
		return err
	}

	// set resource ID accordingly (MAC may have been updated)
	d.SetId(r.Payload.Mac)

	return nil
}
