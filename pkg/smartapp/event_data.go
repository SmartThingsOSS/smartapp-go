// Code generated by go-swagger; DO NOT EDIT.

package smartapp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// EventData The data payload to an execution request with an AppLifecycle of EVENT.
// swagger:model EventData
type EventData struct {

	// An OAuth token to use when calling into SmartThings API's.
	// Required: true
	AuthToken *string `json:"authToken"`

	// events
	Events []*Event `json:"events"`

	// installed app
	// Required: true
	InstalledApp *InstalledApp `json:"installedApp"`
}

// Validate validates this event data
func (m *EventData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthToken(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEvents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstalledApp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EventData) validateAuthToken(formats strfmt.Registry) error {

	if err := validate.Required("authToken", "body", m.AuthToken); err != nil {
		return err
	}

	return nil
}

func (m *EventData) validateEvents(formats strfmt.Registry) error {

	if swag.IsZero(m.Events) { // not required
		return nil
	}

	for i := 0; i < len(m.Events); i++ {
		if swag.IsZero(m.Events[i]) { // not required
			continue
		}

		if m.Events[i] != nil {
			if err := m.Events[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("events" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EventData) validateInstalledApp(formats strfmt.Registry) error {

	if err := validate.Required("installedApp", "body", m.InstalledApp); err != nil {
		return err
	}

	if m.InstalledApp != nil {
		if err := m.InstalledApp.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("installedApp")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EventData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EventData) UnmarshalBinary(b []byte) error {
	var res EventData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}