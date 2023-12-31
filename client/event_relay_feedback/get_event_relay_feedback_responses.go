// Code generated by go-swagger; DO NOT EDIT.

package event_relay_feedback

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

// GetEventRelayFeedbackReader is a Reader for the GetEventRelayFeedback structure.
type GetEventRelayFeedbackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEventRelayFeedbackReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEventRelayFeedbackOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetEventRelayFeedbackDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetEventRelayFeedbackOK creates a GetEventRelayFeedbackOK with default headers values
func NewGetEventRelayFeedbackOK() *GetEventRelayFeedbackOK {
	return &GetEventRelayFeedbackOK{}
}

/*
GetEventRelayFeedbackOK describes a response with status code 200, with default header values.

successful operation
*/
type GetEventRelayFeedbackOK struct {
	Payload *models.EventRelayFeedback
}

func (o *GetEventRelayFeedbackOK) Error() string {
	s := fmt.Sprintf("%+v", o.Payload)
	b, err := json.Marshal(o.Payload)
	if err == nil {
		s = string(b)
	}
	return fmt.Sprintf("[GET /sobjects/EventRelayFeedback/{Id}][%d] getEventRelayFeedbackOK  %s", 200, s)
}
func (o *GetEventRelayFeedbackOK) GetPayload() *models.EventRelayFeedback {
	return o.Payload
}

func (o *GetEventRelayFeedbackOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EventRelayFeedback)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEventRelayFeedbackDefault creates a GetEventRelayFeedbackDefault with default headers values
func NewGetEventRelayFeedbackDefault(code int) *GetEventRelayFeedbackDefault {
	return &GetEventRelayFeedbackDefault{
		_statusCode: code,
	}
}

/*
GetEventRelayFeedbackDefault describes a response with status code -1, with default header values.

Error
*/
type GetEventRelayFeedbackDefault struct {
	_statusCode int

	Payload []*models.ErrorResponse
}

// Code gets the status code for the get event relay feedback default response
func (o *GetEventRelayFeedbackDefault) Code() int {
	return o._statusCode
}

func (o *GetEventRelayFeedbackDefault) Error() string {
	s := fmt.Sprintf("%+v", o.Payload)
	b, err := json.Marshal(o.Payload)
	if err == nil {
		s = string(b)
	}
	return fmt.Sprintf("[GET /sobjects/EventRelayFeedback/{Id}][%d] getEventRelayFeedback default  %s", o._statusCode, s)
}
func (o *GetEventRelayFeedbackDefault) GetPayload() []*models.ErrorResponse {
	return o.Payload
}

func (o *GetEventRelayFeedbackDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
