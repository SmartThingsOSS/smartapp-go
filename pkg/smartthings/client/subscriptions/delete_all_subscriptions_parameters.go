// Code generated by go-swagger; DO NOT EDIT.

package subscriptions

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

// NewDeleteAllSubscriptionsParams creates a new DeleteAllSubscriptionsParams object
// with the default values initialized.
func NewDeleteAllSubscriptionsParams() *DeleteAllSubscriptionsParams {
	var ()
	return &DeleteAllSubscriptionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAllSubscriptionsParamsWithTimeout creates a new DeleteAllSubscriptionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteAllSubscriptionsParamsWithTimeout(timeout time.Duration) *DeleteAllSubscriptionsParams {
	var ()
	return &DeleteAllSubscriptionsParams{

		timeout: timeout,
	}
}

// NewDeleteAllSubscriptionsParamsWithContext creates a new DeleteAllSubscriptionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteAllSubscriptionsParamsWithContext(ctx context.Context) *DeleteAllSubscriptionsParams {
	var ()
	return &DeleteAllSubscriptionsParams{

		Context: ctx,
	}
}

// NewDeleteAllSubscriptionsParamsWithHTTPClient creates a new DeleteAllSubscriptionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteAllSubscriptionsParamsWithHTTPClient(client *http.Client) *DeleteAllSubscriptionsParams {
	var ()
	return &DeleteAllSubscriptionsParams{
		HTTPClient: client,
	}
}

/*DeleteAllSubscriptionsParams contains all the parameters to send to the API endpoint
for the delete all subscriptions operation typically these are written to a http.Request
*/
type DeleteAllSubscriptionsParams struct {

	/*Authorization
	  OAuth token

	*/
	Authorization string
	/*DeviceID
	  Limit deletion to subscriptions for a particular device.

	*/
	DeviceID *string
	/*InstalledAppID
	  The ID of the installed application.

	*/
	InstalledAppID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete all subscriptions params
func (o *DeleteAllSubscriptionsParams) WithTimeout(timeout time.Duration) *DeleteAllSubscriptionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete all subscriptions params
func (o *DeleteAllSubscriptionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete all subscriptions params
func (o *DeleteAllSubscriptionsParams) WithContext(ctx context.Context) *DeleteAllSubscriptionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete all subscriptions params
func (o *DeleteAllSubscriptionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete all subscriptions params
func (o *DeleteAllSubscriptionsParams) WithHTTPClient(client *http.Client) *DeleteAllSubscriptionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete all subscriptions params
func (o *DeleteAllSubscriptionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the delete all subscriptions params
func (o *DeleteAllSubscriptionsParams) WithAuthorization(authorization string) *DeleteAllSubscriptionsParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the delete all subscriptions params
func (o *DeleteAllSubscriptionsParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithDeviceID adds the deviceID to the delete all subscriptions params
func (o *DeleteAllSubscriptionsParams) WithDeviceID(deviceID *string) *DeleteAllSubscriptionsParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the delete all subscriptions params
func (o *DeleteAllSubscriptionsParams) SetDeviceID(deviceID *string) {
	o.DeviceID = deviceID
}

// WithInstalledAppID adds the installedAppID to the delete all subscriptions params
func (o *DeleteAllSubscriptionsParams) WithInstalledAppID(installedAppID string) *DeleteAllSubscriptionsParams {
	o.SetInstalledAppID(installedAppID)
	return o
}

// SetInstalledAppID adds the installedAppId to the delete all subscriptions params
func (o *DeleteAllSubscriptionsParams) SetInstalledAppID(installedAppID string) {
	o.InstalledAppID = installedAppID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAllSubscriptionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if o.DeviceID != nil {

		// query param deviceId
		var qrDeviceID string
		if o.DeviceID != nil {
			qrDeviceID = *o.DeviceID
		}
		qDeviceID := qrDeviceID
		if qDeviceID != "" {
			if err := r.SetQueryParam("deviceId", qDeviceID); err != nil {
				return err
			}
		}

	}

	// path param installedAppId
	if err := r.SetPathParam("installedAppId", o.InstalledAppID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}