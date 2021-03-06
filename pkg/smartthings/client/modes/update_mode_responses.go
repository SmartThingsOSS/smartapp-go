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

// UpdateModeReader is a Reader for the UpdateMode structure.
type UpdateModeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateModeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateModeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUpdateModeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewUpdateModeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewUpdateModeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewUpdateModeUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewUpdateModeTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewUpdateModeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateModeOK creates a UpdateModeOK with default headers values
func NewUpdateModeOK() *UpdateModeOK {
	return &UpdateModeOK{}
}

/*UpdateModeOK handles this case with default header values.

Created successfully.
*/
type UpdateModeOK struct {
	Payload *models.Mode
}

func (o *UpdateModeOK) Error() string {
	return fmt.Sprintf("[PUT /locations/{locationId}/modes/{modeId}][%d] updateModeOK  %+v", 200, o.Payload)
}

func (o *UpdateModeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Mode)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateModeBadRequest creates a UpdateModeBadRequest with default headers values
func NewUpdateModeBadRequest() *UpdateModeBadRequest {
	return &UpdateModeBadRequest{}
}

/*UpdateModeBadRequest handles this case with default header values.

Bad request
*/
type UpdateModeBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *UpdateModeBadRequest) Error() string {
	return fmt.Sprintf("[PUT /locations/{locationId}/modes/{modeId}][%d] updateModeBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateModeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateModeUnauthorized creates a UpdateModeUnauthorized with default headers values
func NewUpdateModeUnauthorized() *UpdateModeUnauthorized {
	return &UpdateModeUnauthorized{}
}

/*UpdateModeUnauthorized handles this case with default header values.

Unauthorized
*/
type UpdateModeUnauthorized struct {
}

func (o *UpdateModeUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /locations/{locationId}/modes/{modeId}][%d] updateModeUnauthorized ", 401)
}

func (o *UpdateModeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateModeForbidden creates a UpdateModeForbidden with default headers values
func NewUpdateModeForbidden() *UpdateModeForbidden {
	return &UpdateModeForbidden{}
}

/*UpdateModeForbidden handles this case with default header values.

Forbidden
*/
type UpdateModeForbidden struct {
}

func (o *UpdateModeForbidden) Error() string {
	return fmt.Sprintf("[PUT /locations/{locationId}/modes/{modeId}][%d] updateModeForbidden ", 403)
}

func (o *UpdateModeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateModeUnprocessableEntity creates a UpdateModeUnprocessableEntity with default headers values
func NewUpdateModeUnprocessableEntity() *UpdateModeUnprocessableEntity {
	return &UpdateModeUnprocessableEntity{}
}

/*UpdateModeUnprocessableEntity handles this case with default header values.

Unprocessable Entity
*/
type UpdateModeUnprocessableEntity struct {
	Payload *models.ErrorResponse
}

func (o *UpdateModeUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /locations/{locationId}/modes/{modeId}][%d] updateModeUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *UpdateModeUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateModeTooManyRequests creates a UpdateModeTooManyRequests with default headers values
func NewUpdateModeTooManyRequests() *UpdateModeTooManyRequests {
	return &UpdateModeTooManyRequests{}
}

/*UpdateModeTooManyRequests handles this case with default header values.

Too many requests
*/
type UpdateModeTooManyRequests struct {
	Payload *models.ErrorResponse
}

func (o *UpdateModeTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /locations/{locationId}/modes/{modeId}][%d] updateModeTooManyRequests  %+v", 429, o.Payload)
}

func (o *UpdateModeTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateModeDefault creates a UpdateModeDefault with default headers values
func NewUpdateModeDefault(code int) *UpdateModeDefault {
	return &UpdateModeDefault{
		_statusCode: code,
	}
}

/*UpdateModeDefault handles this case with default header values.

Unexpected error
*/
type UpdateModeDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the update mode default response
func (o *UpdateModeDefault) Code() int {
	return o._statusCode
}

func (o *UpdateModeDefault) Error() string {
	return fmt.Sprintf("[PUT /locations/{locationId}/modes/{modeId}][%d] updateMode default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateModeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
