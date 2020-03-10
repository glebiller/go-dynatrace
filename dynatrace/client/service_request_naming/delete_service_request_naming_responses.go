// Code generated by go-swagger; DO NOT EDIT.

package service_request_naming

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteServiceRequestNamingReader is a Reader for the DeleteServiceRequestNaming structure.
type DeleteServiceRequestNamingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteServiceRequestNamingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteServiceRequestNamingNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteServiceRequestNamingNoContent creates a DeleteServiceRequestNamingNoContent with default headers values
func NewDeleteServiceRequestNamingNoContent() *DeleteServiceRequestNamingNoContent {
	return &DeleteServiceRequestNamingNoContent{}
}

/*DeleteServiceRequestNamingNoContent handles this case with default header values.

Success. The rule has been deleted. Response doesn't have a body.
*/
type DeleteServiceRequestNamingNoContent struct {
}

func (o *DeleteServiceRequestNamingNoContent) Error() string {
	return fmt.Sprintf("[DELETE /service/requestNaming/{id}][%d] deleteServiceRequestNamingNoContent ", 204)
}

func (o *DeleteServiceRequestNamingNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
