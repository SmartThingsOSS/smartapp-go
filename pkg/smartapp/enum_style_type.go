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

// EnumStyleType Style of the setting.
// swagger:model EnumStyleType
type EnumStyleType string

const (

	// EnumStyleTypeCOMPLETE captures enum value "COMPLETE"
	EnumStyleTypeCOMPLETE EnumStyleType = "COMPLETE"

	// EnumStyleTypeERROR captures enum value "ERROR"
	EnumStyleTypeERROR EnumStyleType = "ERROR"

	// EnumStyleTypeDEFAULT captures enum value "DEFAULT"
	EnumStyleTypeDEFAULT EnumStyleType = "DEFAULT"

	// EnumStyleTypeDROPDOWN captures enum value "DROPDOWN"
	EnumStyleTypeDROPDOWN EnumStyleType = "DROPDOWN"
)

// for schema
var enumStyleTypeEnum []interface{}

func init() {
	var res []EnumStyleType
	if err := json.Unmarshal([]byte(`["COMPLETE","ERROR","DEFAULT","DROPDOWN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		enumStyleTypeEnum = append(enumStyleTypeEnum, v)
	}
}

func (m EnumStyleType) validateEnumStyleTypeEnum(path, location string, value EnumStyleType) error {
	if err := validate.Enum(path, location, value, enumStyleTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this enum style type
func (m EnumStyleType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateEnumStyleTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
