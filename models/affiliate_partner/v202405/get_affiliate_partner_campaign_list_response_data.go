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

            // checks if the AffiliatePartner202405GetAffiliatePartnerCampaignListResponseData type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AffiliatePartner202405GetAffiliatePartnerCampaignListResponseData{}

// AffiliatePartner202405GetAffiliatePartnerCampaignListResponseData struct for AffiliatePartner202405GetAffiliatePartnerCampaignListResponseData
type AffiliatePartner202405GetAffiliatePartnerCampaignListResponseData struct {
    // A list of campaigns. 
    Campaigns []AffiliatePartner202405GetAffiliatePartnerCampaignListResponseDataCampaigns `json:"campaigns,omitempty"`
    // An opaque token used to retrieve the next page of a paginated result set. 
    NextPageToken *string `json:"next_page_token,omitempty"`
    // The total number of campaigns in the list.
    TotalCount *int64 `json:"total_count,omitempty"`
}

// NewAffiliatePartner202405GetAffiliatePartnerCampaignListResponseData instantiates a new AffiliatePartner202405GetAffiliatePartnerCampaignListResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAffiliatePartner202405GetAffiliatePartnerCampaignListResponseData() *AffiliatePartner202405GetAffiliatePartnerCampaignListResponseData {
    this := AffiliatePartner202405GetAffiliatePartnerCampaignListResponseData{}
    return &this
}

// NewAffiliatePartner202405GetAffiliatePartnerCampaignListResponseDataWithDefaults instantiates a new AffiliatePartner202405GetAffiliatePartnerCampaignListResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAffiliatePartner202405GetAffiliatePartnerCampaignListResponseDataWithDefaults() *AffiliatePartner202405GetAffiliatePartnerCampaignListResponseData {
    this := AffiliatePartner202405GetAffiliatePartnerCampaignListResponseData{}
    return &this
}

// GetCampaigns returns the Campaigns field value if set, zero value otherwise.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignListResponseData) GetCampaigns() []AffiliatePartner202405GetAffiliatePartnerCampaignListResponseDataCampaigns {
    if o == nil || utils.IsNil(o.Campaigns) {
        var ret []AffiliatePartner202405GetAffiliatePartnerCampaignListResponseDataCampaigns
        return ret
    }
    return o.Campaigns
}

// GetCampaignsOk returns a tuple with the Campaigns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignListResponseData) GetCampaignsOk() ([]AffiliatePartner202405GetAffiliatePartnerCampaignListResponseDataCampaigns, bool) {
    if o == nil || utils.IsNil(o.Campaigns) {
        return nil, false
    }
    return o.Campaigns, true
}

// HasCampaigns returns a boolean if a field has been set.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignListResponseData) HasCampaigns() bool {
    if o != nil && !utils.IsNil(o.Campaigns) {
        return true
    }

    return false
}

// SetCampaigns gets a reference to the given []AffiliatePartner202405GetAffiliatePartnerCampaignListResponseDataCampaigns and assigns it to the Campaigns field.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignListResponseData) SetCampaigns(v []AffiliatePartner202405GetAffiliatePartnerCampaignListResponseDataCampaigns) {
    o.Campaigns = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignListResponseData) GetNextPageToken() string {
    if o == nil || utils.IsNil(o.NextPageToken) {
        var ret string
        return ret
    }
    return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignListResponseData) GetNextPageTokenOk() (*string, bool) {
    if o == nil || utils.IsNil(o.NextPageToken) {
        return nil, false
    }
    return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignListResponseData) HasNextPageToken() bool {
    if o != nil && !utils.IsNil(o.NextPageToken) {
        return true
    }

    return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignListResponseData) SetNextPageToken(v string) {
    o.NextPageToken = &v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignListResponseData) GetTotalCount() int64 {
    if o == nil || utils.IsNil(o.TotalCount) {
        var ret int64
        return ret
    }
    return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignListResponseData) GetTotalCountOk() (*int64, bool) {
    if o == nil || utils.IsNil(o.TotalCount) {
        return nil, false
    }
    return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignListResponseData) HasTotalCount() bool {
    if o != nil && !utils.IsNil(o.TotalCount) {
        return true
    }

    return false
}

// SetTotalCount gets a reference to the given int64 and assigns it to the TotalCount field.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignListResponseData) SetTotalCount(v int64) {
    o.TotalCount = &v
}

func (o AffiliatePartner202405GetAffiliatePartnerCampaignListResponseData) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o AffiliatePartner202405GetAffiliatePartnerCampaignListResponseData) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.Campaigns) {
        toSerialize["campaigns"] = o.Campaigns
    }
    if !utils.IsNil(o.NextPageToken) {
        toSerialize["next_page_token"] = o.NextPageToken
    }
    if !utils.IsNil(o.TotalCount) {
        toSerialize["total_count"] = o.TotalCount
    }
    return toSerialize, nil
}

type NullableAffiliatePartner202405GetAffiliatePartnerCampaignListResponseData struct {
	value *AffiliatePartner202405GetAffiliatePartnerCampaignListResponseData
	isSet bool
}

func (v NullableAffiliatePartner202405GetAffiliatePartnerCampaignListResponseData) Get() *AffiliatePartner202405GetAffiliatePartnerCampaignListResponseData {
	return v.value
}

func (v *NullableAffiliatePartner202405GetAffiliatePartnerCampaignListResponseData) Set(val *AffiliatePartner202405GetAffiliatePartnerCampaignListResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableAffiliatePartner202405GetAffiliatePartnerCampaignListResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableAffiliatePartner202405GetAffiliatePartnerCampaignListResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAffiliatePartner202405GetAffiliatePartnerCampaignListResponseData(val *AffiliatePartner202405GetAffiliatePartnerCampaignListResponseData) *NullableAffiliatePartner202405GetAffiliatePartnerCampaignListResponseData {
	return &NullableAffiliatePartner202405GetAffiliatePartnerCampaignListResponseData{value: val, isSet: true}
}

func (v NullableAffiliatePartner202405GetAffiliatePartnerCampaignListResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAffiliatePartner202405GetAffiliatePartnerCampaignListResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


