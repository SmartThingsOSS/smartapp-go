// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SceneCapability A device component capability
// swagger:model SceneCapability
type SceneCapability struct {

	// The id of the capability
	CapabilityID string `json:"capabilityId,omitempty"`

	// Capability commands
	Commands map[string]SceneCommand `json:"commands,omitempty"`

	// The status of the capability
	// Enum: [proposed live deprecated dead]
	Status string `json:"status,omitempty"`
}

// Validate validates this scene capability
func (m *SceneCapability) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommands(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SceneCapability) validateCommands(formats strfmt.Registry) error {

	if swag.IsZero(m.Commands) { // not required
		return nil
	}

	for k := range m.Commands {

		if err := validate.Required("commands"+"."+k, "body", m.Commands[k]); err != nil {
			return err
		}
		if val, ok := m.Commands[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

var sceneCapabilityTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["proposed","live","deprecated","dead"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sceneCapabilityTypeStatusPropEnum = append(sceneCapabilityTypeStatusPropEnum, v)
	}
}

const (

	// SceneCapabilityStatusProposed captures enum value "proposed"
	SceneCapabilityStatusProposed string = "proposed"

	// SceneCapabilityStatusLive captures enum value "live"
	SceneCapabilityStatusLive string = "live"

	// SceneCapabilityStatusDeprecated captures enum value "deprecated"
	SceneCapabilityStatusDeprecated string = "deprecated"

	// SceneCapabilityStatusDead captures enum value "dead"
	SceneCapabilityStatusDead string = "dead"
)

// prop value enum
func (m *SceneCapability) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, sceneCapabilityTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SceneCapability) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SceneCapability) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SceneCapability) UnmarshalBinary(b []byte) error {
	var res SceneCapability
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
