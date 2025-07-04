/*
tiktok shop openapi

sdk for apis

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package affiliate_seller_v202505

import (
    "encoding/json"
    "github.com/Yuttachai-101/api_tiktok/utils"
)

            // checks if the AffiliateSeller202505SellerSearchCreatoronMarketplaceResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AffiliateSeller202505SellerSearchCreatoronMarketplaceResponse{}

// AffiliateSeller202505SellerSearchCreatoronMarketplaceResponse struct for AffiliateSeller202505SellerSearchCreatoronMarketplaceResponse
type AffiliateSeller202505SellerSearchCreatoronMarketplaceResponse struct {
    // The success or failure status code returned in API response.
    Code *int32 `json:"code,omitempty"`
    Data *AffiliateSeller202505SellerSearchCreatoronMarketplaceResponseData `json:"data,omitempty"`
    // The success or failure messages returned in API response. Reasons of failure will be described in the message.
    Message *string `json:"message,omitempty"`
    // Request log.
    RequestId *string `json:"request_id,omitempty"`
}

// NewAffiliateSeller202505SellerSearchCreatoronMarketplaceResponse instantiates a new AffiliateSeller202505SellerSearchCreatoronMarketplaceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAffiliateSeller202505SellerSearchCreatoronMarketplaceResponse() *AffiliateSeller202505SellerSearchCreatoronMarketplaceResponse {
    this := AffiliateSeller202505SellerSearchCreatoronMarketplaceResponse{}
    return &this
}

// NewAffiliateSeller202505SellerSearchCreatoronMarketplaceResponseWithDefaults instantiates a new AffiliateSeller202505SellerSearchCreatoronMarketplaceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAffiliateSeller202505SellerSearchCreatoronMarketplaceResponseWithDefaults() *AffiliateSeller202505SellerSearchCreatoronMarketplaceResponse {
    this := AffiliateSeller202505SellerSearchCreatoronMarketplaceResponse{}
    return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *AffiliateSeller202505SellerSearchCreatoronMarketplaceResponse) GetCode() int32 {
    if o == nil || utils.IsNil(o.Code) {
        var ret int32
        return ret
    }
    return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateSeller202505SellerSearchCreatoronMarketplaceResponse) GetCodeOk() (*int32, bool) {
    if o == nil || utils.IsNil(o.Code) {
        return nil, false
    }
    return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *AffiliateSeller202505SellerSearchCreatoronMarketplaceResponse) HasCode() bool {
    if o != nil && !utils.IsNil(o.Code) {
        return true
    }

    return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *AffiliateSeller202505SellerSearchCreatoronMarketplaceResponse) SetCode(v int32) {
    o.Code = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AffiliateSeller202505SellerSearchCreatoronMarketplaceResponse) GetData() AffiliateSeller202505SellerSearchCreatoronMarketplaceResponseData {
    if o == nil || utils.IsNil(o.Data) {
        var ret AffiliateSeller202505SellerSearchCreatoronMarketplaceResponseData
        return ret
    }
    return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateSeller202505SellerSearchCreatoronMarketplaceResponse) GetDataOk() (*AffiliateSeller202505SellerSearchCreatoronMarketplaceResponseData, bool) {
    if o == nil || utils.IsNil(o.Data) {
        return nil, false
    }
    return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AffiliateSeller202505SellerSearchCreatoronMarketplaceResponse) HasData() bool {
    if o != nil && !utils.IsNil(o.Data) {
        return true
    }

    return false
}

// SetData gets a reference to the given AffiliateSeller202505SellerSearchCreatoronMarketplaceResponseData and assigns it to the Data field.
func (o *AffiliateSeller202505SellerSearchCreatoronMarketplaceResponse) SetData(v AffiliateSeller202505SellerSearchCreatoronMarketplaceResponseData) {
    o.Data = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *AffiliateSeller202505SellerSearchCreatoronMarketplaceResponse) GetMessage() string {
    if o == nil || utils.IsNil(o.Message) {
        var ret string
        return ret
    }
    return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateSeller202505SellerSearchCreatoronMarketplaceResponse) GetMessageOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Message) {
        return nil, false
    }
    return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *AffiliateSeller202505SellerSearchCreatoronMarketplaceResponse) HasMessage() bool {
    if o != nil && !utils.IsNil(o.Message) {
        return true
    }

    return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *AffiliateSeller202505SellerSearchCreatoronMarketplaceResponse) SetMessage(v string) {
    o.Message = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *AffiliateSeller202505SellerSearchCreatoronMarketplaceResponse) GetRequestId() string {
    if o == nil || utils.IsNil(o.RequestId) {
        var ret string
        return ret
    }
    return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateSeller202505SellerSearchCreatoronMarketplaceResponse) GetRequestIdOk() (*string, bool) {
    if o == nil || utils.IsNil(o.RequestId) {
        return nil, false
    }
    return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *AffiliateSeller202505SellerSearchCreatoronMarketplaceResponse) HasRequestId() bool {
    if o != nil && !utils.IsNil(o.RequestId) {
        return true
    }

    return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *AffiliateSeller202505SellerSearchCreatoronMarketplaceResponse) SetRequestId(v string) {
    o.RequestId = &v
}

func (o AffiliateSeller202505SellerSearchCreatoronMarketplaceResponse) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o AffiliateSeller202505SellerSearchCreatoronMarketplaceResponse) ToMap() (map[string]interface{}, error) {
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

type NullableAffiliateSeller202505SellerSearchCreatoronMarketplaceResponse struct {
	value *AffiliateSeller202505SellerSearchCreatoronMarketplaceResponse
	isSet bool
}

func (v NullableAffiliateSeller202505SellerSearchCreatoronMarketplaceResponse) Get() *AffiliateSeller202505SellerSearchCreatoronMarketplaceResponse {
	return v.value
}

func (v *NullableAffiliateSeller202505SellerSearchCreatoronMarketplaceResponse) Set(val *AffiliateSeller202505SellerSearchCreatoronMarketplaceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAffiliateSeller202505SellerSearchCreatoronMarketplaceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAffiliateSeller202505SellerSearchCreatoronMarketplaceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAffiliateSeller202505SellerSearchCreatoronMarketplaceResponse(val *AffiliateSeller202505SellerSearchCreatoronMarketplaceResponse) *NullableAffiliateSeller202505SellerSearchCreatoronMarketplaceResponse {
	return &NullableAffiliateSeller202505SellerSearchCreatoronMarketplaceResponse{value: val, isSet: true}
}

func (v NullableAffiliateSeller202505SellerSearchCreatoronMarketplaceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAffiliateSeller202505SellerSearchCreatoronMarketplaceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


