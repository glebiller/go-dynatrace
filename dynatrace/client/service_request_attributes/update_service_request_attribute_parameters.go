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

	"github.com/Kissy/go-dynatrace/dynatrace"
)

// NewUpdateServiceRequestAttributeParams creates a new UpdateServiceRequestAttributeParams object
// with the default values initialized.
func NewUpdateServiceRequestAttributeParams() *UpdateServiceRequestAttributeParams {
	var ()
	return &UpdateServiceRequestAttributeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateServiceRequestAttributeParamsWithTimeout creates a new UpdateServiceRequestAttributeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateServiceRequestAttributeParamsWithTimeout(timeout time.Duration) *UpdateServiceRequestAttributeParams {
	var ()
	return &UpdateServiceRequestAttributeParams{

		timeout: timeout,
	}
}

// NewUpdateServiceRequestAttributeParamsWithContext creates a new UpdateServiceRequestAttributeParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateServiceRequestAttributeParamsWithContext(ctx context.Context) *UpdateServiceRequestAttributeParams {
	var ()
	return &UpdateServiceRequestAttributeParams{

		Context: ctx,
	}
}

// NewUpdateServiceRequestAttributeParamsWithHTTPClient creates a new UpdateServiceRequestAttributeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateServiceRequestAttributeParamsWithHTTPClient(client *http.Client) *UpdateServiceRequestAttributeParams {
	var ()
	return &UpdateServiceRequestAttributeParams{
		HTTPClient: client,
	}
}

/*UpdateServiceRequestAttributeParams contains all the parameters to send to the API endpoint
for the update service request attribute operation typically these are written to a http.Request
*/
type UpdateServiceRequestAttributeParams struct {

	/*Body*/
	Body *dynatrace.RequestAttribute
	/*ID
	 The ID of the request attribute to be updated.

	If you set the ID in the body as well, it must match this ID.

	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update service request attribute params
func (o *UpdateServiceRequestAttributeParams) WithTimeout(timeout time.Duration) *UpdateServiceRequestAttributeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update service request attribute params
func (o *UpdateServiceRequestAttributeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update service request attribute params
func (o *UpdateServiceRequestAttributeParams) WithContext(ctx context.Context) *UpdateServiceRequestAttributeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update service request attribute params
func (o *UpdateServiceRequestAttributeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update service request attribute params
func (o *UpdateServiceRequestAttributeParams) WithHTTPClient(client *http.Client) *UpdateServiceRequestAttributeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update service request attribute params
func (o *UpdateServiceRequestAttributeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update service request attribute params
func (o *UpdateServiceRequestAttributeParams) WithBody(body *dynatrace.RequestAttribute) *UpdateServiceRequestAttributeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update service request attribute params
func (o *UpdateServiceRequestAttributeParams) SetBody(body *dynatrace.RequestAttribute) {
	o.Body = body
}

// WithID adds the id to the update service request attribute params
func (o *UpdateServiceRequestAttributeParams) WithID(id strfmt.UUID) *UpdateServiceRequestAttributeParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update service request attribute params
func (o *UpdateServiceRequestAttributeParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateServiceRequestAttributeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
