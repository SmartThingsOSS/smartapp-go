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

// UpdateInstalledAppTagsReader is a Reader for the UpdateInstalledAppTags structure.
type UpdateInstalledAppTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateInstalledAppTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateInstalledAppTagsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUpdateInstalledAppTagsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewUpdateInstalledAppTagsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewUpdateInstalledAppTagsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewUpdateInstalledAppTagsUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewUpdateInstalledAppTagsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewUpdateInstalledAppTagsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateInstalledAppTagsOK creates a UpdateInstalledAppTagsOK with default headers values
func NewUpdateInstalledAppTagsOK() *UpdateInstalledAppTagsOK {
	return &UpdateInstalledAppTagsOK{}
}

/*UpdateInstalledAppTagsOK handles this case with default header values.

A response containing an InstalledApp's tags.
*/
type UpdateInstalledAppTagsOK struct {
	Payload *models.UpdateTagsResponse
}

func (o *UpdateInstalledAppTagsOK) Error() string {
	return fmt.Sprintf("[PUT /installedapps/{installedAppId}/tags][%d] updateInstalledAppTagsOK  %+v", 200, o.Payload)
}

func (o *UpdateInstalledAppTagsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpdateTagsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateInstalledAppTagsBadRequest creates a UpdateInstalledAppTagsBadRequest with default headers values
func NewUpdateInstalledAppTagsBadRequest() *UpdateInstalledAppTagsBadRequest {
	return &UpdateInstalledAppTagsBadRequest{}
}

/*UpdateInstalledAppTagsBadRequest handles this case with default header values.

Bad request
*/
type UpdateInstalledAppTagsBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *UpdateInstalledAppTagsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /installedapps/{installedAppId}/tags][%d] updateInstalledAppTagsBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateInstalledAppTagsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateInstalledAppTagsUnauthorized creates a UpdateInstalledAppTagsUnauthorized with default headers values
func NewUpdateInstalledAppTagsUnauthorized() *UpdateInstalledAppTagsUnauthorized {
	return &UpdateInstalledAppTagsUnauthorized{}
}

/*UpdateInstalledAppTagsUnauthorized handles this case with default header values.

Unauthorized
*/
type UpdateInstalledAppTagsUnauthorized struct {
}

func (o *UpdateInstalledAppTagsUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /installedapps/{installedAppId}/tags][%d] updateInstalledAppTagsUnauthorized ", 401)
}

func (o *UpdateInstalledAppTagsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateInstalledAppTagsForbidden creates a UpdateInstalledAppTagsForbidden with default headers values
func NewUpdateInstalledAppTagsForbidden() *UpdateInstalledAppTagsForbidden {
	return &UpdateInstalledAppTagsForbidden{}
}

/*UpdateInstalledAppTagsForbidden handles this case with default header values.

Forbidden
*/
type UpdateInstalledAppTagsForbidden struct {
}

func (o *UpdateInstalledAppTagsForbidden) Error() string {
	return fmt.Sprintf("[PUT /installedapps/{installedAppId}/tags][%d] updateInstalledAppTagsForbidden ", 403)
}

func (o *UpdateInstalledAppTagsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateInstalledAppTagsUnprocessableEntity creates a UpdateInstalledAppTagsUnprocessableEntity with default headers values
func NewUpdateInstalledAppTagsUnprocessableEntity() *UpdateInstalledAppTagsUnprocessableEntity {
	return &UpdateInstalledAppTagsUnprocessableEntity{}
}

/*UpdateInstalledAppTagsUnprocessableEntity handles this case with default header values.

Unprocessable Entity
*/
type UpdateInstalledAppTagsUnprocessableEntity struct {
	Payload *models.ErrorResponse
}

func (o *UpdateInstalledAppTagsUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /installedapps/{installedAppId}/tags][%d] updateInstalledAppTagsUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *UpdateInstalledAppTagsUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateInstalledAppTagsTooManyRequests creates a UpdateInstalledAppTagsTooManyRequests with default headers values
func NewUpdateInstalledAppTagsTooManyRequests() *UpdateInstalledAppTagsTooManyRequests {
	return &UpdateInstalledAppTagsTooManyRequests{}
}

/*UpdateInstalledAppTagsTooManyRequests handles this case with default header values.

Too many requests
*/
type UpdateInstalledAppTagsTooManyRequests struct {
	Payload *models.ErrorResponse
}

func (o *UpdateInstalledAppTagsTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /installedapps/{installedAppId}/tags][%d] updateInstalledAppTagsTooManyRequests  %+v", 429, o.Payload)
}

func (o *UpdateInstalledAppTagsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateInstalledAppTagsDefault creates a UpdateInstalledAppTagsDefault with default headers values
func NewUpdateInstalledAppTagsDefault(code int) *UpdateInstalledAppTagsDefault {
	return &UpdateInstalledAppTagsDefault{
		_statusCode: code,
	}
}

/*UpdateInstalledAppTagsDefault handles this case with default header values.

Unexpected error
*/
type UpdateInstalledAppTagsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the update installed app tags default response
func (o *UpdateInstalledAppTagsDefault) Code() int {
	return o._statusCode
}

func (o *UpdateInstalledAppTagsDefault) Error() string {
	return fmt.Sprintf("[PUT /installedapps/{installedAppId}/tags][%d] updateInstalledAppTags default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateInstalledAppTagsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
