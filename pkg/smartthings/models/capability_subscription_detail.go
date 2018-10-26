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

// CapabilitySubscriptionDetail Details of a subscription of source type CAPABILITY. The combination of subscribed values must be unique per installed app.
// swagger:model CapabilitySubscriptionDetail
type CapabilitySubscriptionDetail struct {

	// Name of the capabilities attribute or * for all.
	// Max Length: 256
	// Min Length: 1
	Attribute string `json:"attribute,omitempty"`

	// Name of the capability that is subscribed to.
	// Required: true
	// Max Length: 128
	// Min Length: 1
	Capability *string `json:"capability"`

	// The id of the location that both the app and source device are in.
	// Required: true
	LocationID *string `json:"locationId"`

	// Only execute the subscription if the subscribed event is a state change from previous events.
	StateChangeOnly *bool `json:"stateChangeOnly,omitempty"`

	// A name for the subscription that will be passed to the installed app. Must be unique per installed app.
	SubscriptionName string `json:"subscriptionName,omitempty"`

	// A particular value for the attribute that will trigger the subscription or * for all.
	// Max Length: 256
	Value interface{} `json:"value,omitempty"`
}

// Validate validates this capability subscription detail
func (m *CapabilitySubscriptionDetail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttribute(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCapability(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocationID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CapabilitySubscriptionDetail) validateAttribute(formats strfmt.Registry) error {

	if swag.IsZero(m.Attribute) { // not required
		return nil
	}

	if err := validate.MinLength("attribute", "body", string(m.Attribute), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("attribute", "body", string(m.Attribute), 256); err != nil {
		return err
	}

	return nil
}

func (m *CapabilitySubscriptionDetail) validateCapability(formats strfmt.Registry) error {

	if err := validate.Required("capability", "body", m.Capability); err != nil {
		return err
	}

	if err := validate.MinLength("capability", "body", string(*m.Capability), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("capability", "body", string(*m.Capability), 128); err != nil {
		return err
	}

	return nil
}

func (m *CapabilitySubscriptionDetail) validateLocationID(formats strfmt.Registry) error {

	if err := validate.Required("locationId", "body", m.LocationID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CapabilitySubscriptionDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CapabilitySubscriptionDetail) UnmarshalBinary(b []byte) error {
	var res CapabilitySubscriptionDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
