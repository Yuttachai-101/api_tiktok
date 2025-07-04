/*
tiktok shop openapi

sdk for apis

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package analytics_v202405

import (
    "encoding/json"
    "github.com/Yuttachai-101/api_tiktok/utils"
)

            // checks if the Analytics202405GetShopProductPerformanceListResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Analytics202405GetShopProductPerformanceListResponse{}

// Analytics202405GetShopProductPerformanceListResponse struct for Analytics202405GetShopProductPerformanceListResponse
type Analytics202405GetShopProductPerformanceListResponse struct {
    // The success or failure status code returned in API response.
    Code *int32 `json:"code,omitempty"`
    Data *Analytics202405GetShopProductPerformanceListResponseData `json:"data,omitempty"`
    // The success or failure messages returned in API response. Reasons of failure will be described in the message.
    Message *string `json:"message,omitempty"`
    // Request log.
    RequestId *string `json:"request_id,omitempty"`
}

// NewAnalytics202405GetShopProductPerformanceListResponse instantiates a new Analytics202405GetShopProductPerformanceListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnalytics202405GetShopProductPerformanceListResponse() *Analytics202405GetShopProductPerformanceListResponse {
    this := Analytics202405GetShopProductPerformanceListResponse{}
    return &this
}

// NewAnalytics202405GetShopProductPerformanceListResponseWithDefaults instantiates a new Analytics202405GetShopProductPerformanceListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnalytics202405GetShopProductPerformanceListResponseWithDefaults() *Analytics202405GetShopProductPerformanceListResponse {
    this := Analytics202405GetShopProductPerformanceListResponse{}
    return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *Analytics202405GetShopProductPerformanceListResponse) GetCode() int32 {
    if o == nil || utils.IsNil(o.Code) {
        var ret int32
        return ret
    }
    return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Analytics202405GetShopProductPerformanceListResponse) GetCodeOk() (*int32, bool) {
    if o == nil || utils.IsNil(o.Code) {
        return nil, false
    }
    return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *Analytics202405GetShopProductPerformanceListResponse) HasCode() bool {
    if o != nil && !utils.IsNil(o.Code) {
        return true
    }

    return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *Analytics202405GetShopProductPerformanceListResponse) SetCode(v int32) {
    o.Code = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *Analytics202405GetShopProductPerformanceListResponse) GetData() Analytics202405GetShopProductPerformanceListResponseData {
    if o == nil || utils.IsNil(o.Data) {
        var ret Analytics202405GetShopProductPerformanceListResponseData
        return ret
    }
    return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Analytics202405GetShopProductPerformanceListResponse) GetDataOk() (*Analytics202405GetShopProductPerformanceListResponseData, bool) {
    if o == nil || utils.IsNil(o.Data) {
        return nil, false
    }
    return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *Analytics202405GetShopProductPerformanceListResponse) HasData() bool {
    if o != nil && !utils.IsNil(o.Data) {
        return true
    }

    return false
}

// SetData gets a reference to the given Analytics202405GetShopProductPerformanceListResponseData and assigns it to the Data field.
func (o *Analytics202405GetShopProductPerformanceListResponse) SetData(v Analytics202405GetShopProductPerformanceListResponseData) {
    o.Data = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *Analytics202405GetShopProductPerformanceListResponse) GetMessage() string {
    if o == nil || utils.IsNil(o.Message) {
        var ret string
        return ret
    }
    return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Analytics202405GetShopProductPerformanceListResponse) GetMessageOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Message) {
        return nil, false
    }
    return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *Analytics202405GetShopProductPerformanceListResponse) HasMessage() bool {
    if o != nil && !utils.IsNil(o.Message) {
        return true
    }

    return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *Analytics202405GetShopProductPerformanceListResponse) SetMessage(v string) {
    o.Message = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *Analytics202405GetShopProductPerformanceListResponse) GetRequestId() string {
    if o == nil || utils.IsNil(o.RequestId) {
        var ret string
        return ret
    }
    return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Analytics202405GetShopProductPerformanceListResponse) GetRequestIdOk() (*string, bool) {
    if o == nil || utils.IsNil(o.RequestId) {
        return nil, false
    }
    return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *Analytics202405GetShopProductPerformanceListResponse) HasRequestId() bool {
    if o != nil && !utils.IsNil(o.RequestId) {
        return true
    }

    return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *Analytics202405GetShopProductPerformanceListResponse) SetRequestId(v string) {
    o.RequestId = &v
}

func (o Analytics202405GetShopProductPerformanceListResponse) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Analytics202405GetShopProductPerformanceListResponse) ToMap() (map[string]interface{}, error) {
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

type NullableAnalytics202405GetShopProductPerformanceListResponse struct {
	value *Analytics202405GetShopProductPerformanceListResponse
	isSet bool
}

func (v NullableAnalytics202405GetShopProductPerformanceListResponse) Get() *Analytics202405GetShopProductPerformanceListResponse {
	return v.value
}

func (v *NullableAnalytics202405GetShopProductPerformanceListResponse) Set(val *Analytics202405GetShopProductPerformanceListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAnalytics202405GetShopProductPerformanceListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAnalytics202405GetShopProductPerformanceListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnalytics202405GetShopProductPerformanceListResponse(val *Analytics202405GetShopProductPerformanceListResponse) *NullableAnalytics202405GetShopProductPerformanceListResponse {
	return &NullableAnalytics202405GetShopProductPerformanceListResponse{value: val, isSet: true}
}

func (v NullableAnalytics202405GetShopProductPerformanceListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnalytics202405GetShopProductPerformanceListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


