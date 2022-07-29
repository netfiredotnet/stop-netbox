# WritablePowerFeed

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Display** | Pointer to **string** |  | [optional] [readonly] 
**PowerPanel** | **int32** |  | 
**Rack** | Pointer to **NullableInt32** |  | [optional] 
**Name** | **string** |  | 
**Status** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Supply** | Pointer to **string** |  | [optional] 
**Phase** | Pointer to **string** |  | [optional] 
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

### NewWritablePowerFeed

`func NewWritablePowerFeed(powerPanel int32, name string, ) *WritablePowerFeed`

NewWritablePowerFeed instantiates a new WritablePowerFeed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritablePowerFeedWithDefaults

`func NewWritablePowerFeedWithDefaults() *WritablePowerFeed`

NewWritablePowerFeedWithDefaults instantiates a new WritablePowerFeed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WritablePowerFeed) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WritablePowerFeed) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WritablePowerFeed) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *WritablePowerFeed) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *WritablePowerFeed) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WritablePowerFeed) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WritablePowerFeed) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *WritablePowerFeed) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDisplay

`func (o *WritablePowerFeed) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *WritablePowerFeed) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *WritablePowerFeed) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *WritablePowerFeed) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetPowerPanel

`func (o *WritablePowerFeed) GetPowerPanel() int32`

GetPowerPanel returns the PowerPanel field if non-nil, zero value otherwise.

### GetPowerPanelOk

`func (o *WritablePowerFeed) GetPowerPanelOk() (*int32, bool)`

GetPowerPanelOk returns a tuple with the PowerPanel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerPanel

`func (o *WritablePowerFeed) SetPowerPanel(v int32)`

SetPowerPanel sets PowerPanel field to given value.


### GetRack

`func (o *WritablePowerFeed) GetRack() int32`

GetRack returns the Rack field if non-nil, zero value otherwise.

### GetRackOk

`func (o *WritablePowerFeed) GetRackOk() (*int32, bool)`

GetRackOk returns a tuple with the Rack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRack

`func (o *WritablePowerFeed) SetRack(v int32)`

SetRack sets Rack field to given value.

### HasRack

`func (o *WritablePowerFeed) HasRack() bool`

HasRack returns a boolean if a field has been set.

### SetRackNil

`func (o *WritablePowerFeed) SetRackNil(b bool)`

 SetRackNil sets the value for Rack to be an explicit nil

### UnsetRack
`func (o *WritablePowerFeed) UnsetRack()`

UnsetRack ensures that no value is present for Rack, not even an explicit nil
### GetName

`func (o *WritablePowerFeed) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WritablePowerFeed) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WritablePowerFeed) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *WritablePowerFeed) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WritablePowerFeed) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WritablePowerFeed) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WritablePowerFeed) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *WritablePowerFeed) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WritablePowerFeed) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WritablePowerFeed) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WritablePowerFeed) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSupply

`func (o *WritablePowerFeed) GetSupply() string`

GetSupply returns the Supply field if non-nil, zero value otherwise.

### GetSupplyOk

`func (o *WritablePowerFeed) GetSupplyOk() (*string, bool)`

GetSupplyOk returns a tuple with the Supply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupply

`func (o *WritablePowerFeed) SetSupply(v string)`

SetSupply sets Supply field to given value.

### HasSupply

`func (o *WritablePowerFeed) HasSupply() bool`

HasSupply returns a boolean if a field has been set.

### GetPhase

`func (o *WritablePowerFeed) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *WritablePowerFeed) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *WritablePowerFeed) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *WritablePowerFeed) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetVoltage

`func (o *WritablePowerFeed) GetVoltage() int32`

GetVoltage returns the Voltage field if non-nil, zero value otherwise.

### GetVoltageOk

`func (o *WritablePowerFeed) GetVoltageOk() (*int32, bool)`

GetVoltageOk returns a tuple with the Voltage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoltage

`func (o *WritablePowerFeed) SetVoltage(v int32)`

SetVoltage sets Voltage field to given value.

### HasVoltage

`func (o *WritablePowerFeed) HasVoltage() bool`

HasVoltage returns a boolean if a field has been set.

### GetAmperage

`func (o *WritablePowerFeed) GetAmperage() int32`

GetAmperage returns the Amperage field if non-nil, zero value otherwise.

### GetAmperageOk

`func (o *WritablePowerFeed) GetAmperageOk() (*int32, bool)`

GetAmperageOk returns a tuple with the Amperage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmperage

`func (o *WritablePowerFeed) SetAmperage(v int32)`

SetAmperage sets Amperage field to given value.

### HasAmperage

`func (o *WritablePowerFeed) HasAmperage() bool`

HasAmperage returns a boolean if a field has been set.

### GetMaxUtilization

`func (o *WritablePowerFeed) GetMaxUtilization() int32`

GetMaxUtilization returns the MaxUtilization field if non-nil, zero value otherwise.

### GetMaxUtilizationOk

`func (o *WritablePowerFeed) GetMaxUtilizationOk() (*int32, bool)`

GetMaxUtilizationOk returns a tuple with the MaxUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUtilization

`func (o *WritablePowerFeed) SetMaxUtilization(v int32)`

SetMaxUtilization sets MaxUtilization field to given value.

### HasMaxUtilization

`func (o *WritablePowerFeed) HasMaxUtilization() bool`

HasMaxUtilization returns a boolean if a field has been set.

### GetComments

`func (o *WritablePowerFeed) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *WritablePowerFeed) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *WritablePowerFeed) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *WritablePowerFeed) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetMarkConnected

`func (o *WritablePowerFeed) GetMarkConnected() bool`

GetMarkConnected returns the MarkConnected field if non-nil, zero value otherwise.

### GetMarkConnectedOk

`func (o *WritablePowerFeed) GetMarkConnectedOk() (*bool, bool)`

GetMarkConnectedOk returns a tuple with the MarkConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkConnected

`func (o *WritablePowerFeed) SetMarkConnected(v bool)`

SetMarkConnected sets MarkConnected field to given value.

### HasMarkConnected

`func (o *WritablePowerFeed) HasMarkConnected() bool`

HasMarkConnected returns a boolean if a field has been set.

### GetCable

`func (o *WritablePowerFeed) GetCable() NestedCable`

GetCable returns the Cable field if non-nil, zero value otherwise.

### GetCableOk

`func (o *WritablePowerFeed) GetCableOk() (*NestedCable, bool)`

GetCableOk returns a tuple with the Cable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCable

`func (o *WritablePowerFeed) SetCable(v NestedCable)`

SetCable sets Cable field to given value.

### HasCable

`func (o *WritablePowerFeed) HasCable() bool`

HasCable returns a boolean if a field has been set.

### GetCablePeer

`func (o *WritablePowerFeed) GetCablePeer() map[string]string`

GetCablePeer returns the CablePeer field if non-nil, zero value otherwise.

### GetCablePeerOk

`func (o *WritablePowerFeed) GetCablePeerOk() (*map[string]string, bool)`

GetCablePeerOk returns a tuple with the CablePeer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCablePeer

`func (o *WritablePowerFeed) SetCablePeer(v map[string]string)`

SetCablePeer sets CablePeer field to given value.

### HasCablePeer

`func (o *WritablePowerFeed) HasCablePeer() bool`

HasCablePeer returns a boolean if a field has been set.

### GetCablePeerType

`func (o *WritablePowerFeed) GetCablePeerType() string`

GetCablePeerType returns the CablePeerType field if non-nil, zero value otherwise.

### GetCablePeerTypeOk

`func (o *WritablePowerFeed) GetCablePeerTypeOk() (*string, bool)`

GetCablePeerTypeOk returns a tuple with the CablePeerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCablePeerType

`func (o *WritablePowerFeed) SetCablePeerType(v string)`

SetCablePeerType sets CablePeerType field to given value.

### HasCablePeerType

`func (o *WritablePowerFeed) HasCablePeerType() bool`

HasCablePeerType returns a boolean if a field has been set.

### GetConnectedEndpoint

`func (o *WritablePowerFeed) GetConnectedEndpoint() map[string]string`

GetConnectedEndpoint returns the ConnectedEndpoint field if non-nil, zero value otherwise.

### GetConnectedEndpointOk

`func (o *WritablePowerFeed) GetConnectedEndpointOk() (*map[string]string, bool)`

GetConnectedEndpointOk returns a tuple with the ConnectedEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedEndpoint

`func (o *WritablePowerFeed) SetConnectedEndpoint(v map[string]string)`

SetConnectedEndpoint sets ConnectedEndpoint field to given value.

### HasConnectedEndpoint

`func (o *WritablePowerFeed) HasConnectedEndpoint() bool`

HasConnectedEndpoint returns a boolean if a field has been set.

### GetConnectedEndpointType

`func (o *WritablePowerFeed) GetConnectedEndpointType() string`

GetConnectedEndpointType returns the ConnectedEndpointType field if non-nil, zero value otherwise.

### GetConnectedEndpointTypeOk

`func (o *WritablePowerFeed) GetConnectedEndpointTypeOk() (*string, bool)`

GetConnectedEndpointTypeOk returns a tuple with the ConnectedEndpointType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedEndpointType

`func (o *WritablePowerFeed) SetConnectedEndpointType(v string)`

SetConnectedEndpointType sets ConnectedEndpointType field to given value.

### HasConnectedEndpointType

`func (o *WritablePowerFeed) HasConnectedEndpointType() bool`

HasConnectedEndpointType returns a boolean if a field has been set.

### GetConnectedEndpointReachable

`func (o *WritablePowerFeed) GetConnectedEndpointReachable() bool`

GetConnectedEndpointReachable returns the ConnectedEndpointReachable field if non-nil, zero value otherwise.

### GetConnectedEndpointReachableOk

`func (o *WritablePowerFeed) GetConnectedEndpointReachableOk() (*bool, bool)`

GetConnectedEndpointReachableOk returns a tuple with the ConnectedEndpointReachable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedEndpointReachable

`func (o *WritablePowerFeed) SetConnectedEndpointReachable(v bool)`

SetConnectedEndpointReachable sets ConnectedEndpointReachable field to given value.

### HasConnectedEndpointReachable

`func (o *WritablePowerFeed) HasConnectedEndpointReachable() bool`

HasConnectedEndpointReachable returns a boolean if a field has been set.

### GetTags

`func (o *WritablePowerFeed) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WritablePowerFeed) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WritablePowerFeed) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WritablePowerFeed) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *WritablePowerFeed) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *WritablePowerFeed) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *WritablePowerFeed) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *WritablePowerFeed) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *WritablePowerFeed) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *WritablePowerFeed) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *WritablePowerFeed) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *WritablePowerFeed) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *WritablePowerFeed) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *WritablePowerFeed) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *WritablePowerFeed) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *WritablePowerFeed) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetOccupied

`func (o *WritablePowerFeed) GetOccupied() bool`

GetOccupied returns the Occupied field if non-nil, zero value otherwise.

### GetOccupiedOk

`func (o *WritablePowerFeed) GetOccupiedOk() (*bool, bool)`

GetOccupiedOk returns a tuple with the Occupied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccupied

`func (o *WritablePowerFeed) SetOccupied(v bool)`

SetOccupied sets Occupied field to given value.

### HasOccupied

`func (o *WritablePowerFeed) HasOccupied() bool`

HasOccupied returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


