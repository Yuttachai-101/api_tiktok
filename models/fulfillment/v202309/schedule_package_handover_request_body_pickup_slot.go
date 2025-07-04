/*
tiktok shop openapi

sdk for apis

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fulfillment_v202309

import (
    "encoding/json"
    "github.com/Yuttachai-101/api_tiktok/utils"
)

            // checks if the Fulfillment202309SchedulePackageHandoverRequestBodyPickupSlot type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Fulfillment202309SchedulePackageHandoverRequestBodyPickupSlot{}

// Fulfillment202309SchedulePackageHandoverRequestBodyPickupSlot struct for Fulfillment202309SchedulePackageHandoverRequestBodyPickupSlot
type Fulfillment202309SchedulePackageHandoverRequestBodyPickupSlot struct {
    // The end date and time of the package pickup time slot. 
    EndTime *int64 `json:"end_time,omitempty"`
    // The start date and time of the package pickup time slot.
    StartTime *int64 `json:"start_time,omitempty"`
}

// NewFulfillment202309SchedulePackageHandoverRequestBodyPickupSlot instantiates a new Fulfillment202309SchedulePackageHandoverRequestBodyPickupSlot object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFulfillment202309SchedulePackageHandoverRequestBodyPickupSlot() *Fulfillment202309SchedulePackageHandoverRequestBodyPickupSlot {
    this := Fulfillment202309SchedulePackageHandoverRequestBodyPickupSlot{}
    return &this
}

// NewFulfillment202309SchedulePackageHandoverRequestBodyPickupSlotWithDefaults instantiates a new Fulfillment202309SchedulePackageHandoverRequestBodyPickupSlot object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFulfillment202309SchedulePackageHandoverRequestBodyPickupSlotWithDefaults() *Fulfillment202309SchedulePackageHandoverRequestBodyPickupSlot {
    this := Fulfillment202309SchedulePackageHandoverRequestBodyPickupSlot{}
    return &this
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *Fulfillment202309SchedulePackageHandoverRequestBodyPickupSlot) GetEndTime() int64 {
    if o == nil || utils.IsNil(o.EndTime) {
        var ret int64
        return ret
    }
    return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fulfillment202309SchedulePackageHandoverRequestBodyPickupSlot) GetEndTimeOk() (*int64, bool) {
    if o == nil || utils.IsNil(o.EndTime) {
        return nil, false
    }
    return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *Fulfillment202309SchedulePackageHandoverRequestBodyPickupSlot) HasEndTime() bool {
    if o != nil && !utils.IsNil(o.EndTime) {
        return true
    }

    return false
}

// SetEndTime gets a reference to the given int64 and assigns it to the EndTime field.
func (o *Fulfillment202309SchedulePackageHandoverRequestBodyPickupSlot) SetEndTime(v int64) {
    o.EndTime = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *Fulfillment202309SchedulePackageHandoverRequestBodyPickupSlot) GetStartTime() int64 {
    if o == nil || utils.IsNil(o.StartTime) {
        var ret int64
        return ret
    }
    return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fulfillment202309SchedulePackageHandoverRequestBodyPickupSlot) GetStartTimeOk() (*int64, bool) {
    if o == nil || utils.IsNil(o.StartTime) {
        return nil, false
    }
    return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *Fulfillment202309SchedulePackageHandoverRequestBodyPickupSlot) HasStartTime() bool {
    if o != nil && !utils.IsNil(o.StartTime) {
        return true
    }

    return false
}

// SetStartTime gets a reference to the given int64 and assigns it to the StartTime field.
func (o *Fulfillment202309SchedulePackageHandoverRequestBodyPickupSlot) SetStartTime(v int64) {
    o.StartTime = &v
}

func (o Fulfillment202309SchedulePackageHandoverRequestBodyPickupSlot) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Fulfillment202309SchedulePackageHandoverRequestBodyPickupSlot) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.EndTime) {
        toSerialize["end_time"] = o.EndTime
    }
    if !utils.IsNil(o.StartTime) {
        toSerialize["start_time"] = o.StartTime
    }
    return toSerialize, nil
}

type NullableFulfillment202309SchedulePackageHandoverRequestBodyPickupSlot struct {
	value *Fulfillment202309SchedulePackageHandoverRequestBodyPickupSlot
	isSet bool
}

func (v NullableFulfillment202309SchedulePackageHandoverRequestBodyPickupSlot) Get() *Fulfillment202309SchedulePackageHandoverRequestBodyPickupSlot {
	return v.value
}

func (v *NullableFulfillment202309SchedulePackageHandoverRequestBodyPickupSlot) Set(val *Fulfillment202309SchedulePackageHandoverRequestBodyPickupSlot) {
	v.value = val
	v.isSet = true
}

func (v NullableFulfillment202309SchedulePackageHandoverRequestBodyPickupSlot) IsSet() bool {
	return v.isSet
}

func (v *NullableFulfillment202309SchedulePackageHandoverRequestBodyPickupSlot) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFulfillment202309SchedulePackageHandoverRequestBodyPickupSlot(val *Fulfillment202309SchedulePackageHandoverRequestBodyPickupSlot) *NullableFulfillment202309SchedulePackageHandoverRequestBodyPickupSlot {
	return &NullableFulfillment202309SchedulePackageHandoverRequestBodyPickupSlot{value: val, isSet: true}
}

func (v NullableFulfillment202309SchedulePackageHandoverRequestBodyPickupSlot) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFulfillment202309SchedulePackageHandoverRequestBodyPickupSlot) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


