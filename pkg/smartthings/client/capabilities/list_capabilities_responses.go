// Code generated by go-swagger; DO NOT EDIT.

package capabilities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/SmartThingsOSS/smartapp-go/pkg/smartthings/models"
)

// ListCapabilitiesReader is a Reader for the ListCapabilities structure.
type ListCapabilitiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListCapabilitiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListCapabilitiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewListCapabilitiesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewListCapabilitiesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewListCapabilitiesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewListCapabilitiesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListCapabilitiesOK creates a ListCapabilitiesOK with default headers values
func NewListCapabilitiesOK() *ListCapabilitiesOK {
	return &ListCapabilitiesOK{}
}

/*ListCapabilitiesOK handles this case with default header values.

successful return of all capabilities
*/
type ListCapabilitiesOK struct {
	/*Maximum requests allowed within the rate limit window.
	 */
	XRateLimitLimit int64
	/*Remaining requests available within the window.
	 */
	XRateLimitRemaining int64
	/*Time in milliseconds until the current window expires.
	 */
	XRateLimitReset int64

	Payload *models.PagedCapabilities
}

func (o *ListCapabilitiesOK) Error() string {
	return fmt.Sprintf("[GET /capabilities][%d] listCapabilitiesOK  %+v", 200, o.Payload)
}

func (o *ListCapabilitiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-RateLimit-Limit
	xRateLimitLimit, err := swag.ConvertInt64(response.GetHeader("X-RateLimit-Limit"))
	if err != nil {
		return errors.InvalidType("X-RateLimit-Limit", "header", "int64", response.GetHeader("X-RateLimit-Limit"))
	}
	o.XRateLimitLimit = xRateLimitLimit

	// response header X-RateLimit-Remaining
	xRateLimitRemaining, err := swag.ConvertInt64(response.GetHeader("X-RateLimit-Remaining"))
	if err != nil {
		return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", response.GetHeader("X-RateLimit-Remaining"))
	}
	o.XRateLimitRemaining = xRateLimitRemaining

	// response header X-RateLimit-Reset
	xRateLimitReset, err := swag.ConvertInt64(response.GetHeader("X-RateLimit-Reset"))
	if err != nil {
		return errors.InvalidType("X-RateLimit-Reset", "header", "int64", response.GetHeader("X-RateLimit-Reset"))
	}
	o.XRateLimitReset = xRateLimitReset

	o.Payload = new(models.PagedCapabilities)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCapabilitiesUnauthorized creates a ListCapabilitiesUnauthorized with default headers values
func NewListCapabilitiesUnauthorized() *ListCapabilitiesUnauthorized {
	return &ListCapabilitiesUnauthorized{}
}

/*ListCapabilitiesUnauthorized handles this case with default header values.

Unauthorized
*/
type ListCapabilitiesUnauthorized struct {
}

func (o *ListCapabilitiesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /capabilities][%d] listCapabilitiesUnauthorized ", 401)
}

func (o *ListCapabilitiesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListCapabilitiesForbidden creates a ListCapabilitiesForbidden with default headers values
func NewListCapabilitiesForbidden() *ListCapabilitiesForbidden {
	return &ListCapabilitiesForbidden{}
}

/*ListCapabilitiesForbidden handles this case with default header values.

Forbidden
*/
type ListCapabilitiesForbidden struct {
}

func (o *ListCapabilitiesForbidden) Error() string {
	return fmt.Sprintf("[GET /capabilities][%d] listCapabilitiesForbidden ", 403)
}

func (o *ListCapabilitiesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListCapabilitiesTooManyRequests creates a ListCapabilitiesTooManyRequests with default headers values
func NewListCapabilitiesTooManyRequests() *ListCapabilitiesTooManyRequests {
	return &ListCapabilitiesTooManyRequests{}
}

/*ListCapabilitiesTooManyRequests handles this case with default header values.

Too many requests
*/
type ListCapabilitiesTooManyRequests struct {
	Payload *models.ErrorResponse
}

func (o *ListCapabilitiesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /capabilities][%d] listCapabilitiesTooManyRequests  %+v", 429, o.Payload)
}

func (o *ListCapabilitiesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCapabilitiesDefault creates a ListCapabilitiesDefault with default headers values
func NewListCapabilitiesDefault(code int) *ListCapabilitiesDefault {
	return &ListCapabilitiesDefault{
		_statusCode: code,
	}
}

/*ListCapabilitiesDefault handles this case with default header values.

Unexpected error
*/
type ListCapabilitiesDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list capabilities default response
func (o *ListCapabilitiesDefault) Code() int {
	return o._statusCode
}

func (o *ListCapabilitiesDefault) Error() string {
	return fmt.Sprintf("[GET /capabilities][%d] listCapabilities default  %+v", o._statusCode, o.Payload)
}

func (o *ListCapabilitiesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
