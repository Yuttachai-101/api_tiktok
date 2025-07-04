/*
tiktok shop openapi

sdk for apis

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package affiliate_seller_v202405

import (
    "encoding/json"
    "github.com/Yuttachai-101/api_tiktok/utils"
)

            // checks if the AffiliateSeller202405EditOpenCollaborationSettingsRequestBody type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AffiliateSeller202405EditOpenCollaborationSettingsRequestBody{}

// AffiliateSeller202405EditOpenCollaborationSettingsRequestBody struct for AffiliateSeller202405EditOpenCollaborationSettingsRequestBody
type AffiliateSeller202405EditOpenCollaborationSettingsRequestBody struct {
    AutoAddProduct *AffiliateSeller202405EditOpenCollaborationSettingsRequestBodyAutoAddProduct `json:"auto_add_product,omitempty"`
}

// NewAffiliateSeller202405EditOpenCollaborationSettingsRequestBody instantiates a new AffiliateSeller202405EditOpenCollaborationSettingsRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAffiliateSeller202405EditOpenCollaborationSettingsRequestBody() *AffiliateSeller202405EditOpenCollaborationSettingsRequestBody {
    this := AffiliateSeller202405EditOpenCollaborationSettingsRequestBody{}
    return &this
}

// NewAffiliateSeller202405EditOpenCollaborationSettingsRequestBodyWithDefaults instantiates a new AffiliateSeller202405EditOpenCollaborationSettingsRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAffiliateSeller202405EditOpenCollaborationSettingsRequestBodyWithDefaults() *AffiliateSeller202405EditOpenCollaborationSettingsRequestBody {
    this := AffiliateSeller202405EditOpenCollaborationSettingsRequestBody{}
    return &this
}

// GetAutoAddProduct returns the AutoAddProduct field value if set, zero value otherwise.
func (o *AffiliateSeller202405EditOpenCollaborationSettingsRequestBody) GetAutoAddProduct() AffiliateSeller202405EditOpenCollaborationSettingsRequestBodyAutoAddProduct {
    if o == nil || utils.IsNil(o.AutoAddProduct) {
        var ret AffiliateSeller202405EditOpenCollaborationSettingsRequestBodyAutoAddProduct
        return ret
    }
    return *o.AutoAddProduct
}

// GetAutoAddProductOk returns a tuple with the AutoAddProduct field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateSeller202405EditOpenCollaborationSettingsRequestBody) GetAutoAddProductOk() (*AffiliateSeller202405EditOpenCollaborationSettingsRequestBodyAutoAddProduct, bool) {
    if o == nil || utils.IsNil(o.AutoAddProduct) {
        return nil, false
    }
    return o.AutoAddProduct, true
}

// HasAutoAddProduct returns a boolean if a field has been set.
func (o *AffiliateSeller202405EditOpenCollaborationSettingsRequestBody) HasAutoAddProduct() bool {
    if o != nil && !utils.IsNil(o.AutoAddProduct) {
        return true
    }

    return false
}

// SetAutoAddProduct gets a reference to the given AffiliateSeller202405EditOpenCollaborationSettingsRequestBodyAutoAddProduct and assigns it to the AutoAddProduct field.
func (o *AffiliateSeller202405EditOpenCollaborationSettingsRequestBody) SetAutoAddProduct(v AffiliateSeller202405EditOpenCollaborationSettingsRequestBodyAutoAddProduct) {
    o.AutoAddProduct = &v
}

func (o AffiliateSeller202405EditOpenCollaborationSettingsRequestBody) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o AffiliateSeller202405EditOpenCollaborationSettingsRequestBody) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.AutoAddProduct) {
        toSerialize["auto_add_product"] = o.AutoAddProduct
    }
    return toSerialize, nil
}

type NullableAffiliateSeller202405EditOpenCollaborationSettingsRequestBody struct {
	value *AffiliateSeller202405EditOpenCollaborationSettingsRequestBody
	isSet bool
}

func (v NullableAffiliateSeller202405EditOpenCollaborationSettingsRequestBody) Get() *AffiliateSeller202405EditOpenCollaborationSettingsRequestBody {
	return v.value
}

func (v *NullableAffiliateSeller202405EditOpenCollaborationSettingsRequestBody) Set(val *AffiliateSeller202405EditOpenCollaborationSettingsRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableAffiliateSeller202405EditOpenCollaborationSettingsRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableAffiliateSeller202405EditOpenCollaborationSettingsRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAffiliateSeller202405EditOpenCollaborationSettingsRequestBody(val *AffiliateSeller202405EditOpenCollaborationSettingsRequestBody) *NullableAffiliateSeller202405EditOpenCollaborationSettingsRequestBody {
	return &NullableAffiliateSeller202405EditOpenCollaborationSettingsRequestBody{value: val, isSet: true}
}

func (v NullableAffiliateSeller202405EditOpenCollaborationSettingsRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAffiliateSeller202405EditOpenCollaborationSettingsRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


