// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DHCP d h c p
//
// swagger:model DHCP
type DHCP struct {

	// The status of the DHCPv4 service
	// Enum: [OK KO]
	Status string `json:"status,omitempty"`

	// The version of the DHCPv4 service
	Version string `json:"version,omitempty"`
}

// Validate validates this d h c p
func (m *DHCP) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var dHCPTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["OK","KO"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		dHCPTypeStatusPropEnum = append(dHCPTypeStatusPropEnum, v)
	}
}

const (

	// DHCPStatusOK captures enum value "OK"
	DHCPStatusOK string = "OK"

	// DHCPStatusKO captures enum value "KO"
	DHCPStatusKO string = "KO"
)

// prop value enum
func (m *DHCP) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, dHCPTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DHCP) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this d h c p based on context it is used
func (m *DHCP) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DHCP) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DHCP) UnmarshalBinary(b []byte) error {
	var res DHCP
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}