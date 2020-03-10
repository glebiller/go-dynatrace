// Code generated by go-swagger; DO NOT EDIT.

package credential_vault

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new credential vault API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for credential vault API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateCredential(params *CreateCredentialParams, authInfo runtime.ClientAuthInfoWriter) (*CreateCredentialCreated, error)

	DeleteCredential(params *DeleteCredentialParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteCredentialNoContent, error)

	GetCredential(params *GetCredentialParams, authInfo runtime.ClientAuthInfoWriter) (*GetCredentialOK, error)

	GetCredentials(params *GetCredentialsParams, authInfo runtime.ClientAuthInfoWriter) (*GetCredentialsOK, error)

	UpdateCredential(params *UpdateCredentialParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateCredentialCreated, *UpdateCredentialNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateCredential creates a new credentials set pipe maturity e a r l y a d o p t e r

  The body must not provide an ID. An ID is assigned automatically by the Dynatrace server.
*/
func (a *Client) CreateCredential(params *CreateCredentialParams, authInfo runtime.ClientAuthInfoWriter) (*CreateCredentialCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateCredentialParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createCredential",
		Method:             "POST",
		PathPattern:        "/credentials",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateCredentialReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateCredentialCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createCredential: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteCredential deletes the specified credentials set pipe maturity e a r l y a d o p t e r
*/
func (a *Client) DeleteCredential(params *DeleteCredentialParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteCredentialNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteCredentialParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteCredential",
		Method:             "DELETE",
		PathPattern:        "/credentials/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteCredentialReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteCredentialNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteCredential: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCredential gets the metadata of the specified credentials set pipe maturity e a r l y a d o p t e r

  The credentials set itself (username/certificate and password) is not included in the response.
*/
func (a *Client) GetCredential(params *GetCredentialParams, authInfo runtime.ClientAuthInfoWriter) (*GetCredentialOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCredentialParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCredential",
		Method:             "GET",
		PathPattern:        "/credentials/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCredentialReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCredentialOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCredential: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCredentials lists all sets of credentials for synthetic monitors stored in your environment pipe maturity e a r l y a d o p t e r
*/
func (a *Client) GetCredentials(params *GetCredentialsParams, authInfo runtime.ClientAuthInfoWriter) (*GetCredentialsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCredentialsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCredentials",
		Method:             "GET",
		PathPattern:        "/credentials",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCredentialsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCredentialsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCredentials: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateCredential updates the specified credentials set pipe maturity e a r l y a d o p t e r
*/
func (a *Client) UpdateCredential(params *UpdateCredentialParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateCredentialCreated, *UpdateCredentialNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateCredentialParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateCredential",
		Method:             "PUT",
		PathPattern:        "/credentials/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateCredentialReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *UpdateCredentialCreated:
		return value, nil, nil
	case *UpdateCredentialNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for credential_vault: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
