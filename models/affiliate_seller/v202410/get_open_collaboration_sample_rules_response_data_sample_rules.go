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

            // checks if the AffiliateSeller202410GetOpenCollaborationSampleRulesResponseDataSampleRules type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AffiliateSeller202410GetOpenCollaborationSampleRulesResponseDataSampleRules{}

// AffiliateSeller202410GetOpenCollaborationSampleRulesResponseDataSampleRules struct for AffiliateSeller202410GetOpenCollaborationSampleRulesResponseDataSampleRules
type AffiliateSeller202410GetOpenCollaborationSampleRulesResponseDataSampleRules struct {
    // The quantity of remaining samples available for application.
    AvailableQuantity *int32 `json:"available_quantity,omitempty"`
    // The end time of the period during which a creator can apply for samples.
    EndTime *int64 `json:"end_time,omitempty"`
    // Whether samples are timely unlimited for request: - If true, it means available all the time - If false, it means sample only available from `begin_time` to `end_time`.
    IsSampleTimeUnlimited *bool `json:"is_sample_time_unlimited,omitempty"`
    // The unique id of product.
    ProductId *string `json:"product_id,omitempty"`
    // The sample total quantity provided by seller which creators can apply for.
    SampleQuota *int32 `json:"sample_quota,omitempty"`
    // The start time of the period during which a creator can apply for samples. You must specify `start_time` and `end_time` together.
    StartTime *int64 `json:"start_time,omitempty"`
    // Free sample rule status: - NOT_STARTED. Not yet reached the specified `begin_time`. - ONGOING. The sample rule is effective and the collaboration is ongoing. - `NO_LEFT_COUNT`. `sample_quantity_available` is `0`. - PLAN_EXCEPTION. The product commission collaboration status is abnormal - EXPIRED. Exceeded the `end_time`.
    Status *string `json:"status,omitempty"`
    Thresholds *AffiliateSeller202410GetOpenCollaborationSampleRulesResponseDataSampleRulesThresholds `json:"thresholds,omitempty"`
}

// NewAffiliateSeller202410GetOpenCollaborationSampleRulesResponseDataSampleRules instantiates a new AffiliateSeller202410GetOpenCollaborationSampleRulesResponseDataSampleRules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAffiliateSeller202410GetOpenCollaborationSampleRulesResponseDataSampleRules() *AffiliateSeller202410GetOpenCollaborationSampleRulesResponseDataSampleRules {
    this := AffiliateSeller202410GetOpenCollaborationSampleRulesResponseDataSampleRules{}
    return &this
}

// NewAffiliateSeller202410GetOpenCollaborationSampleRulesResponseDataSampleRulesWithDefaults instantiates a new AffiliateSeller202410GetOpenCollaborationSampleRulesResponseDataSampleRules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAffiliateSeller202410GetOpenCollaborationSampleRulesResponseDataSampleRulesWithDefaults() *AffiliateSeller202410GetOpenCollaborationSampleRulesResponseDataSampleRules {
    this := AffiliateSeller202410GetOpenCollaborationSampleRulesResponseDataSampleRules{}
    return &this
}

// GetAvailableQuantity returns the AvailableQuantity field value if set, zero value otherwise.
func (o *AffiliateSeller202410GetOpenCollaborationSampleRulesResponseDataSampleRules) GetAvailableQuantity() int32 {
    if o == nil || utils.IsNil(o.AvailableQuantity) {
        var ret int32
        return ret
    }
    return *o.AvailableQuantity
}

// GetAvailableQuantityOk returns a tuple with the AvailableQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateSeller202410GetOpenCollaborationSampleRulesResponseDataSampleRules) GetAvailableQuantityOk() (*int32, bool) {
    if o == nil || utils.IsNil(o.AvailableQuantity) {
        return nil, false
    }
    return o.AvailableQuantity, true
}

// HasAvailableQuantity returns a boolean if a field has been set.
func (o *AffiliateSeller202410GetOpenCollaborationSampleRulesResponseDataSampleRules) HasAvailableQuantity() bool {
    if o != nil && !utils.IsNil(o.AvailableQuantity) {
        return true
    }

    return false
}

// SetAvailableQuantity gets a reference to the given int32 and assigns it to the AvailableQuantity field.
func (o *AffiliateSeller202410GetOpenCollaborationSampleRulesResponseDataSampleRules) SetAvailableQuantity(v int32) {
    o.AvailableQuantity = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *AffiliateSeller202410GetOpenCollaborationSampleRulesResponseDataSampleRules) GetEndTime() int64 {
    if o == nil || utils.IsNil(o.EndTime) {
        var ret int64
        return ret
    }
    return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateSeller202410GetOpenCollaborationSampleRulesResponseDataSampleRules) GetEndTimeOk() (*int64, bool) {
    if o == nil || utils.IsNil(o.EndTime) {
        return nil, false
    }
    return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *AffiliateSeller202410GetOpenCollaborationSampleRulesResponseDataSampleRules) HasEndTime() bool {
    if o != nil && !utils.IsNil(o.EndTime) {
        return true
    }

    return false
}

// SetEndTime gets a reference to the given int64 and assigns it to the EndTime field.
func (o *AffiliateSeller202410GetOpenCollaborationSampleRulesResponseDataSampleRules) SetEndTime(v int64) {
    o.EndTime = &v
}

// GetIsSampleTimeUnlimited returns the IsSampleTimeUnlimited field value if set, zero value otherwise.
func (o *AffiliateSeller202410GetOpenCollaborationSampleRulesResponseDataSampleRules) GetIsSampleTimeUnlimited() bool {
    if o == nil || utils.IsNil(o.IsSampleTimeUnlimited) {
        var ret bool
        return ret
    }
    return *o.IsSampleTimeUnlimited
}

// GetIsSampleTimeUnlimitedOk returns a tuple with the IsSampleTimeUnlimited field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateSeller202410GetOpenCollaborationSampleRulesResponseDataSampleRules) GetIsSampleTimeUnlimitedOk() (*bool, bool) {
    if o == nil || utils.IsNil(o.IsSampleTimeUnlimited) {
        return nil, false
    }
    return o.IsSampleTimeUnlimited, true
}

// HasIsSampleTimeUnlimited returns a boolean if a field has been set.
func (o *AffiliateSeller202410GetOpenCollaborationSampleRulesResponseDataSampleRules) HasIsSampleTimeUnlimited() bool {
    if o != nil && !utils.IsNil(o.IsSampleTimeUnlimited) {
        return true
    }

    return false
}

// SetIsSampleTimeUnlimited gets a reference to the given bool and assigns it to the IsSampleTimeUnlimited field.
func (o *AffiliateSeller202410GetOpenCollaborationSampleRulesResponseDataSampleRules) SetIsSampleTimeUnlimited(v bool) {
    o.IsSampleTimeUnlimited = &v
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *AffiliateSeller202410GetOpenCollaborationSampleRulesResponseDataSampleRules) GetProductId() string {
    if o == nil || utils.IsNil(o.ProductId) {
        var ret string
        return ret
    }
    return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateSeller202410GetOpenCollaborationSampleRulesResponseDataSampleRules) GetProductIdOk() (*string, bool) {
    if o == nil || utils.IsNil(o.ProductId) {
        return nil, false
    }
    return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *AffiliateSeller202410GetOpenCollaborationSampleRulesResponseDataSampleRules) HasProductId() bool {
    if o != nil && !utils.IsNil(o.ProductId) {
        return true
    }

    return false
}

// SetProductId gets a reference to the given string and assigns it to the ProductId field.
func (o *AffiliateSeller202410GetOpenCollaborationSampleRulesResponseDataSampleRules) SetProductId(v string) {
    o.ProductId = &v
}

// GetSampleQuota returns the SampleQuota field value if set, zero value otherwise.
func (o *AffiliateSeller202410GetOpenCollaborationSampleRulesResponseDataSampleRules) GetSampleQuota() int32 {
    if o == nil || utils.IsNil(o.SampleQuota) {
        var ret int32
        return ret
    }
    return *o.SampleQuota
}

// GetSampleQuotaOk returns a tuple with the SampleQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateSeller202410GetOpenCollaborationSampleRulesResponseDataSampleRules) GetSampleQuotaOk() (*int32, bool) {
    if o == nil || utils.IsNil(o.SampleQuota) {
        return nil, false
    }
    return o.SampleQuota, true
}

// HasSampleQuota returns a boolean if a field has been set.
func (o *AffiliateSeller202410GetOpenCollaborationSampleRulesResponseDataSampleRules) HasSampleQuota() bool {
    if o != nil && !utils.IsNil(o.SampleQuota) {
        return true
    }

    return false
}

// SetSampleQuota gets a reference to the given int32 and assigns it to the SampleQuota field.
func (o *AffiliateSeller202410GetOpenCollaborationSampleRulesResponseDataSampleRules) SetSampleQuota(v int32) {
    o.SampleQuota = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *AffiliateSeller202410GetOpenCollaborationSampleRulesResponseDataSampleRules) GetStartTime() int64 {
    if o == nil || utils.IsNil(o.StartTime) {
        var ret int64
        return ret
    }
    return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateSeller202410GetOpenCollaborationSampleRulesResponseDataSampleRules) GetStartTimeOk() (*int64, bool) {
    if o == nil || utils.IsNil(o.StartTime) {
        return nil, false
    }
    return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *AffiliateSeller202410GetOpenCollaborationSampleRulesResponseDataSampleRules) HasStartTime() bool {
    if o != nil && !utils.IsNil(o.StartTime) {
        return true
    }

    return false
}

// SetStartTime gets a reference to the given int64 and assigns it to the StartTime field.
func (o *AffiliateSeller202410GetOpenCollaborationSampleRulesResponseDataSampleRules) SetStartTime(v int64) {
    o.StartTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AffiliateSeller202410GetOpenCollaborationSampleRulesResponseDataSampleRules) GetStatus() string {
    if o == nil || utils.IsNil(o.Status) {
        var ret string
        return ret
    }
    return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateSeller202410GetOpenCollaborationSampleRulesResponseDataSampleRules) GetStatusOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Status) {
        return nil, false
    }
    return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AffiliateSeller202410GetOpenCollaborationSampleRulesResponseDataSampleRules) HasStatus() bool {
    if o != nil && !utils.IsNil(o.Status) {
        return true
    }

    return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AffiliateSeller202410GetOpenCollaborationSampleRulesResponseDataSampleRules) SetStatus(v string) {
    o.Status = &v
}

// GetThresholds returns the Thresholds field value if set, zero value otherwise.
func (o *AffiliateSeller202410GetOpenCollaborationSampleRulesResponseDataSampleRules) GetThresholds() AffiliateSeller202410GetOpenCollaborationSampleRulesResponseDataSampleRulesThresholds {
    if o == nil || utils.IsNil(o.Thresholds) {
        var ret AffiliateSeller202410GetOpenCollaborationSampleRulesResponseDataSampleRulesThresholds
        return ret
    }
    return *o.Thresholds
}

// GetThresholdsOk returns a tuple with the Thresholds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateSeller202410GetOpenCollaborationSampleRulesResponseDataSampleRules) GetThresholdsOk() (*AffiliateSeller202410GetOpenCollaborationSampleRulesResponseDataSampleRulesThresholds, bool) {
    if o == nil || utils.IsNil(o.Thresholds) {
        return nil, false
    }
    return o.Thresholds, true
}

// HasThresholds returns a boolean if a field has been set.
func (o *AffiliateSeller202410GetOpenCollaborationSampleRulesResponseDataSampleRules) HasThresholds() bool {
    if o != nil && !utils.IsNil(o.Thresholds) {
        return true
    }

    return false
}

// SetThresholds gets a reference to the given AffiliateSeller202410GetOpenCollaborationSampleRulesResponseDataSampleRulesThresholds and assigns it to the Thresholds field.
func (o *AffiliateSeller202410GetOpenCollaborationSampleRulesResponseDataSampleRules) SetThresholds(v AffiliateSeller202410GetOpenCollaborationSampleRulesResponseDataSampleRulesThresholds) {
    o.Thresholds = &v
}

func (o AffiliateSeller202410GetOpenCollaborationSampleRulesResponseDataSampleRules) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o AffiliateSeller202410GetOpenCollaborationSampleRulesResponseDataSampleRules) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.AvailableQuantity) {
        toSerialize["available_quantity"] = o.AvailableQuantity
    }
    if !utils.IsNil(o.EndTime) {
        toSerialize["end_time"] = o.EndTime
    }
    if !utils.IsNil(o.IsSampleTimeUnlimited) {
        toSerialize["is_sample_time_unlimited"] = o.IsSampleTimeUnlimited
    }
    if !utils.IsNil(o.ProductId) {
        toSerialize["product_id"] = o.ProductId
    }
    if !utils.IsNil(o.SampleQuota) {
        toSerialize["sample_quota"] = o.SampleQuota
    }
    if !utils.IsNil(o.StartTime) {
        toSerialize["start_time"] = o.StartTime
    }
    if !utils.IsNil(o.Status) {
        toSerialize["status"] = o.Status
    }
    if !utils.IsNil(o.Thresholds) {
        toSerialize["thresholds"] = o.Thresholds
    }
    return toSerialize, nil
}

type NullableAffiliateSeller202410GetOpenCollaborationSampleRulesResponseDataSampleRules struct {
	value *AffiliateSeller202410GetOpenCollaborationSampleRulesResponseDataSampleRules
	isSet bool
}

func (v NullableAffiliateSeller202410GetOpenCollaborationSampleRulesResponseDataSampleRules) Get() *AffiliateSeller202410GetOpenCollaborationSampleRulesResponseDataSampleRules {
	return v.value
}

func (v *NullableAffiliateSeller202410GetOpenCollaborationSampleRulesResponseDataSampleRules) Set(val *AffiliateSeller202410GetOpenCollaborationSampleRulesResponseDataSampleRules) {
	v.value = val
	v.isSet = true
}

func (v NullableAffiliateSeller202410GetOpenCollaborationSampleRulesResponseDataSampleRules) IsSet() bool {
	return v.isSet
}

func (v *NullableAffiliateSeller202410GetOpenCollaborationSampleRulesResponseDataSampleRules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAffiliateSeller202410GetOpenCollaborationSampleRulesResponseDataSampleRules(val *AffiliateSeller202410GetOpenCollaborationSampleRulesResponseDataSampleRules) *NullableAffiliateSeller202410GetOpenCollaborationSampleRulesResponseDataSampleRules {
	return &NullableAffiliateSeller202410GetOpenCollaborationSampleRulesResponseDataSampleRules{value: val, isSet: true}
}

func (v NullableAffiliateSeller202410GetOpenCollaborationSampleRulesResponseDataSampleRules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAffiliateSeller202410GetOpenCollaborationSampleRulesResponseDataSampleRules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


