# WritableSiteGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Display** | Pointer to **string** |  | [optional] [readonly] 
**Name** | **string** |  | 
**Slug** | **string** |  | 
**Parent** | Pointer to **NullableInt32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] [readonly] 
**SiteCount** | Pointer to **int32** |  | [optional] [readonly] 
**Depth** | Pointer to **int32** |  | [optional] [readonly] 

## Methods

### NewWritableSiteGroup

`func NewWritableSiteGroup(name string, slug string, ) *WritableSiteGroup`

NewWritableSiteGroup instantiates a new WritableSiteGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritableSiteGroupWithDefaults

`func NewWritableSiteGroupWithDefaults() *WritableSiteGroup`

NewWritableSiteGroupWithDefaults instantiates a new WritableSiteGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WritableSiteGroup) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WritableSiteGroup) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WritableSiteGroup) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *WritableSiteGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *WritableSiteGroup) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WritableSiteGroup) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WritableSiteGroup) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *WritableSiteGroup) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDisplay

`func (o *WritableSiteGroup) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *WritableSiteGroup) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *WritableSiteGroup) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *WritableSiteGroup) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetName

`func (o *WritableSiteGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WritableSiteGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WritableSiteGroup) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *WritableSiteGroup) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *WritableSiteGroup) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *WritableSiteGroup) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetParent

`func (o *WritableSiteGroup) GetParent() int32`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *WritableSiteGroup) GetParentOk() (*int32, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *WritableSiteGroup) SetParent(v int32)`

SetParent sets Parent field to given value.

### HasParent

`func (o *WritableSiteGroup) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *WritableSiteGroup) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *WritableSiteGroup) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetDescription

`func (o *WritableSiteGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WritableSiteGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WritableSiteGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WritableSiteGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCustomFields

`func (o *WritableSiteGroup) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *WritableSiteGroup) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *WritableSiteGroup) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *WritableSiteGroup) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *WritableSiteGroup) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *WritableSiteGroup) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *WritableSiteGroup) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *WritableSiteGroup) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *WritableSiteGroup) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *WritableSiteGroup) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *WritableSiteGroup) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *WritableSiteGroup) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetSiteCount

`func (o *WritableSiteGroup) GetSiteCount() int32`

GetSiteCount returns the SiteCount field if non-nil, zero value otherwise.

### GetSiteCountOk

`func (o *WritableSiteGroup) GetSiteCountOk() (*int32, bool)`

GetSiteCountOk returns a tuple with the SiteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteCount

`func (o *WritableSiteGroup) SetSiteCount(v int32)`

SetSiteCount sets SiteCount field to given value.

### HasSiteCount

`func (o *WritableSiteGroup) HasSiteCount() bool`

HasSiteCount returns a boolean if a field has been set.

### GetDepth

`func (o *WritableSiteGroup) GetDepth() int32`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *WritableSiteGroup) GetDepthOk() (*int32, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *WritableSiteGroup) SetDepth(v int32)`

SetDepth sets Depth field to given value.

### HasDepth

`func (o *WritableSiteGroup) HasDepth() bool`

HasDepth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


