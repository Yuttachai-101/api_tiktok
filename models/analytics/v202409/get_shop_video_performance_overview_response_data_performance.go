/*
tiktok shop openapi

sdk for apis

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package analytics_v202409

import (
    "encoding/json"
    "github.com/Yuttachai-101/api_tiktok/utils"
)

            // checks if the Analytics202409GetShopVideoPerformanceOverviewResponseDataPerformance type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Analytics202409GetShopVideoPerformanceOverviewResponseDataPerformance{}

// Analytics202409GetShopVideoPerformanceOverviewResponseDataPerformance struct for Analytics202409GetShopVideoPerformanceOverviewResponseDataPerformance
type Analytics202409GetShopVideoPerformanceOverviewResponseDataPerformance struct {
    // Same structure as \"intervals\" It contains data for the previous time range with the same range length and granularity of the current time range Example, if current time range (represented in start_time_ge and end_time_lt) is from 2024-09-01 to 2024-09-08) with granularity \"ALL\", the previous_intervals will contain data from 2024-08-25 to 2024-09-01 with granularity \"ALL\"
    ComparisonIntervals []Analytics202409GetShopVideoPerformanceOverviewResponseDataPerformanceComparisonIntervals `json:"comparison_intervals,omitempty"`
    // Interval data for the requested time range. The time range of each interval is determined by the granularity.
    Intervals []Analytics202409GetShopVideoPerformanceOverviewResponseDataPerformanceIntervals `json:"intervals,omitempty"`
}

// NewAnalytics202409GetShopVideoPerformanceOverviewResponseDataPerformance instantiates a new Analytics202409GetShopVideoPerformanceOverviewResponseDataPerformance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnalytics202409GetShopVideoPerformanceOverviewResponseDataPerformance() *Analytics202409GetShopVideoPerformanceOverviewResponseDataPerformance {
    this := Analytics202409GetShopVideoPerformanceOverviewResponseDataPerformance{}
    return &this
}

// NewAnalytics202409GetShopVideoPerformanceOverviewResponseDataPerformanceWithDefaults instantiates a new Analytics202409GetShopVideoPerformanceOverviewResponseDataPerformance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnalytics202409GetShopVideoPerformanceOverviewResponseDataPerformanceWithDefaults() *Analytics202409GetShopVideoPerformanceOverviewResponseDataPerformance {
    this := Analytics202409GetShopVideoPerformanceOverviewResponseDataPerformance{}
    return &this
}

// GetComparisonIntervals returns the ComparisonIntervals field value if set, zero value otherwise.
func (o *Analytics202409GetShopVideoPerformanceOverviewResponseDataPerformance) GetComparisonIntervals() []Analytics202409GetShopVideoPerformanceOverviewResponseDataPerformanceComparisonIntervals {
    if o == nil || utils.IsNil(o.ComparisonIntervals) {
        var ret []Analytics202409GetShopVideoPerformanceOverviewResponseDataPerformanceComparisonIntervals
        return ret
    }
    return o.ComparisonIntervals
}

// GetComparisonIntervalsOk returns a tuple with the ComparisonIntervals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Analytics202409GetShopVideoPerformanceOverviewResponseDataPerformance) GetComparisonIntervalsOk() ([]Analytics202409GetShopVideoPerformanceOverviewResponseDataPerformanceComparisonIntervals, bool) {
    if o == nil || utils.IsNil(o.ComparisonIntervals) {
        return nil, false
    }
    return o.ComparisonIntervals, true
}

// HasComparisonIntervals returns a boolean if a field has been set.
func (o *Analytics202409GetShopVideoPerformanceOverviewResponseDataPerformance) HasComparisonIntervals() bool {
    if o != nil && !utils.IsNil(o.ComparisonIntervals) {
        return true
    }

    return false
}

// SetComparisonIntervals gets a reference to the given []Analytics202409GetShopVideoPerformanceOverviewResponseDataPerformanceComparisonIntervals and assigns it to the ComparisonIntervals field.
func (o *Analytics202409GetShopVideoPerformanceOverviewResponseDataPerformance) SetComparisonIntervals(v []Analytics202409GetShopVideoPerformanceOverviewResponseDataPerformanceComparisonIntervals) {
    o.ComparisonIntervals = v
}

// GetIntervals returns the Intervals field value if set, zero value otherwise.
func (o *Analytics202409GetShopVideoPerformanceOverviewResponseDataPerformance) GetIntervals() []Analytics202409GetShopVideoPerformanceOverviewResponseDataPerformanceIntervals {
    if o == nil || utils.IsNil(o.Intervals) {
        var ret []Analytics202409GetShopVideoPerformanceOverviewResponseDataPerformanceIntervals
        return ret
    }
    return o.Intervals
}

// GetIntervalsOk returns a tuple with the Intervals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Analytics202409GetShopVideoPerformanceOverviewResponseDataPerformance) GetIntervalsOk() ([]Analytics202409GetShopVideoPerformanceOverviewResponseDataPerformanceIntervals, bool) {
    if o == nil || utils.IsNil(o.Intervals) {
        return nil, false
    }
    return o.Intervals, true
}

// HasIntervals returns a boolean if a field has been set.
func (o *Analytics202409GetShopVideoPerformanceOverviewResponseDataPerformance) HasIntervals() bool {
    if o != nil && !utils.IsNil(o.Intervals) {
        return true
    }

    return false
}

// SetIntervals gets a reference to the given []Analytics202409GetShopVideoPerformanceOverviewResponseDataPerformanceIntervals and assigns it to the Intervals field.
func (o *Analytics202409GetShopVideoPerformanceOverviewResponseDataPerformance) SetIntervals(v []Analytics202409GetShopVideoPerformanceOverviewResponseDataPerformanceIntervals) {
    o.Intervals = v
}

func (o Analytics202409GetShopVideoPerformanceOverviewResponseDataPerformance) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Analytics202409GetShopVideoPerformanceOverviewResponseDataPerformance) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.ComparisonIntervals) {
        toSerialize["comparison_intervals"] = o.ComparisonIntervals
    }
    if !utils.IsNil(o.Intervals) {
        toSerialize["intervals"] = o.Intervals
    }
    return toSerialize, nil
}

type NullableAnalytics202409GetShopVideoPerformanceOverviewResponseDataPerformance struct {
	value *Analytics202409GetShopVideoPerformanceOverviewResponseDataPerformance
	isSet bool
}

func (v NullableAnalytics202409GetShopVideoPerformanceOverviewResponseDataPerformance) Get() *Analytics202409GetShopVideoPerformanceOverviewResponseDataPerformance {
	return v.value
}

func (v *NullableAnalytics202409GetShopVideoPerformanceOverviewResponseDataPerformance) Set(val *Analytics202409GetShopVideoPerformanceOverviewResponseDataPerformance) {
	v.value = val
	v.isSet = true
}

func (v NullableAnalytics202409GetShopVideoPerformanceOverviewResponseDataPerformance) IsSet() bool {
	return v.isSet
}

func (v *NullableAnalytics202409GetShopVideoPerformanceOverviewResponseDataPerformance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnalytics202409GetShopVideoPerformanceOverviewResponseDataPerformance(val *Analytics202409GetShopVideoPerformanceOverviewResponseDataPerformance) *NullableAnalytics202409GetShopVideoPerformanceOverviewResponseDataPerformance {
	return &NullableAnalytics202409GetShopVideoPerformanceOverviewResponseDataPerformance{value: val, isSet: true}
}

func (v NullableAnalytics202409GetShopVideoPerformanceOverviewResponseDataPerformance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnalytics202409GetShopVideoPerformanceOverviewResponseDataPerformance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


