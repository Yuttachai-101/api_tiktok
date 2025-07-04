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

            // checks if the Product202309GetProductResponseDataSkusSalesAttributesSkuImg type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Product202309GetProductResponseDataSkusSalesAttributesSkuImg{}

// Product202309GetProductResponseDataSkusSalesAttributesSkuImg struct for Product202309GetProductResponseDataSkusSalesAttributesSkuImg
type Product202309GetProductResponseDataSkusSalesAttributesSkuImg struct {
    // The image height. Unit: px
    Height *int32 `json:"height,omitempty"`
    // The URLs to view the image thumbnails.
    ThumbUrls []string `json:"thumb_urls,omitempty"`
    // The URI of the image.
    Uri *string `json:"uri,omitempty"`
    // The URLs to view the images.
    Urls []string `json:"urls,omitempty"`
    // The image width. Unit: px
    Width *int32 `json:"width,omitempty"`
}

// NewProduct202309GetProductResponseDataSkusSalesAttributesSkuImg instantiates a new Product202309GetProductResponseDataSkusSalesAttributesSkuImg object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProduct202309GetProductResponseDataSkusSalesAttributesSkuImg() *Product202309GetProductResponseDataSkusSalesAttributesSkuImg {
    this := Product202309GetProductResponseDataSkusSalesAttributesSkuImg{}
    return &this
}

// NewProduct202309GetProductResponseDataSkusSalesAttributesSkuImgWithDefaults instantiates a new Product202309GetProductResponseDataSkusSalesAttributesSkuImg object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProduct202309GetProductResponseDataSkusSalesAttributesSkuImgWithDefaults() *Product202309GetProductResponseDataSkusSalesAttributesSkuImg {
    this := Product202309GetProductResponseDataSkusSalesAttributesSkuImg{}
    return &this
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *Product202309GetProductResponseDataSkusSalesAttributesSkuImg) GetHeight() int32 {
    if o == nil || utils.IsNil(o.Height) {
        var ret int32
        return ret
    }
    return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309GetProductResponseDataSkusSalesAttributesSkuImg) GetHeightOk() (*int32, bool) {
    if o == nil || utils.IsNil(o.Height) {
        return nil, false
    }
    return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *Product202309GetProductResponseDataSkusSalesAttributesSkuImg) HasHeight() bool {
    if o != nil && !utils.IsNil(o.Height) {
        return true
    }

    return false
}

// SetHeight gets a reference to the given int32 and assigns it to the Height field.
func (o *Product202309GetProductResponseDataSkusSalesAttributesSkuImg) SetHeight(v int32) {
    o.Height = &v
}

// GetThumbUrls returns the ThumbUrls field value if set, zero value otherwise.
func (o *Product202309GetProductResponseDataSkusSalesAttributesSkuImg) GetThumbUrls() []string {
    if o == nil || utils.IsNil(o.ThumbUrls) {
        var ret []string
        return ret
    }
    return o.ThumbUrls
}

// GetThumbUrlsOk returns a tuple with the ThumbUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309GetProductResponseDataSkusSalesAttributesSkuImg) GetThumbUrlsOk() ([]string, bool) {
    if o == nil || utils.IsNil(o.ThumbUrls) {
        return nil, false
    }
    return o.ThumbUrls, true
}

// HasThumbUrls returns a boolean if a field has been set.
func (o *Product202309GetProductResponseDataSkusSalesAttributesSkuImg) HasThumbUrls() bool {
    if o != nil && !utils.IsNil(o.ThumbUrls) {
        return true
    }

    return false
}

// SetThumbUrls gets a reference to the given []string and assigns it to the ThumbUrls field.
func (o *Product202309GetProductResponseDataSkusSalesAttributesSkuImg) SetThumbUrls(v []string) {
    o.ThumbUrls = v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *Product202309GetProductResponseDataSkusSalesAttributesSkuImg) GetUri() string {
    if o == nil || utils.IsNil(o.Uri) {
        var ret string
        return ret
    }
    return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309GetProductResponseDataSkusSalesAttributesSkuImg) GetUriOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Uri) {
        return nil, false
    }
    return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *Product202309GetProductResponseDataSkusSalesAttributesSkuImg) HasUri() bool {
    if o != nil && !utils.IsNil(o.Uri) {
        return true
    }

    return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *Product202309GetProductResponseDataSkusSalesAttributesSkuImg) SetUri(v string) {
    o.Uri = &v
}

// GetUrls returns the Urls field value if set, zero value otherwise.
func (o *Product202309GetProductResponseDataSkusSalesAttributesSkuImg) GetUrls() []string {
    if o == nil || utils.IsNil(o.Urls) {
        var ret []string
        return ret
    }
    return o.Urls
}

// GetUrlsOk returns a tuple with the Urls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309GetProductResponseDataSkusSalesAttributesSkuImg) GetUrlsOk() ([]string, bool) {
    if o == nil || utils.IsNil(o.Urls) {
        return nil, false
    }
    return o.Urls, true
}

// HasUrls returns a boolean if a field has been set.
func (o *Product202309GetProductResponseDataSkusSalesAttributesSkuImg) HasUrls() bool {
    if o != nil && !utils.IsNil(o.Urls) {
        return true
    }

    return false
}

// SetUrls gets a reference to the given []string and assigns it to the Urls field.
func (o *Product202309GetProductResponseDataSkusSalesAttributesSkuImg) SetUrls(v []string) {
    o.Urls = v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *Product202309GetProductResponseDataSkusSalesAttributesSkuImg) GetWidth() int32 {
    if o == nil || utils.IsNil(o.Width) {
        var ret int32
        return ret
    }
    return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309GetProductResponseDataSkusSalesAttributesSkuImg) GetWidthOk() (*int32, bool) {
    if o == nil || utils.IsNil(o.Width) {
        return nil, false
    }
    return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *Product202309GetProductResponseDataSkusSalesAttributesSkuImg) HasWidth() bool {
    if o != nil && !utils.IsNil(o.Width) {
        return true
    }

    return false
}

// SetWidth gets a reference to the given int32 and assigns it to the Width field.
func (o *Product202309GetProductResponseDataSkusSalesAttributesSkuImg) SetWidth(v int32) {
    o.Width = &v
}

func (o Product202309GetProductResponseDataSkusSalesAttributesSkuImg) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Product202309GetProductResponseDataSkusSalesAttributesSkuImg) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.Height) {
        toSerialize["height"] = o.Height
    }
    if !utils.IsNil(o.ThumbUrls) {
        toSerialize["thumb_urls"] = o.ThumbUrls
    }
    if !utils.IsNil(o.Uri) {
        toSerialize["uri"] = o.Uri
    }
    if !utils.IsNil(o.Urls) {
        toSerialize["urls"] = o.Urls
    }
    if !utils.IsNil(o.Width) {
        toSerialize["width"] = o.Width
    }
    return toSerialize, nil
}

type NullableProduct202309GetProductResponseDataSkusSalesAttributesSkuImg struct {
	value *Product202309GetProductResponseDataSkusSalesAttributesSkuImg
	isSet bool
}

func (v NullableProduct202309GetProductResponseDataSkusSalesAttributesSkuImg) Get() *Product202309GetProductResponseDataSkusSalesAttributesSkuImg {
	return v.value
}

func (v *NullableProduct202309GetProductResponseDataSkusSalesAttributesSkuImg) Set(val *Product202309GetProductResponseDataSkusSalesAttributesSkuImg) {
	v.value = val
	v.isSet = true
}

func (v NullableProduct202309GetProductResponseDataSkusSalesAttributesSkuImg) IsSet() bool {
	return v.isSet
}

func (v *NullableProduct202309GetProductResponseDataSkusSalesAttributesSkuImg) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProduct202309GetProductResponseDataSkusSalesAttributesSkuImg(val *Product202309GetProductResponseDataSkusSalesAttributesSkuImg) *NullableProduct202309GetProductResponseDataSkusSalesAttributesSkuImg {
	return &NullableProduct202309GetProductResponseDataSkusSalesAttributesSkuImg{value: val, isSet: true}
}

func (v NullableProduct202309GetProductResponseDataSkusSalesAttributesSkuImg) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProduct202309GetProductResponseDataSkusSalesAttributesSkuImg) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


