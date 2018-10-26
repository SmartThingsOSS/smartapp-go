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

// SecurityArmStateDetail Details of a subscription of souce type SECURITY_ARM_STATE
// swagger:model SecurityArmStateDetail
type SecurityArmStateDetail struct {

	// The id of the location that both the app and the security system are in.
	// Required: true
	LocationID *string `json:"locationId"`

	// A name for the subscription that will be passed to the installed app.
	SubscriptionName string `json:"subscriptionName,omitempty"`
}

// Validate validates this security arm state detail
func (m *SecurityArmStateDetail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLocationID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityArmStateDetail) validateLocationID(formats strfmt.Registry) error {

	if err := validate.Required("locationId", "body", m.LocationID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SecurityArmStateDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecurityArmStateDetail) UnmarshalBinary(b []byte) error {
	var res SecurityArmStateDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
