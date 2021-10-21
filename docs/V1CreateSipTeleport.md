# V1CreateSipTeleport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of SIPTeleport | 
**Destination** | **string** | Destination of SIPTeleport | 
**Status** | Pointer to **string** | Enum: [ ENABLED, DISABLED ] | [optional] 

## Methods

### NewV1CreateSipTeleport

`func NewV1CreateSipTeleport(name string, destination string, ) *V1CreateSipTeleport`

NewV1CreateSipTeleport instantiates a new V1CreateSipTeleport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1CreateSipTeleportWithDefaults

`func NewV1CreateSipTeleportWithDefaults() *V1CreateSipTeleport`

NewV1CreateSipTeleportWithDefaults instantiates a new V1CreateSipTeleport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *V1CreateSipTeleport) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1CreateSipTeleport) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1CreateSipTeleport) SetName(v string)`

SetName sets Name field to given value.


### GetDestination

`func (o *V1CreateSipTeleport) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *V1CreateSipTeleport) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *V1CreateSipTeleport) SetDestination(v string)`

SetDestination sets Destination field to given value.


### GetStatus

`func (o *V1CreateSipTeleport) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V1CreateSipTeleport) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V1CreateSipTeleport) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V1CreateSipTeleport) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


