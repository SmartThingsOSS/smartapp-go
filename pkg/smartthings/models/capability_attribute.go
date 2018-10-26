// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// CapabilityAttribute capability attribute
// swagger:model CapabilityAttribute
type CapabilityAttribute struct {

	// [JSON schema](http://json-schema.org/specification-links.html#draft-4) for the attribute.
	//
	Schema interface{} `json:"schema,omitempty"`
}

// Validate validates this capability attribute
func (m *CapabilityAttribute) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CapabilityAttribute) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CapabilityAttribute) UnmarshalBinary(b []byte) error {
	var res CapabilityAttribute
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
