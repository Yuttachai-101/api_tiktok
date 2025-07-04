/*
tiktok shop openapi

sdk for apis

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package affiliate_seller_v202409

import (
    "encoding/json"
    "github.com/Yuttachai-101/api_tiktok/utils"
)

            // checks if the AffiliateSeller202409SellerSearchSampleApplicationsResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AffiliateSeller202409SellerSearchSampleApplicationsResponse{}

// AffiliateSeller202409SellerSearchSampleApplicationsResponse struct for AffiliateSeller202409SellerSearchSampleApplicationsResponse
type AffiliateSeller202409SellerSearchSampleApplicationsResponse struct {
    // The success or failure status code returned in API response.
    Code *int32 `json:"code,omitempty"`
    Data *AffiliateSeller202409SellerSearchSampleApplicationsResponseData `json:"data,omitempty"`
    // The success or failure messages returned in API response. Reasons of failure will be described in the message.
    Message *string `json:"message,omitempty"`
    // Request log.
    RequestId *string `json:"request_id,omitempty"`
}

// NewAffiliateSeller202409SellerSearchSampleApplicationsResponse instantiates a new AffiliateSeller202409SellerSearchSampleApplicationsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAffiliateSeller202409SellerSearchSampleApplicationsResponse() *AffiliateSeller202409SellerSearchSampleApplicationsResponse {
    this := AffiliateSeller202409SellerSearchSampleApplicationsResponse{}
    return &this
}

// NewAffiliateSeller202409SellerSearchSampleApplicationsResponseWithDefaults instantiates a new AffiliateSeller202409SellerSearchSampleApplicationsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAffiliateSeller202409SellerSearchSampleApplicationsResponseWithDefaults() *AffiliateSeller202409SellerSearchSampleApplicationsResponse {
    this := AffiliateSeller202409SellerSearchSampleApplicationsResponse{}
    return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *AffiliateSeller202409SellerSearchSampleApplicationsResponse) GetCode() int32 {
    if o == nil || utils.IsNil(o.Code) {
        var ret int32
        return ret
    }
    return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateSeller202409SellerSearchSampleApplicationsResponse) GetCodeOk() (*int32, bool) {
    if o == nil || utils.IsNil(o.Code) {
        return nil, false
    }
    return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *AffiliateSeller202409SellerSearchSampleApplicationsResponse) HasCode() bool {
    if o != nil && !utils.IsNil(o.Code) {
        return true
    }

    return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *AffiliateSeller202409SellerSearchSampleApplicationsResponse) SetCode(v int32) {
    o.Code = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AffiliateSeller202409SellerSearchSampleApplicationsResponse) GetData() AffiliateSeller202409SellerSearchSampleApplicationsResponseData {
    if o == nil || utils.IsNil(o.Data) {
        var ret AffiliateSeller202409SellerSearchSampleApplicationsResponseData
        return ret
    }
    return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateSeller202409SellerSearchSampleApplicationsResponse) GetDataOk() (*AffiliateSeller202409SellerSearchSampleApplicationsResponseData, bool) {
    if o == nil || utils.IsNil(o.Data) {
        return nil, false
    }
    return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AffiliateSeller202409SellerSearchSampleApplicationsResponse) HasData() bool {
    if o != nil && !utils.IsNil(o.Data) {
        return true
    }

    return false
}

// SetData gets a reference to the given AffiliateSeller202409SellerSearchSampleApplicationsResponseData and assigns it to the Data field.
func (o *AffiliateSeller202409SellerSearchSampleApplicationsResponse) SetData(v AffiliateSeller202409SellerSearchSampleApplicationsResponseData) {
    o.Data = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *AffiliateSeller202409SellerSearchSampleApplicationsResponse) GetMessage() string {
    if o == nil || utils.IsNil(o.Message) {
        var ret string
        return ret
    }
    return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateSeller202409SellerSearchSampleApplicationsResponse) GetMessageOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Message) {
        return nil, false
    }
    return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *AffiliateSeller202409SellerSearchSampleApplicationsResponse) HasMessage() bool {
    if o != nil && !utils.IsNil(o.Message) {
        return true
    }

    return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *AffiliateSeller202409SellerSearchSampleApplicationsResponse) SetMessage(v string) {
    o.Message = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *AffiliateSeller202409SellerSearchSampleApplicationsResponse) GetRequestId() string {
    if o == nil || utils.IsNil(o.RequestId) {
        var ret string
        return ret
    }
    return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateSeller202409SellerSearchSampleApplicationsResponse) GetRequestIdOk() (*string, bool) {
    if o == nil || utils.IsNil(o.RequestId) {
        return nil, false
    }
    return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *AffiliateSeller202409SellerSearchSampleApplicationsResponse) HasRequestId() bool {
    if o != nil && !utils.IsNil(o.RequestId) {
        return true
    }

    return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *AffiliateSeller202409SellerSearchSampleApplicationsResponse) SetRequestId(v string) {
    o.RequestId = &v
}

func (o AffiliateSeller202409SellerSearchSampleApplicationsResponse) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o AffiliateSeller202409SellerSearchSampleApplicationsResponse) ToMap() (map[string]interface{}, error) {
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

type NullableAffiliateSeller202409SellerSearchSampleApplicationsResponse struct {
	value *AffiliateSeller202409SellerSearchSampleApplicationsResponse
	isSet bool
}

func (v NullableAffiliateSeller202409SellerSearchSampleApplicationsResponse) Get() *AffiliateSeller202409SellerSearchSampleApplicationsResponse {
	return v.value
}

func (v *NullableAffiliateSeller202409SellerSearchSampleApplicationsResponse) Set(val *AffiliateSeller202409SellerSearchSampleApplicationsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAffiliateSeller202409SellerSearchSampleApplicationsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAffiliateSeller202409SellerSearchSampleApplicationsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAffiliateSeller202409SellerSearchSampleApplicationsResponse(val *AffiliateSeller202409SellerSearchSampleApplicationsResponse) *NullableAffiliateSeller202409SellerSearchSampleApplicationsResponse {
	return &NullableAffiliateSeller202409SellerSearchSampleApplicationsResponse{value: val, isSet: true}
}

func (v NullableAffiliateSeller202409SellerSearchSampleApplicationsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAffiliateSeller202409SellerSearchSampleApplicationsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


