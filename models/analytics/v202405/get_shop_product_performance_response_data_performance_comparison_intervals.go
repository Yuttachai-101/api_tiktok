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

            // checks if the Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals{}

// Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals struct for Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals
type Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals struct {
    // Average daily product page visitor breakdowns.
    AvgPageVisitorBreakdowns []Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervalsAvgPageVisitorBreakdowns `json:"avg_page_visitor_breakdowns,omitempty"`
    // Average number of unique visitors per day for the product.
    AvgPageVisitors *int64 `json:"avg_page_visitors,omitempty"`
    // Ratio of the number of product clicks compared to number of product impressions in raw decimal format. To calculate the percentage, multiple it by 100%. Example: 0.0528 <=> 5.28%
    ClickThroughRate *string `json:"click_through_rate,omitempty"`
    // Click through rate breakdowns.
    ClickThroughRateBreakdowns []Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervalsClickThroughRateBreakdowns `json:"click_through_rate_breakdowns,omitempty"`
    // End date of the interval (ISO 8601 YYYY-MM-DD format) in shop registered timezone, exclusive.
    EndDate *string `json:"end_date,omitempty"`
    Gmv *Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervalsGmv `json:"gmv,omitempty"`
    // GMV breakdowns for the product.
    GmvBreakdowns []Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervalsGmvBreakdowns `json:"gmv_breakdowns,omitempty"`
    // Impression breakdowns.
    ImpressionBreakdowns []Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervalsImpressionBreakdowns `json:"impression_breakdowns,omitempty"`
    // Total impressions for the product.
    Impressions *int64 `json:"impressions,omitempty"`
    // Total (sum of all) orders for the product.
    Orders *int64 `json:"orders,omitempty"`
    // Page view breakdowns.
    PageViewBreakdowns []Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervalsPageViewBreakdowns `json:"page_view_breakdowns,omitempty"`
    // Total page views for the product.
    PageViews *int64 `json:"page_views,omitempty"`
    // Start date of the interval (ISO 8601 YYYY-MM-DD format) in shop registered timezone, inclusive.
    StartDate *string `json:"start_date,omitempty"`
    // Unit sold breakdowns.
    UnitSoldBreakdowns []Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervalsUnitSoldBreakdowns `json:"unit_sold_breakdowns,omitempty"`
    // Number of units sold for the product.
    UnitsSold *int64 `json:"units_sold,omitempty"`
}

// NewAnalytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals instantiates a new Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnalytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals() *Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals {
    this := Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals{}
    return &this
}

// NewAnalytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervalsWithDefaults instantiates a new Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnalytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervalsWithDefaults() *Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals {
    this := Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals{}
    return &this
}

// GetAvgPageVisitorBreakdowns returns the AvgPageVisitorBreakdowns field value if set, zero value otherwise.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals) GetAvgPageVisitorBreakdowns() []Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervalsAvgPageVisitorBreakdowns {
    if o == nil || utils.IsNil(o.AvgPageVisitorBreakdowns) {
        var ret []Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervalsAvgPageVisitorBreakdowns
        return ret
    }
    return o.AvgPageVisitorBreakdowns
}

// GetAvgPageVisitorBreakdownsOk returns a tuple with the AvgPageVisitorBreakdowns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals) GetAvgPageVisitorBreakdownsOk() ([]Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervalsAvgPageVisitorBreakdowns, bool) {
    if o == nil || utils.IsNil(o.AvgPageVisitorBreakdowns) {
        return nil, false
    }
    return o.AvgPageVisitorBreakdowns, true
}

// HasAvgPageVisitorBreakdowns returns a boolean if a field has been set.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals) HasAvgPageVisitorBreakdowns() bool {
    if o != nil && !utils.IsNil(o.AvgPageVisitorBreakdowns) {
        return true
    }

    return false
}

// SetAvgPageVisitorBreakdowns gets a reference to the given []Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervalsAvgPageVisitorBreakdowns and assigns it to the AvgPageVisitorBreakdowns field.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals) SetAvgPageVisitorBreakdowns(v []Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervalsAvgPageVisitorBreakdowns) {
    o.AvgPageVisitorBreakdowns = v
}

// GetAvgPageVisitors returns the AvgPageVisitors field value if set, zero value otherwise.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals) GetAvgPageVisitors() int64 {
    if o == nil || utils.IsNil(o.AvgPageVisitors) {
        var ret int64
        return ret
    }
    return *o.AvgPageVisitors
}

// GetAvgPageVisitorsOk returns a tuple with the AvgPageVisitors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals) GetAvgPageVisitorsOk() (*int64, bool) {
    if o == nil || utils.IsNil(o.AvgPageVisitors) {
        return nil, false
    }
    return o.AvgPageVisitors, true
}

// HasAvgPageVisitors returns a boolean if a field has been set.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals) HasAvgPageVisitors() bool {
    if o != nil && !utils.IsNil(o.AvgPageVisitors) {
        return true
    }

    return false
}

// SetAvgPageVisitors gets a reference to the given int64 and assigns it to the AvgPageVisitors field.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals) SetAvgPageVisitors(v int64) {
    o.AvgPageVisitors = &v
}

// GetClickThroughRate returns the ClickThroughRate field value if set, zero value otherwise.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals) GetClickThroughRate() string {
    if o == nil || utils.IsNil(o.ClickThroughRate) {
        var ret string
        return ret
    }
    return *o.ClickThroughRate
}

// GetClickThroughRateOk returns a tuple with the ClickThroughRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals) GetClickThroughRateOk() (*string, bool) {
    if o == nil || utils.IsNil(o.ClickThroughRate) {
        return nil, false
    }
    return o.ClickThroughRate, true
}

// HasClickThroughRate returns a boolean if a field has been set.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals) HasClickThroughRate() bool {
    if o != nil && !utils.IsNil(o.ClickThroughRate) {
        return true
    }

    return false
}

// SetClickThroughRate gets a reference to the given string and assigns it to the ClickThroughRate field.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals) SetClickThroughRate(v string) {
    o.ClickThroughRate = &v
}

// GetClickThroughRateBreakdowns returns the ClickThroughRateBreakdowns field value if set, zero value otherwise.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals) GetClickThroughRateBreakdowns() []Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervalsClickThroughRateBreakdowns {
    if o == nil || utils.IsNil(o.ClickThroughRateBreakdowns) {
        var ret []Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervalsClickThroughRateBreakdowns
        return ret
    }
    return o.ClickThroughRateBreakdowns
}

// GetClickThroughRateBreakdownsOk returns a tuple with the ClickThroughRateBreakdowns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals) GetClickThroughRateBreakdownsOk() ([]Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervalsClickThroughRateBreakdowns, bool) {
    if o == nil || utils.IsNil(o.ClickThroughRateBreakdowns) {
        return nil, false
    }
    return o.ClickThroughRateBreakdowns, true
}

// HasClickThroughRateBreakdowns returns a boolean if a field has been set.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals) HasClickThroughRateBreakdowns() bool {
    if o != nil && !utils.IsNil(o.ClickThroughRateBreakdowns) {
        return true
    }

    return false
}

// SetClickThroughRateBreakdowns gets a reference to the given []Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervalsClickThroughRateBreakdowns and assigns it to the ClickThroughRateBreakdowns field.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals) SetClickThroughRateBreakdowns(v []Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervalsClickThroughRateBreakdowns) {
    o.ClickThroughRateBreakdowns = v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals) GetEndDate() string {
    if o == nil || utils.IsNil(o.EndDate) {
        var ret string
        return ret
    }
    return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals) GetEndDateOk() (*string, bool) {
    if o == nil || utils.IsNil(o.EndDate) {
        return nil, false
    }
    return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals) HasEndDate() bool {
    if o != nil && !utils.IsNil(o.EndDate) {
        return true
    }

    return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals) SetEndDate(v string) {
    o.EndDate = &v
}

// GetGmv returns the Gmv field value if set, zero value otherwise.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals) GetGmv() Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervalsGmv {
    if o == nil || utils.IsNil(o.Gmv) {
        var ret Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervalsGmv
        return ret
    }
    return *o.Gmv
}

// GetGmvOk returns a tuple with the Gmv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals) GetGmvOk() (*Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervalsGmv, bool) {
    if o == nil || utils.IsNil(o.Gmv) {
        return nil, false
    }
    return o.Gmv, true
}

// HasGmv returns a boolean if a field has been set.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals) HasGmv() bool {
    if o != nil && !utils.IsNil(o.Gmv) {
        return true
    }

    return false
}

// SetGmv gets a reference to the given Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervalsGmv and assigns it to the Gmv field.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals) SetGmv(v Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervalsGmv) {
    o.Gmv = &v
}

// GetGmvBreakdowns returns the GmvBreakdowns field value if set, zero value otherwise.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals) GetGmvBreakdowns() []Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervalsGmvBreakdowns {
    if o == nil || utils.IsNil(o.GmvBreakdowns) {
        var ret []Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervalsGmvBreakdowns
        return ret
    }
    return o.GmvBreakdowns
}

// GetGmvBreakdownsOk returns a tuple with the GmvBreakdowns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals) GetGmvBreakdownsOk() ([]Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervalsGmvBreakdowns, bool) {
    if o == nil || utils.IsNil(o.GmvBreakdowns) {
        return nil, false
    }
    return o.GmvBreakdowns, true
}

// HasGmvBreakdowns returns a boolean if a field has been set.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals) HasGmvBreakdowns() bool {
    if o != nil && !utils.IsNil(o.GmvBreakdowns) {
        return true
    }

    return false
}

// SetGmvBreakdowns gets a reference to the given []Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervalsGmvBreakdowns and assigns it to the GmvBreakdowns field.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals) SetGmvBreakdowns(v []Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervalsGmvBreakdowns) {
    o.GmvBreakdowns = v
}

// GetImpressionBreakdowns returns the ImpressionBreakdowns field value if set, zero value otherwise.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals) GetImpressionBreakdowns() []Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervalsImpressionBreakdowns {
    if o == nil || utils.IsNil(o.ImpressionBreakdowns) {
        var ret []Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervalsImpressionBreakdowns
        return ret
    }
    return o.ImpressionBreakdowns
}

// GetImpressionBreakdownsOk returns a tuple with the ImpressionBreakdowns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals) GetImpressionBreakdownsOk() ([]Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervalsImpressionBreakdowns, bool) {
    if o == nil || utils.IsNil(o.ImpressionBreakdowns) {
        return nil, false
    }
    return o.ImpressionBreakdowns, true
}

// HasImpressionBreakdowns returns a boolean if a field has been set.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals) HasImpressionBreakdowns() bool {
    if o != nil && !utils.IsNil(o.ImpressionBreakdowns) {
        return true
    }

    return false
}

// SetImpressionBreakdowns gets a reference to the given []Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervalsImpressionBreakdowns and assigns it to the ImpressionBreakdowns field.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals) SetImpressionBreakdowns(v []Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervalsImpressionBreakdowns) {
    o.ImpressionBreakdowns = v
}

// GetImpressions returns the Impressions field value if set, zero value otherwise.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals) GetImpressions() int64 {
    if o == nil || utils.IsNil(o.Impressions) {
        var ret int64
        return ret
    }
    return *o.Impressions
}

// GetImpressionsOk returns a tuple with the Impressions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals) GetImpressionsOk() (*int64, bool) {
    if o == nil || utils.IsNil(o.Impressions) {
        return nil, false
    }
    return o.Impressions, true
}

// HasImpressions returns a boolean if a field has been set.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals) HasImpressions() bool {
    if o != nil && !utils.IsNil(o.Impressions) {
        return true
    }

    return false
}

// SetImpressions gets a reference to the given int64 and assigns it to the Impressions field.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals) SetImpressions(v int64) {
    o.Impressions = &v
}

// GetOrders returns the Orders field value if set, zero value otherwise.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals) GetOrders() int64 {
    if o == nil || utils.IsNil(o.Orders) {
        var ret int64
        return ret
    }
    return *o.Orders
}

// GetOrdersOk returns a tuple with the Orders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals) GetOrdersOk() (*int64, bool) {
    if o == nil || utils.IsNil(o.Orders) {
        return nil, false
    }
    return o.Orders, true
}

// HasOrders returns a boolean if a field has been set.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals) HasOrders() bool {
    if o != nil && !utils.IsNil(o.Orders) {
        return true
    }

    return false
}

// SetOrders gets a reference to the given int64 and assigns it to the Orders field.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals) SetOrders(v int64) {
    o.Orders = &v
}

// GetPageViewBreakdowns returns the PageViewBreakdowns field value if set, zero value otherwise.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals) GetPageViewBreakdowns() []Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervalsPageViewBreakdowns {
    if o == nil || utils.IsNil(o.PageViewBreakdowns) {
        var ret []Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervalsPageViewBreakdowns
        return ret
    }
    return o.PageViewBreakdowns
}

// GetPageViewBreakdownsOk returns a tuple with the PageViewBreakdowns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals) GetPageViewBreakdownsOk() ([]Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervalsPageViewBreakdowns, bool) {
    if o == nil || utils.IsNil(o.PageViewBreakdowns) {
        return nil, false
    }
    return o.PageViewBreakdowns, true
}

// HasPageViewBreakdowns returns a boolean if a field has been set.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals) HasPageViewBreakdowns() bool {
    if o != nil && !utils.IsNil(o.PageViewBreakdowns) {
        return true
    }

    return false
}

// SetPageViewBreakdowns gets a reference to the given []Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervalsPageViewBreakdowns and assigns it to the PageViewBreakdowns field.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals) SetPageViewBreakdowns(v []Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervalsPageViewBreakdowns) {
    o.PageViewBreakdowns = v
}

// GetPageViews returns the PageViews field value if set, zero value otherwise.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals) GetPageViews() int64 {
    if o == nil || utils.IsNil(o.PageViews) {
        var ret int64
        return ret
    }
    return *o.PageViews
}

// GetPageViewsOk returns a tuple with the PageViews field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals) GetPageViewsOk() (*int64, bool) {
    if o == nil || utils.IsNil(o.PageViews) {
        return nil, false
    }
    return o.PageViews, true
}

// HasPageViews returns a boolean if a field has been set.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals) HasPageViews() bool {
    if o != nil && !utils.IsNil(o.PageViews) {
        return true
    }

    return false
}

// SetPageViews gets a reference to the given int64 and assigns it to the PageViews field.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals) SetPageViews(v int64) {
    o.PageViews = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals) GetStartDate() string {
    if o == nil || utils.IsNil(o.StartDate) {
        var ret string
        return ret
    }
    return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals) GetStartDateOk() (*string, bool) {
    if o == nil || utils.IsNil(o.StartDate) {
        return nil, false
    }
    return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals) HasStartDate() bool {
    if o != nil && !utils.IsNil(o.StartDate) {
        return true
    }

    return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals) SetStartDate(v string) {
    o.StartDate = &v
}

// GetUnitSoldBreakdowns returns the UnitSoldBreakdowns field value if set, zero value otherwise.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals) GetUnitSoldBreakdowns() []Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervalsUnitSoldBreakdowns {
    if o == nil || utils.IsNil(o.UnitSoldBreakdowns) {
        var ret []Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervalsUnitSoldBreakdowns
        return ret
    }
    return o.UnitSoldBreakdowns
}

// GetUnitSoldBreakdownsOk returns a tuple with the UnitSoldBreakdowns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals) GetUnitSoldBreakdownsOk() ([]Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervalsUnitSoldBreakdowns, bool) {
    if o == nil || utils.IsNil(o.UnitSoldBreakdowns) {
        return nil, false
    }
    return o.UnitSoldBreakdowns, true
}

// HasUnitSoldBreakdowns returns a boolean if a field has been set.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals) HasUnitSoldBreakdowns() bool {
    if o != nil && !utils.IsNil(o.UnitSoldBreakdowns) {
        return true
    }

    return false
}

// SetUnitSoldBreakdowns gets a reference to the given []Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervalsUnitSoldBreakdowns and assigns it to the UnitSoldBreakdowns field.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals) SetUnitSoldBreakdowns(v []Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervalsUnitSoldBreakdowns) {
    o.UnitSoldBreakdowns = v
}

// GetUnitsSold returns the UnitsSold field value if set, zero value otherwise.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals) GetUnitsSold() int64 {
    if o == nil || utils.IsNil(o.UnitsSold) {
        var ret int64
        return ret
    }
    return *o.UnitsSold
}

// GetUnitsSoldOk returns a tuple with the UnitsSold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals) GetUnitsSoldOk() (*int64, bool) {
    if o == nil || utils.IsNil(o.UnitsSold) {
        return nil, false
    }
    return o.UnitsSold, true
}

// HasUnitsSold returns a boolean if a field has been set.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals) HasUnitsSold() bool {
    if o != nil && !utils.IsNil(o.UnitsSold) {
        return true
    }

    return false
}

// SetUnitsSold gets a reference to the given int64 and assigns it to the UnitsSold field.
func (o *Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals) SetUnitsSold(v int64) {
    o.UnitsSold = &v
}

func (o Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals) ToMap() (map[string]interface{}, error) {
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

type NullableAnalytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals struct {
	value *Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals
	isSet bool
}

func (v NullableAnalytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals) Get() *Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals {
	return v.value
}

func (v *NullableAnalytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals) Set(val *Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals) {
	v.value = val
	v.isSet = true
}

func (v NullableAnalytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals) IsSet() bool {
	return v.isSet
}

func (v *NullableAnalytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnalytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals(val *Analytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals) *NullableAnalytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals {
	return &NullableAnalytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals{value: val, isSet: true}
}

func (v NullableAnalytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnalytics202405GetShopProductPerformanceResponseDataPerformanceComparisonIntervals) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


