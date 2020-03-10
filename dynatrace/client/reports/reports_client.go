// Code generated by go-swagger; DO NOT EDIT.

package reports

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new reports API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for reports API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateReport(params *CreateReportParams, authInfo runtime.ClientAuthInfoWriter) (*CreateReportCreated, error)

	DeleteReport(params *DeleteReportParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteReportNoContent, error)

	GetReport(params *GetReportParams, authInfo runtime.ClientAuthInfoWriter) (*GetReportOK, error)

	GetReports(params *GetReportsParams, authInfo runtime.ClientAuthInfoWriter) (*GetReportsOK, error)

	SubscribeReport(params *SubscribeReportParams, authInfo runtime.ClientAuthInfoWriter) (*SubscribeReportCreated, error)

	UnsubscribeReport(params *UnsubscribeReportParams, authInfo runtime.ClientAuthInfoWriter) (*UnsubscribeReportCreated, error)

	UpdateReport(params *UpdateReportParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateReportCreated, *UpdateReportNoContent, error)

	ValidateCreateReport(params *ValidateCreateReportParams, authInfo runtime.ClientAuthInfoWriter) (*ValidateCreateReportNoContent, error)

	ValidateUpdateReport(params *ValidateUpdateReportParams, authInfo runtime.ClientAuthInfoWriter) (*ValidateUpdateReportNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateReport creates a report pipe maturity e a r l y a d o p t e r

  The body must not provide an ID. An ID is assigned automatically by the Dynatrace server.
*/
func (a *Client) CreateReport(params *CreateReportParams, authInfo runtime.ClientAuthInfoWriter) (*CreateReportCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateReportParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createReport",
		Method:             "POST",
		PathPattern:        "/reports",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateReportReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateReportCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createReport: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteReport deletes the specified report pipe maturity e a r l y a d o p t e r
*/
func (a *Client) DeleteReport(params *DeleteReportParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteReportNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteReportParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteReport",
		Method:             "DELETE",
		PathPattern:        "/reports/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteReportReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteReportNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteReport: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetReport gets the properties of the specified report pipe maturity e a r l y a d o p t e r
*/
func (a *Client) GetReport(params *GetReportParams, authInfo runtime.ClientAuthInfoWriter) (*GetReportOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetReportParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getReport",
		Method:             "GET",
		PathPattern:        "/reports/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetReportReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetReportOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getReport: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetReports lists all reports available in your environment pipe maturity e a r l y a d o p t e r
*/
func (a *Client) GetReports(params *GetReportsParams, authInfo runtime.ClientAuthInfoWriter) (*GetReportsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetReportsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getReports",
		Method:             "GET",
		PathPattern:        "/reports",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetReportsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetReportsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getReports: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SubscribeReport subscribes to the specified report pipe maturity e a r l y a d o p t e r
*/
func (a *Client) SubscribeReport(params *SubscribeReportParams, authInfo runtime.ClientAuthInfoWriter) (*SubscribeReportCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSubscribeReportParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "subscribeReport",
		Method:             "POST",
		PathPattern:        "/reports/{id}/subscribe",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SubscribeReportReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SubscribeReportCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for subscribeReport: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UnsubscribeReport unsubscribes from the specified report pipe maturity e a r l y a d o p t e r
*/
func (a *Client) UnsubscribeReport(params *UnsubscribeReportParams, authInfo runtime.ClientAuthInfoWriter) (*UnsubscribeReportCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUnsubscribeReportParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "unsubscribeReport",
		Method:             "POST",
		PathPattern:        "/reports/{id}/unsubscribe",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UnsubscribeReportReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UnsubscribeReportCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for unsubscribeReport: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateReport updates an existing report pipe maturity e a r l y a d o p t e r

  If a report with the specified ID doesn't exist, a new report is created.
*/
func (a *Client) UpdateReport(params *UpdateReportParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateReportCreated, *UpdateReportNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateReportParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateReport",
		Method:             "PUT",
		PathPattern:        "/reports/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateReportReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *UpdateReportCreated:
		return value, nil, nil
	case *UpdateReportNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for reports: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ValidateCreateReport validates the payload for the p o s t reports request pipe maturity e a r l y a d o p t e r

  The body must not provide an ID.
*/
func (a *Client) ValidateCreateReport(params *ValidateCreateReportParams, authInfo runtime.ClientAuthInfoWriter) (*ValidateCreateReportNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewValidateCreateReportParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "validateCreateReport",
		Method:             "POST",
		PathPattern:        "/reports/validator",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ValidateCreateReportReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ValidateCreateReportNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for validateCreateReport: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ValidateUpdateReport validates the payload for the p u t reports id request pipe maturity e a r l y a d o p t e r
*/
func (a *Client) ValidateUpdateReport(params *ValidateUpdateReportParams, authInfo runtime.ClientAuthInfoWriter) (*ValidateUpdateReportNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewValidateUpdateReportParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "validateUpdateReport",
		Method:             "POST",
		PathPattern:        "/reports/{id}/validator",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ValidateUpdateReportReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ValidateUpdateReportNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for validateUpdateReport: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
