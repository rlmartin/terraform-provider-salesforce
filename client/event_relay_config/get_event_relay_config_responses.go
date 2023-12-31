// Code generated by go-swagger; DO NOT EDIT.

package event_relay_config

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

// GetEventRelayConfigReader is a Reader for the GetEventRelayConfig structure.
type GetEventRelayConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEventRelayConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEventRelayConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetEventRelayConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetEventRelayConfigOK creates a GetEventRelayConfigOK with default headers values
func NewGetEventRelayConfigOK() *GetEventRelayConfigOK {
	return &GetEventRelayConfigOK{}
}

/*
GetEventRelayConfigOK describes a response with status code 200, with default header values.

successful operation
*/
type GetEventRelayConfigOK struct {
	Payload *models.EventRelayConfig
}

func (o *GetEventRelayConfigOK) Error() string {
	s := fmt.Sprintf("%+v", o.Payload)
	b, err := json.Marshal(o.Payload)
	if err == nil {
		s = string(b)
	}
	return fmt.Sprintf("[GET /tooling/sobjects/EventRelayConfig/{Id}][%d] getEventRelayConfigOK  %s", 200, s)
}
func (o *GetEventRelayConfigOK) GetPayload() *models.EventRelayConfig {
	return o.Payload
}

func (o *GetEventRelayConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EventRelayConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEventRelayConfigDefault creates a GetEventRelayConfigDefault with default headers values
func NewGetEventRelayConfigDefault(code int) *GetEventRelayConfigDefault {
	return &GetEventRelayConfigDefault{
		_statusCode: code,
	}
}

/*
GetEventRelayConfigDefault describes a response with status code -1, with default header values.

Error
*/
type GetEventRelayConfigDefault struct {
	_statusCode int

	Payload []*models.ErrorResponse
}

// Code gets the status code for the get event relay config default response
func (o *GetEventRelayConfigDefault) Code() int {
	return o._statusCode
}

func (o *GetEventRelayConfigDefault) Error() string {
	s := fmt.Sprintf("%+v", o.Payload)
	b, err := json.Marshal(o.Payload)
	if err == nil {
		s = string(b)
	}
	return fmt.Sprintf("[GET /tooling/sobjects/EventRelayConfig/{Id}][%d] getEventRelayConfig default  %s", o._statusCode, s)
}
func (o *GetEventRelayConfigDefault) GetPayload() []*models.ErrorResponse {
	return o.Payload
}

func (o *GetEventRelayConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
