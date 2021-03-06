// Code generated by go-swagger; DO NOT EDIT.

package deviceprofiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/SmartThingsOSS/smartapp-go/pkg/smartthings/models"
)

// GetDeviceProfileLocalesReader is a Reader for the GetDeviceProfileLocales structure.
type GetDeviceProfileLocalesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeviceProfileLocalesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDeviceProfileLocalesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetDeviceProfileLocalesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetDeviceProfileLocalesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetDeviceProfileLocalesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewGetDeviceProfileLocalesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetDeviceProfileLocalesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDeviceProfileLocalesOK creates a GetDeviceProfileLocalesOK with default headers values
func NewGetDeviceProfileLocalesOK() *GetDeviceProfileLocalesOK {
	return &GetDeviceProfileLocalesOK{}
}

/*GetDeviceProfileLocalesOK handles this case with default header values.

The locale tag list.
*/
type GetDeviceProfileLocalesOK struct {
	Payload *models.DeviceProfileLocaleTranslations
}

func (o *GetDeviceProfileLocalesOK) Error() string {
	return fmt.Sprintf("[GET /deviceprofiles/{deviceProfileId}/i18n][%d] getDeviceProfileLocalesOK  %+v", 200, o.Payload)
}

func (o *GetDeviceProfileLocalesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceProfileLocaleTranslations)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeviceProfileLocalesUnauthorized creates a GetDeviceProfileLocalesUnauthorized with default headers values
func NewGetDeviceProfileLocalesUnauthorized() *GetDeviceProfileLocalesUnauthorized {
	return &GetDeviceProfileLocalesUnauthorized{}
}

/*GetDeviceProfileLocalesUnauthorized handles this case with default header values.

Unauthorized
*/
type GetDeviceProfileLocalesUnauthorized struct {
}

func (o *GetDeviceProfileLocalesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /deviceprofiles/{deviceProfileId}/i18n][%d] getDeviceProfileLocalesUnauthorized ", 401)
}

func (o *GetDeviceProfileLocalesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDeviceProfileLocalesForbidden creates a GetDeviceProfileLocalesForbidden with default headers values
func NewGetDeviceProfileLocalesForbidden() *GetDeviceProfileLocalesForbidden {
	return &GetDeviceProfileLocalesForbidden{}
}

/*GetDeviceProfileLocalesForbidden handles this case with default header values.

Forbidden
*/
type GetDeviceProfileLocalesForbidden struct {
}

func (o *GetDeviceProfileLocalesForbidden) Error() string {
	return fmt.Sprintf("[GET /deviceprofiles/{deviceProfileId}/i18n][%d] getDeviceProfileLocalesForbidden ", 403)
}

func (o *GetDeviceProfileLocalesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDeviceProfileLocalesNotFound creates a GetDeviceProfileLocalesNotFound with default headers values
func NewGetDeviceProfileLocalesNotFound() *GetDeviceProfileLocalesNotFound {
	return &GetDeviceProfileLocalesNotFound{}
}

/*GetDeviceProfileLocalesNotFound handles this case with default header values.

Not found
*/
type GetDeviceProfileLocalesNotFound struct {
	Payload *models.ErrorResponse
}

func (o *GetDeviceProfileLocalesNotFound) Error() string {
	return fmt.Sprintf("[GET /deviceprofiles/{deviceProfileId}/i18n][%d] getDeviceProfileLocalesNotFound  %+v", 404, o.Payload)
}

func (o *GetDeviceProfileLocalesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeviceProfileLocalesTooManyRequests creates a GetDeviceProfileLocalesTooManyRequests with default headers values
func NewGetDeviceProfileLocalesTooManyRequests() *GetDeviceProfileLocalesTooManyRequests {
	return &GetDeviceProfileLocalesTooManyRequests{}
}

/*GetDeviceProfileLocalesTooManyRequests handles this case with default header values.

Too many requests
*/
type GetDeviceProfileLocalesTooManyRequests struct {
	Payload *models.ErrorResponse
}

func (o *GetDeviceProfileLocalesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /deviceprofiles/{deviceProfileId}/i18n][%d] getDeviceProfileLocalesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetDeviceProfileLocalesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeviceProfileLocalesDefault creates a GetDeviceProfileLocalesDefault with default headers values
func NewGetDeviceProfileLocalesDefault(code int) *GetDeviceProfileLocalesDefault {
	return &GetDeviceProfileLocalesDefault{
		_statusCode: code,
	}
}

/*GetDeviceProfileLocalesDefault handles this case with default header values.

Unexpected error
*/
type GetDeviceProfileLocalesDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get device profile locales default response
func (o *GetDeviceProfileLocalesDefault) Code() int {
	return o._statusCode
}

func (o *GetDeviceProfileLocalesDefault) Error() string {
	return fmt.Sprintf("[GET /deviceprofiles/{deviceProfileId}/i18n][%d] getDeviceProfileLocales default  %+v", o._statusCode, o.Payload)
}

func (o *GetDeviceProfileLocalesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
