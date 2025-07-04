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

            // checks if the Product202309GetProductResponseDataSkusGlobalListingPolicy type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Product202309GetProductResponseDataSkusGlobalListingPolicy{}

// Product202309GetProductResponseDataSkusGlobalListingPolicy struct for Product202309GetProductResponseDataSkusGlobalListingPolicy
type Product202309GetProductResponseDataSkusGlobalListingPolicy struct {
    // The type of inventory to synchronize. Possible values: - SHARED: Inventory Area Sharing - EXCLUSIVE: Inventory Exclusive
    InventoryType *string `json:"inventory_type,omitempty"`
    // A flag indicating whether the product price is synchronized.
    PriceSync *bool `json:"price_sync,omitempty"`
    ReplicateSource *Product202309GetProductResponseDataSkusGlobalListingPolicyReplicateSource `json:"replicate_source,omitempty"`
}

// NewProduct202309GetProductResponseDataSkusGlobalListingPolicy instantiates a new Product202309GetProductResponseDataSkusGlobalListingPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProduct202309GetProductResponseDataSkusGlobalListingPolicy() *Product202309GetProductResponseDataSkusGlobalListingPolicy {
    this := Product202309GetProductResponseDataSkusGlobalListingPolicy{}
    return &this
}

// NewProduct202309GetProductResponseDataSkusGlobalListingPolicyWithDefaults instantiates a new Product202309GetProductResponseDataSkusGlobalListingPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProduct202309GetProductResponseDataSkusGlobalListingPolicyWithDefaults() *Product202309GetProductResponseDataSkusGlobalListingPolicy {
    this := Product202309GetProductResponseDataSkusGlobalListingPolicy{}
    return &this
}

// GetInventoryType returns the InventoryType field value if set, zero value otherwise.
func (o *Product202309GetProductResponseDataSkusGlobalListingPolicy) GetInventoryType() string {
    if o == nil || utils.IsNil(o.InventoryType) {
        var ret string
        return ret
    }
    return *o.InventoryType
}

// GetInventoryTypeOk returns a tuple with the InventoryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309GetProductResponseDataSkusGlobalListingPolicy) GetInventoryTypeOk() (*string, bool) {
    if o == nil || utils.IsNil(o.InventoryType) {
        return nil, false
    }
    return o.InventoryType, true
}

// HasInventoryType returns a boolean if a field has been set.
func (o *Product202309GetProductResponseDataSkusGlobalListingPolicy) HasInventoryType() bool {
    if o != nil && !utils.IsNil(o.InventoryType) {
        return true
    }

    return false
}

// SetInventoryType gets a reference to the given string and assigns it to the InventoryType field.
func (o *Product202309GetProductResponseDataSkusGlobalListingPolicy) SetInventoryType(v string) {
    o.InventoryType = &v
}

// GetPriceSync returns the PriceSync field value if set, zero value otherwise.
func (o *Product202309GetProductResponseDataSkusGlobalListingPolicy) GetPriceSync() bool {
    if o == nil || utils.IsNil(o.PriceSync) {
        var ret bool
        return ret
    }
    return *o.PriceSync
}

// GetPriceSyncOk returns a tuple with the PriceSync field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309GetProductResponseDataSkusGlobalListingPolicy) GetPriceSyncOk() (*bool, bool) {
    if o == nil || utils.IsNil(o.PriceSync) {
        return nil, false
    }
    return o.PriceSync, true
}

// HasPriceSync returns a boolean if a field has been set.
func (o *Product202309GetProductResponseDataSkusGlobalListingPolicy) HasPriceSync() bool {
    if o != nil && !utils.IsNil(o.PriceSync) {
        return true
    }

    return false
}

// SetPriceSync gets a reference to the given bool and assigns it to the PriceSync field.
func (o *Product202309GetProductResponseDataSkusGlobalListingPolicy) SetPriceSync(v bool) {
    o.PriceSync = &v
}

// GetReplicateSource returns the ReplicateSource field value if set, zero value otherwise.
func (o *Product202309GetProductResponseDataSkusGlobalListingPolicy) GetReplicateSource() Product202309GetProductResponseDataSkusGlobalListingPolicyReplicateSource {
    if o == nil || utils.IsNil(o.ReplicateSource) {
        var ret Product202309GetProductResponseDataSkusGlobalListingPolicyReplicateSource
        return ret
    }
    return *o.ReplicateSource
}

// GetReplicateSourceOk returns a tuple with the ReplicateSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309GetProductResponseDataSkusGlobalListingPolicy) GetReplicateSourceOk() (*Product202309GetProductResponseDataSkusGlobalListingPolicyReplicateSource, bool) {
    if o == nil || utils.IsNil(o.ReplicateSource) {
        return nil, false
    }
    return o.ReplicateSource, true
}

// HasReplicateSource returns a boolean if a field has been set.
func (o *Product202309GetProductResponseDataSkusGlobalListingPolicy) HasReplicateSource() bool {
    if o != nil && !utils.IsNil(o.ReplicateSource) {
        return true
    }

    return false
}

// SetReplicateSource gets a reference to the given Product202309GetProductResponseDataSkusGlobalListingPolicyReplicateSource and assigns it to the ReplicateSource field.
func (o *Product202309GetProductResponseDataSkusGlobalListingPolicy) SetReplicateSource(v Product202309GetProductResponseDataSkusGlobalListingPolicyReplicateSource) {
    o.ReplicateSource = &v
}

func (o Product202309GetProductResponseDataSkusGlobalListingPolicy) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Product202309GetProductResponseDataSkusGlobalListingPolicy) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.InventoryType) {
        toSerialize["inventory_type"] = o.InventoryType
    }
    if !utils.IsNil(o.PriceSync) {
        toSerialize["price_sync"] = o.PriceSync
    }
    if !utils.IsNil(o.ReplicateSource) {
        toSerialize["replicate_source"] = o.ReplicateSource
    }
    return toSerialize, nil
}

type NullableProduct202309GetProductResponseDataSkusGlobalListingPolicy struct {
	value *Product202309GetProductResponseDataSkusGlobalListingPolicy
	isSet bool
}

func (v NullableProduct202309GetProductResponseDataSkusGlobalListingPolicy) Get() *Product202309GetProductResponseDataSkusGlobalListingPolicy {
	return v.value
}

func (v *NullableProduct202309GetProductResponseDataSkusGlobalListingPolicy) Set(val *Product202309GetProductResponseDataSkusGlobalListingPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableProduct202309GetProductResponseDataSkusGlobalListingPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableProduct202309GetProductResponseDataSkusGlobalListingPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProduct202309GetProductResponseDataSkusGlobalListingPolicy(val *Product202309GetProductResponseDataSkusGlobalListingPolicy) *NullableProduct202309GetProductResponseDataSkusGlobalListingPolicy {
	return &NullableProduct202309GetProductResponseDataSkusGlobalListingPolicy{value: val, isSet: true}
}

func (v NullableProduct202309GetProductResponseDataSkusGlobalListingPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProduct202309GetProductResponseDataSkusGlobalListingPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


