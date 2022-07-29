# InterfaceConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InterfaceA** | Pointer to [**NestedInterface**](NestedInterface.md) |  | [optional] 
**InterfaceB** | [**NestedInterface**](NestedInterface.md) |  | 
**ConnectedEndpointReachable** | Pointer to **bool** |  | [optional] [readonly] 

## Methods

### NewInterfaceConnection

`func NewInterfaceConnection(interfaceB NestedInterface, ) *InterfaceConnection`

NewInterfaceConnection instantiates a new InterfaceConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterfaceConnectionWithDefaults

`func NewInterfaceConnectionWithDefaults() *InterfaceConnection`

NewInterfaceConnectionWithDefaults instantiates a new InterfaceConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterfaceA

`func (o *InterfaceConnection) GetInterfaceA() NestedInterface`

GetInterfaceA returns the InterfaceA field if non-nil, zero value otherwise.

### GetInterfaceAOk

`func (o *InterfaceConnection) GetInterfaceAOk() (*NestedInterface, bool)`

GetInterfaceAOk returns a tuple with the InterfaceA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceA

`func (o *InterfaceConnection) SetInterfaceA(v NestedInterface)`

SetInterfaceA sets InterfaceA field to given value.

### HasInterfaceA

`func (o *InterfaceConnection) HasInterfaceA() bool`

HasInterfaceA returns a boolean if a field has been set.

### GetInterfaceB

`func (o *InterfaceConnection) GetInterfaceB() NestedInterface`

GetInterfaceB returns the InterfaceB field if non-nil, zero value otherwise.

### GetInterfaceBOk

`func (o *InterfaceConnection) GetInterfaceBOk() (*NestedInterface, bool)`

GetInterfaceBOk returns a tuple with the InterfaceB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceB

`func (o *InterfaceConnection) SetInterfaceB(v NestedInterface)`

SetInterfaceB sets InterfaceB field to given value.


### GetConnectedEndpointReachable

`func (o *InterfaceConnection) GetConnectedEndpointReachable() bool`

GetConnectedEndpointReachable returns the ConnectedEndpointReachable field if non-nil, zero value otherwise.

### GetConnectedEndpointReachableOk

`func (o *InterfaceConnection) GetConnectedEndpointReachableOk() (*bool, bool)`

GetConnectedEndpointReachableOk returns a tuple with the ConnectedEndpointReachable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedEndpointReachable

`func (o *InterfaceConnection) SetConnectedEndpointReachable(v bool)`

SetConnectedEndpointReachable sets ConnectedEndpointReachable field to given value.

### HasConnectedEndpointReachable

`func (o *InterfaceConnection) HasConnectedEndpointReachable() bool`

HasConnectedEndpointReachable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


