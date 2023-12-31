// Code generated by go-swagger; DO NOT EDIT.

package platform_event_channel_member

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"vestahealthcare/models"
)

// UpdatePlatformEventChannelMemberReader is a Reader for the UpdatePlatformEventChannelMember structure.
type UpdatePlatformEventChannelMemberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdatePlatformEventChannelMemberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdatePlatformEventChannelMemberNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdatePlatformEventChannelMemberDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdatePlatformEventChannelMemberNoContent creates a UpdatePlatformEventChannelMemberNoContent with default headers values
func NewUpdatePlatformEventChannelMemberNoContent() *UpdatePlatformEventChannelMemberNoContent {
	return &UpdatePlatformEventChannelMemberNoContent{}
}

/*
UpdatePlatformEventChannelMemberNoContent describes a response with status code 204, with default header values.

successful operation
*/
type UpdatePlatformEventChannelMemberNoContent struct {
}

func (o *UpdatePlatformEventChannelMemberNoContent) Error() string {

	return fmt.Sprintf("[PATCH /tooling/sobjects/PlatformEventChannelMember/{Id}][%d] updatePlatformEventChannelMemberNoContent ", 204)
}

func (o *UpdatePlatformEventChannelMemberNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdatePlatformEventChannelMemberDefault creates a UpdatePlatformEventChannelMemberDefault with default headers values
func NewUpdatePlatformEventChannelMemberDefault(code int) *UpdatePlatformEventChannelMemberDefault {
	return &UpdatePlatformEventChannelMemberDefault{
		_statusCode: code,
	}
}

/*
UpdatePlatformEventChannelMemberDefault describes a response with status code -1, with default header values.

Error
*/
type UpdatePlatformEventChannelMemberDefault struct {
	_statusCode int

	Payload []*models.ErrorResponse
}

// Code gets the status code for the update platform event channel member default response
func (o *UpdatePlatformEventChannelMemberDefault) Code() int {
	return o._statusCode
}

func (o *UpdatePlatformEventChannelMemberDefault) Error() string {
	s := fmt.Sprintf("%+v", o.Payload)
	b, err := json.Marshal(o.Payload)
	if err == nil {
		s = string(b)
	}
	return fmt.Sprintf("[PATCH /tooling/sobjects/PlatformEventChannelMember/{Id}][%d] updatePlatformEventChannelMember default  %s", o._statusCode, s)
}
func (o *UpdatePlatformEventChannelMemberDefault) GetPayload() []*models.ErrorResponse {
	return o.Payload
}

func (o *UpdatePlatformEventChannelMemberDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
