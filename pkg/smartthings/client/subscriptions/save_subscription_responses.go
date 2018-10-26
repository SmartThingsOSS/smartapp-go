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

// SaveSubscriptionReader is a Reader for the SaveSubscription structure.
type SaveSubscriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SaveSubscriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSaveSubscriptionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSaveSubscriptionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewSaveSubscriptionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewSaveSubscriptionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewSaveSubscriptionConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewSaveSubscriptionUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewSaveSubscriptionTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewSaveSubscriptionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSaveSubscriptionOK creates a SaveSubscriptionOK with default headers values
func NewSaveSubscriptionOK() *SaveSubscriptionOK {
	return &SaveSubscriptionOK{}
}

/*SaveSubscriptionOK handles this case with default header values.

The subscription
*/
type SaveSubscriptionOK struct {
	Payload *models.Subscription
}

func (o *SaveSubscriptionOK) Error() string {
	return fmt.Sprintf("[POST /installedapps/{installedAppId}/subscriptions][%d] saveSubscriptionOK  %+v", 200, o.Payload)
}

func (o *SaveSubscriptionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Subscription)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSaveSubscriptionBadRequest creates a SaveSubscriptionBadRequest with default headers values
func NewSaveSubscriptionBadRequest() *SaveSubscriptionBadRequest {
	return &SaveSubscriptionBadRequest{}
}

/*SaveSubscriptionBadRequest handles this case with default header values.

Bad request
*/
type SaveSubscriptionBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *SaveSubscriptionBadRequest) Error() string {
	return fmt.Sprintf("[POST /installedapps/{installedAppId}/subscriptions][%d] saveSubscriptionBadRequest  %+v", 400, o.Payload)
}

func (o *SaveSubscriptionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSaveSubscriptionUnauthorized creates a SaveSubscriptionUnauthorized with default headers values
func NewSaveSubscriptionUnauthorized() *SaveSubscriptionUnauthorized {
	return &SaveSubscriptionUnauthorized{}
}

/*SaveSubscriptionUnauthorized handles this case with default header values.

Not authenticated
*/
type SaveSubscriptionUnauthorized struct {
}

func (o *SaveSubscriptionUnauthorized) Error() string {
	return fmt.Sprintf("[POST /installedapps/{installedAppId}/subscriptions][%d] saveSubscriptionUnauthorized ", 401)
}

func (o *SaveSubscriptionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSaveSubscriptionForbidden creates a SaveSubscriptionForbidden with default headers values
func NewSaveSubscriptionForbidden() *SaveSubscriptionForbidden {
	return &SaveSubscriptionForbidden{}
}

/*SaveSubscriptionForbidden handles this case with default header values.

Not authorized
*/
type SaveSubscriptionForbidden struct {
}

func (o *SaveSubscriptionForbidden) Error() string {
	return fmt.Sprintf("[POST /installedapps/{installedAppId}/subscriptions][%d] saveSubscriptionForbidden ", 403)
}

func (o *SaveSubscriptionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSaveSubscriptionConflict creates a SaveSubscriptionConflict with default headers values
func NewSaveSubscriptionConflict() *SaveSubscriptionConflict {
	return &SaveSubscriptionConflict{}
}

/*SaveSubscriptionConflict handles this case with default header values.

Conflict
*/
type SaveSubscriptionConflict struct {
	Payload *models.ErrorResponse
}

func (o *SaveSubscriptionConflict) Error() string {
	return fmt.Sprintf("[POST /installedapps/{installedAppId}/subscriptions][%d] saveSubscriptionConflict  %+v", 409, o.Payload)
}

func (o *SaveSubscriptionConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSaveSubscriptionUnprocessableEntity creates a SaveSubscriptionUnprocessableEntity with default headers values
func NewSaveSubscriptionUnprocessableEntity() *SaveSubscriptionUnprocessableEntity {
	return &SaveSubscriptionUnprocessableEntity{}
}

/*SaveSubscriptionUnprocessableEntity handles this case with default header values.

Unprocessable Entity
*/
type SaveSubscriptionUnprocessableEntity struct {
	Payload *models.ErrorResponse
}

func (o *SaveSubscriptionUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /installedapps/{installedAppId}/subscriptions][%d] saveSubscriptionUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *SaveSubscriptionUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSaveSubscriptionTooManyRequests creates a SaveSubscriptionTooManyRequests with default headers values
func NewSaveSubscriptionTooManyRequests() *SaveSubscriptionTooManyRequests {
	return &SaveSubscriptionTooManyRequests{}
}

/*SaveSubscriptionTooManyRequests handles this case with default header values.

Too many requests
*/
type SaveSubscriptionTooManyRequests struct {
	Payload *models.ErrorResponse
}

func (o *SaveSubscriptionTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /installedapps/{installedAppId}/subscriptions][%d] saveSubscriptionTooManyRequests  %+v", 429, o.Payload)
}

func (o *SaveSubscriptionTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSaveSubscriptionDefault creates a SaveSubscriptionDefault with default headers values
func NewSaveSubscriptionDefault(code int) *SaveSubscriptionDefault {
	return &SaveSubscriptionDefault{
		_statusCode: code,
	}
}

/*SaveSubscriptionDefault handles this case with default header values.

Unexpected error
*/
type SaveSubscriptionDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the save subscription default response
func (o *SaveSubscriptionDefault) Code() int {
	return o._statusCode
}

func (o *SaveSubscriptionDefault) Error() string {
	return fmt.Sprintf("[POST /installedapps/{installedAppId}/subscriptions][%d] saveSubscription default  %+v", o._statusCode, o.Payload)
}

func (o *SaveSubscriptionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
