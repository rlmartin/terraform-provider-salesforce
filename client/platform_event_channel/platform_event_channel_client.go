// Code generated by go-swagger; DO NOT EDIT.

package platform_event_channel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"log"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new platform event channel API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for platform event channel API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
CreatePlatformEventChannel creates platform event channel
*/
func (a *Client) CreatePlatformEventChannel(params *CreatePlatformEventChannelParams) (*CreatePlatformEventChannelOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreatePlatformEventChannelParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createPlatformEventChannel",
		Method:             "POST",
		PathPattern:        "/PlatformEventChannel",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreatePlatformEventChannelReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		log.Printf("[TRACE] err: %+v", err)
		return nil, err
	}
	return result.(*CreatePlatformEventChannelOK), nil

}

/*
DeletePlatformEventChannel deletes platform event channel
*/
func (a *Client) DeletePlatformEventChannel(params *DeletePlatformEventChannelParams) (*DeletePlatformEventChannelNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePlatformEventChannelParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deletePlatformEventChannel",
		Method:             "DELETE",
		PathPattern:        "/PlatformEventChannel/{Id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeletePlatformEventChannelReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		log.Printf("[TRACE] err: %+v", err)
		return nil, err
	}
	return result.(*DeletePlatformEventChannelNoContent), nil

}

/*
GetPlatformEventChannel gets platform event channel
*/
func (a *Client) GetPlatformEventChannel(params *GetPlatformEventChannelParams) (*GetPlatformEventChannelOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPlatformEventChannelParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPlatformEventChannel",
		Method:             "GET",
		PathPattern:        "/PlatformEventChannel/{Id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPlatformEventChannelReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		log.Printf("[TRACE] err: %+v", err)
		return nil, err
	}
	return result.(*GetPlatformEventChannelOK), nil

}

/*
UpdatePlatformEventChannel updates platform event channel
*/
func (a *Client) UpdatePlatformEventChannel(params *UpdatePlatformEventChannelParams) (*UpdatePlatformEventChannelNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdatePlatformEventChannelParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updatePlatformEventChannel",
		Method:             "PATCH",
		PathPattern:        "/PlatformEventChannel/{Id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdatePlatformEventChannelReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		log.Printf("[TRACE] err: %+v", err)
		return nil, err
	}
	return result.(*UpdatePlatformEventChannelNoContent), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}