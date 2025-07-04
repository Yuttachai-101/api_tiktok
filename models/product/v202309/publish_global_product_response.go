/*
tiktok shop openapi

sdk for apis

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package product_v202309

import (
    "encoding/json"
    "github.com/Yuttachai-101/api_tiktok/utils"
)

            // checks if the Product202309PublishGlobalProductResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Product202309PublishGlobalProductResponse{}

// Product202309PublishGlobalProductResponse struct for Product202309PublishGlobalProductResponse
type Product202309PublishGlobalProductResponse struct {
    // The success or failure status code returned in API response.
    Code *int32 `json:"code,omitempty"`
    Data *Product202309PublishGlobalProductResponseData `json:"data,omitempty"`
    // The success or failure messages returned in API response. Reasons of failure will be described in the message.
    Message *string `json:"message,omitempty"`
    // Request log.
    RequestId *string `json:"request_id,omitempty"`
}

// NewProduct202309PublishGlobalProductResponse instantiates a new Product202309PublishGlobalProductResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProduct202309PublishGlobalProductResponse() *Product202309PublishGlobalProductResponse {
    this := Product202309PublishGlobalProductResponse{}
    return &this
}

// NewProduct202309PublishGlobalProductResponseWithDefaults instantiates a new Product202309PublishGlobalProductResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProduct202309PublishGlobalProductResponseWithDefaults() *Product202309PublishGlobalProductResponse {
    this := Product202309PublishGlobalProductResponse{}
    return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *Product202309PublishGlobalProductResponse) GetCode() int32 {
    if o == nil || utils.IsNil(o.Code) {
        var ret int32
        return ret
    }
    return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309PublishGlobalProductResponse) GetCodeOk() (*int32, bool) {
    if o == nil || utils.IsNil(o.Code) {
        return nil, false
    }
    return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *Product202309PublishGlobalProductResponse) HasCode() bool {
    if o != nil && !utils.IsNil(o.Code) {
        return true
    }

    return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *Product202309PublishGlobalProductResponse) SetCode(v int32) {
    o.Code = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *Product202309PublishGlobalProductResponse) GetData() Product202309PublishGlobalProductResponseData {
    if o == nil || utils.IsNil(o.Data) {
        var ret Product202309PublishGlobalProductResponseData
        return ret
    }
    return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309PublishGlobalProductResponse) GetDataOk() (*Product202309PublishGlobalProductResponseData, bool) {
    if o == nil || utils.IsNil(o.Data) {
        return nil, false
    }
    return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *Product202309PublishGlobalProductResponse) HasData() bool {
    if o != nil && !utils.IsNil(o.Data) {
        return true
    }

    return false
}

// SetData gets a reference to the given Product202309PublishGlobalProductResponseData and assigns it to the Data field.
func (o *Product202309PublishGlobalProductResponse) SetData(v Product202309PublishGlobalProductResponseData) {
    o.Data = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *Product202309PublishGlobalProductResponse) GetMessage() string {
    if o == nil || utils.IsNil(o.Message) {
        var ret string
        return ret
    }
    return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309PublishGlobalProductResponse) GetMessageOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Message) {
        return nil, false
    }
    return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *Product202309PublishGlobalProductResponse) HasMessage() bool {
    if o != nil && !utils.IsNil(o.Message) {
        return true
    }

    return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *Product202309PublishGlobalProductResponse) SetMessage(v string) {
    o.Message = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *Product202309PublishGlobalProductResponse) GetRequestId() string {
    if o == nil || utils.IsNil(o.RequestId) {
        var ret string
        return ret
    }
    return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309PublishGlobalProductResponse) GetRequestIdOk() (*string, bool) {
    if o == nil || utils.IsNil(o.RequestId) {
        return nil, false
    }
    return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *Product202309PublishGlobalProductResponse) HasRequestId() bool {
    if o != nil && !utils.IsNil(o.RequestId) {
        return true
    }

    return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *Product202309PublishGlobalProductResponse) SetRequestId(v string) {
    o.RequestId = &v
}

func (o Product202309PublishGlobalProductResponse) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Product202309PublishGlobalProductResponse) ToMap() (map[string]interface{}, error) {
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

type NullableProduct202309PublishGlobalProductResponse struct {
	value *Product202309PublishGlobalProductResponse
	isSet bool
}

func (v NullableProduct202309PublishGlobalProductResponse) Get() *Product202309PublishGlobalProductResponse {
	return v.value
}

func (v *NullableProduct202309PublishGlobalProductResponse) Set(val *Product202309PublishGlobalProductResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProduct202309PublishGlobalProductResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProduct202309PublishGlobalProductResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProduct202309PublishGlobalProductResponse(val *Product202309PublishGlobalProductResponse) *NullableProduct202309PublishGlobalProductResponse {
	return &NullableProduct202309PublishGlobalProductResponse{value: val, isSet: true}
}

func (v NullableProduct202309PublishGlobalProductResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProduct202309PublishGlobalProductResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


