package iris

import (
	"fmt"
	"net/url"
	"sync"

	"github.com/dalet-oss/terraform-provider-iris/sdk"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

const (
	// MimeJSON is JSON MIME-type representation
	MimeJSON = "application/json"

	// KeyIrisProviderURI is the full URI to Iris API server
	KeyIrisProviderURI = "uri"
	// KeyIrisProviderToken is the API key to authenticate with
	KeyIrisProviderToken = "token"
)

// ProviderConfiguration struct for iris-provider
type ProviderConfiguration struct {
	Iris  *sdk.Iris
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
				ValidateFunc: validation.IsURLWithHTTPorHTTPS,
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
			"iris_dhcp_subnet":      resourceIrisDHCPSubnet(),
			"iris_dhcp_reservation": resourceIrisDHCPReservation(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func newIrisClient(uri, token string) (*sdk.Iris, error) {
	if uri == "" || token == "" {
		return nil, fmt.Errorf("The Iris provider needs proper initialization parameters")
	}

	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	r := httptransport.New(u.Host, sdk.DefaultBasePath, []string{u.Scheme})
	r.SetDebug(false)
	r.Consumers[MimeJSON] = runtime.JSONConsumer()
	r.Producers[MimeJSON] = runtime.JSONProducer()
	auths := []runtime.ClientAuthInfoWriter{
		httptransport.APIKeyAuth("x-token", "header", token),
	}
	r.DefaultAuthentication = httptransport.Compose(auths...)

	return sdk.New(r, strfmt.Default), nil
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {

	// check for mandatory requirements
	uri := d.Get(KeyIrisProviderURI).(string)
	token := d.Get(KeyIrisProviderToken).(string)

	iris, err := newIrisClient(uri, token)
	if err != nil {
		return nil, err
	}

	var mut sync.Mutex
	var provider = ProviderConfiguration{
		Iris:  iris,
		Mutex: &mut,
		Cond:  sync.NewCond(&mut),
	}

	return &provider, nil
}
