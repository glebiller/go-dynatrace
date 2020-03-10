// Code generated by go-swagger; DO NOT EDIT.

package dynatrace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DataSource data source
// swagger:model DataSource
type DataSource struct {

	// Specifies the location where the values are captured and stored.
	//
	//  Required if the **source** is one of the following: `GET_PARAMETER`, `URI`, `REQUEST_HEADER`, `RESPONSE_HEADER`.
	//
	//  Not applicable in other cases.
	//
	//  If the **source** value is `REQUEST_HEADER` or `RESPONSE_HEADER`, the `CAPTURE_AND_STORE_ON_BOTH` location is not allowed.
	// Enum: [CAPTURE_AND_STORE_ON_BOTH CAPTURE_AND_STORE_ON_CLIENT CAPTURE_AND_STORE_ON_SERVER CAPTURE_ON_CLIENT_STORE_ON_SERVER]
	CapturingAndStorageLocation string `json:"capturingAndStorageLocation,omitempty"`

	// CICS SDK node name condition for which the value is captured.
	//
	//  This  is required if the **source** is: `CICS_SDK`.
	//
	//  Not applicable in other cases.
	CicsSDKMethodNodeCondition *ValueCondition `json:"cicsSDKMethodNodeCondition,omitempty"`

	// The data source is enbled (`true`) or disabled (`false`).
	// Required: true
	Enabled *bool `json:"enabled"`

	// IBM integration bus node name condition for which the value is captured.
	//
	//  This or `iibNodeType` is required if the **source** is: `IIB_NODE`.
	//
	//  Not applicable in other cases.
	IibMethodNodeCondition *ValueCondition `json:"iibMethodNodeCondition,omitempty"`

	// The IBM integration bus node type for which the value is captured.
	//
	//  This or `iibMethodNodeCondition` is required if the **source** is: `IIB_NODE`.
	//
	//  Not applicable in other cases.
	// Enum: [AGGREGATE_CONTROL_NODE AGGREGATE_REPLY_NODE AGGREGATE_REQUEST_NODE CALLABLE_FLOW_REPLY_NODE COLLECTOR_NODE COMPUTE_NODE DATABASE_NODE DECISION_SERVICE_NODE DOT_NET_COMPUTE_NODE FILE_READ_NODE FILTER_NODE FLOW_ORDER_NODE GROUP_COMPLETE_NODE GROUP_GATHER_NODE GROUP_SCATTER_NODE HTTP_HEADER JAVA_COMPUTE_NODE JMS_CLIENT_RECEIVE JMS_CLIENT_REPLY_NODE JMS_HEADER MQ_GET_NODE MQ_OUTPUT_NODE PASSTHRU_NODE RESET_CONTENT_DESCRIPTOR_NODE RE_SEQUENCE_NODE ROUTE_NODE SAP_REPLY_NODE SCA_REPLY_NODE SECURITY_PEP SEQUENCE_NODE SOAP_EXTRACT_NODE SOAP_REPLY_NODE SOAP_WRAPPER_NODE SR_RETRIEVE_ENTITY_NODE SR_RETRIEVE_IT_SERVICE_NODE THROW_NODE TRACE_NODE TRY_CATCH_NODE VALIDATE_NODE WS_REPLY_NODE XSL_MQSI_NODE]
	IibNodeType string `json:"iibNodeType,omitempty"`

	// The method specification if the **source** value is `METHOD_PARAM`.
	//
	//  Not applicable in other cases.
	Methods []*CapturedMethod `json:"methods"`

	// The name of the web request parameter to capture.
	//
	//  Required if the **source** is one of the following: `POST_PARAMETER`, `GET_PARAMETER`, `REQUEST_HEADER`, `RESPONSE_HEADER`, `CUSTOM_ATTRIBUTE`.
	//
	//  Not applicable in other cases.
	ParameterName string `json:"parameterName,omitempty"`

	// Conditions for data capturing.
	Scope *ScopeConditions `json:"scope,omitempty"`

	// The technology of the session attribute to capture if the **source** value is `SESSION_ATTRIBUTE`. \n\n Not applicable in other cases.
	// Enum: [ASP_NET ASP_NET_CORE JAVA]
	SessionAttributeTechnology string `json:"sessionAttributeTechnology,omitempty"`

	// The source of the attribute to capture. Works in conjunction with **parameterName** or **methods** and **technology**.
	// Required: true
	// Enum: [CICS_SDK CLIENT_IP CUSTOM_ATTRIBUTE IIB_NODE METHOD_PARAM POST_PARAMETER QUERY_PARAMETER REQUEST_HEADER RESPONSE_HEADER SESSION_ATTRIBUTE URI URI_PATH]
	Source *string `json:"source"`

	// The technology of the method to capture if the **source** value is `METHOD_PARAM`. \n\n Not applicable in other cases.
	// Enum: [DOTNET JAVA PHP]
	Technology string `json:"technology,omitempty"`

	// Process values as specified.
	ValueProcessing *ValueProcessing `json:"valueProcessing,omitempty"`
}

// Validate validates this data source
func (m *DataSource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCapturingAndStorageLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCicsSDKMethodNodeCondition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIibMethodNodeCondition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIibNodeType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMethods(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScope(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSessionAttributeTechnology(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTechnology(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValueProcessing(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var dataSourceTypeCapturingAndStorageLocationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CAPTURE_AND_STORE_ON_BOTH","CAPTURE_AND_STORE_ON_CLIENT","CAPTURE_AND_STORE_ON_SERVER","CAPTURE_ON_CLIENT_STORE_ON_SERVER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		dataSourceTypeCapturingAndStorageLocationPropEnum = append(dataSourceTypeCapturingAndStorageLocationPropEnum, v)
	}
}

const (

	// DataSourceCapturingAndStorageLocationCAPTUREANDSTOREONBOTH captures enum value "CAPTURE_AND_STORE_ON_BOTH"
	DataSourceCapturingAndStorageLocationCAPTUREANDSTOREONBOTH string = "CAPTURE_AND_STORE_ON_BOTH"

	// DataSourceCapturingAndStorageLocationCAPTUREANDSTOREONCLIENT captures enum value "CAPTURE_AND_STORE_ON_CLIENT"
	DataSourceCapturingAndStorageLocationCAPTUREANDSTOREONCLIENT string = "CAPTURE_AND_STORE_ON_CLIENT"

	// DataSourceCapturingAndStorageLocationCAPTUREANDSTOREONSERVER captures enum value "CAPTURE_AND_STORE_ON_SERVER"
	DataSourceCapturingAndStorageLocationCAPTUREANDSTOREONSERVER string = "CAPTURE_AND_STORE_ON_SERVER"

	// DataSourceCapturingAndStorageLocationCAPTUREONCLIENTSTOREONSERVER captures enum value "CAPTURE_ON_CLIENT_STORE_ON_SERVER"
	DataSourceCapturingAndStorageLocationCAPTUREONCLIENTSTOREONSERVER string = "CAPTURE_ON_CLIENT_STORE_ON_SERVER"
)

// prop value enum
func (m *DataSource) validateCapturingAndStorageLocationEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, dataSourceTypeCapturingAndStorageLocationPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *DataSource) validateCapturingAndStorageLocation(formats strfmt.Registry) error {

	if swag.IsZero(m.CapturingAndStorageLocation) { // not required
		return nil
	}

	// value enum
	if err := m.validateCapturingAndStorageLocationEnum("capturingAndStorageLocation", "body", m.CapturingAndStorageLocation); err != nil {
		return err
	}

	return nil
}

func (m *DataSource) validateCicsSDKMethodNodeCondition(formats strfmt.Registry) error {

	if swag.IsZero(m.CicsSDKMethodNodeCondition) { // not required
		return nil
	}

	if m.CicsSDKMethodNodeCondition != nil {
		if err := m.CicsSDKMethodNodeCondition.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cicsSDKMethodNodeCondition")
			}
			return err
		}
	}

	return nil
}

func (m *DataSource) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

func (m *DataSource) validateIibMethodNodeCondition(formats strfmt.Registry) error {

	if swag.IsZero(m.IibMethodNodeCondition) { // not required
		return nil
	}

	if m.IibMethodNodeCondition != nil {
		if err := m.IibMethodNodeCondition.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("iibMethodNodeCondition")
			}
			return err
		}
	}

	return nil
}

var dataSourceTypeIibNodeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AGGREGATE_CONTROL_NODE","AGGREGATE_REPLY_NODE","AGGREGATE_REQUEST_NODE","CALLABLE_FLOW_REPLY_NODE","COLLECTOR_NODE","COMPUTE_NODE","DATABASE_NODE","DECISION_SERVICE_NODE","DOT_NET_COMPUTE_NODE","FILE_READ_NODE","FILTER_NODE","FLOW_ORDER_NODE","GROUP_COMPLETE_NODE","GROUP_GATHER_NODE","GROUP_SCATTER_NODE","HTTP_HEADER","JAVA_COMPUTE_NODE","JMS_CLIENT_RECEIVE","JMS_CLIENT_REPLY_NODE","JMS_HEADER","MQ_GET_NODE","MQ_OUTPUT_NODE","PASSTHRU_NODE","RESET_CONTENT_DESCRIPTOR_NODE","RE_SEQUENCE_NODE","ROUTE_NODE","SAP_REPLY_NODE","SCA_REPLY_NODE","SECURITY_PEP","SEQUENCE_NODE","SOAP_EXTRACT_NODE","SOAP_REPLY_NODE","SOAP_WRAPPER_NODE","SR_RETRIEVE_ENTITY_NODE","SR_RETRIEVE_IT_SERVICE_NODE","THROW_NODE","TRACE_NODE","TRY_CATCH_NODE","VALIDATE_NODE","WS_REPLY_NODE","XSL_MQSI_NODE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		dataSourceTypeIibNodeTypePropEnum = append(dataSourceTypeIibNodeTypePropEnum, v)
	}
}

const (

	// DataSourceIibNodeTypeAGGREGATECONTROLNODE captures enum value "AGGREGATE_CONTROL_NODE"
	DataSourceIibNodeTypeAGGREGATECONTROLNODE string = "AGGREGATE_CONTROL_NODE"

	// DataSourceIibNodeTypeAGGREGATEREPLYNODE captures enum value "AGGREGATE_REPLY_NODE"
	DataSourceIibNodeTypeAGGREGATEREPLYNODE string = "AGGREGATE_REPLY_NODE"

	// DataSourceIibNodeTypeAGGREGATEREQUESTNODE captures enum value "AGGREGATE_REQUEST_NODE"
	DataSourceIibNodeTypeAGGREGATEREQUESTNODE string = "AGGREGATE_REQUEST_NODE"

	// DataSourceIibNodeTypeCALLABLEFLOWREPLYNODE captures enum value "CALLABLE_FLOW_REPLY_NODE"
	DataSourceIibNodeTypeCALLABLEFLOWREPLYNODE string = "CALLABLE_FLOW_REPLY_NODE"

	// DataSourceIibNodeTypeCOLLECTORNODE captures enum value "COLLECTOR_NODE"
	DataSourceIibNodeTypeCOLLECTORNODE string = "COLLECTOR_NODE"

	// DataSourceIibNodeTypeCOMPUTENODE captures enum value "COMPUTE_NODE"
	DataSourceIibNodeTypeCOMPUTENODE string = "COMPUTE_NODE"

	// DataSourceIibNodeTypeDATABASENODE captures enum value "DATABASE_NODE"
	DataSourceIibNodeTypeDATABASENODE string = "DATABASE_NODE"

	// DataSourceIibNodeTypeDECISIONSERVICENODE captures enum value "DECISION_SERVICE_NODE"
	DataSourceIibNodeTypeDECISIONSERVICENODE string = "DECISION_SERVICE_NODE"

	// DataSourceIibNodeTypeDOTNETCOMPUTENODE captures enum value "DOT_NET_COMPUTE_NODE"
	DataSourceIibNodeTypeDOTNETCOMPUTENODE string = "DOT_NET_COMPUTE_NODE"

	// DataSourceIibNodeTypeFILEREADNODE captures enum value "FILE_READ_NODE"
	DataSourceIibNodeTypeFILEREADNODE string = "FILE_READ_NODE"

	// DataSourceIibNodeTypeFILTERNODE captures enum value "FILTER_NODE"
	DataSourceIibNodeTypeFILTERNODE string = "FILTER_NODE"

	// DataSourceIibNodeTypeFLOWORDERNODE captures enum value "FLOW_ORDER_NODE"
	DataSourceIibNodeTypeFLOWORDERNODE string = "FLOW_ORDER_NODE"

	// DataSourceIibNodeTypeGROUPCOMPLETENODE captures enum value "GROUP_COMPLETE_NODE"
	DataSourceIibNodeTypeGROUPCOMPLETENODE string = "GROUP_COMPLETE_NODE"

	// DataSourceIibNodeTypeGROUPGATHERNODE captures enum value "GROUP_GATHER_NODE"
	DataSourceIibNodeTypeGROUPGATHERNODE string = "GROUP_GATHER_NODE"

	// DataSourceIibNodeTypeGROUPSCATTERNODE captures enum value "GROUP_SCATTER_NODE"
	DataSourceIibNodeTypeGROUPSCATTERNODE string = "GROUP_SCATTER_NODE"

	// DataSourceIibNodeTypeHTTPHEADER captures enum value "HTTP_HEADER"
	DataSourceIibNodeTypeHTTPHEADER string = "HTTP_HEADER"

	// DataSourceIibNodeTypeJAVACOMPUTENODE captures enum value "JAVA_COMPUTE_NODE"
	DataSourceIibNodeTypeJAVACOMPUTENODE string = "JAVA_COMPUTE_NODE"

	// DataSourceIibNodeTypeJMSCLIENTRECEIVE captures enum value "JMS_CLIENT_RECEIVE"
	DataSourceIibNodeTypeJMSCLIENTRECEIVE string = "JMS_CLIENT_RECEIVE"

	// DataSourceIibNodeTypeJMSCLIENTREPLYNODE captures enum value "JMS_CLIENT_REPLY_NODE"
	DataSourceIibNodeTypeJMSCLIENTREPLYNODE string = "JMS_CLIENT_REPLY_NODE"

	// DataSourceIibNodeTypeJMSHEADER captures enum value "JMS_HEADER"
	DataSourceIibNodeTypeJMSHEADER string = "JMS_HEADER"

	// DataSourceIibNodeTypeMQGETNODE captures enum value "MQ_GET_NODE"
	DataSourceIibNodeTypeMQGETNODE string = "MQ_GET_NODE"

	// DataSourceIibNodeTypeMQOUTPUTNODE captures enum value "MQ_OUTPUT_NODE"
	DataSourceIibNodeTypeMQOUTPUTNODE string = "MQ_OUTPUT_NODE"

	// DataSourceIibNodeTypePASSTHRUNODE captures enum value "PASSTHRU_NODE"
	DataSourceIibNodeTypePASSTHRUNODE string = "PASSTHRU_NODE"

	// DataSourceIibNodeTypeRESETCONTENTDESCRIPTORNODE captures enum value "RESET_CONTENT_DESCRIPTOR_NODE"
	DataSourceIibNodeTypeRESETCONTENTDESCRIPTORNODE string = "RESET_CONTENT_DESCRIPTOR_NODE"

	// DataSourceIibNodeTypeRESEQUENCENODE captures enum value "RE_SEQUENCE_NODE"
	DataSourceIibNodeTypeRESEQUENCENODE string = "RE_SEQUENCE_NODE"

	// DataSourceIibNodeTypeROUTENODE captures enum value "ROUTE_NODE"
	DataSourceIibNodeTypeROUTENODE string = "ROUTE_NODE"

	// DataSourceIibNodeTypeSAPREPLYNODE captures enum value "SAP_REPLY_NODE"
	DataSourceIibNodeTypeSAPREPLYNODE string = "SAP_REPLY_NODE"

	// DataSourceIibNodeTypeSCAREPLYNODE captures enum value "SCA_REPLY_NODE"
	DataSourceIibNodeTypeSCAREPLYNODE string = "SCA_REPLY_NODE"

	// DataSourceIibNodeTypeSECURITYPEP captures enum value "SECURITY_PEP"
	DataSourceIibNodeTypeSECURITYPEP string = "SECURITY_PEP"

	// DataSourceIibNodeTypeSEQUENCENODE captures enum value "SEQUENCE_NODE"
	DataSourceIibNodeTypeSEQUENCENODE string = "SEQUENCE_NODE"

	// DataSourceIibNodeTypeSOAPEXTRACTNODE captures enum value "SOAP_EXTRACT_NODE"
	DataSourceIibNodeTypeSOAPEXTRACTNODE string = "SOAP_EXTRACT_NODE"

	// DataSourceIibNodeTypeSOAPREPLYNODE captures enum value "SOAP_REPLY_NODE"
	DataSourceIibNodeTypeSOAPREPLYNODE string = "SOAP_REPLY_NODE"

	// DataSourceIibNodeTypeSOAPWRAPPERNODE captures enum value "SOAP_WRAPPER_NODE"
	DataSourceIibNodeTypeSOAPWRAPPERNODE string = "SOAP_WRAPPER_NODE"

	// DataSourceIibNodeTypeSRRETRIEVEENTITYNODE captures enum value "SR_RETRIEVE_ENTITY_NODE"
	DataSourceIibNodeTypeSRRETRIEVEENTITYNODE string = "SR_RETRIEVE_ENTITY_NODE"

	// DataSourceIibNodeTypeSRRETRIEVEITSERVICENODE captures enum value "SR_RETRIEVE_IT_SERVICE_NODE"
	DataSourceIibNodeTypeSRRETRIEVEITSERVICENODE string = "SR_RETRIEVE_IT_SERVICE_NODE"

	// DataSourceIibNodeTypeTHROWNODE captures enum value "THROW_NODE"
	DataSourceIibNodeTypeTHROWNODE string = "THROW_NODE"

	// DataSourceIibNodeTypeTRACENODE captures enum value "TRACE_NODE"
	DataSourceIibNodeTypeTRACENODE string = "TRACE_NODE"

	// DataSourceIibNodeTypeTRYCATCHNODE captures enum value "TRY_CATCH_NODE"
	DataSourceIibNodeTypeTRYCATCHNODE string = "TRY_CATCH_NODE"

	// DataSourceIibNodeTypeVALIDATENODE captures enum value "VALIDATE_NODE"
	DataSourceIibNodeTypeVALIDATENODE string = "VALIDATE_NODE"

	// DataSourceIibNodeTypeWSREPLYNODE captures enum value "WS_REPLY_NODE"
	DataSourceIibNodeTypeWSREPLYNODE string = "WS_REPLY_NODE"

	// DataSourceIibNodeTypeXSLMQSINODE captures enum value "XSL_MQSI_NODE"
	DataSourceIibNodeTypeXSLMQSINODE string = "XSL_MQSI_NODE"
)

// prop value enum
func (m *DataSource) validateIibNodeTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, dataSourceTypeIibNodeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *DataSource) validateIibNodeType(formats strfmt.Registry) error {

	if swag.IsZero(m.IibNodeType) { // not required
		return nil
	}

	// value enum
	if err := m.validateIibNodeTypeEnum("iibNodeType", "body", m.IibNodeType); err != nil {
		return err
	}

	return nil
}

func (m *DataSource) validateMethods(formats strfmt.Registry) error {

	if swag.IsZero(m.Methods) { // not required
		return nil
	}

	for i := 0; i < len(m.Methods); i++ {
		if swag.IsZero(m.Methods[i]) { // not required
			continue
		}

		if m.Methods[i] != nil {
			if err := m.Methods[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("methods" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DataSource) validateScope(formats strfmt.Registry) error {

	if swag.IsZero(m.Scope) { // not required
		return nil
	}

	if m.Scope != nil {
		if err := m.Scope.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scope")
			}
			return err
		}
	}

	return nil
}

var dataSourceTypeSessionAttributeTechnologyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ASP_NET","ASP_NET_CORE","JAVA"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		dataSourceTypeSessionAttributeTechnologyPropEnum = append(dataSourceTypeSessionAttributeTechnologyPropEnum, v)
	}
}

const (

	// DataSourceSessionAttributeTechnologyASPNET captures enum value "ASP_NET"
	DataSourceSessionAttributeTechnologyASPNET string = "ASP_NET"

	// DataSourceSessionAttributeTechnologyASPNETCORE captures enum value "ASP_NET_CORE"
	DataSourceSessionAttributeTechnologyASPNETCORE string = "ASP_NET_CORE"

	// DataSourceSessionAttributeTechnologyJAVA captures enum value "JAVA"
	DataSourceSessionAttributeTechnologyJAVA string = "JAVA"
)

// prop value enum
func (m *DataSource) validateSessionAttributeTechnologyEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, dataSourceTypeSessionAttributeTechnologyPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *DataSource) validateSessionAttributeTechnology(formats strfmt.Registry) error {

	if swag.IsZero(m.SessionAttributeTechnology) { // not required
		return nil
	}

	// value enum
	if err := m.validateSessionAttributeTechnologyEnum("sessionAttributeTechnology", "body", m.SessionAttributeTechnology); err != nil {
		return err
	}

	return nil
}

var dataSourceTypeSourcePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CICS_SDK","CLIENT_IP","CUSTOM_ATTRIBUTE","IIB_NODE","METHOD_PARAM","POST_PARAMETER","QUERY_PARAMETER","REQUEST_HEADER","RESPONSE_HEADER","SESSION_ATTRIBUTE","URI","URI_PATH"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		dataSourceTypeSourcePropEnum = append(dataSourceTypeSourcePropEnum, v)
	}
}

const (

	// DataSourceSourceCICSSDK captures enum value "CICS_SDK"
	DataSourceSourceCICSSDK string = "CICS_SDK"

	// DataSourceSourceCLIENTIP captures enum value "CLIENT_IP"
	DataSourceSourceCLIENTIP string = "CLIENT_IP"

	// DataSourceSourceCUSTOMATTRIBUTE captures enum value "CUSTOM_ATTRIBUTE"
	DataSourceSourceCUSTOMATTRIBUTE string = "CUSTOM_ATTRIBUTE"

	// DataSourceSourceIIBNODE captures enum value "IIB_NODE"
	DataSourceSourceIIBNODE string = "IIB_NODE"

	// DataSourceSourceMETHODPARAM captures enum value "METHOD_PARAM"
	DataSourceSourceMETHODPARAM string = "METHOD_PARAM"

	// DataSourceSourcePOSTPARAMETER captures enum value "POST_PARAMETER"
	DataSourceSourcePOSTPARAMETER string = "POST_PARAMETER"

	// DataSourceSourceQUERYPARAMETER captures enum value "QUERY_PARAMETER"
	DataSourceSourceQUERYPARAMETER string = "QUERY_PARAMETER"

	// DataSourceSourceREQUESTHEADER captures enum value "REQUEST_HEADER"
	DataSourceSourceREQUESTHEADER string = "REQUEST_HEADER"

	// DataSourceSourceRESPONSEHEADER captures enum value "RESPONSE_HEADER"
	DataSourceSourceRESPONSEHEADER string = "RESPONSE_HEADER"

	// DataSourceSourceSESSIONATTRIBUTE captures enum value "SESSION_ATTRIBUTE"
	DataSourceSourceSESSIONATTRIBUTE string = "SESSION_ATTRIBUTE"

	// DataSourceSourceURI captures enum value "URI"
	DataSourceSourceURI string = "URI"

	// DataSourceSourceURIPATH captures enum value "URI_PATH"
	DataSourceSourceURIPATH string = "URI_PATH"
)

// prop value enum
func (m *DataSource) validateSourceEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, dataSourceTypeSourcePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *DataSource) validateSource(formats strfmt.Registry) error {

	if err := validate.Required("source", "body", m.Source); err != nil {
		return err
	}

	// value enum
	if err := m.validateSourceEnum("source", "body", *m.Source); err != nil {
		return err
	}

	return nil
}

var dataSourceTypeTechnologyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DOTNET","JAVA","PHP"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		dataSourceTypeTechnologyPropEnum = append(dataSourceTypeTechnologyPropEnum, v)
	}
}

const (

	// DataSourceTechnologyDOTNET captures enum value "DOTNET"
	DataSourceTechnologyDOTNET string = "DOTNET"

	// DataSourceTechnologyJAVA captures enum value "JAVA"
	DataSourceTechnologyJAVA string = "JAVA"

	// DataSourceTechnologyPHP captures enum value "PHP"
	DataSourceTechnologyPHP string = "PHP"
)

// prop value enum
func (m *DataSource) validateTechnologyEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, dataSourceTypeTechnologyPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *DataSource) validateTechnology(formats strfmt.Registry) error {

	if swag.IsZero(m.Technology) { // not required
		return nil
	}

	// value enum
	if err := m.validateTechnologyEnum("technology", "body", m.Technology); err != nil {
		return err
	}

	return nil
}

func (m *DataSource) validateValueProcessing(formats strfmt.Registry) error {

	if swag.IsZero(m.ValueProcessing) { // not required
		return nil
	}

	if m.ValueProcessing != nil {
		if err := m.ValueProcessing.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("valueProcessing")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DataSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataSource) UnmarshalBinary(b []byte) error {
	var res DataSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
