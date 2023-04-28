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

// IntegrationCustomerSessionResponse struct for IntegrationCustomerSessionResponse
type IntegrationCustomerSessionResponse struct {
	CustomerSession *CustomerSessionV2 `json:"customerSession,omitempty"`
	Effects         *[]Effect          `json:"effects,omitempty"`
}

// GetCustomerSession returns the CustomerSession field value if set, zero value otherwise.
func (o *IntegrationCustomerSessionResponse) GetCustomerSession() CustomerSessionV2 {
	if o == nil || o.CustomerSession == nil {
		var ret CustomerSessionV2
		return ret
	}
	return *o.CustomerSession
}

// GetCustomerSessionOk returns a tuple with the CustomerSession field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationCustomerSessionResponse) GetCustomerSessionOk() (CustomerSessionV2, bool) {
	if o == nil || o.CustomerSession == nil {
		var ret CustomerSessionV2
		return ret, false
	}
	return *o.CustomerSession, true
}

// HasCustomerSession returns a boolean if a field has been set.
func (o *IntegrationCustomerSessionResponse) HasCustomerSession() bool {
	if o != nil && o.CustomerSession != nil {
		return true
	}

	return false
}

// SetCustomerSession gets a reference to the given CustomerSessionV2 and assigns it to the CustomerSession field.
func (o *IntegrationCustomerSessionResponse) SetCustomerSession(v CustomerSessionV2) {
	o.CustomerSession = &v
}

// GetEffects returns the Effects field value if set, zero value otherwise.
func (o *IntegrationCustomerSessionResponse) GetEffects() []Effect {
	if o == nil || o.Effects == nil {
		var ret []Effect
		return ret
	}
	return *o.Effects
}

// GetEffectsOk returns a tuple with the Effects field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationCustomerSessionResponse) GetEffectsOk() ([]Effect, bool) {
	if o == nil || o.Effects == nil {
		var ret []Effect
		return ret, false
	}
	return *o.Effects, true
}

// HasEffects returns a boolean if a field has been set.
func (o *IntegrationCustomerSessionResponse) HasEffects() bool {
	if o != nil && o.Effects != nil {
		return true
	}

	return false
}

// SetEffects gets a reference to the given []Effect and assigns it to the Effects field.
func (o *IntegrationCustomerSessionResponse) SetEffects(v []Effect) {
	o.Effects = &v
}

type NullableIntegrationCustomerSessionResponse struct {
	Value        IntegrationCustomerSessionResponse
	ExplicitNull bool
}

func (v NullableIntegrationCustomerSessionResponse) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableIntegrationCustomerSessionResponse) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}