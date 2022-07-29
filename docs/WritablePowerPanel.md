# WritablePowerPanel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Display** | Pointer to **string** |  | [optional] [readonly] 
**Site** | **int32** |  | 
**Location** | Pointer to **NullableInt32** |  | [optional] 
**Name** | **string** |  | 
**Tags** | Pointer to [**[]NestedTag**](NestedTag.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**PowerfeedCount** | Pointer to **int32** |  | [optional] [readonly] 

## Methods

### NewWritablePowerPanel

`func NewWritablePowerPanel(site int32, name string, ) *WritablePowerPanel`

NewWritablePowerPanel instantiates a new WritablePowerPanel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritablePowerPanelWithDefaults

`func NewWritablePowerPanelWithDefaults() *WritablePowerPanel`

NewWritablePowerPanelWithDefaults instantiates a new WritablePowerPanel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WritablePowerPanel) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WritablePowerPanel) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WritablePowerPanel) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *WritablePowerPanel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *WritablePowerPanel) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WritablePowerPanel) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WritablePowerPanel) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *WritablePowerPanel) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDisplay

`func (o *WritablePowerPanel) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *WritablePowerPanel) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *WritablePowerPanel) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *WritablePowerPanel) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetSite

`func (o *WritablePowerPanel) GetSite() int32`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *WritablePowerPanel) GetSiteOk() (*int32, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *WritablePowerPanel) SetSite(v int32)`

SetSite sets Site field to given value.


### GetLocation

`func (o *WritablePowerPanel) GetLocation() int32`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *WritablePowerPanel) GetLocationOk() (*int32, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *WritablePowerPanel) SetLocation(v int32)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *WritablePowerPanel) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *WritablePowerPanel) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *WritablePowerPanel) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetName

`func (o *WritablePowerPanel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WritablePowerPanel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WritablePowerPanel) SetName(v string)`

SetName sets Name field to given value.


### GetTags

`func (o *WritablePowerPanel) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WritablePowerPanel) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WritablePowerPanel) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WritablePowerPanel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *WritablePowerPanel) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *WritablePowerPanel) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *WritablePowerPanel) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *WritablePowerPanel) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetPowerfeedCount

`func (o *WritablePowerPanel) GetPowerfeedCount() int32`

GetPowerfeedCount returns the PowerfeedCount field if non-nil, zero value otherwise.

### GetPowerfeedCountOk

`func (o *WritablePowerPanel) GetPowerfeedCountOk() (*int32, bool)`

GetPowerfeedCountOk returns a tuple with the PowerfeedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerfeedCount

`func (o *WritablePowerPanel) SetPowerfeedCount(v int32)`

SetPowerfeedCount sets PowerfeedCount field to given value.

### HasPowerfeedCount

`func (o *WritablePowerPanel) HasPowerfeedCount() bool`

HasPowerfeedCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


