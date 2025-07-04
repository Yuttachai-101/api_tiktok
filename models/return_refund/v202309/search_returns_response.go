/*
tiktok shop openapi

sdk for apis

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package return_refund_v202309

import (
    "encoding/json"
    "github.com/Yuttachai-101/api_tiktok/utils"
)

            // checks if the ReturnRefund202309SearchReturnsResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ReturnRefund202309SearchReturnsResponse{}

// ReturnRefund202309SearchReturnsResponse struct for ReturnRefund202309SearchReturnsResponse
type ReturnRefund202309SearchReturnsResponse struct {
    // The success or failure status code returned in API response.
    Code *int32 `json:"code,omitempty"`
    Data *ReturnRefund202309SearchReturnsResponseData `json:"data,omitempty"`
    // The success or failure messages returned in API response. Reasons of failure will be described in the message.
    Message *string `json:"message,omitempty"`
    // Request log.
    RequestId *string `json:"request_id,omitempty"`
}

// NewReturnRefund202309SearchReturnsResponse instantiates a new ReturnRefund202309SearchReturnsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReturnRefund202309SearchReturnsResponse() *ReturnRefund202309SearchReturnsResponse {
    this := ReturnRefund202309SearchReturnsResponse{}
    return &this
}

// NewReturnRefund202309SearchReturnsResponseWithDefaults instantiates a new ReturnRefund202309SearchReturnsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReturnRefund202309SearchReturnsResponseWithDefaults() *ReturnRefund202309SearchReturnsResponse {
    this := ReturnRefund202309SearchReturnsResponse{}
    return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ReturnRefund202309SearchReturnsResponse) GetCode() int32 {
    if o == nil || utils.IsNil(o.Code) {
        var ret int32
        return ret
    }
    return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnRefund202309SearchReturnsResponse) GetCodeOk() (*int32, bool) {
    if o == nil || utils.IsNil(o.Code) {
        return nil, false
    }
    return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ReturnRefund202309SearchReturnsResponse) HasCode() bool {
    if o != nil && !utils.IsNil(o.Code) {
        return true
    }

    return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *ReturnRefund202309SearchReturnsResponse) SetCode(v int32) {
    o.Code = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ReturnRefund202309SearchReturnsResponse) GetData() ReturnRefund202309SearchReturnsResponseData {
    if o == nil || utils.IsNil(o.Data) {
        var ret ReturnRefund202309SearchReturnsResponseData
        return ret
    }
    return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnRefund202309SearchReturnsResponse) GetDataOk() (*ReturnRefund202309SearchReturnsResponseData, bool) {
    if o == nil || utils.IsNil(o.Data) {
        return nil, false
    }
    return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ReturnRefund202309SearchReturnsResponse) HasData() bool {
    if o != nil && !utils.IsNil(o.Data) {
        return true
    }

    return false
}

// SetData gets a reference to the given ReturnRefund202309SearchReturnsResponseData and assigns it to the Data field.
func (o *ReturnRefund202309SearchReturnsResponse) SetData(v ReturnRefund202309SearchReturnsResponseData) {
    o.Data = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ReturnRefund202309SearchReturnsResponse) GetMessage() string {
    if o == nil || utils.IsNil(o.Message) {
        var ret string
        return ret
    }
    return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnRefund202309SearchReturnsResponse) GetMessageOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Message) {
        return nil, false
    }
    return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ReturnRefund202309SearchReturnsResponse) HasMessage() bool {
    if o != nil && !utils.IsNil(o.Message) {
        return true
    }

    return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ReturnRefund202309SearchReturnsResponse) SetMessage(v string) {
    o.Message = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *ReturnRefund202309SearchReturnsResponse) GetRequestId() string {
    if o == nil || utils.IsNil(o.RequestId) {
        var ret string
        return ret
    }
    return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnRefund202309SearchReturnsResponse) GetRequestIdOk() (*string, bool) {
    if o == nil || utils.IsNil(o.RequestId) {
        return nil, false
    }
    return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *ReturnRefund202309SearchReturnsResponse) HasRequestId() bool {
    if o != nil && !utils.IsNil(o.RequestId) {
        return true
    }

    return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *ReturnRefund202309SearchReturnsResponse) SetRequestId(v string) {
    o.RequestId = &v
}

func (o ReturnRefund202309SearchReturnsResponse) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o ReturnRefund202309SearchReturnsResponse) ToMap() (map[string]interface{}, error) {
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

type NullableReturnRefund202309SearchReturnsResponse struct {
	value *ReturnRefund202309SearchReturnsResponse
	isSet bool
}

func (v NullableReturnRefund202309SearchReturnsResponse) Get() *ReturnRefund202309SearchReturnsResponse {
	return v.value
}

func (v *NullableReturnRefund202309SearchReturnsResponse) Set(val *ReturnRefund202309SearchReturnsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReturnRefund202309SearchReturnsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReturnRefund202309SearchReturnsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReturnRefund202309SearchReturnsResponse(val *ReturnRefund202309SearchReturnsResponse) *NullableReturnRefund202309SearchReturnsResponse {
	return &NullableReturnRefund202309SearchReturnsResponse{value: val, isSet: true}
}

func (v NullableReturnRefund202309SearchReturnsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReturnRefund202309SearchReturnsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


