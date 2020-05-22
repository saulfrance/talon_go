/*
 * Talon.One API
 *
 * The Talon.One API is used to manage applications and campaigns, as well as to integrate with your application. The operations in the _Integration API_ section are used to integrate with our platform, while the other operations are used to manage applications and campaigns.  ### Where is the API?  The API is available at the same hostname as these docs. For example, if you are reading this page at `https://mycompany.talon.one/docs/api/`, the URL for the [updateCustomerProfile][] operation is `https://mycompany.talon.one/v1/customer_profiles/id`  [updateCustomerProfile]: #operation--v1-customer_profiles--integrationId--put
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package talon

import (
	"bytes"
	"encoding/json"
)

// SamlConnectionMetadata struct for SamlConnectionMetadata
type SamlConnectionMetadata struct {
	// ID of the SAML service.
	Name string `json:"name"`
	// Determines if this SAML connection active.
	Enabled   bool    `json:"enabled"`
	AccountId float32 `json:"accountId"`
	// Identity Provider metadata XML document.
	MetadataDocument string `json:"metadataDocument"`
}

// GetName returns the Name field value
func (o *SamlConnectionMetadata) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// SetName sets field value
func (o *SamlConnectionMetadata) SetName(v string) {
	o.Name = v
}

// GetEnabled returns the Enabled field value
func (o *SamlConnectionMetadata) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// SetEnabled sets field value
func (o *SamlConnectionMetadata) SetEnabled(v bool) {
	o.Enabled = v
}

// GetAccountId returns the AccountId field value
func (o *SamlConnectionMetadata) GetAccountId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.AccountId
}

// SetAccountId sets field value
func (o *SamlConnectionMetadata) SetAccountId(v float32) {
	o.AccountId = v
}

// GetMetadataDocument returns the MetadataDocument field value
func (o *SamlConnectionMetadata) GetMetadataDocument() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetadataDocument
}

// SetMetadataDocument sets field value
func (o *SamlConnectionMetadata) SetMetadataDocument(v string) {
	o.MetadataDocument = v
}

type NullableSamlConnectionMetadata struct {
	Value        SamlConnectionMetadata
	ExplicitNull bool
}

func (v NullableSamlConnectionMetadata) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableSamlConnectionMetadata) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}