// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// SceneArgument A capability command argument
// swagger:model SceneArgument
type SceneArgument struct {

	// the name of the command
	Name string `json:"name,omitempty"`

	// the schema of the command
	Schema interface{} `json:"schema,omitempty"`

	// The value being set for the capability command
	Value interface{} `json:"value,omitempty"`
}

// Validate validates this scene argument
func (m *SceneArgument) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SceneArgument) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SceneArgument) UnmarshalBinary(b []byte) error {
	var res SceneArgument
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
