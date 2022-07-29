# WritableDeviceBay

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Display** | Pointer to **string** |  | [optional] [readonly] 
**Device** | **int32** |  | 
**Name** | **string** |  | 
**Label** | Pointer to **string** | Physical label | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**InstalledDevice** | Pointer to **NullableInt32** |  | [optional] 
**Tags** | Pointer to [**[]NestedTag**](NestedTag.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewWritableDeviceBay

`func NewWritableDeviceBay(device int32, name string, ) *WritableDeviceBay`

NewWritableDeviceBay instantiates a new WritableDeviceBay object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritableDeviceBayWithDefaults

`func NewWritableDeviceBayWithDefaults() *WritableDeviceBay`

NewWritableDeviceBayWithDefaults instantiates a new WritableDeviceBay object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WritableDeviceBay) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WritableDeviceBay) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WritableDeviceBay) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *WritableDeviceBay) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *WritableDeviceBay) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WritableDeviceBay) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WritableDeviceBay) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *WritableDeviceBay) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDisplay

`func (o *WritableDeviceBay) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *WritableDeviceBay) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *WritableDeviceBay) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *WritableDeviceBay) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetDevice

`func (o *WritableDeviceBay) GetDevice() int32`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *WritableDeviceBay) GetDeviceOk() (*int32, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *WritableDeviceBay) SetDevice(v int32)`

SetDevice sets Device field to given value.


### GetName

`func (o *WritableDeviceBay) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WritableDeviceBay) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WritableDeviceBay) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *WritableDeviceBay) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *WritableDeviceBay) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *WritableDeviceBay) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *WritableDeviceBay) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetDescription

`func (o *WritableDeviceBay) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WritableDeviceBay) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WritableDeviceBay) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WritableDeviceBay) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInstalledDevice

`func (o *WritableDeviceBay) GetInstalledDevice() int32`

GetInstalledDevice returns the InstalledDevice field if non-nil, zero value otherwise.

### GetInstalledDeviceOk

`func (o *WritableDeviceBay) GetInstalledDeviceOk() (*int32, bool)`

GetInstalledDeviceOk returns a tuple with the InstalledDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledDevice

`func (o *WritableDeviceBay) SetInstalledDevice(v int32)`

SetInstalledDevice sets InstalledDevice field to given value.

### HasInstalledDevice

`func (o *WritableDeviceBay) HasInstalledDevice() bool`

HasInstalledDevice returns a boolean if a field has been set.

### SetInstalledDeviceNil

`func (o *WritableDeviceBay) SetInstalledDeviceNil(b bool)`

 SetInstalledDeviceNil sets the value for InstalledDevice to be an explicit nil

### UnsetInstalledDevice
`func (o *WritableDeviceBay) UnsetInstalledDevice()`

UnsetInstalledDevice ensures that no value is present for InstalledDevice, not even an explicit nil
### GetTags

`func (o *WritableDeviceBay) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WritableDeviceBay) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WritableDeviceBay) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WritableDeviceBay) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *WritableDeviceBay) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *WritableDeviceBay) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *WritableDeviceBay) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *WritableDeviceBay) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *WritableDeviceBay) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *WritableDeviceBay) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *WritableDeviceBay) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *WritableDeviceBay) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *WritableDeviceBay) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *WritableDeviceBay) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *WritableDeviceBay) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *WritableDeviceBay) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


