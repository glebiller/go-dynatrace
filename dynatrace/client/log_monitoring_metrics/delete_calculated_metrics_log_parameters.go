// Code generated by go-swagger; DO NOT EDIT.

package log_monitoring_metrics

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

// NewDeleteCalculatedMetricsLogParams creates a new DeleteCalculatedMetricsLogParams object
// with the default values initialized.
func NewDeleteCalculatedMetricsLogParams() *DeleteCalculatedMetricsLogParams {
	var ()
	return &DeleteCalculatedMetricsLogParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCalculatedMetricsLogParamsWithTimeout creates a new DeleteCalculatedMetricsLogParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteCalculatedMetricsLogParamsWithTimeout(timeout time.Duration) *DeleteCalculatedMetricsLogParams {
	var ()
	return &DeleteCalculatedMetricsLogParams{

		timeout: timeout,
	}
}

// NewDeleteCalculatedMetricsLogParamsWithContext creates a new DeleteCalculatedMetricsLogParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteCalculatedMetricsLogParamsWithContext(ctx context.Context) *DeleteCalculatedMetricsLogParams {
	var ()
	return &DeleteCalculatedMetricsLogParams{

		Context: ctx,
	}
}

// NewDeleteCalculatedMetricsLogParamsWithHTTPClient creates a new DeleteCalculatedMetricsLogParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteCalculatedMetricsLogParamsWithHTTPClient(client *http.Client) *DeleteCalculatedMetricsLogParams {
	var ()
	return &DeleteCalculatedMetricsLogParams{
		HTTPClient: client,
	}
}

/*DeleteCalculatedMetricsLogParams contains all the parameters to send to the API endpoint
for the delete calculated metrics log operation typically these are written to a http.Request
*/
type DeleteCalculatedMetricsLogParams struct {

	/*MetricKey
	  The key of the custom log metric to be deleted.

	*/
	MetricKey string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete calculated metrics log params
func (o *DeleteCalculatedMetricsLogParams) WithTimeout(timeout time.Duration) *DeleteCalculatedMetricsLogParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete calculated metrics log params
func (o *DeleteCalculatedMetricsLogParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete calculated metrics log params
func (o *DeleteCalculatedMetricsLogParams) WithContext(ctx context.Context) *DeleteCalculatedMetricsLogParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete calculated metrics log params
func (o *DeleteCalculatedMetricsLogParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete calculated metrics log params
func (o *DeleteCalculatedMetricsLogParams) WithHTTPClient(client *http.Client) *DeleteCalculatedMetricsLogParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete calculated metrics log params
func (o *DeleteCalculatedMetricsLogParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMetricKey adds the metricKey to the delete calculated metrics log params
func (o *DeleteCalculatedMetricsLogParams) WithMetricKey(metricKey string) *DeleteCalculatedMetricsLogParams {
	o.SetMetricKey(metricKey)
	return o
}

// SetMetricKey adds the metricKey to the delete calculated metrics log params
func (o *DeleteCalculatedMetricsLogParams) SetMetricKey(metricKey string) {
	o.MetricKey = metricKey
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCalculatedMetricsLogParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param metricKey
	if err := r.SetPathParam("metricKey", o.MetricKey); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
