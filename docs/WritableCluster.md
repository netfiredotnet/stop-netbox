# WritableCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Display** | Pointer to **string** |  | [optional] [readonly] 
**Name** | **string** |  | 
**Type** | **int32** |  | 
**Group** | Pointer to **NullableInt32** |  | [optional] 
**Tenant** | Pointer to **NullableInt32** |  | [optional] 
**Site** | Pointer to **NullableInt32** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTag**](NestedTag.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] [readonly] 
**DeviceCount** | Pointer to **int32** |  | [optional] [readonly] 
**VirtualmachineCount** | Pointer to **int32** |  | [optional] [readonly] 

## Methods

### NewWritableCluster

`func NewWritableCluster(name string, type_ int32, ) *WritableCluster`

NewWritableCluster instantiates a new WritableCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritableClusterWithDefaults

`func NewWritableClusterWithDefaults() *WritableCluster`

NewWritableClusterWithDefaults instantiates a new WritableCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WritableCluster) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WritableCluster) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WritableCluster) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *WritableCluster) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *WritableCluster) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WritableCluster) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WritableCluster) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *WritableCluster) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDisplay

`func (o *WritableCluster) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *WritableCluster) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *WritableCluster) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *WritableCluster) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetName

`func (o *WritableCluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WritableCluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WritableCluster) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *WritableCluster) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WritableCluster) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WritableCluster) SetType(v int32)`

SetType sets Type field to given value.


### GetGroup

`func (o *WritableCluster) GetGroup() int32`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *WritableCluster) GetGroupOk() (*int32, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *WritableCluster) SetGroup(v int32)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *WritableCluster) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *WritableCluster) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *WritableCluster) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetTenant

`func (o *WritableCluster) GetTenant() int32`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *WritableCluster) GetTenantOk() (*int32, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *WritableCluster) SetTenant(v int32)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *WritableCluster) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *WritableCluster) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *WritableCluster) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetSite

`func (o *WritableCluster) GetSite() int32`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *WritableCluster) GetSiteOk() (*int32, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *WritableCluster) SetSite(v int32)`

SetSite sets Site field to given value.

### HasSite

`func (o *WritableCluster) HasSite() bool`

HasSite returns a boolean if a field has been set.

### SetSiteNil

`func (o *WritableCluster) SetSiteNil(b bool)`

 SetSiteNil sets the value for Site to be an explicit nil

### UnsetSite
`func (o *WritableCluster) UnsetSite()`

UnsetSite ensures that no value is present for Site, not even an explicit nil
### GetComments

`func (o *WritableCluster) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *WritableCluster) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *WritableCluster) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *WritableCluster) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *WritableCluster) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WritableCluster) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WritableCluster) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WritableCluster) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *WritableCluster) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *WritableCluster) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *WritableCluster) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *WritableCluster) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *WritableCluster) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *WritableCluster) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *WritableCluster) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *WritableCluster) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *WritableCluster) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *WritableCluster) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *WritableCluster) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *WritableCluster) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetDeviceCount

`func (o *WritableCluster) GetDeviceCount() int32`

GetDeviceCount returns the DeviceCount field if non-nil, zero value otherwise.

### GetDeviceCountOk

`func (o *WritableCluster) GetDeviceCountOk() (*int32, bool)`

GetDeviceCountOk returns a tuple with the DeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceCount

`func (o *WritableCluster) SetDeviceCount(v int32)`

SetDeviceCount sets DeviceCount field to given value.

### HasDeviceCount

`func (o *WritableCluster) HasDeviceCount() bool`

HasDeviceCount returns a boolean if a field has been set.

### GetVirtualmachineCount

`func (o *WritableCluster) GetVirtualmachineCount() int32`

GetVirtualmachineCount returns the VirtualmachineCount field if non-nil, zero value otherwise.

### GetVirtualmachineCountOk

`func (o *WritableCluster) GetVirtualmachineCountOk() (*int32, bool)`

GetVirtualmachineCountOk returns a tuple with the VirtualmachineCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualmachineCount

`func (o *WritableCluster) SetVirtualmachineCount(v int32)`

SetVirtualmachineCount sets VirtualmachineCount field to given value.

### HasVirtualmachineCount

`func (o *WritableCluster) HasVirtualmachineCount() bool`

HasVirtualmachineCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


