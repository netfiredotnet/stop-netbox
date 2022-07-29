# VLANGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Display** | Pointer to **string** |  | [optional] [readonly] 
**Name** | **string** |  | 
**Slug** | **string** |  | 
**ScopeType** | Pointer to **string** |  | [optional] 
**ScopeId** | Pointer to **NullableInt32** |  | [optional] 
**Scope** | Pointer to **string** |  | [optional] [readonly] 
**Description** | Pointer to **string** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] [readonly] 
**VlanCount** | Pointer to **int32** |  | [optional] [readonly] 

## Methods

### NewVLANGroup

`func NewVLANGroup(name string, slug string, ) *VLANGroup`

NewVLANGroup instantiates a new VLANGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVLANGroupWithDefaults

`func NewVLANGroupWithDefaults() *VLANGroup`

NewVLANGroupWithDefaults instantiates a new VLANGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VLANGroup) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VLANGroup) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VLANGroup) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *VLANGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *VLANGroup) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VLANGroup) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VLANGroup) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *VLANGroup) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDisplay

`func (o *VLANGroup) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *VLANGroup) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *VLANGroup) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *VLANGroup) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetName

`func (o *VLANGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VLANGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VLANGroup) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *VLANGroup) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *VLANGroup) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *VLANGroup) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetScopeType

`func (o *VLANGroup) GetScopeType() string`

GetScopeType returns the ScopeType field if non-nil, zero value otherwise.

### GetScopeTypeOk

`func (o *VLANGroup) GetScopeTypeOk() (*string, bool)`

GetScopeTypeOk returns a tuple with the ScopeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeType

`func (o *VLANGroup) SetScopeType(v string)`

SetScopeType sets ScopeType field to given value.

### HasScopeType

`func (o *VLANGroup) HasScopeType() bool`

HasScopeType returns a boolean if a field has been set.

### GetScopeId

`func (o *VLANGroup) GetScopeId() int32`

GetScopeId returns the ScopeId field if non-nil, zero value otherwise.

### GetScopeIdOk

`func (o *VLANGroup) GetScopeIdOk() (*int32, bool)`

GetScopeIdOk returns a tuple with the ScopeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeId

`func (o *VLANGroup) SetScopeId(v int32)`

SetScopeId sets ScopeId field to given value.

### HasScopeId

`func (o *VLANGroup) HasScopeId() bool`

HasScopeId returns a boolean if a field has been set.

### SetScopeIdNil

`func (o *VLANGroup) SetScopeIdNil(b bool)`

 SetScopeIdNil sets the value for ScopeId to be an explicit nil

### UnsetScopeId
`func (o *VLANGroup) UnsetScopeId()`

UnsetScopeId ensures that no value is present for ScopeId, not even an explicit nil
### GetScope

`func (o *VLANGroup) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *VLANGroup) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *VLANGroup) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *VLANGroup) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetDescription

`func (o *VLANGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VLANGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VLANGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VLANGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCustomFields

`func (o *VLANGroup) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *VLANGroup) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *VLANGroup) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *VLANGroup) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *VLANGroup) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *VLANGroup) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *VLANGroup) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *VLANGroup) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *VLANGroup) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *VLANGroup) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *VLANGroup) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *VLANGroup) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetVlanCount

`func (o *VLANGroup) GetVlanCount() int32`

GetVlanCount returns the VlanCount field if non-nil, zero value otherwise.

### GetVlanCountOk

`func (o *VLANGroup) GetVlanCountOk() (*int32, bool)`

GetVlanCountOk returns a tuple with the VlanCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanCount

`func (o *VLANGroup) SetVlanCount(v int32)`

SetVlanCount sets VlanCount field to given value.

### HasVlanCount

`func (o *VLANGroup) HasVlanCount() bool`

HasVlanCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


