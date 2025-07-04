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

            // checks if the AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorData type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorData{}

// AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorData struct for AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorData
type AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorData struct {
    // GMV, the total payment amount corresponding to `actual_order_num`. 
    ActualAmount *string `json:"actual_amount,omitempty"`
    // The actual number of paid orders. 
    ActualOrderNum *string `json:"actual_order_num,omitempty"`
    // The partner commission amount corresponding to actual_order_num. 
    ActualPartnerCommission *string `json:"actual_partner_commission,omitempty"`
    // The total number of creators collaborating on the campaign.
    CollaboratedCreatorsNum *string `json:"collaborated_creators_num,omitempty"`
    // The number of creators credited with at least one product sale.
    CreatorSalesNum *string `json:"creator_sales_num,omitempty"`
    // The total payment amount corresponding to `paid_order_num`. 
    EstimatedAmount *string `json:"estimated_amount,omitempty"`
    // The partner commission amount corresponding to paid_order_num. 
    EstimatedPartnerCommission *string `json:"estimated_partner_commission,omitempty"`
    // The total number of paid orders. `paid_order_num` = `actual_order_num` + {number of paid but returned orders} 
    PaidOrderNum *string `json:"paid_order_num,omitempty"`
    // The number of creators involved in the partner plan. 
    PromotedCreatorNum *string `json:"promoted_creator_num,omitempty"`
    // The total number of creators that applied for a sample.
    SampleRequestedCreatorNum *string `json:"sample_requested_creator_num,omitempty"`
}

// NewAffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorData instantiates a new AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorData() *AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorData {
    this := AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorData{}
    return &this
}

// NewAffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorDataWithDefaults instantiates a new AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorDataWithDefaults() *AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorData {
    this := AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorData{}
    return &this
}

// GetActualAmount returns the ActualAmount field value if set, zero value otherwise.
func (o *AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorData) GetActualAmount() string {
    if o == nil || utils.IsNil(o.ActualAmount) {
        var ret string
        return ret
    }
    return *o.ActualAmount
}

// GetActualAmountOk returns a tuple with the ActualAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorData) GetActualAmountOk() (*string, bool) {
    if o == nil || utils.IsNil(o.ActualAmount) {
        return nil, false
    }
    return o.ActualAmount, true
}

// HasActualAmount returns a boolean if a field has been set.
func (o *AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorData) HasActualAmount() bool {
    if o != nil && !utils.IsNil(o.ActualAmount) {
        return true
    }

    return false
}

// SetActualAmount gets a reference to the given string and assigns it to the ActualAmount field.
func (o *AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorData) SetActualAmount(v string) {
    o.ActualAmount = &v
}

// GetActualOrderNum returns the ActualOrderNum field value if set, zero value otherwise.
func (o *AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorData) GetActualOrderNum() string {
    if o == nil || utils.IsNil(o.ActualOrderNum) {
        var ret string
        return ret
    }
    return *o.ActualOrderNum
}

// GetActualOrderNumOk returns a tuple with the ActualOrderNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorData) GetActualOrderNumOk() (*string, bool) {
    if o == nil || utils.IsNil(o.ActualOrderNum) {
        return nil, false
    }
    return o.ActualOrderNum, true
}

// HasActualOrderNum returns a boolean if a field has been set.
func (o *AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorData) HasActualOrderNum() bool {
    if o != nil && !utils.IsNil(o.ActualOrderNum) {
        return true
    }

    return false
}

// SetActualOrderNum gets a reference to the given string and assigns it to the ActualOrderNum field.
func (o *AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorData) SetActualOrderNum(v string) {
    o.ActualOrderNum = &v
}

// GetActualPartnerCommission returns the ActualPartnerCommission field value if set, zero value otherwise.
func (o *AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorData) GetActualPartnerCommission() string {
    if o == nil || utils.IsNil(o.ActualPartnerCommission) {
        var ret string
        return ret
    }
    return *o.ActualPartnerCommission
}

// GetActualPartnerCommissionOk returns a tuple with the ActualPartnerCommission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorData) GetActualPartnerCommissionOk() (*string, bool) {
    if o == nil || utils.IsNil(o.ActualPartnerCommission) {
        return nil, false
    }
    return o.ActualPartnerCommission, true
}

// HasActualPartnerCommission returns a boolean if a field has been set.
func (o *AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorData) HasActualPartnerCommission() bool {
    if o != nil && !utils.IsNil(o.ActualPartnerCommission) {
        return true
    }

    return false
}

// SetActualPartnerCommission gets a reference to the given string and assigns it to the ActualPartnerCommission field.
func (o *AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorData) SetActualPartnerCommission(v string) {
    o.ActualPartnerCommission = &v
}

// GetCollaboratedCreatorsNum returns the CollaboratedCreatorsNum field value if set, zero value otherwise.
func (o *AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorData) GetCollaboratedCreatorsNum() string {
    if o == nil || utils.IsNil(o.CollaboratedCreatorsNum) {
        var ret string
        return ret
    }
    return *o.CollaboratedCreatorsNum
}

// GetCollaboratedCreatorsNumOk returns a tuple with the CollaboratedCreatorsNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorData) GetCollaboratedCreatorsNumOk() (*string, bool) {
    if o == nil || utils.IsNil(o.CollaboratedCreatorsNum) {
        return nil, false
    }
    return o.CollaboratedCreatorsNum, true
}

// HasCollaboratedCreatorsNum returns a boolean if a field has been set.
func (o *AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorData) HasCollaboratedCreatorsNum() bool {
    if o != nil && !utils.IsNil(o.CollaboratedCreatorsNum) {
        return true
    }

    return false
}

// SetCollaboratedCreatorsNum gets a reference to the given string and assigns it to the CollaboratedCreatorsNum field.
func (o *AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorData) SetCollaboratedCreatorsNum(v string) {
    o.CollaboratedCreatorsNum = &v
}

// GetCreatorSalesNum returns the CreatorSalesNum field value if set, zero value otherwise.
func (o *AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorData) GetCreatorSalesNum() string {
    if o == nil || utils.IsNil(o.CreatorSalesNum) {
        var ret string
        return ret
    }
    return *o.CreatorSalesNum
}

// GetCreatorSalesNumOk returns a tuple with the CreatorSalesNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorData) GetCreatorSalesNumOk() (*string, bool) {
    if o == nil || utils.IsNil(o.CreatorSalesNum) {
        return nil, false
    }
    return o.CreatorSalesNum, true
}

// HasCreatorSalesNum returns a boolean if a field has been set.
func (o *AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorData) HasCreatorSalesNum() bool {
    if o != nil && !utils.IsNil(o.CreatorSalesNum) {
        return true
    }

    return false
}

// SetCreatorSalesNum gets a reference to the given string and assigns it to the CreatorSalesNum field.
func (o *AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorData) SetCreatorSalesNum(v string) {
    o.CreatorSalesNum = &v
}

// GetEstimatedAmount returns the EstimatedAmount field value if set, zero value otherwise.
func (o *AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorData) GetEstimatedAmount() string {
    if o == nil || utils.IsNil(o.EstimatedAmount) {
        var ret string
        return ret
    }
    return *o.EstimatedAmount
}

// GetEstimatedAmountOk returns a tuple with the EstimatedAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorData) GetEstimatedAmountOk() (*string, bool) {
    if o == nil || utils.IsNil(o.EstimatedAmount) {
        return nil, false
    }
    return o.EstimatedAmount, true
}

// HasEstimatedAmount returns a boolean if a field has been set.
func (o *AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorData) HasEstimatedAmount() bool {
    if o != nil && !utils.IsNil(o.EstimatedAmount) {
        return true
    }

    return false
}

// SetEstimatedAmount gets a reference to the given string and assigns it to the EstimatedAmount field.
func (o *AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorData) SetEstimatedAmount(v string) {
    o.EstimatedAmount = &v
}

// GetEstimatedPartnerCommission returns the EstimatedPartnerCommission field value if set, zero value otherwise.
func (o *AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorData) GetEstimatedPartnerCommission() string {
    if o == nil || utils.IsNil(o.EstimatedPartnerCommission) {
        var ret string
        return ret
    }
    return *o.EstimatedPartnerCommission
}

// GetEstimatedPartnerCommissionOk returns a tuple with the EstimatedPartnerCommission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorData) GetEstimatedPartnerCommissionOk() (*string, bool) {
    if o == nil || utils.IsNil(o.EstimatedPartnerCommission) {
        return nil, false
    }
    return o.EstimatedPartnerCommission, true
}

// HasEstimatedPartnerCommission returns a boolean if a field has been set.
func (o *AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorData) HasEstimatedPartnerCommission() bool {
    if o != nil && !utils.IsNil(o.EstimatedPartnerCommission) {
        return true
    }

    return false
}

// SetEstimatedPartnerCommission gets a reference to the given string and assigns it to the EstimatedPartnerCommission field.
func (o *AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorData) SetEstimatedPartnerCommission(v string) {
    o.EstimatedPartnerCommission = &v
}

// GetPaidOrderNum returns the PaidOrderNum field value if set, zero value otherwise.
func (o *AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorData) GetPaidOrderNum() string {
    if o == nil || utils.IsNil(o.PaidOrderNum) {
        var ret string
        return ret
    }
    return *o.PaidOrderNum
}

// GetPaidOrderNumOk returns a tuple with the PaidOrderNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorData) GetPaidOrderNumOk() (*string, bool) {
    if o == nil || utils.IsNil(o.PaidOrderNum) {
        return nil, false
    }
    return o.PaidOrderNum, true
}

// HasPaidOrderNum returns a boolean if a field has been set.
func (o *AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorData) HasPaidOrderNum() bool {
    if o != nil && !utils.IsNil(o.PaidOrderNum) {
        return true
    }

    return false
}

// SetPaidOrderNum gets a reference to the given string and assigns it to the PaidOrderNum field.
func (o *AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorData) SetPaidOrderNum(v string) {
    o.PaidOrderNum = &v
}

// GetPromotedCreatorNum returns the PromotedCreatorNum field value if set, zero value otherwise.
func (o *AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorData) GetPromotedCreatorNum() string {
    if o == nil || utils.IsNil(o.PromotedCreatorNum) {
        var ret string
        return ret
    }
    return *o.PromotedCreatorNum
}

// GetPromotedCreatorNumOk returns a tuple with the PromotedCreatorNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorData) GetPromotedCreatorNumOk() (*string, bool) {
    if o == nil || utils.IsNil(o.PromotedCreatorNum) {
        return nil, false
    }
    return o.PromotedCreatorNum, true
}

// HasPromotedCreatorNum returns a boolean if a field has been set.
func (o *AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorData) HasPromotedCreatorNum() bool {
    if o != nil && !utils.IsNil(o.PromotedCreatorNum) {
        return true
    }

    return false
}

// SetPromotedCreatorNum gets a reference to the given string and assigns it to the PromotedCreatorNum field.
func (o *AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorData) SetPromotedCreatorNum(v string) {
    o.PromotedCreatorNum = &v
}

// GetSampleRequestedCreatorNum returns the SampleRequestedCreatorNum field value if set, zero value otherwise.
func (o *AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorData) GetSampleRequestedCreatorNum() string {
    if o == nil || utils.IsNil(o.SampleRequestedCreatorNum) {
        var ret string
        return ret
    }
    return *o.SampleRequestedCreatorNum
}

// GetSampleRequestedCreatorNumOk returns a tuple with the SampleRequestedCreatorNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorData) GetSampleRequestedCreatorNumOk() (*string, bool) {
    if o == nil || utils.IsNil(o.SampleRequestedCreatorNum) {
        return nil, false
    }
    return o.SampleRequestedCreatorNum, true
}

// HasSampleRequestedCreatorNum returns a boolean if a field has been set.
func (o *AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorData) HasSampleRequestedCreatorNum() bool {
    if o != nil && !utils.IsNil(o.SampleRequestedCreatorNum) {
        return true
    }

    return false
}

// SetSampleRequestedCreatorNum gets a reference to the given string and assigns it to the SampleRequestedCreatorNum field.
func (o *AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorData) SetSampleRequestedCreatorNum(v string) {
    o.SampleRequestedCreatorNum = &v
}

func (o AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorData) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorData) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.ActualAmount) {
        toSerialize["actual_amount"] = o.ActualAmount
    }
    if !utils.IsNil(o.ActualOrderNum) {
        toSerialize["actual_order_num"] = o.ActualOrderNum
    }
    if !utils.IsNil(o.ActualPartnerCommission) {
        toSerialize["actual_partner_commission"] = o.ActualPartnerCommission
    }
    if !utils.IsNil(o.CollaboratedCreatorsNum) {
        toSerialize["collaborated_creators_num"] = o.CollaboratedCreatorsNum
    }
    if !utils.IsNil(o.CreatorSalesNum) {
        toSerialize["creator_sales_num"] = o.CreatorSalesNum
    }
    if !utils.IsNil(o.EstimatedAmount) {
        toSerialize["estimated_amount"] = o.EstimatedAmount
    }
    if !utils.IsNil(o.EstimatedPartnerCommission) {
        toSerialize["estimated_partner_commission"] = o.EstimatedPartnerCommission
    }
    if !utils.IsNil(o.PaidOrderNum) {
        toSerialize["paid_order_num"] = o.PaidOrderNum
    }
    if !utils.IsNil(o.PromotedCreatorNum) {
        toSerialize["promoted_creator_num"] = o.PromotedCreatorNum
    }
    if !utils.IsNil(o.SampleRequestedCreatorNum) {
        toSerialize["sample_requested_creator_num"] = o.SampleRequestedCreatorNum
    }
    return toSerialize, nil
}

type NullableAffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorData struct {
	value *AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorData
	isSet bool
}

func (v NullableAffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorData) Get() *AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorData {
	return v.value
}

func (v *NullableAffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorData) Set(val *AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorData) {
	v.value = val
	v.isSet = true
}

func (v NullableAffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorData) IsSet() bool {
	return v.isSet
}

func (v *NullableAffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorData(val *AffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorData) *NullableAffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorData {
	return &NullableAffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorData{value: val, isSet: true}
}

func (v NullableAffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAffiliatePartner202501GetAffiliateCampaignCreatorFulfillmentStatusListResponseDataCampaignProductStatisticsCampaignProductDetailIndicatorData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


