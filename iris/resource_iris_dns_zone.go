package iris

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	"github.com/dalet-oss/iris-api/client/dns"
	"github.com/dalet-oss/iris-api/models"
)

const (
	// KeyName corresponds to the associated resource schema key
	KeyName = "name"
)

func resourceIrisDNSZone() *schema.Resource {
	return &schema.Resource{
		Create: resourceDNSZoneCreate,
		Read:   resourceDNSZoneRead,
		Update: resourceDNSZoneUpdate,
		Delete: resourceDNSZoneDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			KeyName: {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.All(validation.StringIsNotEmpty, validation.StringIsNotWhiteSpace),
			},
		},
	}
}

func newZone(d *schema.ResourceData) models.Zone {
	name := d.Get(KeyName).(string)
	return models.Zone{
		Name: &name,
	}
}

func zoneToResource(z *models.Zone, d *schema.ResourceData) {
	// set object params
	d.Set(KeyName, *z.Name)
}

func resourceDNSZoneCreate(d *schema.ResourceData, meta interface{}) error {
	pconf := meta.(*ProviderConfiguration)

	pconf.Mutex.Lock()
	defer pconf.Mutex.Unlock()

	// create a new zone
	zone := newZone(d)
	params := dns.NewCreateDNSZoneParams().WithBody(&zone)
	z, err := pconf.Iris.DNS.CreateDNSZone(params, nil)
	if err != nil {
		return err
	}

	// set resource ID accordingly
	d.SetId(z.Payload.ID)

	return err
}

func resourceDNSZoneRead(d *schema.ResourceData, meta interface{}) error {
	pconf := meta.(*ProviderConfiguration)

	pconf.Mutex.Lock()
	defer pconf.Mutex.Unlock()

	params := dns.NewGetDNSZoneParams().WithZoneID(d.Id())
	z, err := pconf.Iris.DNS.GetDNSZone(params, nil)
	if err != nil {
		return err
	}

	// set object params
	zoneToResource(z.Payload, d)

	return nil
}

func resourceDNSZoneDelete(d *schema.ResourceData, meta interface{}) error {
	pconf := meta.(*ProviderConfiguration)

	pconf.Mutex.Lock()
	defer pconf.Mutex.Unlock()

	params := dns.NewDeleteDNSZoneParams().WithZoneID(d.Id())
	_, err := pconf.Iris.DNS.DeleteDNSZone(params, nil)
	if err != nil {
		return err
	}

	return nil
}

func resourceDNSZoneUpdate(d *schema.ResourceData, meta interface{}) error {
	pconf := meta.(*ProviderConfiguration)

	pconf.Mutex.Lock()
	defer pconf.Mutex.Unlock()

	// update an existing zone
	zone := newZone(d)
	params := dns.NewUpdateDNSZoneParams().WithZoneID(d.Id()).WithBody(&zone)
	_, err := pconf.Iris.DNS.UpdateDNSZone(params, nil)
	if err != nil {
		return err
	}

	return nil
}
