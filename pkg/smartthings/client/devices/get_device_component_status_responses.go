// Code generated by go-swagger; DO NOT EDIT.

package devices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/SmartThingsOSS/smartapp-go/pkg/smartthings/models"
)

// GetDeviceComponentStatusReader is a Reader for the GetDeviceComponentStatus structure.
type GetDeviceComponentStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeviceComponentStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDeviceComponentStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetDeviceComponentStatusBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetDeviceComponentStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetDeviceComponentStatusForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewGetDeviceComponentStatusTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetDeviceComponentStatusDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDeviceComponentStatusOK creates a GetDeviceComponentStatusOK with default headers values
func NewGetDeviceComponentStatusOK() *GetDeviceComponentStatusOK {
	return &GetDeviceComponentStatusOK{}
}

/*GetDeviceComponentStatusOK handles this case with default header values.

Successful return  current status of device component's attributes.
*/
type GetDeviceComponentStatusOK struct {
	Payload models.ComponentStatus
}

func (o *GetDeviceComponentStatusOK) Error() string {
	return fmt.Sprintf("[GET /devices/{deviceId}/components/{componentId}/status][%d] getDeviceComponentStatusOK  %+v", 200, o.Payload)
}

func (o *GetDeviceComponentStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeviceComponentStatusBadRequest creates a GetDeviceComponentStatusBadRequest with default headers values
func NewGetDeviceComponentStatusBadRequest() *GetDeviceComponentStatusBadRequest {
	return &GetDeviceComponentStatusBadRequest{}
}

/*GetDeviceComponentStatusBadRequest handles this case with default header values.

Bad request
*/
type GetDeviceComponentStatusBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *GetDeviceComponentStatusBadRequest) Error() string {
	return fmt.Sprintf("[GET /devices/{deviceId}/components/{componentId}/status][%d] getDeviceComponentStatusBadRequest  %+v", 400, o.Payload)
}

func (o *GetDeviceComponentStatusBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeviceComponentStatusUnauthorized creates a GetDeviceComponentStatusUnauthorized with default headers values
func NewGetDeviceComponentStatusUnauthorized() *GetDeviceComponentStatusUnauthorized {
	return &GetDeviceComponentStatusUnauthorized{}
}

/*GetDeviceComponentStatusUnauthorized handles this case with default header values.

Not authenticated
*/
type GetDeviceComponentStatusUnauthorized struct {
}

func (o *GetDeviceComponentStatusUnauthorized) Error() string {
	return fmt.Sprintf("[GET /devices/{deviceId}/components/{componentId}/status][%d] getDeviceComponentStatusUnauthorized ", 401)
}

func (o *GetDeviceComponentStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDeviceComponentStatusForbidden creates a GetDeviceComponentStatusForbidden with default headers values
func NewGetDeviceComponentStatusForbidden() *GetDeviceComponentStatusForbidden {
	return &GetDeviceComponentStatusForbidden{}
}

/*GetDeviceComponentStatusForbidden handles this case with default header values.

Not authorized
*/
type GetDeviceComponentStatusForbidden struct {
}

func (o *GetDeviceComponentStatusForbidden) Error() string {
	return fmt.Sprintf("[GET /devices/{deviceId}/components/{componentId}/status][%d] getDeviceComponentStatusForbidden ", 403)
}

func (o *GetDeviceComponentStatusForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDeviceComponentStatusTooManyRequests creates a GetDeviceComponentStatusTooManyRequests with default headers values
func NewGetDeviceComponentStatusTooManyRequests() *GetDeviceComponentStatusTooManyRequests {
	return &GetDeviceComponentStatusTooManyRequests{}
}

/*GetDeviceComponentStatusTooManyRequests handles this case with default header values.

Too many requests
*/
type GetDeviceComponentStatusTooManyRequests struct {
	Payload *models.ErrorResponse
}

func (o *GetDeviceComponentStatusTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /devices/{deviceId}/components/{componentId}/status][%d] getDeviceComponentStatusTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetDeviceComponentStatusTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeviceComponentStatusDefault creates a GetDeviceComponentStatusDefault with default headers values
func NewGetDeviceComponentStatusDefault(code int) *GetDeviceComponentStatusDefault {
	return &GetDeviceComponentStatusDefault{
		_statusCode: code,
	}
}

/*GetDeviceComponentStatusDefault handles this case with default header values.

Unexpected error
*/
type GetDeviceComponentStatusDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get device component status default response
func (o *GetDeviceComponentStatusDefault) Code() int {
	return o._statusCode
}

func (o *GetDeviceComponentStatusDefault) Error() string {
	return fmt.Sprintf("[GET /devices/{deviceId}/components/{componentId}/status][%d] getDeviceComponentStatus default  %+v", o._statusCode, o.Payload)
}

func (o *GetDeviceComponentStatusDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}