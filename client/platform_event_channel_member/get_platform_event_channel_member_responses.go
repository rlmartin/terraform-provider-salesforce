// Code generated by go-swagger; DO NOT EDIT.

package platform_event_channel_member

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"vestahealthcare/models"
)

// GetPlatformEventChannelMemberReader is a Reader for the GetPlatformEventChannelMember structure.
type GetPlatformEventChannelMemberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPlatformEventChannelMemberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPlatformEventChannelMemberOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetPlatformEventChannelMemberDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetPlatformEventChannelMemberOK creates a GetPlatformEventChannelMemberOK with default headers values
func NewGetPlatformEventChannelMemberOK() *GetPlatformEventChannelMemberOK {
	return &GetPlatformEventChannelMemberOK{}
}

/*
GetPlatformEventChannelMemberOK describes a response with status code 200, with default header values.

successful operation
*/
type GetPlatformEventChannelMemberOK struct {
	Payload *models.PlatformEventChannelMember
}

func (o *GetPlatformEventChannelMemberOK) Error() string {
	return fmt.Sprintf("[GET /PlatformEventChannelMember/{Id}][%d] getPlatformEventChannelMemberOK  %+v", 200, o.Payload)
}
func (o *GetPlatformEventChannelMemberOK) GetPayload() *models.PlatformEventChannelMember {
	return o.Payload
}

func (o *GetPlatformEventChannelMemberOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PlatformEventChannelMember)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPlatformEventChannelMemberDefault creates a GetPlatformEventChannelMemberDefault with default headers values
func NewGetPlatformEventChannelMemberDefault(code int) *GetPlatformEventChannelMemberDefault {
	return &GetPlatformEventChannelMemberDefault{
		_statusCode: code,
	}
}

/*
GetPlatformEventChannelMemberDefault describes a response with status code -1, with default header values.

Error
*/
type GetPlatformEventChannelMemberDefault struct {
	_statusCode int

	Payload []*models.ErrorResponse
}

// Code gets the status code for the get platform event channel member default response
func (o *GetPlatformEventChannelMemberDefault) Code() int {
	return o._statusCode
}

func (o *GetPlatformEventChannelMemberDefault) Error() string {
	return fmt.Sprintf("[GET /PlatformEventChannelMember/{Id}][%d] getPlatformEventChannelMember default  %+v", o._statusCode, o.Payload)
}
func (o *GetPlatformEventChannelMemberDefault) GetPayload() []*models.ErrorResponse {
	return o.Payload
}

func (o *GetPlatformEventChannelMemberDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
