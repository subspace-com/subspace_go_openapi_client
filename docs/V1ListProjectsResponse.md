# V1ListProjectsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Projects** | Pointer to [**[]V1Project**](V1Project.md) |  | [optional] 
**NextPage** | Pointer to [**V1NextPage**](V1NextPage.md) |  | [optional] 

## Methods

### NewV1ListProjectsResponse

`func NewV1ListProjectsResponse() *V1ListProjectsResponse`

NewV1ListProjectsResponse instantiates a new V1ListProjectsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ListProjectsResponseWithDefaults

`func NewV1ListProjectsResponseWithDefaults() *V1ListProjectsResponse`

NewV1ListProjectsResponseWithDefaults instantiates a new V1ListProjectsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjects

`func (o *V1ListProjectsResponse) GetProjects() []V1Project`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *V1ListProjectsResponse) GetProjectsOk() (*[]V1Project, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *V1ListProjectsResponse) SetProjects(v []V1Project)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *V1ListProjectsResponse) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetNextPage

`func (o *V1ListProjectsResponse) GetNextPage() V1NextPage`

GetNextPage returns the NextPage field if non-nil, zero value otherwise.

### GetNextPageOk

`func (o *V1ListProjectsResponse) GetNextPageOk() (*V1NextPage, bool)`

GetNextPageOk returns a tuple with the NextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPage

`func (o *V1ListProjectsResponse) SetNextPage(v V1NextPage)`

SetNextPage sets NextPage field to given value.

### HasNextPage

`func (o *V1ListProjectsResponse) HasNextPage() bool`

HasNextPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


