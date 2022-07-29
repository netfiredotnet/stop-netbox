# WritablePlatform

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Display** | Pointer to **string** |  | [optional] [readonly] 
**Name** | **string** |  | 
**Slug** | **string** |  | 
**Manufacturer** | Pointer to **NullableInt32** | Optionally limit this platform to devices of a certain manufacturer | [optional] 
**NapalmDriver** | Pointer to **string** | The name of the NAPALM driver to use when interacting with devices | [optional] 
**NapalmArgs** | Pointer to **NullableString** | Additional arguments to pass when initiating the NAPALM driver (JSON format) | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] [readonly] 
**DeviceCount** | Pointer to **int32** |  | [optional] [readonly] 
**VirtualmachineCount** | Pointer to **int32** |  | [optional] [readonly] 

## Methods

### NewWritablePlatform

`func NewWritablePlatform(name string, slug string, ) *WritablePlatform`

NewWritablePlatform instantiates a new WritablePlatform object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritablePlatformWithDefaults

`func NewWritablePlatformWithDefaults() *WritablePlatform`

NewWritablePlatformWithDefaults instantiates a new WritablePlatform object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WritablePlatform) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WritablePlatform) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WritablePlatform) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *WritablePlatform) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *WritablePlatform) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WritablePlatform) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WritablePlatform) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *WritablePlatform) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDisplay

`func (o *WritablePlatform) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *WritablePlatform) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *WritablePlatform) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *WritablePlatform) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetName

`func (o *WritablePlatform) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WritablePlatform) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WritablePlatform) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *WritablePlatform) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *WritablePlatform) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *WritablePlatform) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetManufacturer

`func (o *WritablePlatform) GetManufacturer() int32`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *WritablePlatform) GetManufacturerOk() (*int32, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *WritablePlatform) SetManufacturer(v int32)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *WritablePlatform) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### SetManufacturerNil

`func (o *WritablePlatform) SetManufacturerNil(b bool)`

 SetManufacturerNil sets the value for Manufacturer to be an explicit nil

### UnsetManufacturer
`func (o *WritablePlatform) UnsetManufacturer()`

UnsetManufacturer ensures that no value is present for Manufacturer, not even an explicit nil
### GetNapalmDriver

`func (o *WritablePlatform) GetNapalmDriver() string`

GetNapalmDriver returns the NapalmDriver field if non-nil, zero value otherwise.

### GetNapalmDriverOk

`func (o *WritablePlatform) GetNapalmDriverOk() (*string, bool)`

GetNapalmDriverOk returns a tuple with the NapalmDriver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNapalmDriver

`func (o *WritablePlatform) SetNapalmDriver(v string)`

SetNapalmDriver sets NapalmDriver field to given value.

### HasNapalmDriver

`func (o *WritablePlatform) HasNapalmDriver() bool`

HasNapalmDriver returns a boolean if a field has been set.

### GetNapalmArgs

`func (o *WritablePlatform) GetNapalmArgs() string`

GetNapalmArgs returns the NapalmArgs field if non-nil, zero value otherwise.

### GetNapalmArgsOk

`func (o *WritablePlatform) GetNapalmArgsOk() (*string, bool)`

GetNapalmArgsOk returns a tuple with the NapalmArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNapalmArgs

`func (o *WritablePlatform) SetNapalmArgs(v string)`

SetNapalmArgs sets NapalmArgs field to given value.

### HasNapalmArgs

`func (o *WritablePlatform) HasNapalmArgs() bool`

HasNapalmArgs returns a boolean if a field has been set.

### SetNapalmArgsNil

`func (o *WritablePlatform) SetNapalmArgsNil(b bool)`

 SetNapalmArgsNil sets the value for NapalmArgs to be an explicit nil

### UnsetNapalmArgs
`func (o *WritablePlatform) UnsetNapalmArgs()`

UnsetNapalmArgs ensures that no value is present for NapalmArgs, not even an explicit nil
### GetDescription

`func (o *WritablePlatform) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WritablePlatform) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WritablePlatform) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WritablePlatform) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCustomFields

`func (o *WritablePlatform) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *WritablePlatform) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *WritablePlatform) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *WritablePlatform) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *WritablePlatform) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *WritablePlatform) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *WritablePlatform) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *WritablePlatform) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *WritablePlatform) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *WritablePlatform) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *WritablePlatform) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *WritablePlatform) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetDeviceCount

`func (o *WritablePlatform) GetDeviceCount() int32`

GetDeviceCount returns the DeviceCount field if non-nil, zero value otherwise.

### GetDeviceCountOk

`func (o *WritablePlatform) GetDeviceCountOk() (*int32, bool)`

GetDeviceCountOk returns a tuple with the DeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceCount

`func (o *WritablePlatform) SetDeviceCount(v int32)`

SetDeviceCount sets DeviceCount field to given value.

### HasDeviceCount

`func (o *WritablePlatform) HasDeviceCount() bool`

HasDeviceCount returns a boolean if a field has been set.

### GetVirtualmachineCount

`func (o *WritablePlatform) GetVirtualmachineCount() int32`

GetVirtualmachineCount returns the VirtualmachineCount field if non-nil, zero value otherwise.

### GetVirtualmachineCountOk

`func (o *WritablePlatform) GetVirtualmachineCountOk() (*int32, bool)`

GetVirtualmachineCountOk returns a tuple with the VirtualmachineCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualmachineCount

`func (o *WritablePlatform) SetVirtualmachineCount(v int32)`

SetVirtualmachineCount sets VirtualmachineCount field to given value.

### HasVirtualmachineCount

`func (o *WritablePlatform) HasVirtualmachineCount() bool`

HasVirtualmachineCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


