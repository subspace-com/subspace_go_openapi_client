# V1UpdateSipTeleport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of SIPTeleport | 
**Destination** | **string** | Destination of SIPTeleport | 
**Status** | [**V1SipTeleportStatus**](V1SipTeleportStatus.md) |  | [default to DISABLED]

## Methods

### NewV1UpdateSipTeleport

`func NewV1UpdateSipTeleport(name string, destination string, status V1SipTeleportStatus, ) *V1UpdateSipTeleport`

NewV1UpdateSipTeleport instantiates a new V1UpdateSipTeleport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1UpdateSipTeleportWithDefaults

`func NewV1UpdateSipTeleportWithDefaults() *V1UpdateSipTeleport`

NewV1UpdateSipTeleportWithDefaults instantiates a new V1UpdateSipTeleport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *V1UpdateSipTeleport) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1UpdateSipTeleport) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1UpdateSipTeleport) SetName(v string)`

SetName sets Name field to given value.


### GetDestination

`func (o *V1UpdateSipTeleport) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *V1UpdateSipTeleport) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *V1UpdateSipTeleport) SetDestination(v string)`

SetDestination sets Destination field to given value.


### GetStatus

`func (o *V1UpdateSipTeleport) GetStatus() V1SipTeleportStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V1UpdateSipTeleport) GetStatusOk() (*V1SipTeleportStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V1UpdateSipTeleport) SetStatus(v V1SipTeleportStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


