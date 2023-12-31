// Code generated by go-swagger; DO NOT EDIT.

package event_relay_feedback_lookup

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

// GetQueryReader is a Reader for the GetQuery structure.
type GetQueryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetQueryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetQueryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetQueryDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetQueryOK creates a GetQueryOK with default headers values
func NewGetQueryOK() *GetQueryOK {
	return &GetQueryOK{}
}

/*
GetQueryOK describes a response with status code 200, with default header values.

successful operation
*/
type GetQueryOK struct {
	Payload *models.EventRelayFeedbackLookup
}

func (o *GetQueryOK) Error() string {
	s := fmt.Sprintf("%+v", o.Payload)
	b, err := json.Marshal(o.Payload)
	if err == nil {
		s = string(b)
	}
	return fmt.Sprintf("[GET /query][%d] getQueryOK  %s", 200, s)
}
func (o *GetQueryOK) GetPayload() *models.EventRelayFeedbackLookup {
	return o.Payload
}

func (o *GetQueryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EventRelayFeedbackLookup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQueryDefault creates a GetQueryDefault with default headers values
func NewGetQueryDefault(code int) *GetQueryDefault {
	return &GetQueryDefault{
		_statusCode: code,
	}
}

/*
GetQueryDefault describes a response with status code -1, with default header values.

Error
*/
type GetQueryDefault struct {
	_statusCode int

	Payload []*models.ErrorResponse
}

// Code gets the status code for the get query default response
func (o *GetQueryDefault) Code() int {
	return o._statusCode
}

func (o *GetQueryDefault) Error() string {
	s := fmt.Sprintf("%+v", o.Payload)
	b, err := json.Marshal(o.Payload)
	if err == nil {
		s = string(b)
	}
	return fmt.Sprintf("[GET /query][%d] getQuery default  %s", o._statusCode, s)
}
func (o *GetQueryDefault) GetPayload() []*models.ErrorResponse {
	return o.Payload
}

func (o *GetQueryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
