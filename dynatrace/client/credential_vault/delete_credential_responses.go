// Code generated by go-swagger; DO NOT EDIT.

package credential_vault

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteCredentialReader is a Reader for the DeleteCredential structure.
type DeleteCredentialReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCredentialReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteCredentialNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteCredentialNoContent creates a DeleteCredentialNoContent with default headers values
func NewDeleteCredentialNoContent() *DeleteCredentialNoContent {
	return &DeleteCredentialNoContent{}
}

/*DeleteCredentialNoContent handles this case with default header values.

Success. The credentials set has been deleted. Response doesn't have a body.
*/
type DeleteCredentialNoContent struct {
}

func (o *DeleteCredentialNoContent) Error() string {
	return fmt.Sprintf("[DELETE /credentials/{id}][%d] deleteCredentialNoContent ", 204)
}

func (o *DeleteCredentialNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}