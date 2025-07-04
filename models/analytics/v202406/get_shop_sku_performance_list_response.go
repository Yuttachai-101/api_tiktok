/*
tiktok shop openapi

sdk for apis

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package analytics_v202406

import (
    "encoding/json"
    "github.com/Yuttachai-101/api_tiktok/utils"
)

            // checks if the Analytics202406GetShopSKUPerformanceListResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Analytics202406GetShopSKUPerformanceListResponse{}

// Analytics202406GetShopSKUPerformanceListResponse struct for Analytics202406GetShopSKUPerformanceListResponse
type Analytics202406GetShopSKUPerformanceListResponse struct {
    // The success or failure status code returned in API response.
    Code *int32 `json:"code,omitempty"`
    Data *Analytics202406GetShopSKUPerformanceListResponseData `json:"data,omitempty"`
    // The success or failure messages returned in API response. Reasons of failure will be described in the message.
    Message *string `json:"message,omitempty"`
    // Request log.
    RequestId *string `json:"request_id,omitempty"`
}

// NewAnalytics202406GetShopSKUPerformanceListResponse instantiates a new Analytics202406GetShopSKUPerformanceListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnalytics202406GetShopSKUPerformanceListResponse() *Analytics202406GetShopSKUPerformanceListResponse {
    this := Analytics202406GetShopSKUPerformanceListResponse{}
    return &this
}

// NewAnalytics202406GetShopSKUPerformanceListResponseWithDefaults instantiates a new Analytics202406GetShopSKUPerformanceListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnalytics202406GetShopSKUPerformanceListResponseWithDefaults() *Analytics202406GetShopSKUPerformanceListResponse {
    this := Analytics202406GetShopSKUPerformanceListResponse{}
    return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *Analytics202406GetShopSKUPerformanceListResponse) GetCode() int32 {
    if o == nil || utils.IsNil(o.Code) {
        var ret int32
        return ret
    }
    return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Analytics202406GetShopSKUPerformanceListResponse) GetCodeOk() (*int32, bool) {
    if o == nil || utils.IsNil(o.Code) {
        return nil, false
    }
    return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *Analytics202406GetShopSKUPerformanceListResponse) HasCode() bool {
    if o != nil && !utils.IsNil(o.Code) {
        return true
    }

    return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *Analytics202406GetShopSKUPerformanceListResponse) SetCode(v int32) {
    o.Code = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *Analytics202406GetShopSKUPerformanceListResponse) GetData() Analytics202406GetShopSKUPerformanceListResponseData {
    if o == nil || utils.IsNil(o.Data) {
        var ret Analytics202406GetShopSKUPerformanceListResponseData
        return ret
    }
    return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Analytics202406GetShopSKUPerformanceListResponse) GetDataOk() (*Analytics202406GetShopSKUPerformanceListResponseData, bool) {
    if o == nil || utils.IsNil(o.Data) {
        return nil, false
    }
    return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *Analytics202406GetShopSKUPerformanceListResponse) HasData() bool {
    if o != nil && !utils.IsNil(o.Data) {
        return true
    }

    return false
}

// SetData gets a reference to the given Analytics202406GetShopSKUPerformanceListResponseData and assigns it to the Data field.
func (o *Analytics202406GetShopSKUPerformanceListResponse) SetData(v Analytics202406GetShopSKUPerformanceListResponseData) {
    o.Data = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *Analytics202406GetShopSKUPerformanceListResponse) GetMessage() string {
    if o == nil || utils.IsNil(o.Message) {
        var ret string
        return ret
    }
    return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Analytics202406GetShopSKUPerformanceListResponse) GetMessageOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Message) {
        return nil, false
    }
    return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *Analytics202406GetShopSKUPerformanceListResponse) HasMessage() bool {
    if o != nil && !utils.IsNil(o.Message) {
        return true
    }

    return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *Analytics202406GetShopSKUPerformanceListResponse) SetMessage(v string) {
    o.Message = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *Analytics202406GetShopSKUPerformanceListResponse) GetRequestId() string {
    if o == nil || utils.IsNil(o.RequestId) {
        var ret string
        return ret
    }
    return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Analytics202406GetShopSKUPerformanceListResponse) GetRequestIdOk() (*string, bool) {
    if o == nil || utils.IsNil(o.RequestId) {
        return nil, false
    }
    return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *Analytics202406GetShopSKUPerformanceListResponse) HasRequestId() bool {
    if o != nil && !utils.IsNil(o.RequestId) {
        return true
    }

    return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *Analytics202406GetShopSKUPerformanceListResponse) SetRequestId(v string) {
    o.RequestId = &v
}

func (o Analytics202406GetShopSKUPerformanceListResponse) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Analytics202406GetShopSKUPerformanceListResponse) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.Code) {
        toSerialize["code"] = o.Code
    }
    if !utils.IsNil(o.Data) {
        toSerialize["data"] = o.Data
    }
    if !utils.IsNil(o.Message) {
        toSerialize["message"] = o.Message
    }
    if !utils.IsNil(o.RequestId) {
        toSerialize["request_id"] = o.RequestId
    }
    return toSerialize, nil
}

type NullableAnalytics202406GetShopSKUPerformanceListResponse struct {
	value *Analytics202406GetShopSKUPerformanceListResponse
	isSet bool
}

func (v NullableAnalytics202406GetShopSKUPerformanceListResponse) Get() *Analytics202406GetShopSKUPerformanceListResponse {
	return v.value
}

func (v *NullableAnalytics202406GetShopSKUPerformanceListResponse) Set(val *Analytics202406GetShopSKUPerformanceListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAnalytics202406GetShopSKUPerformanceListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAnalytics202406GetShopSKUPerformanceListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnalytics202406GetShopSKUPerformanceListResponse(val *Analytics202406GetShopSKUPerformanceListResponse) *NullableAnalytics202406GetShopSKUPerformanceListResponse {
	return &NullableAnalytics202406GetShopSKUPerformanceListResponse{value: val, isSet: true}
}

func (v NullableAnalytics202406GetShopSKUPerformanceListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnalytics202406GetShopSKUPerformanceListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


