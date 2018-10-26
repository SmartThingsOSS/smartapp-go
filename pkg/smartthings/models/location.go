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

// Location location
// swagger:model Location
type Location struct {

	// A geographical latitude.
	Latitude float32 `json:"latitude,omitempty"`

	// An IETF BCP 47 language tag representing the chosen locale for this location.
	Locale string `json:"locale,omitempty"`

	// The ID of the location.
	// Format: uuid
	LocationID strfmt.UUID `json:"locationId,omitempty"`

	// A geographical longitude.
	Longitude float32 `json:"longitude,omitempty"`

	// A nickname given for the location (eg. Home)
	Name string `json:"name,omitempty"`

	// The radius in meters around latitude and longitude which defines this location.
	RegionRadius int32 `json:"regionRadius,omitempty"`

	// temperature scale
	// Enum: [F C]
	TemperatureScale string `json:"temperatureScale,omitempty"`

	// An ID matching the Java Time Zone ID of the location.  Only available if latitude and longitude have been
	// provided.
	//
	TimeZoneID string `json:"timeZoneId,omitempty"`
}

// Validate validates this location
func (m *Location) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLocationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTemperatureScale(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Location) validateLocationID(formats strfmt.Registry) error {

	if swag.IsZero(m.LocationID) { // not required
		return nil
	}

	if err := validate.FormatOf("locationId", "body", "uuid", m.LocationID.String(), formats); err != nil {
		return err
	}

	return nil
}

var locationTypeTemperatureScalePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["F","C"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		locationTypeTemperatureScalePropEnum = append(locationTypeTemperatureScalePropEnum, v)
	}
}

const (

	// LocationTemperatureScaleF captures enum value "F"
	LocationTemperatureScaleF string = "F"

	// LocationTemperatureScaleC captures enum value "C"
	LocationTemperatureScaleC string = "C"
)

// prop value enum
func (m *Location) validateTemperatureScaleEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, locationTypeTemperatureScalePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Location) validateTemperatureScale(formats strfmt.Registry) error {

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
func (m *Location) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Location) UnmarshalBinary(b []byte) error {
	var res Location
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
