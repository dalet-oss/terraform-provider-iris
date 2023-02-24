package iris

import (
	"fmt"
	"sync"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

const (
	// KeyIrisProviderURI is the full URI to Iris API server
	KeyIrisProviderURI = "uri"
	// KeyIrisProviderToken is the API key to authenticate with
	KeyIrisProviderToken = "token"
)

// ProviderConfiguration struct for iris-provider
type ProviderConfiguration struct {
	Mutex *sync.Mutex
	Cond  *sync.Cond
}

// Provider iris
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			KeyIrisProviderURI: {
				Type:         schema.TypeString,
				Required:     true,
				DefaultFunc:  schema.EnvDefaultFunc("IRIS_URI", nil),
				ValidateFunc: validation.IsURLWithHTTPS,
				Description:  "Iris platform URI",
			},
			KeyIrisProviderToken: {
				Type:         schema.TypeString,
				Required:     true,
				DefaultFunc:  schema.EnvDefaultFunc("IRIS_TOKEN", nil),
				ValidateFunc: validation.All(validation.StringIsNotEmpty),
				Description:  "Iris platform token (API key)",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"iris_dhcp_reservation": resourceIrisDHCPReservation(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {

	// check for mandatory requirements
	uri := d.Get(KeyIrisProviderURI).(string)
	token := d.Get(KeyIrisProviderToken).(string)

	if uri == "" || token == "" {
		return nil, fmt.Errorf("The Iris provider needs proper initialization parameters")
	}

	var mut sync.Mutex
	var provider = ProviderConfiguration{
		Mutex: &mut,
		Cond:  sync.NewCond(&mut),
	}

	return &provider, nil
}
