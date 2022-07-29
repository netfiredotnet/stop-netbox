# PowerFeed

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Display** | Pointer to **string** |  | [optional] [readonly] 
**PowerPanel** | [**NestedPowerPanel**](NestedPowerPanel.md) |  | 
**Rack** | Pointer to [**NullableNestedRack**](NestedRack.md) |  | [optional] 
**Name** | **string** |  | 
**Status** | Pointer to [**Status3**](Status3.md) |  | [optional] 
**Type** | Pointer to [**Type4**](Type4.md) |  | [optional] 
**Supply** | Pointer to [**Supply**](Supply.md) |  | [optional] 
**Phase** | Pointer to [**Phase**](Phase.md) |  | [optional] 
**Voltage** | Pointer to **int32** |  | [optional] 
**Amperage** | Pointer to **int32** |  | [optional] 
**MaxUtilization** | Pointer to **int32** | Maximum permissible draw (percentage) | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**MarkConnected** | Pointer to **bool** | Treat as if a cable is connected | [optional] 
**Cable** | Pointer to [**NestedCable**](NestedCable.md) |  | [optional] 
**CablePeer** | Pointer to **map[string]string** |  Return the appropriate serializer for the cable termination model.  | [optional] [readonly] 
**CablePeerType** | Pointer to **string** |  | [optional] [readonly] 
**ConnectedEndpoint** | Pointer to **map[string]string** |  Return the appropriate serializer for the type of connected object.  | [optional] [readonly] 
**ConnectedEndpointType** | Pointer to **string** |  | [optional] [readonly] 
**ConnectedEndpointReachable** | Pointer to **bool** |  | [optional] [readonly] 
**Tags** | Pointer to [**[]NestedTag**](NestedTag.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] [readonly] 
**Occupied** | Pointer to **bool** |  | [optional] [readonly] 

## Methods

### NewPowerFeed

`func NewPowerFeed(powerPanel NestedPowerPanel, name string, ) *PowerFeed`

NewPowerFeed instantiates a new PowerFeed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPowerFeedWithDefaults

`func NewPowerFeedWithDefaults() *PowerFeed`

NewPowerFeedWithDefaults instantiates a new PowerFeed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PowerFeed) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PowerFeed) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PowerFeed) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PowerFeed) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *PowerFeed) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PowerFeed) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PowerFeed) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PowerFeed) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDisplay

`func (o *PowerFeed) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *PowerFeed) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *PowerFeed) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *PowerFeed) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetPowerPanel

`func (o *PowerFeed) GetPowerPanel() NestedPowerPanel`

GetPowerPanel returns the PowerPanel field if non-nil, zero value otherwise.

### GetPowerPanelOk

`func (o *PowerFeed) GetPowerPanelOk() (*NestedPowerPanel, bool)`

GetPowerPanelOk returns a tuple with the PowerPanel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerPanel

`func (o *PowerFeed) SetPowerPanel(v NestedPowerPanel)`

SetPowerPanel sets PowerPanel field to given value.


### GetRack

`func (o *PowerFeed) GetRack() NestedRack`

GetRack returns the Rack field if non-nil, zero value otherwise.

### GetRackOk

`func (o *PowerFeed) GetRackOk() (*NestedRack, bool)`

GetRackOk returns a tuple with the Rack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRack

`func (o *PowerFeed) SetRack(v NestedRack)`

SetRack sets Rack field to given value.

### HasRack

`func (o *PowerFeed) HasRack() bool`

HasRack returns a boolean if a field has been set.

### SetRackNil

`func (o *PowerFeed) SetRackNil(b bool)`

 SetRackNil sets the value for Rack to be an explicit nil

### UnsetRack
`func (o *PowerFeed) UnsetRack()`

UnsetRack ensures that no value is present for Rack, not even an explicit nil
### GetName

`func (o *PowerFeed) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PowerFeed) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PowerFeed) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *PowerFeed) GetStatus() Status3`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PowerFeed) GetStatusOk() (*Status3, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PowerFeed) SetStatus(v Status3)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PowerFeed) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *PowerFeed) GetType() Type4`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PowerFeed) GetTypeOk() (*Type4, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PowerFeed) SetType(v Type4)`

SetType sets Type field to given value.

### HasType

`func (o *PowerFeed) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSupply

`func (o *PowerFeed) GetSupply() Supply`

GetSupply returns the Supply field if non-nil, zero value otherwise.

### GetSupplyOk

`func (o *PowerFeed) GetSupplyOk() (*Supply, bool)`

GetSupplyOk returns a tuple with the Supply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupply

`func (o *PowerFeed) SetSupply(v Supply)`

SetSupply sets Supply field to given value.

### HasSupply

`func (o *PowerFeed) HasSupply() bool`

HasSupply returns a boolean if a field has been set.

### GetPhase

`func (o *PowerFeed) GetPhase() Phase`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *PowerFeed) GetPhaseOk() (*Phase, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *PowerFeed) SetPhase(v Phase)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *PowerFeed) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetVoltage

`func (o *PowerFeed) GetVoltage() int32`

GetVoltage returns the Voltage field if non-nil, zero value otherwise.

### GetVoltageOk

`func (o *PowerFeed) GetVoltageOk() (*int32, bool)`

GetVoltageOk returns a tuple with the Voltage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoltage

`func (o *PowerFeed) SetVoltage(v int32)`

SetVoltage sets Voltage field to given value.

### HasVoltage

`func (o *PowerFeed) HasVoltage() bool`

HasVoltage returns a boolean if a field has been set.

### GetAmperage

`func (o *PowerFeed) GetAmperage() int32`

GetAmperage returns the Amperage field if non-nil, zero value otherwise.

### GetAmperageOk

`func (o *PowerFeed) GetAmperageOk() (*int32, bool)`

GetAmperageOk returns a tuple with the Amperage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmperage

`func (o *PowerFeed) SetAmperage(v int32)`

SetAmperage sets Amperage field to given value.

### HasAmperage

`func (o *PowerFeed) HasAmperage() bool`

HasAmperage returns a boolean if a field has been set.

### GetMaxUtilization

`func (o *PowerFeed) GetMaxUtilization() int32`

GetMaxUtilization returns the MaxUtilization field if non-nil, zero value otherwise.

### GetMaxUtilizationOk

`func (o *PowerFeed) GetMaxUtilizationOk() (*int32, bool)`

GetMaxUtilizationOk returns a tuple with the MaxUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUtilization

`func (o *PowerFeed) SetMaxUtilization(v int32)`

SetMaxUtilization sets MaxUtilization field to given value.

### HasMaxUtilization

`func (o *PowerFeed) HasMaxUtilization() bool`

HasMaxUtilization returns a boolean if a field has been set.

### GetComments

`func (o *PowerFeed) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PowerFeed) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PowerFeed) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PowerFeed) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetMarkConnected

`func (o *PowerFeed) GetMarkConnected() bool`

GetMarkConnected returns the MarkConnected field if non-nil, zero value otherwise.

### GetMarkConnectedOk

`func (o *PowerFeed) GetMarkConnectedOk() (*bool, bool)`

GetMarkConnectedOk returns a tuple with the MarkConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkConnected

`func (o *PowerFeed) SetMarkConnected(v bool)`

SetMarkConnected sets MarkConnected field to given value.

### HasMarkConnected

`func (o *PowerFeed) HasMarkConnected() bool`

HasMarkConnected returns a boolean if a field has been set.

### GetCable

`func (o *PowerFeed) GetCable() NestedCable`

GetCable returns the Cable field if non-nil, zero value otherwise.

### GetCableOk

`func (o *PowerFeed) GetCableOk() (*NestedCable, bool)`

GetCableOk returns a tuple with the Cable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCable

`func (o *PowerFeed) SetCable(v NestedCable)`

SetCable sets Cable field to given value.

### HasCable

`func (o *PowerFeed) HasCable() bool`

HasCable returns a boolean if a field has been set.

### GetCablePeer

`func (o *PowerFeed) GetCablePeer() map[string]string`

GetCablePeer returns the CablePeer field if non-nil, zero value otherwise.

### GetCablePeerOk

`func (o *PowerFeed) GetCablePeerOk() (*map[string]string, bool)`

GetCablePeerOk returns a tuple with the CablePeer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCablePeer

`func (o *PowerFeed) SetCablePeer(v map[string]string)`

SetCablePeer sets CablePeer field to given value.

### HasCablePeer

`func (o *PowerFeed) HasCablePeer() bool`

HasCablePeer returns a boolean if a field has been set.

### GetCablePeerType

`func (o *PowerFeed) GetCablePeerType() string`

GetCablePeerType returns the CablePeerType field if non-nil, zero value otherwise.

### GetCablePeerTypeOk

`func (o *PowerFeed) GetCablePeerTypeOk() (*string, bool)`

GetCablePeerTypeOk returns a tuple with the CablePeerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCablePeerType

`func (o *PowerFeed) SetCablePeerType(v string)`

SetCablePeerType sets CablePeerType field to given value.

### HasCablePeerType

`func (o *PowerFeed) HasCablePeerType() bool`

HasCablePeerType returns a boolean if a field has been set.

### GetConnectedEndpoint

`func (o *PowerFeed) GetConnectedEndpoint() map[string]string`

GetConnectedEndpoint returns the ConnectedEndpoint field if non-nil, zero value otherwise.

### GetConnectedEndpointOk

`func (o *PowerFeed) GetConnectedEndpointOk() (*map[string]string, bool)`

GetConnectedEndpointOk returns a tuple with the ConnectedEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedEndpoint

`func (o *PowerFeed) SetConnectedEndpoint(v map[string]string)`

SetConnectedEndpoint sets ConnectedEndpoint field to given value.

### HasConnectedEndpoint

`func (o *PowerFeed) HasConnectedEndpoint() bool`

HasConnectedEndpoint returns a boolean if a field has been set.

### GetConnectedEndpointType

`func (o *PowerFeed) GetConnectedEndpointType() string`

GetConnectedEndpointType returns the ConnectedEndpointType field if non-nil, zero value otherwise.

### GetConnectedEndpointTypeOk

`func (o *PowerFeed) GetConnectedEndpointTypeOk() (*string, bool)`

GetConnectedEndpointTypeOk returns a tuple with the ConnectedEndpointType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedEndpointType

`func (o *PowerFeed) SetConnectedEndpointType(v string)`

SetConnectedEndpointType sets ConnectedEndpointType field to given value.

### HasConnectedEndpointType

`func (o *PowerFeed) HasConnectedEndpointType() bool`

HasConnectedEndpointType returns a boolean if a field has been set.

### GetConnectedEndpointReachable

`func (o *PowerFeed) GetConnectedEndpointReachable() bool`

GetConnectedEndpointReachable returns the ConnectedEndpointReachable field if non-nil, zero value otherwise.

### GetConnectedEndpointReachableOk

`func (o *PowerFeed) GetConnectedEndpointReachableOk() (*bool, bool)`

GetConnectedEndpointReachableOk returns a tuple with the ConnectedEndpointReachable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedEndpointReachable

`func (o *PowerFeed) SetConnectedEndpointReachable(v bool)`

SetConnectedEndpointReachable sets ConnectedEndpointReachable field to given value.

### HasConnectedEndpointReachable

`func (o *PowerFeed) HasConnectedEndpointReachable() bool`

HasConnectedEndpointReachable returns a boolean if a field has been set.

### GetTags

`func (o *PowerFeed) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PowerFeed) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PowerFeed) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PowerFeed) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PowerFeed) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PowerFeed) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PowerFeed) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PowerFeed) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *PowerFeed) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PowerFeed) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PowerFeed) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *PowerFeed) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *PowerFeed) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *PowerFeed) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *PowerFeed) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *PowerFeed) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetOccupied

`func (o *PowerFeed) GetOccupied() bool`

GetOccupied returns the Occupied field if non-nil, zero value otherwise.

### GetOccupiedOk

`func (o *PowerFeed) GetOccupiedOk() (*bool, bool)`

GetOccupiedOk returns a tuple with the Occupied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccupied

`func (o *PowerFeed) SetOccupied(v bool)`

SetOccupied sets Occupied field to given value.

### HasOccupied

`func (o *PowerFeed) HasOccupied() bool`

HasOccupied returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


