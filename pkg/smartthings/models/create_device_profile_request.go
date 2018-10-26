// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateDeviceProfileRequest create device profile request
// swagger:model CreateDeviceProfileRequest
type CreateDeviceProfileRequest struct {

	// components
	// Required: true
	// Max Items: 10
	// Min Items: 1
	Components []*DeviceComponent `json:"components"`

	// metadata
	Metadata DeviceProfileMetadata `json:"metadata,omitempty"`

	// The name of the device profile.
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this create device profile request
func (m *CreateDeviceProfileRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComponents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateDeviceProfileRequest) validateComponents(formats strfmt.Registry) error {

	if err := validate.Required("components", "body", m.Components); err != nil {
		return err
	}

	iComponentsSize := int64(len(m.Components))

	if err := validate.MinItems("components", "body", iComponentsSize, 1); err != nil {
		return err
	}

	if err := validate.MaxItems("components", "body", iComponentsSize, 10); err != nil {
		return err
	}

	for i := 0; i < len(m.Components); i++ {
		if swag.IsZero(m.Components[i]) { // not required
			continue
		}

		if m.Components[i] != nil {
			if err := m.Components[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("components" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CreateDeviceProfileRequest) validateMetadata(formats strfmt.Registry) error {

	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if err := m.Metadata.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("metadata")
		}
		return err
	}

	return nil
}

func (m *CreateDeviceProfileRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateDeviceProfileRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateDeviceProfileRequest) UnmarshalBinary(b []byte) error {
	var res CreateDeviceProfileRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}