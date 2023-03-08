package iris

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	"github.com/dalet-oss/iris-api/client/dns"
	"github.com/dalet-oss/iris-api/models"
)

const (
	// KeyZone corresponds to the associated resource schema key
	KeyZone = "zone"
	// KeyRecord corresponds to the associated resource schema key
	KeyRecord = "record"
	// KeyType corresponds to the associated resource schema key
	KeyType = "type"
	// KeyTTL corresponds to the associated resource schema key
	KeyTTL = "ttl"
	// KeyValues corresponds to the associated resource schema key
	KeyValues = "values"
)

func resourceIrisDNSRecord() *schema.Resource {
	return &schema.Resource{
		Create: resourceDNSRecordCreate,
		Read:   resourceDNSRecordRead,
		Update: resourceDNSRecordUpdate,
		Delete: resourceDNSRecordDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			KeyZone: {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.All(validation.StringIsNotEmpty, validation.StringIsNotWhiteSpace),
			},
			KeyRecord: {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.All(validation.StringIsNotEmpty, validation.StringIsNotWhiteSpace),
			},
			KeyType: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.All(validation.StringIsNotEmpty, validation.StringIsNotWhiteSpace),
			},
			KeyTTL: {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validation.NoZeroValues,
			},

			KeyValues: {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validation.All(validation.StringIsNotEmpty, validation.StringIsNotWhiteSpace),
				},
			},
		},
	}
}

func dnsResourceID(zone, record string) string {
	return fmt.Sprintf("%s.%s", record, zone)
}

func parseDNSResourceID(resID string) (string, string, error) {
	if !strings.Contains(resID, ".") {
		return "", "", fmt.Errorf("invalid resource format: %s. must be record.zone", resID)
	}
	s := strings.Split(resID, ".")
	return s[0], strings.Join(s[1:], "."), nil
}

func newRecord(d *schema.ResourceData) models.Record {
	t := d.Get(KeyType).(string)
	rec := models.Record{
		ID:   d.Get(KeyRecord).(string),
		TTL:  int32(d.Get(KeyTTL).(int)),
		Type: &t,
	}
	for _, v := range d.Get(KeyValues).([]interface{}) {
		rec.Values = append(rec.Values, v.(string))
	}

	return rec
}

func recordToResource(zone string, r *models.Record, d *schema.ResourceData) {
	// set object params
	d.Set(KeyZone, zone)
	d.Set(KeyRecord, strings.Split(r.ID, ".")[0])
	d.Set(KeyType, *r.Type)
	d.Set(KeyTTL, r.TTL)

	var values []interface{}
	for _, v := range r.Values {
		values = append(values, v)
	}
	d.Set(KeyValues, values)
}

func resourceDNSRecordCreate(d *schema.ResourceData, meta interface{}) error {
	pconf := meta.(*ProviderConfiguration)

	pconf.Mutex.Lock()
	defer pconf.Mutex.Unlock()

	// create a new reservation
	rec := newRecord(d)
	zone := d.Get(KeyZone).(string)
	params := dns.NewCreateDNSZoneRecordParams().WithZoneID(zone).WithBody(&rec)
	r, err := pconf.Iris.DNS.CreateDNSZoneRecord(params, nil)
	if err != nil {
		return err
	}

	// set resource ID accordingly
	d.SetId(dnsResourceID(zone, r.Payload.ID))

	return err
}

func resourceDNSRecordRead(d *schema.ResourceData, meta interface{}) error {
	pconf := meta.(*ProviderConfiguration)

	pconf.Mutex.Lock()
	defer pconf.Mutex.Unlock()

	zone := d.Get(KeyZone).(string)
	record, _, err := parseDNSResourceID(d.Id())
	if err != nil {
		d.SetId("")
		return err
	}
	params := dns.NewGetDNSZoneRecordParams().WithZoneID(zone).WithRecordID(record)
	r, err := pconf.Iris.DNS.GetDNSZoneRecord(params, nil)
	if err != nil {
		return err
	}

	// set object params
	recordToResource(zone, r.Payload, d)

	return nil
}

func resourceDNSRecordDelete(d *schema.ResourceData, meta interface{}) error {
	pconf := meta.(*ProviderConfiguration)

	pconf.Mutex.Lock()
	defer pconf.Mutex.Unlock()

	zone := d.Get(KeyZone).(string)
	t := d.Get(KeyType).(string)
	record, _, err := parseDNSResourceID(d.Id())
	if err != nil {
		d.SetId("")
		return err
	}
	params := dns.NewDeleteDNSZoneRecordParams().WithZoneID(zone).WithRecordID(record).WithType(t)
	_, err = pconf.Iris.DNS.DeleteDNSZoneRecord(params, nil)
	if err != nil {
		return err
	}

	return nil
}

func resourceDNSRecordUpdate(d *schema.ResourceData, meta interface{}) error {
	pconf := meta.(*ProviderConfiguration)

	pconf.Mutex.Lock()
	defer pconf.Mutex.Unlock()

	// update an existing record
	rec := newRecord(d)
	zone := d.Get(KeyZone).(string)
	record, _, err := parseDNSResourceID(d.Id())
	if err != nil {
		d.SetId("")
		return err
	}

	params := dns.NewUpdateDNSZoneRecordParams().WithZoneID(zone).WithRecordID(record).WithBody(&rec)
	_, err = pconf.Iris.DNS.UpdateDNSZoneRecord(params, nil)
	if err != nil {
		return err
	}

	// set resource ID accordingly (ID may have been updated)
	d.SetId(dnsResourceID(zone, record))

	return nil
}
