/*
tiktok shop openapi

sdk for apis

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fulfillment_v202309

import (
    "encoding/json"
    "github.com/Yuttachai-101/api_tiktok/utils"
)

            // checks if the Fulfillment202309SearchPackageResponseDataPackages type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Fulfillment202309SearchPackageResponseDataPackages{}

// Fulfillment202309SearchPackageResponseDataPackages struct for Fulfillment202309SearchPackageResponseDataPackages
type Fulfillment202309SearchPackageResponseDataPackages struct {
    // Package creation time. Unix timestamp.
    CreateTime *int64 `json:"create_time,omitempty"`
    // Package ID.
    Id *string `json:"id,omitempty"`
    // The order line item ID contained in the package. 
    OrderLineItemIds []string `json:"order_line_item_ids,omitempty"`
    // The response list of TikTok Shop orders. 
    Orders []Fulfillment202309SearchPackageResponseDataPackagesOrders `json:"orders,omitempty"`
    // Package shipping provider ID.
    ShippingProviderId *string `json:"shipping_provider_id,omitempty"`
    // Package shipping provider.
    ShippingProviderName *string `json:"shipping_provider_name,omitempty"`
    // Possible values: - `PROCESSING`: Package has been arranged by seller. Waiting for carrier to collect the parcel. - `FULFILLING`: Package has been collected by carrier and in transit. - `COMPLETED`: Package has been delivered. - `CANCELLED`: Package has been canceled. Normally, the package is canceled due to the package being lost or damaged.
    Status *string `json:"status,omitempty"`
    // Package tracking number.
    TrackingNumber *string `json:"tracking_number,omitempty"`
    // Package latest update time. Unix timestamp
    UpdateTime *int64 `json:"update_time,omitempty"`
}

// NewFulfillment202309SearchPackageResponseDataPackages instantiates a new Fulfillment202309SearchPackageResponseDataPackages object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFulfillment202309SearchPackageResponseDataPackages() *Fulfillment202309SearchPackageResponseDataPackages {
    this := Fulfillment202309SearchPackageResponseDataPackages{}
    return &this
}

// NewFulfillment202309SearchPackageResponseDataPackagesWithDefaults instantiates a new Fulfillment202309SearchPackageResponseDataPackages object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFulfillment202309SearchPackageResponseDataPackagesWithDefaults() *Fulfillment202309SearchPackageResponseDataPackages {
    this := Fulfillment202309SearchPackageResponseDataPackages{}
    return &this
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *Fulfillment202309SearchPackageResponseDataPackages) GetCreateTime() int64 {
    if o == nil || utils.IsNil(o.CreateTime) {
        var ret int64
        return ret
    }
    return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fulfillment202309SearchPackageResponseDataPackages) GetCreateTimeOk() (*int64, bool) {
    if o == nil || utils.IsNil(o.CreateTime) {
        return nil, false
    }
    return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *Fulfillment202309SearchPackageResponseDataPackages) HasCreateTime() bool {
    if o != nil && !utils.IsNil(o.CreateTime) {
        return true
    }

    return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *Fulfillment202309SearchPackageResponseDataPackages) SetCreateTime(v int64) {
    o.CreateTime = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Fulfillment202309SearchPackageResponseDataPackages) GetId() string {
    if o == nil || utils.IsNil(o.Id) {
        var ret string
        return ret
    }
    return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fulfillment202309SearchPackageResponseDataPackages) GetIdOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Id) {
        return nil, false
    }
    return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Fulfillment202309SearchPackageResponseDataPackages) HasId() bool {
    if o != nil && !utils.IsNil(o.Id) {
        return true
    }

    return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Fulfillment202309SearchPackageResponseDataPackages) SetId(v string) {
    o.Id = &v
}

// GetOrderLineItemIds returns the OrderLineItemIds field value if set, zero value otherwise.
func (o *Fulfillment202309SearchPackageResponseDataPackages) GetOrderLineItemIds() []string {
    if o == nil || utils.IsNil(o.OrderLineItemIds) {
        var ret []string
        return ret
    }
    return o.OrderLineItemIds
}

// GetOrderLineItemIdsOk returns a tuple with the OrderLineItemIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fulfillment202309SearchPackageResponseDataPackages) GetOrderLineItemIdsOk() ([]string, bool) {
    if o == nil || utils.IsNil(o.OrderLineItemIds) {
        return nil, false
    }
    return o.OrderLineItemIds, true
}

// HasOrderLineItemIds returns a boolean if a field has been set.
func (o *Fulfillment202309SearchPackageResponseDataPackages) HasOrderLineItemIds() bool {
    if o != nil && !utils.IsNil(o.OrderLineItemIds) {
        return true
    }

    return false
}

// SetOrderLineItemIds gets a reference to the given []string and assigns it to the OrderLineItemIds field.
func (o *Fulfillment202309SearchPackageResponseDataPackages) SetOrderLineItemIds(v []string) {
    o.OrderLineItemIds = v
}

// GetOrders returns the Orders field value if set, zero value otherwise.
func (o *Fulfillment202309SearchPackageResponseDataPackages) GetOrders() []Fulfillment202309SearchPackageResponseDataPackagesOrders {
    if o == nil || utils.IsNil(o.Orders) {
        var ret []Fulfillment202309SearchPackageResponseDataPackagesOrders
        return ret
    }
    return o.Orders
}

// GetOrdersOk returns a tuple with the Orders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fulfillment202309SearchPackageResponseDataPackages) GetOrdersOk() ([]Fulfillment202309SearchPackageResponseDataPackagesOrders, bool) {
    if o == nil || utils.IsNil(o.Orders) {
        return nil, false
    }
    return o.Orders, true
}

// HasOrders returns a boolean if a field has been set.
func (o *Fulfillment202309SearchPackageResponseDataPackages) HasOrders() bool {
    if o != nil && !utils.IsNil(o.Orders) {
        return true
    }

    return false
}

// SetOrders gets a reference to the given []Fulfillment202309SearchPackageResponseDataPackagesOrders and assigns it to the Orders field.
func (o *Fulfillment202309SearchPackageResponseDataPackages) SetOrders(v []Fulfillment202309SearchPackageResponseDataPackagesOrders) {
    o.Orders = v
}

// GetShippingProviderId returns the ShippingProviderId field value if set, zero value otherwise.
func (o *Fulfillment202309SearchPackageResponseDataPackages) GetShippingProviderId() string {
    if o == nil || utils.IsNil(o.ShippingProviderId) {
        var ret string
        return ret
    }
    return *o.ShippingProviderId
}

// GetShippingProviderIdOk returns a tuple with the ShippingProviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fulfillment202309SearchPackageResponseDataPackages) GetShippingProviderIdOk() (*string, bool) {
    if o == nil || utils.IsNil(o.ShippingProviderId) {
        return nil, false
    }
    return o.ShippingProviderId, true
}

// HasShippingProviderId returns a boolean if a field has been set.
func (o *Fulfillment202309SearchPackageResponseDataPackages) HasShippingProviderId() bool {
    if o != nil && !utils.IsNil(o.ShippingProviderId) {
        return true
    }

    return false
}

// SetShippingProviderId gets a reference to the given string and assigns it to the ShippingProviderId field.
func (o *Fulfillment202309SearchPackageResponseDataPackages) SetShippingProviderId(v string) {
    o.ShippingProviderId = &v
}

// GetShippingProviderName returns the ShippingProviderName field value if set, zero value otherwise.
func (o *Fulfillment202309SearchPackageResponseDataPackages) GetShippingProviderName() string {
    if o == nil || utils.IsNil(o.ShippingProviderName) {
        var ret string
        return ret
    }
    return *o.ShippingProviderName
}

// GetShippingProviderNameOk returns a tuple with the ShippingProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fulfillment202309SearchPackageResponseDataPackages) GetShippingProviderNameOk() (*string, bool) {
    if o == nil || utils.IsNil(o.ShippingProviderName) {
        return nil, false
    }
    return o.ShippingProviderName, true
}

// HasShippingProviderName returns a boolean if a field has been set.
func (o *Fulfillment202309SearchPackageResponseDataPackages) HasShippingProviderName() bool {
    if o != nil && !utils.IsNil(o.ShippingProviderName) {
        return true
    }

    return false
}

// SetShippingProviderName gets a reference to the given string and assigns it to the ShippingProviderName field.
func (o *Fulfillment202309SearchPackageResponseDataPackages) SetShippingProviderName(v string) {
    o.ShippingProviderName = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Fulfillment202309SearchPackageResponseDataPackages) GetStatus() string {
    if o == nil || utils.IsNil(o.Status) {
        var ret string
        return ret
    }
    return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fulfillment202309SearchPackageResponseDataPackages) GetStatusOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Status) {
        return nil, false
    }
    return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Fulfillment202309SearchPackageResponseDataPackages) HasStatus() bool {
    if o != nil && !utils.IsNil(o.Status) {
        return true
    }

    return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Fulfillment202309SearchPackageResponseDataPackages) SetStatus(v string) {
    o.Status = &v
}

// GetTrackingNumber returns the TrackingNumber field value if set, zero value otherwise.
func (o *Fulfillment202309SearchPackageResponseDataPackages) GetTrackingNumber() string {
    if o == nil || utils.IsNil(o.TrackingNumber) {
        var ret string
        return ret
    }
    return *o.TrackingNumber
}

// GetTrackingNumberOk returns a tuple with the TrackingNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fulfillment202309SearchPackageResponseDataPackages) GetTrackingNumberOk() (*string, bool) {
    if o == nil || utils.IsNil(o.TrackingNumber) {
        return nil, false
    }
    return o.TrackingNumber, true
}

// HasTrackingNumber returns a boolean if a field has been set.
func (o *Fulfillment202309SearchPackageResponseDataPackages) HasTrackingNumber() bool {
    if o != nil && !utils.IsNil(o.TrackingNumber) {
        return true
    }

    return false
}

// SetTrackingNumber gets a reference to the given string and assigns it to the TrackingNumber field.
func (o *Fulfillment202309SearchPackageResponseDataPackages) SetTrackingNumber(v string) {
    o.TrackingNumber = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *Fulfillment202309SearchPackageResponseDataPackages) GetUpdateTime() int64 {
    if o == nil || utils.IsNil(o.UpdateTime) {
        var ret int64
        return ret
    }
    return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fulfillment202309SearchPackageResponseDataPackages) GetUpdateTimeOk() (*int64, bool) {
    if o == nil || utils.IsNil(o.UpdateTime) {
        return nil, false
    }
    return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *Fulfillment202309SearchPackageResponseDataPackages) HasUpdateTime() bool {
    if o != nil && !utils.IsNil(o.UpdateTime) {
        return true
    }

    return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *Fulfillment202309SearchPackageResponseDataPackages) SetUpdateTime(v int64) {
    o.UpdateTime = &v
}

func (o Fulfillment202309SearchPackageResponseDataPackages) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Fulfillment202309SearchPackageResponseDataPackages) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.CreateTime) {
        toSerialize["create_time"] = o.CreateTime
    }
    if !utils.IsNil(o.Id) {
        toSerialize["id"] = o.Id
    }
    if !utils.IsNil(o.OrderLineItemIds) {
        toSerialize["order_line_item_ids"] = o.OrderLineItemIds
    }
    if !utils.IsNil(o.Orders) {
        toSerialize["orders"] = o.Orders
    }
    if !utils.IsNil(o.ShippingProviderId) {
        toSerialize["shipping_provider_id"] = o.ShippingProviderId
    }
    if !utils.IsNil(o.ShippingProviderName) {
        toSerialize["shipping_provider_name"] = o.ShippingProviderName
    }
    if !utils.IsNil(o.Status) {
        toSerialize["status"] = o.Status
    }
    if !utils.IsNil(o.TrackingNumber) {
        toSerialize["tracking_number"] = o.TrackingNumber
    }
    if !utils.IsNil(o.UpdateTime) {
        toSerialize["update_time"] = o.UpdateTime
    }
    return toSerialize, nil
}

type NullableFulfillment202309SearchPackageResponseDataPackages struct {
	value *Fulfillment202309SearchPackageResponseDataPackages
	isSet bool
}

func (v NullableFulfillment202309SearchPackageResponseDataPackages) Get() *Fulfillment202309SearchPackageResponseDataPackages {
	return v.value
}

func (v *NullableFulfillment202309SearchPackageResponseDataPackages) Set(val *Fulfillment202309SearchPackageResponseDataPackages) {
	v.value = val
	v.isSet = true
}

func (v NullableFulfillment202309SearchPackageResponseDataPackages) IsSet() bool {
	return v.isSet
}

func (v *NullableFulfillment202309SearchPackageResponseDataPackages) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFulfillment202309SearchPackageResponseDataPackages(val *Fulfillment202309SearchPackageResponseDataPackages) *NullableFulfillment202309SearchPackageResponseDataPackages {
	return &NullableFulfillment202309SearchPackageResponseDataPackages{value: val, isSet: true}
}

func (v NullableFulfillment202309SearchPackageResponseDataPackages) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFulfillment202309SearchPackageResponseDataPackages) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


