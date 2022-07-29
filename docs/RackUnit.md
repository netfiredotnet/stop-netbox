# RackUnit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] [readonly] 
**Face** | Pointer to [**Face1**](Face1.md) |  | [optional] 
**Device** | Pointer to [**NestedDevice**](NestedDevice.md) |  | [optional] 
**Occupied** | Pointer to **bool** |  | [optional] [readonly] 
**Display** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewRackUnit

`func NewRackUnit() *RackUnit`

NewRackUnit instantiates a new RackUnit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRackUnitWithDefaults

`func NewRackUnitWithDefaults() *RackUnit`

NewRackUnitWithDefaults instantiates a new RackUnit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RackUnit) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RackUnit) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RackUnit) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RackUnit) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RackUnit) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RackUnit) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RackUnit) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RackUnit) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFace

`func (o *RackUnit) GetFace() Face1`

GetFace returns the Face field if non-nil, zero value otherwise.

### GetFaceOk

`func (o *RackUnit) GetFaceOk() (*Face1, bool)`

GetFaceOk returns a tuple with the Face field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFace

`func (o *RackUnit) SetFace(v Face1)`

SetFace sets Face field to given value.

### HasFace

`func (o *RackUnit) HasFace() bool`

HasFace returns a boolean if a field has been set.

### GetDevice

`func (o *RackUnit) GetDevice() NestedDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *RackUnit) GetDeviceOk() (*NestedDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *RackUnit) SetDevice(v NestedDevice)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *RackUnit) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetOccupied

`func (o *RackUnit) GetOccupied() bool`

GetOccupied returns the Occupied field if non-nil, zero value otherwise.

### GetOccupiedOk

`func (o *RackUnit) GetOccupiedOk() (*bool, bool)`

GetOccupiedOk returns a tuple with the Occupied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccupied

`func (o *RackUnit) SetOccupied(v bool)`

SetOccupied sets Occupied field to given value.

### HasOccupied

`func (o *RackUnit) HasOccupied() bool`

HasOccupied returns a boolean if a field has been set.

### GetDisplay

`func (o *RackUnit) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *RackUnit) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *RackUnit) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *RackUnit) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


