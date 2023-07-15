// Code generated by go-swagger; DO NOT EDIT.

package named_credential

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetNamedCredentialParams creates a new GetNamedCredentialParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNamedCredentialParams() *GetNamedCredentialParams {
	return &GetNamedCredentialParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNamedCredentialParamsWithTimeout creates a new GetNamedCredentialParams object
// with the ability to set a timeout on a request.
func NewGetNamedCredentialParamsWithTimeout(timeout time.Duration) *GetNamedCredentialParams {
	return &GetNamedCredentialParams{
		timeout: timeout,
	}
}

// NewGetNamedCredentialParamsWithContext creates a new GetNamedCredentialParams object
// with the ability to set a context for a request.
func NewGetNamedCredentialParamsWithContext(ctx context.Context) *GetNamedCredentialParams {
	return &GetNamedCredentialParams{
		Context: ctx,
	}
}

// NewGetNamedCredentialParamsWithHTTPClient creates a new GetNamedCredentialParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNamedCredentialParamsWithHTTPClient(client *http.Client) *GetNamedCredentialParams {
	return &GetNamedCredentialParams{
		HTTPClient: client,
	}
}

/*
GetNamedCredentialParams contains all the parameters to send to the API endpoint

	for the get named credential operation.

	Typically these are written to a http.Request.
*/
type GetNamedCredentialParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get named credential params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNamedCredentialParams) WithDefaults() *GetNamedCredentialParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get named credential params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNamedCredentialParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get named credential params
func (o *GetNamedCredentialParams) WithTimeout(timeout time.Duration) *GetNamedCredentialParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get named credential params
func (o *GetNamedCredentialParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get named credential params
func (o *GetNamedCredentialParams) WithContext(ctx context.Context) *GetNamedCredentialParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get named credential params
func (o *GetNamedCredentialParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get named credential params
func (o *GetNamedCredentialParams) WithHTTPClient(client *http.Client) *GetNamedCredentialParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get named credential params
func (o *GetNamedCredentialParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get named credential params
func (o *GetNamedCredentialParams) WithID(id string) *GetNamedCredentialParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get named credential params
func (o *GetNamedCredentialParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetNamedCredentialParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param Id
	if err := r.SetPathParam("Id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
