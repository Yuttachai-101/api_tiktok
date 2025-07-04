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

            // checks if the Product202309InventorySearchResponseDataInventorySkusTotalAvailableInventoryDistributionCampaignInventory type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Product202309InventorySearchResponseDataInventorySkusTotalAvailableInventoryDistributionCampaignInventory{}

// Product202309InventorySearchResponseDataInventorySkusTotalAvailableInventoryDistributionCampaignInventory struct for Product202309InventorySearchResponseDataInventorySkusTotalAvailableInventoryDistributionCampaignInventory
type Product202309InventorySearchResponseDataInventorySkusTotalAvailableInventoryDistributionCampaignInventory struct {
    // The name of the associated campaign.
    CampaignName *string `json:"campaign_name,omitempty"`
    // The number of units allocated.
    Quantity *int32 `json:"quantity,omitempty"`
}

// NewProduct202309InventorySearchResponseDataInventorySkusTotalAvailableInventoryDistributionCampaignInventory instantiates a new Product202309InventorySearchResponseDataInventorySkusTotalAvailableInventoryDistributionCampaignInventory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProduct202309InventorySearchResponseDataInventorySkusTotalAvailableInventoryDistributionCampaignInventory() *Product202309InventorySearchResponseDataInventorySkusTotalAvailableInventoryDistributionCampaignInventory {
    this := Product202309InventorySearchResponseDataInventorySkusTotalAvailableInventoryDistributionCampaignInventory{}
    return &this
}

// NewProduct202309InventorySearchResponseDataInventorySkusTotalAvailableInventoryDistributionCampaignInventoryWithDefaults instantiates a new Product202309InventorySearchResponseDataInventorySkusTotalAvailableInventoryDistributionCampaignInventory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProduct202309InventorySearchResponseDataInventorySkusTotalAvailableInventoryDistributionCampaignInventoryWithDefaults() *Product202309InventorySearchResponseDataInventorySkusTotalAvailableInventoryDistributionCampaignInventory {
    this := Product202309InventorySearchResponseDataInventorySkusTotalAvailableInventoryDistributionCampaignInventory{}
    return &this
}

// GetCampaignName returns the CampaignName field value if set, zero value otherwise.
func (o *Product202309InventorySearchResponseDataInventorySkusTotalAvailableInventoryDistributionCampaignInventory) GetCampaignName() string {
    if o == nil || utils.IsNil(o.CampaignName) {
        var ret string
        return ret
    }
    return *o.CampaignName
}

// GetCampaignNameOk returns a tuple with the CampaignName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309InventorySearchResponseDataInventorySkusTotalAvailableInventoryDistributionCampaignInventory) GetCampaignNameOk() (*string, bool) {
    if o == nil || utils.IsNil(o.CampaignName) {
        return nil, false
    }
    return o.CampaignName, true
}

// HasCampaignName returns a boolean if a field has been set.
func (o *Product202309InventorySearchResponseDataInventorySkusTotalAvailableInventoryDistributionCampaignInventory) HasCampaignName() bool {
    if o != nil && !utils.IsNil(o.CampaignName) {
        return true
    }

    return false
}

// SetCampaignName gets a reference to the given string and assigns it to the CampaignName field.
func (o *Product202309InventorySearchResponseDataInventorySkusTotalAvailableInventoryDistributionCampaignInventory) SetCampaignName(v string) {
    o.CampaignName = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *Product202309InventorySearchResponseDataInventorySkusTotalAvailableInventoryDistributionCampaignInventory) GetQuantity() int32 {
    if o == nil || utils.IsNil(o.Quantity) {
        var ret int32
        return ret
    }
    return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309InventorySearchResponseDataInventorySkusTotalAvailableInventoryDistributionCampaignInventory) GetQuantityOk() (*int32, bool) {
    if o == nil || utils.IsNil(o.Quantity) {
        return nil, false
    }
    return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *Product202309InventorySearchResponseDataInventorySkusTotalAvailableInventoryDistributionCampaignInventory) HasQuantity() bool {
    if o != nil && !utils.IsNil(o.Quantity) {
        return true
    }

    return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *Product202309InventorySearchResponseDataInventorySkusTotalAvailableInventoryDistributionCampaignInventory) SetQuantity(v int32) {
    o.Quantity = &v
}

func (o Product202309InventorySearchResponseDataInventorySkusTotalAvailableInventoryDistributionCampaignInventory) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Product202309InventorySearchResponseDataInventorySkusTotalAvailableInventoryDistributionCampaignInventory) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.CampaignName) {
        toSerialize["campaign_name"] = o.CampaignName
    }
    if !utils.IsNil(o.Quantity) {
        toSerialize["quantity"] = o.Quantity
    }
    return toSerialize, nil
}

type NullableProduct202309InventorySearchResponseDataInventorySkusTotalAvailableInventoryDistributionCampaignInventory struct {
	value *Product202309InventorySearchResponseDataInventorySkusTotalAvailableInventoryDistributionCampaignInventory
	isSet bool
}

func (v NullableProduct202309InventorySearchResponseDataInventorySkusTotalAvailableInventoryDistributionCampaignInventory) Get() *Product202309InventorySearchResponseDataInventorySkusTotalAvailableInventoryDistributionCampaignInventory {
	return v.value
}

func (v *NullableProduct202309InventorySearchResponseDataInventorySkusTotalAvailableInventoryDistributionCampaignInventory) Set(val *Product202309InventorySearchResponseDataInventorySkusTotalAvailableInventoryDistributionCampaignInventory) {
	v.value = val
	v.isSet = true
}

func (v NullableProduct202309InventorySearchResponseDataInventorySkusTotalAvailableInventoryDistributionCampaignInventory) IsSet() bool {
	return v.isSet
}

func (v *NullableProduct202309InventorySearchResponseDataInventorySkusTotalAvailableInventoryDistributionCampaignInventory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProduct202309InventorySearchResponseDataInventorySkusTotalAvailableInventoryDistributionCampaignInventory(val *Product202309InventorySearchResponseDataInventorySkusTotalAvailableInventoryDistributionCampaignInventory) *NullableProduct202309InventorySearchResponseDataInventorySkusTotalAvailableInventoryDistributionCampaignInventory {
	return &NullableProduct202309InventorySearchResponseDataInventorySkusTotalAvailableInventoryDistributionCampaignInventory{value: val, isSet: true}
}

func (v NullableProduct202309InventorySearchResponseDataInventorySkusTotalAvailableInventoryDistributionCampaignInventory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProduct202309InventorySearchResponseDataInventorySkusTotalAvailableInventoryDistributionCampaignInventory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


