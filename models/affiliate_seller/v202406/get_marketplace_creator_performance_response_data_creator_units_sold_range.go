/*
tiktok shop openapi

sdk for apis

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package affiliate_seller_v202406

import (
    "encoding/json"
    "github.com/Yuttachai-101/api_tiktok/utils"
)

            // checks if the AffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorUnitsSoldRange type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorUnitsSoldRange{}

// AffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorUnitsSoldRange struct for AffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorUnitsSoldRange
type AffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorUnitsSoldRange struct {
    // The highest number of units sold.
    MaximumAmount *int64 `json:"maximum_amount,omitempty"`
    // The lowest number of units sold.
    MinimumAmount *int64 `json:"minimum_amount,omitempty"`
}

// NewAffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorUnitsSoldRange instantiates a new AffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorUnitsSoldRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorUnitsSoldRange() *AffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorUnitsSoldRange {
    this := AffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorUnitsSoldRange{}
    return &this
}

// NewAffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorUnitsSoldRangeWithDefaults instantiates a new AffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorUnitsSoldRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorUnitsSoldRangeWithDefaults() *AffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorUnitsSoldRange {
    this := AffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorUnitsSoldRange{}
    return &this
}

// GetMaximumAmount returns the MaximumAmount field value if set, zero value otherwise.
func (o *AffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorUnitsSoldRange) GetMaximumAmount() int64 {
    if o == nil || utils.IsNil(o.MaximumAmount) {
        var ret int64
        return ret
    }
    return *o.MaximumAmount
}

// GetMaximumAmountOk returns a tuple with the MaximumAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorUnitsSoldRange) GetMaximumAmountOk() (*int64, bool) {
    if o == nil || utils.IsNil(o.MaximumAmount) {
        return nil, false
    }
    return o.MaximumAmount, true
}

// HasMaximumAmount returns a boolean if a field has been set.
func (o *AffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorUnitsSoldRange) HasMaximumAmount() bool {
    if o != nil && !utils.IsNil(o.MaximumAmount) {
        return true
    }

    return false
}

// SetMaximumAmount gets a reference to the given int64 and assigns it to the MaximumAmount field.
func (o *AffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorUnitsSoldRange) SetMaximumAmount(v int64) {
    o.MaximumAmount = &v
}

// GetMinimumAmount returns the MinimumAmount field value if set, zero value otherwise.
func (o *AffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorUnitsSoldRange) GetMinimumAmount() int64 {
    if o == nil || utils.IsNil(o.MinimumAmount) {
        var ret int64
        return ret
    }
    return *o.MinimumAmount
}

// GetMinimumAmountOk returns a tuple with the MinimumAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorUnitsSoldRange) GetMinimumAmountOk() (*int64, bool) {
    if o == nil || utils.IsNil(o.MinimumAmount) {
        return nil, false
    }
    return o.MinimumAmount, true
}

// HasMinimumAmount returns a boolean if a field has been set.
func (o *AffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorUnitsSoldRange) HasMinimumAmount() bool {
    if o != nil && !utils.IsNil(o.MinimumAmount) {
        return true
    }

    return false
}

// SetMinimumAmount gets a reference to the given int64 and assigns it to the MinimumAmount field.
func (o *AffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorUnitsSoldRange) SetMinimumAmount(v int64) {
    o.MinimumAmount = &v
}

func (o AffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorUnitsSoldRange) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o AffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorUnitsSoldRange) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.MaximumAmount) {
        toSerialize["maximum_amount"] = o.MaximumAmount
    }
    if !utils.IsNil(o.MinimumAmount) {
        toSerialize["minimum_amount"] = o.MinimumAmount
    }
    return toSerialize, nil
}

type NullableAffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorUnitsSoldRange struct {
	value *AffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorUnitsSoldRange
	isSet bool
}

func (v NullableAffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorUnitsSoldRange) Get() *AffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorUnitsSoldRange {
	return v.value
}

func (v *NullableAffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorUnitsSoldRange) Set(val *AffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorUnitsSoldRange) {
	v.value = val
	v.isSet = true
}

func (v NullableAffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorUnitsSoldRange) IsSet() bool {
	return v.isSet
}

func (v *NullableAffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorUnitsSoldRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorUnitsSoldRange(val *AffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorUnitsSoldRange) *NullableAffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorUnitsSoldRange {
	return &NullableAffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorUnitsSoldRange{value: val, isSet: true}
}

func (v NullableAffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorUnitsSoldRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorUnitsSoldRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


