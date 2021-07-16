# V1ListAcceleratorsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accelerators** | Pointer to [**[]V1Accelerator**](V1Accelerator.md) |  | [optional] 
**NextPage** | Pointer to [**V1NextPage**](V1NextPage.md) |  | [optional] 

## Methods

### NewV1ListAcceleratorsResponse

`func NewV1ListAcceleratorsResponse() *V1ListAcceleratorsResponse`

NewV1ListAcceleratorsResponse instantiates a new V1ListAcceleratorsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ListAcceleratorsResponseWithDefaults

`func NewV1ListAcceleratorsResponseWithDefaults() *V1ListAcceleratorsResponse`

NewV1ListAcceleratorsResponseWithDefaults instantiates a new V1ListAcceleratorsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccelerators

`func (o *V1ListAcceleratorsResponse) GetAccelerators() []V1Accelerator`

GetAccelerators returns the Accelerators field if non-nil, zero value otherwise.

### GetAcceleratorsOk

`func (o *V1ListAcceleratorsResponse) GetAcceleratorsOk() (*[]V1Accelerator, bool)`

GetAcceleratorsOk returns a tuple with the Accelerators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccelerators

`func (o *V1ListAcceleratorsResponse) SetAccelerators(v []V1Accelerator)`

SetAccelerators sets Accelerators field to given value.

### HasAccelerators

`func (o *V1ListAcceleratorsResponse) HasAccelerators() bool`

HasAccelerators returns a boolean if a field has been set.

### GetNextPage

`func (o *V1ListAcceleratorsResponse) GetNextPage() V1NextPage`

GetNextPage returns the NextPage field if non-nil, zero value otherwise.

### GetNextPageOk

`func (o *V1ListAcceleratorsResponse) GetNextPageOk() (*V1NextPage, bool)`

GetNextPageOk returns a tuple with the NextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPage

`func (o *V1ListAcceleratorsResponse) SetNextPage(v V1NextPage)`

SetNextPage sets NextPage field to given value.

### HasNextPage

`func (o *V1ListAcceleratorsResponse) HasNextPage() bool`

HasNextPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


