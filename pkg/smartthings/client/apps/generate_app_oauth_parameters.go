// Code generated by go-swagger; DO NOT EDIT.

package apps

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

// NewGenerateAppOauthParams creates a new GenerateAppOauthParams object
// with the default values initialized.
func NewGenerateAppOauthParams() *GenerateAppOauthParams {
	var ()
	return &GenerateAppOauthParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGenerateAppOauthParamsWithTimeout creates a new GenerateAppOauthParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGenerateAppOauthParamsWithTimeout(timeout time.Duration) *GenerateAppOauthParams {
	var ()
	return &GenerateAppOauthParams{

		timeout: timeout,
	}
}

// NewGenerateAppOauthParamsWithContext creates a new GenerateAppOauthParams object
// with the default values initialized, and the ability to set a context for a request
func NewGenerateAppOauthParamsWithContext(ctx context.Context) *GenerateAppOauthParams {
	var ()
	return &GenerateAppOauthParams{

		Context: ctx,
	}
}

// NewGenerateAppOauthParamsWithHTTPClient creates a new GenerateAppOauthParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGenerateAppOauthParamsWithHTTPClient(client *http.Client) *GenerateAppOauthParams {
	var ()
	return &GenerateAppOauthParams{
		HTTPClient: client,
	}
}

/*GenerateAppOauthParams contains all the parameters to send to the API endpoint
for the generate app oauth operation typically these are written to a http.Request
*/
type GenerateAppOauthParams struct {

	/*Authorization
	  OAuth token

	*/
	Authorization string
	/*AppNameOrID
	  The appName or appId field of an app.

	*/
	AppNameOrID string
	/*GenerateAppOAuthRequest*/
	GenerateAppOAuthRequest *models.GenerateAppOAuthRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the generate app oauth params
func (o *GenerateAppOauthParams) WithTimeout(timeout time.Duration) *GenerateAppOauthParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the generate app oauth params
func (o *GenerateAppOauthParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the generate app oauth params
func (o *GenerateAppOauthParams) WithContext(ctx context.Context) *GenerateAppOauthParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the generate app oauth params
func (o *GenerateAppOauthParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the generate app oauth params
func (o *GenerateAppOauthParams) WithHTTPClient(client *http.Client) *GenerateAppOauthParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the generate app oauth params
func (o *GenerateAppOauthParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the generate app oauth params
func (o *GenerateAppOauthParams) WithAuthorization(authorization string) *GenerateAppOauthParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the generate app oauth params
func (o *GenerateAppOauthParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithAppNameOrID adds the appNameOrID to the generate app oauth params
func (o *GenerateAppOauthParams) WithAppNameOrID(appNameOrID string) *GenerateAppOauthParams {
	o.SetAppNameOrID(appNameOrID)
	return o
}

// SetAppNameOrID adds the appNameOrId to the generate app oauth params
func (o *GenerateAppOauthParams) SetAppNameOrID(appNameOrID string) {
	o.AppNameOrID = appNameOrID
}

// WithGenerateAppOAuthRequest adds the generateAppOAuthRequest to the generate app oauth params
func (o *GenerateAppOauthParams) WithGenerateAppOAuthRequest(generateAppOAuthRequest *models.GenerateAppOAuthRequest) *GenerateAppOauthParams {
	o.SetGenerateAppOAuthRequest(generateAppOAuthRequest)
	return o
}

// SetGenerateAppOAuthRequest adds the generateAppOAuthRequest to the generate app oauth params
func (o *GenerateAppOauthParams) SetGenerateAppOAuthRequest(generateAppOAuthRequest *models.GenerateAppOAuthRequest) {
	o.GenerateAppOAuthRequest = generateAppOAuthRequest
}

// WriteToRequest writes these params to a swagger request
func (o *GenerateAppOauthParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// path param appNameOrId
	if err := r.SetPathParam("appNameOrId", o.AppNameOrID); err != nil {
		return err
	}

	if o.GenerateAppOAuthRequest != nil {
		if err := r.SetBodyParam(o.GenerateAppOAuthRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}