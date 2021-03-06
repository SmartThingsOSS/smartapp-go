// Code generated by go-swagger; DO NOT EDIT.

package smartapp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// BasicButtonPosition Position of where the button should be rendered within card.
// swagger:model BasicButtonPosition
type BasicButtonPosition string

const (

	// BasicButtonPositionLEFT captures enum value "LEFT"
	BasicButtonPositionLEFT BasicButtonPosition = "LEFT"

	// BasicButtonPositionCENTER captures enum value "CENTER"
	BasicButtonPositionCENTER BasicButtonPosition = "CENTER"

	// BasicButtonPositionRIGHT captures enum value "RIGHT"
	BasicButtonPositionRIGHT BasicButtonPosition = "RIGHT"
)

// for schema
var basicButtonPositionEnum []interface{}

func init() {
	var res []BasicButtonPosition
	if err := json.Unmarshal([]byte(`["LEFT","CENTER","RIGHT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		basicButtonPositionEnum = append(basicButtonPositionEnum, v)
	}
}

func (m BasicButtonPosition) validateBasicButtonPositionEnum(path, location string, value BasicButtonPosition) error {
	if err := validate.Enum(path, location, value, basicButtonPositionEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this basic button position
func (m BasicButtonPosition) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateBasicButtonPositionEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
