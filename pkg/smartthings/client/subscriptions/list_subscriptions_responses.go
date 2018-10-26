// Code generated by go-swagger; DO NOT EDIT.

package subscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/SmartThingsOSS/smartapp-go/pkg/smartthings/models"
)

// ListSubscriptionsReader is a Reader for the ListSubscriptions structure.
type ListSubscriptionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListSubscriptionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListSubscriptionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewListSubscriptionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewListSubscriptionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewListSubscriptionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewListSubscriptionsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewListSubscriptionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListSubscriptionsOK creates a ListSubscriptionsOK with default headers values
func NewListSubscriptionsOK() *ListSubscriptionsOK {
	return &ListSubscriptionsOK{}
}

/*ListSubscriptionsOK handles this case with default header values.

An array of subscriptions
*/
type ListSubscriptionsOK struct {
	Payload *models.PagedSubscriptions
}

func (o *ListSubscriptionsOK) Error() string {
	return fmt.Sprintf("[GET /installedapps/{installedAppId}/subscriptions][%d] listSubscriptionsOK  %+v", 200, o.Payload)
}

func (o *ListSubscriptionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PagedSubscriptions)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSubscriptionsBadRequest creates a ListSubscriptionsBadRequest with default headers values
func NewListSubscriptionsBadRequest() *ListSubscriptionsBadRequest {
	return &ListSubscriptionsBadRequest{}
}

/*ListSubscriptionsBadRequest handles this case with default header values.

Bad request
*/
type ListSubscriptionsBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *ListSubscriptionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /installedapps/{installedAppId}/subscriptions][%d] listSubscriptionsBadRequest  %+v", 400, o.Payload)
}

func (o *ListSubscriptionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSubscriptionsUnauthorized creates a ListSubscriptionsUnauthorized with default headers values
func NewListSubscriptionsUnauthorized() *ListSubscriptionsUnauthorized {
	return &ListSubscriptionsUnauthorized{}
}

/*ListSubscriptionsUnauthorized handles this case with default header values.

Not authenticated
*/
type ListSubscriptionsUnauthorized struct {
}

func (o *ListSubscriptionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /installedapps/{installedAppId}/subscriptions][%d] listSubscriptionsUnauthorized ", 401)
}

func (o *ListSubscriptionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListSubscriptionsForbidden creates a ListSubscriptionsForbidden with default headers values
func NewListSubscriptionsForbidden() *ListSubscriptionsForbidden {
	return &ListSubscriptionsForbidden{}
}

/*ListSubscriptionsForbidden handles this case with default header values.

Not authorized
*/
type ListSubscriptionsForbidden struct {
}

func (o *ListSubscriptionsForbidden) Error() string {
	return fmt.Sprintf("[GET /installedapps/{installedAppId}/subscriptions][%d] listSubscriptionsForbidden ", 403)
}

func (o *ListSubscriptionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListSubscriptionsTooManyRequests creates a ListSubscriptionsTooManyRequests with default headers values
func NewListSubscriptionsTooManyRequests() *ListSubscriptionsTooManyRequests {
	return &ListSubscriptionsTooManyRequests{}
}

/*ListSubscriptionsTooManyRequests handles this case with default header values.

Too many requests
*/
type ListSubscriptionsTooManyRequests struct {
	Payload *models.ErrorResponse
}

func (o *ListSubscriptionsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /installedapps/{installedAppId}/subscriptions][%d] listSubscriptionsTooManyRequests  %+v", 429, o.Payload)
}

func (o *ListSubscriptionsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSubscriptionsDefault creates a ListSubscriptionsDefault with default headers values
func NewListSubscriptionsDefault(code int) *ListSubscriptionsDefault {
	return &ListSubscriptionsDefault{
		_statusCode: code,
	}
}

/*ListSubscriptionsDefault handles this case with default header values.

Unexpected error
*/
type ListSubscriptionsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list subscriptions default response
func (o *ListSubscriptionsDefault) Code() int {
	return o._statusCode
}

func (o *ListSubscriptionsDefault) Error() string {
	return fmt.Sprintf("[GET /installedapps/{installedAppId}/subscriptions][%d] listSubscriptions default  %+v", o._statusCode, o.Payload)
}

func (o *ListSubscriptionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}