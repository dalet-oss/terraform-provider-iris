// Code generated by go-swagger; DO NOT EDIT.

package dhcp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// EnableDHCPReader is a Reader for the EnableDHCP structure.
type EnableDHCPReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EnableDHCPReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewEnableDHCPCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 409:
		result := NewEnableDHCPConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEnableDHCPInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewEnableDHCPCreated creates a EnableDHCPCreated with default headers values
func NewEnableDHCPCreated() *EnableDHCPCreated {
	return &EnableDHCPCreated{}
}

/*
EnableDHCPCreated describes a response with status code 201, with default header values.

The DHCPv4 service has been enabled.
*/
type EnableDHCPCreated struct {
}

// IsSuccess returns true when this enable d h c p created response has a 2xx status code
func (o *EnableDHCPCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this enable d h c p created response has a 3xx status code
func (o *EnableDHCPCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enable d h c p created response has a 4xx status code
func (o *EnableDHCPCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this enable d h c p created response has a 5xx status code
func (o *EnableDHCPCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this enable d h c p created response a status code equal to that given
func (o *EnableDHCPCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the enable d h c p created response
func (o *EnableDHCPCreated) Code() int {
	return 201
}

func (o *EnableDHCPCreated) Error() string {
	return fmt.Sprintf("[POST /dhcp/enable][%d] enableDHCPCreated ", 201)
}

func (o *EnableDHCPCreated) String() string {
	return fmt.Sprintf("[POST /dhcp/enable][%d] enableDHCPCreated ", 201)
}

func (o *EnableDHCPCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEnableDHCPConflict creates a EnableDHCPConflict with default headers values
func NewEnableDHCPConflict() *EnableDHCPConflict {
	return &EnableDHCPConflict{}
}

/*
EnableDHCPConflict describes a response with status code 409, with default header values.

The DHCPv4 service was already enable.
*/
type EnableDHCPConflict struct {
}

// IsSuccess returns true when this enable d h c p conflict response has a 2xx status code
func (o *EnableDHCPConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this enable d h c p conflict response has a 3xx status code
func (o *EnableDHCPConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enable d h c p conflict response has a 4xx status code
func (o *EnableDHCPConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this enable d h c p conflict response has a 5xx status code
func (o *EnableDHCPConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this enable d h c p conflict response a status code equal to that given
func (o *EnableDHCPConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the enable d h c p conflict response
func (o *EnableDHCPConflict) Code() int {
	return 409
}

func (o *EnableDHCPConflict) Error() string {
	return fmt.Sprintf("[POST /dhcp/enable][%d] enableDHCPConflict ", 409)
}

func (o *EnableDHCPConflict) String() string {
	return fmt.Sprintf("[POST /dhcp/enable][%d] enableDHCPConflict ", 409)
}

func (o *EnableDHCPConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEnableDHCPInternalServerError creates a EnableDHCPInternalServerError with default headers values
func NewEnableDHCPInternalServerError() *EnableDHCPInternalServerError {
	return &EnableDHCPInternalServerError{}
}

/*
EnableDHCPInternalServerError describes a response with status code 500, with default header values.

Unable to enable the DHCPv4 service.
*/
type EnableDHCPInternalServerError struct {
}

// IsSuccess returns true when this enable d h c p internal server error response has a 2xx status code
func (o *EnableDHCPInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this enable d h c p internal server error response has a 3xx status code
func (o *EnableDHCPInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enable d h c p internal server error response has a 4xx status code
func (o *EnableDHCPInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this enable d h c p internal server error response has a 5xx status code
func (o *EnableDHCPInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this enable d h c p internal server error response a status code equal to that given
func (o *EnableDHCPInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the enable d h c p internal server error response
func (o *EnableDHCPInternalServerError) Code() int {
	return 500
}

func (o *EnableDHCPInternalServerError) Error() string {
	return fmt.Sprintf("[POST /dhcp/enable][%d] enableDHCPInternalServerError ", 500)
}

func (o *EnableDHCPInternalServerError) String() string {
	return fmt.Sprintf("[POST /dhcp/enable][%d] enableDHCPInternalServerError ", 500)
}

func (o *EnableDHCPInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}