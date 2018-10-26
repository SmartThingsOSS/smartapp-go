// Code generated by go-swagger; DO NOT EDIT.

package modes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/SmartThingsOSS/smartapp-go/pkg/smartthings/models"
)

// GetCurrentModeReader is a Reader for the GetCurrentMode structure.
type GetCurrentModeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCurrentModeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCurrentModeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetCurrentModeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetCurrentModeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetCurrentModeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewGetCurrentModeTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetCurrentModeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCurrentModeOK creates a GetCurrentModeOK with default headers values
func NewGetCurrentModeOK() *GetCurrentModeOK {
	return &GetCurrentModeOK{}
}

/*GetCurrentModeOK handles this case with default header values.

The Mode.
*/
type GetCurrentModeOK struct {
	/*This header field describes the natural language(s) of the intended audience for the representation. This can have multiple values as per [RFC 7231, section 3.1.3.2](https://tools.ietf.org/html/rfc7231#section-3.1.3.2)
	 */
	ContentLanguage string

	Payload *models.Mode
}

func (o *GetCurrentModeOK) Error() string {
	return fmt.Sprintf("[GET /locations/{locationId}/modes/current][%d] getCurrentModeOK  %+v", 200, o.Payload)
}

func (o *GetCurrentModeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Content-Language
	o.ContentLanguage = response.GetHeader("Content-Language")

	o.Payload = new(models.Mode)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCurrentModeBadRequest creates a GetCurrentModeBadRequest with default headers values
func NewGetCurrentModeBadRequest() *GetCurrentModeBadRequest {
	return &GetCurrentModeBadRequest{}
}

/*GetCurrentModeBadRequest handles this case with default header values.

Bad request
*/
type GetCurrentModeBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *GetCurrentModeBadRequest) Error() string {
	return fmt.Sprintf("[GET /locations/{locationId}/modes/current][%d] getCurrentModeBadRequest  %+v", 400, o.Payload)
}

func (o *GetCurrentModeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCurrentModeUnauthorized creates a GetCurrentModeUnauthorized with default headers values
func NewGetCurrentModeUnauthorized() *GetCurrentModeUnauthorized {
	return &GetCurrentModeUnauthorized{}
}

/*GetCurrentModeUnauthorized handles this case with default header values.

Unauthorized
*/
type GetCurrentModeUnauthorized struct {
}

func (o *GetCurrentModeUnauthorized) Error() string {
	return fmt.Sprintf("[GET /locations/{locationId}/modes/current][%d] getCurrentModeUnauthorized ", 401)
}

func (o *GetCurrentModeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCurrentModeForbidden creates a GetCurrentModeForbidden with default headers values
func NewGetCurrentModeForbidden() *GetCurrentModeForbidden {
	return &GetCurrentModeForbidden{}
}

/*GetCurrentModeForbidden handles this case with default header values.

Forbidden
*/
type GetCurrentModeForbidden struct {
}

func (o *GetCurrentModeForbidden) Error() string {
	return fmt.Sprintf("[GET /locations/{locationId}/modes/current][%d] getCurrentModeForbidden ", 403)
}

func (o *GetCurrentModeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCurrentModeTooManyRequests creates a GetCurrentModeTooManyRequests with default headers values
func NewGetCurrentModeTooManyRequests() *GetCurrentModeTooManyRequests {
	return &GetCurrentModeTooManyRequests{}
}

/*GetCurrentModeTooManyRequests handles this case with default header values.

Too many requests
*/
type GetCurrentModeTooManyRequests struct {
	Payload *models.ErrorResponse
}

func (o *GetCurrentModeTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /locations/{locationId}/modes/current][%d] getCurrentModeTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetCurrentModeTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCurrentModeDefault creates a GetCurrentModeDefault with default headers values
func NewGetCurrentModeDefault(code int) *GetCurrentModeDefault {
	return &GetCurrentModeDefault{
		_statusCode: code,
	}
}

/*GetCurrentModeDefault handles this case with default header values.

Unexpected error
*/
type GetCurrentModeDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get current mode default response
func (o *GetCurrentModeDefault) Code() int {
	return o._statusCode
}

func (o *GetCurrentModeDefault) Error() string {
	return fmt.Sprintf("[GET /locations/{locationId}/modes/current][%d] getCurrentMode default  %+v", o._statusCode, o.Payload)
}

func (o *GetCurrentModeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
