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

// ComparisonInfo Type-specific comparison for attributes. The actual set of fields depends on the `type` of the comparison.
// swagger:discriminator ComparisonInfo type
type ComparisonInfo interface {
	runtime.Validatable

	// Operator of the comparision. You can reverse it by setting **negate** to `true`.
	// Required: true
	Comparison() Enum
	SetComparison(Enum)

	// Reverse the comparison **operator**. For example, it turns **equals** into **does not equal**.
	// Required: true
	Negate() *bool
	SetNegate(*bool)

	// Defines the actual set of fields depending on the value:
	//
	// STRING -> StringComparisonInfo
	// NUMBER -> NumberComparisonInfo
	// BOOLEAN -> BooleanComparisonInfo
	// HTTP_METHOD -> HttpMethodComparisonInfo
	// STRING_REQUEST_ATTRIBUTE -> StringRequestAttributeComparisonInfo
	// NUMBER_REQUEST_ATTRIBUTE -> NumberRequestAttributeComparisonInfo
	// ZOS_CALL_TYPE -> ZosComparisonInfo
	// IIB_INPUT_NODE_TYPE -> IIBInputNodeTypeComparisonInfo
	// ESB_INPUT_NODE_TYPE -> ESBInputNodeTypeComparisonInfo
	// FAILED_STATE -> FailedStateComparisonInfo
	// FLAW_STATE -> FlawStateComparisonInfo
	// FAILURE_REASON -> FailureReasonComparisonInfo
	// HTTP_STATUS_CLASS -> HttpStatusClassComparisonInfo
	// TAG -> TagComparisonInfo
	// FAST_STRING -> FastStringComparisonInfo
	// SERVICE_TYPE -> ServiceTypeComparisonInfo
	//
	// Required: true
	// Enum: [STRING NUMBER BOOLEAN HTTP_METHOD STRING_REQUEST_ATTRIBUTE NUMBER_REQUEST_ATTRIBUTE ZOS_CALL_TYPE IIB_INPUT_NODE_TYPE ESB_INPUT_NODE_TYPE FAILED_STATE FLAW_STATE FAILURE_REASON HTTP_STATUS_CLASS TAG FAST_STRING SERVICE_TYPE]
	Type() string
	SetType(string)

	// The value to compare to.
	Value() interface{}
	SetValue(interface{})
}

type comparisonInfo struct {
	comparisonField Enum

	negateField *bool

	typeField string

	valueField interface{}
}

// Comparison gets the comparison of this polymorphic type
func (m *comparisonInfo) Comparison() Enum {
	return m.comparisonField
}

// SetComparison sets the comparison of this polymorphic type
func (m *comparisonInfo) SetComparison(val Enum) {
	m.comparisonField = val
}

// Negate gets the negate of this polymorphic type
func (m *comparisonInfo) Negate() *bool {
	return m.negateField
}

// SetNegate sets the negate of this polymorphic type
func (m *comparisonInfo) SetNegate(val *bool) {
	m.negateField = val
}

// Type gets the type of this polymorphic type
func (m *comparisonInfo) Type() string {
	return "ComparisonInfo"
}

// SetType sets the type of this polymorphic type
func (m *comparisonInfo) SetType(val string) {

}

// Value gets the value of this polymorphic type
func (m *comparisonInfo) Value() interface{} {
	return m.valueField
}

// SetValue sets the value of this polymorphic type
func (m *comparisonInfo) SetValue(val interface{}) {
	m.valueField = val
}

// UnmarshalComparisonInfoSlice unmarshals polymorphic slices of ComparisonInfo
func UnmarshalComparisonInfoSlice(reader io.Reader, consumer runtime.Consumer) ([]ComparisonInfo, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []ComparisonInfo
	for _, element := range elements {
		obj, err := unmarshalComparisonInfo(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalComparisonInfo unmarshals polymorphic ComparisonInfo
func UnmarshalComparisonInfo(reader io.Reader, consumer runtime.Consumer) (ComparisonInfo, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalComparisonInfo(data, consumer)
}

func unmarshalComparisonInfo(data []byte, consumer runtime.Consumer) (ComparisonInfo, error) {
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
	case "BooleanComparisonInfo":
		var result BooleanComparisonInfo
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "ComparisonInfo":
		var result comparisonInfo
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "ESBInputNodeTypeComparisonInfo":
		var result ESBInputNodeTypeComparisonInfo
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "FailedStateComparisonInfo":
		var result FailedStateComparisonInfo
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "FailureReasonComparisonInfo":
		var result FailureReasonComparisonInfo
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "FastStringComparisonInfo":
		var result FastStringComparisonInfo
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "FlawStateComparisonInfo":
		var result FlawStateComparisonInfo
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "HttpMethodComparisonInfo":
		var result HTTPMethodComparisonInfo
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "HttpStatusClassComparisonInfo":
		var result HTTPStatusClassComparisonInfo
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "IIBInputNodeTypeComparisonInfo":
		var result IIBInputNodeTypeComparisonInfo
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "NumberComparisonInfo":
		var result NumberComparisonInfo
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "NumberRequestAttributeComparisonInfo":
		var result NumberRequestAttributeComparisonInfo
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "ServiceTypeComparisonInfo":
		var result ServiceTypeComparisonInfo
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "StringComparisonInfo":
		var result StringComparisonInfo
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "StringRequestAttributeComparisonInfo":
		var result StringRequestAttributeComparisonInfo
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "TagComparisonInfo":
		var result TagComparisonInfo
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "ZosComparisonInfo":
		var result ZosComparisonInfo
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	}
	return nil, errors.New(422, "invalid type value: %q", getType.Type)

}

// Validate validates this comparison info
func (m *comparisonInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComparison(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNegate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *comparisonInfo) validateComparison(formats strfmt.Registry) error {

	if err := validate.Required("comparison", "body", m.Comparison()); err != nil {
		return err
	}

	return nil
}

func (m *comparisonInfo) validateNegate(formats strfmt.Registry) error {

	if err := validate.Required("negate", "body", m.Negate()); err != nil {
		return err
	}

	return nil
}

var comparisonInfoTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["STRING","NUMBER","BOOLEAN","HTTP_METHOD","STRING_REQUEST_ATTRIBUTE","NUMBER_REQUEST_ATTRIBUTE","ZOS_CALL_TYPE","IIB_INPUT_NODE_TYPE","ESB_INPUT_NODE_TYPE","FAILED_STATE","FLAW_STATE","FAILURE_REASON","HTTP_STATUS_CLASS","TAG","FAST_STRING","SERVICE_TYPE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		comparisonInfoTypeTypePropEnum = append(comparisonInfoTypeTypePropEnum, v)
	}
}

const (

	// ComparisonInfoTypeSTRING captures enum value "STRING"
	ComparisonInfoTypeSTRING string = "STRING"

	// ComparisonInfoTypeNUMBER captures enum value "NUMBER"
	ComparisonInfoTypeNUMBER string = "NUMBER"

	// ComparisonInfoTypeBOOLEAN captures enum value "BOOLEAN"
	ComparisonInfoTypeBOOLEAN string = "BOOLEAN"

	// ComparisonInfoTypeHTTPMETHOD captures enum value "HTTP_METHOD"
	ComparisonInfoTypeHTTPMETHOD string = "HTTP_METHOD"

	// ComparisonInfoTypeSTRINGREQUESTATTRIBUTE captures enum value "STRING_REQUEST_ATTRIBUTE"
	ComparisonInfoTypeSTRINGREQUESTATTRIBUTE string = "STRING_REQUEST_ATTRIBUTE"

	// ComparisonInfoTypeNUMBERREQUESTATTRIBUTE captures enum value "NUMBER_REQUEST_ATTRIBUTE"
	ComparisonInfoTypeNUMBERREQUESTATTRIBUTE string = "NUMBER_REQUEST_ATTRIBUTE"

	// ComparisonInfoTypeZOSCALLTYPE captures enum value "ZOS_CALL_TYPE"
	ComparisonInfoTypeZOSCALLTYPE string = "ZOS_CALL_TYPE"

	// ComparisonInfoTypeIIBINPUTNODETYPE captures enum value "IIB_INPUT_NODE_TYPE"
	ComparisonInfoTypeIIBINPUTNODETYPE string = "IIB_INPUT_NODE_TYPE"

	// ComparisonInfoTypeESBINPUTNODETYPE captures enum value "ESB_INPUT_NODE_TYPE"
	ComparisonInfoTypeESBINPUTNODETYPE string = "ESB_INPUT_NODE_TYPE"

	// ComparisonInfoTypeFAILEDSTATE captures enum value "FAILED_STATE"
	ComparisonInfoTypeFAILEDSTATE string = "FAILED_STATE"

	// ComparisonInfoTypeFLAWSTATE captures enum value "FLAW_STATE"
	ComparisonInfoTypeFLAWSTATE string = "FLAW_STATE"

	// ComparisonInfoTypeFAILUREREASON captures enum value "FAILURE_REASON"
	ComparisonInfoTypeFAILUREREASON string = "FAILURE_REASON"

	// ComparisonInfoTypeHTTPSTATUSCLASS captures enum value "HTTP_STATUS_CLASS"
	ComparisonInfoTypeHTTPSTATUSCLASS string = "HTTP_STATUS_CLASS"

	// ComparisonInfoTypeTAG captures enum value "TAG"
	ComparisonInfoTypeTAG string = "TAG"

	// ComparisonInfoTypeFASTSTRING captures enum value "FAST_STRING"
	ComparisonInfoTypeFASTSTRING string = "FAST_STRING"

	// ComparisonInfoTypeSERVICETYPE captures enum value "SERVICE_TYPE"
	ComparisonInfoTypeSERVICETYPE string = "SERVICE_TYPE"
)

// prop value enum
func (m *comparisonInfo) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, comparisonInfoTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}
