// Code generated by go-swagger; DO NOT EDIT.

package apps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/SmartThingsOSS/smartapp-go/pkg/smartthings/models"
)

// GetAppTagsReader is a Reader for the GetAppTags structure.
type GetAppTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAppTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAppTagsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetAppTagsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetAppTagsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetAppTagsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewGetAppTagsUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewGetAppTagsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetAppTagsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAppTagsOK creates a GetAppTagsOK with default headers values
func NewGetAppTagsOK() *GetAppTagsOK {
	return &GetAppTagsOK{}
}

/*GetAppTagsOK handles this case with default header values.

A response containing an App's tags.
*/
type GetAppTagsOK struct {
	Payload *models.GetTagsResponse
}

func (o *GetAppTagsOK) Error() string {
	return fmt.Sprintf("[GET /apps/{appNameOrId}/tags][%d] getAppTagsOK  %+v", 200, o.Payload)
}

func (o *GetAppTagsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetTagsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppTagsBadRequest creates a GetAppTagsBadRequest with default headers values
func NewGetAppTagsBadRequest() *GetAppTagsBadRequest {
	return &GetAppTagsBadRequest{}
}

/*GetAppTagsBadRequest handles this case with default header values.

Bad request
*/
type GetAppTagsBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *GetAppTagsBadRequest) Error() string {
	return fmt.Sprintf("[GET /apps/{appNameOrId}/tags][%d] getAppTagsBadRequest  %+v", 400, o.Payload)
}

func (o *GetAppTagsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppTagsUnauthorized creates a GetAppTagsUnauthorized with default headers values
func NewGetAppTagsUnauthorized() *GetAppTagsUnauthorized {
	return &GetAppTagsUnauthorized{}
}

/*GetAppTagsUnauthorized handles this case with default header values.

Unauthorized
*/
type GetAppTagsUnauthorized struct {
}

func (o *GetAppTagsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apps/{appNameOrId}/tags][%d] getAppTagsUnauthorized ", 401)
}

func (o *GetAppTagsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAppTagsForbidden creates a GetAppTagsForbidden with default headers values
func NewGetAppTagsForbidden() *GetAppTagsForbidden {
	return &GetAppTagsForbidden{}
}

/*GetAppTagsForbidden handles this case with default header values.

Forbidden
*/
type GetAppTagsForbidden struct {
}

func (o *GetAppTagsForbidden) Error() string {
	return fmt.Sprintf("[GET /apps/{appNameOrId}/tags][%d] getAppTagsForbidden ", 403)
}

func (o *GetAppTagsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAppTagsUnprocessableEntity creates a GetAppTagsUnprocessableEntity with default headers values
func NewGetAppTagsUnprocessableEntity() *GetAppTagsUnprocessableEntity {
	return &GetAppTagsUnprocessableEntity{}
}

/*GetAppTagsUnprocessableEntity handles this case with default header values.

Unprocessable Entity
*/
type GetAppTagsUnprocessableEntity struct {
	Payload *models.ErrorResponse
}

func (o *GetAppTagsUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /apps/{appNameOrId}/tags][%d] getAppTagsUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *GetAppTagsUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppTagsTooManyRequests creates a GetAppTagsTooManyRequests with default headers values
func NewGetAppTagsTooManyRequests() *GetAppTagsTooManyRequests {
	return &GetAppTagsTooManyRequests{}
}

/*GetAppTagsTooManyRequests handles this case with default header values.

Too many requests
*/
type GetAppTagsTooManyRequests struct {
	Payload *models.ErrorResponse
}

func (o *GetAppTagsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /apps/{appNameOrId}/tags][%d] getAppTagsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetAppTagsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppTagsDefault creates a GetAppTagsDefault with default headers values
func NewGetAppTagsDefault(code int) *GetAppTagsDefault {
	return &GetAppTagsDefault{
		_statusCode: code,
	}
}

/*GetAppTagsDefault handles this case with default header values.

Unexpected error
*/
type GetAppTagsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get app tags default response
func (o *GetAppTagsDefault) Code() int {
	return o._statusCode
}

func (o *GetAppTagsDefault) Error() string {
	return fmt.Sprintf("[GET /apps/{appNameOrId}/tags][%d] getAppTags default  %+v", o._statusCode, o.Payload)
}

func (o *GetAppTagsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
