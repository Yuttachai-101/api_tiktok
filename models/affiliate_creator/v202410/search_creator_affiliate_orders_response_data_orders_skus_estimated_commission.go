/*
tiktok shop openapi

sdk for apis

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package affiliate_creator_v202410

import (
    "encoding/json"
    "github.com/Yuttachai-101/api_tiktok/utils"
)

            // checks if the AffiliateCreator202410SearchCreatorAffiliateOrdersResponseDataOrdersSkusEstimatedCommission type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AffiliateCreator202410SearchCreatorAffiliateOrdersResponseDataOrdersSkusEstimatedCommission{}

// AffiliateCreator202410SearchCreatorAffiliateOrdersResponseDataOrdersSkusEstimatedCommission struct for AffiliateCreator202410SearchCreatorAffiliateOrdersResponseDataOrdersSkusEstimatedCommission
type AffiliateCreator202410SearchCreatorAffiliateOrdersResponseDataOrdersSkusEstimatedCommission struct {
    // The estimated commission amount.
    Amount *string `json:"amount,omitempty"`
    // The currency code.
    Currency *string `json:"currency,omitempty"`
}

// NewAffiliateCreator202410SearchCreatorAffiliateOrdersResponseDataOrdersSkusEstimatedCommission instantiates a new AffiliateCreator202410SearchCreatorAffiliateOrdersResponseDataOrdersSkusEstimatedCommission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAffiliateCreator202410SearchCreatorAffiliateOrdersResponseDataOrdersSkusEstimatedCommission() *AffiliateCreator202410SearchCreatorAffiliateOrdersResponseDataOrdersSkusEstimatedCommission {
    this := AffiliateCreator202410SearchCreatorAffiliateOrdersResponseDataOrdersSkusEstimatedCommission{}
    return &this
}

// NewAffiliateCreator202410SearchCreatorAffiliateOrdersResponseDataOrdersSkusEstimatedCommissionWithDefaults instantiates a new AffiliateCreator202410SearchCreatorAffiliateOrdersResponseDataOrdersSkusEstimatedCommission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAffiliateCreator202410SearchCreatorAffiliateOrdersResponseDataOrdersSkusEstimatedCommissionWithDefaults() *AffiliateCreator202410SearchCreatorAffiliateOrdersResponseDataOrdersSkusEstimatedCommission {
    this := AffiliateCreator202410SearchCreatorAffiliateOrdersResponseDataOrdersSkusEstimatedCommission{}
    return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *AffiliateCreator202410SearchCreatorAffiliateOrdersResponseDataOrdersSkusEstimatedCommission) GetAmount() string {
    if o == nil || utils.IsNil(o.Amount) {
        var ret string
        return ret
    }
    return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateCreator202410SearchCreatorAffiliateOrdersResponseDataOrdersSkusEstimatedCommission) GetAmountOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Amount) {
        return nil, false
    }
    return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *AffiliateCreator202410SearchCreatorAffiliateOrdersResponseDataOrdersSkusEstimatedCommission) HasAmount() bool {
    if o != nil && !utils.IsNil(o.Amount) {
        return true
    }

    return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *AffiliateCreator202410SearchCreatorAffiliateOrdersResponseDataOrdersSkusEstimatedCommission) SetAmount(v string) {
    o.Amount = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *AffiliateCreator202410SearchCreatorAffiliateOrdersResponseDataOrdersSkusEstimatedCommission) GetCurrency() string {
    if o == nil || utils.IsNil(o.Currency) {
        var ret string
        return ret
    }
    return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateCreator202410SearchCreatorAffiliateOrdersResponseDataOrdersSkusEstimatedCommission) GetCurrencyOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Currency) {
        return nil, false
    }
    return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *AffiliateCreator202410SearchCreatorAffiliateOrdersResponseDataOrdersSkusEstimatedCommission) HasCurrency() bool {
    if o != nil && !utils.IsNil(o.Currency) {
        return true
    }

    return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *AffiliateCreator202410SearchCreatorAffiliateOrdersResponseDataOrdersSkusEstimatedCommission) SetCurrency(v string) {
    o.Currency = &v
}

func (o AffiliateCreator202410SearchCreatorAffiliateOrdersResponseDataOrdersSkusEstimatedCommission) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o AffiliateCreator202410SearchCreatorAffiliateOrdersResponseDataOrdersSkusEstimatedCommission) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.Amount) {
        toSerialize["amount"] = o.Amount
    }
    if !utils.IsNil(o.Currency) {
        toSerialize["currency"] = o.Currency
    }
    return toSerialize, nil
}

type NullableAffiliateCreator202410SearchCreatorAffiliateOrdersResponseDataOrdersSkusEstimatedCommission struct {
	value *AffiliateCreator202410SearchCreatorAffiliateOrdersResponseDataOrdersSkusEstimatedCommission
	isSet bool
}

func (v NullableAffiliateCreator202410SearchCreatorAffiliateOrdersResponseDataOrdersSkusEstimatedCommission) Get() *AffiliateCreator202410SearchCreatorAffiliateOrdersResponseDataOrdersSkusEstimatedCommission {
	return v.value
}

func (v *NullableAffiliateCreator202410SearchCreatorAffiliateOrdersResponseDataOrdersSkusEstimatedCommission) Set(val *AffiliateCreator202410SearchCreatorAffiliateOrdersResponseDataOrdersSkusEstimatedCommission) {
	v.value = val
	v.isSet = true
}

func (v NullableAffiliateCreator202410SearchCreatorAffiliateOrdersResponseDataOrdersSkusEstimatedCommission) IsSet() bool {
	return v.isSet
}

func (v *NullableAffiliateCreator202410SearchCreatorAffiliateOrdersResponseDataOrdersSkusEstimatedCommission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAffiliateCreator202410SearchCreatorAffiliateOrdersResponseDataOrdersSkusEstimatedCommission(val *AffiliateCreator202410SearchCreatorAffiliateOrdersResponseDataOrdersSkusEstimatedCommission) *NullableAffiliateCreator202410SearchCreatorAffiliateOrdersResponseDataOrdersSkusEstimatedCommission {
	return &NullableAffiliateCreator202410SearchCreatorAffiliateOrdersResponseDataOrdersSkusEstimatedCommission{value: val, isSet: true}
}

func (v NullableAffiliateCreator202410SearchCreatorAffiliateOrdersResponseDataOrdersSkusEstimatedCommission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAffiliateCreator202410SearchCreatorAffiliateOrdersResponseDataOrdersSkusEstimatedCommission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


