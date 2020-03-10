// Code generated by go-swagger; DO NOT EDIT.

package a_w_s_credentials_configuration

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

// NewCreateAwsCredentialParams creates a new CreateAwsCredentialParams object
// with the default values initialized.
func NewCreateAwsCredentialParams() *CreateAwsCredentialParams {
	var ()
	return &CreateAwsCredentialParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateAwsCredentialParamsWithTimeout creates a new CreateAwsCredentialParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateAwsCredentialParamsWithTimeout(timeout time.Duration) *CreateAwsCredentialParams {
	var ()
	return &CreateAwsCredentialParams{

		timeout: timeout,
	}
}

// NewCreateAwsCredentialParamsWithContext creates a new CreateAwsCredentialParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateAwsCredentialParamsWithContext(ctx context.Context) *CreateAwsCredentialParams {
	var ()
	return &CreateAwsCredentialParams{

		Context: ctx,
	}
}

// NewCreateAwsCredentialParamsWithHTTPClient creates a new CreateAwsCredentialParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateAwsCredentialParamsWithHTTPClient(client *http.Client) *CreateAwsCredentialParams {
	var ()
	return &CreateAwsCredentialParams{
		HTTPClient: client,
	}
}

/*CreateAwsCredentialParams contains all the parameters to send to the API endpoint
for the create aws credential operation typically these are written to a http.Request
*/
type CreateAwsCredentialParams struct {

	/*Body
	  The JSON body of the request. Contains parameters of the new AWS credentials configuration.

	*/
	Body *dynatrace.AwsCredentialsConfig

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create aws credential params
func (o *CreateAwsCredentialParams) WithTimeout(timeout time.Duration) *CreateAwsCredentialParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create aws credential params
func (o *CreateAwsCredentialParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create aws credential params
func (o *CreateAwsCredentialParams) WithContext(ctx context.Context) *CreateAwsCredentialParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create aws credential params
func (o *CreateAwsCredentialParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create aws credential params
func (o *CreateAwsCredentialParams) WithHTTPClient(client *http.Client) *CreateAwsCredentialParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create aws credential params
func (o *CreateAwsCredentialParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create aws credential params
func (o *CreateAwsCredentialParams) WithBody(body *dynatrace.AwsCredentialsConfig) *CreateAwsCredentialParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create aws credential params
func (o *CreateAwsCredentialParams) SetBody(body *dynatrace.AwsCredentialsConfig) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAwsCredentialParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
