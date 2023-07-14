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

// DeletePlatformEventChannelMemberReader is a Reader for the DeletePlatformEventChannelMember structure.
type DeletePlatformEventChannelMemberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePlatformEventChannelMemberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeletePlatformEventChannelMemberNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeletePlatformEventChannelMemberDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeletePlatformEventChannelMemberNoContent creates a DeletePlatformEventChannelMemberNoContent with default headers values
func NewDeletePlatformEventChannelMemberNoContent() *DeletePlatformEventChannelMemberNoContent {
	return &DeletePlatformEventChannelMemberNoContent{}
}

/*
DeletePlatformEventChannelMemberNoContent describes a response with status code 204, with default header values.

successful operation
*/
type DeletePlatformEventChannelMemberNoContent struct {
}

func (o *DeletePlatformEventChannelMemberNoContent) Error() string {
	return fmt.Sprintf("[DELETE /PlatformEventChannelMember/{Id}][%d] deletePlatformEventChannelMemberNoContent ", 204)
}

func (o *DeletePlatformEventChannelMemberNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	//  bytes, _ := io.ReadAll(response.Body())
	//  log.Printf("[TRACE] response body: ", string(bytes))
	//  log.Printf("[TRACE] response code: ", response.Code())

	return nil
}

// NewDeletePlatformEventChannelMemberDefault creates a DeletePlatformEventChannelMemberDefault with default headers values
func NewDeletePlatformEventChannelMemberDefault(code int) *DeletePlatformEventChannelMemberDefault {
	return &DeletePlatformEventChannelMemberDefault{
		_statusCode: code,
	}
}

/*
DeletePlatformEventChannelMemberDefault describes a response with status code -1, with default header values.

Error
*/
type DeletePlatformEventChannelMemberDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the delete platform event channel member default response
func (o *DeletePlatformEventChannelMemberDefault) Code() int {
	return o._statusCode
}

func (o *DeletePlatformEventChannelMemberDefault) Error() string {
	return fmt.Sprintf("[DELETE /PlatformEventChannelMember/{Id}][%d] deletePlatformEventChannelMember default  %+v", o._statusCode, o.Payload)
}
func (o *DeletePlatformEventChannelMemberDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeletePlatformEventChannelMemberDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	//  bytes, _ := io.ReadAll(response.Body())
	//  log.Printf("[TRACE] response body: ", string(bytes))
	//  log.Printf("[TRACE] response code: ", response.Code())

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
