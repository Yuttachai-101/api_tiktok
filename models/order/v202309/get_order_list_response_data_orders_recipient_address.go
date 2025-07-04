/*
tiktok shop openapi

sdk for apis

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package order_v202309

import (
    "encoding/json"
    "github.com/Yuttachai-101/api_tiktok/utils"
)

            // checks if the Order202309GetOrderListResponseDataOrdersRecipientAddress type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Order202309GetOrderListResponseDataOrdersRecipientAddress{}

// Order202309GetOrderListResponseDataOrdersRecipientAddress struct for Order202309GetOrderListResponseDataOrdersRecipientAddress
type Order202309GetOrderListResponseDataOrdersRecipientAddress struct {
    // Full recipient detailed address.
    AddressDetail *string `json:"address_detail,omitempty"`
    // The first line of the street address.
    AddressLine1 *string `json:"address_line1,omitempty"`
    // The second line of the street address.
    AddressLine2 *string `json:"address_line2,omitempty"`
    // The third line of the street address. Usually only for the Brazilian market
    AddressLine3 *string `json:"address_line3,omitempty"`
    // The fourth line of the street address. Usually only for the Brazilian market
    AddressLine4 *string `json:"address_line4,omitempty"`
    DeliveryPreferences *Order202309GetOrderListResponseDataOrdersRecipientAddressDeliveryPreferences `json:"delivery_preferences,omitempty"`
    // `district_info` is unavailable under `UNPAID` and `ON_HOLD` statuses.
    DistrictInfo []Order202309GetOrderListResponseDataOrdersRecipientAddressDistrictInfo `json:"district_info,omitempty"`
    // Recipient first name.  If the recipient first and last names are not provided separately, this parameter will have the same value as the `name` parameter.   **Note**: Applicable only for the US, UK, and JP markets.
    FirstName *string `json:"first_name,omitempty"`
    // Recipient first name in katakana. **Note**: Applicable only for the JP market.
    FirstNameLocalScript *string `json:"first_name_local_script,omitempty"`
    // Complete recipient address information.
    FullAddress *string `json:"full_address,omitempty"`
    // Recipient last name.   If the recipient first and last names are not provided separately, this parameter will be empty.   **Note**: Applicable only for the US, UK, and JP markets.
    LastName *string `json:"last_name,omitempty"`
    // Recipient last name in katakana. **Note**: Applicable only for the JP market.
    LastNameLocalScript *string `json:"last_name_local_script,omitempty"`
    // Recipient name.  **Note**: If this order uses platform logistics, the recipient name will be desensitized.
    Name *string `json:"name,omitempty"`
    // Recipient telephone number.  **Note**: If this order uses platform logistics, the phone number will be desensitized.
    PhoneNumber *string `json:"phone_number,omitempty"`
    // The postal code that can be used by seller for shipping. For the US market, this refers to the ZIP Code.
    PostalCode *string `json:"postal_code,omitempty"`
    // Region code.
    RegionCode *string `json:"region_code,omitempty"`
}

// NewOrder202309GetOrderListResponseDataOrdersRecipientAddress instantiates a new Order202309GetOrderListResponseDataOrdersRecipientAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrder202309GetOrderListResponseDataOrdersRecipientAddress() *Order202309GetOrderListResponseDataOrdersRecipientAddress {
    this := Order202309GetOrderListResponseDataOrdersRecipientAddress{}
    return &this
}

// NewOrder202309GetOrderListResponseDataOrdersRecipientAddressWithDefaults instantiates a new Order202309GetOrderListResponseDataOrdersRecipientAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrder202309GetOrderListResponseDataOrdersRecipientAddressWithDefaults() *Order202309GetOrderListResponseDataOrdersRecipientAddress {
    this := Order202309GetOrderListResponseDataOrdersRecipientAddress{}
    return &this
}

// GetAddressDetail returns the AddressDetail field value if set, zero value otherwise.
func (o *Order202309GetOrderListResponseDataOrdersRecipientAddress) GetAddressDetail() string {
    if o == nil || utils.IsNil(o.AddressDetail) {
        var ret string
        return ret
    }
    return *o.AddressDetail
}

// GetAddressDetailOk returns a tuple with the AddressDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order202309GetOrderListResponseDataOrdersRecipientAddress) GetAddressDetailOk() (*string, bool) {
    if o == nil || utils.IsNil(o.AddressDetail) {
        return nil, false
    }
    return o.AddressDetail, true
}

// HasAddressDetail returns a boolean if a field has been set.
func (o *Order202309GetOrderListResponseDataOrdersRecipientAddress) HasAddressDetail() bool {
    if o != nil && !utils.IsNil(o.AddressDetail) {
        return true
    }

    return false
}

// SetAddressDetail gets a reference to the given string and assigns it to the AddressDetail field.
func (o *Order202309GetOrderListResponseDataOrdersRecipientAddress) SetAddressDetail(v string) {
    o.AddressDetail = &v
}

// GetAddressLine1 returns the AddressLine1 field value if set, zero value otherwise.
func (o *Order202309GetOrderListResponseDataOrdersRecipientAddress) GetAddressLine1() string {
    if o == nil || utils.IsNil(o.AddressLine1) {
        var ret string
        return ret
    }
    return *o.AddressLine1
}

// GetAddressLine1Ok returns a tuple with the AddressLine1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order202309GetOrderListResponseDataOrdersRecipientAddress) GetAddressLine1Ok() (*string, bool) {
    if o == nil || utils.IsNil(o.AddressLine1) {
        return nil, false
    }
    return o.AddressLine1, true
}

// HasAddressLine1 returns a boolean if a field has been set.
func (o *Order202309GetOrderListResponseDataOrdersRecipientAddress) HasAddressLine1() bool {
    if o != nil && !utils.IsNil(o.AddressLine1) {
        return true
    }

    return false
}

// SetAddressLine1 gets a reference to the given string and assigns it to the AddressLine1 field.
func (o *Order202309GetOrderListResponseDataOrdersRecipientAddress) SetAddressLine1(v string) {
    o.AddressLine1 = &v
}

// GetAddressLine2 returns the AddressLine2 field value if set, zero value otherwise.
func (o *Order202309GetOrderListResponseDataOrdersRecipientAddress) GetAddressLine2() string {
    if o == nil || utils.IsNil(o.AddressLine2) {
        var ret string
        return ret
    }
    return *o.AddressLine2
}

// GetAddressLine2Ok returns a tuple with the AddressLine2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order202309GetOrderListResponseDataOrdersRecipientAddress) GetAddressLine2Ok() (*string, bool) {
    if o == nil || utils.IsNil(o.AddressLine2) {
        return nil, false
    }
    return o.AddressLine2, true
}

// HasAddressLine2 returns a boolean if a field has been set.
func (o *Order202309GetOrderListResponseDataOrdersRecipientAddress) HasAddressLine2() bool {
    if o != nil && !utils.IsNil(o.AddressLine2) {
        return true
    }

    return false
}

// SetAddressLine2 gets a reference to the given string and assigns it to the AddressLine2 field.
func (o *Order202309GetOrderListResponseDataOrdersRecipientAddress) SetAddressLine2(v string) {
    o.AddressLine2 = &v
}

// GetAddressLine3 returns the AddressLine3 field value if set, zero value otherwise.
func (o *Order202309GetOrderListResponseDataOrdersRecipientAddress) GetAddressLine3() string {
    if o == nil || utils.IsNil(o.AddressLine3) {
        var ret string
        return ret
    }
    return *o.AddressLine3
}

// GetAddressLine3Ok returns a tuple with the AddressLine3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order202309GetOrderListResponseDataOrdersRecipientAddress) GetAddressLine3Ok() (*string, bool) {
    if o == nil || utils.IsNil(o.AddressLine3) {
        return nil, false
    }
    return o.AddressLine3, true
}

// HasAddressLine3 returns a boolean if a field has been set.
func (o *Order202309GetOrderListResponseDataOrdersRecipientAddress) HasAddressLine3() bool {
    if o != nil && !utils.IsNil(o.AddressLine3) {
        return true
    }

    return false
}

// SetAddressLine3 gets a reference to the given string and assigns it to the AddressLine3 field.
func (o *Order202309GetOrderListResponseDataOrdersRecipientAddress) SetAddressLine3(v string) {
    o.AddressLine3 = &v
}

// GetAddressLine4 returns the AddressLine4 field value if set, zero value otherwise.
func (o *Order202309GetOrderListResponseDataOrdersRecipientAddress) GetAddressLine4() string {
    if o == nil || utils.IsNil(o.AddressLine4) {
        var ret string
        return ret
    }
    return *o.AddressLine4
}

// GetAddressLine4Ok returns a tuple with the AddressLine4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order202309GetOrderListResponseDataOrdersRecipientAddress) GetAddressLine4Ok() (*string, bool) {
    if o == nil || utils.IsNil(o.AddressLine4) {
        return nil, false
    }
    return o.AddressLine4, true
}

// HasAddressLine4 returns a boolean if a field has been set.
func (o *Order202309GetOrderListResponseDataOrdersRecipientAddress) HasAddressLine4() bool {
    if o != nil && !utils.IsNil(o.AddressLine4) {
        return true
    }

    return false
}

// SetAddressLine4 gets a reference to the given string and assigns it to the AddressLine4 field.
func (o *Order202309GetOrderListResponseDataOrdersRecipientAddress) SetAddressLine4(v string) {
    o.AddressLine4 = &v
}

// GetDeliveryPreferences returns the DeliveryPreferences field value if set, zero value otherwise.
func (o *Order202309GetOrderListResponseDataOrdersRecipientAddress) GetDeliveryPreferences() Order202309GetOrderListResponseDataOrdersRecipientAddressDeliveryPreferences {
    if o == nil || utils.IsNil(o.DeliveryPreferences) {
        var ret Order202309GetOrderListResponseDataOrdersRecipientAddressDeliveryPreferences
        return ret
    }
    return *o.DeliveryPreferences
}

// GetDeliveryPreferencesOk returns a tuple with the DeliveryPreferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order202309GetOrderListResponseDataOrdersRecipientAddress) GetDeliveryPreferencesOk() (*Order202309GetOrderListResponseDataOrdersRecipientAddressDeliveryPreferences, bool) {
    if o == nil || utils.IsNil(o.DeliveryPreferences) {
        return nil, false
    }
    return o.DeliveryPreferences, true
}

// HasDeliveryPreferences returns a boolean if a field has been set.
func (o *Order202309GetOrderListResponseDataOrdersRecipientAddress) HasDeliveryPreferences() bool {
    if o != nil && !utils.IsNil(o.DeliveryPreferences) {
        return true
    }

    return false
}

// SetDeliveryPreferences gets a reference to the given Order202309GetOrderListResponseDataOrdersRecipientAddressDeliveryPreferences and assigns it to the DeliveryPreferences field.
func (o *Order202309GetOrderListResponseDataOrdersRecipientAddress) SetDeliveryPreferences(v Order202309GetOrderListResponseDataOrdersRecipientAddressDeliveryPreferences) {
    o.DeliveryPreferences = &v
}

// GetDistrictInfo returns the DistrictInfo field value if set, zero value otherwise.
func (o *Order202309GetOrderListResponseDataOrdersRecipientAddress) GetDistrictInfo() []Order202309GetOrderListResponseDataOrdersRecipientAddressDistrictInfo {
    if o == nil || utils.IsNil(o.DistrictInfo) {
        var ret []Order202309GetOrderListResponseDataOrdersRecipientAddressDistrictInfo
        return ret
    }
    return o.DistrictInfo
}

// GetDistrictInfoOk returns a tuple with the DistrictInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order202309GetOrderListResponseDataOrdersRecipientAddress) GetDistrictInfoOk() ([]Order202309GetOrderListResponseDataOrdersRecipientAddressDistrictInfo, bool) {
    if o == nil || utils.IsNil(o.DistrictInfo) {
        return nil, false
    }
    return o.DistrictInfo, true
}

// HasDistrictInfo returns a boolean if a field has been set.
func (o *Order202309GetOrderListResponseDataOrdersRecipientAddress) HasDistrictInfo() bool {
    if o != nil && !utils.IsNil(o.DistrictInfo) {
        return true
    }

    return false
}

// SetDistrictInfo gets a reference to the given []Order202309GetOrderListResponseDataOrdersRecipientAddressDistrictInfo and assigns it to the DistrictInfo field.
func (o *Order202309GetOrderListResponseDataOrdersRecipientAddress) SetDistrictInfo(v []Order202309GetOrderListResponseDataOrdersRecipientAddressDistrictInfo) {
    o.DistrictInfo = v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *Order202309GetOrderListResponseDataOrdersRecipientAddress) GetFirstName() string {
    if o == nil || utils.IsNil(o.FirstName) {
        var ret string
        return ret
    }
    return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order202309GetOrderListResponseDataOrdersRecipientAddress) GetFirstNameOk() (*string, bool) {
    if o == nil || utils.IsNil(o.FirstName) {
        return nil, false
    }
    return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *Order202309GetOrderListResponseDataOrdersRecipientAddress) HasFirstName() bool {
    if o != nil && !utils.IsNil(o.FirstName) {
        return true
    }

    return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *Order202309GetOrderListResponseDataOrdersRecipientAddress) SetFirstName(v string) {
    o.FirstName = &v
}

// GetFirstNameLocalScript returns the FirstNameLocalScript field value if set, zero value otherwise.
func (o *Order202309GetOrderListResponseDataOrdersRecipientAddress) GetFirstNameLocalScript() string {
    if o == nil || utils.IsNil(o.FirstNameLocalScript) {
        var ret string
        return ret
    }
    return *o.FirstNameLocalScript
}

// GetFirstNameLocalScriptOk returns a tuple with the FirstNameLocalScript field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order202309GetOrderListResponseDataOrdersRecipientAddress) GetFirstNameLocalScriptOk() (*string, bool) {
    if o == nil || utils.IsNil(o.FirstNameLocalScript) {
        return nil, false
    }
    return o.FirstNameLocalScript, true
}

// HasFirstNameLocalScript returns a boolean if a field has been set.
func (o *Order202309GetOrderListResponseDataOrdersRecipientAddress) HasFirstNameLocalScript() bool {
    if o != nil && !utils.IsNil(o.FirstNameLocalScript) {
        return true
    }

    return false
}

// SetFirstNameLocalScript gets a reference to the given string and assigns it to the FirstNameLocalScript field.
func (o *Order202309GetOrderListResponseDataOrdersRecipientAddress) SetFirstNameLocalScript(v string) {
    o.FirstNameLocalScript = &v
}

// GetFullAddress returns the FullAddress field value if set, zero value otherwise.
func (o *Order202309GetOrderListResponseDataOrdersRecipientAddress) GetFullAddress() string {
    if o == nil || utils.IsNil(o.FullAddress) {
        var ret string
        return ret
    }
    return *o.FullAddress
}

// GetFullAddressOk returns a tuple with the FullAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order202309GetOrderListResponseDataOrdersRecipientAddress) GetFullAddressOk() (*string, bool) {
    if o == nil || utils.IsNil(o.FullAddress) {
        return nil, false
    }
    return o.FullAddress, true
}

// HasFullAddress returns a boolean if a field has been set.
func (o *Order202309GetOrderListResponseDataOrdersRecipientAddress) HasFullAddress() bool {
    if o != nil && !utils.IsNil(o.FullAddress) {
        return true
    }

    return false
}

// SetFullAddress gets a reference to the given string and assigns it to the FullAddress field.
func (o *Order202309GetOrderListResponseDataOrdersRecipientAddress) SetFullAddress(v string) {
    o.FullAddress = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *Order202309GetOrderListResponseDataOrdersRecipientAddress) GetLastName() string {
    if o == nil || utils.IsNil(o.LastName) {
        var ret string
        return ret
    }
    return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order202309GetOrderListResponseDataOrdersRecipientAddress) GetLastNameOk() (*string, bool) {
    if o == nil || utils.IsNil(o.LastName) {
        return nil, false
    }
    return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *Order202309GetOrderListResponseDataOrdersRecipientAddress) HasLastName() bool {
    if o != nil && !utils.IsNil(o.LastName) {
        return true
    }

    return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *Order202309GetOrderListResponseDataOrdersRecipientAddress) SetLastName(v string) {
    o.LastName = &v
}

// GetLastNameLocalScript returns the LastNameLocalScript field value if set, zero value otherwise.
func (o *Order202309GetOrderListResponseDataOrdersRecipientAddress) GetLastNameLocalScript() string {
    if o == nil || utils.IsNil(o.LastNameLocalScript) {
        var ret string
        return ret
    }
    return *o.LastNameLocalScript
}

// GetLastNameLocalScriptOk returns a tuple with the LastNameLocalScript field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order202309GetOrderListResponseDataOrdersRecipientAddress) GetLastNameLocalScriptOk() (*string, bool) {
    if o == nil || utils.IsNil(o.LastNameLocalScript) {
        return nil, false
    }
    return o.LastNameLocalScript, true
}

// HasLastNameLocalScript returns a boolean if a field has been set.
func (o *Order202309GetOrderListResponseDataOrdersRecipientAddress) HasLastNameLocalScript() bool {
    if o != nil && !utils.IsNil(o.LastNameLocalScript) {
        return true
    }

    return false
}

// SetLastNameLocalScript gets a reference to the given string and assigns it to the LastNameLocalScript field.
func (o *Order202309GetOrderListResponseDataOrdersRecipientAddress) SetLastNameLocalScript(v string) {
    o.LastNameLocalScript = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Order202309GetOrderListResponseDataOrdersRecipientAddress) GetName() string {
    if o == nil || utils.IsNil(o.Name) {
        var ret string
        return ret
    }
    return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order202309GetOrderListResponseDataOrdersRecipientAddress) GetNameOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Name) {
        return nil, false
    }
    return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Order202309GetOrderListResponseDataOrdersRecipientAddress) HasName() bool {
    if o != nil && !utils.IsNil(o.Name) {
        return true
    }

    return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Order202309GetOrderListResponseDataOrdersRecipientAddress) SetName(v string) {
    o.Name = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *Order202309GetOrderListResponseDataOrdersRecipientAddress) GetPhoneNumber() string {
    if o == nil || utils.IsNil(o.PhoneNumber) {
        var ret string
        return ret
    }
    return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order202309GetOrderListResponseDataOrdersRecipientAddress) GetPhoneNumberOk() (*string, bool) {
    if o == nil || utils.IsNil(o.PhoneNumber) {
        return nil, false
    }
    return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *Order202309GetOrderListResponseDataOrdersRecipientAddress) HasPhoneNumber() bool {
    if o != nil && !utils.IsNil(o.PhoneNumber) {
        return true
    }

    return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *Order202309GetOrderListResponseDataOrdersRecipientAddress) SetPhoneNumber(v string) {
    o.PhoneNumber = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *Order202309GetOrderListResponseDataOrdersRecipientAddress) GetPostalCode() string {
    if o == nil || utils.IsNil(o.PostalCode) {
        var ret string
        return ret
    }
    return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order202309GetOrderListResponseDataOrdersRecipientAddress) GetPostalCodeOk() (*string, bool) {
    if o == nil || utils.IsNil(o.PostalCode) {
        return nil, false
    }
    return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *Order202309GetOrderListResponseDataOrdersRecipientAddress) HasPostalCode() bool {
    if o != nil && !utils.IsNil(o.PostalCode) {
        return true
    }

    return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *Order202309GetOrderListResponseDataOrdersRecipientAddress) SetPostalCode(v string) {
    o.PostalCode = &v
}

// GetRegionCode returns the RegionCode field value if set, zero value otherwise.
func (o *Order202309GetOrderListResponseDataOrdersRecipientAddress) GetRegionCode() string {
    if o == nil || utils.IsNil(o.RegionCode) {
        var ret string
        return ret
    }
    return *o.RegionCode
}

// GetRegionCodeOk returns a tuple with the RegionCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order202309GetOrderListResponseDataOrdersRecipientAddress) GetRegionCodeOk() (*string, bool) {
    if o == nil || utils.IsNil(o.RegionCode) {
        return nil, false
    }
    return o.RegionCode, true
}

// HasRegionCode returns a boolean if a field has been set.
func (o *Order202309GetOrderListResponseDataOrdersRecipientAddress) HasRegionCode() bool {
    if o != nil && !utils.IsNil(o.RegionCode) {
        return true
    }

    return false
}

// SetRegionCode gets a reference to the given string and assigns it to the RegionCode field.
func (o *Order202309GetOrderListResponseDataOrdersRecipientAddress) SetRegionCode(v string) {
    o.RegionCode = &v
}

func (o Order202309GetOrderListResponseDataOrdersRecipientAddress) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Order202309GetOrderListResponseDataOrdersRecipientAddress) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.AddressDetail) {
        toSerialize["address_detail"] = o.AddressDetail
    }
    if !utils.IsNil(o.AddressLine1) {
        toSerialize["address_line1"] = o.AddressLine1
    }
    if !utils.IsNil(o.AddressLine2) {
        toSerialize["address_line2"] = o.AddressLine2
    }
    if !utils.IsNil(o.AddressLine3) {
        toSerialize["address_line3"] = o.AddressLine3
    }
    if !utils.IsNil(o.AddressLine4) {
        toSerialize["address_line4"] = o.AddressLine4
    }
    if !utils.IsNil(o.DeliveryPreferences) {
        toSerialize["delivery_preferences"] = o.DeliveryPreferences
    }
    if !utils.IsNil(o.DistrictInfo) {
        toSerialize["district_info"] = o.DistrictInfo
    }
    if !utils.IsNil(o.FirstName) {
        toSerialize["first_name"] = o.FirstName
    }
    if !utils.IsNil(o.FirstNameLocalScript) {
        toSerialize["first_name_local_script"] = o.FirstNameLocalScript
    }
    if !utils.IsNil(o.FullAddress) {
        toSerialize["full_address"] = o.FullAddress
    }
    if !utils.IsNil(o.LastName) {
        toSerialize["last_name"] = o.LastName
    }
    if !utils.IsNil(o.LastNameLocalScript) {
        toSerialize["last_name_local_script"] = o.LastNameLocalScript
    }
    if !utils.IsNil(o.Name) {
        toSerialize["name"] = o.Name
    }
    if !utils.IsNil(o.PhoneNumber) {
        toSerialize["phone_number"] = o.PhoneNumber
    }
    if !utils.IsNil(o.PostalCode) {
        toSerialize["postal_code"] = o.PostalCode
    }
    if !utils.IsNil(o.RegionCode) {
        toSerialize["region_code"] = o.RegionCode
    }
    return toSerialize, nil
}

type NullableOrder202309GetOrderListResponseDataOrdersRecipientAddress struct {
	value *Order202309GetOrderListResponseDataOrdersRecipientAddress
	isSet bool
}

func (v NullableOrder202309GetOrderListResponseDataOrdersRecipientAddress) Get() *Order202309GetOrderListResponseDataOrdersRecipientAddress {
	return v.value
}

func (v *NullableOrder202309GetOrderListResponseDataOrdersRecipientAddress) Set(val *Order202309GetOrderListResponseDataOrdersRecipientAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableOrder202309GetOrderListResponseDataOrdersRecipientAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableOrder202309GetOrderListResponseDataOrdersRecipientAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrder202309GetOrderListResponseDataOrdersRecipientAddress(val *Order202309GetOrderListResponseDataOrdersRecipientAddress) *NullableOrder202309GetOrderListResponseDataOrdersRecipientAddress {
	return &NullableOrder202309GetOrderListResponseDataOrdersRecipientAddress{value: val, isSet: true}
}

func (v NullableOrder202309GetOrderListResponseDataOrdersRecipientAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrder202309GetOrderListResponseDataOrdersRecipientAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


