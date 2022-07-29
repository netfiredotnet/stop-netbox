# WritableLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Display** | Pointer to **string** |  | [optional] [readonly] 
**Name** | **string** |  | 
**Slug** | **string** |  | 
**Site** | **int32** |  | 
**Parent** | Pointer to **NullableInt32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] [readonly] 
**RackCount** | Pointer to **int32** |  | [optional] [readonly] 
**DeviceCount** | Pointer to **int32** |  | [optional] [readonly] 
**Depth** | Pointer to **int32** |  | [optional] [readonly] 

## Methods

### NewWritableLocation

`func NewWritableLocation(name string, slug string, site int32, ) *WritableLocation`

NewWritableLocation instantiates a new WritableLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritableLocationWithDefaults

`func NewWritableLocationWithDefaults() *WritableLocation`

NewWritableLocationWithDefaults instantiates a new WritableLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WritableLocation) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WritableLocation) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WritableLocation) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *WritableLocation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *WritableLocation) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WritableLocation) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WritableLocation) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *WritableLocation) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDisplay

`func (o *WritableLocation) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *WritableLocation) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *WritableLocation) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *WritableLocation) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetName

`func (o *WritableLocation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WritableLocation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WritableLocation) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *WritableLocation) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *WritableLocation) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *WritableLocation) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetSite

`func (o *WritableLocation) GetSite() int32`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *WritableLocation) GetSiteOk() (*int32, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *WritableLocation) SetSite(v int32)`

SetSite sets Site field to given value.


### GetParent

`func (o *WritableLocation) GetParent() int32`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *WritableLocation) GetParentOk() (*int32, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *WritableLocation) SetParent(v int32)`

SetParent sets Parent field to given value.

### HasParent

`func (o *WritableLocation) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *WritableLocation) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *WritableLocation) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetDescription

`func (o *WritableLocation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WritableLocation) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WritableLocation) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WritableLocation) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCustomFields

`func (o *WritableLocation) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *WritableLocation) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *WritableLocation) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *WritableLocation) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *WritableLocation) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *WritableLocation) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *WritableLocation) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *WritableLocation) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *WritableLocation) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *WritableLocation) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *WritableLocation) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *WritableLocation) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetRackCount

`func (o *WritableLocation) GetRackCount() int32`

GetRackCount returns the RackCount field if non-nil, zero value otherwise.

### GetRackCountOk

`func (o *WritableLocation) GetRackCountOk() (*int32, bool)`

GetRackCountOk returns a tuple with the RackCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRackCount

`func (o *WritableLocation) SetRackCount(v int32)`

SetRackCount sets RackCount field to given value.

### HasRackCount

`func (o *WritableLocation) HasRackCount() bool`

HasRackCount returns a boolean if a field has been set.

### GetDeviceCount

`func (o *WritableLocation) GetDeviceCount() int32`

GetDeviceCount returns the DeviceCount field if non-nil, zero value otherwise.

### GetDeviceCountOk

`func (o *WritableLocation) GetDeviceCountOk() (*int32, bool)`

GetDeviceCountOk returns a tuple with the DeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceCount

`func (o *WritableLocation) SetDeviceCount(v int32)`

SetDeviceCount sets DeviceCount field to given value.

### HasDeviceCount

`func (o *WritableLocation) HasDeviceCount() bool`

HasDeviceCount returns a boolean if a field has been set.

### GetDepth

`func (o *WritableLocation) GetDepth() int32`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *WritableLocation) GetDepthOk() (*int32, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *WritableLocation) SetDepth(v int32)`

SetDepth sets Depth field to given value.

### HasDepth

`func (o *WritableLocation) HasDepth() bool`

HasDepth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


