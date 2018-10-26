// Code generated by go-swagger; DO NOT EDIT.

package deviceprofiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteDeviceProfileTranslationsParams creates a new DeleteDeviceProfileTranslationsParams object
// with the default values initialized.
func NewDeleteDeviceProfileTranslationsParams() *DeleteDeviceProfileTranslationsParams {
	var ()
	return &DeleteDeviceProfileTranslationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDeviceProfileTranslationsParamsWithTimeout creates a new DeleteDeviceProfileTranslationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteDeviceProfileTranslationsParamsWithTimeout(timeout time.Duration) *DeleteDeviceProfileTranslationsParams {
	var ()
	return &DeleteDeviceProfileTranslationsParams{

		timeout: timeout,
	}
}

// NewDeleteDeviceProfileTranslationsParamsWithContext creates a new DeleteDeviceProfileTranslationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteDeviceProfileTranslationsParamsWithContext(ctx context.Context) *DeleteDeviceProfileTranslationsParams {
	var ()
	return &DeleteDeviceProfileTranslationsParams{

		Context: ctx,
	}
}

// NewDeleteDeviceProfileTranslationsParamsWithHTTPClient creates a new DeleteDeviceProfileTranslationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteDeviceProfileTranslationsParamsWithHTTPClient(client *http.Client) *DeleteDeviceProfileTranslationsParams {
	var ()
	return &DeleteDeviceProfileTranslationsParams{
		HTTPClient: client,
	}
}

/*DeleteDeviceProfileTranslationsParams contains all the parameters to send to the API endpoint
for the delete device profile translations operation typically these are written to a http.Request
*/
type DeleteDeviceProfileTranslationsParams struct {

	/*Authorization
	  OAuth token

	*/
	Authorization string
	/*DeviceProfileID
	  the device profile ID

	*/
	DeviceProfileID string
	/*Locale
	  The tag of the locale as defined in [RFC bcp47](http://www.rfc-editor.org/rfc/bcp/bcp47.txt).

	*/
	Locale string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete device profile translations params
func (o *DeleteDeviceProfileTranslationsParams) WithTimeout(timeout time.Duration) *DeleteDeviceProfileTranslationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete device profile translations params
func (o *DeleteDeviceProfileTranslationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete device profile translations params
func (o *DeleteDeviceProfileTranslationsParams) WithContext(ctx context.Context) *DeleteDeviceProfileTranslationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete device profile translations params
func (o *DeleteDeviceProfileTranslationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete device profile translations params
func (o *DeleteDeviceProfileTranslationsParams) WithHTTPClient(client *http.Client) *DeleteDeviceProfileTranslationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete device profile translations params
func (o *DeleteDeviceProfileTranslationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the delete device profile translations params
func (o *DeleteDeviceProfileTranslationsParams) WithAuthorization(authorization string) *DeleteDeviceProfileTranslationsParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the delete device profile translations params
func (o *DeleteDeviceProfileTranslationsParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithDeviceProfileID adds the deviceProfileID to the delete device profile translations params
func (o *DeleteDeviceProfileTranslationsParams) WithDeviceProfileID(deviceProfileID string) *DeleteDeviceProfileTranslationsParams {
	o.SetDeviceProfileID(deviceProfileID)
	return o
}

// SetDeviceProfileID adds the deviceProfileId to the delete device profile translations params
func (o *DeleteDeviceProfileTranslationsParams) SetDeviceProfileID(deviceProfileID string) {
	o.DeviceProfileID = deviceProfileID
}

// WithLocale adds the locale to the delete device profile translations params
func (o *DeleteDeviceProfileTranslationsParams) WithLocale(locale string) *DeleteDeviceProfileTranslationsParams {
	o.SetLocale(locale)
	return o
}

// SetLocale adds the locale to the delete device profile translations params
func (o *DeleteDeviceProfileTranslationsParams) SetLocale(locale string) {
	o.Locale = locale
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDeviceProfileTranslationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// path param deviceProfileId
	if err := r.SetPathParam("deviceProfileId", o.DeviceProfileID); err != nil {
		return err
	}

	// path param locale
	if err := r.SetPathParam("locale", o.Locale); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
