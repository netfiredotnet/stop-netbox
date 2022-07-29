# RearPort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Display** | Pointer to **string** |  | [optional] [readonly] 
**Device** | [**NestedDevice**](NestedDevice.md) |  | 
**Name** | **string** |  | 
**Label** | Pointer to **string** | Physical label | [optional] 
**Type** | [**Type1**](Type1.md) |  | 
**Positions** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**MarkConnected** | Pointer to **bool** | Treat as if a cable is connected | [optional] 
**Cable** | Pointer to [**NestedCable**](NestedCable.md) |  | [optional] 
**CablePeer** | Pointer to **map[string]string** |  Return the appropriate serializer for the cable termination model.  | [optional] [readonly] 
**CablePeerType** | Pointer to **string** |  | [optional] [readonly] 
**Tags** | Pointer to [**[]NestedTag**](NestedTag.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] [readonly] 
**Occupied** | Pointer to **bool** |  | [optional] [readonly] 

## Methods

### NewRearPort

`func NewRearPort(device NestedDevice, name string, type_ Type1, ) *RearPort`

NewRearPort instantiates a new RearPort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRearPortWithDefaults

`func NewRearPortWithDefaults() *RearPort`

NewRearPortWithDefaults instantiates a new RearPort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RearPort) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RearPort) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RearPort) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RearPort) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *RearPort) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RearPort) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RearPort) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *RearPort) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDisplay

`func (o *RearPort) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *RearPort) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *RearPort) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *RearPort) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetDevice

`func (o *RearPort) GetDevice() NestedDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *RearPort) GetDeviceOk() (*NestedDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *RearPort) SetDevice(v NestedDevice)`

SetDevice sets Device field to given value.


### GetName

`func (o *RearPort) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RearPort) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RearPort) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *RearPort) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *RearPort) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *RearPort) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *RearPort) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetType

`func (o *RearPort) GetType() Type1`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RearPort) GetTypeOk() (*Type1, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RearPort) SetType(v Type1)`

SetType sets Type field to given value.


### GetPositions

`func (o *RearPort) GetPositions() int32`

GetPositions returns the Positions field if non-nil, zero value otherwise.

### GetPositionsOk

`func (o *RearPort) GetPositionsOk() (*int32, bool)`

GetPositionsOk returns a tuple with the Positions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositions

`func (o *RearPort) SetPositions(v int32)`

SetPositions sets Positions field to given value.

### HasPositions

`func (o *RearPort) HasPositions() bool`

HasPositions returns a boolean if a field has been set.

### GetDescription

`func (o *RearPort) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RearPort) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RearPort) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RearPort) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMarkConnected

`func (o *RearPort) GetMarkConnected() bool`

GetMarkConnected returns the MarkConnected field if non-nil, zero value otherwise.

### GetMarkConnectedOk

`func (o *RearPort) GetMarkConnectedOk() (*bool, bool)`

GetMarkConnectedOk returns a tuple with the MarkConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkConnected

`func (o *RearPort) SetMarkConnected(v bool)`

SetMarkConnected sets MarkConnected field to given value.

### HasMarkConnected

`func (o *RearPort) HasMarkConnected() bool`

HasMarkConnected returns a boolean if a field has been set.

### GetCable

`func (o *RearPort) GetCable() NestedCable`

GetCable returns the Cable field if non-nil, zero value otherwise.

### GetCableOk

`func (o *RearPort) GetCableOk() (*NestedCable, bool)`

GetCableOk returns a tuple with the Cable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCable

`func (o *RearPort) SetCable(v NestedCable)`

SetCable sets Cable field to given value.

### HasCable

`func (o *RearPort) HasCable() bool`

HasCable returns a boolean if a field has been set.

### GetCablePeer

`func (o *RearPort) GetCablePeer() map[string]string`

GetCablePeer returns the CablePeer field if non-nil, zero value otherwise.

### GetCablePeerOk

`func (o *RearPort) GetCablePeerOk() (*map[string]string, bool)`

GetCablePeerOk returns a tuple with the CablePeer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCablePeer

`func (o *RearPort) SetCablePeer(v map[string]string)`

SetCablePeer sets CablePeer field to given value.

### HasCablePeer

`func (o *RearPort) HasCablePeer() bool`

HasCablePeer returns a boolean if a field has been set.

### GetCablePeerType

`func (o *RearPort) GetCablePeerType() string`

GetCablePeerType returns the CablePeerType field if non-nil, zero value otherwise.

### GetCablePeerTypeOk

`func (o *RearPort) GetCablePeerTypeOk() (*string, bool)`

GetCablePeerTypeOk returns a tuple with the CablePeerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCablePeerType

`func (o *RearPort) SetCablePeerType(v string)`

SetCablePeerType sets CablePeerType field to given value.

### HasCablePeerType

`func (o *RearPort) HasCablePeerType() bool`

HasCablePeerType returns a boolean if a field has been set.

### GetTags

`func (o *RearPort) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RearPort) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RearPort) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RearPort) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *RearPort) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *RearPort) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *RearPort) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *RearPort) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *RearPort) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *RearPort) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *RearPort) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *RearPort) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *RearPort) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *RearPort) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *RearPort) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *RearPort) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetOccupied

`func (o *RearPort) GetOccupied() bool`

GetOccupied returns the Occupied field if non-nil, zero value otherwise.

### GetOccupiedOk

`func (o *RearPort) GetOccupiedOk() (*bool, bool)`

GetOccupiedOk returns a tuple with the Occupied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccupied

`func (o *RearPort) SetOccupied(v bool)`

SetOccupied sets Occupied field to given value.

### HasOccupied

`func (o *RearPort) HasOccupied() bool`

HasOccupied returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


