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

// NewGetDeviceProfileTranslationsParams creates a new GetDeviceProfileTranslationsParams object
// with the default values initialized.
func NewGetDeviceProfileTranslationsParams() *GetDeviceProfileTranslationsParams {
	var ()
	return &GetDeviceProfileTranslationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeviceProfileTranslationsParamsWithTimeout creates a new GetDeviceProfileTranslationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDeviceProfileTranslationsParamsWithTimeout(timeout time.Duration) *GetDeviceProfileTranslationsParams {
	var ()
	return &GetDeviceProfileTranslationsParams{

		timeout: timeout,
	}
}

// NewGetDeviceProfileTranslationsParamsWithContext creates a new GetDeviceProfileTranslationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDeviceProfileTranslationsParamsWithContext(ctx context.Context) *GetDeviceProfileTranslationsParams {
	var ()
	return &GetDeviceProfileTranslationsParams{

		Context: ctx,
	}
}

// NewGetDeviceProfileTranslationsParamsWithHTTPClient creates a new GetDeviceProfileTranslationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDeviceProfileTranslationsParamsWithHTTPClient(client *http.Client) *GetDeviceProfileTranslationsParams {
	var ()
	return &GetDeviceProfileTranslationsParams{
		HTTPClient: client,
	}
}

/*GetDeviceProfileTranslationsParams contains all the parameters to send to the API endpoint
for the get device profile translations operation typically these are written to a http.Request
*/
type GetDeviceProfileTranslationsParams struct {

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

// WithTimeout adds the timeout to the get device profile translations params
func (o *GetDeviceProfileTranslationsParams) WithTimeout(timeout time.Duration) *GetDeviceProfileTranslationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get device profile translations params
func (o *GetDeviceProfileTranslationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get device profile translations params
func (o *GetDeviceProfileTranslationsParams) WithContext(ctx context.Context) *GetDeviceProfileTranslationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get device profile translations params
func (o *GetDeviceProfileTranslationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get device profile translations params
func (o *GetDeviceProfileTranslationsParams) WithHTTPClient(client *http.Client) *GetDeviceProfileTranslationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get device profile translations params
func (o *GetDeviceProfileTranslationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get device profile translations params
func (o *GetDeviceProfileTranslationsParams) WithAuthorization(authorization string) *GetDeviceProfileTranslationsParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get device profile translations params
func (o *GetDeviceProfileTranslationsParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithDeviceProfileID adds the deviceProfileID to the get device profile translations params
func (o *GetDeviceProfileTranslationsParams) WithDeviceProfileID(deviceProfileID string) *GetDeviceProfileTranslationsParams {
	o.SetDeviceProfileID(deviceProfileID)
	return o
}

// SetDeviceProfileID adds the deviceProfileId to the get device profile translations params
func (o *GetDeviceProfileTranslationsParams) SetDeviceProfileID(deviceProfileID string) {
	o.DeviceProfileID = deviceProfileID
}

// WithLocale adds the locale to the get device profile translations params
func (o *GetDeviceProfileTranslationsParams) WithLocale(locale string) *GetDeviceProfileTranslationsParams {
	o.SetLocale(locale)
	return o
}

// SetLocale adds the locale to the get device profile translations params
func (o *GetDeviceProfileTranslationsParams) SetLocale(locale string) {
	o.Locale = locale
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeviceProfileTranslationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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