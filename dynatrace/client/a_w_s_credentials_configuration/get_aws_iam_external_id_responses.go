// Code generated by go-swagger; DO NOT EDIT.

package a_w_s_credentials_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/Kissy/go-dynatrace/dynatrace"
)

// GetAwsIamExternalIDReader is a Reader for the GetAwsIamExternalID structure.
type GetAwsIamExternalIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAwsIamExternalIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAwsIamExternalIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAwsIamExternalIDOK creates a GetAwsIamExternalIDOK with default headers values
func NewGetAwsIamExternalIDOK() *GetAwsIamExternalIDOK {
	return &GetAwsIamExternalIDOK{}
}

/*GetAwsIamExternalIDOK handles this case with default header values.

successful operation
*/
type GetAwsIamExternalIDOK struct {
	Payload *dynatrace.AwsIamToken
}

func (o *GetAwsIamExternalIDOK) Error() string {
	return fmt.Sprintf("[GET /aws/iamExternalId][%d] getAwsIamExternalIdOK  %+v", 200, o.Payload)
}

func (o *GetAwsIamExternalIDOK) GetPayload() *dynatrace.AwsIamToken {
	return o.Payload
}

func (o *GetAwsIamExternalIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.AwsIamToken)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
