# WritablePowerPort

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
**MaximumDraw** | Pointer to **NullableInt32** | Maximum power draw (watts) | [optional] 
**AllocatedDraw** | Pointer to **NullableInt32** | Allocated power draw (watts) | [optional] 
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

### NewWritablePowerPort

`func NewWritablePowerPort(device int32, name string, ) *WritablePowerPort`

NewWritablePowerPort instantiates a new WritablePowerPort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritablePowerPortWithDefaults

`func NewWritablePowerPortWithDefaults() *WritablePowerPort`

NewWritablePowerPortWithDefaults instantiates a new WritablePowerPort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WritablePowerPort) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WritablePowerPort) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WritablePowerPort) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *WritablePowerPort) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *WritablePowerPort) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WritablePowerPort) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WritablePowerPort) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *WritablePowerPort) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDisplay

`func (o *WritablePowerPort) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *WritablePowerPort) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *WritablePowerPort) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *WritablePowerPort) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetDevice

`func (o *WritablePowerPort) GetDevice() int32`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *WritablePowerPort) GetDeviceOk() (*int32, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *WritablePowerPort) SetDevice(v int32)`

SetDevice sets Device field to given value.


### GetName

`func (o *WritablePowerPort) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WritablePowerPort) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WritablePowerPort) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *WritablePowerPort) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *WritablePowerPort) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *WritablePowerPort) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *WritablePowerPort) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetType

`func (o *WritablePowerPort) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WritablePowerPort) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WritablePowerPort) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WritablePowerPort) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMaximumDraw

`func (o *WritablePowerPort) GetMaximumDraw() int32`

GetMaximumDraw returns the MaximumDraw field if non-nil, zero value otherwise.

### GetMaximumDrawOk

`func (o *WritablePowerPort) GetMaximumDrawOk() (*int32, bool)`

GetMaximumDrawOk returns a tuple with the MaximumDraw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumDraw

`func (o *WritablePowerPort) SetMaximumDraw(v int32)`

SetMaximumDraw sets MaximumDraw field to given value.

### HasMaximumDraw

`func (o *WritablePowerPort) HasMaximumDraw() bool`

HasMaximumDraw returns a boolean if a field has been set.

### SetMaximumDrawNil

`func (o *WritablePowerPort) SetMaximumDrawNil(b bool)`

 SetMaximumDrawNil sets the value for MaximumDraw to be an explicit nil

### UnsetMaximumDraw
`func (o *WritablePowerPort) UnsetMaximumDraw()`

UnsetMaximumDraw ensures that no value is present for MaximumDraw, not even an explicit nil
### GetAllocatedDraw

`func (o *WritablePowerPort) GetAllocatedDraw() int32`

GetAllocatedDraw returns the AllocatedDraw field if non-nil, zero value otherwise.

### GetAllocatedDrawOk

`func (o *WritablePowerPort) GetAllocatedDrawOk() (*int32, bool)`

GetAllocatedDrawOk returns a tuple with the AllocatedDraw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedDraw

`func (o *WritablePowerPort) SetAllocatedDraw(v int32)`

SetAllocatedDraw sets AllocatedDraw field to given value.

### HasAllocatedDraw

`func (o *WritablePowerPort) HasAllocatedDraw() bool`

HasAllocatedDraw returns a boolean if a field has been set.

### SetAllocatedDrawNil

`func (o *WritablePowerPort) SetAllocatedDrawNil(b bool)`

 SetAllocatedDrawNil sets the value for AllocatedDraw to be an explicit nil

### UnsetAllocatedDraw
`func (o *WritablePowerPort) UnsetAllocatedDraw()`

UnsetAllocatedDraw ensures that no value is present for AllocatedDraw, not even an explicit nil
### GetDescription

`func (o *WritablePowerPort) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WritablePowerPort) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WritablePowerPort) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WritablePowerPort) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMarkConnected

`func (o *WritablePowerPort) GetMarkConnected() bool`

GetMarkConnected returns the MarkConnected field if non-nil, zero value otherwise.

### GetMarkConnectedOk

`func (o *WritablePowerPort) GetMarkConnectedOk() (*bool, bool)`

GetMarkConnectedOk returns a tuple with the MarkConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkConnected

`func (o *WritablePowerPort) SetMarkConnected(v bool)`

SetMarkConnected sets MarkConnected field to given value.

### HasMarkConnected

`func (o *WritablePowerPort) HasMarkConnected() bool`

HasMarkConnected returns a boolean if a field has been set.

### GetCable

`func (o *WritablePowerPort) GetCable() NestedCable`

GetCable returns the Cable field if non-nil, zero value otherwise.

### GetCableOk

`func (o *WritablePowerPort) GetCableOk() (*NestedCable, bool)`

GetCableOk returns a tuple with the Cable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCable

`func (o *WritablePowerPort) SetCable(v NestedCable)`

SetCable sets Cable field to given value.

### HasCable

`func (o *WritablePowerPort) HasCable() bool`

HasCable returns a boolean if a field has been set.

### GetCablePeer

`func (o *WritablePowerPort) GetCablePeer() map[string]string`

GetCablePeer returns the CablePeer field if non-nil, zero value otherwise.

### GetCablePeerOk

`func (o *WritablePowerPort) GetCablePeerOk() (*map[string]string, bool)`

GetCablePeerOk returns a tuple with the CablePeer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCablePeer

`func (o *WritablePowerPort) SetCablePeer(v map[string]string)`

SetCablePeer sets CablePeer field to given value.

### HasCablePeer

`func (o *WritablePowerPort) HasCablePeer() bool`

HasCablePeer returns a boolean if a field has been set.

### GetCablePeerType

`func (o *WritablePowerPort) GetCablePeerType() string`

GetCablePeerType returns the CablePeerType field if non-nil, zero value otherwise.

### GetCablePeerTypeOk

`func (o *WritablePowerPort) GetCablePeerTypeOk() (*string, bool)`

GetCablePeerTypeOk returns a tuple with the CablePeerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCablePeerType

`func (o *WritablePowerPort) SetCablePeerType(v string)`

SetCablePeerType sets CablePeerType field to given value.

### HasCablePeerType

`func (o *WritablePowerPort) HasCablePeerType() bool`

HasCablePeerType returns a boolean if a field has been set.

### GetConnectedEndpoint

`func (o *WritablePowerPort) GetConnectedEndpoint() map[string]string`

GetConnectedEndpoint returns the ConnectedEndpoint field if non-nil, zero value otherwise.

### GetConnectedEndpointOk

`func (o *WritablePowerPort) GetConnectedEndpointOk() (*map[string]string, bool)`

GetConnectedEndpointOk returns a tuple with the ConnectedEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedEndpoint

`func (o *WritablePowerPort) SetConnectedEndpoint(v map[string]string)`

SetConnectedEndpoint sets ConnectedEndpoint field to given value.

### HasConnectedEndpoint

`func (o *WritablePowerPort) HasConnectedEndpoint() bool`

HasConnectedEndpoint returns a boolean if a field has been set.

### GetConnectedEndpointType

`func (o *WritablePowerPort) GetConnectedEndpointType() string`

GetConnectedEndpointType returns the ConnectedEndpointType field if non-nil, zero value otherwise.

### GetConnectedEndpointTypeOk

`func (o *WritablePowerPort) GetConnectedEndpointTypeOk() (*string, bool)`

GetConnectedEndpointTypeOk returns a tuple with the ConnectedEndpointType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedEndpointType

`func (o *WritablePowerPort) SetConnectedEndpointType(v string)`

SetConnectedEndpointType sets ConnectedEndpointType field to given value.

### HasConnectedEndpointType

`func (o *WritablePowerPort) HasConnectedEndpointType() bool`

HasConnectedEndpointType returns a boolean if a field has been set.

### GetConnectedEndpointReachable

`func (o *WritablePowerPort) GetConnectedEndpointReachable() bool`

GetConnectedEndpointReachable returns the ConnectedEndpointReachable field if non-nil, zero value otherwise.

### GetConnectedEndpointReachableOk

`func (o *WritablePowerPort) GetConnectedEndpointReachableOk() (*bool, bool)`

GetConnectedEndpointReachableOk returns a tuple with the ConnectedEndpointReachable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedEndpointReachable

`func (o *WritablePowerPort) SetConnectedEndpointReachable(v bool)`

SetConnectedEndpointReachable sets ConnectedEndpointReachable field to given value.

### HasConnectedEndpointReachable

`func (o *WritablePowerPort) HasConnectedEndpointReachable() bool`

HasConnectedEndpointReachable returns a boolean if a field has been set.

### GetTags

`func (o *WritablePowerPort) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WritablePowerPort) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WritablePowerPort) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WritablePowerPort) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *WritablePowerPort) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *WritablePowerPort) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *WritablePowerPort) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *WritablePowerPort) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *WritablePowerPort) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *WritablePowerPort) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *WritablePowerPort) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *WritablePowerPort) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *WritablePowerPort) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *WritablePowerPort) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *WritablePowerPort) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *WritablePowerPort) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetOccupied

`func (o *WritablePowerPort) GetOccupied() bool`

GetOccupied returns the Occupied field if non-nil, zero value otherwise.

### GetOccupiedOk

`func (o *WritablePowerPort) GetOccupiedOk() (*bool, bool)`

GetOccupiedOk returns a tuple with the Occupied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccupied

`func (o *WritablePowerPort) SetOccupied(v bool)`

SetOccupied sets Occupied field to given value.

### HasOccupied

`func (o *WritablePowerPort) HasOccupied() bool`

HasOccupied returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


