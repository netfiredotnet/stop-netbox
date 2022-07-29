# SiteGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Display** | Pointer to **string** |  | [optional] [readonly] 
**Name** | **string** |  | 
**Slug** | **string** |  | 
**Parent** | Pointer to [**NullableNestedSiteGroup**](NestedSiteGroup.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] [readonly] 
**SiteCount** | Pointer to **int32** |  | [optional] [readonly] 
**Depth** | Pointer to **int32** |  | [optional] [readonly] 

## Methods

### NewSiteGroup

`func NewSiteGroup(name string, slug string, ) *SiteGroup`

NewSiteGroup instantiates a new SiteGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteGroupWithDefaults

`func NewSiteGroupWithDefaults() *SiteGroup`

NewSiteGroupWithDefaults instantiates a new SiteGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SiteGroup) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SiteGroup) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SiteGroup) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SiteGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *SiteGroup) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SiteGroup) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SiteGroup) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SiteGroup) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDisplay

`func (o *SiteGroup) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *SiteGroup) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *SiteGroup) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *SiteGroup) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetName

`func (o *SiteGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SiteGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SiteGroup) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *SiteGroup) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *SiteGroup) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *SiteGroup) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetParent

`func (o *SiteGroup) GetParent() NestedSiteGroup`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *SiteGroup) GetParentOk() (*NestedSiteGroup, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *SiteGroup) SetParent(v NestedSiteGroup)`

SetParent sets Parent field to given value.

### HasParent

`func (o *SiteGroup) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *SiteGroup) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *SiteGroup) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetDescription

`func (o *SiteGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SiteGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SiteGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SiteGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCustomFields

`func (o *SiteGroup) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *SiteGroup) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *SiteGroup) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *SiteGroup) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *SiteGroup) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *SiteGroup) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *SiteGroup) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *SiteGroup) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *SiteGroup) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *SiteGroup) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *SiteGroup) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *SiteGroup) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetSiteCount

`func (o *SiteGroup) GetSiteCount() int32`

GetSiteCount returns the SiteCount field if non-nil, zero value otherwise.

### GetSiteCountOk

`func (o *SiteGroup) GetSiteCountOk() (*int32, bool)`

GetSiteCountOk returns a tuple with the SiteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteCount

`func (o *SiteGroup) SetSiteCount(v int32)`

SetSiteCount sets SiteCount field to given value.

### HasSiteCount

`func (o *SiteGroup) HasSiteCount() bool`

HasSiteCount returns a boolean if a field has been set.

### GetDepth

`func (o *SiteGroup) GetDepth() int32`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *SiteGroup) GetDepthOk() (*int32, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *SiteGroup) SetDepth(v int32)`

SetDepth sets Depth field to given value.

### HasDepth

`func (o *SiteGroup) HasDepth() bool`

HasDepth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


