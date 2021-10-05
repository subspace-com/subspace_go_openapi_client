# V1ListSessionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]V1Session**](V1Session.md) |  | [optional] 
**NextPage** | Pointer to [**V1NextPage**](V1NextPage.md) |  | [optional] 

## Methods

### NewV1ListSessionsResponse

`func NewV1ListSessionsResponse() *V1ListSessionsResponse`

NewV1ListSessionsResponse instantiates a new V1ListSessionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ListSessionsResponseWithDefaults

`func NewV1ListSessionsResponseWithDefaults() *V1ListSessionsResponse`

NewV1ListSessionsResponseWithDefaults instantiates a new V1ListSessionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *V1ListSessionsResponse) GetData() []V1Session`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *V1ListSessionsResponse) GetDataOk() (*[]V1Session, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *V1ListSessionsResponse) SetData(v []V1Session)`

SetData sets Data field to given value.

### HasData

`func (o *V1ListSessionsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetNextPage

`func (o *V1ListSessionsResponse) GetNextPage() V1NextPage`

GetNextPage returns the NextPage field if non-nil, zero value otherwise.

### GetNextPageOk

`func (o *V1ListSessionsResponse) GetNextPageOk() (*V1NextPage, bool)`

GetNextPageOk returns a tuple with the NextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPage

`func (o *V1ListSessionsResponse) SetNextPage(v V1NextPage)`

SetNextPage sets NextPage field to given value.

### HasNextPage

`func (o *V1ListSessionsResponse) HasNextPage() bool`

HasNextPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


