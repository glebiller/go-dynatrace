// Code generated by go-swagger; DO NOT EDIT.

package service_i_b_m_m_q_tracing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new service i b m m q tracing API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for service i b m m q tracing API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateServiceMqtTracingImsEntryQueue(params *CreateServiceMqtTracingImsEntryQueueParams, authInfo runtime.ClientAuthInfoWriter) (*CreateServiceMqtTracingImsEntryQueueCreated, error)

	DeleteServiceMqtTracingImsEntryQueue(params *DeleteServiceMqtTracingImsEntryQueueParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteServiceMqtTracingImsEntryQueueNoContent, error)

	DeleteServiceMqtTracingQueueManager(params *DeleteServiceMqtTracingQueueManagerParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteServiceMqtTracingQueueManagerNoContent, error)

	GetServiceMqtTracingImsEntryQueue(params *GetServiceMqtTracingImsEntryQueueParams, authInfo runtime.ClientAuthInfoWriter) (*GetServiceMqtTracingImsEntryQueueOK, error)

	GetServiceMqtTracingImsEntryQueues(params *GetServiceMqtTracingImsEntryQueuesParams, authInfo runtime.ClientAuthInfoWriter) (*GetServiceMqtTracingImsEntryQueuesOK, error)

	GetServiceMqtTracingQueueManager(params *GetServiceMqtTracingQueueManagerParams, authInfo runtime.ClientAuthInfoWriter) (*GetServiceMqtTracingQueueManagerOK, error)

	GetServiceMqtTracingQueueManagers(params *GetServiceMqtTracingQueueManagersParams, authInfo runtime.ClientAuthInfoWriter) (*GetServiceMqtTracingQueueManagersOK, error)

	UpdateServiceMqtTracingImsEntryQueue(params *UpdateServiceMqtTracingImsEntryQueueParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateServiceMqtTracingImsEntryQueueCreated, *UpdateServiceMqtTracingImsEntryQueueNoContent, error)

	UpdateServiceMqtTracingQueueManager(params *UpdateServiceMqtTracingQueueManagerParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateServiceMqtTracingQueueManagerCreated, *UpdateServiceMqtTracingQueueManagerNoContent, error)

	ValidateCreateServiceMqtTracingImsEntryQueue(params *ValidateCreateServiceMqtTracingImsEntryQueueParams, authInfo runtime.ClientAuthInfoWriter) (*ValidateCreateServiceMqtTracingImsEntryQueueNoContent, error)

	ValidateUpdateServiceMqtTracingImsEntryQueue(params *ValidateUpdateServiceMqtTracingImsEntryQueueParams, authInfo runtime.ClientAuthInfoWriter) (*ValidateUpdateServiceMqtTracingImsEntryQueueNoContent, error)

	ValidateUpdateServiceMqtTracingQueueManager(params *ValidateUpdateServiceMqtTracingQueueManagerParams, authInfo runtime.ClientAuthInfoWriter) (*ValidateUpdateServiceMqtTracingQueueManagerNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateServiceMqtTracingImsEntryQueue creates an i b m i m s entry queue

  The body must not provide an ID as it will be automatically assigned by the Dynatrace server.
*/
func (a *Client) CreateServiceMqtTracingImsEntryQueue(params *CreateServiceMqtTracingImsEntryQueueParams, authInfo runtime.ClientAuthInfoWriter) (*CreateServiceMqtTracingImsEntryQueueCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateServiceMqtTracingImsEntryQueueParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createServiceMqtTracingImsEntryQueue",
		Method:             "POST",
		PathPattern:        "/service/ibmMQTracing/imsEntryQueue",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateServiceMqtTracingImsEntryQueueReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateServiceMqtTracingImsEntryQueueCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createServiceMqtTracingImsEntryQueue: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteServiceMqtTracingImsEntryQueue deletes the specified i b m i m s entry queue
*/
func (a *Client) DeleteServiceMqtTracingImsEntryQueue(params *DeleteServiceMqtTracingImsEntryQueueParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteServiceMqtTracingImsEntryQueueNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteServiceMqtTracingImsEntryQueueParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteServiceMqtTracingImsEntryQueue",
		Method:             "DELETE",
		PathPattern:        "/service/ibmMQTracing/imsEntryQueue/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteServiceMqtTracingImsEntryQueueReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteServiceMqtTracingImsEntryQueueNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteServiceMqtTracingImsEntryQueue: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteServiceMqtTracingQueueManager deletes the specified queue manager
*/
func (a *Client) DeleteServiceMqtTracingQueueManager(params *DeleteServiceMqtTracingQueueManagerParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteServiceMqtTracingQueueManagerNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteServiceMqtTracingQueueManagerParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteServiceMqtTracingQueueManager",
		Method:             "DELETE",
		PathPattern:        "/service/ibmMQTracing/queueManager/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteServiceMqtTracingQueueManagerReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteServiceMqtTracingQueueManagerNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteServiceMqtTracingQueueManager: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetServiceMqtTracingImsEntryQueue gets the properties of the specified i b m i m s entry queue
*/
func (a *Client) GetServiceMqtTracingImsEntryQueue(params *GetServiceMqtTracingImsEntryQueueParams, authInfo runtime.ClientAuthInfoWriter) (*GetServiceMqtTracingImsEntryQueueOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetServiceMqtTracingImsEntryQueueParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getServiceMqtTracingImsEntryQueue",
		Method:             "GET",
		PathPattern:        "/service/ibmMQTracing/imsEntryQueue/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetServiceMqtTracingImsEntryQueueReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetServiceMqtTracingImsEntryQueueOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getServiceMqtTracingImsEntryQueue: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetServiceMqtTracingImsEntryQueues lists all i b m i m s entry queues

  This endpoint is used to provide Dynatrace with all IBM MQ queues (defined by QueueManagerName and QueueName) which are used to send messages to IBM IMS on the mainframe.

 This is required to facilitate end to end tracing for messages sent via IBM MQ to IBM IMS.
*/
func (a *Client) GetServiceMqtTracingImsEntryQueues(params *GetServiceMqtTracingImsEntryQueuesParams, authInfo runtime.ClientAuthInfoWriter) (*GetServiceMqtTracingImsEntryQueuesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetServiceMqtTracingImsEntryQueuesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getServiceMqtTracingImsEntryQueues",
		Method:             "GET",
		PathPattern:        "/service/ibmMQTracing/imsEntryQueue",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetServiceMqtTracingImsEntryQueuesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetServiceMqtTracingImsEntryQueuesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getServiceMqtTracingImsEntryQueues: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetServiceMqtTracingQueueManager gets the parameters of the specified queue manager
*/
func (a *Client) GetServiceMqtTracingQueueManager(params *GetServiceMqtTracingQueueManagerParams, authInfo runtime.ClientAuthInfoWriter) (*GetServiceMqtTracingQueueManagerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetServiceMqtTracingQueueManagerParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getServiceMqtTracingQueueManager",
		Method:             "GET",
		PathPattern:        "/service/ibmMQTracing/queueManager/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetServiceMqtTracingQueueManagerReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetServiceMqtTracingQueueManagerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getServiceMqtTracingQueueManager: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetServiceMqtTracingQueueManagers lists all available queue managers

  This endpoint is used to provide Dynatrace with your IBM MQ setup regarding alias, remote and cluster queues.

 This is required to facilitate end to end tracing for messages send via IBM MQ where sender and receiver use different queue names. Without this information Dynatrace would still trace all requests, but would not be able to stitch service calls that use these IBM MQ features.
*/
func (a *Client) GetServiceMqtTracingQueueManagers(params *GetServiceMqtTracingQueueManagersParams, authInfo runtime.ClientAuthInfoWriter) (*GetServiceMqtTracingQueueManagersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetServiceMqtTracingQueueManagersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getServiceMqtTracingQueueManagers",
		Method:             "GET",
		PathPattern:        "/service/ibmMQTracing/queueManager",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetServiceMqtTracingQueueManagersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetServiceMqtTracingQueueManagersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getServiceMqtTracingQueueManagers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateServiceMqtTracingImsEntryQueue updates an existing i b m i m s entry queue or creates a new one

  If the IBM IMS entry queue with the specified ID does not exist, a new IBM IMS entry queue will be created.
*/
func (a *Client) UpdateServiceMqtTracingImsEntryQueue(params *UpdateServiceMqtTracingImsEntryQueueParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateServiceMqtTracingImsEntryQueueCreated, *UpdateServiceMqtTracingImsEntryQueueNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateServiceMqtTracingImsEntryQueueParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateServiceMqtTracingImsEntryQueue",
		Method:             "PUT",
		PathPattern:        "/service/ibmMQTracing/imsEntryQueue/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateServiceMqtTracingImsEntryQueueReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *UpdateServiceMqtTracingImsEntryQueueCreated:
		return value, nil, nil
	case *UpdateServiceMqtTracingImsEntryQueueNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for service_i_b_m_m_q_tracing: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateServiceMqtTracingQueueManager updates the specified queue manager or creates a new one

  If the queue manager with the specified ID doesn’t exist, a new queue manager will be created.
*/
func (a *Client) UpdateServiceMqtTracingQueueManager(params *UpdateServiceMqtTracingQueueManagerParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateServiceMqtTracingQueueManagerCreated, *UpdateServiceMqtTracingQueueManagerNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateServiceMqtTracingQueueManagerParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateServiceMqtTracingQueueManager",
		Method:             "PUT",
		PathPattern:        "/service/ibmMQTracing/queueManager/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateServiceMqtTracingQueueManagerReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *UpdateServiceMqtTracingQueueManagerCreated:
		return value, nil, nil
	case *UpdateServiceMqtTracingQueueManagerNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for service_i_b_m_m_q_tracing: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ValidateCreateServiceMqtTracingImsEntryQueue validates new i b m i m s entry queues for the p o s t service ibm m q tracing ims entry queue request
*/
func (a *Client) ValidateCreateServiceMqtTracingImsEntryQueue(params *ValidateCreateServiceMqtTracingImsEntryQueueParams, authInfo runtime.ClientAuthInfoWriter) (*ValidateCreateServiceMqtTracingImsEntryQueueNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewValidateCreateServiceMqtTracingImsEntryQueueParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "validateCreateServiceMqtTracingImsEntryQueue",
		Method:             "POST",
		PathPattern:        "/service/ibmMQTracing/imsEntryQueue/validator",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ValidateCreateServiceMqtTracingImsEntryQueueReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ValidateCreateServiceMqtTracingImsEntryQueueNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for validateCreateServiceMqtTracingImsEntryQueue: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ValidateUpdateServiceMqtTracingImsEntryQueue validates update of existing i b m i m s entry queues for the p u t service ibm m q tracing ims entry queue id request
*/
func (a *Client) ValidateUpdateServiceMqtTracingImsEntryQueue(params *ValidateUpdateServiceMqtTracingImsEntryQueueParams, authInfo runtime.ClientAuthInfoWriter) (*ValidateUpdateServiceMqtTracingImsEntryQueueNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewValidateUpdateServiceMqtTracingImsEntryQueueParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "validateUpdateServiceMqtTracingImsEntryQueue",
		Method:             "POST",
		PathPattern:        "/service/ibmMQTracing/imsEntryQueue/{id}/validator",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ValidateUpdateServiceMqtTracingImsEntryQueueReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ValidateUpdateServiceMqtTracingImsEntryQueueNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for validateUpdateServiceMqtTracingImsEntryQueue: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ValidateUpdateServiceMqtTracingQueueManager validates the queue manager update for the p u t service ibm m q tracing queue manager name request
*/
func (a *Client) ValidateUpdateServiceMqtTracingQueueManager(params *ValidateUpdateServiceMqtTracingQueueManagerParams, authInfo runtime.ClientAuthInfoWriter) (*ValidateUpdateServiceMqtTracingQueueManagerNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewValidateUpdateServiceMqtTracingQueueManagerParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "validateUpdateServiceMqtTracingQueueManager",
		Method:             "POST",
		PathPattern:        "/service/ibmMQTracing/queueManager/{name}/validator",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ValidateUpdateServiceMqtTracingQueueManagerReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ValidateUpdateServiceMqtTracingQueueManagerNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for validateUpdateServiceMqtTracingQueueManager: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
