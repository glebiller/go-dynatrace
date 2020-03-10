// Code generated by go-swagger; DO NOT EDIT.

package cloud_foundry_credentials_configuration

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

// NewDeleteCloudFoundryCredentialParams creates a new DeleteCloudFoundryCredentialParams object
// with the default values initialized.
func NewDeleteCloudFoundryCredentialParams() *DeleteCloudFoundryCredentialParams {
	var ()
	return &DeleteCloudFoundryCredentialParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCloudFoundryCredentialParamsWithTimeout creates a new DeleteCloudFoundryCredentialParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteCloudFoundryCredentialParamsWithTimeout(timeout time.Duration) *DeleteCloudFoundryCredentialParams {
	var ()
	return &DeleteCloudFoundryCredentialParams{

		timeout: timeout,
	}
}

// NewDeleteCloudFoundryCredentialParamsWithContext creates a new DeleteCloudFoundryCredentialParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteCloudFoundryCredentialParamsWithContext(ctx context.Context) *DeleteCloudFoundryCredentialParams {
	var ()
	return &DeleteCloudFoundryCredentialParams{

		Context: ctx,
	}
}

// NewDeleteCloudFoundryCredentialParamsWithHTTPClient creates a new DeleteCloudFoundryCredentialParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteCloudFoundryCredentialParamsWithHTTPClient(client *http.Client) *DeleteCloudFoundryCredentialParams {
	var ()
	return &DeleteCloudFoundryCredentialParams{
		HTTPClient: client,
	}
}

/*DeleteCloudFoundryCredentialParams contains all the parameters to send to the API endpoint
for the delete cloud foundry credential operation typically these are written to a http.Request
*/
type DeleteCloudFoundryCredentialParams struct {

	/*ID
	  The ID of the Cloud Foundry foundation credentials to be deleted.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete cloud foundry credential params
func (o *DeleteCloudFoundryCredentialParams) WithTimeout(timeout time.Duration) *DeleteCloudFoundryCredentialParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete cloud foundry credential params
func (o *DeleteCloudFoundryCredentialParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete cloud foundry credential params
func (o *DeleteCloudFoundryCredentialParams) WithContext(ctx context.Context) *DeleteCloudFoundryCredentialParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete cloud foundry credential params
func (o *DeleteCloudFoundryCredentialParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete cloud foundry credential params
func (o *DeleteCloudFoundryCredentialParams) WithHTTPClient(client *http.Client) *DeleteCloudFoundryCredentialParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete cloud foundry credential params
func (o *DeleteCloudFoundryCredentialParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete cloud foundry credential params
func (o *DeleteCloudFoundryCredentialParams) WithID(id string) *DeleteCloudFoundryCredentialParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete cloud foundry credential params
func (o *DeleteCloudFoundryCredentialParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCloudFoundryCredentialParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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