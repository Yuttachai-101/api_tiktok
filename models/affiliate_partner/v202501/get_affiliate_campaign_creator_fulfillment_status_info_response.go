/*
tiktok shop openapi

sdk for apis

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package affiliate_partner_v202501

import (
    "encoding/json"
    "github.com/Yuttachai-101/api_tiktok/utils"
)

            // checks if the AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusInfoResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusInfoResponse{}

// AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusInfoResponse struct for AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusInfoResponse
type AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusInfoResponse struct {
    // The success or failure status code returned in API response.
    Code *int32 `json:"code,omitempty"`
    Data *AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusInfoResponseData `json:"data,omitempty"`
    // The success or failure messages returned in API response. Reasons of failure will be described in the message.
    Message *string `json:"message,omitempty"`
    // Request log.
    RequestId *string `json:"request_id,omitempty"`
}

// NewAffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusInfoResponse instantiates a new AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusInfoResponse() *AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusInfoResponse {
    this := AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusInfoResponse{}
    return &this
}

// NewAffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusInfoResponseWithDefaults instantiates a new AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusInfoResponseWithDefaults() *AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusInfoResponse {
    this := AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusInfoResponse{}
    return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusInfoResponse) GetCode() int32 {
    if o == nil || utils.IsNil(o.Code) {
        var ret int32
        return ret
    }
    return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusInfoResponse) GetCodeOk() (*int32, bool) {
    if o == nil || utils.IsNil(o.Code) {
        return nil, false
    }
    return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusInfoResponse) HasCode() bool {
    if o != nil && !utils.IsNil(o.Code) {
        return true
    }

    return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusInfoResponse) SetCode(v int32) {
    o.Code = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusInfoResponse) GetData() AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusInfoResponseData {
    if o == nil || utils.IsNil(o.Data) {
        var ret AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusInfoResponseData
        return ret
    }
    return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusInfoResponse) GetDataOk() (*AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusInfoResponseData, bool) {
    if o == nil || utils.IsNil(o.Data) {
        return nil, false
    }
    return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusInfoResponse) HasData() bool {
    if o != nil && !utils.IsNil(o.Data) {
        return true
    }

    return false
}

// SetData gets a reference to the given AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusInfoResponseData and assigns it to the Data field.
func (o *AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusInfoResponse) SetData(v AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusInfoResponseData) {
    o.Data = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusInfoResponse) GetMessage() string {
    if o == nil || utils.IsNil(o.Message) {
        var ret string
        return ret
    }
    return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusInfoResponse) GetMessageOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Message) {
        return nil, false
    }
    return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusInfoResponse) HasMessage() bool {
    if o != nil && !utils.IsNil(o.Message) {
        return true
    }

    return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusInfoResponse) SetMessage(v string) {
    o.Message = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusInfoResponse) GetRequestId() string {
    if o == nil || utils.IsNil(o.RequestId) {
        var ret string
        return ret
    }
    return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusInfoResponse) GetRequestIdOk() (*string, bool) {
    if o == nil || utils.IsNil(o.RequestId) {
        return nil, false
    }
    return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusInfoResponse) HasRequestId() bool {
    if o != nil && !utils.IsNil(o.RequestId) {
        return true
    }

    return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusInfoResponse) SetRequestId(v string) {
    o.RequestId = &v
}

func (o AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusInfoResponse) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusInfoResponse) ToMap() (map[string]interface{}, error) {
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

type NullableAffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusInfoResponse struct {
	value *AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusInfoResponse
	isSet bool
}

func (v NullableAffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusInfoResponse) Get() *AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusInfoResponse {
	return v.value
}

func (v *NullableAffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusInfoResponse) Set(val *AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusInfoResponse(val *AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusInfoResponse) *NullableAffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusInfoResponse {
	return &NullableAffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusInfoResponse{value: val, isSet: true}
}

func (v NullableAffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


