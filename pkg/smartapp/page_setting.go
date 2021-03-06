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

// PageSetting Jump to page Setting
// swagger:model PageSetting
type PageSetting struct {
	SectionSetting

	// The image url.
	// Max Length: 2048
	Image string `json:"image,omitempty"`

	// The page to navigate to.
	// Max Length: 128
	Page string `json:"page,omitempty"`

	// style
	Style StyleType `json:"style,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *PageSetting) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 SectionSetting
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.SectionSetting = aO0

	// AO1
	var dataAO1 struct {
		Image string `json:"image,omitempty"`

		Page string `json:"page,omitempty"`

		Style StyleType `json:"style,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Image = dataAO1.Image

	m.Page = dataAO1.Page

	m.Style = dataAO1.Style

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m PageSetting) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.SectionSetting)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Image string `json:"image,omitempty"`

		Page string `json:"page,omitempty"`

		Style StyleType `json:"style,omitempty"`
	}

	dataAO1.Image = m.Image

	dataAO1.Page = m.Page

	dataAO1.Style = m.Style

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this page setting
func (m *PageSetting) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with SectionSetting
	if err := m.SectionSetting.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStyle(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PageSetting) validateImage(formats strfmt.Registry) error {

	if swag.IsZero(m.Image) { // not required
		return nil
	}

	if err := validate.MaxLength("image", "body", string(m.Image), 2048); err != nil {
		return err
	}

	return nil
}

func (m *PageSetting) validatePage(formats strfmt.Registry) error {

	if swag.IsZero(m.Page) { // not required
		return nil
	}

	if err := validate.MaxLength("page", "body", string(m.Page), 128); err != nil {
		return err
	}

	return nil
}

func (m *PageSetting) validateStyle(formats strfmt.Registry) error {

	if swag.IsZero(m.Style) { // not required
		return nil
	}

	if err := m.Style.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("style")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PageSetting) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PageSetting) UnmarshalBinary(b []byte) error {
	var res PageSetting
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
