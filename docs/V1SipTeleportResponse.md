# V1SipTeleportResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Destination** | Pointer to **string** |  | [optional] 
**TeleportEntryPoints** | Pointer to [**[]V1TeleportAddresses**](V1TeleportAddresses.md) |  | [optional] 
**Status** | Pointer to [**V1SipTeleportStatus**](V1SipTeleportStatus.md) |  | [optional] [default to UNKNOWN]
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**DateUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewV1SipTeleportResponse

`func NewV1SipTeleportResponse() *V1SipTeleportResponse`

NewV1SipTeleportResponse instantiates a new V1SipTeleportResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1SipTeleportResponseWithDefaults

`func NewV1SipTeleportResponseWithDefaults() *V1SipTeleportResponse`

NewV1SipTeleportResponseWithDefaults instantiates a new V1SipTeleportResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *V1SipTeleportResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V1SipTeleportResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V1SipTeleportResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *V1SipTeleportResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *V1SipTeleportResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1SipTeleportResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1SipTeleportResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1SipTeleportResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDestination

`func (o *V1SipTeleportResponse) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *V1SipTeleportResponse) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *V1SipTeleportResponse) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *V1SipTeleportResponse) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetTeleportEntryPoints

`func (o *V1SipTeleportResponse) GetTeleportEntryPoints() []V1TeleportAddresses`

GetTeleportEntryPoints returns the TeleportEntryPoints field if non-nil, zero value otherwise.

### GetTeleportEntryPointsOk

`func (o *V1SipTeleportResponse) GetTeleportEntryPointsOk() (*[]V1TeleportAddresses, bool)`

GetTeleportEntryPointsOk returns a tuple with the TeleportEntryPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeleportEntryPoints

`func (o *V1SipTeleportResponse) SetTeleportEntryPoints(v []V1TeleportAddresses)`

SetTeleportEntryPoints sets TeleportEntryPoints field to given value.

### HasTeleportEntryPoints

`func (o *V1SipTeleportResponse) HasTeleportEntryPoints() bool`

HasTeleportEntryPoints returns a boolean if a field has been set.

### GetStatus

`func (o *V1SipTeleportResponse) GetStatus() V1SipTeleportStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V1SipTeleportResponse) GetStatusOk() (*V1SipTeleportStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V1SipTeleportResponse) SetStatus(v V1SipTeleportStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V1SipTeleportResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDateCreated

`func (o *V1SipTeleportResponse) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *V1SipTeleportResponse) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *V1SipTeleportResponse) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *V1SipTeleportResponse) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetDateUpdated

`func (o *V1SipTeleportResponse) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *V1SipTeleportResponse) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *V1SipTeleportResponse) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *V1SipTeleportResponse) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


