// Code generated by go-swagger; DO NOT EDIT.

package service_request_attributes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteServiceRequestAttributeParams creates a new DeleteServiceRequestAttributeParams object
// with the default values initialized.
func NewDeleteServiceRequestAttributeParams() *DeleteServiceRequestAttributeParams {
	var ()
	return &DeleteServiceRequestAttributeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteServiceRequestAttributeParamsWithTimeout creates a new DeleteServiceRequestAttributeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteServiceRequestAttributeParamsWithTimeout(timeout time.Duration) *DeleteServiceRequestAttributeParams {
	var ()
	return &DeleteServiceRequestAttributeParams{

		timeout: timeout,
	}
}

// NewDeleteServiceRequestAttributeParamsWithContext creates a new DeleteServiceRequestAttributeParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteServiceRequestAttributeParamsWithContext(ctx context.Context) *DeleteServiceRequestAttributeParams {
	var ()
	return &DeleteServiceRequestAttributeParams{

		Context: ctx,
	}
}

// NewDeleteServiceRequestAttributeParamsWithHTTPClient creates a new DeleteServiceRequestAttributeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteServiceRequestAttributeParamsWithHTTPClient(client *http.Client) *DeleteServiceRequestAttributeParams {
	var ()
	return &DeleteServiceRequestAttributeParams{
		HTTPClient: client,
	}
}

/*DeleteServiceRequestAttributeParams contains all the parameters to send to the API endpoint
for the delete service request attribute operation typically these are written to a http.Request
*/
type DeleteServiceRequestAttributeParams struct {

	/*ID
	  The ID of the request attribute to be deleted.

	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete service request attribute params
func (o *DeleteServiceRequestAttributeParams) WithTimeout(timeout time.Duration) *DeleteServiceRequestAttributeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete service request attribute params
func (o *DeleteServiceRequestAttributeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete service request attribute params
func (o *DeleteServiceRequestAttributeParams) WithContext(ctx context.Context) *DeleteServiceRequestAttributeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete service request attribute params
func (o *DeleteServiceRequestAttributeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete service request attribute params
func (o *DeleteServiceRequestAttributeParams) WithHTTPClient(client *http.Client) *DeleteServiceRequestAttributeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete service request attribute params
func (o *DeleteServiceRequestAttributeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete service request attribute params
func (o *DeleteServiceRequestAttributeParams) WithID(id strfmt.UUID) *DeleteServiceRequestAttributeParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete service request attribute params
func (o *DeleteServiceRequestAttributeParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteServiceRequestAttributeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}