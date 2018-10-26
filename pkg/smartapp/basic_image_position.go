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

// BasicImagePosition Position of where the image should be rendered within card.
// swagger:model BasicImagePosition
type BasicImagePosition string

const (

	// BasicImagePositionLEFT captures enum value "LEFT"
	BasicImagePositionLEFT BasicImagePosition = "LEFT"

	// BasicImagePositionCENTER captures enum value "CENTER"
	BasicImagePositionCENTER BasicImagePosition = "CENTER"

	// BasicImagePositionRIGHT captures enum value "RIGHT"
	BasicImagePositionRIGHT BasicImagePosition = "RIGHT"
)

// for schema
var basicImagePositionEnum []interface{}

func init() {
	var res []BasicImagePosition
	if err := json.Unmarshal([]byte(`["LEFT","CENTER","RIGHT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		basicImagePositionEnum = append(basicImagePositionEnum, v)
	}
}

func (m BasicImagePosition) validateBasicImagePositionEnum(path, location string, value BasicImagePosition) error {
	if err := validate.Enum(path, location, value, basicImagePositionEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this basic image position
func (m BasicImagePosition) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateBasicImagePositionEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
