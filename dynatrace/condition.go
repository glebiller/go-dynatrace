// Code generated by go-swagger; DO NOT EDIT.

package dynatrace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Condition A condition of a rule usage.
// swagger:model Condition
type Condition struct {

	// The attribute to be matched.
	// Required: true
	// Enum: [ACTOR_SYSTEM AKKA_ACTOR_CLASS_NAME AKKA_ACTOR_MESSAGE_TYPE AKKA_ACTOR_PATH AZURE_FUNCTIONS_FUNCTION_NAME AZURE_FUNCTIONS_SITE_NAME CICS_PROGRAM_NAME CICS_SYSTEM_ID CICS_TASK_ID CICS_TRANSACTION_ID CPU_TIME CTG_GATEWAY_URL CTG_PROGRAM CTG_SERVER_NAME CTG_TRANSACTION_ID CUSTOMSERVICE_CLASS CUSTOMSERVICE_METHOD DATABASE_CHILD_CALL_COUNT DATABASE_CHILD_CALL_TIME ERROR_COUNT ESB_APPLICATION_NAME ESB_INPUT_TYPE ESB_LIBRARY_NAME ESB_MESSAGE_FLOW_NAME EXCEPTION_CLASS EXCEPTION_MESSAGE FAILED_STATE FAILURE_REASON FLAW_STATE HTTP_REQUEST_METHOD HTTP_STATUS HTTP_STATUS_CLASS IMS_PROGRAM_NAME IMS_TRANSACTION_ID IO_TIME LOCK_TIME MESSAGING_DESTINATION_TYPE MESSAGING_IS_TEMPORARY_QUEUE MESSAGING_QUEUE_NAME MESSAGING_QUEUE_VENDOR NON_DATABASE_CHILD_CALL_COUNT NON_DATABASE_CHILD_CALL_TIME PROCESS_GROUP_NAME PROCESS_GROUP_TAG REMOTE_ENDPOINT REMOTE_METHOD REMOTE_SERVICE_NAME REQUEST_NAME REQUEST_TYPE RESPONSE_TIME RESPONSE_TIME_CLIENT RMI_CLASS RMI_METHOD SERVICE_DISPLAY_NAME SERVICE_NAME SERVICE_PORT SERVICE_PUBLIC_DOMAIN_NAME SERVICE_REQUEST_ATTRIBUTE SERVICE_TAG SERVICE_TYPE SERVICE_WEB_APPLICATION_ID SERVICE_WEB_CONTEXT_ROOT SERVICE_WEB_SERVER_NAME SERVICE_WEB_SERVICE_NAME SERVICE_WEB_SERVICE_NAMESPACE SUSPENSION_TIME TOTAL_PROCESSING_TIME WAIT_TIME WEBREQUEST_QUERY WEBREQUEST_URL WEBREQUEST_URL_HOST WEBREQUEST_URL_PATH WEBREQUEST_URL_PORT WEBSERVICE_ENDPOINT WEBSERVICE_METHOD ZOS_CALL_TYPE]
	Attribute *string `json:"attribute"`

	comparisonInfoField ComparisonInfo
}

// ComparisonInfo gets the comparison info of this base type
func (m *Condition) ComparisonInfo() ComparisonInfo {
	return m.comparisonInfoField
}

// SetComparisonInfo sets the comparison info of this base type
func (m *Condition) SetComparisonInfo(val ComparisonInfo) {
	m.comparisonInfoField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *Condition) UnmarshalJSON(raw []byte) error {
	var data struct {
		Attribute *string `json:"attribute"`

		ComparisonInfo json.RawMessage `json:"comparisonInfo"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	propComparisonInfo, err := UnmarshalComparisonInfo(bytes.NewBuffer(data.ComparisonInfo), runtime.JSONConsumer())
	if err != nil && err != io.EOF {
		return err
	}

	var result Condition

	// attribute
	result.Attribute = data.Attribute

	// comparisonInfo
	result.comparisonInfoField = propComparisonInfo

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m Condition) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
		Attribute *string `json:"attribute"`
	}{

		Attribute: m.Attribute,
	},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		ComparisonInfo ComparisonInfo `json:"comparisonInfo"`
	}{

		ComparisonInfo: m.comparisonInfoField,
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this condition
func (m *Condition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttribute(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComparisonInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var conditionTypeAttributePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ACTOR_SYSTEM","AKKA_ACTOR_CLASS_NAME","AKKA_ACTOR_MESSAGE_TYPE","AKKA_ACTOR_PATH","AZURE_FUNCTIONS_FUNCTION_NAME","AZURE_FUNCTIONS_SITE_NAME","CICS_PROGRAM_NAME","CICS_SYSTEM_ID","CICS_TASK_ID","CICS_TRANSACTION_ID","CPU_TIME","CTG_GATEWAY_URL","CTG_PROGRAM","CTG_SERVER_NAME","CTG_TRANSACTION_ID","CUSTOMSERVICE_CLASS","CUSTOMSERVICE_METHOD","DATABASE_CHILD_CALL_COUNT","DATABASE_CHILD_CALL_TIME","ERROR_COUNT","ESB_APPLICATION_NAME","ESB_INPUT_TYPE","ESB_LIBRARY_NAME","ESB_MESSAGE_FLOW_NAME","EXCEPTION_CLASS","EXCEPTION_MESSAGE","FAILED_STATE","FAILURE_REASON","FLAW_STATE","HTTP_REQUEST_METHOD","HTTP_STATUS","HTTP_STATUS_CLASS","IMS_PROGRAM_NAME","IMS_TRANSACTION_ID","IO_TIME","LOCK_TIME","MESSAGING_DESTINATION_TYPE","MESSAGING_IS_TEMPORARY_QUEUE","MESSAGING_QUEUE_NAME","MESSAGING_QUEUE_VENDOR","NON_DATABASE_CHILD_CALL_COUNT","NON_DATABASE_CHILD_CALL_TIME","PROCESS_GROUP_NAME","PROCESS_GROUP_TAG","REMOTE_ENDPOINT","REMOTE_METHOD","REMOTE_SERVICE_NAME","REQUEST_NAME","REQUEST_TYPE","RESPONSE_TIME","RESPONSE_TIME_CLIENT","RMI_CLASS","RMI_METHOD","SERVICE_DISPLAY_NAME","SERVICE_NAME","SERVICE_PORT","SERVICE_PUBLIC_DOMAIN_NAME","SERVICE_REQUEST_ATTRIBUTE","SERVICE_TAG","SERVICE_TYPE","SERVICE_WEB_APPLICATION_ID","SERVICE_WEB_CONTEXT_ROOT","SERVICE_WEB_SERVER_NAME","SERVICE_WEB_SERVICE_NAME","SERVICE_WEB_SERVICE_NAMESPACE","SUSPENSION_TIME","TOTAL_PROCESSING_TIME","WAIT_TIME","WEBREQUEST_QUERY","WEBREQUEST_URL","WEBREQUEST_URL_HOST","WEBREQUEST_URL_PATH","WEBREQUEST_URL_PORT","WEBSERVICE_ENDPOINT","WEBSERVICE_METHOD","ZOS_CALL_TYPE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		conditionTypeAttributePropEnum = append(conditionTypeAttributePropEnum, v)
	}
}

const (

	// ConditionAttributeACTORSYSTEM captures enum value "ACTOR_SYSTEM"
	ConditionAttributeACTORSYSTEM string = "ACTOR_SYSTEM"

	// ConditionAttributeAKKAACTORCLASSNAME captures enum value "AKKA_ACTOR_CLASS_NAME"
	ConditionAttributeAKKAACTORCLASSNAME string = "AKKA_ACTOR_CLASS_NAME"

	// ConditionAttributeAKKAACTORMESSAGETYPE captures enum value "AKKA_ACTOR_MESSAGE_TYPE"
	ConditionAttributeAKKAACTORMESSAGETYPE string = "AKKA_ACTOR_MESSAGE_TYPE"

	// ConditionAttributeAKKAACTORPATH captures enum value "AKKA_ACTOR_PATH"
	ConditionAttributeAKKAACTORPATH string = "AKKA_ACTOR_PATH"

	// ConditionAttributeAZUREFUNCTIONSFUNCTIONNAME captures enum value "AZURE_FUNCTIONS_FUNCTION_NAME"
	ConditionAttributeAZUREFUNCTIONSFUNCTIONNAME string = "AZURE_FUNCTIONS_FUNCTION_NAME"

	// ConditionAttributeAZUREFUNCTIONSSITENAME captures enum value "AZURE_FUNCTIONS_SITE_NAME"
	ConditionAttributeAZUREFUNCTIONSSITENAME string = "AZURE_FUNCTIONS_SITE_NAME"

	// ConditionAttributeCICSPROGRAMNAME captures enum value "CICS_PROGRAM_NAME"
	ConditionAttributeCICSPROGRAMNAME string = "CICS_PROGRAM_NAME"

	// ConditionAttributeCICSSYSTEMID captures enum value "CICS_SYSTEM_ID"
	ConditionAttributeCICSSYSTEMID string = "CICS_SYSTEM_ID"

	// ConditionAttributeCICSTASKID captures enum value "CICS_TASK_ID"
	ConditionAttributeCICSTASKID string = "CICS_TASK_ID"

	// ConditionAttributeCICSTRANSACTIONID captures enum value "CICS_TRANSACTION_ID"
	ConditionAttributeCICSTRANSACTIONID string = "CICS_TRANSACTION_ID"

	// ConditionAttributeCPUTIME captures enum value "CPU_TIME"
	ConditionAttributeCPUTIME string = "CPU_TIME"

	// ConditionAttributeCTGGATEWAYURL captures enum value "CTG_GATEWAY_URL"
	ConditionAttributeCTGGATEWAYURL string = "CTG_GATEWAY_URL"

	// ConditionAttributeCTGPROGRAM captures enum value "CTG_PROGRAM"
	ConditionAttributeCTGPROGRAM string = "CTG_PROGRAM"

	// ConditionAttributeCTGSERVERNAME captures enum value "CTG_SERVER_NAME"
	ConditionAttributeCTGSERVERNAME string = "CTG_SERVER_NAME"

	// ConditionAttributeCTGTRANSACTIONID captures enum value "CTG_TRANSACTION_ID"
	ConditionAttributeCTGTRANSACTIONID string = "CTG_TRANSACTION_ID"

	// ConditionAttributeCUSTOMSERVICECLASS captures enum value "CUSTOMSERVICE_CLASS"
	ConditionAttributeCUSTOMSERVICECLASS string = "CUSTOMSERVICE_CLASS"

	// ConditionAttributeCUSTOMSERVICEMETHOD captures enum value "CUSTOMSERVICE_METHOD"
	ConditionAttributeCUSTOMSERVICEMETHOD string = "CUSTOMSERVICE_METHOD"

	// ConditionAttributeDATABASECHILDCALLCOUNT captures enum value "DATABASE_CHILD_CALL_COUNT"
	ConditionAttributeDATABASECHILDCALLCOUNT string = "DATABASE_CHILD_CALL_COUNT"

	// ConditionAttributeDATABASECHILDCALLTIME captures enum value "DATABASE_CHILD_CALL_TIME"
	ConditionAttributeDATABASECHILDCALLTIME string = "DATABASE_CHILD_CALL_TIME"

	// ConditionAttributeERRORCOUNT captures enum value "ERROR_COUNT"
	ConditionAttributeERRORCOUNT string = "ERROR_COUNT"

	// ConditionAttributeESBAPPLICATIONNAME captures enum value "ESB_APPLICATION_NAME"
	ConditionAttributeESBAPPLICATIONNAME string = "ESB_APPLICATION_NAME"

	// ConditionAttributeESBINPUTTYPE captures enum value "ESB_INPUT_TYPE"
	ConditionAttributeESBINPUTTYPE string = "ESB_INPUT_TYPE"

	// ConditionAttributeESBLIBRARYNAME captures enum value "ESB_LIBRARY_NAME"
	ConditionAttributeESBLIBRARYNAME string = "ESB_LIBRARY_NAME"

	// ConditionAttributeESBMESSAGEFLOWNAME captures enum value "ESB_MESSAGE_FLOW_NAME"
	ConditionAttributeESBMESSAGEFLOWNAME string = "ESB_MESSAGE_FLOW_NAME"

	// ConditionAttributeEXCEPTIONCLASS captures enum value "EXCEPTION_CLASS"
	ConditionAttributeEXCEPTIONCLASS string = "EXCEPTION_CLASS"

	// ConditionAttributeEXCEPTIONMESSAGE captures enum value "EXCEPTION_MESSAGE"
	ConditionAttributeEXCEPTIONMESSAGE string = "EXCEPTION_MESSAGE"

	// ConditionAttributeFAILEDSTATE captures enum value "FAILED_STATE"
	ConditionAttributeFAILEDSTATE string = "FAILED_STATE"

	// ConditionAttributeFAILUREREASON captures enum value "FAILURE_REASON"
	ConditionAttributeFAILUREREASON string = "FAILURE_REASON"

	// ConditionAttributeFLAWSTATE captures enum value "FLAW_STATE"
	ConditionAttributeFLAWSTATE string = "FLAW_STATE"

	// ConditionAttributeHTTPREQUESTMETHOD captures enum value "HTTP_REQUEST_METHOD"
	ConditionAttributeHTTPREQUESTMETHOD string = "HTTP_REQUEST_METHOD"

	// ConditionAttributeHTTPSTATUS captures enum value "HTTP_STATUS"
	ConditionAttributeHTTPSTATUS string = "HTTP_STATUS"

	// ConditionAttributeHTTPSTATUSCLASS captures enum value "HTTP_STATUS_CLASS"
	ConditionAttributeHTTPSTATUSCLASS string = "HTTP_STATUS_CLASS"

	// ConditionAttributeIMSPROGRAMNAME captures enum value "IMS_PROGRAM_NAME"
	ConditionAttributeIMSPROGRAMNAME string = "IMS_PROGRAM_NAME"

	// ConditionAttributeIMSTRANSACTIONID captures enum value "IMS_TRANSACTION_ID"
	ConditionAttributeIMSTRANSACTIONID string = "IMS_TRANSACTION_ID"

	// ConditionAttributeIOTIME captures enum value "IO_TIME"
	ConditionAttributeIOTIME string = "IO_TIME"

	// ConditionAttributeLOCKTIME captures enum value "LOCK_TIME"
	ConditionAttributeLOCKTIME string = "LOCK_TIME"

	// ConditionAttributeMESSAGINGDESTINATIONTYPE captures enum value "MESSAGING_DESTINATION_TYPE"
	ConditionAttributeMESSAGINGDESTINATIONTYPE string = "MESSAGING_DESTINATION_TYPE"

	// ConditionAttributeMESSAGINGISTEMPORARYQUEUE captures enum value "MESSAGING_IS_TEMPORARY_QUEUE"
	ConditionAttributeMESSAGINGISTEMPORARYQUEUE string = "MESSAGING_IS_TEMPORARY_QUEUE"

	// ConditionAttributeMESSAGINGQUEUENAME captures enum value "MESSAGING_QUEUE_NAME"
	ConditionAttributeMESSAGINGQUEUENAME string = "MESSAGING_QUEUE_NAME"

	// ConditionAttributeMESSAGINGQUEUEVENDOR captures enum value "MESSAGING_QUEUE_VENDOR"
	ConditionAttributeMESSAGINGQUEUEVENDOR string = "MESSAGING_QUEUE_VENDOR"

	// ConditionAttributeNONDATABASECHILDCALLCOUNT captures enum value "NON_DATABASE_CHILD_CALL_COUNT"
	ConditionAttributeNONDATABASECHILDCALLCOUNT string = "NON_DATABASE_CHILD_CALL_COUNT"

	// ConditionAttributeNONDATABASECHILDCALLTIME captures enum value "NON_DATABASE_CHILD_CALL_TIME"
	ConditionAttributeNONDATABASECHILDCALLTIME string = "NON_DATABASE_CHILD_CALL_TIME"

	// ConditionAttributePROCESSGROUPNAME captures enum value "PROCESS_GROUP_NAME"
	ConditionAttributePROCESSGROUPNAME string = "PROCESS_GROUP_NAME"

	// ConditionAttributePROCESSGROUPTAG captures enum value "PROCESS_GROUP_TAG"
	ConditionAttributePROCESSGROUPTAG string = "PROCESS_GROUP_TAG"

	// ConditionAttributeREMOTEENDPOINT captures enum value "REMOTE_ENDPOINT"
	ConditionAttributeREMOTEENDPOINT string = "REMOTE_ENDPOINT"

	// ConditionAttributeREMOTEMETHOD captures enum value "REMOTE_METHOD"
	ConditionAttributeREMOTEMETHOD string = "REMOTE_METHOD"

	// ConditionAttributeREMOTESERVICENAME captures enum value "REMOTE_SERVICE_NAME"
	ConditionAttributeREMOTESERVICENAME string = "REMOTE_SERVICE_NAME"

	// ConditionAttributeREQUESTNAME captures enum value "REQUEST_NAME"
	ConditionAttributeREQUESTNAME string = "REQUEST_NAME"

	// ConditionAttributeREQUESTTYPE captures enum value "REQUEST_TYPE"
	ConditionAttributeREQUESTTYPE string = "REQUEST_TYPE"

	// ConditionAttributeRESPONSETIME captures enum value "RESPONSE_TIME"
	ConditionAttributeRESPONSETIME string = "RESPONSE_TIME"

	// ConditionAttributeRESPONSETIMECLIENT captures enum value "RESPONSE_TIME_CLIENT"
	ConditionAttributeRESPONSETIMECLIENT string = "RESPONSE_TIME_CLIENT"

	// ConditionAttributeRMICLASS captures enum value "RMI_CLASS"
	ConditionAttributeRMICLASS string = "RMI_CLASS"

	// ConditionAttributeRMIMETHOD captures enum value "RMI_METHOD"
	ConditionAttributeRMIMETHOD string = "RMI_METHOD"

	// ConditionAttributeSERVICEDISPLAYNAME captures enum value "SERVICE_DISPLAY_NAME"
	ConditionAttributeSERVICEDISPLAYNAME string = "SERVICE_DISPLAY_NAME"

	// ConditionAttributeSERVICENAME captures enum value "SERVICE_NAME"
	ConditionAttributeSERVICENAME string = "SERVICE_NAME"

	// ConditionAttributeSERVICEPORT captures enum value "SERVICE_PORT"
	ConditionAttributeSERVICEPORT string = "SERVICE_PORT"

	// ConditionAttributeSERVICEPUBLICDOMAINNAME captures enum value "SERVICE_PUBLIC_DOMAIN_NAME"
	ConditionAttributeSERVICEPUBLICDOMAINNAME string = "SERVICE_PUBLIC_DOMAIN_NAME"

	// ConditionAttributeSERVICEREQUESTATTRIBUTE captures enum value "SERVICE_REQUEST_ATTRIBUTE"
	ConditionAttributeSERVICEREQUESTATTRIBUTE string = "SERVICE_REQUEST_ATTRIBUTE"

	// ConditionAttributeSERVICETAG captures enum value "SERVICE_TAG"
	ConditionAttributeSERVICETAG string = "SERVICE_TAG"

	// ConditionAttributeSERVICETYPE captures enum value "SERVICE_TYPE"
	ConditionAttributeSERVICETYPE string = "SERVICE_TYPE"

	// ConditionAttributeSERVICEWEBAPPLICATIONID captures enum value "SERVICE_WEB_APPLICATION_ID"
	ConditionAttributeSERVICEWEBAPPLICATIONID string = "SERVICE_WEB_APPLICATION_ID"

	// ConditionAttributeSERVICEWEBCONTEXTROOT captures enum value "SERVICE_WEB_CONTEXT_ROOT"
	ConditionAttributeSERVICEWEBCONTEXTROOT string = "SERVICE_WEB_CONTEXT_ROOT"

	// ConditionAttributeSERVICEWEBSERVERNAME captures enum value "SERVICE_WEB_SERVER_NAME"
	ConditionAttributeSERVICEWEBSERVERNAME string = "SERVICE_WEB_SERVER_NAME"

	// ConditionAttributeSERVICEWEBSERVICENAME captures enum value "SERVICE_WEB_SERVICE_NAME"
	ConditionAttributeSERVICEWEBSERVICENAME string = "SERVICE_WEB_SERVICE_NAME"

	// ConditionAttributeSERVICEWEBSERVICENAMESPACE captures enum value "SERVICE_WEB_SERVICE_NAMESPACE"
	ConditionAttributeSERVICEWEBSERVICENAMESPACE string = "SERVICE_WEB_SERVICE_NAMESPACE"

	// ConditionAttributeSUSPENSIONTIME captures enum value "SUSPENSION_TIME"
	ConditionAttributeSUSPENSIONTIME string = "SUSPENSION_TIME"

	// ConditionAttributeTOTALPROCESSINGTIME captures enum value "TOTAL_PROCESSING_TIME"
	ConditionAttributeTOTALPROCESSINGTIME string = "TOTAL_PROCESSING_TIME"

	// ConditionAttributeWAITTIME captures enum value "WAIT_TIME"
	ConditionAttributeWAITTIME string = "WAIT_TIME"

	// ConditionAttributeWEBREQUESTQUERY captures enum value "WEBREQUEST_QUERY"
	ConditionAttributeWEBREQUESTQUERY string = "WEBREQUEST_QUERY"

	// ConditionAttributeWEBREQUESTURL captures enum value "WEBREQUEST_URL"
	ConditionAttributeWEBREQUESTURL string = "WEBREQUEST_URL"

	// ConditionAttributeWEBREQUESTURLHOST captures enum value "WEBREQUEST_URL_HOST"
	ConditionAttributeWEBREQUESTURLHOST string = "WEBREQUEST_URL_HOST"

	// ConditionAttributeWEBREQUESTURLPATH captures enum value "WEBREQUEST_URL_PATH"
	ConditionAttributeWEBREQUESTURLPATH string = "WEBREQUEST_URL_PATH"

	// ConditionAttributeWEBREQUESTURLPORT captures enum value "WEBREQUEST_URL_PORT"
	ConditionAttributeWEBREQUESTURLPORT string = "WEBREQUEST_URL_PORT"

	// ConditionAttributeWEBSERVICEENDPOINT captures enum value "WEBSERVICE_ENDPOINT"
	ConditionAttributeWEBSERVICEENDPOINT string = "WEBSERVICE_ENDPOINT"

	// ConditionAttributeWEBSERVICEMETHOD captures enum value "WEBSERVICE_METHOD"
	ConditionAttributeWEBSERVICEMETHOD string = "WEBSERVICE_METHOD"

	// ConditionAttributeZOSCALLTYPE captures enum value "ZOS_CALL_TYPE"
	ConditionAttributeZOSCALLTYPE string = "ZOS_CALL_TYPE"
)

// prop value enum
func (m *Condition) validateAttributeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, conditionTypeAttributePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Condition) validateAttribute(formats strfmt.Registry) error {

	if err := validate.Required("attribute", "body", m.Attribute); err != nil {
		return err
	}

	// value enum
	if err := m.validateAttributeEnum("attribute", "body", *m.Attribute); err != nil {
		return err
	}

	return nil
}

func (m *Condition) validateComparisonInfo(formats strfmt.Registry) error {

	if err := validate.Required("comparisonInfo", "body", m.ComparisonInfo()); err != nil {
		return err
	}

	if err := m.ComparisonInfo().Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("comparisonInfo")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Condition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Condition) UnmarshalBinary(b []byte) error {
	var res Condition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
