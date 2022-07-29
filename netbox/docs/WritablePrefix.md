# WritablePrefix

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Display** | Pointer to **string** |  | [optional] [readonly] 
**Family** | Pointer to **string** |  | [optional] [readonly] 
**Prefix** | **string** | IPv4 or IPv6 network with mask | 
**Site** | Pointer to **NullableInt32** |  | [optional] 
**Vrf** | Pointer to **NullableInt32** |  | [optional] 
**Tenant** | Pointer to **NullableInt32** |  | [optional] 
**Vlan** | Pointer to **NullableInt32** |  | [optional] 
**Status** | Pointer to **string** | Operational status of this prefix | [optional] 
**Role** | Pointer to **NullableInt32** | The primary function of this prefix | [optional] 
**IsPool** | Pointer to **bool** | All IP addresses within this prefix are considered usable | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTag**](NestedTag.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] [readonly] 
**Children** | Pointer to **int32** |  | [optional] [readonly] 
**Depth** | Pointer to **int32** |  | [optional] [readonly] 

## Methods

### NewWritablePrefix

`func NewWritablePrefix(prefix string, ) *WritablePrefix`

NewWritablePrefix instantiates a new WritablePrefix object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritablePrefixWithDefaults

`func NewWritablePrefixWithDefaults() *WritablePrefix`

NewWritablePrefixWithDefaults instantiates a new WritablePrefix object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WritablePrefix) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WritablePrefix) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WritablePrefix) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *WritablePrefix) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *WritablePrefix) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WritablePrefix) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WritablePrefix) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *WritablePrefix) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDisplay

`func (o *WritablePrefix) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *WritablePrefix) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *WritablePrefix) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *WritablePrefix) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetFamily

`func (o *WritablePrefix) GetFamily() string`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *WritablePrefix) GetFamilyOk() (*string, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *WritablePrefix) SetFamily(v string)`

SetFamily sets Family field to given value.

### HasFamily

`func (o *WritablePrefix) HasFamily() bool`

HasFamily returns a boolean if a field has been set.

### GetPrefix

`func (o *WritablePrefix) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *WritablePrefix) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *WritablePrefix) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.


### GetSite

`func (o *WritablePrefix) GetSite() int32`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *WritablePrefix) GetSiteOk() (*int32, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *WritablePrefix) SetSite(v int32)`

SetSite sets Site field to given value.

### HasSite

`func (o *WritablePrefix) HasSite() bool`

HasSite returns a boolean if a field has been set.

### SetSiteNil

`func (o *WritablePrefix) SetSiteNil(b bool)`

 SetSiteNil sets the value for Site to be an explicit nil

### UnsetSite
`func (o *WritablePrefix) UnsetSite()`

UnsetSite ensures that no value is present for Site, not even an explicit nil
### GetVrf

`func (o *WritablePrefix) GetVrf() int32`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *WritablePrefix) GetVrfOk() (*int32, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *WritablePrefix) SetVrf(v int32)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *WritablePrefix) HasVrf() bool`

HasVrf returns a boolean if a field has been set.

### SetVrfNil

`func (o *WritablePrefix) SetVrfNil(b bool)`

 SetVrfNil sets the value for Vrf to be an explicit nil

### UnsetVrf
`func (o *WritablePrefix) UnsetVrf()`

UnsetVrf ensures that no value is present for Vrf, not even an explicit nil
### GetTenant

`func (o *WritablePrefix) GetTenant() int32`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *WritablePrefix) GetTenantOk() (*int32, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *WritablePrefix) SetTenant(v int32)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *WritablePrefix) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *WritablePrefix) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *WritablePrefix) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetVlan

`func (o *WritablePrefix) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *WritablePrefix) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *WritablePrefix) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *WritablePrefix) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### SetVlanNil

`func (o *WritablePrefix) SetVlanNil(b bool)`

 SetVlanNil sets the value for Vlan to be an explicit nil

### UnsetVlan
`func (o *WritablePrefix) UnsetVlan()`

UnsetVlan ensures that no value is present for Vlan, not even an explicit nil
### GetStatus

`func (o *WritablePrefix) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WritablePrefix) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WritablePrefix) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WritablePrefix) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRole

`func (o *WritablePrefix) GetRole() int32`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *WritablePrefix) GetRoleOk() (*int32, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *WritablePrefix) SetRole(v int32)`

SetRole sets Role field to given value.

### HasRole

`func (o *WritablePrefix) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *WritablePrefix) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *WritablePrefix) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetIsPool

`func (o *WritablePrefix) GetIsPool() bool`

GetIsPool returns the IsPool field if non-nil, zero value otherwise.

### GetIsPoolOk

`func (o *WritablePrefix) GetIsPoolOk() (*bool, bool)`

GetIsPoolOk returns a tuple with the IsPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPool

`func (o *WritablePrefix) SetIsPool(v bool)`

SetIsPool sets IsPool field to given value.

### HasIsPool

`func (o *WritablePrefix) HasIsPool() bool`

HasIsPool returns a boolean if a field has been set.

### GetDescription

`func (o *WritablePrefix) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WritablePrefix) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WritablePrefix) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WritablePrefix) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *WritablePrefix) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WritablePrefix) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WritablePrefix) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WritablePrefix) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *WritablePrefix) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *WritablePrefix) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *WritablePrefix) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *WritablePrefix) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *WritablePrefix) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *WritablePrefix) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *WritablePrefix) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *WritablePrefix) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *WritablePrefix) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *WritablePrefix) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *WritablePrefix) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *WritablePrefix) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetChildren

`func (o *WritablePrefix) GetChildren() int32`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *WritablePrefix) GetChildrenOk() (*int32, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *WritablePrefix) SetChildren(v int32)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *WritablePrefix) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetDepth

`func (o *WritablePrefix) GetDepth() int32`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *WritablePrefix) GetDepthOk() (*int32, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *WritablePrefix) SetDepth(v int32)`

SetDepth sets Depth field to given value.

### HasDepth

`func (o *WritablePrefix) HasDepth() bool`

HasDepth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


