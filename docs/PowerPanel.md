# PowerPanel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Display** | Pointer to **string** |  | [optional] [readonly] 
**Site** | [**NestedSite**](NestedSite.md) |  | 
**Location** | Pointer to [**NullableNestedLocation**](NestedLocation.md) |  | [optional] 
**Name** | **string** |  | 
**Tags** | Pointer to [**[]NestedTag**](NestedTag.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**PowerfeedCount** | Pointer to **int32** |  | [optional] [readonly] 

## Methods

### NewPowerPanel

`func NewPowerPanel(site NestedSite, name string, ) *PowerPanel`

NewPowerPanel instantiates a new PowerPanel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPowerPanelWithDefaults

`func NewPowerPanelWithDefaults() *PowerPanel`

NewPowerPanelWithDefaults instantiates a new PowerPanel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PowerPanel) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PowerPanel) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PowerPanel) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PowerPanel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *PowerPanel) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PowerPanel) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PowerPanel) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PowerPanel) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDisplay

`func (o *PowerPanel) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *PowerPanel) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *PowerPanel) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *PowerPanel) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetSite

`func (o *PowerPanel) GetSite() NestedSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *PowerPanel) GetSiteOk() (*NestedSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *PowerPanel) SetSite(v NestedSite)`

SetSite sets Site field to given value.


### GetLocation

`func (o *PowerPanel) GetLocation() NestedLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PowerPanel) GetLocationOk() (*NestedLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PowerPanel) SetLocation(v NestedLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *PowerPanel) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *PowerPanel) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *PowerPanel) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetName

`func (o *PowerPanel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PowerPanel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PowerPanel) SetName(v string)`

SetName sets Name field to given value.


### GetTags

`func (o *PowerPanel) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PowerPanel) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PowerPanel) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PowerPanel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PowerPanel) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PowerPanel) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PowerPanel) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PowerPanel) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetPowerfeedCount

`func (o *PowerPanel) GetPowerfeedCount() int32`

GetPowerfeedCount returns the PowerfeedCount field if non-nil, zero value otherwise.

### GetPowerfeedCountOk

`func (o *PowerPanel) GetPowerfeedCountOk() (*int32, bool)`

GetPowerfeedCountOk returns a tuple with the PowerfeedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerfeedCount

`func (o *PowerPanel) SetPowerfeedCount(v int32)`

SetPowerfeedCount sets PowerfeedCount field to given value.

### HasPowerfeedCount

`func (o *PowerPanel) HasPowerfeedCount() bool`

HasPowerfeedCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


