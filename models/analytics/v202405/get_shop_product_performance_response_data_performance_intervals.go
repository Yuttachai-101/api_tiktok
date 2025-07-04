/*
tiktok shop openapi

sdk for apis

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package analytics_v202405

import (
    "encoding/json"
    "github.com/Yuttachai-101/api_tiktok/utils"
)

            // checks if the Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals{}

// Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals struct for Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals
type Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals struct {
    // Average daily product page visitor breakdowns.
    AvgPageVisitorBreakdowns []Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervalsAvgPageVisitorBreakdowns `json:"avg_page_visitor_breakdowns,omitempty"`
    // Average number of unique visitors per day for the product.
    AvgPageVisitors *int64 `json:"avg_page_visitors,omitempty"`
    // Ratio of the number of product clicks compared to number of product impressions in raw decimal format. To calculate the percentage, multiple it by 100%. Example: 0.0528 <=> 5.28%
    ClickThroughRate *string `json:"click_through_rate,omitempty"`
    // Click through rate breakdowns.
    ClickThroughRateBreakdowns []Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervalsClickThroughRateBreakdowns `json:"click_through_rate_breakdowns,omitempty"`
    // End date of the interval (ISO 8601 YYYY-MM-DD format) in shop registered timezone, exclusive.
    EndDate *string `json:"end_date,omitempty"`
    Gmv *Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervalsGmv `json:"gmv,omitempty"`
    // GMV breakdowns for the product.
    GmvBreakdowns []Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervalsGmvBreakdowns `json:"gmv_breakdowns,omitempty"`
    // Impression breakdowns.
    ImpressionBreakdowns []Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervalsImpressionBreakdowns `json:"impression_breakdowns,omitempty"`
    // Total impressions for the product.
    Impressions *int64 `json:"impressions,omitempty"`
    // Total (sum of all) orders for the product.
    Orders *int64 `json:"orders,omitempty"`
    // Page view breakdowns.
    PageViewBreakdowns []Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervalsPageViewBreakdowns `json:"page_view_breakdowns,omitempty"`
    // Total page views for the product.
    PageViews *int64 `json:"page_views,omitempty"`
    // Start date of the interval (ISO 8601 YYYY-MM-DD format) in shop registered timezone, inclusive.
    StartDate *string `json:"start_date,omitempty"`
    // Unit sold breakdowns.
    UnitSoldBreakdowns []Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervalsUnitSoldBreakdowns `json:"unit_sold_breakdowns,omitempty"`
    // Number of units sold for the product.
    UnitsSold *int64 `json:"units_sold,omitempty"`
}

// NewAnalytics202405GetShopProductPerformanceResponseDataPerformanceIntervals instantiates a new Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnalytics202405GetShopProductPerformanceResponseDataPerformanceIntervals() *Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals {
    this := Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals{}
    return &this
}

// NewAnalytics202405GetShopProductPerformanceResponseDataPerformanceIntervalsWithDefaults instantiates a new Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnalytics202405GetShopProductPerformanceResponseDataPerformanceIntervalsWithDefaults() *Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals {
    this := Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals{}
    return &this
}

// GetAvgPageVisitorBreakdowns returns the AvgPageVisitorBreakdowns field value if set, zero value otherwise.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals) GetAvgPageVisitorBreakdowns() []Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervalsAvgPageVisitorBreakdowns {
    if o == nil || utils.IsNil(o.AvgPageVisitorBreakdowns) {
        var ret []Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervalsAvgPageVisitorBreakdowns
        return ret
    }
    return o.AvgPageVisitorBreakdowns
}

// GetAvgPageVisitorBreakdownsOk returns a tuple with the AvgPageVisitorBreakdowns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals) GetAvgPageVisitorBreakdownsOk() ([]Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervalsAvgPageVisitorBreakdowns, bool) {
    if o == nil || utils.IsNil(o.AvgPageVisitorBreakdowns) {
        return nil, false
    }
    return o.AvgPageVisitorBreakdowns, true
}

// HasAvgPageVisitorBreakdowns returns a boolean if a field has been set.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals) HasAvgPageVisitorBreakdowns() bool {
    if o != nil && !utils.IsNil(o.AvgPageVisitorBreakdowns) {
        return true
    }

    return false
}

// SetAvgPageVisitorBreakdowns gets a reference to the given []Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervalsAvgPageVisitorBreakdowns and assigns it to the AvgPageVisitorBreakdowns field.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals) SetAvgPageVisitorBreakdowns(v []Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervalsAvgPageVisitorBreakdowns) {
    o.AvgPageVisitorBreakdowns = v
}

// GetAvgPageVisitors returns the AvgPageVisitors field value if set, zero value otherwise.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals) GetAvgPageVisitors() int64 {
    if o == nil || utils.IsNil(o.AvgPageVisitors) {
        var ret int64
        return ret
    }
    return *o.AvgPageVisitors
}

// GetAvgPageVisitorsOk returns a tuple with the AvgPageVisitors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals) GetAvgPageVisitorsOk() (*int64, bool) {
    if o == nil || utils.IsNil(o.AvgPageVisitors) {
        return nil, false
    }
    return o.AvgPageVisitors, true
}

// HasAvgPageVisitors returns a boolean if a field has been set.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals) HasAvgPageVisitors() bool {
    if o != nil && !utils.IsNil(o.AvgPageVisitors) {
        return true
    }

    return false
}

// SetAvgPageVisitors gets a reference to the given int64 and assigns it to the AvgPageVisitors field.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals) SetAvgPageVisitors(v int64) {
    o.AvgPageVisitors = &v
}

// GetClickThroughRate returns the ClickThroughRate field value if set, zero value otherwise.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals) GetClickThroughRate() string {
    if o == nil || utils.IsNil(o.ClickThroughRate) {
        var ret string
        return ret
    }
    return *o.ClickThroughRate
}

// GetClickThroughRateOk returns a tuple with the ClickThroughRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals) GetClickThroughRateOk() (*string, bool) {
    if o == nil || utils.IsNil(o.ClickThroughRate) {
        return nil, false
    }
    return o.ClickThroughRate, true
}

// HasClickThroughRate returns a boolean if a field has been set.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals) HasClickThroughRate() bool {
    if o != nil && !utils.IsNil(o.ClickThroughRate) {
        return true
    }

    return false
}

// SetClickThroughRate gets a reference to the given string and assigns it to the ClickThroughRate field.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals) SetClickThroughRate(v string) {
    o.ClickThroughRate = &v
}

// GetClickThroughRateBreakdowns returns the ClickThroughRateBreakdowns field value if set, zero value otherwise.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals) GetClickThroughRateBreakdowns() []Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervalsClickThroughRateBreakdowns {
    if o == nil || utils.IsNil(o.ClickThroughRateBreakdowns) {
        var ret []Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervalsClickThroughRateBreakdowns
        return ret
    }
    return o.ClickThroughRateBreakdowns
}

// GetClickThroughRateBreakdownsOk returns a tuple with the ClickThroughRateBreakdowns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals) GetClickThroughRateBreakdownsOk() ([]Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervalsClickThroughRateBreakdowns, bool) {
    if o == nil || utils.IsNil(o.ClickThroughRateBreakdowns) {
        return nil, false
    }
    return o.ClickThroughRateBreakdowns, true
}

// HasClickThroughRateBreakdowns returns a boolean if a field has been set.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals) HasClickThroughRateBreakdowns() bool {
    if o != nil && !utils.IsNil(o.ClickThroughRateBreakdowns) {
        return true
    }

    return false
}

// SetClickThroughRateBreakdowns gets a reference to the given []Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervalsClickThroughRateBreakdowns and assigns it to the ClickThroughRateBreakdowns field.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals) SetClickThroughRateBreakdowns(v []Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervalsClickThroughRateBreakdowns) {
    o.ClickThroughRateBreakdowns = v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals) GetEndDate() string {
    if o == nil || utils.IsNil(o.EndDate) {
        var ret string
        return ret
    }
    return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals) GetEndDateOk() (*string, bool) {
    if o == nil || utils.IsNil(o.EndDate) {
        return nil, false
    }
    return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals) HasEndDate() bool {
    if o != nil && !utils.IsNil(o.EndDate) {
        return true
    }

    return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals) SetEndDate(v string) {
    o.EndDate = &v
}

// GetGmv returns the Gmv field value if set, zero value otherwise.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals) GetGmv() Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervalsGmv {
    if o == nil || utils.IsNil(o.Gmv) {
        var ret Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervalsGmv
        return ret
    }
    return *o.Gmv
}

// GetGmvOk returns a tuple with the Gmv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals) GetGmvOk() (*Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervalsGmv, bool) {
    if o == nil || utils.IsNil(o.Gmv) {
        return nil, false
    }
    return o.Gmv, true
}

// HasGmv returns a boolean if a field has been set.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals) HasGmv() bool {
    if o != nil && !utils.IsNil(o.Gmv) {
        return true
    }

    return false
}

// SetGmv gets a reference to the given Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervalsGmv and assigns it to the Gmv field.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals) SetGmv(v Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervalsGmv) {
    o.Gmv = &v
}

// GetGmvBreakdowns returns the GmvBreakdowns field value if set, zero value otherwise.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals) GetGmvBreakdowns() []Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervalsGmvBreakdowns {
    if o == nil || utils.IsNil(o.GmvBreakdowns) {
        var ret []Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervalsGmvBreakdowns
        return ret
    }
    return o.GmvBreakdowns
}

// GetGmvBreakdownsOk returns a tuple with the GmvBreakdowns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals) GetGmvBreakdownsOk() ([]Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervalsGmvBreakdowns, bool) {
    if o == nil || utils.IsNil(o.GmvBreakdowns) {
        return nil, false
    }
    return o.GmvBreakdowns, true
}

// HasGmvBreakdowns returns a boolean if a field has been set.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals) HasGmvBreakdowns() bool {
    if o != nil && !utils.IsNil(o.GmvBreakdowns) {
        return true
    }

    return false
}

// SetGmvBreakdowns gets a reference to the given []Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervalsGmvBreakdowns and assigns it to the GmvBreakdowns field.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals) SetGmvBreakdowns(v []Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervalsGmvBreakdowns) {
    o.GmvBreakdowns = v
}

// GetImpressionBreakdowns returns the ImpressionBreakdowns field value if set, zero value otherwise.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals) GetImpressionBreakdowns() []Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervalsImpressionBreakdowns {
    if o == nil || utils.IsNil(o.ImpressionBreakdowns) {
        var ret []Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervalsImpressionBreakdowns
        return ret
    }
    return o.ImpressionBreakdowns
}

// GetImpressionBreakdownsOk returns a tuple with the ImpressionBreakdowns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals) GetImpressionBreakdownsOk() ([]Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervalsImpressionBreakdowns, bool) {
    if o == nil || utils.IsNil(o.ImpressionBreakdowns) {
        return nil, false
    }
    return o.ImpressionBreakdowns, true
}

// HasImpressionBreakdowns returns a boolean if a field has been set.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals) HasImpressionBreakdowns() bool {
    if o != nil && !utils.IsNil(o.ImpressionBreakdowns) {
        return true
    }

    return false
}

// SetImpressionBreakdowns gets a reference to the given []Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervalsImpressionBreakdowns and assigns it to the ImpressionBreakdowns field.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals) SetImpressionBreakdowns(v []Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervalsImpressionBreakdowns) {
    o.ImpressionBreakdowns = v
}

// GetImpressions returns the Impressions field value if set, zero value otherwise.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals) GetImpressions() int64 {
    if o == nil || utils.IsNil(o.Impressions) {
        var ret int64
        return ret
    }
    return *o.Impressions
}

// GetImpressionsOk returns a tuple with the Impressions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals) GetImpressionsOk() (*int64, bool) {
    if o == nil || utils.IsNil(o.Impressions) {
        return nil, false
    }
    return o.Impressions, true
}

// HasImpressions returns a boolean if a field has been set.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals) HasImpressions() bool {
    if o != nil && !utils.IsNil(o.Impressions) {
        return true
    }

    return false
}

// SetImpressions gets a reference to the given int64 and assigns it to the Impressions field.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals) SetImpressions(v int64) {
    o.Impressions = &v
}

// GetOrders returns the Orders field value if set, zero value otherwise.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals) GetOrders() int64 {
    if o == nil || utils.IsNil(o.Orders) {
        var ret int64
        return ret
    }
    return *o.Orders
}

// GetOrdersOk returns a tuple with the Orders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals) GetOrdersOk() (*int64, bool) {
    if o == nil || utils.IsNil(o.Orders) {
        return nil, false
    }
    return o.Orders, true
}

// HasOrders returns a boolean if a field has been set.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals) HasOrders() bool {
    if o != nil && !utils.IsNil(o.Orders) {
        return true
    }

    return false
}

// SetOrders gets a reference to the given int64 and assigns it to the Orders field.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals) SetOrders(v int64) {
    o.Orders = &v
}

// GetPageViewBreakdowns returns the PageViewBreakdowns field value if set, zero value otherwise.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals) GetPageViewBreakdowns() []Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervalsPageViewBreakdowns {
    if o == nil || utils.IsNil(o.PageViewBreakdowns) {
        var ret []Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervalsPageViewBreakdowns
        return ret
    }
    return o.PageViewBreakdowns
}

// GetPageViewBreakdownsOk returns a tuple with the PageViewBreakdowns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals) GetPageViewBreakdownsOk() ([]Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervalsPageViewBreakdowns, bool) {
    if o == nil || utils.IsNil(o.PageViewBreakdowns) {
        return nil, false
    }
    return o.PageViewBreakdowns, true
}

// HasPageViewBreakdowns returns a boolean if a field has been set.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals) HasPageViewBreakdowns() bool {
    if o != nil && !utils.IsNil(o.PageViewBreakdowns) {
        return true
    }

    return false
}

// SetPageViewBreakdowns gets a reference to the given []Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervalsPageViewBreakdowns and assigns it to the PageViewBreakdowns field.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals) SetPageViewBreakdowns(v []Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervalsPageViewBreakdowns) {
    o.PageViewBreakdowns = v
}

// GetPageViews returns the PageViews field value if set, zero value otherwise.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals) GetPageViews() int64 {
    if o == nil || utils.IsNil(o.PageViews) {
        var ret int64
        return ret
    }
    return *o.PageViews
}

// GetPageViewsOk returns a tuple with the PageViews field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals) GetPageViewsOk() (*int64, bool) {
    if o == nil || utils.IsNil(o.PageViews) {
        return nil, false
    }
    return o.PageViews, true
}

// HasPageViews returns a boolean if a field has been set.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals) HasPageViews() bool {
    if o != nil && !utils.IsNil(o.PageViews) {
        return true
    }

    return false
}

// SetPageViews gets a reference to the given int64 and assigns it to the PageViews field.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals) SetPageViews(v int64) {
    o.PageViews = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals) GetStartDate() string {
    if o == nil || utils.IsNil(o.StartDate) {
        var ret string
        return ret
    }
    return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals) GetStartDateOk() (*string, bool) {
    if o == nil || utils.IsNil(o.StartDate) {
        return nil, false
    }
    return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals) HasStartDate() bool {
    if o != nil && !utils.IsNil(o.StartDate) {
        return true
    }

    return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals) SetStartDate(v string) {
    o.StartDate = &v
}

// GetUnitSoldBreakdowns returns the UnitSoldBreakdowns field value if set, zero value otherwise.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals) GetUnitSoldBreakdowns() []Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervalsUnitSoldBreakdowns {
    if o == nil || utils.IsNil(o.UnitSoldBreakdowns) {
        var ret []Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervalsUnitSoldBreakdowns
        return ret
    }
    return o.UnitSoldBreakdowns
}

// GetUnitSoldBreakdownsOk returns a tuple with the UnitSoldBreakdowns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals) GetUnitSoldBreakdownsOk() ([]Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervalsUnitSoldBreakdowns, bool) {
    if o == nil || utils.IsNil(o.UnitSoldBreakdowns) {
        return nil, false
    }
    return o.UnitSoldBreakdowns, true
}

// HasUnitSoldBreakdowns returns a boolean if a field has been set.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals) HasUnitSoldBreakdowns() bool {
    if o != nil && !utils.IsNil(o.UnitSoldBreakdowns) {
        return true
    }

    return false
}

// SetUnitSoldBreakdowns gets a reference to the given []Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervalsUnitSoldBreakdowns and assigns it to the UnitSoldBreakdowns field.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals) SetUnitSoldBreakdowns(v []Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervalsUnitSoldBreakdowns) {
    o.UnitSoldBreakdowns = v
}

// GetUnitsSold returns the UnitsSold field value if set, zero value otherwise.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals) GetUnitsSold() int64 {
    if o == nil || utils.IsNil(o.UnitsSold) {
        var ret int64
        return ret
    }
    return *o.UnitsSold
}

// GetUnitsSoldOk returns a tuple with the UnitsSold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals) GetUnitsSoldOk() (*int64, bool) {
    if o == nil || utils.IsNil(o.UnitsSold) {
        return nil, false
    }
    return o.UnitsSold, true
}

// HasUnitsSold returns a boolean if a field has been set.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals) HasUnitsSold() bool {
    if o != nil && !utils.IsNil(o.UnitsSold) {
        return true
    }

    return false
}

// SetUnitsSold gets a reference to the given int64 and assigns it to the UnitsSold field.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals) SetUnitsSold(v int64) {
    o.UnitsSold = &v
}

func (o Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.AvgPageVisitorBreakdowns) {
        toSerialize["avg_page_visitor_breakdowns"] = o.AvgPageVisitorBreakdowns
    }
    if !utils.IsNil(o.AvgPageVisitors) {
        toSerialize["avg_page_visitors"] = o.AvgPageVisitors
    }
    if !utils.IsNil(o.ClickThroughRate) {
        toSerialize["click_through_rate"] = o.ClickThroughRate
    }
    if !utils.IsNil(o.ClickThroughRateBreakdowns) {
        toSerialize["click_through_rate_breakdowns"] = o.ClickThroughRateBreakdowns
    }
    if !utils.IsNil(o.EndDate) {
        toSerialize["end_date"] = o.EndDate
    }
    if !utils.IsNil(o.Gmv) {
        toSerialize["gmv"] = o.Gmv
    }
    if !utils.IsNil(o.GmvBreakdowns) {
        toSerialize["gmv_breakdowns"] = o.GmvBreakdowns
    }
    if !utils.IsNil(o.ImpressionBreakdowns) {
        toSerialize["impression_breakdowns"] = o.ImpressionBreakdowns
    }
    if !utils.IsNil(o.Impressions) {
        toSerialize["impressions"] = o.Impressions
    }
    if !utils.IsNil(o.Orders) {
        toSerialize["orders"] = o.Orders
    }
    if !utils.IsNil(o.PageViewBreakdowns) {
        toSerialize["page_view_breakdowns"] = o.PageViewBreakdowns
    }
    if !utils.IsNil(o.PageViews) {
        toSerialize["page_views"] = o.PageViews
    }
    if !utils.IsNil(o.StartDate) {
        toSerialize["start_date"] = o.StartDate
    }
    if !utils.IsNil(o.UnitSoldBreakdowns) {
        toSerialize["unit_sold_breakdowns"] = o.UnitSoldBreakdowns
    }
    if !utils.IsNil(o.UnitsSold) {
        toSerialize["units_sold"] = o.UnitsSold
    }
    return toSerialize, nil
}

type NullableAnalytics202405GetShopProductPerformanceResponseDataPerformanceIntervals struct {
	value *Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals
	isSet bool
}

func (v NullableAnalytics202405GetShopProductPerformanceResponseDataPerformanceIntervals) Get() *Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals {
	return v.value
}

func (v *NullableAnalytics202405GetShopProductPerformanceResponseDataPerformanceIntervals) Set(val *Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals) {
	v.value = val
	v.isSet = true
}

func (v NullableAnalytics202405GetShopProductPerformanceResponseDataPerformanceIntervals) IsSet() bool {
	return v.isSet
}

func (v *NullableAnalytics202405GetShopProductPerformanceResponseDataPerformanceIntervals) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnalytics202405GetShopProductPerformanceResponseDataPerformanceIntervals(val *Analytics202405GetShopProductPerformanceResponseDataPerformanceIntervals) *NullableAnalytics202405GetShopProductPerformanceResponseDataPerformanceIntervals {
	return &NullableAnalytics202405GetShopProductPerformanceResponseDataPerformanceIntervals{value: val, isSet: true}
}

func (v NullableAnalytics202405GetShopProductPerformanceResponseDataPerformanceIntervals) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnalytics202405GetShopProductPerformanceResponseDataPerformanceIntervals) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


