// Code generated by go-swagger; DO NOT EDIT.

package smartapp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// DeviceConfig A device configuration.
// swagger:model DeviceConfig
type DeviceConfig struct {

	// The component ID on the device.
	ComponentID string `json:"componentId,omitempty"`

	// The ID of the device.
	DeviceID string `json:"deviceId,omitempty"`
}

// Validate validates this device config
func (m *DeviceConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeviceConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceConfig) UnmarshalBinary(b []byte) error {
	var res DeviceConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
