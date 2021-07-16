# Body

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | **string** | udp or tcp | 
**Name** | **string** | Name of PacketAccelerator | 
**DestinationIp** | **string** | Destination IP, e.g. 1.2.3.4 | 
**DestinationPort** | **int32** | Destination port, e.g. 443 | 
**SubspacePort** | Pointer to **int32** | Desired Subspace-assigned ingress port, optional | [optional] 

## Methods

### NewBody

`func NewBody(protocol string, name string, destinationIp string, destinationPort int32, ) *Body`

NewBody instantiates a new Body object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBodyWithDefaults

`func NewBodyWithDefaults() *Body`

NewBodyWithDefaults instantiates a new Body object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *Body) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *Body) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *Body) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetName

`func (o *Body) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Body) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Body) SetName(v string)`

SetName sets Name field to given value.


### GetDestinationIp

`func (o *Body) GetDestinationIp() string`

GetDestinationIp returns the DestinationIp field if non-nil, zero value otherwise.

### GetDestinationIpOk

`func (o *Body) GetDestinationIpOk() (*string, bool)`

GetDestinationIpOk returns a tuple with the DestinationIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationIp

`func (o *Body) SetDestinationIp(v string)`

SetDestinationIp sets DestinationIp field to given value.


### GetDestinationPort

`func (o *Body) GetDestinationPort() int32`

GetDestinationPort returns the DestinationPort field if non-nil, zero value otherwise.

### GetDestinationPortOk

`func (o *Body) GetDestinationPortOk() (*int32, bool)`

GetDestinationPortOk returns a tuple with the DestinationPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPort

`func (o *Body) SetDestinationPort(v int32)`

SetDestinationPort sets DestinationPort field to given value.


### GetSubspacePort

`func (o *Body) GetSubspacePort() int32`

GetSubspacePort returns the SubspacePort field if non-nil, zero value otherwise.

### GetSubspacePortOk

`func (o *Body) GetSubspacePortOk() (*int32, bool)`

GetSubspacePortOk returns a tuple with the SubspacePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubspacePort

`func (o *Body) SetSubspacePort(v int32)`

SetSubspacePort sets SubspacePort field to given value.

### HasSubspacePort

`func (o *Body) HasSubspacePort() bool`

HasSubspacePort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


