// Code generated by go-swagger; DO NOT EDIT.

package notifications

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/Kissy/go-dynatrace/dynatrace"
)

// GetNotificationReader is a Reader for the GetNotification structure.
type GetNotificationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNotificationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNotificationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetNotificationOK creates a GetNotificationOK with default headers values
func NewGetNotificationOK() *GetNotificationOK {
	return &GetNotificationOK{}
}

/*GetNotificationOK handles this case with default header values.

successful operation
*/
type GetNotificationOK struct {
	Payload dynatrace.NotificationConfig
}

func (o *GetNotificationOK) Error() string {
	return fmt.Sprintf("[GET /notifications/{id}][%d] getNotificationOK  %+v", 200, o.Payload)
}

func (o *GetNotificationOK) GetPayload() dynatrace.NotificationConfig {
	return o.Payload
}

func (o *GetNotificationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := dynatrace.UnmarshalNotificationConfig(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}
