/*
tiktok shop openapi

sdk for apis

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package affiliate_partner_v202411

import (
    "encoding/json"
    "github.com/Yuttachai-101/api_tiktok/utils"
)

            // checks if the AffiliatePartner202411SearchTapAffiliateOrdersResponseDataOrdersSkusEstimatedCommissionBase type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AffiliatePartner202411SearchTapAffiliateOrdersResponseDataOrdersSkusEstimatedCommissionBase{}

// AffiliatePartner202411SearchTapAffiliateOrdersResponseDataOrdersSkusEstimatedCommissionBase struct for AffiliatePartner202411SearchTapAffiliateOrdersResponseDataOrdersSkusEstimatedCommissionBase
type AffiliatePartner202411SearchTapAffiliateOrdersResponseDataOrdersSkusEstimatedCommissionBase struct {
    // Price amount for product, such as Rp 1000
    Amount *string `json:"amount,omitempty"`
    // Type of currency use
    Currency *string `json:"currency,omitempty"`
}

// NewAffiliatePartner202411SearchTapAffiliateOrdersResponseDataOrdersSkusEstimatedCommissionBase instantiates a new AffiliatePartner202411SearchTapAffiliateOrdersResponseDataOrdersSkusEstimatedCommissionBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAffiliatePartner202411SearchTapAffiliateOrdersResponseDataOrdersSkusEstimatedCommissionBase() *AffiliatePartner202411SearchTapAffiliateOrdersResponseDataOrdersSkusEstimatedCommissionBase {
    this := AffiliatePartner202411SearchTapAffiliateOrdersResponseDataOrdersSkusEstimatedCommissionBase{}
    return &this
}

// NewAffiliatePartner202411SearchTapAffiliateOrdersResponseDataOrdersSkusEstimatedCommissionBaseWithDefaults instantiates a new AffiliatePartner202411SearchTapAffiliateOrdersResponseDataOrdersSkusEstimatedCommissionBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAffiliatePartner202411SearchTapAffiliateOrdersResponseDataOrdersSkusEstimatedCommissionBaseWithDefaults() *AffiliatePartner202411SearchTapAffiliateOrdersResponseDataOrdersSkusEstimatedCommissionBase {
    this := AffiliatePartner202411SearchTapAffiliateOrdersResponseDataOrdersSkusEstimatedCommissionBase{}
    return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *AffiliatePartner202411SearchTapAffiliateOrdersResponseDataOrdersSkusEstimatedCommissionBase) GetAmount() string {
    if o == nil || utils.IsNil(o.Amount) {
        var ret string
        return ret
    }
    return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliatePartner202411SearchTapAffiliateOrdersResponseDataOrdersSkusEstimatedCommissionBase) GetAmountOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Amount) {
        return nil, false
    }
    return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *AffiliatePartner202411SearchTapAffiliateOrdersResponseDataOrdersSkusEstimatedCommissionBase) HasAmount() bool {
    if o != nil && !utils.IsNil(o.Amount) {
        return true
    }

    return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *AffiliatePartner202411SearchTapAffiliateOrdersResponseDataOrdersSkusEstimatedCommissionBase) SetAmount(v string) {
    o.Amount = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *AffiliatePartner202411SearchTapAffiliateOrdersResponseDataOrdersSkusEstimatedCommissionBase) GetCurrency() string {
    if o == nil || utils.IsNil(o.Currency) {
        var ret string
        return ret
    }
    return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliatePartner202411SearchTapAffiliateOrdersResponseDataOrdersSkusEstimatedCommissionBase) GetCurrencyOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Currency) {
        return nil, false
    }
    return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *AffiliatePartner202411SearchTapAffiliateOrdersResponseDataOrdersSkusEstimatedCommissionBase) HasCurrency() bool {
    if o != nil && !utils.IsNil(o.Currency) {
        return true
    }

    return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *AffiliatePartner202411SearchTapAffiliateOrdersResponseDataOrdersSkusEstimatedCommissionBase) SetCurrency(v string) {
    o.Currency = &v
}

func (o AffiliatePartner202411SearchTapAffiliateOrdersResponseDataOrdersSkusEstimatedCommissionBase) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o AffiliatePartner202411SearchTapAffiliateOrdersResponseDataOrdersSkusEstimatedCommissionBase) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.Amount) {
        toSerialize["amount"] = o.Amount
    }
    if !utils.IsNil(o.Currency) {
        toSerialize["currency"] = o.Currency
    }
    return toSerialize, nil
}

type NullableAffiliatePartner202411SearchTapAffiliateOrdersResponseDataOrdersSkusEstimatedCommissionBase struct {
	value *AffiliatePartner202411SearchTapAffiliateOrdersResponseDataOrdersSkusEstimatedCommissionBase
	isSet bool
}

func (v NullableAffiliatePartner202411SearchTapAffiliateOrdersResponseDataOrdersSkusEstimatedCommissionBase) Get() *AffiliatePartner202411SearchTapAffiliateOrdersResponseDataOrdersSkusEstimatedCommissionBase {
	return v.value
}

func (v *NullableAffiliatePartner202411SearchTapAffiliateOrdersResponseDataOrdersSkusEstimatedCommissionBase) Set(val *AffiliatePartner202411SearchTapAffiliateOrdersResponseDataOrdersSkusEstimatedCommissionBase) {
	v.value = val
	v.isSet = true
}

func (v NullableAffiliatePartner202411SearchTapAffiliateOrdersResponseDataOrdersSkusEstimatedCommissionBase) IsSet() bool {
	return v.isSet
}

func (v *NullableAffiliatePartner202411SearchTapAffiliateOrdersResponseDataOrdersSkusEstimatedCommissionBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAffiliatePartner202411SearchTapAffiliateOrdersResponseDataOrdersSkusEstimatedCommissionBase(val *AffiliatePartner202411SearchTapAffiliateOrdersResponseDataOrdersSkusEstimatedCommissionBase) *NullableAffiliatePartner202411SearchTapAffiliateOrdersResponseDataOrdersSkusEstimatedCommissionBase {
	return &NullableAffiliatePartner202411SearchTapAffiliateOrdersResponseDataOrdersSkusEstimatedCommissionBase{value: val, isSet: true}
}

func (v NullableAffiliatePartner202411SearchTapAffiliateOrdersResponseDataOrdersSkusEstimatedCommissionBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAffiliatePartner202411SearchTapAffiliateOrdersResponseDataOrdersSkusEstimatedCommissionBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


