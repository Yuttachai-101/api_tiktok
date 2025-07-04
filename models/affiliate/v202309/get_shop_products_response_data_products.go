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

            // checks if the Affiliate202309GetShopProductsResponseDataProducts type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Affiliate202309GetShopProductsResponseDataProducts{}

// Affiliate202309GetShopProductsResponseDataProducts struct for Affiliate202309GetShopProductsResponseDataProducts
type Affiliate202309GetShopProductsResponseDataProducts struct {
    // Showcase add status with possible values: - ADDABLE - ADDED - REJECTED
    AddedStatus *string `json:"added_status,omitempty"`
    // The brand name a seller has set for a product.
    BrandName *string `json:"brand_name,omitempty"`
    // TikTok product ID.
    Id *string `json:"id,omitempty"`
    // Images of a product.
    Images []Affiliate202309GetShopProductsResponseDataProductsImages `json:"images,omitempty"`
    Price *Affiliate202309GetShopProductsResponseDataProductsPrice `json:"price,omitempty"`
    // The number of products that have been sold.
    SalesCount *int32 `json:"sales_count,omitempty"`
    // Product name.
    Title *string `json:"title,omitempty"`
}

// NewAffiliate202309GetShopProductsResponseDataProducts instantiates a new Affiliate202309GetShopProductsResponseDataProducts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAffiliate202309GetShopProductsResponseDataProducts() *Affiliate202309GetShopProductsResponseDataProducts {
    this := Affiliate202309GetShopProductsResponseDataProducts{}
    return &this
}

// NewAffiliate202309GetShopProductsResponseDataProductsWithDefaults instantiates a new Affiliate202309GetShopProductsResponseDataProducts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAffiliate202309GetShopProductsResponseDataProductsWithDefaults() *Affiliate202309GetShopProductsResponseDataProducts {
    this := Affiliate202309GetShopProductsResponseDataProducts{}
    return &this
}

// GetAddedStatus returns the AddedStatus field value if set, zero value otherwise.
func (o *Affiliate202309GetShopProductsResponseDataProducts) GetAddedStatus() string {
    if o == nil || utils.IsNil(o.AddedStatus) {
        var ret string
        return ret
    }
    return *o.AddedStatus
}

// GetAddedStatusOk returns a tuple with the AddedStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Affiliate202309GetShopProductsResponseDataProducts) GetAddedStatusOk() (*string, bool) {
    if o == nil || utils.IsNil(o.AddedStatus) {
        return nil, false
    }
    return o.AddedStatus, true
}

// HasAddedStatus returns a boolean if a field has been set.
func (o *Affiliate202309GetShopProductsResponseDataProducts) HasAddedStatus() bool {
    if o != nil && !utils.IsNil(o.AddedStatus) {
        return true
    }

    return false
}

// SetAddedStatus gets a reference to the given string and assigns it to the AddedStatus field.
func (o *Affiliate202309GetShopProductsResponseDataProducts) SetAddedStatus(v string) {
    o.AddedStatus = &v
}

// GetBrandName returns the BrandName field value if set, zero value otherwise.
func (o *Affiliate202309GetShopProductsResponseDataProducts) GetBrandName() string {
    if o == nil || utils.IsNil(o.BrandName) {
        var ret string
        return ret
    }
    return *o.BrandName
}

// GetBrandNameOk returns a tuple with the BrandName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Affiliate202309GetShopProductsResponseDataProducts) GetBrandNameOk() (*string, bool) {
    if o == nil || utils.IsNil(o.BrandName) {
        return nil, false
    }
    return o.BrandName, true
}

// HasBrandName returns a boolean if a field has been set.
func (o *Affiliate202309GetShopProductsResponseDataProducts) HasBrandName() bool {
    if o != nil && !utils.IsNil(o.BrandName) {
        return true
    }

    return false
}

// SetBrandName gets a reference to the given string and assigns it to the BrandName field.
func (o *Affiliate202309GetShopProductsResponseDataProducts) SetBrandName(v string) {
    o.BrandName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Affiliate202309GetShopProductsResponseDataProducts) GetId() string {
    if o == nil || utils.IsNil(o.Id) {
        var ret string
        return ret
    }
    return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Affiliate202309GetShopProductsResponseDataProducts) GetIdOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Id) {
        return nil, false
    }
    return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Affiliate202309GetShopProductsResponseDataProducts) HasId() bool {
    if o != nil && !utils.IsNil(o.Id) {
        return true
    }

    return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Affiliate202309GetShopProductsResponseDataProducts) SetId(v string) {
    o.Id = &v
}

// GetImages returns the Images field value if set, zero value otherwise.
func (o *Affiliate202309GetShopProductsResponseDataProducts) GetImages() []Affiliate202309GetShopProductsResponseDataProductsImages {
    if o == nil || utils.IsNil(o.Images) {
        var ret []Affiliate202309GetShopProductsResponseDataProductsImages
        return ret
    }
    return o.Images
}

// GetImagesOk returns a tuple with the Images field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Affiliate202309GetShopProductsResponseDataProducts) GetImagesOk() ([]Affiliate202309GetShopProductsResponseDataProductsImages, bool) {
    if o == nil || utils.IsNil(o.Images) {
        return nil, false
    }
    return o.Images, true
}

// HasImages returns a boolean if a field has been set.
func (o *Affiliate202309GetShopProductsResponseDataProducts) HasImages() bool {
    if o != nil && !utils.IsNil(o.Images) {
        return true
    }

    return false
}

// SetImages gets a reference to the given []Affiliate202309GetShopProductsResponseDataProductsImages and assigns it to the Images field.
func (o *Affiliate202309GetShopProductsResponseDataProducts) SetImages(v []Affiliate202309GetShopProductsResponseDataProductsImages) {
    o.Images = v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *Affiliate202309GetShopProductsResponseDataProducts) GetPrice() Affiliate202309GetShopProductsResponseDataProductsPrice {
    if o == nil || utils.IsNil(o.Price) {
        var ret Affiliate202309GetShopProductsResponseDataProductsPrice
        return ret
    }
    return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Affiliate202309GetShopProductsResponseDataProducts) GetPriceOk() (*Affiliate202309GetShopProductsResponseDataProductsPrice, bool) {
    if o == nil || utils.IsNil(o.Price) {
        return nil, false
    }
    return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *Affiliate202309GetShopProductsResponseDataProducts) HasPrice() bool {
    if o != nil && !utils.IsNil(o.Price) {
        return true
    }

    return false
}

// SetPrice gets a reference to the given Affiliate202309GetShopProductsResponseDataProductsPrice and assigns it to the Price field.
func (o *Affiliate202309GetShopProductsResponseDataProducts) SetPrice(v Affiliate202309GetShopProductsResponseDataProductsPrice) {
    o.Price = &v
}

// GetSalesCount returns the SalesCount field value if set, zero value otherwise.
func (o *Affiliate202309GetShopProductsResponseDataProducts) GetSalesCount() int32 {
    if o == nil || utils.IsNil(o.SalesCount) {
        var ret int32
        return ret
    }
    return *o.SalesCount
}

// GetSalesCountOk returns a tuple with the SalesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Affiliate202309GetShopProductsResponseDataProducts) GetSalesCountOk() (*int32, bool) {
    if o == nil || utils.IsNil(o.SalesCount) {
        return nil, false
    }
    return o.SalesCount, true
}

// HasSalesCount returns a boolean if a field has been set.
func (o *Affiliate202309GetShopProductsResponseDataProducts) HasSalesCount() bool {
    if o != nil && !utils.IsNil(o.SalesCount) {
        return true
    }

    return false
}

// SetSalesCount gets a reference to the given int32 and assigns it to the SalesCount field.
func (o *Affiliate202309GetShopProductsResponseDataProducts) SetSalesCount(v int32) {
    o.SalesCount = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *Affiliate202309GetShopProductsResponseDataProducts) GetTitle() string {
    if o == nil || utils.IsNil(o.Title) {
        var ret string
        return ret
    }
    return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Affiliate202309GetShopProductsResponseDataProducts) GetTitleOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Title) {
        return nil, false
    }
    return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *Affiliate202309GetShopProductsResponseDataProducts) HasTitle() bool {
    if o != nil && !utils.IsNil(o.Title) {
        return true
    }

    return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *Affiliate202309GetShopProductsResponseDataProducts) SetTitle(v string) {
    o.Title = &v
}

func (o Affiliate202309GetShopProductsResponseDataProducts) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Affiliate202309GetShopProductsResponseDataProducts) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.AddedStatus) {
        toSerialize["added_status"] = o.AddedStatus
    }
    if !utils.IsNil(o.BrandName) {
        toSerialize["brand_name"] = o.BrandName
    }
    if !utils.IsNil(o.Id) {
        toSerialize["id"] = o.Id
    }
    if !utils.IsNil(o.Images) {
        toSerialize["images"] = o.Images
    }
    if !utils.IsNil(o.Price) {
        toSerialize["price"] = o.Price
    }
    if !utils.IsNil(o.SalesCount) {
        toSerialize["sales_count"] = o.SalesCount
    }
    if !utils.IsNil(o.Title) {
        toSerialize["title"] = o.Title
    }
    return toSerialize, nil
}

type NullableAffiliate202309GetShopProductsResponseDataProducts struct {
	value *Affiliate202309GetShopProductsResponseDataProducts
	isSet bool
}

func (v NullableAffiliate202309GetShopProductsResponseDataProducts) Get() *Affiliate202309GetShopProductsResponseDataProducts {
	return v.value
}

func (v *NullableAffiliate202309GetShopProductsResponseDataProducts) Set(val *Affiliate202309GetShopProductsResponseDataProducts) {
	v.value = val
	v.isSet = true
}

func (v NullableAffiliate202309GetShopProductsResponseDataProducts) IsSet() bool {
	return v.isSet
}

func (v *NullableAffiliate202309GetShopProductsResponseDataProducts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAffiliate202309GetShopProductsResponseDataProducts(val *Affiliate202309GetShopProductsResponseDataProducts) *NullableAffiliate202309GetShopProductsResponseDataProducts {
	return &NullableAffiliate202309GetShopProductsResponseDataProducts{value: val, isSet: true}
}

func (v NullableAffiliate202309GetShopProductsResponseDataProducts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAffiliate202309GetShopProductsResponseDataProducts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


