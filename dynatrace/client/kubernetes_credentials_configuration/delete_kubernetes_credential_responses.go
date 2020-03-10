// Code generated by go-swagger; DO NOT EDIT.

package kubernetes_credentials_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/Kissy/go-dynatrace/dynatrace"
)

// DeleteKubernetesCredentialReader is a Reader for the DeleteKubernetesCredential structure.
type DeleteKubernetesCredentialReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteKubernetesCredentialReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteKubernetesCredentialNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteKubernetesCredentialBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteKubernetesCredentialNoContent creates a DeleteKubernetesCredentialNoContent with default headers values
func NewDeleteKubernetesCredentialNoContent() *DeleteKubernetesCredentialNoContent {
	return &DeleteKubernetesCredentialNoContent{}
}

/*DeleteKubernetesCredentialNoContent handles this case with default header values.

Deletes the specified Kubernetes credentials configuration
*/
type DeleteKubernetesCredentialNoContent struct {
}

func (o *DeleteKubernetesCredentialNoContent) Error() string {
	return fmt.Sprintf("[DELETE /kubernetes/credentials/{id}][%d] deleteKubernetesCredentialNoContent ", 204)
}

func (o *DeleteKubernetesCredentialNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteKubernetesCredentialBadRequest creates a DeleteKubernetesCredentialBadRequest with default headers values
func NewDeleteKubernetesCredentialBadRequest() *DeleteKubernetesCredentialBadRequest {
	return &DeleteKubernetesCredentialBadRequest{}
}

/*DeleteKubernetesCredentialBadRequest handles this case with default header values.

Failed. The input is invalid.
*/
type DeleteKubernetesCredentialBadRequest struct {
	Payload *dynatrace.ErrorEnvelope
}

func (o *DeleteKubernetesCredentialBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /kubernetes/credentials/{id}][%d] deleteKubernetesCredentialBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteKubernetesCredentialBadRequest) GetPayload() *dynatrace.ErrorEnvelope {
	return o.Payload
}

func (o *DeleteKubernetesCredentialBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.ErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}