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

// UpdateLocationRequest update location request
// swagger:model UpdateLocationRequest
type UpdateLocationRequest struct {

	// A geographical latitude.
	Latitude float32 `json:"latitude,omitempty"`

	// An IETF BCP 47 language tag representing the chosen locale for this location.
	Locale string `json:"locale,omitempty"`

	// A geographical longitude.
	Longitude float32 `json:"longitude,omitempty"`

	// A nickname for the location.
	Name string `json:"name,omitempty"`

	// The radius in meters around latitude and longitude which defines this location.
	RegionRadius int64 `json:"regionRadius,omitempty"`

	// The desired temperature scale used within location.
	// Enum: [F C]
	TemperatureScale string `json:"temperatureScale,omitempty"`
}

// Validate validates this update location request
func (m *UpdateLocationRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTemperatureScale(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var updateLocationRequestTypeTemperatureScalePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["F","C"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateLocationRequestTypeTemperatureScalePropEnum = append(updateLocationRequestTypeTemperatureScalePropEnum, v)
	}
}

const (

	// UpdateLocationRequestTemperatureScaleF captures enum value "F"
	UpdateLocationRequestTemperatureScaleF string = "F"

	// UpdateLocationRequestTemperatureScaleC captures enum value "C"
	UpdateLocationRequestTemperatureScaleC string = "C"
)

// prop value enum
func (m *UpdateLocationRequest) validateTemperatureScaleEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, updateLocationRequestTypeTemperatureScalePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *UpdateLocationRequest) validateTemperatureScale(formats strfmt.Registry) error {

	if swag.IsZero(m.TemperatureScale) { // not required
		return nil
	}

	// value enum
	if err := m.validateTemperatureScaleEnum("temperatureScale", "body", m.TemperatureScale); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateLocationRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateLocationRequest) UnmarshalBinary(b []byte) error {
	var res UpdateLocationRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
