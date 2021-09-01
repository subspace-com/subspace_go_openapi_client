# V1Project

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**AcceleratorQuota** | Pointer to **int64** |  | [optional] 
**AcceleratorRequestPort** | Pointer to **bool** |  | [optional] 
**GlobalturnLimitGb** | Pointer to **int64** |  | [optional] 
**SipteleportQuota** | Pointer to **int64** |  | [optional] 
**SipteleportCallQuota** | Pointer to **int64** |  | [optional] 
**RtpspeedLimitGb** | Pointer to **int64** |  | [optional] 

## Methods

### NewV1Project

`func NewV1Project() *V1Project`

NewV1Project instantiates a new V1Project object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ProjectWithDefaults

`func NewV1ProjectWithDefaults() *V1Project`

NewV1ProjectWithDefaults instantiates a new V1Project object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *V1Project) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V1Project) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V1Project) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *V1Project) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *V1Project) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1Project) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1Project) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1Project) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAcceleratorQuota

`func (o *V1Project) GetAcceleratorQuota() int64`

GetAcceleratorQuota returns the AcceleratorQuota field if non-nil, zero value otherwise.

### GetAcceleratorQuotaOk

`func (o *V1Project) GetAcceleratorQuotaOk() (*int64, bool)`

GetAcceleratorQuotaOk returns a tuple with the AcceleratorQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceleratorQuota

`func (o *V1Project) SetAcceleratorQuota(v int64)`

SetAcceleratorQuota sets AcceleratorQuota field to given value.

### HasAcceleratorQuota

`func (o *V1Project) HasAcceleratorQuota() bool`

HasAcceleratorQuota returns a boolean if a field has been set.

### GetAcceleratorRequestPort

`func (o *V1Project) GetAcceleratorRequestPort() bool`

GetAcceleratorRequestPort returns the AcceleratorRequestPort field if non-nil, zero value otherwise.

### GetAcceleratorRequestPortOk

`func (o *V1Project) GetAcceleratorRequestPortOk() (*bool, bool)`

GetAcceleratorRequestPortOk returns a tuple with the AcceleratorRequestPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceleratorRequestPort

`func (o *V1Project) SetAcceleratorRequestPort(v bool)`

SetAcceleratorRequestPort sets AcceleratorRequestPort field to given value.

### HasAcceleratorRequestPort

`func (o *V1Project) HasAcceleratorRequestPort() bool`

HasAcceleratorRequestPort returns a boolean if a field has been set.

### GetGlobalturnLimitGb

`func (o *V1Project) GetGlobalturnLimitGb() int64`

GetGlobalturnLimitGb returns the GlobalturnLimitGb field if non-nil, zero value otherwise.

### GetGlobalturnLimitGbOk

`func (o *V1Project) GetGlobalturnLimitGbOk() (*int64, bool)`

GetGlobalturnLimitGbOk returns a tuple with the GlobalturnLimitGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalturnLimitGb

`func (o *V1Project) SetGlobalturnLimitGb(v int64)`

SetGlobalturnLimitGb sets GlobalturnLimitGb field to given value.

### HasGlobalturnLimitGb

`func (o *V1Project) HasGlobalturnLimitGb() bool`

HasGlobalturnLimitGb returns a boolean if a field has been set.

### GetSipteleportQuota

`func (o *V1Project) GetSipteleportQuota() int64`

GetSipteleportQuota returns the SipteleportQuota field if non-nil, zero value otherwise.

### GetSipteleportQuotaOk

`func (o *V1Project) GetSipteleportQuotaOk() (*int64, bool)`

GetSipteleportQuotaOk returns a tuple with the SipteleportQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipteleportQuota

`func (o *V1Project) SetSipteleportQuota(v int64)`

SetSipteleportQuota sets SipteleportQuota field to given value.

### HasSipteleportQuota

`func (o *V1Project) HasSipteleportQuota() bool`

HasSipteleportQuota returns a boolean if a field has been set.

### GetSipteleportCallQuota

`func (o *V1Project) GetSipteleportCallQuota() int64`

GetSipteleportCallQuota returns the SipteleportCallQuota field if non-nil, zero value otherwise.

### GetSipteleportCallQuotaOk

`func (o *V1Project) GetSipteleportCallQuotaOk() (*int64, bool)`

GetSipteleportCallQuotaOk returns a tuple with the SipteleportCallQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipteleportCallQuota

`func (o *V1Project) SetSipteleportCallQuota(v int64)`

SetSipteleportCallQuota sets SipteleportCallQuota field to given value.

### HasSipteleportCallQuota

`func (o *V1Project) HasSipteleportCallQuota() bool`

HasSipteleportCallQuota returns a boolean if a field has been set.

### GetRtpspeedLimitGb

`func (o *V1Project) GetRtpspeedLimitGb() int64`

GetRtpspeedLimitGb returns the RtpspeedLimitGb field if non-nil, zero value otherwise.

### GetRtpspeedLimitGbOk

`func (o *V1Project) GetRtpspeedLimitGbOk() (*int64, bool)`

GetRtpspeedLimitGbOk returns a tuple with the RtpspeedLimitGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRtpspeedLimitGb

`func (o *V1Project) SetRtpspeedLimitGb(v int64)`

SetRtpspeedLimitGb sets RtpspeedLimitGb field to given value.

### HasRtpspeedLimitGb

`func (o *V1Project) HasRtpspeedLimitGb() bool`

HasRtpspeedLimitGb returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


