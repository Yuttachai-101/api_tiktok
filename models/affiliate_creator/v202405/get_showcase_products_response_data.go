/*
tiktok shop openapi

sdk for apis

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package affiliate_creator_v202405

import (
    "encoding/json"
    "github.com/Yuttachai-101/api_tiktok/utils"
)

            // checks if the AffiliateCreator202405GetShowcaseProductsResponseData type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AffiliateCreator202405GetShowcaseProductsResponseData{}

// AffiliateCreator202405GetShowcaseProductsResponseData struct for AffiliateCreator202405GetShowcaseProductsResponseData
type AffiliateCreator202405GetShowcaseProductsResponseData struct {
    // An opaque token used to retrieve the next page of a paginated result set.
    NextPageToken *string `json:"next_page_token,omitempty"`
    // A list of products.
    Products []AffiliateCreator202405GetShowcaseProductsResponseDataProducts `json:"products,omitempty"`
    // Total count of products in the response.
    TotalCount *int64 `json:"total_count,omitempty"`
}

// NewAffiliateCreator202405GetShowcaseProductsResponseData instantiates a new AffiliateCreator202405GetShowcaseProductsResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAffiliateCreator202405GetShowcaseProductsResponseData() *AffiliateCreator202405GetShowcaseProductsResponseData {
    this := AffiliateCreator202405GetShowcaseProductsResponseData{}
    return &this
}

// NewAffiliateCreator202405GetShowcaseProductsResponseDataWithDefaults instantiates a new AffiliateCreator202405GetShowcaseProductsResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAffiliateCreator202405GetShowcaseProductsResponseDataWithDefaults() *AffiliateCreator202405GetShowcaseProductsResponseData {
    this := AffiliateCreator202405GetShowcaseProductsResponseData{}
    return &this
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *AffiliateCreator202405GetShowcaseProductsResponseData) GetNextPageToken() string {
    if o == nil || utils.IsNil(o.NextPageToken) {
        var ret string
        return ret
    }
    return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateCreator202405GetShowcaseProductsResponseData) GetNextPageTokenOk() (*string, bool) {
    if o == nil || utils.IsNil(o.NextPageToken) {
        return nil, false
    }
    return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *AffiliateCreator202405GetShowcaseProductsResponseData) HasNextPageToken() bool {
    if o != nil && !utils.IsNil(o.NextPageToken) {
        return true
    }

    return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *AffiliateCreator202405GetShowcaseProductsResponseData) SetNextPageToken(v string) {
    o.NextPageToken = &v
}

// GetProducts returns the Products field value if set, zero value otherwise.
func (o *AffiliateCreator202405GetShowcaseProductsResponseData) GetProducts() []AffiliateCreator202405GetShowcaseProductsResponseDataProducts {
    if o == nil || utils.IsNil(o.Products) {
        var ret []AffiliateCreator202405GetShowcaseProductsResponseDataProducts
        return ret
    }
    return o.Products
}

// GetProductsOk returns a tuple with the Products field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateCreator202405GetShowcaseProductsResponseData) GetProductsOk() ([]AffiliateCreator202405GetShowcaseProductsResponseDataProducts, bool) {
    if o == nil || utils.IsNil(o.Products) {
        return nil, false
    }
    return o.Products, true
}

// HasProducts returns a boolean if a field has been set.
func (o *AffiliateCreator202405GetShowcaseProductsResponseData) HasProducts() bool {
    if o != nil && !utils.IsNil(o.Products) {
        return true
    }

    return false
}

// SetProducts gets a reference to the given []AffiliateCreator202405GetShowcaseProductsResponseDataProducts and assigns it to the Products field.
func (o *AffiliateCreator202405GetShowcaseProductsResponseData) SetProducts(v []AffiliateCreator202405GetShowcaseProductsResponseDataProducts) {
    o.Products = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *AffiliateCreator202405GetShowcaseProductsResponseData) GetTotalCount() int64 {
    if o == nil || utils.IsNil(o.TotalCount) {
        var ret int64
        return ret
    }
    return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateCreator202405GetShowcaseProductsResponseData) GetTotalCountOk() (*int64, bool) {
    if o == nil || utils.IsNil(o.TotalCount) {
        return nil, false
    }
    return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *AffiliateCreator202405GetShowcaseProductsResponseData) HasTotalCount() bool {
    if o != nil && !utils.IsNil(o.TotalCount) {
        return true
    }

    return false
}

// SetTotalCount gets a reference to the given int64 and assigns it to the TotalCount field.
func (o *AffiliateCreator202405GetShowcaseProductsResponseData) SetTotalCount(v int64) {
    o.TotalCount = &v
}

func (o AffiliateCreator202405GetShowcaseProductsResponseData) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o AffiliateCreator202405GetShowcaseProductsResponseData) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.NextPageToken) {
        toSerialize["next_page_token"] = o.NextPageToken
    }
    if !utils.IsNil(o.Products) {
        toSerialize["products"] = o.Products
    }
    if !utils.IsNil(o.TotalCount) {
        toSerialize["total_count"] = o.TotalCount
    }
    return toSerialize, nil
}

type NullableAffiliateCreator202405GetShowcaseProductsResponseData struct {
	value *AffiliateCreator202405GetShowcaseProductsResponseData
	isSet bool
}

func (v NullableAffiliateCreator202405GetShowcaseProductsResponseData) Get() *AffiliateCreator202405GetShowcaseProductsResponseData {
	return v.value
}

func (v *NullableAffiliateCreator202405GetShowcaseProductsResponseData) Set(val *AffiliateCreator202405GetShowcaseProductsResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableAffiliateCreator202405GetShowcaseProductsResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableAffiliateCreator202405GetShowcaseProductsResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAffiliateCreator202405GetShowcaseProductsResponseData(val *AffiliateCreator202405GetShowcaseProductsResponseData) *NullableAffiliateCreator202405GetShowcaseProductsResponseData {
	return &NullableAffiliateCreator202405GetShowcaseProductsResponseData{value: val, isSet: true}
}

func (v NullableAffiliateCreator202405GetShowcaseProductsResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAffiliateCreator202405GetShowcaseProductsResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


