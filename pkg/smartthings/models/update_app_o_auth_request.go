// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UpdateAppOAuthRequest update app o auth request
// swagger:model UpdateAppOAuthRequest
type UpdateAppOAuthRequest struct {

	// A name given to the OAuth Client.
	// Required: true
	ClientName *string `json:"clientName"`

	// A list of SmartThings API OAuth scope identifiers that maybe required to execute your integration.
	// Required: true
	Scope []string `json:"scope"`
}

// Validate validates this update app o auth request
func (m *UpdateAppOAuthRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClientName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScope(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateAppOAuthRequest) validateClientName(formats strfmt.Registry) error {

	if err := validate.Required("clientName", "body", m.ClientName); err != nil {
		return err
	}

	return nil
}

func (m *UpdateAppOAuthRequest) validateScope(formats strfmt.Registry) error {

	if err := validate.Required("scope", "body", m.Scope); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateAppOAuthRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateAppOAuthRequest) UnmarshalBinary(b []byte) error {
	var res UpdateAppOAuthRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
