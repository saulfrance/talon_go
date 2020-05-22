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

// DeductLoyaltyPointsEffectProps The properties specific to the \"deductLoyaltyPoints\" effect. This gets triggered whenever a validated rule contained a condition to only trigger when the given number of loyalty points could be deduced. These points are automatically stored and managed inside Talon.One.
type DeductLoyaltyPointsEffectProps struct {
	// The title of the rule that contained triggered this points deduction
	RuleTitle string `json:"ruleTitle"`
	// The ID of the loyalty program where these points were added
	ProgramId int32 `json:"programId"`
	// The ID of the subledger within the loyalty program where these points were added
	SubLedgerId string `json:"subLedgerId"`
	// The amount of points that were deducted
	Value float32 `json:"value"`
}

// GetRuleTitle returns the RuleTitle field value
func (o *DeductLoyaltyPointsEffectProps) GetRuleTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RuleTitle
}

// SetRuleTitle sets field value
func (o *DeductLoyaltyPointsEffectProps) SetRuleTitle(v string) {
	o.RuleTitle = v
}

// GetProgramId returns the ProgramId field value
func (o *DeductLoyaltyPointsEffectProps) GetProgramId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProgramId
}

// SetProgramId sets field value
func (o *DeductLoyaltyPointsEffectProps) SetProgramId(v int32) {
	o.ProgramId = v
}

// GetSubLedgerId returns the SubLedgerId field value
func (o *DeductLoyaltyPointsEffectProps) GetSubLedgerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubLedgerId
}

// SetSubLedgerId sets field value
func (o *DeductLoyaltyPointsEffectProps) SetSubLedgerId(v string) {
	o.SubLedgerId = v
}

// GetValue returns the Value field value
func (o *DeductLoyaltyPointsEffectProps) GetValue() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Value
}

// SetValue sets field value
func (o *DeductLoyaltyPointsEffectProps) SetValue(v float32) {
	o.Value = v
}

type NullableDeductLoyaltyPointsEffectProps struct {
	Value        DeductLoyaltyPointsEffectProps
	ExplicitNull bool
}

func (v NullableDeductLoyaltyPointsEffectProps) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableDeductLoyaltyPointsEffectProps) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}