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

            // checks if the AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyerRange type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyerRange{}

// AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyerRange struct for AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyerRange
type AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyerRange struct {
    // The currency code.
    Currency *string `json:"currency,omitempty"`
    // The formatted range of average GMV per buyer associated with the creator.
    FormattedRange *string `json:"formatted_range,omitempty"`
    // The largest average GMV per buyer.
    MaximumAmount *string `json:"maximum_amount,omitempty"`
    // The smallest average GMV per buyer.
    MinimumAmount *string `json:"minimum_amount,omitempty"`
}

// NewAffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyerRange instantiates a new AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyerRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyerRange() *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyerRange {
    this := AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyerRange{}
    return &this
}

// NewAffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyerRangeWithDefaults instantiates a new AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyerRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyerRangeWithDefaults() *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyerRange {
    this := AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyerRange{}
    return &this
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyerRange) GetCurrency() string {
    if o == nil || utils.IsNil(o.Currency) {
        var ret string
        return ret
    }
    return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyerRange) GetCurrencyOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Currency) {
        return nil, false
    }
    return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyerRange) HasCurrency() bool {
    if o != nil && !utils.IsNil(o.Currency) {
        return true
    }

    return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyerRange) SetCurrency(v string) {
    o.Currency = &v
}

// GetFormattedRange returns the FormattedRange field value if set, zero value otherwise.
func (o *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyerRange) GetFormattedRange() string {
    if o == nil || utils.IsNil(o.FormattedRange) {
        var ret string
        return ret
    }
    return *o.FormattedRange
}

// GetFormattedRangeOk returns a tuple with the FormattedRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyerRange) GetFormattedRangeOk() (*string, bool) {
    if o == nil || utils.IsNil(o.FormattedRange) {
        return nil, false
    }
    return o.FormattedRange, true
}

// HasFormattedRange returns a boolean if a field has been set.
func (o *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyerRange) HasFormattedRange() bool {
    if o != nil && !utils.IsNil(o.FormattedRange) {
        return true
    }

    return false
}

// SetFormattedRange gets a reference to the given string and assigns it to the FormattedRange field.
func (o *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyerRange) SetFormattedRange(v string) {
    o.FormattedRange = &v
}

// GetMaximumAmount returns the MaximumAmount field value if set, zero value otherwise.
func (o *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyerRange) GetMaximumAmount() string {
    if o == nil || utils.IsNil(o.MaximumAmount) {
        var ret string
        return ret
    }
    return *o.MaximumAmount
}

// GetMaximumAmountOk returns a tuple with the MaximumAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyerRange) GetMaximumAmountOk() (*string, bool) {
    if o == nil || utils.IsNil(o.MaximumAmount) {
        return nil, false
    }
    return o.MaximumAmount, true
}

// HasMaximumAmount returns a boolean if a field has been set.
func (o *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyerRange) HasMaximumAmount() bool {
    if o != nil && !utils.IsNil(o.MaximumAmount) {
        return true
    }

    return false
}

// SetMaximumAmount gets a reference to the given string and assigns it to the MaximumAmount field.
func (o *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyerRange) SetMaximumAmount(v string) {
    o.MaximumAmount = &v
}

// GetMinimumAmount returns the MinimumAmount field value if set, zero value otherwise.
func (o *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyerRange) GetMinimumAmount() string {
    if o == nil || utils.IsNil(o.MinimumAmount) {
        var ret string
        return ret
    }
    return *o.MinimumAmount
}

// GetMinimumAmountOk returns a tuple with the MinimumAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyerRange) GetMinimumAmountOk() (*string, bool) {
    if o == nil || utils.IsNil(o.MinimumAmount) {
        return nil, false
    }
    return o.MinimumAmount, true
}

// HasMinimumAmount returns a boolean if a field has been set.
func (o *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyerRange) HasMinimumAmount() bool {
    if o != nil && !utils.IsNil(o.MinimumAmount) {
        return true
    }

    return false
}

// SetMinimumAmount gets a reference to the given string and assigns it to the MinimumAmount field.
func (o *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyerRange) SetMinimumAmount(v string) {
    o.MinimumAmount = &v
}

func (o AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyerRange) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyerRange) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.Currency) {
        toSerialize["currency"] = o.Currency
    }
    if !utils.IsNil(o.FormattedRange) {
        toSerialize["formatted_range"] = o.FormattedRange
    }
    if !utils.IsNil(o.MaximumAmount) {
        toSerialize["maximum_amount"] = o.MaximumAmount
    }
    if !utils.IsNil(o.MinimumAmount) {
        toSerialize["minimum_amount"] = o.MinimumAmount
    }
    return toSerialize, nil
}

type NullableAffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyerRange struct {
	value *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyerRange
	isSet bool
}

func (v NullableAffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyerRange) Get() *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyerRange {
	return v.value
}

func (v *NullableAffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyerRange) Set(val *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyerRange) {
	v.value = val
	v.isSet = true
}

func (v NullableAffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyerRange) IsSet() bool {
	return v.isSet
}

func (v *NullableAffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyerRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyerRange(val *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyerRange) *NullableAffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyerRange {
	return &NullableAffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyerRange{value: val, isSet: true}
}

func (v NullableAffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyerRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorAvgGmvPerBuyerRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


