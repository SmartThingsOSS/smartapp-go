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

// Setting setting
// swagger:model Setting
type Setting struct {

	// Description of the app to be configured.
	// Max Length: 2048
	Description string `json:"description,omitempty"`

	// A developer defined configuration ID.
	// Max Length: 128
	ID string `json:"id,omitempty"`

	// Name of the setting to be configured.
	// Max Length: 128
	Name string `json:"name,omitempty"`
}

// Validate validates this setting
func (m *Setting) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Setting) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", string(m.Description), 2048); err != nil {
		return err
	}

	return nil
}

func (m *Setting) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.MaxLength("id", "body", string(m.ID), 128); err != nil {
		return err
	}

	return nil
}

func (m *Setting) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MaxLength("name", "body", string(m.Name), 128); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Setting) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Setting) UnmarshalBinary(b []byte) error {
	var res Setting
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
