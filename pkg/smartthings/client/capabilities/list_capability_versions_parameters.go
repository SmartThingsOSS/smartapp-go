// Code generated by go-swagger; DO NOT EDIT.

package capabilities

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

// NewListCapabilityVersionsParams creates a new ListCapabilityVersionsParams object
// with the default values initialized.
func NewListCapabilityVersionsParams() *ListCapabilityVersionsParams {
	var ()
	return &ListCapabilityVersionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListCapabilityVersionsParamsWithTimeout creates a new ListCapabilityVersionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListCapabilityVersionsParamsWithTimeout(timeout time.Duration) *ListCapabilityVersionsParams {
	var ()
	return &ListCapabilityVersionsParams{

		timeout: timeout,
	}
}

// NewListCapabilityVersionsParamsWithContext creates a new ListCapabilityVersionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListCapabilityVersionsParamsWithContext(ctx context.Context) *ListCapabilityVersionsParams {
	var ()
	return &ListCapabilityVersionsParams{

		Context: ctx,
	}
}

// NewListCapabilityVersionsParamsWithHTTPClient creates a new ListCapabilityVersionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListCapabilityVersionsParamsWithHTTPClient(client *http.Client) *ListCapabilityVersionsParams {
	var ()
	return &ListCapabilityVersionsParams{
		HTTPClient: client,
	}
}

/*ListCapabilityVersionsParams contains all the parameters to send to the API endpoint
for the list capability versions operation typically these are written to a http.Request
*/
type ListCapabilityVersionsParams struct {

	/*Authorization
	  OAuth token

	*/
	Authorization string
	/*CapabilityID
	  The ID of the capability

	*/
	CapabilityID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list capability versions params
func (o *ListCapabilityVersionsParams) WithTimeout(timeout time.Duration) *ListCapabilityVersionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list capability versions params
func (o *ListCapabilityVersionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list capability versions params
func (o *ListCapabilityVersionsParams) WithContext(ctx context.Context) *ListCapabilityVersionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list capability versions params
func (o *ListCapabilityVersionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list capability versions params
func (o *ListCapabilityVersionsParams) WithHTTPClient(client *http.Client) *ListCapabilityVersionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list capability versions params
func (o *ListCapabilityVersionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the list capability versions params
func (o *ListCapabilityVersionsParams) WithAuthorization(authorization string) *ListCapabilityVersionsParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the list capability versions params
func (o *ListCapabilityVersionsParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithCapabilityID adds the capabilityID to the list capability versions params
func (o *ListCapabilityVersionsParams) WithCapabilityID(capabilityID string) *ListCapabilityVersionsParams {
	o.SetCapabilityID(capabilityID)
	return o
}

// SetCapabilityID adds the capabilityId to the list capability versions params
func (o *ListCapabilityVersionsParams) SetCapabilityID(capabilityID string) {
	o.CapabilityID = capabilityID
}

// WriteToRequest writes these params to a swagger request
func (o *ListCapabilityVersionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// path param capabilityId
	if err := r.SetPathParam("capabilityId", o.CapabilityID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
