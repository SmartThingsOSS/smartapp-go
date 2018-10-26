// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// DthDeviceDetails dth device details
// swagger:model DthDeviceDetails
type DthDeviceDetails struct {

	// The device network type.
	DeviceNetworkType string `json:"deviceNetworkType,omitempty"`

	// The identifier for the device's DeviceType.
	DeviceTypeID string `json:"deviceTypeId,omitempty"`

	// The name for the device's DeviceType.
	DeviceTypeName string `json:"deviceTypeName,omitempty"`
}

// Validate validates this dth device details
func (m *DthDeviceDetails) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DthDeviceDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DthDeviceDetails) UnmarshalBinary(b []byte) error {
	var res DthDeviceDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
