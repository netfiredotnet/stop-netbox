# DeviceBayTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Display** | Pointer to **string** |  | [optional] [readonly] 
**DeviceType** | [**NestedDeviceType**](NestedDeviceType.md) |  | 
**Name** | **string** |  | 
**Label** | Pointer to **string** | Physical label | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewDeviceBayTemplate

`func NewDeviceBayTemplate(deviceType NestedDeviceType, name string, ) *DeviceBayTemplate`

NewDeviceBayTemplate instantiates a new DeviceBayTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceBayTemplateWithDefaults

`func NewDeviceBayTemplateWithDefaults() *DeviceBayTemplate`

NewDeviceBayTemplateWithDefaults instantiates a new DeviceBayTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeviceBayTemplate) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceBayTemplate) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceBayTemplate) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *DeviceBayTemplate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *DeviceBayTemplate) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DeviceBayTemplate) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DeviceBayTemplate) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *DeviceBayTemplate) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDisplay

`func (o *DeviceBayTemplate) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *DeviceBayTemplate) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *DeviceBayTemplate) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *DeviceBayTemplate) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetDeviceType

`func (o *DeviceBayTemplate) GetDeviceType() NestedDeviceType`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *DeviceBayTemplate) GetDeviceTypeOk() (*NestedDeviceType, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *DeviceBayTemplate) SetDeviceType(v NestedDeviceType)`

SetDeviceType sets DeviceType field to given value.


### GetName

`func (o *DeviceBayTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceBayTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceBayTemplate) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *DeviceBayTemplate) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *DeviceBayTemplate) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *DeviceBayTemplate) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *DeviceBayTemplate) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetDescription

`func (o *DeviceBayTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DeviceBayTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DeviceBayTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DeviceBayTemplate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreated

`func (o *DeviceBayTemplate) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *DeviceBayTemplate) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *DeviceBayTemplate) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *DeviceBayTemplate) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *DeviceBayTemplate) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *DeviceBayTemplate) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *DeviceBayTemplate) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *DeviceBayTemplate) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


