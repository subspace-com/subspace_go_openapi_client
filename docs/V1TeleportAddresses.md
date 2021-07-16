# V1TeleportAddresses

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** |  | [optional] 
**TransportType** | Pointer to [**V1TransportType**](V1TransportType.md) |  | [optional] [default to UDP_TCP]

## Methods

### NewV1TeleportAddresses

`func NewV1TeleportAddresses() *V1TeleportAddresses`

NewV1TeleportAddresses instantiates a new V1TeleportAddresses object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1TeleportAddressesWithDefaults

`func NewV1TeleportAddressesWithDefaults() *V1TeleportAddresses`

NewV1TeleportAddressesWithDefaults instantiates a new V1TeleportAddresses object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *V1TeleportAddresses) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *V1TeleportAddresses) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *V1TeleportAddresses) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *V1TeleportAddresses) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetTransportType

`func (o *V1TeleportAddresses) GetTransportType() V1TransportType`

GetTransportType returns the TransportType field if non-nil, zero value otherwise.

### GetTransportTypeOk

`func (o *V1TeleportAddresses) GetTransportTypeOk() (*V1TransportType, bool)`

GetTransportTypeOk returns a tuple with the TransportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportType

`func (o *V1TeleportAddresses) SetTransportType(v V1TransportType)`

SetTransportType sets TransportType field to given value.

### HasTransportType

`func (o *V1TeleportAddresses) HasTransportType() bool`

HasTransportType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


