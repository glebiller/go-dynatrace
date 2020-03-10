// Code generated by go-swagger; DO NOT EDIT.

package automatically_applied_tags

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

// NewCreateAutoTagParams creates a new CreateAutoTagParams object
// with the default values initialized.
func NewCreateAutoTagParams() *CreateAutoTagParams {
	var ()
	return &CreateAutoTagParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateAutoTagParamsWithTimeout creates a new CreateAutoTagParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateAutoTagParamsWithTimeout(timeout time.Duration) *CreateAutoTagParams {
	var ()
	return &CreateAutoTagParams{

		timeout: timeout,
	}
}

// NewCreateAutoTagParamsWithContext creates a new CreateAutoTagParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateAutoTagParamsWithContext(ctx context.Context) *CreateAutoTagParams {
	var ()
	return &CreateAutoTagParams{

		Context: ctx,
	}
}

// NewCreateAutoTagParamsWithHTTPClient creates a new CreateAutoTagParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateAutoTagParamsWithHTTPClient(client *http.Client) *CreateAutoTagParams {
	var ()
	return &CreateAutoTagParams{
		HTTPClient: client,
	}
}

/*CreateAutoTagParams contains all the parameters to send to the API endpoint
for the create auto tag operation typically these are written to a http.Request
*/
type CreateAutoTagParams struct {

	/*Body
	  The JSON body of the request. Contains parameters of the new auto-tag.

	*/
	Body *dynatrace.AutoTag

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create auto tag params
func (o *CreateAutoTagParams) WithTimeout(timeout time.Duration) *CreateAutoTagParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create auto tag params
func (o *CreateAutoTagParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create auto tag params
func (o *CreateAutoTagParams) WithContext(ctx context.Context) *CreateAutoTagParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create auto tag params
func (o *CreateAutoTagParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create auto tag params
func (o *CreateAutoTagParams) WithHTTPClient(client *http.Client) *CreateAutoTagParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create auto tag params
func (o *CreateAutoTagParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create auto tag params
func (o *CreateAutoTagParams) WithBody(body *dynatrace.AutoTag) *CreateAutoTagParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create auto tag params
func (o *CreateAutoTagParams) SetBody(body *dynatrace.AutoTag) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAutoTagParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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