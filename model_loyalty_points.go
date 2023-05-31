/*
 * Talon.One API
 *
 * Use the Talon.One API to integrate with your application and to manage applications and campaigns:  - Use the operations in the [Integration API section](#integration-api) are used to integrate with our platform - Use the operation in the [Management API section](#management-api) to manage applications and campaigns.  ## Determining the base URL of the endpoints  The API is available at the same hostname as your Campaign Manager deployment. For example, if you are reading this page at `https://mycompany.talon.one/docs/api/`, the URL for the [updateCustomerSession](https://docs.talon.one/integration-api/#operation/updateCustomerSessionV2) endpoint is `https://mycompany.talon.one/v2/customer_sessions/{Id}`
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package talon

import (
	"bytes"
	"encoding/json"
)

// LoyaltyPoints Points to add or deduct.
type LoyaltyPoints struct {
	// Amount of loyalty points.
	Points float32 `json:"points"`
	// Allows to specify a name/reason for the point addition or deduction.
	Name *string `json:"name,omitempty"`
	// Indicates the duration after which the added loyalty points should expire. The format is a number followed by one letter indicating the time unit, like '1h' or '40m' (defined by Go time package).
	ValidityDuration *string `json:"validityDuration,omitempty"`
	// Indicates the amount of time before the points are considered valid. The format is a number followed by one letter indicating the time unit, like '1h' or '40m' (defined by Go time package).
	PendingDuration *string `json:"pendingDuration,omitempty"`
	// This specifies if we are adding loyalty points to the main ledger or a subledger.
	SubLedgerID *string `json:"subLedgerID,omitempty"`
}

// GetPoints returns the Points field value
func (o *LoyaltyPoints) GetPoints() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Points
}

// SetPoints sets field value
func (o *LoyaltyPoints) SetPoints(v float32) {
	o.Points = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LoyaltyPoints) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LoyaltyPoints) GetNameOk() (string, bool) {
	if o == nil || o.Name == nil {
		var ret string
		return ret, false
	}
	return *o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LoyaltyPoints) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LoyaltyPoints) SetName(v string) {
	o.Name = &v
}

// GetValidityDuration returns the ValidityDuration field value if set, zero value otherwise.
func (o *LoyaltyPoints) GetValidityDuration() string {
	if o == nil || o.ValidityDuration == nil {
		var ret string
		return ret
	}
	return *o.ValidityDuration
}

// GetValidityDurationOk returns a tuple with the ValidityDuration field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LoyaltyPoints) GetValidityDurationOk() (string, bool) {
	if o == nil || o.ValidityDuration == nil {
		var ret string
		return ret, false
	}
	return *o.ValidityDuration, true
}

// HasValidityDuration returns a boolean if a field has been set.
func (o *LoyaltyPoints) HasValidityDuration() bool {
	if o != nil && o.ValidityDuration != nil {
		return true
	}

	return false
}

// SetValidityDuration gets a reference to the given string and assigns it to the ValidityDuration field.
func (o *LoyaltyPoints) SetValidityDuration(v string) {
	o.ValidityDuration = &v
}

// GetPendingDuration returns the PendingDuration field value if set, zero value otherwise.
func (o *LoyaltyPoints) GetPendingDuration() string {
	if o == nil || o.PendingDuration == nil {
		var ret string
		return ret
	}
	return *o.PendingDuration
}

// GetPendingDurationOk returns a tuple with the PendingDuration field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LoyaltyPoints) GetPendingDurationOk() (string, bool) {
	if o == nil || o.PendingDuration == nil {
		var ret string
		return ret, false
	}
	return *o.PendingDuration, true
}

// HasPendingDuration returns a boolean if a field has been set.
func (o *LoyaltyPoints) HasPendingDuration() bool {
	if o != nil && o.PendingDuration != nil {
		return true
	}

	return false
}

// SetPendingDuration gets a reference to the given string and assigns it to the PendingDuration field.
func (o *LoyaltyPoints) SetPendingDuration(v string) {
	o.PendingDuration = &v
}

// GetSubLedgerID returns the SubLedgerID field value if set, zero value otherwise.
func (o *LoyaltyPoints) GetSubLedgerID() string {
	if o == nil || o.SubLedgerID == nil {
		var ret string
		return ret
	}
	return *o.SubLedgerID
}

// GetSubLedgerIDOk returns a tuple with the SubLedgerID field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LoyaltyPoints) GetSubLedgerIDOk() (string, bool) {
	if o == nil || o.SubLedgerID == nil {
		var ret string
		return ret, false
	}
	return *o.SubLedgerID, true
}

// HasSubLedgerID returns a boolean if a field has been set.
func (o *LoyaltyPoints) HasSubLedgerID() bool {
	if o != nil && o.SubLedgerID != nil {
		return true
	}

	return false
}

// SetSubLedgerID gets a reference to the given string and assigns it to the SubLedgerID field.
func (o *LoyaltyPoints) SetSubLedgerID(v string) {
	o.SubLedgerID = &v
}

type NullableLoyaltyPoints struct {
	Value        LoyaltyPoints
	ExplicitNull bool
}

func (v NullableLoyaltyPoints) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableLoyaltyPoints) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}