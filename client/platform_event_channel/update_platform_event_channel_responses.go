// Code generated by go-swagger; DO NOT EDIT.

package platform_event_channel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"vestahealthcare/models"
)

// UpdatePlatformEventChannelReader is a Reader for the UpdatePlatformEventChannel structure.
type UpdatePlatformEventChannelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdatePlatformEventChannelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdatePlatformEventChannelNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdatePlatformEventChannelDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdatePlatformEventChannelNoContent creates a UpdatePlatformEventChannelNoContent with default headers values
func NewUpdatePlatformEventChannelNoContent() *UpdatePlatformEventChannelNoContent {
	return &UpdatePlatformEventChannelNoContent{}
}

/*
UpdatePlatformEventChannelNoContent describes a response with status code 204, with default header values.

successful operation
*/
type UpdatePlatformEventChannelNoContent struct {
}

func (o *UpdatePlatformEventChannelNoContent) Error() string {
	return fmt.Sprintf("[PATCH /PlatformEventChannel/{Id}][%d] updatePlatformEventChannelNoContent ", 204)
}

func (o *UpdatePlatformEventChannelNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdatePlatformEventChannelDefault creates a UpdatePlatformEventChannelDefault with default headers values
func NewUpdatePlatformEventChannelDefault(code int) *UpdatePlatformEventChannelDefault {
	return &UpdatePlatformEventChannelDefault{
		_statusCode: code,
	}
}

/*
UpdatePlatformEventChannelDefault describes a response with status code -1, with default header values.

Error
*/
type UpdatePlatformEventChannelDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the update platform event channel default response
func (o *UpdatePlatformEventChannelDefault) Code() int {
	return o._statusCode
}

func (o *UpdatePlatformEventChannelDefault) Error() string {
	return fmt.Sprintf("[PATCH /PlatformEventChannel/{Id}][%d] updatePlatformEventChannel default  %+v", o._statusCode, o.Payload)
}
func (o *UpdatePlatformEventChannelDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdatePlatformEventChannelDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
