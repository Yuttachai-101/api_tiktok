/*
tiktok shop openapi

sdk for apis

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apis

import (
    "bytes"
    "context"
    "io"
    "net/http"
    "net/url"
    "strings"

    "github.com/Yuttachai-101/api_tiktok/models/analytics/v202409"
)


// AnalyticsV202409APIService AnalyticsV202409API service
type AnalyticsV202409APIService service

type ApiAnalytics202409ShopVideosOverviewPerformanceGetRequest struct {
    ctx context.Context
    ApiService *AnalyticsV202409APIService
    startDateGe *string
    endDateLt *string
    xTtsAccessToken *string
    contentType *string
    withComparison *bool
    granularity *string
    currency *string
    accountType *string
    shopCipher *string
}

// Start date (ISO 8601 YYYY-MM-DD format) in shop registered timezone. In the parameter name, \&quot;ge\&quot; refers to \&quot;greater than or equal to\&quot; (inclusive)
func (r ApiAnalytics202409ShopVideosOverviewPerformanceGetRequest) StartDateGe(startDateGe string) ApiAnalytics202409ShopVideosOverviewPerformanceGetRequest {
    r.startDateGe = &startDateGe
    return r
}
// End date (ISO 8601 YYYY-MM-DD format) in shop registered timezone. In the parameter name, \&quot;lt\&quot; refers to \&quot;less than\&quot; (exclusive)
func (r ApiAnalytics202409ShopVideosOverviewPerformanceGetRequest) EndDateLt(endDateLt string) ApiAnalytics202409ShopVideosOverviewPerformanceGetRequest {
    r.endDateLt = &endDateLt
    return r
}
// 
func (r ApiAnalytics202409ShopVideosOverviewPerformanceGetRequest) XTtsAccessToken(xTtsAccessToken string) ApiAnalytics202409ShopVideosOverviewPerformanceGetRequest {
    r.xTtsAccessToken = &xTtsAccessToken
    return r
}
// Allowed type: application/json
func (r ApiAnalytics202409ShopVideosOverviewPerformanceGetRequest) ContentType(contentType string) ApiAnalytics202409ShopVideosOverviewPerformanceGetRequest {
    r.contentType = &contentType
    return r
}
// Whether previous period data is returned for comparison. Available values: true, false Default value: false The previous period has the same length and granularity as the current period with end time being the same as the start time of the current period. Example: If start_time_ge &#x3D; 2024-04-01 and end_time_lt &#x3D; 2024-04-08, the previous period data will be from 2024-03-25 to 2024-04-01.
func (r ApiAnalytics202409ShopVideosOverviewPerformanceGetRequest) WithComparison(withComparison bool) ApiAnalytics202409ShopVideosOverviewPerformanceGetRequest {
    r.withComparison = &withComparison
    return r
}
// Granularity of the data. Available values: ALL, 1D Default value: ALL * ALL: aggregate * 1D: daily
func (r ApiAnalytics202409ShopVideosOverviewPerformanceGetRequest) Granularity(granularity string) ApiAnalytics202409ShopVideosOverviewPerformanceGetRequest {
    r.granularity = &granularity
    return r
}
// Currency. Available values: USD, LOCAL Default value: LOCAL * USD: US dollars * LOCAL: local currency where the shop is located
func (r ApiAnalytics202409ShopVideosOverviewPerformanceGetRequest) Currency(currency string) ApiAnalytics202409ShopVideosOverviewPerformanceGetRequest {
    r.currency = &currency
    return r
}
// Types of the accounts under which videos were created. Available values: ALL, LINKED_ACCOUNTS, AFFILIATES Default value: ALL * ALL: all account types * LINKED_ACCOUNTS: linked account types, includes official and marketing account types * AFFILIATES: affiliate account type
func (r ApiAnalytics202409ShopVideosOverviewPerformanceGetRequest) AccountType(accountType string) ApiAnalytics202409ShopVideosOverviewPerformanceGetRequest {
    r.accountType = &accountType
    return r
}
// 
func (r ApiAnalytics202409ShopVideosOverviewPerformanceGetRequest) ShopCipher(shopCipher string) ApiAnalytics202409ShopVideosOverviewPerformanceGetRequest {
    r.shopCipher = &shopCipher
    return r
}
func (r ApiAnalytics202409ShopVideosOverviewPerformanceGetRequest) Execute() (*analytics_v202409.Analytics202409GetShopVideoPerformanceOverviewResponse, *http.Response, error) {
    return r.ApiService.Analytics202409ShopVideosOverviewPerformanceGetExecute(r)
}

/*
Analytics202409ShopVideosOverviewPerformanceGet GetShopVideoPerformanceOverview
Returns overall performance metrics for all videos under a shop. This API currently provides data only for shops registered in the United States, United Kingdom, Singapore, Vietnam, and Thailand.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return ApiAnalytics202409ShopVideosOverviewPerformanceGetRequest
*/
func (a *AnalyticsV202409APIService) Analytics202409ShopVideosOverviewPerformanceGet(ctx context.Context) ApiAnalytics202409ShopVideosOverviewPerformanceGetRequest {
    return ApiAnalytics202409ShopVideosOverviewPerformanceGetRequest{
        ApiService: a,
        ctx: ctx,
    }
}

// Execute executes the request
//  @return Analytics202409GetShopVideoPerformanceOverviewResponse
func (a *AnalyticsV202409APIService) Analytics202409ShopVideosOverviewPerformanceGetExecute(r ApiAnalytics202409ShopVideosOverviewPerformanceGetRequest) (*analytics_v202409.Analytics202409GetShopVideoPerformanceOverviewResponse, *http.Response, error) {
    var (
        localVarHTTPMethod   = http.MethodGet
        localVarPostBody     interface{}
        formFiles            []formFile
        localVarReturnValue  *analytics_v202409.Analytics202409GetShopVideoPerformanceOverviewResponse
    )

    localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnalyticsV202409APIService.Analytics202409ShopVideosOverviewPerformanceGet")
    if err != nil {
        return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
    }

    localVarPath := localBasePath + "/analytics/202409/shop_videos/overview_performance"

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := url.Values{}
    if r.startDateGe == nil {
        return localVarReturnValue, nil, reportError("startDateGe is required and must be specified")
    }
    if r.endDateLt == nil {
        return localVarReturnValue, nil, reportError("endDateLt is required and must be specified")
    }
    if r.xTtsAccessToken == nil {
        return localVarReturnValue, nil, reportError("xTtsAccessToken is required and must be specified")
    }
    if r.contentType == nil {
        return localVarReturnValue, nil, reportError("contentType is required and must be specified")
    }

    parameterAddToHeaderOrQuery(localVarQueryParams, "start_date_ge", r.startDateGe, "")
    parameterAddToHeaderOrQuery(localVarQueryParams, "end_date_lt", r.endDateLt, "")
    if r.withComparison != nil {
    parameterAddToHeaderOrQuery(localVarQueryParams, "with_comparison", r.withComparison, "")
    }
    if r.granularity != nil {
    parameterAddToHeaderOrQuery(localVarQueryParams, "granularity", r.granularity, "")
    }
    if r.currency != nil {
    parameterAddToHeaderOrQuery(localVarQueryParams, "currency", r.currency, "")
    }
    if r.accountType != nil {
    parameterAddToHeaderOrQuery(localVarQueryParams, "account_type", r.accountType, "")
    }
    if r.shopCipher != nil {
    parameterAddToHeaderOrQuery(localVarQueryParams, "shop_cipher", r.shopCipher, "")
    }
    // to determine the Content-Type header
    localVarHTTPContentTypes := []string{}

    // set Content-Type header
    localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
    if localVarHTTPContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHTTPContentType
    }

    // to determine the Accept header
    localVarHTTPHeaderAccepts := []string{"application/json"}

    // set Accept header
    localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
    if localVarHTTPHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
    }
    parameterAddToHeaderOrQuery(localVarHeaderParams, "x-tts-access-token", r.xTtsAccessToken, "")
    parameterAddToHeaderOrQuery(localVarHeaderParams, "Content-Type", r.contentType, "")
    req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
    if err != nil {
        return localVarReturnValue, nil, err
    }

    localVarHTTPResponse, err := a.client.callAPI(req)
    if err != nil || localVarHTTPResponse == nil {
        return localVarReturnValue, localVarHTTPResponse, err
    }

    localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
    localVarHTTPResponse.Body.Close()
    localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
    if err != nil {
        return localVarReturnValue, localVarHTTPResponse, err
    }

    if localVarHTTPResponse.StatusCode >= 300 {
        newErr := &GenericOpenAPIError{
            body:  localVarBody,
            error: localVarHTTPResponse.Status,
        }
        return localVarReturnValue, localVarHTTPResponse, newErr
    }

    err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
    if err != nil {
        newErr := &GenericOpenAPIError{
            body:  localVarBody,
            error: err.Error(),
        }
        return localVarReturnValue, localVarHTTPResponse, newErr
    }

    return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiAnalytics202409ShopVideosPerformanceGetRequest struct {
    ctx context.Context
    ApiService *AnalyticsV202409APIService
    startDateGe *string
    endDateLt *string
    xTtsAccessToken *string
    contentType *string
    pageSize *int32
    sortField *string
    sortOrder *string
    currency *string
    pageToken *string
    accountType *string
    shopCipher *string
}

// Start date (ISO 8601 YYYY-MM-DD format) in shop registered timezone. In the parameter name, \&quot;ge\&quot; refers to \&quot;greater than or equal to\&quot; (inclusive)
func (r ApiAnalytics202409ShopVideosPerformanceGetRequest) StartDateGe(startDateGe string) ApiAnalytics202409ShopVideosPerformanceGetRequest {
    r.startDateGe = &startDateGe
    return r
}
// End date (ISO 8601 YYYY-MM-DD format) in shop registered timezone. In the parameter name, \&quot;lt\&quot; refers to \&quot;less than\&quot; (exclusive)
func (r ApiAnalytics202409ShopVideosPerformanceGetRequest) EndDateLt(endDateLt string) ApiAnalytics202409ShopVideosPerformanceGetRequest {
    r.endDateLt = &endDateLt
    return r
}
// 
func (r ApiAnalytics202409ShopVideosPerformanceGetRequest) XTtsAccessToken(xTtsAccessToken string) ApiAnalytics202409ShopVideosPerformanceGetRequest {
    r.xTtsAccessToken = &xTtsAccessToken
    return r
}
// Allowed type: application/json
func (r ApiAnalytics202409ShopVideosPerformanceGetRequest) ContentType(contentType string) ApiAnalytics202409ShopVideosPerformanceGetRequest {
    r.contentType = &contentType
    return r
}
// Number of products per page. Max value: 100 Default value: 10
func (r ApiAnalytics202409ShopVideosPerformanceGetRequest) PageSize(pageSize int32) ApiAnalytics202409ShopVideosPerformanceGetRequest {
    r.pageSize = &pageSize
    return r
}
// Sort on. Available values: gmv, sku_orders, units_sold, views, click_through_rate Default value: gmv
func (r ApiAnalytics202409ShopVideosPerformanceGetRequest) SortField(sortField string) ApiAnalytics202409ShopVideosPerformanceGetRequest {
    r.sortField = &sortField
    return r
}
// Sort direction. Available values: ASC, DESC Default value: DESC * ASC: ascending * DESC: descending
func (r ApiAnalytics202409ShopVideosPerformanceGetRequest) SortOrder(sortOrder string) ApiAnalytics202409ShopVideosPerformanceGetRequest {
    r.sortOrder = &sortOrder
    return r
}
// Currency. Available values: USD, LOCAL Default value: LOCAL * USD: US dollars * LOCAL: local currency where the shop is located
func (r ApiAnalytics202409ShopVideosPerformanceGetRequest) Currency(currency string) ApiAnalytics202409ShopVideosPerformanceGetRequest {
    r.currency = &currency
    return r
}
// Page token, indicating the current position. Used for requesting next page data. Leave this field empty for first time queries.
func (r ApiAnalytics202409ShopVideosPerformanceGetRequest) PageToken(pageToken string) ApiAnalytics202409ShopVideosPerformanceGetRequest {
    r.pageToken = &pageToken
    return r
}
// Types of the accounts under which videos were created. Available values: ALL, LINKED_ACCOUNTS, AFFILIATES Default value: ALL * ALL: all account types * LINKED_ACCOUNTS: linked account types, includes official and marketing account types * AFFILIATES: affiliate account type
func (r ApiAnalytics202409ShopVideosPerformanceGetRequest) AccountType(accountType string) ApiAnalytics202409ShopVideosPerformanceGetRequest {
    r.accountType = &accountType
    return r
}
// 
func (r ApiAnalytics202409ShopVideosPerformanceGetRequest) ShopCipher(shopCipher string) ApiAnalytics202409ShopVideosPerformanceGetRequest {
    r.shopCipher = &shopCipher
    return r
}
func (r ApiAnalytics202409ShopVideosPerformanceGetRequest) Execute() (*analytics_v202409.Analytics202409GetShopVideoPerformanceListResponse, *http.Response, error) {
    return r.ApiService.Analytics202409ShopVideosPerformanceGetExecute(r)
}

/*
Analytics202409ShopVideosPerformanceGet GetShopVideoPerformanceList
Returns a list of videos and associated metrics for a shop. This API currently provides data only for shops registered in the United States, United Kingdom, Singapore, Vietnam, and Thailand.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return ApiAnalytics202409ShopVideosPerformanceGetRequest
*/
func (a *AnalyticsV202409APIService) Analytics202409ShopVideosPerformanceGet(ctx context.Context) ApiAnalytics202409ShopVideosPerformanceGetRequest {
    return ApiAnalytics202409ShopVideosPerformanceGetRequest{
        ApiService: a,
        ctx: ctx,
    }
}

// Execute executes the request
//  @return Analytics202409GetShopVideoPerformanceListResponse
func (a *AnalyticsV202409APIService) Analytics202409ShopVideosPerformanceGetExecute(r ApiAnalytics202409ShopVideosPerformanceGetRequest) (*analytics_v202409.Analytics202409GetShopVideoPerformanceListResponse, *http.Response, error) {
    var (
        localVarHTTPMethod   = http.MethodGet
        localVarPostBody     interface{}
        formFiles            []formFile
        localVarReturnValue  *analytics_v202409.Analytics202409GetShopVideoPerformanceListResponse
    )

    localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnalyticsV202409APIService.Analytics202409ShopVideosPerformanceGet")
    if err != nil {
        return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
    }

    localVarPath := localBasePath + "/analytics/202409/shop_videos/performance"

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := url.Values{}
    if r.startDateGe == nil {
        return localVarReturnValue, nil, reportError("startDateGe is required and must be specified")
    }
    if r.endDateLt == nil {
        return localVarReturnValue, nil, reportError("endDateLt is required and must be specified")
    }
    if r.xTtsAccessToken == nil {
        return localVarReturnValue, nil, reportError("xTtsAccessToken is required and must be specified")
    }
    if r.contentType == nil {
        return localVarReturnValue, nil, reportError("contentType is required and must be specified")
    }

    parameterAddToHeaderOrQuery(localVarQueryParams, "start_date_ge", r.startDateGe, "")
    parameterAddToHeaderOrQuery(localVarQueryParams, "end_date_lt", r.endDateLt, "")
    if r.pageSize != nil {
    parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize, "")
    }
    if r.sortField != nil {
    parameterAddToHeaderOrQuery(localVarQueryParams, "sort_field", r.sortField, "")
    }
    if r.sortOrder != nil {
    parameterAddToHeaderOrQuery(localVarQueryParams, "sort_order", r.sortOrder, "")
    }
    if r.currency != nil {
    parameterAddToHeaderOrQuery(localVarQueryParams, "currency", r.currency, "")
    }
    if r.pageToken != nil {
    parameterAddToHeaderOrQuery(localVarQueryParams, "page_token", r.pageToken, "")
    }
    if r.accountType != nil {
    parameterAddToHeaderOrQuery(localVarQueryParams, "account_type", r.accountType, "")
    }
    if r.shopCipher != nil {
    parameterAddToHeaderOrQuery(localVarQueryParams, "shop_cipher", r.shopCipher, "")
    }
    // to determine the Content-Type header
    localVarHTTPContentTypes := []string{}

    // set Content-Type header
    localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
    if localVarHTTPContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHTTPContentType
    }

    // to determine the Accept header
    localVarHTTPHeaderAccepts := []string{"application/json"}

    // set Accept header
    localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
    if localVarHTTPHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
    }
    parameterAddToHeaderOrQuery(localVarHeaderParams, "x-tts-access-token", r.xTtsAccessToken, "")
    parameterAddToHeaderOrQuery(localVarHeaderParams, "Content-Type", r.contentType, "")
    req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
    if err != nil {
        return localVarReturnValue, nil, err
    }

    localVarHTTPResponse, err := a.client.callAPI(req)
    if err != nil || localVarHTTPResponse == nil {
        return localVarReturnValue, localVarHTTPResponse, err
    }

    localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
    localVarHTTPResponse.Body.Close()
    localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
    if err != nil {
        return localVarReturnValue, localVarHTTPResponse, err
    }

    if localVarHTTPResponse.StatusCode >= 300 {
        newErr := &GenericOpenAPIError{
            body:  localVarBody,
            error: localVarHTTPResponse.Status,
        }
        return localVarReturnValue, localVarHTTPResponse, newErr
    }

    err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
    if err != nil {
        newErr := &GenericOpenAPIError{
            body:  localVarBody,
            error: err.Error(),
        }
        return localVarReturnValue, localVarHTTPResponse, newErr
    }

    return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiAnalytics202409ShopVideosVideoIdPerformanceGetRequest struct {
    ctx context.Context
    ApiService *AnalyticsV202409APIService
    videoId string
    startDateGe *string
    endDateLt *string
    xTtsAccessToken *string
    contentType *string
    withComparison *bool
    granularity *string
    currency *string
    shopCipher *string
}

// Start date (ISO 8601 YYYY-MM-DD format) in shop registered timezone. In the parameter name, \&quot;ge\&quot; refers to \&quot;greater than or equal to\&quot; (inclusive)
func (r ApiAnalytics202409ShopVideosVideoIdPerformanceGetRequest) StartDateGe(startDateGe string) ApiAnalytics202409ShopVideosVideoIdPerformanceGetRequest {
    r.startDateGe = &startDateGe
    return r
}
// End date (ISO 8601 YYYY-MM-DD format) in shop registered timezone. In the parameter name, \&quot;lt\&quot; refers to \&quot;less than\&quot; (exclusive)
func (r ApiAnalytics202409ShopVideosVideoIdPerformanceGetRequest) EndDateLt(endDateLt string) ApiAnalytics202409ShopVideosVideoIdPerformanceGetRequest {
    r.endDateLt = &endDateLt
    return r
}
// 
func (r ApiAnalytics202409ShopVideosVideoIdPerformanceGetRequest) XTtsAccessToken(xTtsAccessToken string) ApiAnalytics202409ShopVideosVideoIdPerformanceGetRequest {
    r.xTtsAccessToken = &xTtsAccessToken
    return r
}
// Allowed type: application/json
func (r ApiAnalytics202409ShopVideosVideoIdPerformanceGetRequest) ContentType(contentType string) ApiAnalytics202409ShopVideosVideoIdPerformanceGetRequest {
    r.contentType = &contentType
    return r
}
// Whether previous period data is returned for comparison. Available values: true, false Default value: false The previous period has the same length and granularity as the current period with end time being the same as the start time of the current period. Example: If start_time_ge &#x3D; 2024-04-01 and end_time_lt &#x3D; 2024-04-08, the previous period data will be from 2024-03-25 to 2024-04-01.
func (r ApiAnalytics202409ShopVideosVideoIdPerformanceGetRequest) WithComparison(withComparison bool) ApiAnalytics202409ShopVideosVideoIdPerformanceGetRequest {
    r.withComparison = &withComparison
    return r
}
// Granularity of the data. Available values: ALL, 1D Default value: ALL * ALL: aggregate * 1D: daily
func (r ApiAnalytics202409ShopVideosVideoIdPerformanceGetRequest) Granularity(granularity string) ApiAnalytics202409ShopVideosVideoIdPerformanceGetRequest {
    r.granularity = &granularity
    return r
}
// Currency. Available values: USD, LOCAL Default value: LOCAL * USD: US dollars * LOCAL: local currency where the shop is located
func (r ApiAnalytics202409ShopVideosVideoIdPerformanceGetRequest) Currency(currency string) ApiAnalytics202409ShopVideosVideoIdPerformanceGetRequest {
    r.currency = &currency
    return r
}
// 
func (r ApiAnalytics202409ShopVideosVideoIdPerformanceGetRequest) ShopCipher(shopCipher string) ApiAnalytics202409ShopVideosVideoIdPerformanceGetRequest {
    r.shopCipher = &shopCipher
    return r
}
func (r ApiAnalytics202409ShopVideosVideoIdPerformanceGetRequest) Execute() (*analytics_v202409.Analytics202409GetShopVideoPerformanceDetailsResponse, *http.Response, error) {
    return r.ApiService.Analytics202409ShopVideosVideoIdPerformanceGetExecute(r)
}

/*
Analytics202409ShopVideosVideoIdPerformanceGet GetShopVideoPerformanceDetails
Returns detailed performance metrics for a (requested) video. This API currently provides data only for shops registered in the United States, United Kingdom, Singapore, Vietnam, and Thailand.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param videoId Video ID
@return ApiAnalytics202409ShopVideosVideoIdPerformanceGetRequest
*/
func (a *AnalyticsV202409APIService) Analytics202409ShopVideosVideoIdPerformanceGet(ctx context.Context, videoId string) ApiAnalytics202409ShopVideosVideoIdPerformanceGetRequest {
    return ApiAnalytics202409ShopVideosVideoIdPerformanceGetRequest{
        ApiService: a,
        ctx: ctx,
        videoId: videoId,
    }
}

// Execute executes the request
//  @return Analytics202409GetShopVideoPerformanceDetailsResponse
func (a *AnalyticsV202409APIService) Analytics202409ShopVideosVideoIdPerformanceGetExecute(r ApiAnalytics202409ShopVideosVideoIdPerformanceGetRequest) (*analytics_v202409.Analytics202409GetShopVideoPerformanceDetailsResponse, *http.Response, error) {
    var (
        localVarHTTPMethod   = http.MethodGet
        localVarPostBody     interface{}
        formFiles            []formFile
        localVarReturnValue  *analytics_v202409.Analytics202409GetShopVideoPerformanceDetailsResponse
    )

    localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnalyticsV202409APIService.Analytics202409ShopVideosVideoIdPerformanceGet")
    if err != nil {
        return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
    }

    localVarPath := localBasePath + "/analytics/202409/shop_videos/{video_id}/performance"
    localVarPath = strings.Replace(localVarPath, "{"+"video_id"+"}", url.PathEscape(parameterValueToString(r.videoId, "videoId")), -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := url.Values{}
    if r.startDateGe == nil {
        return localVarReturnValue, nil, reportError("startDateGe is required and must be specified")
    }
    if r.endDateLt == nil {
        return localVarReturnValue, nil, reportError("endDateLt is required and must be specified")
    }
    if r.xTtsAccessToken == nil {
        return localVarReturnValue, nil, reportError("xTtsAccessToken is required and must be specified")
    }
    if r.contentType == nil {
        return localVarReturnValue, nil, reportError("contentType is required and must be specified")
    }

    parameterAddToHeaderOrQuery(localVarQueryParams, "start_date_ge", r.startDateGe, "")
    parameterAddToHeaderOrQuery(localVarQueryParams, "end_date_lt", r.endDateLt, "")
    if r.withComparison != nil {
    parameterAddToHeaderOrQuery(localVarQueryParams, "with_comparison", r.withComparison, "")
    }
    if r.granularity != nil {
    parameterAddToHeaderOrQuery(localVarQueryParams, "granularity", r.granularity, "")
    }
    if r.currency != nil {
    parameterAddToHeaderOrQuery(localVarQueryParams, "currency", r.currency, "")
    }
    if r.shopCipher != nil {
    parameterAddToHeaderOrQuery(localVarQueryParams, "shop_cipher", r.shopCipher, "")
    }
    // to determine the Content-Type header
    localVarHTTPContentTypes := []string{}

    // set Content-Type header
    localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
    if localVarHTTPContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHTTPContentType
    }

    // to determine the Accept header
    localVarHTTPHeaderAccepts := []string{"application/json"}

    // set Accept header
    localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
    if localVarHTTPHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
    }
    parameterAddToHeaderOrQuery(localVarHeaderParams, "x-tts-access-token", r.xTtsAccessToken, "")
    parameterAddToHeaderOrQuery(localVarHeaderParams, "Content-Type", r.contentType, "")
    req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
    if err != nil {
        return localVarReturnValue, nil, err
    }

    localVarHTTPResponse, err := a.client.callAPI(req)
    if err != nil || localVarHTTPResponse == nil {
        return localVarReturnValue, localVarHTTPResponse, err
    }

    localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
    localVarHTTPResponse.Body.Close()
    localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
    if err != nil {
        return localVarReturnValue, localVarHTTPResponse, err
    }

    if localVarHTTPResponse.StatusCode >= 300 {
        newErr := &GenericOpenAPIError{
            body:  localVarBody,
            error: localVarHTTPResponse.Status,
        }
        return localVarReturnValue, localVarHTTPResponse, newErr
    }

    err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
    if err != nil {
        newErr := &GenericOpenAPIError{
            body:  localVarBody,
            error: err.Error(),
        }
        return localVarReturnValue, localVarHTTPResponse, newErr
    }

    return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiAnalytics202409ShopVideosVideoIdProductsPerformanceGetRequest struct {
    ctx context.Context
    ApiService *AnalyticsV202409APIService
    videoId string
    endDateLt *string
    startDateGe *string
    xTtsAccessToken *string
    contentType *string
    pageSize *int32
    sortField *string
    sortOrder *string
    currency *string
    pageToken *string
    shopCipher *string
}

// End date (ISO 8601 YYYY-MM-DD format) in shop registered timezone. In the parameter name, \&quot;lt\&quot; refers to \&quot;less than\&quot; (exclusive)
func (r ApiAnalytics202409ShopVideosVideoIdProductsPerformanceGetRequest) EndDateLt(endDateLt string) ApiAnalytics202409ShopVideosVideoIdProductsPerformanceGetRequest {
    r.endDateLt = &endDateLt
    return r
}
// Start date (ISO 8601 YYYY-MM-DD format) in shop registered timezone. In the parameter name, \&quot;ge\&quot; refers to \&quot;greater than or equal to\&quot; (inclusive)
func (r ApiAnalytics202409ShopVideosVideoIdProductsPerformanceGetRequest) StartDateGe(startDateGe string) ApiAnalytics202409ShopVideosVideoIdProductsPerformanceGetRequest {
    r.startDateGe = &startDateGe
    return r
}
// 
func (r ApiAnalytics202409ShopVideosVideoIdProductsPerformanceGetRequest) XTtsAccessToken(xTtsAccessToken string) ApiAnalytics202409ShopVideosVideoIdProductsPerformanceGetRequest {
    r.xTtsAccessToken = &xTtsAccessToken
    return r
}
// Allowed type: application/json
func (r ApiAnalytics202409ShopVideosVideoIdProductsPerformanceGetRequest) ContentType(contentType string) ApiAnalytics202409ShopVideosVideoIdProductsPerformanceGetRequest {
    r.contentType = &contentType
    return r
}
// Number of products per page. Max value: 100 Default value: 10
func (r ApiAnalytics202409ShopVideosVideoIdProductsPerformanceGetRequest) PageSize(pageSize int32) ApiAnalytics202409ShopVideosVideoIdProductsPerformanceGetRequest {
    r.pageSize = &pageSize
    return r
}
// Sort on. Available values: gmv, units_sold, daily_avg_buyers Default value: gmv
func (r ApiAnalytics202409ShopVideosVideoIdProductsPerformanceGetRequest) SortField(sortField string) ApiAnalytics202409ShopVideosVideoIdProductsPerformanceGetRequest {
    r.sortField = &sortField
    return r
}
// Sort direction. Available values: ASC, DESC Default value: DESC * ASC: ascending * DESC: descending
func (r ApiAnalytics202409ShopVideosVideoIdProductsPerformanceGetRequest) SortOrder(sortOrder string) ApiAnalytics202409ShopVideosVideoIdProductsPerformanceGetRequest {
    r.sortOrder = &sortOrder
    return r
}
// Currency. Available values: USD, LOCAL Default value: LOCAL * USD: US dollars * LOCAL: local currency where the shop is located
func (r ApiAnalytics202409ShopVideosVideoIdProductsPerformanceGetRequest) Currency(currency string) ApiAnalytics202409ShopVideosVideoIdProductsPerformanceGetRequest {
    r.currency = &currency
    return r
}
// Page token, indicating the current position. Used for requesting next page data. Leave this field empty for first time queries.
func (r ApiAnalytics202409ShopVideosVideoIdProductsPerformanceGetRequest) PageToken(pageToken string) ApiAnalytics202409ShopVideosVideoIdProductsPerformanceGetRequest {
    r.pageToken = &pageToken
    return r
}
// 
func (r ApiAnalytics202409ShopVideosVideoIdProductsPerformanceGetRequest) ShopCipher(shopCipher string) ApiAnalytics202409ShopVideosVideoIdProductsPerformanceGetRequest {
    r.shopCipher = &shopCipher
    return r
}
func (r ApiAnalytics202409ShopVideosVideoIdProductsPerformanceGetRequest) Execute() (*analytics_v202409.Analytics202409GetShopVideoProductPerformanceListResponse, *http.Response, error) {
    return r.ApiService.Analytics202409ShopVideosVideoIdProductsPerformanceGetExecute(r)
}

/*
Analytics202409ShopVideosVideoIdProductsPerformanceGet GetShopVideoProductPerformanceList
Returns performance metrics for a list of products promoted in a given video. This API currently provides data only for shops registered in the United States, United Kingdom, Singapore, Vietnam, and Thailand.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param videoId 
@return ApiAnalytics202409ShopVideosVideoIdProductsPerformanceGetRequest
*/
func (a *AnalyticsV202409APIService) Analytics202409ShopVideosVideoIdProductsPerformanceGet(ctx context.Context, videoId string) ApiAnalytics202409ShopVideosVideoIdProductsPerformanceGetRequest {
    return ApiAnalytics202409ShopVideosVideoIdProductsPerformanceGetRequest{
        ApiService: a,
        ctx: ctx,
        videoId: videoId,
    }
}

// Execute executes the request
//  @return Analytics202409GetShopVideoProductPerformanceListResponse
func (a *AnalyticsV202409APIService) Analytics202409ShopVideosVideoIdProductsPerformanceGetExecute(r ApiAnalytics202409ShopVideosVideoIdProductsPerformanceGetRequest) (*analytics_v202409.Analytics202409GetShopVideoProductPerformanceListResponse, *http.Response, error) {
    var (
        localVarHTTPMethod   = http.MethodGet
        localVarPostBody     interface{}
        formFiles            []formFile
        localVarReturnValue  *analytics_v202409.Analytics202409GetShopVideoProductPerformanceListResponse
    )

    localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnalyticsV202409APIService.Analytics202409ShopVideosVideoIdProductsPerformanceGet")
    if err != nil {
        return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
    }

    localVarPath := localBasePath + "/analytics/202409/shop_videos/{video_id}/products/performance"
    localVarPath = strings.Replace(localVarPath, "{"+"video_id"+"}", url.PathEscape(parameterValueToString(r.videoId, "videoId")), -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := url.Values{}
    if r.endDateLt == nil {
        return localVarReturnValue, nil, reportError("endDateLt is required and must be specified")
    }
    if r.startDateGe == nil {
        return localVarReturnValue, nil, reportError("startDateGe is required and must be specified")
    }
    if r.xTtsAccessToken == nil {
        return localVarReturnValue, nil, reportError("xTtsAccessToken is required and must be specified")
    }
    if r.contentType == nil {
        return localVarReturnValue, nil, reportError("contentType is required and must be specified")
    }

    parameterAddToHeaderOrQuery(localVarQueryParams, "end_date_lt", r.endDateLt, "")
    if r.pageSize != nil {
    parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize, "")
    }
    if r.sortField != nil {
    parameterAddToHeaderOrQuery(localVarQueryParams, "sort_field", r.sortField, "")
    }
    if r.sortOrder != nil {
    parameterAddToHeaderOrQuery(localVarQueryParams, "sort_order", r.sortOrder, "")
    }
    if r.currency != nil {
    parameterAddToHeaderOrQuery(localVarQueryParams, "currency", r.currency, "")
    }
    if r.pageToken != nil {
    parameterAddToHeaderOrQuery(localVarQueryParams, "page_token", r.pageToken, "")
    }
    parameterAddToHeaderOrQuery(localVarQueryParams, "start_date_ge", r.startDateGe, "")
    if r.shopCipher != nil {
    parameterAddToHeaderOrQuery(localVarQueryParams, "shop_cipher", r.shopCipher, "")
    }
    // to determine the Content-Type header
    localVarHTTPContentTypes := []string{}

    // set Content-Type header
    localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
    if localVarHTTPContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHTTPContentType
    }

    // to determine the Accept header
    localVarHTTPHeaderAccepts := []string{"application/json"}

    // set Accept header
    localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
    if localVarHTTPHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
    }
    parameterAddToHeaderOrQuery(localVarHeaderParams, "x-tts-access-token", r.xTtsAccessToken, "")
    parameterAddToHeaderOrQuery(localVarHeaderParams, "Content-Type", r.contentType, "")
    req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
    if err != nil {
        return localVarReturnValue, nil, err
    }

    localVarHTTPResponse, err := a.client.callAPI(req)
    if err != nil || localVarHTTPResponse == nil {
        return localVarReturnValue, localVarHTTPResponse, err
    }

    localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
    localVarHTTPResponse.Body.Close()
    localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
    if err != nil {
        return localVarReturnValue, localVarHTTPResponse, err
    }

    if localVarHTTPResponse.StatusCode >= 300 {
        newErr := &GenericOpenAPIError{
            body:  localVarBody,
            error: localVarHTTPResponse.Status,
        }
        return localVarReturnValue, localVarHTTPResponse, newErr
    }

    err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
    if err != nil {
        newErr := &GenericOpenAPIError{
            body:  localVarBody,
            error: err.Error(),
        }
        return localVarReturnValue, localVarHTTPResponse, newErr
    }

    return localVarReturnValue, localVarHTTPResponse, nil
}
