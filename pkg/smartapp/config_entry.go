// Code generated by go-swagger; DO NOT EDIT.

package smartapp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ConfigEntry A configuration value type and the correspodning configuration.
// swagger:model ConfigEntry
type ConfigEntry struct {

	// The config if valueType is DEVICE, meaningless otherwise
	DeviceConfig *DeviceConfig `json:"deviceConfig,omitempty"`

	// The config if valueType is MODE, meaningless otherwise
	ModeConfig *ModeConfig `json:"modeConfig,omitempty"`

	// The config if valueType is STRING, meaningless otherwise
	StringConfig *StringConfig `json:"stringConfig,omitempty"`

	// The value type.
	// Enum: [STRING DEVICE MODE]
	ValueType *string `json:"valueType,omitempty"`
}

// Validate validates this config entry
func (m *ConfigEntry) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeviceConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModeConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStringConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValueType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigEntry) validateDeviceConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.DeviceConfig) { // not required
		return nil
	}

	if m.DeviceConfig != nil {
		if err := m.DeviceConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deviceConfig")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigEntry) validateModeConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.ModeConfig) { // not required
		return nil
	}

	if m.ModeConfig != nil {
		if err := m.ModeConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("modeConfig")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigEntry) validateStringConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.StringConfig) { // not required
		return nil
	}

	if m.StringConfig != nil {
		if err := m.StringConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stringConfig")
			}
			return err
		}
	}

	return nil
}

var configEntryTypeValueTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["STRING","DEVICE","MODE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		configEntryTypeValueTypePropEnum = append(configEntryTypeValueTypePropEnum, v)
	}
}

const (

	// ConfigEntryValueTypeSTRING captures enum value "STRING"
	ConfigEntryValueTypeSTRING string = "STRING"

	// ConfigEntryValueTypeDEVICE captures enum value "DEVICE"
	ConfigEntryValueTypeDEVICE string = "DEVICE"

	// ConfigEntryValueTypeMODE captures enum value "MODE"
	ConfigEntryValueTypeMODE string = "MODE"
)

// prop value enum
func (m *ConfigEntry) validateValueTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, configEntryTypeValueTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ConfigEntry) validateValueType(formats strfmt.Registry) error {

	if swag.IsZero(m.ValueType) { // not required
		return nil
	}

	// value enum
	if err := m.validateValueTypeEnum("valueType", "body", *m.ValueType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigEntry) UnmarshalBinary(b []byte) error {
	var res ConfigEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}