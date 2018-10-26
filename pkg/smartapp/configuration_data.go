// Code generated by go-swagger; DO NOT EDIT.

package smartapp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ConfigurationData A request for a pages flow to drive app installation.  This will only be available for executions of type "CONFIGURATION".
//
// swagger:model ConfigurationData
type ConfigurationData struct {

	// Settings currently configured by user.
	Config ConfigMap `json:"config,omitempty"`

	// The id of the installed app.
	InstalledAppID string `json:"installedAppId,omitempty"`

	// A developer defined page ID. Must be URL safe characters.
	PageID string `json:"pageId,omitempty"`

	// phase
	// Required: true
	Phase ConfigurationPhase `json:"phase"`

	// The previous page the user completed. Must be URL safe characters.
	PreviousPageID string `json:"previousPageId,omitempty"`
}

// Validate validates this configuration data
func (m *ConfigurationData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhase(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigurationData) validateConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.Config) { // not required
		return nil
	}

	if err := m.Config.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("config")
		}
		return err
	}

	return nil
}

func (m *ConfigurationData) validatePhase(formats strfmt.Registry) error {

	if err := m.Phase.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("phase")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigurationData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigurationData) UnmarshalBinary(b []byte) error {
	var res ConfigurationData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}