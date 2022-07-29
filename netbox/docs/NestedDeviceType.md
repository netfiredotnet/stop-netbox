# NestedDeviceType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Display** | Pointer to **string** |  | [optional] [readonly] 
**Manufacturer** | Pointer to [**NestedManufacturer**](NestedManufacturer.md) |  | [optional] 
**Model** | **string** |  | 
**Slug** | **string** |  | 
**DisplayName** | Pointer to **string** |  | [optional] [readonly] 
**DeviceCount** | Pointer to **int32** |  | [optional] [readonly] 

## Methods

### NewNestedDeviceType

`func NewNestedDeviceType(model string, slug string, ) *NestedDeviceType`

NewNestedDeviceType instantiates a new NestedDeviceType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNestedDeviceTypeWithDefaults

`func NewNestedDeviceTypeWithDefaults() *NestedDeviceType`

NewNestedDeviceTypeWithDefaults instantiates a new NestedDeviceType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NestedDeviceType) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NestedDeviceType) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NestedDeviceType) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *NestedDeviceType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *NestedDeviceType) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NestedDeviceType) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NestedDeviceType) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *NestedDeviceType) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDisplay

`func (o *NestedDeviceType) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *NestedDeviceType) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *NestedDeviceType) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *NestedDeviceType) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetManufacturer

`func (o *NestedDeviceType) GetManufacturer() NestedManufacturer`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *NestedDeviceType) GetManufacturerOk() (*NestedManufacturer, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *NestedDeviceType) SetManufacturer(v NestedManufacturer)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *NestedDeviceType) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### GetModel

`func (o *NestedDeviceType) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *NestedDeviceType) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *NestedDeviceType) SetModel(v string)`

SetModel sets Model field to given value.


### GetSlug

`func (o *NestedDeviceType) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *NestedDeviceType) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *NestedDeviceType) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetDisplayName

`func (o *NestedDeviceType) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *NestedDeviceType) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *NestedDeviceType) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *NestedDeviceType) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDeviceCount

`func (o *NestedDeviceType) GetDeviceCount() int32`

GetDeviceCount returns the DeviceCount field if non-nil, zero value otherwise.

### GetDeviceCountOk

`func (o *NestedDeviceType) GetDeviceCountOk() (*int32, bool)`

GetDeviceCountOk returns a tuple with the DeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceCount

`func (o *NestedDeviceType) SetDeviceCount(v int32)`

SetDeviceCount sets DeviceCount field to given value.

### HasDeviceCount

`func (o *NestedDeviceType) HasDeviceCount() bool`

HasDeviceCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


