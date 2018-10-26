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

	models "github.com/SmartThingsOSS/smartapp-go/pkg/smartthings/models"
)

// NewInstallDeviceParams creates a new InstallDeviceParams object
// with the default values initialized.
func NewInstallDeviceParams() *InstallDeviceParams {
	var ()
	return &InstallDeviceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewInstallDeviceParamsWithTimeout creates a new InstallDeviceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewInstallDeviceParamsWithTimeout(timeout time.Duration) *InstallDeviceParams {
	var ()
	return &InstallDeviceParams{

		timeout: timeout,
	}
}

// NewInstallDeviceParamsWithContext creates a new InstallDeviceParams object
// with the default values initialized, and the ability to set a context for a request
func NewInstallDeviceParamsWithContext(ctx context.Context) *InstallDeviceParams {
	var ()
	return &InstallDeviceParams{

		Context: ctx,
	}
}

// NewInstallDeviceParamsWithHTTPClient creates a new InstallDeviceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewInstallDeviceParamsWithHTTPClient(client *http.Client) *InstallDeviceParams {
	var ()
	return &InstallDeviceParams{
		HTTPClient: client,
	}
}

/*InstallDeviceParams contains all the parameters to send to the API endpoint
for the install device operation typically these are written to a http.Request
*/
type InstallDeviceParams struct {

	/*Authorization
	  OAuth token

	*/
	Authorization string
	/*InstallationRequest
	  Installation Request

	*/
	InstallationRequest *models.DeviceInstallRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the install device params
func (o *InstallDeviceParams) WithTimeout(timeout time.Duration) *InstallDeviceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the install device params
func (o *InstallDeviceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the install device params
func (o *InstallDeviceParams) WithContext(ctx context.Context) *InstallDeviceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the install device params
func (o *InstallDeviceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the install device params
func (o *InstallDeviceParams) WithHTTPClient(client *http.Client) *InstallDeviceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the install device params
func (o *InstallDeviceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the install device params
func (o *InstallDeviceParams) WithAuthorization(authorization string) *InstallDeviceParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the install device params
func (o *InstallDeviceParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithInstallationRequest adds the installationRequest to the install device params
func (o *InstallDeviceParams) WithInstallationRequest(installationRequest *models.DeviceInstallRequest) *InstallDeviceParams {
	o.SetInstallationRequest(installationRequest)
	return o
}

// SetInstallationRequest adds the installationRequest to the install device params
func (o *InstallDeviceParams) SetInstallationRequest(installationRequest *models.DeviceInstallRequest) {
	o.InstallationRequest = installationRequest
}

// WriteToRequest writes these params to a swagger request
func (o *InstallDeviceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if o.InstallationRequest != nil {
		if err := r.SetBodyParam(o.InstallationRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
