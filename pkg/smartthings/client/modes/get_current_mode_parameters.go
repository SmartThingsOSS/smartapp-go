// Code generated by go-swagger; DO NOT EDIT.

package modes

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

// NewGetCurrentModeParams creates a new GetCurrentModeParams object
// with the default values initialized.
func NewGetCurrentModeParams() *GetCurrentModeParams {
	var ()
	return &GetCurrentModeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCurrentModeParamsWithTimeout creates a new GetCurrentModeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCurrentModeParamsWithTimeout(timeout time.Duration) *GetCurrentModeParams {
	var ()
	return &GetCurrentModeParams{

		timeout: timeout,
	}
}

// NewGetCurrentModeParamsWithContext creates a new GetCurrentModeParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCurrentModeParamsWithContext(ctx context.Context) *GetCurrentModeParams {
	var ()
	return &GetCurrentModeParams{

		Context: ctx,
	}
}

// NewGetCurrentModeParamsWithHTTPClient creates a new GetCurrentModeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCurrentModeParamsWithHTTPClient(client *http.Client) *GetCurrentModeParams {
	var ()
	return &GetCurrentModeParams{
		HTTPClient: client,
	}
}

/*GetCurrentModeParams contains all the parameters to send to the API endpoint
for the get current mode operation typically these are written to a http.Request
*/
type GetCurrentModeParams struct {

	/*AcceptLanguage
	  Language header representing the client's preferred language. The format of the `Accept-Language` header follows what is defined in [RFC 7231, section 5.3.5](https://tools.ietf.org/html/rfc7231#section-5.3.5)

	*/
	AcceptLanguage *string
	/*Authorization
	  OAuth token

	*/
	Authorization string
	/*LocationID
	  The ID of the location.

	*/
	LocationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get current mode params
func (o *GetCurrentModeParams) WithTimeout(timeout time.Duration) *GetCurrentModeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get current mode params
func (o *GetCurrentModeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get current mode params
func (o *GetCurrentModeParams) WithContext(ctx context.Context) *GetCurrentModeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get current mode params
func (o *GetCurrentModeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get current mode params
func (o *GetCurrentModeParams) WithHTTPClient(client *http.Client) *GetCurrentModeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get current mode params
func (o *GetCurrentModeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAcceptLanguage adds the acceptLanguage to the get current mode params
func (o *GetCurrentModeParams) WithAcceptLanguage(acceptLanguage *string) *GetCurrentModeParams {
	o.SetAcceptLanguage(acceptLanguage)
	return o
}

// SetAcceptLanguage adds the acceptLanguage to the get current mode params
func (o *GetCurrentModeParams) SetAcceptLanguage(acceptLanguage *string) {
	o.AcceptLanguage = acceptLanguage
}

// WithAuthorization adds the authorization to the get current mode params
func (o *GetCurrentModeParams) WithAuthorization(authorization string) *GetCurrentModeParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get current mode params
func (o *GetCurrentModeParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithLocationID adds the locationID to the get current mode params
func (o *GetCurrentModeParams) WithLocationID(locationID string) *GetCurrentModeParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the get current mode params
func (o *GetCurrentModeParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetCurrentModeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AcceptLanguage != nil {

		// header param Accept-Language
		if err := r.SetHeaderParam("Accept-Language", *o.AcceptLanguage); err != nil {
			return err
		}

	}

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// path param locationId
	if err := r.SetPathParam("locationId", o.LocationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
