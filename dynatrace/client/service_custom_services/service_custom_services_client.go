// Code generated by go-swagger; DO NOT EDIT.

package service_custom_services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new service custom services API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for service custom services API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateServiceCustomService(params *CreateServiceCustomServiceParams, authInfo runtime.ClientAuthInfoWriter) (*CreateServiceCustomServiceCreated, error)

	DeleteServiceCustomService(params *DeleteServiceCustomServiceParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteServiceCustomServiceNoContent, error)

	GetServiceCustomService(params *GetServiceCustomServiceParams, authInfo runtime.ClientAuthInfoWriter) (*GetServiceCustomServiceOK, error)

	GetServiceCustomServices(params *GetServiceCustomServicesParams, authInfo runtime.ClientAuthInfoWriter) (*GetServiceCustomServicesOK, error)

	UpdateServiceCustomService(params *UpdateServiceCustomServiceParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateServiceCustomServiceCreated, *UpdateServiceCustomServiceNoContent, error)

	UpdateServiceCustomServicesOrder(params *UpdateServiceCustomServicesOrderParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateServiceCustomServicesOrderNoContent, error)

	ValidateCreateServiceCustomService(params *ValidateCreateServiceCustomServiceParams, authInfo runtime.ClientAuthInfoWriter) (*ValidateCreateServiceCustomServiceNoContent, error)

	ValidateUpdateServiceCustomService(params *ValidateUpdateServiceCustomServiceParams, authInfo runtime.ClientAuthInfoWriter) (*ValidateUpdateServiceCustomServiceNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateServiceCustomService creates a custom service

  In the body of the request, neither custom service nor its rules can have the ID. All IDs will be generated automatically by Dynatrace.
*/
func (a *Client) CreateServiceCustomService(params *CreateServiceCustomServiceParams, authInfo runtime.ClientAuthInfoWriter) (*CreateServiceCustomServiceCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateServiceCustomServiceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createServiceCustomService",
		Method:             "POST",
		PathPattern:        "/service/customServices/{technology}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateServiceCustomServiceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateServiceCustomServiceCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createServiceCustomService: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteServiceCustomService deletes the specified custom service
*/
func (a *Client) DeleteServiceCustomService(params *DeleteServiceCustomServiceParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteServiceCustomServiceNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteServiceCustomServiceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteServiceCustomService",
		Method:             "DELETE",
		PathPattern:        "/service/customServices/{technology}/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteServiceCustomServiceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteServiceCustomServiceNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteServiceCustomService: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetServiceCustomService gets the definition of the specified custom service
*/
func (a *Client) GetServiceCustomService(params *GetServiceCustomServiceParams, authInfo runtime.ClientAuthInfoWriter) (*GetServiceCustomServiceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetServiceCustomServiceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getServiceCustomService",
		Method:             "GET",
		PathPattern:        "/service/customServices/{technology}/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetServiceCustomServiceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetServiceCustomServiceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getServiceCustomService: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetServiceCustomServices lists all custom services of the specified technology
*/
func (a *Client) GetServiceCustomServices(params *GetServiceCustomServicesParams, authInfo runtime.ClientAuthInfoWriter) (*GetServiceCustomServicesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetServiceCustomServicesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getServiceCustomServices",
		Method:             "GET",
		PathPattern:        "/service/customServices/{technology}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetServiceCustomServicesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetServiceCustomServicesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getServiceCustomServices: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateServiceCustomService updates the specified custom service or create a new one

  Will use the config's ´order´ attribute if supplied, otherwise keeps the order of the existing config or appends if no existing config with the supplied ID was found.
*/
func (a *Client) UpdateServiceCustomService(params *UpdateServiceCustomServiceParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateServiceCustomServiceCreated, *UpdateServiceCustomServiceNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateServiceCustomServiceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateServiceCustomService",
		Method:             "PUT",
		PathPattern:        "/service/customServices/{technology}/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateServiceCustomServiceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *UpdateServiceCustomServiceCreated:
		return value, nil, nil
	case *UpdateServiceCustomServiceNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for service_custom_services: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateServiceCustomServicesOrder reorders the custom services of the specified technology

  This request reorders the custom services of the specified technology according to the given list of IDs. Custom services not present in the body of the request will retain their relative ordering but will be ordered *after* all those present in the request.
*/
func (a *Client) UpdateServiceCustomServicesOrder(params *UpdateServiceCustomServicesOrderParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateServiceCustomServicesOrderNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateServiceCustomServicesOrderParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateServiceCustomServicesOrder",
		Method:             "PUT",
		PathPattern:        "/service/customServices/{technology}/order",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateServiceCustomServicesOrderReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateServiceCustomServicesOrderNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateServiceCustomServicesOrder: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ValidateCreateServiceCustomService validates the new custom service for the p o s t custom services technology request
*/
func (a *Client) ValidateCreateServiceCustomService(params *ValidateCreateServiceCustomServiceParams, authInfo runtime.ClientAuthInfoWriter) (*ValidateCreateServiceCustomServiceNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewValidateCreateServiceCustomServiceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "validateCreateServiceCustomService",
		Method:             "POST",
		PathPattern:        "/service/customServices/{technology}/validator",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ValidateCreateServiceCustomServiceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ValidateCreateServiceCustomServiceNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for validateCreateServiceCustomService: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ValidateUpdateServiceCustomService validates the new custom service for the p u t custom services technology id request
*/
func (a *Client) ValidateUpdateServiceCustomService(params *ValidateUpdateServiceCustomServiceParams, authInfo runtime.ClientAuthInfoWriter) (*ValidateUpdateServiceCustomServiceNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewValidateUpdateServiceCustomServiceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "validateUpdateServiceCustomService",
		Method:             "POST",
		PathPattern:        "/service/customServices/{technology}/{id}/validator",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ValidateUpdateServiceCustomServiceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ValidateUpdateServiceCustomServiceNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for validateUpdateServiceCustomService: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}