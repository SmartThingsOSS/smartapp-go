// Code generated by go-swagger; DO NOT EDIT.

package installedapps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/SmartThingsOSS/smartapp-go/pkg/smartthings/models"
)

// GetInstallationConfigReader is a Reader for the GetInstallationConfig structure.
type GetInstallationConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInstallationConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetInstallationConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetInstallationConfigBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetInstallationConfigUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetInstallationConfigForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewGetInstallationConfigTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetInstallationConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetInstallationConfigOK creates a GetInstallationConfigOK with default headers values
func NewGetInstallationConfigOK() *GetInstallationConfigOK {
	return &GetInstallationConfigOK{}
}

/*GetInstallationConfigOK handles this case with default header values.

An install configuration detail model.
*/
type GetInstallationConfigOK struct {
	Payload *models.InstallConfigurationDetail
}

func (o *GetInstallationConfigOK) Error() string {
	return fmt.Sprintf("[GET /installedapps/{installedAppId}/configs/{configurationId}][%d] getInstallationConfigOK  %+v", 200, o.Payload)
}

func (o *GetInstallationConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InstallConfigurationDetail)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstallationConfigBadRequest creates a GetInstallationConfigBadRequest with default headers values
func NewGetInstallationConfigBadRequest() *GetInstallationConfigBadRequest {
	return &GetInstallationConfigBadRequest{}
}

/*GetInstallationConfigBadRequest handles this case with default header values.

Bad request
*/
type GetInstallationConfigBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *GetInstallationConfigBadRequest) Error() string {
	return fmt.Sprintf("[GET /installedapps/{installedAppId}/configs/{configurationId}][%d] getInstallationConfigBadRequest  %+v", 400, o.Payload)
}

func (o *GetInstallationConfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstallationConfigUnauthorized creates a GetInstallationConfigUnauthorized with default headers values
func NewGetInstallationConfigUnauthorized() *GetInstallationConfigUnauthorized {
	return &GetInstallationConfigUnauthorized{}
}

/*GetInstallationConfigUnauthorized handles this case with default header values.

Not authenticated
*/
type GetInstallationConfigUnauthorized struct {
}

func (o *GetInstallationConfigUnauthorized) Error() string {
	return fmt.Sprintf("[GET /installedapps/{installedAppId}/configs/{configurationId}][%d] getInstallationConfigUnauthorized ", 401)
}

func (o *GetInstallationConfigUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstallationConfigForbidden creates a GetInstallationConfigForbidden with default headers values
func NewGetInstallationConfigForbidden() *GetInstallationConfigForbidden {
	return &GetInstallationConfigForbidden{}
}

/*GetInstallationConfigForbidden handles this case with default header values.

Not authorized
*/
type GetInstallationConfigForbidden struct {
}

func (o *GetInstallationConfigForbidden) Error() string {
	return fmt.Sprintf("[GET /installedapps/{installedAppId}/configs/{configurationId}][%d] getInstallationConfigForbidden ", 403)
}

func (o *GetInstallationConfigForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstallationConfigTooManyRequests creates a GetInstallationConfigTooManyRequests with default headers values
func NewGetInstallationConfigTooManyRequests() *GetInstallationConfigTooManyRequests {
	return &GetInstallationConfigTooManyRequests{}
}

/*GetInstallationConfigTooManyRequests handles this case with default header values.

Too many requests
*/
type GetInstallationConfigTooManyRequests struct {
	Payload *models.ErrorResponse
}

func (o *GetInstallationConfigTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /installedapps/{installedAppId}/configs/{configurationId}][%d] getInstallationConfigTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetInstallationConfigTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstallationConfigDefault creates a GetInstallationConfigDefault with default headers values
func NewGetInstallationConfigDefault(code int) *GetInstallationConfigDefault {
	return &GetInstallationConfigDefault{
		_statusCode: code,
	}
}

/*GetInstallationConfigDefault handles this case with default header values.

Unexpected error
*/
type GetInstallationConfigDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get installation config default response
func (o *GetInstallationConfigDefault) Code() int {
	return o._statusCode
}

func (o *GetInstallationConfigDefault) Error() string {
	return fmt.Sprintf("[GET /installedapps/{installedAppId}/configs/{configurationId}][%d] getInstallationConfig default  %+v", o._statusCode, o.Payload)
}

func (o *GetInstallationConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
