// Code generated by go-swagger; DO NOT EDIT.

package dynatrace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// JiraNotificationConfig jira notification config
// swagger:model JiraNotificationConfig
type JiraNotificationConfig struct {
	activeField *bool

	alertingProfileField *strfmt.UUID

	idField strfmt.UUID

	nameField *string

	// The description of the Jira issue to be created by this notification.
	//
	//  You can use same placeholders as in issue summary.
	// Required: true
	// Min Length: 1
	Description *string `json:"description"`

	// The type of the Jira issue to be created by this notification.
	// Required: true
	// Max Length: 1000
	// Min Length: 1
	IssueType *string `json:"issueType"`

	// The password for the Jira profile.
	// Max Length: 1000
	// Min Length: 1
	Password string `json:"password,omitempty"`

	// The project key of the Jira issue to be created by this notification.
	// Required: true
	// Max Length: 1000
	// Min Length: 1
	ProjectKey *string `json:"projectKey"`

	// The summary of the Jira issue to be created by this notification.
	//
	//  You can use the following placeholders:
	// * `{ImpactedEntity}`: The entity impacted by the problem or *X* impacted entities.
	// * `{PID}`: The ID of the reported problem.
	// * `{ProblemDetailsText}`: All problem event details, including root cause, as a text-formatted string.
	// * `{ProblemID}`: The display number of the reported problem.
	// * `{ProblemImpact}`: The [impact level](https://www.dynatrace.com/support/help/shortlink/impact-analysis) of the problem. Possible values are `APPLICATION`, `SERVICE`, and `INFRASTRUCTURE`.
	// * `{ProblemSeverity}`: The [severity level](https://www.dynatrace.com/support/help/shortlink/event-types) of the problem. Possible values are `AVAILABILITY`, `ERROR`, `PERFORMANCE`, `RESOURCE_CONTENTION`, and `CUSTOM_ALERT`.
	// * `{ProblemTitle}`: A short description of the problem.
	// * `{ProblemURL}`: The URL of the problem within Dynatrace.
	// * `{State}`: The state of the problem. Possible values are `OPEN` and `RESOLVED`. If the problem has been merged into another problem, it has the `MERGED` value.
	// * `{Tags}`: The list of tags that are defined for all impacted entities, separated by commas.
	//
	// Required: true
	// Max Length: 1000
	// Min Length: 1
	Summary *string `json:"summary"`

	// The URL of the Jira API endpoint.
	// Required: true
	// Min Length: 1
	URL *string `json:"url"`

	// The username of the Jira profile.
	// Required: true
	// Max Length: 1000
	// Min Length: 1
	Username *string `json:"username"`
}

// Active gets the active of this subtype
func (m *JiraNotificationConfig) Active() *bool {
	return m.activeField
}

// SetActive sets the active of this subtype
func (m *JiraNotificationConfig) SetActive(val *bool) {
	m.activeField = val
}

// AlertingProfile gets the alerting profile of this subtype
func (m *JiraNotificationConfig) AlertingProfile() *strfmt.UUID {
	return m.alertingProfileField
}

// SetAlertingProfile sets the alerting profile of this subtype
func (m *JiraNotificationConfig) SetAlertingProfile(val *strfmt.UUID) {
	m.alertingProfileField = val
}

// ID gets the id of this subtype
func (m *JiraNotificationConfig) ID() strfmt.UUID {
	return m.idField
}

// SetID sets the id of this subtype
func (m *JiraNotificationConfig) SetID(val strfmt.UUID) {
	m.idField = val
}

// Name gets the name of this subtype
func (m *JiraNotificationConfig) Name() *string {
	return m.nameField
}

// SetName sets the name of this subtype
func (m *JiraNotificationConfig) SetName(val *string) {
	m.nameField = val
}

// Type gets the type of this subtype
func (m *JiraNotificationConfig) Type() string {
	return "JiraNotificationConfig"
}

// SetType sets the type of this subtype
func (m *JiraNotificationConfig) SetType(val string) {

}

// Description gets the description of this subtype

// IssueType gets the issue type of this subtype

// Password gets the password of this subtype

// ProjectKey gets the project key of this subtype

// Summary gets the summary of this subtype

// URL gets the url of this subtype

// Username gets the username of this subtype

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *JiraNotificationConfig) UnmarshalJSON(raw []byte) error {
	var data struct {

		// The description of the Jira issue to be created by this notification.
		//
		//  You can use same placeholders as in issue summary.
		// Required: true
		// Min Length: 1
		Description *string `json:"description"`

		// The type of the Jira issue to be created by this notification.
		// Required: true
		// Max Length: 1000
		// Min Length: 1
		IssueType *string `json:"issueType"`

		// The password for the Jira profile.
		// Max Length: 1000
		// Min Length: 1
		Password string `json:"password,omitempty"`

		// The project key of the Jira issue to be created by this notification.
		// Required: true
		// Max Length: 1000
		// Min Length: 1
		ProjectKey *string `json:"projectKey"`

		// The summary of the Jira issue to be created by this notification.
		//
		//  You can use the following placeholders:
		// * `{ImpactedEntity}`: The entity impacted by the problem or *X* impacted entities.
		// * `{PID}`: The ID of the reported problem.
		// * `{ProblemDetailsText}`: All problem event details, including root cause, as a text-formatted string.
		// * `{ProblemID}`: The display number of the reported problem.
		// * `{ProblemImpact}`: The [impact level](https://www.dynatrace.com/support/help/shortlink/impact-analysis) of the problem. Possible values are `APPLICATION`, `SERVICE`, and `INFRASTRUCTURE`.
		// * `{ProblemSeverity}`: The [severity level](https://www.dynatrace.com/support/help/shortlink/event-types) of the problem. Possible values are `AVAILABILITY`, `ERROR`, `PERFORMANCE`, `RESOURCE_CONTENTION`, and `CUSTOM_ALERT`.
		// * `{ProblemTitle}`: A short description of the problem.
		// * `{ProblemURL}`: The URL of the problem within Dynatrace.
		// * `{State}`: The state of the problem. Possible values are `OPEN` and `RESOLVED`. If the problem has been merged into another problem, it has the `MERGED` value.
		// * `{Tags}`: The list of tags that are defined for all impacted entities, separated by commas.
		//
		// Required: true
		// Max Length: 1000
		// Min Length: 1
		Summary *string `json:"summary"`

		// The URL of the Jira API endpoint.
		// Required: true
		// Min Length: 1
		URL *string `json:"url"`

		// The username of the Jira profile.
		// Required: true
		// Max Length: 1000
		// Min Length: 1
		Username *string `json:"username"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		Active *bool `json:"active"`

		AlertingProfile *strfmt.UUID `json:"alertingProfile"`

		ID strfmt.UUID `json:"id,omitempty"`

		Name *string `json:"name"`

		Type string `json:"type"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result JiraNotificationConfig

	result.activeField = base.Active

	result.alertingProfileField = base.AlertingProfile

	result.idField = base.ID

	result.nameField = base.Name

	if base.Type != result.Type() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid type value: %q", base.Type)
	}

	result.Description = data.Description

	result.IssueType = data.IssueType

	result.Password = data.Password

	result.ProjectKey = data.ProjectKey

	result.Summary = data.Summary

	result.URL = data.URL

	result.Username = data.Username

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m JiraNotificationConfig) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// The description of the Jira issue to be created by this notification.
		//
		//  You can use same placeholders as in issue summary.
		// Required: true
		// Min Length: 1
		Description *string `json:"description"`

		// The type of the Jira issue to be created by this notification.
		// Required: true
		// Max Length: 1000
		// Min Length: 1
		IssueType *string `json:"issueType"`

		// The password for the Jira profile.
		// Max Length: 1000
		// Min Length: 1
		Password string `json:"password,omitempty"`

		// The project key of the Jira issue to be created by this notification.
		// Required: true
		// Max Length: 1000
		// Min Length: 1
		ProjectKey *string `json:"projectKey"`

		// The summary of the Jira issue to be created by this notification.
		//
		//  You can use the following placeholders:
		// * `{ImpactedEntity}`: The entity impacted by the problem or *X* impacted entities.
		// * `{PID}`: The ID of the reported problem.
		// * `{ProblemDetailsText}`: All problem event details, including root cause, as a text-formatted string.
		// * `{ProblemID}`: The display number of the reported problem.
		// * `{ProblemImpact}`: The [impact level](https://www.dynatrace.com/support/help/shortlink/impact-analysis) of the problem. Possible values are `APPLICATION`, `SERVICE`, and `INFRASTRUCTURE`.
		// * `{ProblemSeverity}`: The [severity level](https://www.dynatrace.com/support/help/shortlink/event-types) of the problem. Possible values are `AVAILABILITY`, `ERROR`, `PERFORMANCE`, `RESOURCE_CONTENTION`, and `CUSTOM_ALERT`.
		// * `{ProblemTitle}`: A short description of the problem.
		// * `{ProblemURL}`: The URL of the problem within Dynatrace.
		// * `{State}`: The state of the problem. Possible values are `OPEN` and `RESOLVED`. If the problem has been merged into another problem, it has the `MERGED` value.
		// * `{Tags}`: The list of tags that are defined for all impacted entities, separated by commas.
		//
		// Required: true
		// Max Length: 1000
		// Min Length: 1
		Summary *string `json:"summary"`

		// The URL of the Jira API endpoint.
		// Required: true
		// Min Length: 1
		URL *string `json:"url"`

		// The username of the Jira profile.
		// Required: true
		// Max Length: 1000
		// Min Length: 1
		Username *string `json:"username"`
	}{

		Description: m.Description,

		IssueType: m.IssueType,

		Password: m.Password,

		ProjectKey: m.ProjectKey,

		Summary: m.Summary,

		URL: m.URL,

		Username: m.Username,
	},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Active *bool `json:"active"`

		AlertingProfile *strfmt.UUID `json:"alertingProfile"`

		ID strfmt.UUID `json:"id,omitempty"`

		Name *string `json:"name"`

		Type string `json:"type"`
	}{

		Active: m.Active(),

		AlertingProfile: m.AlertingProfile(),

		ID: m.ID(),

		Name: m.Name(),

		Type: m.Type(),
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this jira notification config
func (m *JiraNotificationConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActive(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAlertingProfile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIssueType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSummary(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JiraNotificationConfig) validateActive(formats strfmt.Registry) error {

	if err := validate.Required("active", "body", m.Active()); err != nil {
		return err
	}

	return nil
}

func (m *JiraNotificationConfig) validateAlertingProfile(formats strfmt.Registry) error {

	if err := validate.Required("alertingProfile", "body", m.AlertingProfile()); err != nil {
		return err
	}

	if err := validate.FormatOf("alertingProfile", "body", "uuid", m.AlertingProfile().String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *JiraNotificationConfig) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID()) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID().String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *JiraNotificationConfig) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name()); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name()), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name()), 100); err != nil {
		return err
	}

	return nil
}

func (m *JiraNotificationConfig) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	if err := validate.MinLength("description", "body", string(*m.Description), 1); err != nil {
		return err
	}

	return nil
}

func (m *JiraNotificationConfig) validateIssueType(formats strfmt.Registry) error {

	if err := validate.Required("issueType", "body", m.IssueType); err != nil {
		return err
	}

	if err := validate.MinLength("issueType", "body", string(*m.IssueType), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("issueType", "body", string(*m.IssueType), 1000); err != nil {
		return err
	}

	return nil
}

func (m *JiraNotificationConfig) validatePassword(formats strfmt.Registry) error {

	if swag.IsZero(m.Password) { // not required
		return nil
	}

	if err := validate.MinLength("password", "body", string(m.Password), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("password", "body", string(m.Password), 1000); err != nil {
		return err
	}

	return nil
}

func (m *JiraNotificationConfig) validateProjectKey(formats strfmt.Registry) error {

	if err := validate.Required("projectKey", "body", m.ProjectKey); err != nil {
		return err
	}

	if err := validate.MinLength("projectKey", "body", string(*m.ProjectKey), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("projectKey", "body", string(*m.ProjectKey), 1000); err != nil {
		return err
	}

	return nil
}

func (m *JiraNotificationConfig) validateSummary(formats strfmt.Registry) error {

	if err := validate.Required("summary", "body", m.Summary); err != nil {
		return err
	}

	if err := validate.MinLength("summary", "body", string(*m.Summary), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("summary", "body", string(*m.Summary), 1000); err != nil {
		return err
	}

	return nil
}

func (m *JiraNotificationConfig) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("url", "body", m.URL); err != nil {
		return err
	}

	if err := validate.MinLength("url", "body", string(*m.URL), 1); err != nil {
		return err
	}

	return nil
}

func (m *JiraNotificationConfig) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.Username); err != nil {
		return err
	}

	if err := validate.MinLength("username", "body", string(*m.Username), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("username", "body", string(*m.Username), 1000); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *JiraNotificationConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JiraNotificationConfig) UnmarshalBinary(b []byte) error {
	var res JiraNotificationConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
