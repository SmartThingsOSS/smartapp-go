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

// GetModesReader is a Reader for the GetModes structure.
type GetModesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetModesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetModesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetModesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetModesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetModesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewGetModesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetModesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetModesOK creates a GetModesOK with default headers values
func NewGetModesOK() *GetModesOK {
	return &GetModesOK{}
}

/*GetModesOK handles this case with default header values.

A List of Modes.
*/
type GetModesOK struct {
	/*This header field describes the natural language(s) of the intended audience for the representation. This can have multiple values as per [RFC 7231, section 3.1.3.2](https://tools.ietf.org/html/rfc7231#section-3.1.3.2)
	 */
	ContentLanguage string

	Payload *models.ModesResponse
}

func (o *GetModesOK) Error() string {
	return fmt.Sprintf("[GET /locations/{locationId}/modes][%d] getModesOK  %+v", 200, o.Payload)
}

func (o *GetModesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Content-Language
	o.ContentLanguage = response.GetHeader("Content-Language")

	o.Payload = new(models.ModesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetModesBadRequest creates a GetModesBadRequest with default headers values
func NewGetModesBadRequest() *GetModesBadRequest {
	return &GetModesBadRequest{}
}

/*GetModesBadRequest handles this case with default header values.

Bad request
*/
type GetModesBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *GetModesBadRequest) Error() string {
	return fmt.Sprintf("[GET /locations/{locationId}/modes][%d] getModesBadRequest  %+v", 400, o.Payload)
}

func (o *GetModesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetModesUnauthorized creates a GetModesUnauthorized with default headers values
func NewGetModesUnauthorized() *GetModesUnauthorized {
	return &GetModesUnauthorized{}
}

/*GetModesUnauthorized handles this case with default header values.

Unauthorized
*/
type GetModesUnauthorized struct {
}

func (o *GetModesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /locations/{locationId}/modes][%d] getModesUnauthorized ", 401)
}

func (o *GetModesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetModesForbidden creates a GetModesForbidden with default headers values
func NewGetModesForbidden() *GetModesForbidden {
	return &GetModesForbidden{}
}

/*GetModesForbidden handles this case with default header values.

Forbidden
*/
type GetModesForbidden struct {
}

func (o *GetModesForbidden) Error() string {
	return fmt.Sprintf("[GET /locations/{locationId}/modes][%d] getModesForbidden ", 403)
}

func (o *GetModesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetModesTooManyRequests creates a GetModesTooManyRequests with default headers values
func NewGetModesTooManyRequests() *GetModesTooManyRequests {
	return &GetModesTooManyRequests{}
}

/*GetModesTooManyRequests handles this case with default header values.

Too many requests
*/
type GetModesTooManyRequests struct {
	Payload *models.ErrorResponse
}

func (o *GetModesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /locations/{locationId}/modes][%d] getModesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetModesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetModesDefault creates a GetModesDefault with default headers values
func NewGetModesDefault(code int) *GetModesDefault {
	return &GetModesDefault{
		_statusCode: code,
	}
}

/*GetModesDefault handles this case with default header values.

Unexpected error
*/
type GetModesDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get modes default response
func (o *GetModesDefault) Code() int {
	return o._statusCode
}

func (o *GetModesDefault) Error() string {
	return fmt.Sprintf("[GET /locations/{locationId}/modes][%d] getModes default  %+v", o._statusCode, o.Payload)
}

func (o *GetModesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
