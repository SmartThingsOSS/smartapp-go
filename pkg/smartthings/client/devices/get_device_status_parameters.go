// Code generated by go-swagger; DO NOT EDIT.

package devices

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

// NewGetDeviceStatusParams creates a new GetDeviceStatusParams object
// with the default values initialized.
func NewGetDeviceStatusParams() *GetDeviceStatusParams {
	var ()
	return &GetDeviceStatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeviceStatusParamsWithTimeout creates a new GetDeviceStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDeviceStatusParamsWithTimeout(timeout time.Duration) *GetDeviceStatusParams {
	var ()
	return &GetDeviceStatusParams{

		timeout: timeout,
	}
}

// NewGetDeviceStatusParamsWithContext creates a new GetDeviceStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDeviceStatusParamsWithContext(ctx context.Context) *GetDeviceStatusParams {
	var ()
	return &GetDeviceStatusParams{

		Context: ctx,
	}
}

// NewGetDeviceStatusParamsWithHTTPClient creates a new GetDeviceStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDeviceStatusParamsWithHTTPClient(client *http.Client) *GetDeviceStatusParams {
	var ()
	return &GetDeviceStatusParams{
		HTTPClient: client,
	}
}

/*GetDeviceStatusParams contains all the parameters to send to the API endpoint
for the get device status operation typically these are written to a http.Request
*/
type GetDeviceStatusParams struct {

	/*Authorization
	  OAuth token

	*/
	Authorization string
	/*DeviceID
	  the device ID

	*/
	DeviceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get device status params
func (o *GetDeviceStatusParams) WithTimeout(timeout time.Duration) *GetDeviceStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get device status params
func (o *GetDeviceStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get device status params
func (o *GetDeviceStatusParams) WithContext(ctx context.Context) *GetDeviceStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get device status params
func (o *GetDeviceStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get device status params
func (o *GetDeviceStatusParams) WithHTTPClient(client *http.Client) *GetDeviceStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get device status params
func (o *GetDeviceStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get device status params
func (o *GetDeviceStatusParams) WithAuthorization(authorization string) *GetDeviceStatusParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get device status params
func (o *GetDeviceStatusParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithDeviceID adds the deviceID to the get device status params
func (o *GetDeviceStatusParams) WithDeviceID(deviceID string) *GetDeviceStatusParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the get device status params
func (o *GetDeviceStatusParams) SetDeviceID(deviceID string) {
	o.DeviceID = deviceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeviceStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// path param deviceId
	if err := r.SetPathParam("deviceId", o.DeviceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}