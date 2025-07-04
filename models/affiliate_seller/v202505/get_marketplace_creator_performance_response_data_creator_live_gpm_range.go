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

            // checks if the AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorLiveGpmRange type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorLiveGpmRange{}

// AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorLiveGpmRange struct for AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorLiveGpmRange
type AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorLiveGpmRange struct {
    // The currency code.
    Currency *string `json:"currency,omitempty"`
    // The formatted range of livestream GPM values associted with the creator. 
    FormattedRange *string `json:"formatted_range,omitempty"`
    // The highest livestream GPM value.
    MaximumAmount *string `json:"maximum_amount,omitempty"`
    // The lowest livestream GPM value.
    MinimumAmount *string `json:"minimum_amount,omitempty"`
}

// NewAffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorLiveGpmRange instantiates a new AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorLiveGpmRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorLiveGpmRange() *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorLiveGpmRange {
    this := AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorLiveGpmRange{}
    return &this
}

// NewAffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorLiveGpmRangeWithDefaults instantiates a new AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorLiveGpmRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorLiveGpmRangeWithDefaults() *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorLiveGpmRange {
    this := AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorLiveGpmRange{}
    return &this
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorLiveGpmRange) GetCurrency() string {
    if o == nil || utils.IsNil(o.Currency) {
        var ret string
        return ret
    }
    return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorLiveGpmRange) GetCurrencyOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Currency) {
        return nil, false
    }
    return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorLiveGpmRange) HasCurrency() bool {
    if o != nil && !utils.IsNil(o.Currency) {
        return true
    }

    return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorLiveGpmRange) SetCurrency(v string) {
    o.Currency = &v
}

// GetFormattedRange returns the FormattedRange field value if set, zero value otherwise.
func (o *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorLiveGpmRange) GetFormattedRange() string {
    if o == nil || utils.IsNil(o.FormattedRange) {
        var ret string
        return ret
    }
    return *o.FormattedRange
}

// GetFormattedRangeOk returns a tuple with the FormattedRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorLiveGpmRange) GetFormattedRangeOk() (*string, bool) {
    if o == nil || utils.IsNil(o.FormattedRange) {
        return nil, false
    }
    return o.FormattedRange, true
}

// HasFormattedRange returns a boolean if a field has been set.
func (o *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorLiveGpmRange) HasFormattedRange() bool {
    if o != nil && !utils.IsNil(o.FormattedRange) {
        return true
    }

    return false
}

// SetFormattedRange gets a reference to the given string and assigns it to the FormattedRange field.
func (o *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorLiveGpmRange) SetFormattedRange(v string) {
    o.FormattedRange = &v
}

// GetMaximumAmount returns the MaximumAmount field value if set, zero value otherwise.
func (o *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorLiveGpmRange) GetMaximumAmount() string {
    if o == nil || utils.IsNil(o.MaximumAmount) {
        var ret string
        return ret
    }
    return *o.MaximumAmount
}

// GetMaximumAmountOk returns a tuple with the MaximumAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorLiveGpmRange) GetMaximumAmountOk() (*string, bool) {
    if o == nil || utils.IsNil(o.MaximumAmount) {
        return nil, false
    }
    return o.MaximumAmount, true
}

// HasMaximumAmount returns a boolean if a field has been set.
func (o *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorLiveGpmRange) HasMaximumAmount() bool {
    if o != nil && !utils.IsNil(o.MaximumAmount) {
        return true
    }

    return false
}

// SetMaximumAmount gets a reference to the given string and assigns it to the MaximumAmount field.
func (o *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorLiveGpmRange) SetMaximumAmount(v string) {
    o.MaximumAmount = &v
}

// GetMinimumAmount returns the MinimumAmount field value if set, zero value otherwise.
func (o *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorLiveGpmRange) GetMinimumAmount() string {
    if o == nil || utils.IsNil(o.MinimumAmount) {
        var ret string
        return ret
    }
    return *o.MinimumAmount
}

// GetMinimumAmountOk returns a tuple with the MinimumAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorLiveGpmRange) GetMinimumAmountOk() (*string, bool) {
    if o == nil || utils.IsNil(o.MinimumAmount) {
        return nil, false
    }
    return o.MinimumAmount, true
}

// HasMinimumAmount returns a boolean if a field has been set.
func (o *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorLiveGpmRange) HasMinimumAmount() bool {
    if o != nil && !utils.IsNil(o.MinimumAmount) {
        return true
    }

    return false
}

// SetMinimumAmount gets a reference to the given string and assigns it to the MinimumAmount field.
func (o *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorLiveGpmRange) SetMinimumAmount(v string) {
    o.MinimumAmount = &v
}

func (o AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorLiveGpmRange) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorLiveGpmRange) ToMap() (map[string]interface{}, error) {
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

type NullableAffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorLiveGpmRange struct {
	value *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorLiveGpmRange
	isSet bool
}

func (v NullableAffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorLiveGpmRange) Get() *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorLiveGpmRange {
	return v.value
}

func (v *NullableAffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorLiveGpmRange) Set(val *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorLiveGpmRange) {
	v.value = val
	v.isSet = true
}

func (v NullableAffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorLiveGpmRange) IsSet() bool {
	return v.isSet
}

func (v *NullableAffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorLiveGpmRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorLiveGpmRange(val *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorLiveGpmRange) *NullableAffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorLiveGpmRange {
	return &NullableAffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorLiveGpmRange{value: val, isSet: true}
}

func (v NullableAffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorLiveGpmRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorLiveGpmRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


