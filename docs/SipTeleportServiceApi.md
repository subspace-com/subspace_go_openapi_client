# \SipTeleportServiceApi

All URIs are relative to *https://api.subspace.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SipTeleportServiceCreate**](SipTeleportServiceApi.md#SipTeleportServiceCreate) | **Post** /v1/sipteleport | 
[**SipTeleportServiceDelete**](SipTeleportServiceApi.md#SipTeleportServiceDelete) | **Delete** /v1/sipteleport/{id} | 
[**SipTeleportServiceGet**](SipTeleportServiceApi.md#SipTeleportServiceGet) | **Get** /v1/sipteleport/{id} | 
[**SipTeleportServiceList**](SipTeleportServiceApi.md#SipTeleportServiceList) | **Get** /v1/sipteleport | 
[**SipTeleportServiceUpdate**](SipTeleportServiceApi.md#SipTeleportServiceUpdate) | **Put** /v1/sipteleport/{id} | 



## SipTeleportServiceCreate

> V1SipTeleportResponse SipTeleportServiceCreate(ctx).V1CreateSipTeleport(v1CreateSipTeleport).IdempotencyKey(idempotencyKey).Execute()



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
    v1CreateSipTeleport := *openapiclient.NewV1CreateSipTeleport("Name_example", "Destination_example") // V1CreateSipTeleport | Required parameters to create a new SIPTeleport
    idempotencyKey := "idempotencyKey_example" // string | Value is the returned etag of a get request.  If a retry sends an Idempotency-Key that has been seen before, the existing teleport is returned with the status code of 200 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SipTeleportServiceApi.SipTeleportServiceCreate(context.Background()).V1CreateSipTeleport(v1CreateSipTeleport).IdempotencyKey(idempotencyKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SipTeleportServiceApi.SipTeleportServiceCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SipTeleportServiceCreate`: V1SipTeleportResponse
    fmt.Fprintf(os.Stdout, "Response from `SipTeleportServiceApi.SipTeleportServiceCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSipTeleportServiceCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1CreateSipTeleport** | [**V1CreateSipTeleport**](V1CreateSipTeleport.md) | Required parameters to create a new SIPTeleport | 
 **idempotencyKey** | **string** | Value is the returned etag of a get request.  If a retry sends an Idempotency-Key that has been seen before, the existing teleport is returned with the status code of 200 | 

### Return type

[**V1SipTeleportResponse**](V1SipTeleportResponse.md)

### Authorization

[accessCode](../README.md#accessCode)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SipTeleportServiceDelete

> V1SipTeleportResponse SipTeleportServiceDelete(ctx, id).Execute()



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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SipTeleportServiceApi.SipTeleportServiceDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SipTeleportServiceApi.SipTeleportServiceDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SipTeleportServiceDelete`: V1SipTeleportResponse
    fmt.Fprintf(os.Stdout, "Response from `SipTeleportServiceApi.SipTeleportServiceDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSipTeleportServiceDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1SipTeleportResponse**](V1SipTeleportResponse.md)

### Authorization

[accessCode](../README.md#accessCode)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SipTeleportServiceGet

> V1SipTeleportResponse SipTeleportServiceGet(ctx, id).Execute()



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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SipTeleportServiceApi.SipTeleportServiceGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SipTeleportServiceApi.SipTeleportServiceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SipTeleportServiceGet`: V1SipTeleportResponse
    fmt.Fprintf(os.Stdout, "Response from `SipTeleportServiceApi.SipTeleportServiceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSipTeleportServiceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1SipTeleportResponse**](V1SipTeleportResponse.md)

### Authorization

[accessCode](../README.md#accessCode)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SipTeleportServiceList

> V1ListSipTeleportResponse SipTeleportServiceList(ctx).Before(before).Limit(limit).Execute()



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
    before := "before_example" // string |  (optional)
    limit := int64(789) // int64 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SipTeleportServiceApi.SipTeleportServiceList(context.Background()).Before(before).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SipTeleportServiceApi.SipTeleportServiceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SipTeleportServiceList`: V1ListSipTeleportResponse
    fmt.Fprintf(os.Stdout, "Response from `SipTeleportServiceApi.SipTeleportServiceList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSipTeleportServiceListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **before** | **string** |  | 
 **limit** | **int64** |  | 

### Return type

[**V1ListSipTeleportResponse**](V1ListSipTeleportResponse.md)

### Authorization

[accessCode](../README.md#accessCode)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SipTeleportServiceUpdate

> V1SipTeleportResponse SipTeleportServiceUpdate(ctx, id).V1UpdateSipTeleport(v1UpdateSipTeleport).Execute()



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
    id := "id_example" // string | 
    v1UpdateSipTeleport := *openapiclient.NewV1UpdateSipTeleport("Name_example", "Destination_example", openapiclient.v1SipTeleportStatus("UNKNOWN")) // V1UpdateSipTeleport | Parameters to update an existing SIPTeleport, minimum requirement of one of them defined to update

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SipTeleportServiceApi.SipTeleportServiceUpdate(context.Background(), id).V1UpdateSipTeleport(v1UpdateSipTeleport).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SipTeleportServiceApi.SipTeleportServiceUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SipTeleportServiceUpdate`: V1SipTeleportResponse
    fmt.Fprintf(os.Stdout, "Response from `SipTeleportServiceApi.SipTeleportServiceUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSipTeleportServiceUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v1UpdateSipTeleport** | [**V1UpdateSipTeleport**](V1UpdateSipTeleport.md) | Parameters to update an existing SIPTeleport, minimum requirement of one of them defined to update | 

### Return type

[**V1SipTeleportResponse**](V1SipTeleportResponse.md)

### Authorization

[accessCode](../README.md#accessCode)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

