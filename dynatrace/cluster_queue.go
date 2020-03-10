// Code generated by go-swagger; DO NOT EDIT.

package dynatrace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ClusterQueue This local queue is visible in these [clusters](https://www.ibm.com/support/knowledgecenter/en/SSFKSJ_7.5.0/com.ibm.mq.pro.doc/q002750_.htm). The queue manager must be part of these clusters.
// swagger:model ClusterQueue
type ClusterQueue struct {

	// The local queue is visible in these [clusters](https://www.ibm.com/support/knowledgecenter/en/SSFKSJ_7.5.0/com.ibm.mq.pro.doc/q002750_.htm). The queue manager must be part of these clusters.
	// Required: true
	// Max Items: 10000
	// Min Items: 1
	ClusterVisibility []string `json:"clusterVisibility"`

	// The name of the local queue.
	// Required: true
	// Max Length: 500
	// Min Length: 1
	LocalQueue *string `json:"localQueue"`
}

// Validate validates this cluster queue
func (m *ClusterQueue) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterVisibility(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalQueue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterQueue) validateClusterVisibility(formats strfmt.Registry) error {

	if err := validate.Required("clusterVisibility", "body", m.ClusterVisibility); err != nil {
		return err
	}

	iClusterVisibilitySize := int64(len(m.ClusterVisibility))

	if err := validate.MinItems("clusterVisibility", "body", iClusterVisibilitySize, 1); err != nil {
		return err
	}

	if err := validate.MaxItems("clusterVisibility", "body", iClusterVisibilitySize, 10000); err != nil {
		return err
	}

	return nil
}

func (m *ClusterQueue) validateLocalQueue(formats strfmt.Registry) error {

	if err := validate.Required("localQueue", "body", m.LocalQueue); err != nil {
		return err
	}

	if err := validate.MinLength("localQueue", "body", string(*m.LocalQueue), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("localQueue", "body", string(*m.LocalQueue), 500); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterQueue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterQueue) UnmarshalBinary(b []byte) error {
	var res ClusterQueue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
