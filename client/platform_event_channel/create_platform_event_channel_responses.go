// Code generated by go-swagger; DO NOT EDIT.

package platform_event_channel

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

// CreatePlatformEventChannelReader is a Reader for the CreatePlatformEventChannel structure.
type CreatePlatformEventChannelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePlatformEventChannelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreatePlatformEventChannelCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreatePlatformEventChannelDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreatePlatformEventChannelCreated creates a CreatePlatformEventChannelCreated with default headers values
func NewCreatePlatformEventChannelCreated() *CreatePlatformEventChannelCreated {
	return &CreatePlatformEventChannelCreated{}
}

/*
CreatePlatformEventChannelCreated describes a response with status code 201, with default header values.

successful operation
*/
type CreatePlatformEventChannelCreated struct {
	Payload *models.PlatformEventChannelCreateResponse
}

func (o *CreatePlatformEventChannelCreated) Error() string {
	s := fmt.Sprintf("%+v", o.Payload)
	b, err := json.Marshal(o.Payload)
	if err == nil {
		s = string(b)
	}
	return fmt.Sprintf("[POST /tooling/sobjects/PlatformEventChannel][%d] createPlatformEventChannelCreated  %s", 201, s)
}
func (o *CreatePlatformEventChannelCreated) GetPayload() *models.PlatformEventChannelCreateResponse {
	return o.Payload
}

func (o *CreatePlatformEventChannelCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PlatformEventChannelCreateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePlatformEventChannelDefault creates a CreatePlatformEventChannelDefault with default headers values
func NewCreatePlatformEventChannelDefault(code int) *CreatePlatformEventChannelDefault {
	return &CreatePlatformEventChannelDefault{
		_statusCode: code,
	}
}

/*
CreatePlatformEventChannelDefault describes a response with status code -1, with default header values.

Error
*/
type CreatePlatformEventChannelDefault struct {
	_statusCode int

	Payload []*models.ErrorResponse
}

// Code gets the status code for the create platform event channel default response
func (o *CreatePlatformEventChannelDefault) Code() int {
	return o._statusCode
}

func (o *CreatePlatformEventChannelDefault) Error() string {
	s := fmt.Sprintf("%+v", o.Payload)
	b, err := json.Marshal(o.Payload)
	if err == nil {
		s = string(b)
	}
	return fmt.Sprintf("[POST /tooling/sobjects/PlatformEventChannel][%d] createPlatformEventChannel default  %s", o._statusCode, s)
}
func (o *CreatePlatformEventChannelDefault) GetPayload() []*models.ErrorResponse {
	return o.Payload
}

func (o *CreatePlatformEventChannelDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
