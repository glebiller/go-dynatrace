// Code generated by go-swagger; DO NOT EDIT.

package azure_credentials_configuration

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

// NewDeleteAzureCredentialParams creates a new DeleteAzureCredentialParams object
// with the default values initialized.
func NewDeleteAzureCredentialParams() *DeleteAzureCredentialParams {
	var ()
	return &DeleteAzureCredentialParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAzureCredentialParamsWithTimeout creates a new DeleteAzureCredentialParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteAzureCredentialParamsWithTimeout(timeout time.Duration) *DeleteAzureCredentialParams {
	var ()
	return &DeleteAzureCredentialParams{

		timeout: timeout,
	}
}

// NewDeleteAzureCredentialParamsWithContext creates a new DeleteAzureCredentialParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteAzureCredentialParamsWithContext(ctx context.Context) *DeleteAzureCredentialParams {
	var ()
	return &DeleteAzureCredentialParams{

		Context: ctx,
	}
}

// NewDeleteAzureCredentialParamsWithHTTPClient creates a new DeleteAzureCredentialParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteAzureCredentialParamsWithHTTPClient(client *http.Client) *DeleteAzureCredentialParams {
	var ()
	return &DeleteAzureCredentialParams{
		HTTPClient: client,
	}
}

/*DeleteAzureCredentialParams contains all the parameters to send to the API endpoint
for the delete azure credential operation typically these are written to a http.Request
*/
type DeleteAzureCredentialParams struct {

	/*ID
	  The ID of the Azure credentials configuration to be deleted.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete azure credential params
func (o *DeleteAzureCredentialParams) WithTimeout(timeout time.Duration) *DeleteAzureCredentialParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete azure credential params
func (o *DeleteAzureCredentialParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete azure credential params
func (o *DeleteAzureCredentialParams) WithContext(ctx context.Context) *DeleteAzureCredentialParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete azure credential params
func (o *DeleteAzureCredentialParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete azure credential params
func (o *DeleteAzureCredentialParams) WithHTTPClient(client *http.Client) *DeleteAzureCredentialParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete azure credential params
func (o *DeleteAzureCredentialParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete azure credential params
func (o *DeleteAzureCredentialParams) WithID(id string) *DeleteAzureCredentialParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete azure credential params
func (o *DeleteAzureCredentialParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAzureCredentialParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
