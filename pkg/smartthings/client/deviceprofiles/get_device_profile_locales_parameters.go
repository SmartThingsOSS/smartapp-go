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

// NewGetDeviceProfileLocalesParams creates a new GetDeviceProfileLocalesParams object
// with the default values initialized.
func NewGetDeviceProfileLocalesParams() *GetDeviceProfileLocalesParams {
	var ()
	return &GetDeviceProfileLocalesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeviceProfileLocalesParamsWithTimeout creates a new GetDeviceProfileLocalesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDeviceProfileLocalesParamsWithTimeout(timeout time.Duration) *GetDeviceProfileLocalesParams {
	var ()
	return &GetDeviceProfileLocalesParams{

		timeout: timeout,
	}
}

// NewGetDeviceProfileLocalesParamsWithContext creates a new GetDeviceProfileLocalesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDeviceProfileLocalesParamsWithContext(ctx context.Context) *GetDeviceProfileLocalesParams {
	var ()
	return &GetDeviceProfileLocalesParams{

		Context: ctx,
	}
}

// NewGetDeviceProfileLocalesParamsWithHTTPClient creates a new GetDeviceProfileLocalesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDeviceProfileLocalesParamsWithHTTPClient(client *http.Client) *GetDeviceProfileLocalesParams {
	var ()
	return &GetDeviceProfileLocalesParams{
		HTTPClient: client,
	}
}

/*GetDeviceProfileLocalesParams contains all the parameters to send to the API endpoint
for the get device profile locales operation typically these are written to a http.Request
*/
type GetDeviceProfileLocalesParams struct {

	/*Authorization
	  OAuth token

	*/
	Authorization string
	/*DeviceProfileID
	  the device profile ID

	*/
	DeviceProfileID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get device profile locales params
func (o *GetDeviceProfileLocalesParams) WithTimeout(timeout time.Duration) *GetDeviceProfileLocalesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get device profile locales params
func (o *GetDeviceProfileLocalesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get device profile locales params
func (o *GetDeviceProfileLocalesParams) WithContext(ctx context.Context) *GetDeviceProfileLocalesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get device profile locales params
func (o *GetDeviceProfileLocalesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get device profile locales params
func (o *GetDeviceProfileLocalesParams) WithHTTPClient(client *http.Client) *GetDeviceProfileLocalesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get device profile locales params
func (o *GetDeviceProfileLocalesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get device profile locales params
func (o *GetDeviceProfileLocalesParams) WithAuthorization(authorization string) *GetDeviceProfileLocalesParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get device profile locales params
func (o *GetDeviceProfileLocalesParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithDeviceProfileID adds the deviceProfileID to the get device profile locales params
func (o *GetDeviceProfileLocalesParams) WithDeviceProfileID(deviceProfileID string) *GetDeviceProfileLocalesParams {
	o.SetDeviceProfileID(deviceProfileID)
	return o
}

// SetDeviceProfileID adds the deviceProfileId to the get device profile locales params
func (o *GetDeviceProfileLocalesParams) SetDeviceProfileID(deviceProfileID string) {
	o.DeviceProfileID = deviceProfileID
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeviceProfileLocalesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
