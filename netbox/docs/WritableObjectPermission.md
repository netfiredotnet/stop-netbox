# WritableObjectPermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Display** | Pointer to **string** |  | [optional] [readonly] 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**ObjectTypes** | **[]string** |  | 
**Groups** | Pointer to **[]int32** |  | [optional] 
**Users** | Pointer to **[]int32** |  | [optional] 
**Actions** | **[]string** | The list of actions granted by this permission | 
**Constraints** | Pointer to **NullableString** | Queryset filter matching the applicable objects of the selected type(s) | [optional] 

## Methods

### NewWritableObjectPermission

`func NewWritableObjectPermission(name string, objectTypes []string, actions []string, ) *WritableObjectPermission`

NewWritableObjectPermission instantiates a new WritableObjectPermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritableObjectPermissionWithDefaults

`func NewWritableObjectPermissionWithDefaults() *WritableObjectPermission`

NewWritableObjectPermissionWithDefaults instantiates a new WritableObjectPermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WritableObjectPermission) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WritableObjectPermission) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WritableObjectPermission) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *WritableObjectPermission) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *WritableObjectPermission) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WritableObjectPermission) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WritableObjectPermission) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *WritableObjectPermission) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDisplay

`func (o *WritableObjectPermission) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *WritableObjectPermission) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *WritableObjectPermission) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *WritableObjectPermission) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetName

`func (o *WritableObjectPermission) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WritableObjectPermission) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WritableObjectPermission) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *WritableObjectPermission) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WritableObjectPermission) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WritableObjectPermission) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WritableObjectPermission) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *WritableObjectPermission) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WritableObjectPermission) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WritableObjectPermission) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *WritableObjectPermission) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetObjectTypes

`func (o *WritableObjectPermission) GetObjectTypes() []string`

GetObjectTypes returns the ObjectTypes field if non-nil, zero value otherwise.

### GetObjectTypesOk

`func (o *WritableObjectPermission) GetObjectTypesOk() (*[]string, bool)`

GetObjectTypesOk returns a tuple with the ObjectTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypes

`func (o *WritableObjectPermission) SetObjectTypes(v []string)`

SetObjectTypes sets ObjectTypes field to given value.


### GetGroups

`func (o *WritableObjectPermission) GetGroups() []int32`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *WritableObjectPermission) GetGroupsOk() (*[]int32, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *WritableObjectPermission) SetGroups(v []int32)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *WritableObjectPermission) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetUsers

`func (o *WritableObjectPermission) GetUsers() []int32`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *WritableObjectPermission) GetUsersOk() (*[]int32, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *WritableObjectPermission) SetUsers(v []int32)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *WritableObjectPermission) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetActions

`func (o *WritableObjectPermission) GetActions() []string`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *WritableObjectPermission) GetActionsOk() (*[]string, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *WritableObjectPermission) SetActions(v []string)`

SetActions sets Actions field to given value.


### GetConstraints

`func (o *WritableObjectPermission) GetConstraints() string`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *WritableObjectPermission) GetConstraintsOk() (*string, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *WritableObjectPermission) SetConstraints(v string)`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *WritableObjectPermission) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.

### SetConstraintsNil

`func (o *WritableObjectPermission) SetConstraintsNil(b bool)`

 SetConstraintsNil sets the value for Constraints to be an explicit nil

### UnsetConstraints
`func (o *WritableObjectPermission) UnsetConstraints()`

UnsetConstraints ensures that no value is present for Constraints, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


