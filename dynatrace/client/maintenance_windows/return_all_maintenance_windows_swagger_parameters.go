// Code generated by go-swagger; DO NOT EDIT.

package maintenance_windows

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

// NewReturnAllMaintenanceWindowsParams creates a new ReturnAllMaintenanceWindowsParams object
// with the default values initialized.
func NewReturnAllMaintenanceWindowsParams() *ReturnAllMaintenanceWindowsParams {

	return &ReturnAllMaintenanceWindowsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReturnAllMaintenanceWindowsParamsWithTimeout creates a new ReturnAllMaintenanceWindowsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReturnAllMaintenanceWindowsParamsWithTimeout(timeout time.Duration) *ReturnAllMaintenanceWindowsParams {

	return &ReturnAllMaintenanceWindowsParams{

		timeout: timeout,
	}
}

// NewReturnAllMaintenanceWindowsParamsWithContext creates a new ReturnAllMaintenanceWindowsParams object
// with the default values initialized, and the ability to set a context for a request
func NewReturnAllMaintenanceWindowsParamsWithContext(ctx context.Context) *ReturnAllMaintenanceWindowsParams {

	return &ReturnAllMaintenanceWindowsParams{

		Context: ctx,
	}
}

// NewReturnAllMaintenanceWindowsParamsWithHTTPClient creates a new ReturnAllMaintenanceWindowsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReturnAllMaintenanceWindowsParamsWithHTTPClient(client *http.Client) *ReturnAllMaintenanceWindowsParams {

	return &ReturnAllMaintenanceWindowsParams{
		HTTPClient: client,
	}
}

/*ReturnAllMaintenanceWindowsParams contains all the parameters to send to the API endpoint
for the return all maintenance windows operation typically these are written to a http.Request
*/
type ReturnAllMaintenanceWindowsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the return all maintenance windows params
func (o *ReturnAllMaintenanceWindowsParams) WithTimeout(timeout time.Duration) *ReturnAllMaintenanceWindowsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the return all maintenance windows params
func (o *ReturnAllMaintenanceWindowsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the return all maintenance windows params
func (o *ReturnAllMaintenanceWindowsParams) WithContext(ctx context.Context) *ReturnAllMaintenanceWindowsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the return all maintenance windows params
func (o *ReturnAllMaintenanceWindowsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the return all maintenance windows params
func (o *ReturnAllMaintenanceWindowsParams) WithHTTPClient(client *http.Client) *ReturnAllMaintenanceWindowsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the return all maintenance windows params
func (o *ReturnAllMaintenanceWindowsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ReturnAllMaintenanceWindowsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
