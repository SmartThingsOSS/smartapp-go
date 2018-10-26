// Code generated by go-swagger; DO NOT EDIT.

package smartapp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// BasicBody Body definition for a Basic V1 template.
// swagger:model BasicBody
type BasicBody struct {

	// image
	Image *BasicImage `json:"image,omitempty"`

	// text
	Text *BasicText `json:"text,omitempty"`
}

// Validate validates this basic body
func (m *BasicBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateText(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BasicBody) validateImage(formats strfmt.Registry) error {

	if swag.IsZero(m.Image) { // not required
		return nil
	}

	if m.Image != nil {
		if err := m.Image.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("image")
			}
			return err
		}
	}

	return nil
}

func (m *BasicBody) validateText(formats strfmt.Registry) error {

	if swag.IsZero(m.Text) { // not required
		return nil
	}

	if m.Text != nil {
		if err := m.Text.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("text")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BasicBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BasicBody) UnmarshalBinary(b []byte) error {
	var res BasicBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
