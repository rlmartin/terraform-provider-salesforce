// Code generated by go-swagger; DO NOT EDIT.

package event_relay_config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"vestahealthcare/models"
)

// CreateEventRelayConfigReader is a Reader for the CreateEventRelayConfig structure.
type CreateEventRelayConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateEventRelayConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateEventRelayConfigCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateEventRelayConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateEventRelayConfigCreated creates a CreateEventRelayConfigCreated with default headers values
func NewCreateEventRelayConfigCreated() *CreateEventRelayConfigCreated {
	return &CreateEventRelayConfigCreated{}
}

/*
CreateEventRelayConfigCreated describes a response with status code 201, with default header values.

successful operation
*/
type CreateEventRelayConfigCreated struct {
	Payload *models.EventRelayConfigCreateResponse
}

func (o *CreateEventRelayConfigCreated) Error() string {
	return fmt.Sprintf("[POST /EventRelayConfig][%d] createEventRelayConfigCreated  %+v", 201, o.Payload)
}
func (o *CreateEventRelayConfigCreated) GetPayload() *models.EventRelayConfigCreateResponse {
	return o.Payload
}

func (o *CreateEventRelayConfigCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EventRelayConfigCreateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEventRelayConfigDefault creates a CreateEventRelayConfigDefault with default headers values
func NewCreateEventRelayConfigDefault(code int) *CreateEventRelayConfigDefault {
	return &CreateEventRelayConfigDefault{
		_statusCode: code,
	}
}

/*
CreateEventRelayConfigDefault describes a response with status code -1, with default header values.

Error
*/
type CreateEventRelayConfigDefault struct {
	_statusCode int

	Payload []*models.ErrorResponse
}

// Code gets the status code for the create event relay config default response
func (o *CreateEventRelayConfigDefault) Code() int {
	return o._statusCode
}

func (o *CreateEventRelayConfigDefault) Error() string {
	return fmt.Sprintf("[POST /EventRelayConfig][%d] createEventRelayConfig default  %+v", o._statusCode, o.Payload)
}
func (o *CreateEventRelayConfigDefault) GetPayload() []*models.ErrorResponse {
	return o.Payload
}

func (o *CreateEventRelayConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
