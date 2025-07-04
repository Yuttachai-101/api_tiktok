/*
tiktok shop openapi

sdk for apis

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package affiliate_seller_v202410

import (
    "encoding/json"
    "github.com/Yuttachai-101/api_tiktok/utils"
)

            // checks if the AffiliateSeller202410SearchSellerAffiliateOrdersResponseDataOrdersSkusActualCofundedCreatorBonusAmount type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AffiliateSeller202410SearchSellerAffiliateOrdersResponseDataOrdersSkusActualCofundedCreatorBonusAmount{}

// AffiliateSeller202410SearchSellerAffiliateOrdersResponseDataOrdersSkusActualCofundedCreatorBonusAmount struct for AffiliateSeller202410SearchSellerAffiliateOrdersResponseDataOrdersSkusActualCofundedCreatorBonusAmount
type AffiliateSeller202410SearchSellerAffiliateOrdersResponseDataOrdersSkusActualCofundedCreatorBonusAmount struct {
    // The commission amount.
    Amount *string `json:"amount,omitempty"`
    // The commission currency.
    Currency *string `json:"currency,omitempty"`
}

// NewAffiliateSeller202410SearchSellerAffiliateOrdersResponseDataOrdersSkusActualCofundedCreatorBonusAmount instantiates a new AffiliateSeller202410SearchSellerAffiliateOrdersResponseDataOrdersSkusActualCofundedCreatorBonusAmount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAffiliateSeller202410SearchSellerAffiliateOrdersResponseDataOrdersSkusActualCofundedCreatorBonusAmount() *AffiliateSeller202410SearchSellerAffiliateOrdersResponseDataOrdersSkusActualCofundedCreatorBonusAmount {
    this := AffiliateSeller202410SearchSellerAffiliateOrdersResponseDataOrdersSkusActualCofundedCreatorBonusAmount{}
    return &this
}

// NewAffiliateSeller202410SearchSellerAffiliateOrdersResponseDataOrdersSkusActualCofundedCreatorBonusAmountWithDefaults instantiates a new AffiliateSeller202410SearchSellerAffiliateOrdersResponseDataOrdersSkusActualCofundedCreatorBonusAmount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAffiliateSeller202410SearchSellerAffiliateOrdersResponseDataOrdersSkusActualCofundedCreatorBonusAmountWithDefaults() *AffiliateSeller202410SearchSellerAffiliateOrdersResponseDataOrdersSkusActualCofundedCreatorBonusAmount {
    this := AffiliateSeller202410SearchSellerAffiliateOrdersResponseDataOrdersSkusActualCofundedCreatorBonusAmount{}
    return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *AffiliateSeller202410SearchSellerAffiliateOrdersResponseDataOrdersSkusActualCofundedCreatorBonusAmount) GetAmount() string {
    if o == nil || utils.IsNil(o.Amount) {
        var ret string
        return ret
    }
    return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateSeller202410SearchSellerAffiliateOrdersResponseDataOrdersSkusActualCofundedCreatorBonusAmount) GetAmountOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Amount) {
        return nil, false
    }
    return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *AffiliateSeller202410SearchSellerAffiliateOrdersResponseDataOrdersSkusActualCofundedCreatorBonusAmount) HasAmount() bool {
    if o != nil && !utils.IsNil(o.Amount) {
        return true
    }

    return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *AffiliateSeller202410SearchSellerAffiliateOrdersResponseDataOrdersSkusActualCofundedCreatorBonusAmount) SetAmount(v string) {
    o.Amount = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *AffiliateSeller202410SearchSellerAffiliateOrdersResponseDataOrdersSkusActualCofundedCreatorBonusAmount) GetCurrency() string {
    if o == nil || utils.IsNil(o.Currency) {
        var ret string
        return ret
    }
    return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateSeller202410SearchSellerAffiliateOrdersResponseDataOrdersSkusActualCofundedCreatorBonusAmount) GetCurrencyOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Currency) {
        return nil, false
    }
    return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *AffiliateSeller202410SearchSellerAffiliateOrdersResponseDataOrdersSkusActualCofundedCreatorBonusAmount) HasCurrency() bool {
    if o != nil && !utils.IsNil(o.Currency) {
        return true
    }

    return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *AffiliateSeller202410SearchSellerAffiliateOrdersResponseDataOrdersSkusActualCofundedCreatorBonusAmount) SetCurrency(v string) {
    o.Currency = &v
}

func (o AffiliateSeller202410SearchSellerAffiliateOrdersResponseDataOrdersSkusActualCofundedCreatorBonusAmount) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o AffiliateSeller202410SearchSellerAffiliateOrdersResponseDataOrdersSkusActualCofundedCreatorBonusAmount) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.Amount) {
        toSerialize["amount"] = o.Amount
    }
    if !utils.IsNil(o.Currency) {
        toSerialize["currency"] = o.Currency
    }
    return toSerialize, nil
}

type NullableAffiliateSeller202410SearchSellerAffiliateOrdersResponseDataOrdersSkusActualCofundedCreatorBonusAmount struct {
	value *AffiliateSeller202410SearchSellerAffiliateOrdersResponseDataOrdersSkusActualCofundedCreatorBonusAmount
	isSet bool
}

func (v NullableAffiliateSeller202410SearchSellerAffiliateOrdersResponseDataOrdersSkusActualCofundedCreatorBonusAmount) Get() *AffiliateSeller202410SearchSellerAffiliateOrdersResponseDataOrdersSkusActualCofundedCreatorBonusAmount {
	return v.value
}

func (v *NullableAffiliateSeller202410SearchSellerAffiliateOrdersResponseDataOrdersSkusActualCofundedCreatorBonusAmount) Set(val *AffiliateSeller202410SearchSellerAffiliateOrdersResponseDataOrdersSkusActualCofundedCreatorBonusAmount) {
	v.value = val
	v.isSet = true
}

func (v NullableAffiliateSeller202410SearchSellerAffiliateOrdersResponseDataOrdersSkusActualCofundedCreatorBonusAmount) IsSet() bool {
	return v.isSet
}

func (v *NullableAffiliateSeller202410SearchSellerAffiliateOrdersResponseDataOrdersSkusActualCofundedCreatorBonusAmount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAffiliateSeller202410SearchSellerAffiliateOrdersResponseDataOrdersSkusActualCofundedCreatorBonusAmount(val *AffiliateSeller202410SearchSellerAffiliateOrdersResponseDataOrdersSkusActualCofundedCreatorBonusAmount) *NullableAffiliateSeller202410SearchSellerAffiliateOrdersResponseDataOrdersSkusActualCofundedCreatorBonusAmount {
	return &NullableAffiliateSeller202410SearchSellerAffiliateOrdersResponseDataOrdersSkusActualCofundedCreatorBonusAmount{value: val, isSet: true}
}

func (v NullableAffiliateSeller202410SearchSellerAffiliateOrdersResponseDataOrdersSkusActualCofundedCreatorBonusAmount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAffiliateSeller202410SearchSellerAffiliateOrdersResponseDataOrdersSkusActualCofundedCreatorBonusAmount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


