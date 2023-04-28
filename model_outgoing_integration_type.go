/*
 * Talon.One API
 *
 * Use the Talon.One API to integrate with your application and to manage applications and campaigns:  - Use the operations in the [Integration API section](#integration-api) are used to integrate with our platform - Use the operation in the [Management API section](#management-api) to manage applications and campaigns.  ## Determining the base URL of the endpoints  The API is available at the same hostname as your Campaign Manager deployment. For example, if you access the Campaign Manager at `https://yourbaseurl.talon.one/`, the URL for the [updateCustomerSessionV2](https://docs.talon.one/integration-api#operation/updateCustomerSessionV2) endpoint is `https://yourbaseurl.talon.one/v2/customer_sessions/{Id}`
 *
 * API version:
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package talon

import (
	"bytes"
	"encoding/json"
)

// OutgoingIntegrationType struct for OutgoingIntegrationType
type OutgoingIntegrationType struct {
	// Unique ID for this entity.
	Id int32 `json:"id"`
	// Name of the outgoing integration.
	Name string `json:"name"`
	// Description of the outgoing integration.
	Description *string `json:"description,omitempty"`
	// Category of the outgoing integration.
	Category *string `json:"category,omitempty"`
	// Http link to the outgoing integration's documentation.
	DocumentationLink *string `json:"documentationLink,omitempty"`
}

// GetId returns the Id field value
func (o *OutgoingIntegrationType) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// SetId sets field value
func (o *OutgoingIntegrationType) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *OutgoingIntegrationType) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// SetName sets field value
func (o *OutgoingIntegrationType) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OutgoingIntegrationType) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *OutgoingIntegrationType) GetDescriptionOk() (string, bool) {
	if o == nil || o.Description == nil {
		var ret string
		return ret, false
	}
	return *o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OutgoingIntegrationType) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OutgoingIntegrationType) SetDescription(v string) {
	o.Description = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *OutgoingIntegrationType) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *OutgoingIntegrationType) GetCategoryOk() (string, bool) {
	if o == nil || o.Category == nil {
		var ret string
		return ret, false
	}
	return *o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *OutgoingIntegrationType) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *OutgoingIntegrationType) SetCategory(v string) {
	o.Category = &v
}

// GetDocumentationLink returns the DocumentationLink field value if set, zero value otherwise.
func (o *OutgoingIntegrationType) GetDocumentationLink() string {
	if o == nil || o.DocumentationLink == nil {
		var ret string
		return ret
	}
	return *o.DocumentationLink
}

// GetDocumentationLinkOk returns a tuple with the DocumentationLink field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *OutgoingIntegrationType) GetDocumentationLinkOk() (string, bool) {
	if o == nil || o.DocumentationLink == nil {
		var ret string
		return ret, false
	}
	return *o.DocumentationLink, true
}

// HasDocumentationLink returns a boolean if a field has been set.
func (o *OutgoingIntegrationType) HasDocumentationLink() bool {
	if o != nil && o.DocumentationLink != nil {
		return true
	}

	return false
}

// SetDocumentationLink gets a reference to the given string and assigns it to the DocumentationLink field.
func (o *OutgoingIntegrationType) SetDocumentationLink(v string) {
	o.DocumentationLink = &v
}

type NullableOutgoingIntegrationType struct {
	Value        OutgoingIntegrationType
	ExplicitNull bool
}

func (v NullableOutgoingIntegrationType) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableOutgoingIntegrationType) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}