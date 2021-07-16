# Body1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | Pointer to **string** | udp or tcp | [optional] 
**Name** | Pointer to **string** | Name of PacketAccelerator | [optional] 
**DestinationIp** | Pointer to **string** | Destination IP, e.g. 1.2.3.4 | [optional] 
**DestinationPort** | Pointer to **int32** | Destination port, e.g. 443 | [optional] 

## Methods

### NewBody1

`func NewBody1() *Body1`

NewBody1 instantiates a new Body1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBody1WithDefaults

`func NewBody1WithDefaults() *Body1`

NewBody1WithDefaults instantiates a new Body1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *Body1) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *Body1) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *Body1) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *Body1) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetName

`func (o *Body1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Body1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Body1) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Body1) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDestinationIp

`func (o *Body1) GetDestinationIp() string`

GetDestinationIp returns the DestinationIp field if non-nil, zero value otherwise.

### GetDestinationIpOk

`func (o *Body1) GetDestinationIpOk() (*string, bool)`

GetDestinationIpOk returns a tuple with the DestinationIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationIp

`func (o *Body1) SetDestinationIp(v string)`

SetDestinationIp sets DestinationIp field to given value.

### HasDestinationIp

`func (o *Body1) HasDestinationIp() bool`

HasDestinationIp returns a boolean if a field has been set.

### GetDestinationPort

`func (o *Body1) GetDestinationPort() int32`

GetDestinationPort returns the DestinationPort field if non-nil, zero value otherwise.

### GetDestinationPortOk

`func (o *Body1) GetDestinationPortOk() (*int32, bool)`

GetDestinationPortOk returns a tuple with the DestinationPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPort

`func (o *Body1) SetDestinationPort(v int32)`

SetDestinationPort sets DestinationPort field to given value.

### HasDestinationPort

`func (o *Body1) HasDestinationPort() bool`

HasDestinationPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


