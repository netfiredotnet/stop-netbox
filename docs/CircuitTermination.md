# CircuitTermination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Display** | Pointer to **string** |  | [optional] [readonly] 
**Circuit** | [**NestedCircuit**](NestedCircuit.md) |  | 
**TermSide** | **string** |  | 
**Site** | Pointer to [**NestedSite**](NestedSite.md) |  | [optional] 
**ProviderNetwork** | Pointer to [**NestedProviderNetwork**](NestedProviderNetwork.md) |  | [optional] 
**PortSpeed** | Pointer to **NullableInt32** |  | [optional] 
**UpstreamSpeed** | Pointer to **NullableInt32** | Upstream speed, if different from port speed | [optional] 
**XconnectId** | Pointer to **string** |  | [optional] 
**PpInfo** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**MarkConnected** | Pointer to **bool** | Treat as if a cable is connected | [optional] 
**Cable** | Pointer to [**NestedCable**](NestedCable.md) |  | [optional] 
**CablePeer** | Pointer to **map[string]string** |  Return the appropriate serializer for the cable termination model.  | [optional] [readonly] 
**CablePeerType** | Pointer to **string** |  | [optional] [readonly] 
**Occupied** | Pointer to **bool** |  | [optional] [readonly] 

## Methods

### NewCircuitTermination

`func NewCircuitTermination(circuit NestedCircuit, termSide string, ) *CircuitTermination`

NewCircuitTermination instantiates a new CircuitTermination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCircuitTerminationWithDefaults

`func NewCircuitTerminationWithDefaults() *CircuitTermination`

NewCircuitTerminationWithDefaults instantiates a new CircuitTermination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CircuitTermination) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CircuitTermination) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CircuitTermination) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CircuitTermination) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *CircuitTermination) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CircuitTermination) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CircuitTermination) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CircuitTermination) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDisplay

`func (o *CircuitTermination) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *CircuitTermination) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *CircuitTermination) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *CircuitTermination) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetCircuit

`func (o *CircuitTermination) GetCircuit() NestedCircuit`

GetCircuit returns the Circuit field if non-nil, zero value otherwise.

### GetCircuitOk

`func (o *CircuitTermination) GetCircuitOk() (*NestedCircuit, bool)`

GetCircuitOk returns a tuple with the Circuit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuit

`func (o *CircuitTermination) SetCircuit(v NestedCircuit)`

SetCircuit sets Circuit field to given value.


### GetTermSide

`func (o *CircuitTermination) GetTermSide() string`

GetTermSide returns the TermSide field if non-nil, zero value otherwise.

### GetTermSideOk

`func (o *CircuitTermination) GetTermSideOk() (*string, bool)`

GetTermSideOk returns a tuple with the TermSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermSide

`func (o *CircuitTermination) SetTermSide(v string)`

SetTermSide sets TermSide field to given value.


### GetSite

`func (o *CircuitTermination) GetSite() NestedSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *CircuitTermination) GetSiteOk() (*NestedSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *CircuitTermination) SetSite(v NestedSite)`

SetSite sets Site field to given value.

### HasSite

`func (o *CircuitTermination) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetProviderNetwork

`func (o *CircuitTermination) GetProviderNetwork() NestedProviderNetwork`

GetProviderNetwork returns the ProviderNetwork field if non-nil, zero value otherwise.

### GetProviderNetworkOk

`func (o *CircuitTermination) GetProviderNetworkOk() (*NestedProviderNetwork, bool)`

GetProviderNetworkOk returns a tuple with the ProviderNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderNetwork

`func (o *CircuitTermination) SetProviderNetwork(v NestedProviderNetwork)`

SetProviderNetwork sets ProviderNetwork field to given value.

### HasProviderNetwork

`func (o *CircuitTermination) HasProviderNetwork() bool`

HasProviderNetwork returns a boolean if a field has been set.

### GetPortSpeed

`func (o *CircuitTermination) GetPortSpeed() int32`

GetPortSpeed returns the PortSpeed field if non-nil, zero value otherwise.

### GetPortSpeedOk

`func (o *CircuitTermination) GetPortSpeedOk() (*int32, bool)`

GetPortSpeedOk returns a tuple with the PortSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortSpeed

`func (o *CircuitTermination) SetPortSpeed(v int32)`

SetPortSpeed sets PortSpeed field to given value.

### HasPortSpeed

`func (o *CircuitTermination) HasPortSpeed() bool`

HasPortSpeed returns a boolean if a field has been set.

### SetPortSpeedNil

`func (o *CircuitTermination) SetPortSpeedNil(b bool)`

 SetPortSpeedNil sets the value for PortSpeed to be an explicit nil

### UnsetPortSpeed
`func (o *CircuitTermination) UnsetPortSpeed()`

UnsetPortSpeed ensures that no value is present for PortSpeed, not even an explicit nil
### GetUpstreamSpeed

`func (o *CircuitTermination) GetUpstreamSpeed() int32`

GetUpstreamSpeed returns the UpstreamSpeed field if non-nil, zero value otherwise.

### GetUpstreamSpeedOk

`func (o *CircuitTermination) GetUpstreamSpeedOk() (*int32, bool)`

GetUpstreamSpeedOk returns a tuple with the UpstreamSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamSpeed

`func (o *CircuitTermination) SetUpstreamSpeed(v int32)`

SetUpstreamSpeed sets UpstreamSpeed field to given value.

### HasUpstreamSpeed

`func (o *CircuitTermination) HasUpstreamSpeed() bool`

HasUpstreamSpeed returns a boolean if a field has been set.

### SetUpstreamSpeedNil

`func (o *CircuitTermination) SetUpstreamSpeedNil(b bool)`

 SetUpstreamSpeedNil sets the value for UpstreamSpeed to be an explicit nil

### UnsetUpstreamSpeed
`func (o *CircuitTermination) UnsetUpstreamSpeed()`

UnsetUpstreamSpeed ensures that no value is present for UpstreamSpeed, not even an explicit nil
### GetXconnectId

`func (o *CircuitTermination) GetXconnectId() string`

GetXconnectId returns the XconnectId field if non-nil, zero value otherwise.

### GetXconnectIdOk

`func (o *CircuitTermination) GetXconnectIdOk() (*string, bool)`

GetXconnectIdOk returns a tuple with the XconnectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXconnectId

`func (o *CircuitTermination) SetXconnectId(v string)`

SetXconnectId sets XconnectId field to given value.

### HasXconnectId

`func (o *CircuitTermination) HasXconnectId() bool`

HasXconnectId returns a boolean if a field has been set.

### GetPpInfo

`func (o *CircuitTermination) GetPpInfo() string`

GetPpInfo returns the PpInfo field if non-nil, zero value otherwise.

### GetPpInfoOk

`func (o *CircuitTermination) GetPpInfoOk() (*string, bool)`

GetPpInfoOk returns a tuple with the PpInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpInfo

`func (o *CircuitTermination) SetPpInfo(v string)`

SetPpInfo sets PpInfo field to given value.

### HasPpInfo

`func (o *CircuitTermination) HasPpInfo() bool`

HasPpInfo returns a boolean if a field has been set.

### GetDescription

`func (o *CircuitTermination) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CircuitTermination) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CircuitTermination) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CircuitTermination) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMarkConnected

`func (o *CircuitTermination) GetMarkConnected() bool`

GetMarkConnected returns the MarkConnected field if non-nil, zero value otherwise.

### GetMarkConnectedOk

`func (o *CircuitTermination) GetMarkConnectedOk() (*bool, bool)`

GetMarkConnectedOk returns a tuple with the MarkConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkConnected

`func (o *CircuitTermination) SetMarkConnected(v bool)`

SetMarkConnected sets MarkConnected field to given value.

### HasMarkConnected

`func (o *CircuitTermination) HasMarkConnected() bool`

HasMarkConnected returns a boolean if a field has been set.

### GetCable

`func (o *CircuitTermination) GetCable() NestedCable`

GetCable returns the Cable field if non-nil, zero value otherwise.

### GetCableOk

`func (o *CircuitTermination) GetCableOk() (*NestedCable, bool)`

GetCableOk returns a tuple with the Cable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCable

`func (o *CircuitTermination) SetCable(v NestedCable)`

SetCable sets Cable field to given value.

### HasCable

`func (o *CircuitTermination) HasCable() bool`

HasCable returns a boolean if a field has been set.

### GetCablePeer

`func (o *CircuitTermination) GetCablePeer() map[string]string`

GetCablePeer returns the CablePeer field if non-nil, zero value otherwise.

### GetCablePeerOk

`func (o *CircuitTermination) GetCablePeerOk() (*map[string]string, bool)`

GetCablePeerOk returns a tuple with the CablePeer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCablePeer

`func (o *CircuitTermination) SetCablePeer(v map[string]string)`

SetCablePeer sets CablePeer field to given value.

### HasCablePeer

`func (o *CircuitTermination) HasCablePeer() bool`

HasCablePeer returns a boolean if a field has been set.

### GetCablePeerType

`func (o *CircuitTermination) GetCablePeerType() string`

GetCablePeerType returns the CablePeerType field if non-nil, zero value otherwise.

### GetCablePeerTypeOk

`func (o *CircuitTermination) GetCablePeerTypeOk() (*string, bool)`

GetCablePeerTypeOk returns a tuple with the CablePeerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCablePeerType

`func (o *CircuitTermination) SetCablePeerType(v string)`

SetCablePeerType sets CablePeerType field to given value.

### HasCablePeerType

`func (o *CircuitTermination) HasCablePeerType() bool`

HasCablePeerType returns a boolean if a field has been set.

### GetOccupied

`func (o *CircuitTermination) GetOccupied() bool`

GetOccupied returns the Occupied field if non-nil, zero value otherwise.

### GetOccupiedOk

`func (o *CircuitTermination) GetOccupiedOk() (*bool, bool)`

GetOccupiedOk returns a tuple with the Occupied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccupied

`func (o *CircuitTermination) SetOccupied(v bool)`

SetOccupied sets Occupied field to given value.

### HasOccupied

`func (o *CircuitTermination) HasOccupied() bool`

HasOccupied returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


