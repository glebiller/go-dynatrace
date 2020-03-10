// Code generated by go-swagger; DO NOT EDIT.

package dynatrace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// CompareOperation The condition of the rule.
//
// The actual set of fields depends on the `type` of the condition.
// swagger:discriminator CompareOperation type
type CompareOperation interface {
	runtime.Validatable

	// Defines the actual set of fields depending on the value:
	//
	// EQUALS -> EqualsCompareOperation
	// STRING_CONTAINS -> StringContainsCompareOperation
	// STARTS_WITH -> StartsWithCompareOperation
	// ENDS_WITH -> EndsWithCompareOperation
	// EXISTS -> ExistsCompareOperation
	// IP_IN_RANGE -> IpInRangeCompareOperation
	// LESS_THAN -> LessThanCompareOperation
	// GREATER_THAN -> GreaterThanCompareOperation
	// INT_EQUALS -> IntEqualsCompareOperation
	// STRING_EQUALS -> StringEqualsCompareOperation
	// TAG -> TagCompareOperation
	//
	// Required: true
	// Enum: [EQUALS STRING_CONTAINS STARTS_WITH ENDS_WITH EXISTS IP_IN_RANGE LESS_THAN GREATER_THAN INT_EQUALS STRING_EQUALS TAG]
	Type() string
	SetType(string)
}

type compareOperation struct {
	typeField string
}

// Type gets the type of this polymorphic type
func (m *compareOperation) Type() string {
	return "CompareOperation"
}

// SetType sets the type of this polymorphic type
func (m *compareOperation) SetType(val string) {

}

// UnmarshalCompareOperationSlice unmarshals polymorphic slices of CompareOperation
func UnmarshalCompareOperationSlice(reader io.Reader, consumer runtime.Consumer) ([]CompareOperation, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []CompareOperation
	for _, element := range elements {
		obj, err := unmarshalCompareOperation(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalCompareOperation unmarshals polymorphic CompareOperation
func UnmarshalCompareOperation(reader io.Reader, consumer runtime.Consumer) (CompareOperation, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalCompareOperation(data, consumer)
}

func unmarshalCompareOperation(data []byte, consumer runtime.Consumer) (CompareOperation, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the type property.
	var getType struct {
		Type string `json:"type"`
	}
	if err := consumer.Consume(buf, &getType); err != nil {
		return nil, err
	}

	if err := validate.RequiredString("type", "body", getType.Type); err != nil {
		return nil, err
	}

	// The value of type is used to determine which type to create and unmarshal the data into
	switch getType.Type {
	case "CompareOperation":
		var result compareOperation
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "EndsWithCompareOperation":
		var result EndsWithCompareOperation
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "EqualsCompareOperation":
		var result EqualsCompareOperation
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "ExistsCompareOperation":
		var result ExistsCompareOperation
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "GreaterThanCompareOperation":
		var result GreaterThanCompareOperation
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "IntEqualsCompareOperation":
		var result IntEqualsCompareOperation
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "IpInRangeCompareOperation":
		var result IPInRangeCompareOperation
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "LessThanCompareOperation":
		var result LessThanCompareOperation
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "StartsWithCompareOperation":
		var result StartsWithCompareOperation
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "StringContainsCompareOperation":
		var result StringContainsCompareOperation
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "StringEqualsCompareOperation":
		var result StringEqualsCompareOperation
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "TagCompareOperation":
		var result TagCompareOperation
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	}
	return nil, errors.New(422, "invalid type value: %q", getType.Type)

}

// Validate validates this compare operation
func (m *compareOperation) Validate(formats strfmt.Registry) error {
	return nil
}