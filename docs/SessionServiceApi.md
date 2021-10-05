# \SessionServiceApi

All URIs are relative to *https://api.subspace.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SessionServiceList**](SessionServiceApi.md#SessionServiceList) | **Get** /v1/accelerator/{accelerator_id}/session | 



## SessionServiceList

> V1ListSessionsResponse SessionServiceList(ctx, acceleratorId).Before(before).Limit(limit).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    acceleratorId := "acceleratorId_example" // string | 
    before := "before_example" // string |  (optional)
    limit := int64(789) // int64 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SessionServiceApi.SessionServiceList(context.Background(), acceleratorId).Before(before).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionServiceApi.SessionServiceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SessionServiceList`: V1ListSessionsResponse
    fmt.Fprintf(os.Stdout, "Response from `SessionServiceApi.SessionServiceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**acceleratorId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSessionServiceListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **before** | **string** |  | 
 **limit** | **int64** |  | 

### Return type

[**V1ListSessionsResponse**](V1ListSessionsResponse.md)

### Authorization

[accessCode](../README.md#accessCode)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

