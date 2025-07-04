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

            // checks if the Affiliate202309GetShowcaseProductsoldResponseDataProductsMainImages type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Affiliate202309GetShowcaseProductsoldResponseDataProductsMainImages{}

// Affiliate202309GetShowcaseProductsoldResponseDataProductsMainImages struct for Affiliate202309GetShowcaseProductsoldResponseDataProductsMainImages
type Affiliate202309GetShowcaseProductsoldResponseDataProductsMainImages struct {
    // The image height in pixels
    Height *int32 `json:"height,omitempty"`
    // The image's URL
    Url *string `json:"url,omitempty"`
    // The image width in pixels
    Width *int32 `json:"width,omitempty"`
}

// NewAffiliate202309GetShowcaseProductsoldResponseDataProductsMainImages instantiates a new Affiliate202309GetShowcaseProductsoldResponseDataProductsMainImages object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAffiliate202309GetShowcaseProductsoldResponseDataProductsMainImages() *Affiliate202309GetShowcaseProductsoldResponseDataProductsMainImages {
    this := Affiliate202309GetShowcaseProductsoldResponseDataProductsMainImages{}
    return &this
}

// NewAffiliate202309GetShowcaseProductsoldResponseDataProductsMainImagesWithDefaults instantiates a new Affiliate202309GetShowcaseProductsoldResponseDataProductsMainImages object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAffiliate202309GetShowcaseProductsoldResponseDataProductsMainImagesWithDefaults() *Affiliate202309GetShowcaseProductsoldResponseDataProductsMainImages {
    this := Affiliate202309GetShowcaseProductsoldResponseDataProductsMainImages{}
    return &this
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *Affiliate202309GetShowcaseProductsoldResponseDataProductsMainImages) GetHeight() int32 {
    if o == nil || utils.IsNil(o.Height) {
        var ret int32
        return ret
    }
    return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Affiliate202309GetShowcaseProductsoldResponseDataProductsMainImages) GetHeightOk() (*int32, bool) {
    if o == nil || utils.IsNil(o.Height) {
        return nil, false
    }
    return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *Affiliate202309GetShowcaseProductsoldResponseDataProductsMainImages) HasHeight() bool {
    if o != nil && !utils.IsNil(o.Height) {
        return true
    }

    return false
}

// SetHeight gets a reference to the given int32 and assigns it to the Height field.
func (o *Affiliate202309GetShowcaseProductsoldResponseDataProductsMainImages) SetHeight(v int32) {
    o.Height = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *Affiliate202309GetShowcaseProductsoldResponseDataProductsMainImages) GetUrl() string {
    if o == nil || utils.IsNil(o.Url) {
        var ret string
        return ret
    }
    return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Affiliate202309GetShowcaseProductsoldResponseDataProductsMainImages) GetUrlOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Url) {
        return nil, false
    }
    return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *Affiliate202309GetShowcaseProductsoldResponseDataProductsMainImages) HasUrl() bool {
    if o != nil && !utils.IsNil(o.Url) {
        return true
    }

    return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *Affiliate202309GetShowcaseProductsoldResponseDataProductsMainImages) SetUrl(v string) {
    o.Url = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *Affiliate202309GetShowcaseProductsoldResponseDataProductsMainImages) GetWidth() int32 {
    if o == nil || utils.IsNil(o.Width) {
        var ret int32
        return ret
    }
    return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Affiliate202309GetShowcaseProductsoldResponseDataProductsMainImages) GetWidthOk() (*int32, bool) {
    if o == nil || utils.IsNil(o.Width) {
        return nil, false
    }
    return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *Affiliate202309GetShowcaseProductsoldResponseDataProductsMainImages) HasWidth() bool {
    if o != nil && !utils.IsNil(o.Width) {
        return true
    }

    return false
}

// SetWidth gets a reference to the given int32 and assigns it to the Width field.
func (o *Affiliate202309GetShowcaseProductsoldResponseDataProductsMainImages) SetWidth(v int32) {
    o.Width = &v
}

func (o Affiliate202309GetShowcaseProductsoldResponseDataProductsMainImages) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Affiliate202309GetShowcaseProductsoldResponseDataProductsMainImages) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.Height) {
        toSerialize["height"] = o.Height
    }
    if !utils.IsNil(o.Url) {
        toSerialize["url"] = o.Url
    }
    if !utils.IsNil(o.Width) {
        toSerialize["width"] = o.Width
    }
    return toSerialize, nil
}

type NullableAffiliate202309GetShowcaseProductsoldResponseDataProductsMainImages struct {
	value *Affiliate202309GetShowcaseProductsoldResponseDataProductsMainImages
	isSet bool
}

func (v NullableAffiliate202309GetShowcaseProductsoldResponseDataProductsMainImages) Get() *Affiliate202309GetShowcaseProductsoldResponseDataProductsMainImages {
	return v.value
}

func (v *NullableAffiliate202309GetShowcaseProductsoldResponseDataProductsMainImages) Set(val *Affiliate202309GetShowcaseProductsoldResponseDataProductsMainImages) {
	v.value = val
	v.isSet = true
}

func (v NullableAffiliate202309GetShowcaseProductsoldResponseDataProductsMainImages) IsSet() bool {
	return v.isSet
}

func (v *NullableAffiliate202309GetShowcaseProductsoldResponseDataProductsMainImages) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAffiliate202309GetShowcaseProductsoldResponseDataProductsMainImages(val *Affiliate202309GetShowcaseProductsoldResponseDataProductsMainImages) *NullableAffiliate202309GetShowcaseProductsoldResponseDataProductsMainImages {
	return &NullableAffiliate202309GetShowcaseProductsoldResponseDataProductsMainImages{value: val, isSet: true}
}

func (v NullableAffiliate202309GetShowcaseProductsoldResponseDataProductsMainImages) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAffiliate202309GetShowcaseProductsoldResponseDataProductsMainImages) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


