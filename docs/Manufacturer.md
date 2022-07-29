# Manufacturer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Display** | Pointer to **string** |  | [optional] [readonly] 
**Name** | **string** |  | 
**Slug** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] [readonly] 
**DevicetypeCount** | Pointer to **int32** |  | [optional] [readonly] 
**InventoryitemCount** | Pointer to **int32** |  | [optional] [readonly] 
**PlatformCount** | Pointer to **int32** |  | [optional] [readonly] 

## Methods

### NewManufacturer

`func NewManufacturer(name string, slug string, ) *Manufacturer`

NewManufacturer instantiates a new Manufacturer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManufacturerWithDefaults

`func NewManufacturerWithDefaults() *Manufacturer`

NewManufacturerWithDefaults instantiates a new Manufacturer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Manufacturer) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Manufacturer) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Manufacturer) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Manufacturer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *Manufacturer) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Manufacturer) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Manufacturer) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Manufacturer) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDisplay

`func (o *Manufacturer) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *Manufacturer) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *Manufacturer) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *Manufacturer) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetName

`func (o *Manufacturer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Manufacturer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Manufacturer) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *Manufacturer) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *Manufacturer) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *Manufacturer) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetDescription

`func (o *Manufacturer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Manufacturer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Manufacturer) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Manufacturer) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCustomFields

`func (o *Manufacturer) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Manufacturer) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Manufacturer) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Manufacturer) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *Manufacturer) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Manufacturer) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Manufacturer) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Manufacturer) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *Manufacturer) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Manufacturer) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Manufacturer) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *Manufacturer) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetDevicetypeCount

`func (o *Manufacturer) GetDevicetypeCount() int32`

GetDevicetypeCount returns the DevicetypeCount field if non-nil, zero value otherwise.

### GetDevicetypeCountOk

`func (o *Manufacturer) GetDevicetypeCountOk() (*int32, bool)`

GetDevicetypeCountOk returns a tuple with the DevicetypeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicetypeCount

`func (o *Manufacturer) SetDevicetypeCount(v int32)`

SetDevicetypeCount sets DevicetypeCount field to given value.

### HasDevicetypeCount

`func (o *Manufacturer) HasDevicetypeCount() bool`

HasDevicetypeCount returns a boolean if a field has been set.

### GetInventoryitemCount

`func (o *Manufacturer) GetInventoryitemCount() int32`

GetInventoryitemCount returns the InventoryitemCount field if non-nil, zero value otherwise.

### GetInventoryitemCountOk

`func (o *Manufacturer) GetInventoryitemCountOk() (*int32, bool)`

GetInventoryitemCountOk returns a tuple with the InventoryitemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryitemCount

`func (o *Manufacturer) SetInventoryitemCount(v int32)`

SetInventoryitemCount sets InventoryitemCount field to given value.

### HasInventoryitemCount

`func (o *Manufacturer) HasInventoryitemCount() bool`

HasInventoryitemCount returns a boolean if a field has been set.

### GetPlatformCount

`func (o *Manufacturer) GetPlatformCount() int32`

GetPlatformCount returns the PlatformCount field if non-nil, zero value otherwise.

### GetPlatformCountOk

`func (o *Manufacturer) GetPlatformCountOk() (*int32, bool)`

GetPlatformCountOk returns a tuple with the PlatformCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformCount

`func (o *Manufacturer) SetPlatformCount(v int32)`

SetPlatformCount sets PlatformCount field to given value.

### HasPlatformCount

`func (o *Manufacturer) HasPlatformCount() bool`

HasPlatformCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


