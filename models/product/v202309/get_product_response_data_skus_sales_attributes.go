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

            // checks if the Product202309GetProductResponseDataSkusSalesAttributes type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Product202309GetProductResponseDataSkusSalesAttributes{}

// Product202309GetProductResponseDataSkusSalesAttributes struct for Product202309GetProductResponseDataSkusSalesAttributes
type Product202309GetProductResponseDataSkusSalesAttributes struct {
    // The sales attribute ID.
    Id *string `json:"id,omitempty"`
    // The sales attribute name.
    Name *string `json:"name,omitempty"`
    SkuImg *Product202309GetProductResponseDataSkusSalesAttributesSkuImg `json:"sku_img,omitempty"`
    // A list of supplementary images for each value (e.g. red) of the primary sales attribute (e.g. color) to provide multiple views or details of the product for that attribute value. These appear in the product options gallery on TikTok Shop.  Applicable only for the US market.
    SupplementarySkuImages []Product202309GetProductResponseDataSkusSalesAttributesSupplementarySkuImages `json:"supplementary_sku_images,omitempty"`
    // The sales attribute value ID.
    ValueId *string `json:"value_id,omitempty"`
    // The sales attribute value name.
    ValueName *string `json:"value_name,omitempty"`
}

// NewProduct202309GetProductResponseDataSkusSalesAttributes instantiates a new Product202309GetProductResponseDataSkusSalesAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProduct202309GetProductResponseDataSkusSalesAttributes() *Product202309GetProductResponseDataSkusSalesAttributes {
    this := Product202309GetProductResponseDataSkusSalesAttributes{}
    return &this
}

// NewProduct202309GetProductResponseDataSkusSalesAttributesWithDefaults instantiates a new Product202309GetProductResponseDataSkusSalesAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProduct202309GetProductResponseDataSkusSalesAttributesWithDefaults() *Product202309GetProductResponseDataSkusSalesAttributes {
    this := Product202309GetProductResponseDataSkusSalesAttributes{}
    return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Product202309GetProductResponseDataSkusSalesAttributes) GetId() string {
    if o == nil || utils.IsNil(o.Id) {
        var ret string
        return ret
    }
    return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309GetProductResponseDataSkusSalesAttributes) GetIdOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Id) {
        return nil, false
    }
    return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Product202309GetProductResponseDataSkusSalesAttributes) HasId() bool {
    if o != nil && !utils.IsNil(o.Id) {
        return true
    }

    return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Product202309GetProductResponseDataSkusSalesAttributes) SetId(v string) {
    o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Product202309GetProductResponseDataSkusSalesAttributes) GetName() string {
    if o == nil || utils.IsNil(o.Name) {
        var ret string
        return ret
    }
    return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309GetProductResponseDataSkusSalesAttributes) GetNameOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Name) {
        return nil, false
    }
    return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Product202309GetProductResponseDataSkusSalesAttributes) HasName() bool {
    if o != nil && !utils.IsNil(o.Name) {
        return true
    }

    return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Product202309GetProductResponseDataSkusSalesAttributes) SetName(v string) {
    o.Name = &v
}

// GetSkuImg returns the SkuImg field value if set, zero value otherwise.
func (o *Product202309GetProductResponseDataSkusSalesAttributes) GetSkuImg() Product202309GetProductResponseDataSkusSalesAttributesSkuImg {
    if o == nil || utils.IsNil(o.SkuImg) {
        var ret Product202309GetProductResponseDataSkusSalesAttributesSkuImg
        return ret
    }
    return *o.SkuImg
}

// GetSkuImgOk returns a tuple with the SkuImg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309GetProductResponseDataSkusSalesAttributes) GetSkuImgOk() (*Product202309GetProductResponseDataSkusSalesAttributesSkuImg, bool) {
    if o == nil || utils.IsNil(o.SkuImg) {
        return nil, false
    }
    return o.SkuImg, true
}

// HasSkuImg returns a boolean if a field has been set.
func (o *Product202309GetProductResponseDataSkusSalesAttributes) HasSkuImg() bool {
    if o != nil && !utils.IsNil(o.SkuImg) {
        return true
    }

    return false
}

// SetSkuImg gets a reference to the given Product202309GetProductResponseDataSkusSalesAttributesSkuImg and assigns it to the SkuImg field.
func (o *Product202309GetProductResponseDataSkusSalesAttributes) SetSkuImg(v Product202309GetProductResponseDataSkusSalesAttributesSkuImg) {
    o.SkuImg = &v
}

// GetSupplementarySkuImages returns the SupplementarySkuImages field value if set, zero value otherwise.
func (o *Product202309GetProductResponseDataSkusSalesAttributes) GetSupplementarySkuImages() []Product202309GetProductResponseDataSkusSalesAttributesSupplementarySkuImages {
    if o == nil || utils.IsNil(o.SupplementarySkuImages) {
        var ret []Product202309GetProductResponseDataSkusSalesAttributesSupplementarySkuImages
        return ret
    }
    return o.SupplementarySkuImages
}

// GetSupplementarySkuImagesOk returns a tuple with the SupplementarySkuImages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309GetProductResponseDataSkusSalesAttributes) GetSupplementarySkuImagesOk() ([]Product202309GetProductResponseDataSkusSalesAttributesSupplementarySkuImages, bool) {
    if o == nil || utils.IsNil(o.SupplementarySkuImages) {
        return nil, false
    }
    return o.SupplementarySkuImages, true
}

// HasSupplementarySkuImages returns a boolean if a field has been set.
func (o *Product202309GetProductResponseDataSkusSalesAttributes) HasSupplementarySkuImages() bool {
    if o != nil && !utils.IsNil(o.SupplementarySkuImages) {
        return true
    }

    return false
}

// SetSupplementarySkuImages gets a reference to the given []Product202309GetProductResponseDataSkusSalesAttributesSupplementarySkuImages and assigns it to the SupplementarySkuImages field.
func (o *Product202309GetProductResponseDataSkusSalesAttributes) SetSupplementarySkuImages(v []Product202309GetProductResponseDataSkusSalesAttributesSupplementarySkuImages) {
    o.SupplementarySkuImages = v
}

// GetValueId returns the ValueId field value if set, zero value otherwise.
func (o *Product202309GetProductResponseDataSkusSalesAttributes) GetValueId() string {
    if o == nil || utils.IsNil(o.ValueId) {
        var ret string
        return ret
    }
    return *o.ValueId
}

// GetValueIdOk returns a tuple with the ValueId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309GetProductResponseDataSkusSalesAttributes) GetValueIdOk() (*string, bool) {
    if o == nil || utils.IsNil(o.ValueId) {
        return nil, false
    }
    return o.ValueId, true
}

// HasValueId returns a boolean if a field has been set.
func (o *Product202309GetProductResponseDataSkusSalesAttributes) HasValueId() bool {
    if o != nil && !utils.IsNil(o.ValueId) {
        return true
    }

    return false
}

// SetValueId gets a reference to the given string and assigns it to the ValueId field.
func (o *Product202309GetProductResponseDataSkusSalesAttributes) SetValueId(v string) {
    o.ValueId = &v
}

// GetValueName returns the ValueName field value if set, zero value otherwise.
func (o *Product202309GetProductResponseDataSkusSalesAttributes) GetValueName() string {
    if o == nil || utils.IsNil(o.ValueName) {
        var ret string
        return ret
    }
    return *o.ValueName
}

// GetValueNameOk returns a tuple with the ValueName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309GetProductResponseDataSkusSalesAttributes) GetValueNameOk() (*string, bool) {
    if o == nil || utils.IsNil(o.ValueName) {
        return nil, false
    }
    return o.ValueName, true
}

// HasValueName returns a boolean if a field has been set.
func (o *Product202309GetProductResponseDataSkusSalesAttributes) HasValueName() bool {
    if o != nil && !utils.IsNil(o.ValueName) {
        return true
    }

    return false
}

// SetValueName gets a reference to the given string and assigns it to the ValueName field.
func (o *Product202309GetProductResponseDataSkusSalesAttributes) SetValueName(v string) {
    o.ValueName = &v
}

func (o Product202309GetProductResponseDataSkusSalesAttributes) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Product202309GetProductResponseDataSkusSalesAttributes) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.Id) {
        toSerialize["id"] = o.Id
    }
    if !utils.IsNil(o.Name) {
        toSerialize["name"] = o.Name
    }
    if !utils.IsNil(o.SkuImg) {
        toSerialize["sku_img"] = o.SkuImg
    }
    if !utils.IsNil(o.SupplementarySkuImages) {
        toSerialize["supplementary_sku_images"] = o.SupplementarySkuImages
    }
    if !utils.IsNil(o.ValueId) {
        toSerialize["value_id"] = o.ValueId
    }
    if !utils.IsNil(o.ValueName) {
        toSerialize["value_name"] = o.ValueName
    }
    return toSerialize, nil
}

type NullableProduct202309GetProductResponseDataSkusSalesAttributes struct {
	value *Product202309GetProductResponseDataSkusSalesAttributes
	isSet bool
}

func (v NullableProduct202309GetProductResponseDataSkusSalesAttributes) Get() *Product202309GetProductResponseDataSkusSalesAttributes {
	return v.value
}

func (v *NullableProduct202309GetProductResponseDataSkusSalesAttributes) Set(val *Product202309GetProductResponseDataSkusSalesAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableProduct202309GetProductResponseDataSkusSalesAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableProduct202309GetProductResponseDataSkusSalesAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProduct202309GetProductResponseDataSkusSalesAttributes(val *Product202309GetProductResponseDataSkusSalesAttributes) *NullableProduct202309GetProductResponseDataSkusSalesAttributes {
	return &NullableProduct202309GetProductResponseDataSkusSalesAttributes{value: val, isSet: true}
}

func (v NullableProduct202309GetProductResponseDataSkusSalesAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProduct202309GetProductResponseDataSkusSalesAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


