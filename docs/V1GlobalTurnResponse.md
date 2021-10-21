# V1GlobalTurnResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IceServers** | Pointer to [**[]V1GlobalTurnServer**](V1GlobalTurnServer.md) |  | [optional] 
**Ttl** | Pointer to **int64** |  | [optional] 

## Methods

### NewV1GlobalTurnResponse

`func NewV1GlobalTurnResponse() *V1GlobalTurnResponse`

NewV1GlobalTurnResponse instantiates a new V1GlobalTurnResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1GlobalTurnResponseWithDefaults

`func NewV1GlobalTurnResponseWithDefaults() *V1GlobalTurnResponse`

NewV1GlobalTurnResponseWithDefaults instantiates a new V1GlobalTurnResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIceServers

`func (o *V1GlobalTurnResponse) GetIceServers() []V1GlobalTurnServer`

GetIceServers returns the IceServers field if non-nil, zero value otherwise.

### GetIceServersOk

`func (o *V1GlobalTurnResponse) GetIceServersOk() (*[]V1GlobalTurnServer, bool)`

GetIceServersOk returns a tuple with the IceServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIceServers

`func (o *V1GlobalTurnResponse) SetIceServers(v []V1GlobalTurnServer)`

SetIceServers sets IceServers field to given value.

### HasIceServers

`func (o *V1GlobalTurnResponse) HasIceServers() bool`

HasIceServers returns a boolean if a field has been set.

### GetTtl

`func (o *V1GlobalTurnResponse) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *V1GlobalTurnResponse) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *V1GlobalTurnResponse) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *V1GlobalTurnResponse) HasTtl() bool`

HasTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


