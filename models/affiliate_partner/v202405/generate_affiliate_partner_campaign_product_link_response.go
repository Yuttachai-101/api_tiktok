/*
tiktok shop openapi

sdk for apis

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package affiliate_partner_v202405

import (
    "encoding/json"
    "github.com/Yuttachai-101/api_tiktok/utils"
)

            // checks if the AffiliatePartner202405GenerateAffiliatePartnerCampaignProductLinkResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AffiliatePartner202405GenerateAffiliatePartnerCampaignProductLinkResponse{}

// AffiliatePartner202405GenerateAffiliatePartnerCampaignProductLinkResponse struct for AffiliatePartner202405GenerateAffiliatePartnerCampaignProductLinkResponse
type AffiliatePartner202405GenerateAffiliatePartnerCampaignProductLinkResponse struct {
    // The success or failure status code returned in API response.
    Code *int32 `json:"code,omitempty"`
    Data *AffiliatePartner202405GenerateAffiliatePartnerCampaignProductLinkResponseData `json:"data,omitempty"`
    // The success or failure messages returned in API response. Reasons of failure will be described in the message.
    Message *string `json:"message,omitempty"`
    // Request log.
    RequestId *string `json:"request_id,omitempty"`
}

// NewAffiliatePartner202405GenerateAffiliatePartnerCampaignProductLinkResponse instantiates a new AffiliatePartner202405GenerateAffiliatePartnerCampaignProductLinkResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAffiliatePartner202405GenerateAffiliatePartnerCampaignProductLinkResponse() *AffiliatePartner202405GenerateAffiliatePartnerCampaignProductLinkResponse {
    this := AffiliatePartner202405GenerateAffiliatePartnerCampaignProductLinkResponse{}
    return &this
}

// NewAffiliatePartner202405GenerateAffiliatePartnerCampaignProductLinkResponseWithDefaults instantiates a new AffiliatePartner202405GenerateAffiliatePartnerCampaignProductLinkResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAffiliatePartner202405GenerateAffiliatePartnerCampaignProductLinkResponseWithDefaults() *AffiliatePartner202405GenerateAffiliatePartnerCampaignProductLinkResponse {
    this := AffiliatePartner202405GenerateAffiliatePartnerCampaignProductLinkResponse{}
    return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *AffiliatePartner202405GenerateAffiliatePartnerCampaignProductLinkResponse) GetCode() int32 {
    if o == nil || utils.IsNil(o.Code) {
        var ret int32
        return ret
    }
    return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliatePartner202405GenerateAffiliatePartnerCampaignProductLinkResponse) GetCodeOk() (*int32, bool) {
    if o == nil || utils.IsNil(o.Code) {
        return nil, false
    }
    return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *AffiliatePartner202405GenerateAffiliatePartnerCampaignProductLinkResponse) HasCode() bool {
    if o != nil && !utils.IsNil(o.Code) {
        return true
    }

    return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *AffiliatePartner202405GenerateAffiliatePartnerCampaignProductLinkResponse) SetCode(v int32) {
    o.Code = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AffiliatePartner202405GenerateAffiliatePartnerCampaignProductLinkResponse) GetData() AffiliatePartner202405GenerateAffiliatePartnerCampaignProductLinkResponseData {
    if o == nil || utils.IsNil(o.Data) {
        var ret AffiliatePartner202405GenerateAffiliatePartnerCampaignProductLinkResponseData
        return ret
    }
    return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliatePartner202405GenerateAffiliatePartnerCampaignProductLinkResponse) GetDataOk() (*AffiliatePartner202405GenerateAffiliatePartnerCampaignProductLinkResponseData, bool) {
    if o == nil || utils.IsNil(o.Data) {
        return nil, false
    }
    return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AffiliatePartner202405GenerateAffiliatePartnerCampaignProductLinkResponse) HasData() bool {
    if o != nil && !utils.IsNil(o.Data) {
        return true
    }

    return false
}

// SetData gets a reference to the given AffiliatePartner202405GenerateAffiliatePartnerCampaignProductLinkResponseData and assigns it to the Data field.
func (o *AffiliatePartner202405GenerateAffiliatePartnerCampaignProductLinkResponse) SetData(v AffiliatePartner202405GenerateAffiliatePartnerCampaignProductLinkResponseData) {
    o.Data = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *AffiliatePartner202405GenerateAffiliatePartnerCampaignProductLinkResponse) GetMessage() string {
    if o == nil || utils.IsNil(o.Message) {
        var ret string
        return ret
    }
    return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliatePartner202405GenerateAffiliatePartnerCampaignProductLinkResponse) GetMessageOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Message) {
        return nil, false
    }
    return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *AffiliatePartner202405GenerateAffiliatePartnerCampaignProductLinkResponse) HasMessage() bool {
    if o != nil && !utils.IsNil(o.Message) {
        return true
    }

    return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *AffiliatePartner202405GenerateAffiliatePartnerCampaignProductLinkResponse) SetMessage(v string) {
    o.Message = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *AffiliatePartner202405GenerateAffiliatePartnerCampaignProductLinkResponse) GetRequestId() string {
    if o == nil || utils.IsNil(o.RequestId) {
        var ret string
        return ret
    }
    return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliatePartner202405GenerateAffiliatePartnerCampaignProductLinkResponse) GetRequestIdOk() (*string, bool) {
    if o == nil || utils.IsNil(o.RequestId) {
        return nil, false
    }
    return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *AffiliatePartner202405GenerateAffiliatePartnerCampaignProductLinkResponse) HasRequestId() bool {
    if o != nil && !utils.IsNil(o.RequestId) {
        return true
    }

    return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *AffiliatePartner202405GenerateAffiliatePartnerCampaignProductLinkResponse) SetRequestId(v string) {
    o.RequestId = &v
}

func (o AffiliatePartner202405GenerateAffiliatePartnerCampaignProductLinkResponse) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o AffiliatePartner202405GenerateAffiliatePartnerCampaignProductLinkResponse) ToMap() (map[string]interface{}, error) {
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

type NullableAffiliatePartner202405GenerateAffiliatePartnerCampaignProductLinkResponse struct {
	value *AffiliatePartner202405GenerateAffiliatePartnerCampaignProductLinkResponse
	isSet bool
}

func (v NullableAffiliatePartner202405GenerateAffiliatePartnerCampaignProductLinkResponse) Get() *AffiliatePartner202405GenerateAffiliatePartnerCampaignProductLinkResponse {
	return v.value
}

func (v *NullableAffiliatePartner202405GenerateAffiliatePartnerCampaignProductLinkResponse) Set(val *AffiliatePartner202405GenerateAffiliatePartnerCampaignProductLinkResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAffiliatePartner202405GenerateAffiliatePartnerCampaignProductLinkResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAffiliatePartner202405GenerateAffiliatePartnerCampaignProductLinkResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAffiliatePartner202405GenerateAffiliatePartnerCampaignProductLinkResponse(val *AffiliatePartner202405GenerateAffiliatePartnerCampaignProductLinkResponse) *NullableAffiliatePartner202405GenerateAffiliatePartnerCampaignProductLinkResponse {
	return &NullableAffiliatePartner202405GenerateAffiliatePartnerCampaignProductLinkResponse{value: val, isSet: true}
}

func (v NullableAffiliatePartner202405GenerateAffiliatePartnerCampaignProductLinkResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAffiliatePartner202405GenerateAffiliatePartnerCampaignProductLinkResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


