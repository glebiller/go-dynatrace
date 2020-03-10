// Code generated by go-swagger; DO NOT EDIT.

package remote_environments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteRemoteEnvironmentReader is a Reader for the DeleteRemoteEnvironment structure.
type DeleteRemoteEnvironmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRemoteEnvironmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteRemoteEnvironmentNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteRemoteEnvironmentNoContent creates a DeleteRemoteEnvironmentNoContent with default headers values
func NewDeleteRemoteEnvironmentNoContent() *DeleteRemoteEnvironmentNoContent {
	return &DeleteRemoteEnvironmentNoContent{}
}

/*DeleteRemoteEnvironmentNoContent handles this case with default header values.

Success. The configuration has been deleted. The response doesn't have a body.
*/
type DeleteRemoteEnvironmentNoContent struct {
}

func (o *DeleteRemoteEnvironmentNoContent) Error() string {
	return fmt.Sprintf("[DELETE /remoteEnvironments/{id}][%d] deleteRemoteEnvironmentNoContent ", 204)
}

func (o *DeleteRemoteEnvironmentNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
