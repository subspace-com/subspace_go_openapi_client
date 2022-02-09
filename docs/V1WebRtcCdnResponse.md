# V1WebRtcCdnResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IceServers** | Pointer to [**[]V1WebRtcCdnServer**](V1WebRtcCdnServer.md) |  | [optional] 
**Ttl** | Pointer to **int64** |  | [optional] 

## Methods

### NewV1WebRtcCdnResponse

`func NewV1WebRtcCdnResponse() *V1WebRtcCdnResponse`

NewV1WebRtcCdnResponse instantiates a new V1WebRtcCdnResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1WebRtcCdnResponseWithDefaults

`func NewV1WebRtcCdnResponseWithDefaults() *V1WebRtcCdnResponse`

NewV1WebRtcCdnResponseWithDefaults instantiates a new V1WebRtcCdnResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIceServers

`func (o *V1WebRtcCdnResponse) GetIceServers() []V1WebRtcCdnServer`

GetIceServers returns the IceServers field if non-nil, zero value otherwise.

### GetIceServersOk

`func (o *V1WebRtcCdnResponse) GetIceServersOk() (*[]V1WebRtcCdnServer, bool)`

GetIceServersOk returns a tuple with the IceServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIceServers

`func (o *V1WebRtcCdnResponse) SetIceServers(v []V1WebRtcCdnServer)`

SetIceServers sets IceServers field to given value.

### HasIceServers

`func (o *V1WebRtcCdnResponse) HasIceServers() bool`

HasIceServers returns a boolean if a field has been set.

### GetTtl

`func (o *V1WebRtcCdnResponse) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *V1WebRtcCdnResponse) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *V1WebRtcCdnResponse) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *V1WebRtcCdnResponse) HasTtl() bool`

HasTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


