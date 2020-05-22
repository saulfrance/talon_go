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

// AddLoyaltyPointsEffectProps The properties specific to the \"addLoyaltyPoints\" effect. This gets triggered whenever a validated rule contained an \"add loyalty\" effect. These points are automatically stored and managed inside Talon.One.
type AddLoyaltyPointsEffectProps struct {
	// The name/description of this loyalty point addition
	Name string `json:"name"`
	// The ID of the loyalty program where these points were added
	ProgramId int32 `json:"programId"`
	// The ID of the subledger within the loyalty program where these points were added
	SubLedgerId string `json:"subLedgerId"`
	// The amount of points that were added
	Value float32 `json:"value"`
	// The user for whom these points were added
	RecipientIntegrationId string `json:"recipientIntegrationId"`
	// The amount of time (in days) these points are valid
	ExpiryCondition string `json:"expiryCondition"`
}

// GetName returns the Name field value
func (o *AddLoyaltyPointsEffectProps) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// SetName sets field value
func (o *AddLoyaltyPointsEffectProps) SetName(v string) {
	o.Name = v
}

// GetProgramId returns the ProgramId field value
func (o *AddLoyaltyPointsEffectProps) GetProgramId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProgramId
}

// SetProgramId sets field value
func (o *AddLoyaltyPointsEffectProps) SetProgramId(v int32) {
	o.ProgramId = v
}

// GetSubLedgerId returns the SubLedgerId field value
func (o *AddLoyaltyPointsEffectProps) GetSubLedgerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubLedgerId
}

// SetSubLedgerId sets field value
func (o *AddLoyaltyPointsEffectProps) SetSubLedgerId(v string) {
	o.SubLedgerId = v
}

// GetValue returns the Value field value
func (o *AddLoyaltyPointsEffectProps) GetValue() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Value
}

// SetValue sets field value
func (o *AddLoyaltyPointsEffectProps) SetValue(v float32) {
	o.Value = v
}

// GetRecipientIntegrationId returns the RecipientIntegrationId field value
func (o *AddLoyaltyPointsEffectProps) GetRecipientIntegrationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecipientIntegrationId
}

// SetRecipientIntegrationId sets field value
func (o *AddLoyaltyPointsEffectProps) SetRecipientIntegrationId(v string) {
	o.RecipientIntegrationId = v
}

// GetExpiryCondition returns the ExpiryCondition field value
func (o *AddLoyaltyPointsEffectProps) GetExpiryCondition() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExpiryCondition
}

// SetExpiryCondition sets field value
func (o *AddLoyaltyPointsEffectProps) SetExpiryCondition(v string) {
	o.ExpiryCondition = v
}

type NullableAddLoyaltyPointsEffectProps struct {
	Value        AddLoyaltyPointsEffectProps
	ExplicitNull bool
}

func (v NullableAddLoyaltyPointsEffectProps) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableAddLoyaltyPointsEffectProps) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}