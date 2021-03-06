# \AcceleratorServiceApi

All URIs are relative to *https://api.subspace.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcceleratorServiceCreate**](AcceleratorServiceApi.md#AcceleratorServiceCreate) | **Post** /v1/accelerator | 
[**AcceleratorServiceDelete**](AcceleratorServiceApi.md#AcceleratorServiceDelete) | **Delete** /v1/accelerator/{id} | 
[**AcceleratorServiceGet**](AcceleratorServiceApi.md#AcceleratorServiceGet) | **Get** /v1/accelerator/{id} | 
[**AcceleratorServiceList**](AcceleratorServiceApi.md#AcceleratorServiceList) | **Get** /v1/accelerator | 
[**AcceleratorServiceUpdate**](AcceleratorServiceApi.md#AcceleratorServiceUpdate) | **Put** /v1/accelerator/{id} | 



## AcceleratorServiceCreate

> V1Accelerator AcceleratorServiceCreate(ctx).Body(body).IdempotencyKey(idempotencyKey).Execute()



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
    body := *openapiclient.NewBody("Name_example", "DestinationIp_example", int32(123)) // Body | Required parameters to create a new PacketAccelerator.
    idempotencyKey := "idempotencyKey_example" // string | Value is the returned etag of a get request.  If a retry sends an Idempotency-Key that has been seen before, the existing accelerator is returned with the status code of 200 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AcceleratorServiceApi.AcceleratorServiceCreate(context.Background()).Body(body).IdempotencyKey(idempotencyKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcceleratorServiceApi.AcceleratorServiceCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcceleratorServiceCreate`: V1Accelerator
    fmt.Fprintf(os.Stdout, "Response from `AcceleratorServiceApi.AcceleratorServiceCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAcceleratorServiceCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Body**](Body.md) | Required parameters to create a new PacketAccelerator. | 
 **idempotencyKey** | **string** | Value is the returned etag of a get request.  If a retry sends an Idempotency-Key that has been seen before, the existing accelerator is returned with the status code of 200 | 

### Return type

[**V1Accelerator**](V1Accelerator.md)

### Authorization

[accessCode](../README.md#accessCode)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AcceleratorServiceDelete

> map[string]interface{} AcceleratorServiceDelete(ctx, id).Execute()



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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AcceleratorServiceApi.AcceleratorServiceDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcceleratorServiceApi.AcceleratorServiceDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcceleratorServiceDelete`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AcceleratorServiceApi.AcceleratorServiceDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcceleratorServiceDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[accessCode](../README.md#accessCode)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AcceleratorServiceGet

> V1Accelerator AcceleratorServiceGet(ctx, id).Execute()



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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AcceleratorServiceApi.AcceleratorServiceGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcceleratorServiceApi.AcceleratorServiceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcceleratorServiceGet`: V1Accelerator
    fmt.Fprintf(os.Stdout, "Response from `AcceleratorServiceApi.AcceleratorServiceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcceleratorServiceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1Accelerator**](V1Accelerator.md)

### Authorization

[accessCode](../README.md#accessCode)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AcceleratorServiceList

> V1ListAcceleratorResponse AcceleratorServiceList(ctx).Before(before).Limit(limit).Name(name).Execute()



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
    name := "name_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AcceleratorServiceApi.AcceleratorServiceList(context.Background()).Before(before).Limit(limit).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcceleratorServiceApi.AcceleratorServiceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcceleratorServiceList`: V1ListAcceleratorResponse
    fmt.Fprintf(os.Stdout, "Response from `AcceleratorServiceApi.AcceleratorServiceList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAcceleratorServiceListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **before** | **string** |  | 
 **limit** | **int64** |  | 
 **name** | **string** |  | 

### Return type

[**V1ListAcceleratorResponse**](V1ListAcceleratorResponse.md)

### Authorization

[accessCode](../README.md#accessCode)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AcceleratorServiceUpdate

> V1Accelerator AcceleratorServiceUpdate(ctx, id).Body1(body1).IfMatch(ifMatch).Execute()



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
    body1 := *openapiclient.NewBody1() // Body1 | Parameters to update an existing PacketAccelerator
    ifMatch := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AcceleratorServiceApi.AcceleratorServiceUpdate(context.Background(), id).Body1(body1).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcceleratorServiceApi.AcceleratorServiceUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcceleratorServiceUpdate`: V1Accelerator
    fmt.Fprintf(os.Stdout, "Response from `AcceleratorServiceApi.AcceleratorServiceUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcceleratorServiceUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body1** | [**Body1**](Body1.md) | Parameters to update an existing PacketAccelerator | 
 **ifMatch** | **int32** |  | 

### Return type

[**V1Accelerator**](V1Accelerator.md)

### Authorization

[accessCode](../README.md#accessCode)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

