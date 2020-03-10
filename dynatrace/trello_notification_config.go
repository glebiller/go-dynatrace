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

// TrelloNotificationConfig trello notification config
// swagger:model TrelloNotificationConfig
type TrelloNotificationConfig struct {
	activeField *bool

	alertingProfileField *strfmt.UUID

	idField strfmt.UUID

	nameField *string

	// The application key for the Trello account.
	// Required: true
	// Max Length: 1000
	// Min Length: 1
	ApplicationKey *string `json:"applicationKey"`

	// The application token for the Trello account.
	// Max Length: 1000
	// Min Length: 1
	AuthorizationToken string `json:"authorizationToken,omitempty"`

	// The Trello board to which the card should be assigned.
	// Required: true
	// Min Length: 1
	BoardID *string `json:"boardId"`

	// The description of the Trello card.
	//
	//  You can use same placeholders as in card text.
	// Required: true
	// Min Length: 1
	Description *string `json:"description"`

	// The Trello list to which the card should be assigned.
	// Required: true
	// Min Length: 1
	ListID *string `json:"listId"`

	// The Trello list to which the card of the resolved problem should be assigned.
	// Required: true
	// Min Length: 1
	ResolvedListID *string `json:"resolvedListId"`

	// The text of the generated Trello card.
	//
	//  You can use the following placeholders:
	// * `{ImpactedEntity}`: The entity impacted by the problem or *X* impacted entities.
	// * `{PID}`: The ID of the reported problem.
	// * `{ProblemDetailsMarkdown}`: All problem event details, including root cause, as a [Markdown-formatted](https://www.markdownguide.org/cheat-sheet/) string.
	// * `{ProblemID}`: The display number of the reported problem.
	// * `{ProblemImpact}`: The [impact level](https://www.dynatrace.com/support/help/shortlink/impact-analysis) of the problem. Possible values are `APPLICATION`, `SERVICE`, and `INFRASTRUCTURE`.
	// * `{ProblemSeverity}`: The [severity level](https://www.dynatrace.com/support/help/shortlink/event-types) of the problem. Possible values are `AVAILABILITY`, `ERROR`, `PERFORMANCE`, `RESOURCE_CONTENTION`, and `CUSTOM_ALERT`.
	// * `{ProblemTitle}`: A short description of the problem.
	// * `{ProblemURL}`: The URL of the problem within Dynatrace.
	// * `{State}`: The state of the problem. Possible values are `OPEN` and `RESOLVED`. If the problem has been merged into another problem, it has the `MERGED` value.
	// * `{Tags}`: The list of tags that are defined for all impacted entities, separated by commas.
	//
	// Required: true
	// Min Length: 1
	Text *string `json:"text"`
}

// Active gets the active of this subtype
func (m *TrelloNotificationConfig) Active() *bool {
	return m.activeField
}

// SetActive sets the active of this subtype
func (m *TrelloNotificationConfig) SetActive(val *bool) {
	m.activeField = val
}

// AlertingProfile gets the alerting profile of this subtype
func (m *TrelloNotificationConfig) AlertingProfile() *strfmt.UUID {
	return m.alertingProfileField
}

// SetAlertingProfile sets the alerting profile of this subtype
func (m *TrelloNotificationConfig) SetAlertingProfile(val *strfmt.UUID) {
	m.alertingProfileField = val
}

// ID gets the id of this subtype
func (m *TrelloNotificationConfig) ID() strfmt.UUID {
	return m.idField
}

// SetID sets the id of this subtype
func (m *TrelloNotificationConfig) SetID(val strfmt.UUID) {
	m.idField = val
}

// Name gets the name of this subtype
func (m *TrelloNotificationConfig) Name() *string {
	return m.nameField
}

// SetName sets the name of this subtype
func (m *TrelloNotificationConfig) SetName(val *string) {
	m.nameField = val
}

// Type gets the type of this subtype
func (m *TrelloNotificationConfig) Type() string {
	return "TrelloNotificationConfig"
}

// SetType sets the type of this subtype
func (m *TrelloNotificationConfig) SetType(val string) {

}

// ApplicationKey gets the application key of this subtype

// AuthorizationToken gets the authorization token of this subtype

// BoardID gets the board Id of this subtype

// Description gets the description of this subtype

// ListID gets the list Id of this subtype

// ResolvedListID gets the resolved list Id of this subtype

// Text gets the text of this subtype

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *TrelloNotificationConfig) UnmarshalJSON(raw []byte) error {
	var data struct {

		// The application key for the Trello account.
		// Required: true
		// Max Length: 1000
		// Min Length: 1
		ApplicationKey *string `json:"applicationKey"`

		// The application token for the Trello account.
		// Max Length: 1000
		// Min Length: 1
		AuthorizationToken string `json:"authorizationToken,omitempty"`

		// The Trello board to which the card should be assigned.
		// Required: true
		// Min Length: 1
		BoardID *string `json:"boardId"`

		// The description of the Trello card.
		//
		//  You can use same placeholders as in card text.
		// Required: true
		// Min Length: 1
		Description *string `json:"description"`

		// The Trello list to which the card should be assigned.
		// Required: true
		// Min Length: 1
		ListID *string `json:"listId"`

		// The Trello list to which the card of the resolved problem should be assigned.
		// Required: true
		// Min Length: 1
		ResolvedListID *string `json:"resolvedListId"`

		// The text of the generated Trello card.
		//
		//  You can use the following placeholders:
		// * `{ImpactedEntity}`: The entity impacted by the problem or *X* impacted entities.
		// * `{PID}`: The ID of the reported problem.
		// * `{ProblemDetailsMarkdown}`: All problem event details, including root cause, as a [Markdown-formatted](https://www.markdownguide.org/cheat-sheet/) string.
		// * `{ProblemID}`: The display number of the reported problem.
		// * `{ProblemImpact}`: The [impact level](https://www.dynatrace.com/support/help/shortlink/impact-analysis) of the problem. Possible values are `APPLICATION`, `SERVICE`, and `INFRASTRUCTURE`.
		// * `{ProblemSeverity}`: The [severity level](https://www.dynatrace.com/support/help/shortlink/event-types) of the problem. Possible values are `AVAILABILITY`, `ERROR`, `PERFORMANCE`, `RESOURCE_CONTENTION`, and `CUSTOM_ALERT`.
		// * `{ProblemTitle}`: A short description of the problem.
		// * `{ProblemURL}`: The URL of the problem within Dynatrace.
		// * `{State}`: The state of the problem. Possible values are `OPEN` and `RESOLVED`. If the problem has been merged into another problem, it has the `MERGED` value.
		// * `{Tags}`: The list of tags that are defined for all impacted entities, separated by commas.
		//
		// Required: true
		// Min Length: 1
		Text *string `json:"text"`
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

	var result TrelloNotificationConfig

	result.activeField = base.Active

	result.alertingProfileField = base.AlertingProfile

	result.idField = base.ID

	result.nameField = base.Name

	if base.Type != result.Type() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid type value: %q", base.Type)
	}

	result.ApplicationKey = data.ApplicationKey

	result.AuthorizationToken = data.AuthorizationToken

	result.BoardID = data.BoardID

	result.Description = data.Description

	result.ListID = data.ListID

	result.ResolvedListID = data.ResolvedListID

	result.Text = data.Text

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m TrelloNotificationConfig) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// The application key for the Trello account.
		// Required: true
		// Max Length: 1000
		// Min Length: 1
		ApplicationKey *string `json:"applicationKey"`

		// The application token for the Trello account.
		// Max Length: 1000
		// Min Length: 1
		AuthorizationToken string `json:"authorizationToken,omitempty"`

		// The Trello board to which the card should be assigned.
		// Required: true
		// Min Length: 1
		BoardID *string `json:"boardId"`

		// The description of the Trello card.
		//
		//  You can use same placeholders as in card text.
		// Required: true
		// Min Length: 1
		Description *string `json:"description"`

		// The Trello list to which the card should be assigned.
		// Required: true
		// Min Length: 1
		ListID *string `json:"listId"`

		// The Trello list to which the card of the resolved problem should be assigned.
		// Required: true
		// Min Length: 1
		ResolvedListID *string `json:"resolvedListId"`

		// The text of the generated Trello card.
		//
		//  You can use the following placeholders:
		// * `{ImpactedEntity}`: The entity impacted by the problem or *X* impacted entities.
		// * `{PID}`: The ID of the reported problem.
		// * `{ProblemDetailsMarkdown}`: All problem event details, including root cause, as a [Markdown-formatted](https://www.markdownguide.org/cheat-sheet/) string.
		// * `{ProblemID}`: The display number of the reported problem.
		// * `{ProblemImpact}`: The [impact level](https://www.dynatrace.com/support/help/shortlink/impact-analysis) of the problem. Possible values are `APPLICATION`, `SERVICE`, and `INFRASTRUCTURE`.
		// * `{ProblemSeverity}`: The [severity level](https://www.dynatrace.com/support/help/shortlink/event-types) of the problem. Possible values are `AVAILABILITY`, `ERROR`, `PERFORMANCE`, `RESOURCE_CONTENTION`, and `CUSTOM_ALERT`.
		// * `{ProblemTitle}`: A short description of the problem.
		// * `{ProblemURL}`: The URL of the problem within Dynatrace.
		// * `{State}`: The state of the problem. Possible values are `OPEN` and `RESOLVED`. If the problem has been merged into another problem, it has the `MERGED` value.
		// * `{Tags}`: The list of tags that are defined for all impacted entities, separated by commas.
		//
		// Required: true
		// Min Length: 1
		Text *string `json:"text"`
	}{

		ApplicationKey: m.ApplicationKey,

		AuthorizationToken: m.AuthorizationToken,

		BoardID: m.BoardID,

		Description: m.Description,

		ListID: m.ListID,

		ResolvedListID: m.ResolvedListID,

		Text: m.Text,
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

// Validate validates this trello notification config
func (m *TrelloNotificationConfig) Validate(formats strfmt.Registry) error {
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

	if err := m.validateApplicationKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthorizationToken(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBoardID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateListID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResolvedListID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateText(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TrelloNotificationConfig) validateActive(formats strfmt.Registry) error {

	if err := validate.Required("active", "body", m.Active()); err != nil {
		return err
	}

	return nil
}

func (m *TrelloNotificationConfig) validateAlertingProfile(formats strfmt.Registry) error {

	if err := validate.Required("alertingProfile", "body", m.AlertingProfile()); err != nil {
		return err
	}

	if err := validate.FormatOf("alertingProfile", "body", "uuid", m.AlertingProfile().String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TrelloNotificationConfig) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID()) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID().String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TrelloNotificationConfig) validateName(formats strfmt.Registry) error {

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

func (m *TrelloNotificationConfig) validateApplicationKey(formats strfmt.Registry) error {

	if err := validate.Required("applicationKey", "body", m.ApplicationKey); err != nil {
		return err
	}

	if err := validate.MinLength("applicationKey", "body", string(*m.ApplicationKey), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("applicationKey", "body", string(*m.ApplicationKey), 1000); err != nil {
		return err
	}

	return nil
}

func (m *TrelloNotificationConfig) validateAuthorizationToken(formats strfmt.Registry) error {

	if swag.IsZero(m.AuthorizationToken) { // not required
		return nil
	}

	if err := validate.MinLength("authorizationToken", "body", string(m.AuthorizationToken), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("authorizationToken", "body", string(m.AuthorizationToken), 1000); err != nil {
		return err
	}

	return nil
}

func (m *TrelloNotificationConfig) validateBoardID(formats strfmt.Registry) error {

	if err := validate.Required("boardId", "body", m.BoardID); err != nil {
		return err
	}

	if err := validate.MinLength("boardId", "body", string(*m.BoardID), 1); err != nil {
		return err
	}

	return nil
}

func (m *TrelloNotificationConfig) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	if err := validate.MinLength("description", "body", string(*m.Description), 1); err != nil {
		return err
	}

	return nil
}

func (m *TrelloNotificationConfig) validateListID(formats strfmt.Registry) error {

	if err := validate.Required("listId", "body", m.ListID); err != nil {
		return err
	}

	if err := validate.MinLength("listId", "body", string(*m.ListID), 1); err != nil {
		return err
	}

	return nil
}

func (m *TrelloNotificationConfig) validateResolvedListID(formats strfmt.Registry) error {

	if err := validate.Required("resolvedListId", "body", m.ResolvedListID); err != nil {
		return err
	}

	if err := validate.MinLength("resolvedListId", "body", string(*m.ResolvedListID), 1); err != nil {
		return err
	}

	return nil
}

func (m *TrelloNotificationConfig) validateText(formats strfmt.Registry) error {

	if err := validate.Required("text", "body", m.Text); err != nil {
		return err
	}

	if err := validate.MinLength("text", "body", string(*m.Text), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TrelloNotificationConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TrelloNotificationConfig) UnmarshalBinary(b []byte) error {
	var res TrelloNotificationConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
