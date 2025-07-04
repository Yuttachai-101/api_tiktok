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

    "github.com/Yuttachai-101/api_tiktok/models/fulfillment/v202502"
)


// FulfillmentV202502APIService FulfillmentV202502API service
type FulfillmentV202502APIService service

type ApiFulfillment202502InvoiceUploadPostRequest struct {
    ctx context.Context
    ApiService *FulfillmentV202502APIService
    xTtsAccessToken *string
    contentType *string
    shopCipher *string
    fulfillment202502UploadInvoiceRequestBody *fulfillment_v202502.Fulfillment202502UploadInvoiceRequestBody
}

// 
func (r ApiFulfillment202502InvoiceUploadPostRequest) XTtsAccessToken(xTtsAccessToken string) ApiFulfillment202502InvoiceUploadPostRequest {
    r.xTtsAccessToken = &xTtsAccessToken
    return r
}
// Allowed type: application/json
func (r ApiFulfillment202502InvoiceUploadPostRequest) ContentType(contentType string) ApiFulfillment202502InvoiceUploadPostRequest {
    r.contentType = &contentType
    return r
}
// 
func (r ApiFulfillment202502InvoiceUploadPostRequest) ShopCipher(shopCipher string) ApiFulfillment202502InvoiceUploadPostRequest {
    r.shopCipher = &shopCipher
    return r
}
func (r ApiFulfillment202502InvoiceUploadPostRequest) Fulfillment202502UploadInvoiceRequestBody(fulfillment202502UploadInvoiceRequestBody fulfillment_v202502.Fulfillment202502UploadInvoiceRequestBody) ApiFulfillment202502InvoiceUploadPostRequest {
    r.fulfillment202502UploadInvoiceRequestBody = &fulfillment202502UploadInvoiceRequestBody
    return r
}
func (r ApiFulfillment202502InvoiceUploadPostRequest) Execute() (*fulfillment_v202502.Fulfillment202502UploadInvoiceResponse, *http.Response, error) {
    return r.ApiService.Fulfillment202502InvoiceUploadPostExecute(r)
}

/*
Fulfillment202502InvoiceUploadPost UploadInvoice
Upload the invoice document.
**Note**: Applicable only for local sellers in the Brazil market.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return ApiFulfillment202502InvoiceUploadPostRequest
*/
func (a *FulfillmentV202502APIService) Fulfillment202502InvoiceUploadPost(ctx context.Context) ApiFulfillment202502InvoiceUploadPostRequest {
    return ApiFulfillment202502InvoiceUploadPostRequest{
        ApiService: a,
        ctx: ctx,
    }
}

// Execute executes the request
//  @return Fulfillment202502UploadInvoiceResponse
func (a *FulfillmentV202502APIService) Fulfillment202502InvoiceUploadPostExecute(r ApiFulfillment202502InvoiceUploadPostRequest) (*fulfillment_v202502.Fulfillment202502UploadInvoiceResponse, *http.Response, error) {
    var (
        localVarHTTPMethod   = http.MethodPost
        localVarPostBody     interface{}
        formFiles            []formFile
        localVarReturnValue  *fulfillment_v202502.Fulfillment202502UploadInvoiceResponse
    )

    localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FulfillmentV202502APIService.Fulfillment202502InvoiceUploadPost")
    if err != nil {
        return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
    }

    localVarPath := localBasePath + "/fulfillment/202502/invoice/upload"

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := url.Values{}
    if r.xTtsAccessToken == nil {
        return localVarReturnValue, nil, reportError("xTtsAccessToken is required and must be specified")
    }
    if r.contentType == nil {
        return localVarReturnValue, nil, reportError("contentType is required and must be specified")
    }

    if r.shopCipher != nil {
    parameterAddToHeaderOrQuery(localVarQueryParams, "shop_cipher", r.shopCipher, "")
    }
    // to determine the Content-Type header
    localVarHTTPContentTypes := []string{"application/json"}

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
    // body params
    localVarPostBody = r.fulfillment202502UploadInvoiceRequestBody
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
