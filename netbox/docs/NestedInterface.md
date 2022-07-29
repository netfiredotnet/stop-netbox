# NestedInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Display** | Pointer to **string** |  | [optional] [readonly] 
**Device** | Pointer to [**NestedDevice**](NestedDevice.md) |  | [optional] 
**Name** | **string** |  | 
**Cable** | Pointer to **NullableInt32** |  | [optional] 
**Occupied** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewNestedInterface

`func NewNestedInterface(name string, ) *NestedInterface`

NewNestedInterface instantiates a new NestedInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNestedInterfaceWithDefaults

`func NewNestedInterfaceWithDefaults() *NestedInterface`

NewNestedInterfaceWithDefaults instantiates a new NestedInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NestedInterface) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NestedInterface) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NestedInterface) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *NestedInterface) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *NestedInterface) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NestedInterface) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NestedInterface) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *NestedInterface) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDisplay

`func (o *NestedInterface) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *NestedInterface) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *NestedInterface) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *NestedInterface) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetDevice

`func (o *NestedInterface) GetDevice() NestedDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *NestedInterface) GetDeviceOk() (*NestedDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *NestedInterface) SetDevice(v NestedDevice)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *NestedInterface) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetName

`func (o *NestedInterface) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NestedInterface) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NestedInterface) SetName(v string)`

SetName sets Name field to given value.


### GetCable

`func (o *NestedInterface) GetCable() int32`

GetCable returns the Cable field if non-nil, zero value otherwise.

### GetCableOk

`func (o *NestedInterface) GetCableOk() (*int32, bool)`

GetCableOk returns a tuple with the Cable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCable

`func (o *NestedInterface) SetCable(v int32)`

SetCable sets Cable field to given value.

### HasCable

`func (o *NestedInterface) HasCable() bool`

HasCable returns a boolean if a field has been set.

### SetCableNil

`func (o *NestedInterface) SetCableNil(b bool)`

 SetCableNil sets the value for Cable to be an explicit nil

### UnsetCable
`func (o *NestedInterface) UnsetCable()`

UnsetCable ensures that no value is present for Cable, not even an explicit nil
### GetOccupied

`func (o *NestedInterface) GetOccupied() string`

GetOccupied returns the Occupied field if non-nil, zero value otherwise.

### GetOccupiedOk

`func (o *NestedInterface) GetOccupiedOk() (*string, bool)`

GetOccupiedOk returns a tuple with the Occupied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccupied

`func (o *NestedInterface) SetOccupied(v string)`

SetOccupied sets Occupied field to given value.

### HasOccupied

`func (o *NestedInterface) HasOccupied() bool`

HasOccupied returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


