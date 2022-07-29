# WritablePowerOutlet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Display** | Pointer to **string** |  | [optional] [readonly] 
**Device** | **int32** |  | 
**Name** | **string** |  | 
**Label** | Pointer to **string** | Physical label | [optional] 
**Type** | Pointer to **string** | Physical port type | [optional] 
**PowerPort** | Pointer to **NullableInt32** |  | [optional] 
**FeedLeg** | Pointer to **string** | Phase (for three-phase feeds) | [optional] 
**Description** | Pointer to **string** |  | [optional] 
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

### NewWritablePowerOutlet

`func NewWritablePowerOutlet(device int32, name string, ) *WritablePowerOutlet`

NewWritablePowerOutlet instantiates a new WritablePowerOutlet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritablePowerOutletWithDefaults

`func NewWritablePowerOutletWithDefaults() *WritablePowerOutlet`

NewWritablePowerOutletWithDefaults instantiates a new WritablePowerOutlet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WritablePowerOutlet) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WritablePowerOutlet) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WritablePowerOutlet) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *WritablePowerOutlet) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *WritablePowerOutlet) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WritablePowerOutlet) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WritablePowerOutlet) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *WritablePowerOutlet) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDisplay

`func (o *WritablePowerOutlet) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *WritablePowerOutlet) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *WritablePowerOutlet) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *WritablePowerOutlet) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetDevice

`func (o *WritablePowerOutlet) GetDevice() int32`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *WritablePowerOutlet) GetDeviceOk() (*int32, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *WritablePowerOutlet) SetDevice(v int32)`

SetDevice sets Device field to given value.


### GetName

`func (o *WritablePowerOutlet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WritablePowerOutlet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WritablePowerOutlet) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *WritablePowerOutlet) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *WritablePowerOutlet) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *WritablePowerOutlet) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *WritablePowerOutlet) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetType

`func (o *WritablePowerOutlet) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WritablePowerOutlet) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WritablePowerOutlet) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WritablePowerOutlet) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPowerPort

`func (o *WritablePowerOutlet) GetPowerPort() int32`

GetPowerPort returns the PowerPort field if non-nil, zero value otherwise.

### GetPowerPortOk

`func (o *WritablePowerOutlet) GetPowerPortOk() (*int32, bool)`

GetPowerPortOk returns a tuple with the PowerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerPort

`func (o *WritablePowerOutlet) SetPowerPort(v int32)`

SetPowerPort sets PowerPort field to given value.

### HasPowerPort

`func (o *WritablePowerOutlet) HasPowerPort() bool`

HasPowerPort returns a boolean if a field has been set.

### SetPowerPortNil

`func (o *WritablePowerOutlet) SetPowerPortNil(b bool)`

 SetPowerPortNil sets the value for PowerPort to be an explicit nil

### UnsetPowerPort
`func (o *WritablePowerOutlet) UnsetPowerPort()`

UnsetPowerPort ensures that no value is present for PowerPort, not even an explicit nil
### GetFeedLeg

`func (o *WritablePowerOutlet) GetFeedLeg() string`

GetFeedLeg returns the FeedLeg field if non-nil, zero value otherwise.

### GetFeedLegOk

`func (o *WritablePowerOutlet) GetFeedLegOk() (*string, bool)`

GetFeedLegOk returns a tuple with the FeedLeg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedLeg

`func (o *WritablePowerOutlet) SetFeedLeg(v string)`

SetFeedLeg sets FeedLeg field to given value.

### HasFeedLeg

`func (o *WritablePowerOutlet) HasFeedLeg() bool`

HasFeedLeg returns a boolean if a field has been set.

### GetDescription

`func (o *WritablePowerOutlet) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WritablePowerOutlet) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WritablePowerOutlet) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WritablePowerOutlet) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMarkConnected

`func (o *WritablePowerOutlet) GetMarkConnected() bool`

GetMarkConnected returns the MarkConnected field if non-nil, zero value otherwise.

### GetMarkConnectedOk

`func (o *WritablePowerOutlet) GetMarkConnectedOk() (*bool, bool)`

GetMarkConnectedOk returns a tuple with the MarkConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkConnected

`func (o *WritablePowerOutlet) SetMarkConnected(v bool)`

SetMarkConnected sets MarkConnected field to given value.

### HasMarkConnected

`func (o *WritablePowerOutlet) HasMarkConnected() bool`

HasMarkConnected returns a boolean if a field has been set.

### GetCable

`func (o *WritablePowerOutlet) GetCable() NestedCable`

GetCable returns the Cable field if non-nil, zero value otherwise.

### GetCableOk

`func (o *WritablePowerOutlet) GetCableOk() (*NestedCable, bool)`

GetCableOk returns a tuple with the Cable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCable

`func (o *WritablePowerOutlet) SetCable(v NestedCable)`

SetCable sets Cable field to given value.

### HasCable

`func (o *WritablePowerOutlet) HasCable() bool`

HasCable returns a boolean if a field has been set.

### GetCablePeer

`func (o *WritablePowerOutlet) GetCablePeer() map[string]string`

GetCablePeer returns the CablePeer field if non-nil, zero value otherwise.

### GetCablePeerOk

`func (o *WritablePowerOutlet) GetCablePeerOk() (*map[string]string, bool)`

GetCablePeerOk returns a tuple with the CablePeer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCablePeer

`func (o *WritablePowerOutlet) SetCablePeer(v map[string]string)`

SetCablePeer sets CablePeer field to given value.

### HasCablePeer

`func (o *WritablePowerOutlet) HasCablePeer() bool`

HasCablePeer returns a boolean if a field has been set.

### GetCablePeerType

`func (o *WritablePowerOutlet) GetCablePeerType() string`

GetCablePeerType returns the CablePeerType field if non-nil, zero value otherwise.

### GetCablePeerTypeOk

`func (o *WritablePowerOutlet) GetCablePeerTypeOk() (*string, bool)`

GetCablePeerTypeOk returns a tuple with the CablePeerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCablePeerType

`func (o *WritablePowerOutlet) SetCablePeerType(v string)`

SetCablePeerType sets CablePeerType field to given value.

### HasCablePeerType

`func (o *WritablePowerOutlet) HasCablePeerType() bool`

HasCablePeerType returns a boolean if a field has been set.

### GetConnectedEndpoint

`func (o *WritablePowerOutlet) GetConnectedEndpoint() map[string]string`

GetConnectedEndpoint returns the ConnectedEndpoint field if non-nil, zero value otherwise.

### GetConnectedEndpointOk

`func (o *WritablePowerOutlet) GetConnectedEndpointOk() (*map[string]string, bool)`

GetConnectedEndpointOk returns a tuple with the ConnectedEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedEndpoint

`func (o *WritablePowerOutlet) SetConnectedEndpoint(v map[string]string)`

SetConnectedEndpoint sets ConnectedEndpoint field to given value.

### HasConnectedEndpoint

`func (o *WritablePowerOutlet) HasConnectedEndpoint() bool`

HasConnectedEndpoint returns a boolean if a field has been set.

### GetConnectedEndpointType

`func (o *WritablePowerOutlet) GetConnectedEndpointType() string`

GetConnectedEndpointType returns the ConnectedEndpointType field if non-nil, zero value otherwise.

### GetConnectedEndpointTypeOk

`func (o *WritablePowerOutlet) GetConnectedEndpointTypeOk() (*string, bool)`

GetConnectedEndpointTypeOk returns a tuple with the ConnectedEndpointType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedEndpointType

`func (o *WritablePowerOutlet) SetConnectedEndpointType(v string)`

SetConnectedEndpointType sets ConnectedEndpointType field to given value.

### HasConnectedEndpointType

`func (o *WritablePowerOutlet) HasConnectedEndpointType() bool`

HasConnectedEndpointType returns a boolean if a field has been set.

### GetConnectedEndpointReachable

`func (o *WritablePowerOutlet) GetConnectedEndpointReachable() bool`

GetConnectedEndpointReachable returns the ConnectedEndpointReachable field if non-nil, zero value otherwise.

### GetConnectedEndpointReachableOk

`func (o *WritablePowerOutlet) GetConnectedEndpointReachableOk() (*bool, bool)`

GetConnectedEndpointReachableOk returns a tuple with the ConnectedEndpointReachable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedEndpointReachable

`func (o *WritablePowerOutlet) SetConnectedEndpointReachable(v bool)`

SetConnectedEndpointReachable sets ConnectedEndpointReachable field to given value.

### HasConnectedEndpointReachable

`func (o *WritablePowerOutlet) HasConnectedEndpointReachable() bool`

HasConnectedEndpointReachable returns a boolean if a field has been set.

### GetTags

`func (o *WritablePowerOutlet) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WritablePowerOutlet) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WritablePowerOutlet) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WritablePowerOutlet) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *WritablePowerOutlet) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *WritablePowerOutlet) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *WritablePowerOutlet) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *WritablePowerOutlet) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *WritablePowerOutlet) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *WritablePowerOutlet) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *WritablePowerOutlet) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *WritablePowerOutlet) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *WritablePowerOutlet) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *WritablePowerOutlet) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *WritablePowerOutlet) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *WritablePowerOutlet) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetOccupied

`func (o *WritablePowerOutlet) GetOccupied() bool`

GetOccupied returns the Occupied field if non-nil, zero value otherwise.

### GetOccupiedOk

`func (o *WritablePowerOutlet) GetOccupiedOk() (*bool, bool)`

GetOccupiedOk returns a tuple with the Occupied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccupied

`func (o *WritablePowerOutlet) SetOccupied(v bool)`

SetOccupied sets Occupied field to given value.

### HasOccupied

`func (o *WritablePowerOutlet) HasOccupied() bool`

HasOccupied returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


