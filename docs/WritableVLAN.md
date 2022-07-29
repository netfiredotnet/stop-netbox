# WritableVLAN

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Display** | Pointer to **string** |  | [optional] [readonly] 
**Site** | Pointer to **NullableInt32** |  | [optional] 
**Group** | Pointer to **NullableInt32** |  | [optional] 
**Vid** | **int32** |  | 
**Name** | **string** |  | 
**Tenant** | Pointer to **NullableInt32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Role** | Pointer to **NullableInt32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTag**](NestedTag.md) |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] [readonly] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] [readonly] 
**PrefixCount** | Pointer to **int32** |  | [optional] [readonly] 

## Methods

### NewWritableVLAN

`func NewWritableVLAN(vid int32, name string, ) *WritableVLAN`

NewWritableVLAN instantiates a new WritableVLAN object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritableVLANWithDefaults

`func NewWritableVLANWithDefaults() *WritableVLAN`

NewWritableVLANWithDefaults instantiates a new WritableVLAN object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WritableVLAN) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WritableVLAN) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WritableVLAN) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *WritableVLAN) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *WritableVLAN) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WritableVLAN) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WritableVLAN) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *WritableVLAN) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDisplay

`func (o *WritableVLAN) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *WritableVLAN) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *WritableVLAN) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *WritableVLAN) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetSite

`func (o *WritableVLAN) GetSite() int32`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *WritableVLAN) GetSiteOk() (*int32, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *WritableVLAN) SetSite(v int32)`

SetSite sets Site field to given value.

### HasSite

`func (o *WritableVLAN) HasSite() bool`

HasSite returns a boolean if a field has been set.

### SetSiteNil

`func (o *WritableVLAN) SetSiteNil(b bool)`

 SetSiteNil sets the value for Site to be an explicit nil

### UnsetSite
`func (o *WritableVLAN) UnsetSite()`

UnsetSite ensures that no value is present for Site, not even an explicit nil
### GetGroup

`func (o *WritableVLAN) GetGroup() int32`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *WritableVLAN) GetGroupOk() (*int32, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *WritableVLAN) SetGroup(v int32)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *WritableVLAN) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *WritableVLAN) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *WritableVLAN) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetVid

`func (o *WritableVLAN) GetVid() int32`

GetVid returns the Vid field if non-nil, zero value otherwise.

### GetVidOk

`func (o *WritableVLAN) GetVidOk() (*int32, bool)`

GetVidOk returns a tuple with the Vid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVid

`func (o *WritableVLAN) SetVid(v int32)`

SetVid sets Vid field to given value.


### GetName

`func (o *WritableVLAN) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WritableVLAN) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WritableVLAN) SetName(v string)`

SetName sets Name field to given value.


### GetTenant

`func (o *WritableVLAN) GetTenant() int32`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *WritableVLAN) GetTenantOk() (*int32, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *WritableVLAN) SetTenant(v int32)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *WritableVLAN) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *WritableVLAN) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *WritableVLAN) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetStatus

`func (o *WritableVLAN) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WritableVLAN) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WritableVLAN) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WritableVLAN) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRole

`func (o *WritableVLAN) GetRole() int32`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *WritableVLAN) GetRoleOk() (*int32, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *WritableVLAN) SetRole(v int32)`

SetRole sets Role field to given value.

### HasRole

`func (o *WritableVLAN) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *WritableVLAN) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *WritableVLAN) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetDescription

`func (o *WritableVLAN) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WritableVLAN) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WritableVLAN) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WritableVLAN) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *WritableVLAN) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WritableVLAN) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WritableVLAN) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WritableVLAN) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetDisplayName

`func (o *WritableVLAN) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *WritableVLAN) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *WritableVLAN) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *WritableVLAN) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetCustomFields

`func (o *WritableVLAN) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *WritableVLAN) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *WritableVLAN) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *WritableVLAN) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *WritableVLAN) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *WritableVLAN) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *WritableVLAN) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *WritableVLAN) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *WritableVLAN) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *WritableVLAN) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *WritableVLAN) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *WritableVLAN) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetPrefixCount

`func (o *WritableVLAN) GetPrefixCount() int32`

GetPrefixCount returns the PrefixCount field if non-nil, zero value otherwise.

### GetPrefixCountOk

`func (o *WritableVLAN) GetPrefixCountOk() (*int32, bool)`

GetPrefixCountOk returns a tuple with the PrefixCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixCount

`func (o *WritableVLAN) SetPrefixCount(v int32)`

SetPrefixCount sets PrefixCount field to given value.

### HasPrefixCount

`func (o *WritableVLAN) HasPrefixCount() bool`

HasPrefixCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


