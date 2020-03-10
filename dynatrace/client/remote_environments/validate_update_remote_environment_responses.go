// Code generated by go-swagger; DO NOT EDIT.

package remote_environments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/Kissy/go-dynatrace/dynatrace"
)

// ValidateUpdateRemoteEnvironmentReader is a Reader for the ValidateUpdateRemoteEnvironment structure.
type ValidateUpdateRemoteEnvironmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ValidateUpdateRemoteEnvironmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewValidateUpdateRemoteEnvironmentNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewValidateUpdateRemoteEnvironmentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewValidateUpdateRemoteEnvironmentNoContent creates a ValidateUpdateRemoteEnvironmentNoContent with default headers values
func NewValidateUpdateRemoteEnvironmentNoContent() *ValidateUpdateRemoteEnvironmentNoContent {
	return &ValidateUpdateRemoteEnvironmentNoContent{}
}

/*ValidateUpdateRemoteEnvironmentNoContent handles this case with default header values.

Validated. The submitted configuration is valid. Response doesn't have a body
*/
type ValidateUpdateRemoteEnvironmentNoContent struct {
}

func (o *ValidateUpdateRemoteEnvironmentNoContent) Error() string {
	return fmt.Sprintf("[POST /remoteEnvironments/{id}/validator][%d] validateUpdateRemoteEnvironmentNoContent ", 204)
}

func (o *ValidateUpdateRemoteEnvironmentNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewValidateUpdateRemoteEnvironmentBadRequest creates a ValidateUpdateRemoteEnvironmentBadRequest with default headers values
func NewValidateUpdateRemoteEnvironmentBadRequest() *ValidateUpdateRemoteEnvironmentBadRequest {
	return &ValidateUpdateRemoteEnvironmentBadRequest{}
}

/*ValidateUpdateRemoteEnvironmentBadRequest handles this case with default header values.

Failed. The input is invalid.
*/
type ValidateUpdateRemoteEnvironmentBadRequest struct {
	Payload *dynatrace.ErrorEnvelope
}

func (o *ValidateUpdateRemoteEnvironmentBadRequest) Error() string {
	return fmt.Sprintf("[POST /remoteEnvironments/{id}/validator][%d] validateUpdateRemoteEnvironmentBadRequest  %+v", 400, o.Payload)
}

func (o *ValidateUpdateRemoteEnvironmentBadRequest) GetPayload() *dynatrace.ErrorEnvelope {
	return o.Payload
}

func (o *ValidateUpdateRemoteEnvironmentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.ErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
