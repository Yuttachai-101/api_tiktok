/*
tiktok shop openapi

sdk for apis

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package affiliate_v202309

import (
    "encoding/json"
    "github.com/Yuttachai-101/api_tiktok/utils"
)

            // checks if the Affiliate202309GetLiveRoomInfoResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Affiliate202309GetLiveRoomInfoResponse{}

// Affiliate202309GetLiveRoomInfoResponse struct for Affiliate202309GetLiveRoomInfoResponse
type Affiliate202309GetLiveRoomInfoResponse struct {
    // The success or failure status code returned in API response.
    Code *int32 `json:"code,omitempty"`
    Data *Affiliate202309GetLiveRoomInfoResponseData `json:"data,omitempty"`
    // The success or failure messages returned in API response. Reasons of failure will be described in the message.
    Message *string `json:"message,omitempty"`
    // Request log.
    RequestId *string `json:"request_id,omitempty"`
}

// NewAffiliate202309GetLiveRoomInfoResponse instantiates a new Affiliate202309GetLiveRoomInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAffiliate202309GetLiveRoomInfoResponse() *Affiliate202309GetLiveRoomInfoResponse {
    this := Affiliate202309GetLiveRoomInfoResponse{}
    return &this
}

// NewAffiliate202309GetLiveRoomInfoResponseWithDefaults instantiates a new Affiliate202309GetLiveRoomInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAffiliate202309GetLiveRoomInfoResponseWithDefaults() *Affiliate202309GetLiveRoomInfoResponse {
    this := Affiliate202309GetLiveRoomInfoResponse{}
    return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *Affiliate202309GetLiveRoomInfoResponse) GetCode() int32 {
    if o == nil || utils.IsNil(o.Code) {
        var ret int32
        return ret
    }
    return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Affiliate202309GetLiveRoomInfoResponse) GetCodeOk() (*int32, bool) {
    if o == nil || utils.IsNil(o.Code) {
        return nil, false
    }
    return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *Affiliate202309GetLiveRoomInfoResponse) HasCode() bool {
    if o != nil && !utils.IsNil(o.Code) {
        return true
    }

    return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *Affiliate202309GetLiveRoomInfoResponse) SetCode(v int32) {
    o.Code = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *Affiliate202309GetLiveRoomInfoResponse) GetData() Affiliate202309GetLiveRoomInfoResponseData {
    if o == nil || utils.IsNil(o.Data) {
        var ret Affiliate202309GetLiveRoomInfoResponseData
        return ret
    }
    return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Affiliate202309GetLiveRoomInfoResponse) GetDataOk() (*Affiliate202309GetLiveRoomInfoResponseData, bool) {
    if o == nil || utils.IsNil(o.Data) {
        return nil, false
    }
    return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *Affiliate202309GetLiveRoomInfoResponse) HasData() bool {
    if o != nil && !utils.IsNil(o.Data) {
        return true
    }

    return false
}

// SetData gets a reference to the given Affiliate202309GetLiveRoomInfoResponseData and assigns it to the Data field.
func (o *Affiliate202309GetLiveRoomInfoResponse) SetData(v Affiliate202309GetLiveRoomInfoResponseData) {
    o.Data = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *Affiliate202309GetLiveRoomInfoResponse) GetMessage() string {
    if o == nil || utils.IsNil(o.Message) {
        var ret string
        return ret
    }
    return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Affiliate202309GetLiveRoomInfoResponse) GetMessageOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Message) {
        return nil, false
    }
    return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *Affiliate202309GetLiveRoomInfoResponse) HasMessage() bool {
    if o != nil && !utils.IsNil(o.Message) {
        return true
    }

    return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *Affiliate202309GetLiveRoomInfoResponse) SetMessage(v string) {
    o.Message = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *Affiliate202309GetLiveRoomInfoResponse) GetRequestId() string {
    if o == nil || utils.IsNil(o.RequestId) {
        var ret string
        return ret
    }
    return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Affiliate202309GetLiveRoomInfoResponse) GetRequestIdOk() (*string, bool) {
    if o == nil || utils.IsNil(o.RequestId) {
        return nil, false
    }
    return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *Affiliate202309GetLiveRoomInfoResponse) HasRequestId() bool {
    if o != nil && !utils.IsNil(o.RequestId) {
        return true
    }

    return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *Affiliate202309GetLiveRoomInfoResponse) SetRequestId(v string) {
    o.RequestId = &v
}

func (o Affiliate202309GetLiveRoomInfoResponse) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Affiliate202309GetLiveRoomInfoResponse) ToMap() (map[string]interface{}, error) {
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

type NullableAffiliate202309GetLiveRoomInfoResponse struct {
	value *Affiliate202309GetLiveRoomInfoResponse
	isSet bool
}

func (v NullableAffiliate202309GetLiveRoomInfoResponse) Get() *Affiliate202309GetLiveRoomInfoResponse {
	return v.value
}

func (v *NullableAffiliate202309GetLiveRoomInfoResponse) Set(val *Affiliate202309GetLiveRoomInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAffiliate202309GetLiveRoomInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAffiliate202309GetLiveRoomInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAffiliate202309GetLiveRoomInfoResponse(val *Affiliate202309GetLiveRoomInfoResponse) *NullableAffiliate202309GetLiveRoomInfoResponse {
	return &NullableAffiliate202309GetLiveRoomInfoResponse{value: val, isSet: true}
}

func (v NullableAffiliate202309GetLiveRoomInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAffiliate202309GetLiveRoomInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


