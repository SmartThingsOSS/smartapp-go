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

// SettingType Denotes the type of setting.
// swagger:model SettingType
type SettingType string

const (

	// SettingTypeDEVICE captures enum value "DEVICE"
	SettingTypeDEVICE SettingType = "DEVICE"

	// SettingTypeTEXT captures enum value "TEXT"
	SettingTypeTEXT SettingType = "TEXT"

	// SettingTypePASSWORD captures enum value "PASSWORD"
	SettingTypePASSWORD SettingType = "PASSWORD"

	// SettingTypeBOOLEAN captures enum value "BOOLEAN"
	SettingTypeBOOLEAN SettingType = "BOOLEAN"

	// SettingTypeENUM captures enum value "ENUM"
	SettingTypeENUM SettingType = "ENUM"

	// SettingTypeMODE captures enum value "MODE"
	SettingTypeMODE SettingType = "MODE"

	// SettingTypeSCENE captures enum value "SCENE"
	SettingTypeSCENE SettingType = "SCENE"

	// SettingTypeLINK captures enum value "LINK"
	SettingTypeLINK SettingType = "LINK"

	// SettingTypePAGE captures enum value "PAGE"
	SettingTypePAGE SettingType = "PAGE"

	// SettingTypeIMAGE captures enum value "IMAGE"
	SettingTypeIMAGE SettingType = "IMAGE"

	// SettingTypeIMAGES captures enum value "IMAGES"
	SettingTypeIMAGES SettingType = "IMAGES"

	// SettingTypeVIDEO captures enum value "VIDEO"
	SettingTypeVIDEO SettingType = "VIDEO"

	// SettingTypeTIME captures enum value "TIME"
	SettingTypeTIME SettingType = "TIME"

	// SettingTypePARAGRAPH captures enum value "PARAGRAPH"
	SettingTypePARAGRAPH SettingType = "PARAGRAPH"

	// SettingTypeEMAIL captures enum value "EMAIL"
	SettingTypeEMAIL SettingType = "EMAIL"

	// SettingTypeDECIMAL captures enum value "DECIMAL"
	SettingTypeDECIMAL SettingType = "DECIMAL"

	// SettingTypeNUMBER captures enum value "NUMBER"
	SettingTypeNUMBER SettingType = "NUMBER"

	// SettingTypePHONE captures enum value "PHONE"
	SettingTypePHONE SettingType = "PHONE"

	// SettingTypeOAUTH captures enum value "OAUTH"
	SettingTypeOAUTH SettingType = "OAUTH"
)

// for schema
var settingTypeEnum []interface{}

func init() {
	var res []SettingType
	if err := json.Unmarshal([]byte(`["DEVICE","TEXT","PASSWORD","BOOLEAN","ENUM","MODE","SCENE","LINK","PAGE","IMAGE","IMAGES","VIDEO","TIME","PARAGRAPH","EMAIL","DECIMAL","NUMBER","PHONE","OAUTH"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		settingTypeEnum = append(settingTypeEnum, v)
	}
}

func (m SettingType) validateSettingTypeEnum(path, location string, value SettingType) error {
	if err := validate.Enum(path, location, value, settingTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this setting type
func (m SettingType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateSettingTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
