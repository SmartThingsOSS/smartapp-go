// Code generated by go-swagger; DO NOT EDIT.

package smartapp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LaunchPluginAction Launch a backing UI plugin in SmartThings Client.
// swagger:model LaunchPluginAction
type LaunchPluginAction struct {

	// The ID of the plugin to launch.
	PluginID string `json:"pluginId,omitempty"`
}

// Validate validates this launch plugin action
func (m *LaunchPluginAction) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LaunchPluginAction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LaunchPluginAction) UnmarshalBinary(b []byte) error {
	var res LaunchPluginAction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
