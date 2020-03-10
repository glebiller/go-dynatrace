// Code generated by go-swagger; DO NOT EDIT.

package service_request_naming

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

// NewUpdateServiceRequestNamingOrderParams creates a new UpdateServiceRequestNamingOrderParams object
// with the default values initialized.
func NewUpdateServiceRequestNamingOrderParams() *UpdateServiceRequestNamingOrderParams {
	var ()
	return &UpdateServiceRequestNamingOrderParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateServiceRequestNamingOrderParamsWithTimeout creates a new UpdateServiceRequestNamingOrderParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateServiceRequestNamingOrderParamsWithTimeout(timeout time.Duration) *UpdateServiceRequestNamingOrderParams {
	var ()
	return &UpdateServiceRequestNamingOrderParams{

		timeout: timeout,
	}
}

// NewUpdateServiceRequestNamingOrderParamsWithContext creates a new UpdateServiceRequestNamingOrderParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateServiceRequestNamingOrderParamsWithContext(ctx context.Context) *UpdateServiceRequestNamingOrderParams {
	var ()
	return &UpdateServiceRequestNamingOrderParams{

		Context: ctx,
	}
}

// NewUpdateServiceRequestNamingOrderParamsWithHTTPClient creates a new UpdateServiceRequestNamingOrderParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateServiceRequestNamingOrderParamsWithHTTPClient(client *http.Client) *UpdateServiceRequestNamingOrderParams {
	var ()
	return &UpdateServiceRequestNamingOrderParams{
		HTTPClient: client,
	}
}

/*UpdateServiceRequestNamingOrderParams contains all the parameters to send to the API endpoint
for the update service request naming order operation typically these are written to a http.Request
*/
type UpdateServiceRequestNamingOrderParams struct {

	/*Body
	  JSON body of the request containing the IDs of the request naming rules in the desired order. Any further properties (*name*, *description*) will be ignored.

	*/
	Body *dynatrace.StubList

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update service request naming order params
func (o *UpdateServiceRequestNamingOrderParams) WithTimeout(timeout time.Duration) *UpdateServiceRequestNamingOrderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update service request naming order params
func (o *UpdateServiceRequestNamingOrderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update service request naming order params
func (o *UpdateServiceRequestNamingOrderParams) WithContext(ctx context.Context) *UpdateServiceRequestNamingOrderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update service request naming order params
func (o *UpdateServiceRequestNamingOrderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update service request naming order params
func (o *UpdateServiceRequestNamingOrderParams) WithHTTPClient(client *http.Client) *UpdateServiceRequestNamingOrderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update service request naming order params
func (o *UpdateServiceRequestNamingOrderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update service request naming order params
func (o *UpdateServiceRequestNamingOrderParams) WithBody(body *dynatrace.StubList) *UpdateServiceRequestNamingOrderParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update service request naming order params
func (o *UpdateServiceRequestNamingOrderParams) SetBody(body *dynatrace.StubList) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateServiceRequestNamingOrderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
