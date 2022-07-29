# VRF

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Display** | Pointer to **string** |  | [optional] [readonly] 
**Name** | **string** |  | 
**Rd** | Pointer to **NullableString** | Unique route distinguisher (as defined in RFC 4364) | [optional] 
**Tenant** | Pointer to [**NullableNestedTenant**](NestedTenant.md) |  | [optional] 
**EnforceUnique** | Pointer to **bool** | Prevent duplicate prefixes/IP addresses within this VRF | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ImportTargets** | Pointer to [**[]NestedRouteTarget**](NestedRouteTarget.md) |  | [optional] 
**ExportTargets** | Pointer to [**[]NestedRouteTarget**](NestedRouteTarget.md) |  | [optional] 
**Tags** | Pointer to [**[]NestedTag**](NestedTag.md) |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] [readonly] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] [readonly] 
**IpaddressCount** | Pointer to **int32** |  | [optional] [readonly] 
**PrefixCount** | Pointer to **int32** |  | [optional] [readonly] 

## Methods

### NewVRF

`func NewVRF(name string, ) *VRF`

NewVRF instantiates a new VRF object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVRFWithDefaults

`func NewVRFWithDefaults() *VRF`

NewVRFWithDefaults instantiates a new VRF object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VRF) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VRF) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VRF) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *VRF) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *VRF) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VRF) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VRF) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *VRF) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDisplay

`func (o *VRF) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *VRF) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *VRF) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *VRF) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetName

`func (o *VRF) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VRF) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VRF) SetName(v string)`

SetName sets Name field to given value.


### GetRd

`func (o *VRF) GetRd() string`

GetRd returns the Rd field if non-nil, zero value otherwise.

### GetRdOk

`func (o *VRF) GetRdOk() (*string, bool)`

GetRdOk returns a tuple with the Rd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRd

`func (o *VRF) SetRd(v string)`

SetRd sets Rd field to given value.

### HasRd

`func (o *VRF) HasRd() bool`

HasRd returns a boolean if a field has been set.

### SetRdNil

`func (o *VRF) SetRdNil(b bool)`

 SetRdNil sets the value for Rd to be an explicit nil

### UnsetRd
`func (o *VRF) UnsetRd()`

UnsetRd ensures that no value is present for Rd, not even an explicit nil
### GetTenant

`func (o *VRF) GetTenant() NestedTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *VRF) GetTenantOk() (*NestedTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *VRF) SetTenant(v NestedTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *VRF) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *VRF) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *VRF) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetEnforceUnique

`func (o *VRF) GetEnforceUnique() bool`

GetEnforceUnique returns the EnforceUnique field if non-nil, zero value otherwise.

### GetEnforceUniqueOk

`func (o *VRF) GetEnforceUniqueOk() (*bool, bool)`

GetEnforceUniqueOk returns a tuple with the EnforceUnique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceUnique

`func (o *VRF) SetEnforceUnique(v bool)`

SetEnforceUnique sets EnforceUnique field to given value.

### HasEnforceUnique

`func (o *VRF) HasEnforceUnique() bool`

HasEnforceUnique returns a boolean if a field has been set.

### GetDescription

`func (o *VRF) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VRF) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VRF) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VRF) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetImportTargets

`func (o *VRF) GetImportTargets() []NestedRouteTarget`

GetImportTargets returns the ImportTargets field if non-nil, zero value otherwise.

### GetImportTargetsOk

`func (o *VRF) GetImportTargetsOk() (*[]NestedRouteTarget, bool)`

GetImportTargetsOk returns a tuple with the ImportTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportTargets

`func (o *VRF) SetImportTargets(v []NestedRouteTarget)`

SetImportTargets sets ImportTargets field to given value.

### HasImportTargets

`func (o *VRF) HasImportTargets() bool`

HasImportTargets returns a boolean if a field has been set.

### GetExportTargets

`func (o *VRF) GetExportTargets() []NestedRouteTarget`

GetExportTargets returns the ExportTargets field if non-nil, zero value otherwise.

### GetExportTargetsOk

`func (o *VRF) GetExportTargetsOk() (*[]NestedRouteTarget, bool)`

GetExportTargetsOk returns a tuple with the ExportTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportTargets

`func (o *VRF) SetExportTargets(v []NestedRouteTarget)`

SetExportTargets sets ExportTargets field to given value.

### HasExportTargets

`func (o *VRF) HasExportTargets() bool`

HasExportTargets returns a boolean if a field has been set.

### GetTags

`func (o *VRF) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VRF) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VRF) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VRF) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetDisplayName

`func (o *VRF) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *VRF) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *VRF) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *VRF) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetCustomFields

`func (o *VRF) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *VRF) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *VRF) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *VRF) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *VRF) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *VRF) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *VRF) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *VRF) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *VRF) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *VRF) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *VRF) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *VRF) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetIpaddressCount

`func (o *VRF) GetIpaddressCount() int32`

GetIpaddressCount returns the IpaddressCount field if non-nil, zero value otherwise.

### GetIpaddressCountOk

`func (o *VRF) GetIpaddressCountOk() (*int32, bool)`

GetIpaddressCountOk returns a tuple with the IpaddressCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpaddressCount

`func (o *VRF) SetIpaddressCount(v int32)`

SetIpaddressCount sets IpaddressCount field to given value.

### HasIpaddressCount

`func (o *VRF) HasIpaddressCount() bool`

HasIpaddressCount returns a boolean if a field has been set.

### GetPrefixCount

`func (o *VRF) GetPrefixCount() int32`

GetPrefixCount returns the PrefixCount field if non-nil, zero value otherwise.

### GetPrefixCountOk

`func (o *VRF) GetPrefixCountOk() (*int32, bool)`

GetPrefixCountOk returns a tuple with the PrefixCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixCount

`func (o *VRF) SetPrefixCount(v int32)`

SetPrefixCount sets PrefixCount field to given value.

### HasPrefixCount

`func (o *VRF) HasPrefixCount() bool`

HasPrefixCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


