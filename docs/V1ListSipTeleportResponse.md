# V1ListSipTeleportResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]V1SipTeleportResponse**](V1SipTeleportResponse.md) |  | [optional] 
**NextPage** | Pointer to [**V1NextPage**](V1NextPage.md) |  | [optional] 

## Methods

### NewV1ListSipTeleportResponse

`func NewV1ListSipTeleportResponse() *V1ListSipTeleportResponse`

NewV1ListSipTeleportResponse instantiates a new V1ListSipTeleportResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ListSipTeleportResponseWithDefaults

`func NewV1ListSipTeleportResponseWithDefaults() *V1ListSipTeleportResponse`

NewV1ListSipTeleportResponseWithDefaults instantiates a new V1ListSipTeleportResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *V1ListSipTeleportResponse) GetData() []V1SipTeleportResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *V1ListSipTeleportResponse) GetDataOk() (*[]V1SipTeleportResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *V1ListSipTeleportResponse) SetData(v []V1SipTeleportResponse)`

SetData sets Data field to given value.

### HasData

`func (o *V1ListSipTeleportResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetNextPage

`func (o *V1ListSipTeleportResponse) GetNextPage() V1NextPage`

GetNextPage returns the NextPage field if non-nil, zero value otherwise.

### GetNextPageOk

`func (o *V1ListSipTeleportResponse) GetNextPageOk() (*V1NextPage, bool)`

GetNextPageOk returns a tuple with the NextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPage

`func (o *V1ListSipTeleportResponse) SetNextPage(v V1NextPage)`

SetNextPage sets NextPage field to given value.

### HasNextPage

`func (o *V1ListSipTeleportResponse) HasNextPage() bool`

HasNextPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


