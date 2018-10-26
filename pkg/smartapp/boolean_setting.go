// Code generated by go-swagger; DO NOT EDIT.

package smartapp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BooleanSetting Boolean Setting
// swagger:model BooleanSetting
type BooleanSetting struct {
	SectionSetting

	// The image url.
	// Max Length: 2048
	Image string `json:"image,omitempty"`

	// Indicates if this input should refresh configs after a change in value.
	SubmitOnChange *bool `json:"submitOnChange,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *BooleanSetting) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 SectionSetting
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.SectionSetting = aO0

	// AO1
	var dataAO1 struct {
		Image string `json:"image,omitempty"`

		SubmitOnChange *bool `json:"submitOnChange,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Image = dataAO1.Image

	m.SubmitOnChange = dataAO1.SubmitOnChange

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m BooleanSetting) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.SectionSetting)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Image string `json:"image,omitempty"`

		SubmitOnChange *bool `json:"submitOnChange,omitempty"`
	}

	dataAO1.Image = m.Image

	dataAO1.SubmitOnChange = m.SubmitOnChange

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this boolean setting
func (m *BooleanSetting) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with SectionSetting
	if err := m.SectionSetting.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BooleanSetting) validateImage(formats strfmt.Registry) error {

	if swag.IsZero(m.Image) { // not required
		return nil
	}

	if err := validate.MaxLength("image", "body", string(m.Image), 2048); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BooleanSetting) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BooleanSetting) UnmarshalBinary(b []byte) error {
	var res BooleanSetting
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
