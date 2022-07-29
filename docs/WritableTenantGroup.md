# WritableTenantGroup

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
**TenantCount** | Pointer to **int32** |  | [optional] [readonly] 
**Depth** | Pointer to **int32** |  | [optional] [readonly] 

## Methods

### NewWritableTenantGroup

`func NewWritableTenantGroup(name string, slug string, ) *WritableTenantGroup`

NewWritableTenantGroup instantiates a new WritableTenantGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritableTenantGroupWithDefaults

`func NewWritableTenantGroupWithDefaults() *WritableTenantGroup`

NewWritableTenantGroupWithDefaults instantiates a new WritableTenantGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WritableTenantGroup) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WritableTenantGroup) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WritableTenantGroup) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *WritableTenantGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *WritableTenantGroup) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WritableTenantGroup) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WritableTenantGroup) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *WritableTenantGroup) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDisplay

`func (o *WritableTenantGroup) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *WritableTenantGroup) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *WritableTenantGroup) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *WritableTenantGroup) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetName

`func (o *WritableTenantGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WritableTenantGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WritableTenantGroup) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *WritableTenantGroup) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *WritableTenantGroup) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *WritableTenantGroup) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetParent

`func (o *WritableTenantGroup) GetParent() int32`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *WritableTenantGroup) GetParentOk() (*int32, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *WritableTenantGroup) SetParent(v int32)`

SetParent sets Parent field to given value.

### HasParent

`func (o *WritableTenantGroup) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *WritableTenantGroup) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *WritableTenantGroup) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetDescription

`func (o *WritableTenantGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WritableTenantGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WritableTenantGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WritableTenantGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCustomFields

`func (o *WritableTenantGroup) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *WritableTenantGroup) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *WritableTenantGroup) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *WritableTenantGroup) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *WritableTenantGroup) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *WritableTenantGroup) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *WritableTenantGroup) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *WritableTenantGroup) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *WritableTenantGroup) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *WritableTenantGroup) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *WritableTenantGroup) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *WritableTenantGroup) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetTenantCount

`func (o *WritableTenantGroup) GetTenantCount() int32`

GetTenantCount returns the TenantCount field if non-nil, zero value otherwise.

### GetTenantCountOk

`func (o *WritableTenantGroup) GetTenantCountOk() (*int32, bool)`

GetTenantCountOk returns a tuple with the TenantCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantCount

`func (o *WritableTenantGroup) SetTenantCount(v int32)`

SetTenantCount sets TenantCount field to given value.

### HasTenantCount

`func (o *WritableTenantGroup) HasTenantCount() bool`

HasTenantCount returns a boolean if a field has been set.

### GetDepth

`func (o *WritableTenantGroup) GetDepth() int32`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *WritableTenantGroup) GetDepthOk() (*int32, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *WritableTenantGroup) SetDepth(v int32)`

SetDepth sets Depth field to given value.

### HasDepth

`func (o *WritableTenantGroup) HasDepth() bool`

HasDepth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


