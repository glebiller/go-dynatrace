// Code generated by go-swagger; DO NOT EDIT.

package mobile_deobfuscation_and_symbolication

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

// NewDeleteSymfileParams creates a new DeleteSymfileParams object
// with the default values initialized.
func NewDeleteSymfileParams() *DeleteSymfileParams {
	var ()
	return &DeleteSymfileParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSymfileParamsWithTimeout creates a new DeleteSymfileParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteSymfileParamsWithTimeout(timeout time.Duration) *DeleteSymfileParams {
	var ()
	return &DeleteSymfileParams{

		timeout: timeout,
	}
}

// NewDeleteSymfileParamsWithContext creates a new DeleteSymfileParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteSymfileParamsWithContext(ctx context.Context) *DeleteSymfileParams {
	var ()
	return &DeleteSymfileParams{

		Context: ctx,
	}
}

// NewDeleteSymfileParamsWithHTTPClient creates a new DeleteSymfileParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteSymfileParamsWithHTTPClient(client *http.Client) *DeleteSymfileParams {
	var ()
	return &DeleteSymfileParams{
		HTTPClient: client,
	}
}

/*DeleteSymfileParams contains all the parameters to send to the API endpoint
for the delete symfile operation typically these are written to a http.Request
*/
type DeleteSymfileParams struct {

	/*ApplicationID
	  The application id used in Dynatrace of the mobile application where a file should be deleted

	*/
	ApplicationID strfmt.UUID
	/*Os
	  The operating system the file belongs to

	*/
	Os string
	/*PackageName
	  The CFBundleIdentifier (iOS) or the package name (Android) which identifies the app associated with the file to be deleted

	*/
	PackageName string
	/*VersionCode
	  The version code (Android) / CFBundleVersion (iOS) of the file to be deleted

	*/
	VersionCode string
	/*VersionName
	  The version name (Android) / CFBundleShortVersionString (iOS) of the file to be deleted

	*/
	VersionName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete symfile params
func (o *DeleteSymfileParams) WithTimeout(timeout time.Duration) *DeleteSymfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete symfile params
func (o *DeleteSymfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete symfile params
func (o *DeleteSymfileParams) WithContext(ctx context.Context) *DeleteSymfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete symfile params
func (o *DeleteSymfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete symfile params
func (o *DeleteSymfileParams) WithHTTPClient(client *http.Client) *DeleteSymfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete symfile params
func (o *DeleteSymfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplicationID adds the applicationID to the delete symfile params
func (o *DeleteSymfileParams) WithApplicationID(applicationID strfmt.UUID) *DeleteSymfileParams {
	o.SetApplicationID(applicationID)
	return o
}

// SetApplicationID adds the applicationId to the delete symfile params
func (o *DeleteSymfileParams) SetApplicationID(applicationID strfmt.UUID) {
	o.ApplicationID = applicationID
}

// WithOs adds the os to the delete symfile params
func (o *DeleteSymfileParams) WithOs(os string) *DeleteSymfileParams {
	o.SetOs(os)
	return o
}

// SetOs adds the os to the delete symfile params
func (o *DeleteSymfileParams) SetOs(os string) {
	o.Os = os
}

// WithPackageName adds the packageName to the delete symfile params
func (o *DeleteSymfileParams) WithPackageName(packageName string) *DeleteSymfileParams {
	o.SetPackageName(packageName)
	return o
}

// SetPackageName adds the packageName to the delete symfile params
func (o *DeleteSymfileParams) SetPackageName(packageName string) {
	o.PackageName = packageName
}

// WithVersionCode adds the versionCode to the delete symfile params
func (o *DeleteSymfileParams) WithVersionCode(versionCode string) *DeleteSymfileParams {
	o.SetVersionCode(versionCode)
	return o
}

// SetVersionCode adds the versionCode to the delete symfile params
func (o *DeleteSymfileParams) SetVersionCode(versionCode string) {
	o.VersionCode = versionCode
}

// WithVersionName adds the versionName to the delete symfile params
func (o *DeleteSymfileParams) WithVersionName(versionName string) *DeleteSymfileParams {
	o.SetVersionName(versionName)
	return o
}

// SetVersionName adds the versionName to the delete symfile params
func (o *DeleteSymfileParams) SetVersionName(versionName string) {
	o.VersionName = versionName
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSymfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param applicationId
	if err := r.SetPathParam("applicationId", o.ApplicationID.String()); err != nil {
		return err
	}

	// path param os
	if err := r.SetPathParam("os", o.Os); err != nil {
		return err
	}

	// path param packageName
	if err := r.SetPathParam("packageName", o.PackageName); err != nil {
		return err
	}

	// path param versionCode
	if err := r.SetPathParam("versionCode", o.VersionCode); err != nil {
		return err
	}

	// path param versionName
	if err := r.SetPathParam("versionName", o.VersionName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
