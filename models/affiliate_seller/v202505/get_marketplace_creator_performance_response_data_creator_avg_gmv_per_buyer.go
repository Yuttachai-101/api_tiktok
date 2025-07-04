/*
tiktok shop openapi

sdk for apis

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package affiliate_seller_v202505

import (
    "encoding/json"
    "github.com/Yuttachai-101/api_tiktok/utils"
)

            // checks if the AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyer type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyer{}

// AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyer struct for AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyer
type AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyer struct {
    // The average GMV per buyer amount.
    Amount *string `json:"amount,omitempty"`
    // The currency code.
    Currency *string `json:"currency,omitempty"`
}

// NewAffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyer instantiates a new AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyer() *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyer {
    this := AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyer{}
    return &this
}

// NewAffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyerWithDefaults instantiates a new AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyerWithDefaults() *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyer {
    this := AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyer{}
    return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyer) GetAmount() string {
    if o == nil || utils.IsNil(o.Amount) {
        var ret string
        return ret
    }
    return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyer) GetAmountOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Amount) {
        return nil, false
    }
    return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyer) HasAmount() bool {
    if o != nil && !utils.IsNil(o.Amount) {
        return true
    }

    return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyer) SetAmount(v string) {
    o.Amount = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyer) GetCurrency() string {
    if o == nil || utils.IsNil(o.Currency) {
        var ret string
        return ret
    }
    return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyer) GetCurrencyOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Currency) {
        return nil, false
    }
    return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyer) HasCurrency() bool {
    if o != nil && !utils.IsNil(o.Currency) {
        return true
    }

    return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyer) SetCurrency(v string) {
    o.Currency = &v
}

func (o AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyer) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyer) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.Amount) {
        toSerialize["amount"] = o.Amount
    }
    if !utils.IsNil(o.Currency) {
        toSerialize["currency"] = o.Currency
    }
    return toSerialize, nil
}

type NullableAffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyer struct {
	value *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyer
	isSet bool
}

func (v NullableAffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyer) Get() *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyer {
	return v.value
}

func (v *NullableAffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyer) Set(val *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyer) {
	v.value = val
	v.isSet = true
}

func (v NullableAffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyer) IsSet() bool {
	return v.isSet
}

func (v *NullableAffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyer(val *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyer) *NullableAffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyer {
	return &NullableAffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyer{value: val, isSet: true}
}

func (v NullableAffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


