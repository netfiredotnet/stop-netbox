# WritableTenant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Display** | Pointer to **string** |  | [optional] [readonly] 
**Name** | **string** |  | 
**Slug** | **string** |  | 
**Group** | Pointer to **NullableInt32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTag**](NestedTag.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] [readonly] 
**CircuitCount** | Pointer to **int32** |  | [optional] [readonly] 
**DeviceCount** | Pointer to **int32** |  | [optional] [readonly] 
**IpaddressCount** | Pointer to **int32** |  | [optional] [readonly] 
**PrefixCount** | Pointer to **int32** |  | [optional] [readonly] 
**RackCount** | Pointer to **int32** |  | [optional] [readonly] 
**SiteCount** | Pointer to **int32** |  | [optional] [readonly] 
**VirtualmachineCount** | Pointer to **int32** |  | [optional] [readonly] 
**VlanCount** | Pointer to **int32** |  | [optional] [readonly] 
**VrfCount** | Pointer to **int32** |  | [optional] [readonly] 
**ClusterCount** | Pointer to **int32** |  | [optional] [readonly] 

## Methods

### NewWritableTenant

`func NewWritableTenant(name string, slug string, ) *WritableTenant`

NewWritableTenant instantiates a new WritableTenant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritableTenantWithDefaults

`func NewWritableTenantWithDefaults() *WritableTenant`

NewWritableTenantWithDefaults instantiates a new WritableTenant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WritableTenant) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WritableTenant) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WritableTenant) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *WritableTenant) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *WritableTenant) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WritableTenant) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WritableTenant) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *WritableTenant) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDisplay

`func (o *WritableTenant) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *WritableTenant) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *WritableTenant) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *WritableTenant) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetName

`func (o *WritableTenant) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WritableTenant) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WritableTenant) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *WritableTenant) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *WritableTenant) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *WritableTenant) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetGroup

`func (o *WritableTenant) GetGroup() int32`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *WritableTenant) GetGroupOk() (*int32, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *WritableTenant) SetGroup(v int32)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *WritableTenant) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *WritableTenant) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *WritableTenant) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetDescription

`func (o *WritableTenant) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WritableTenant) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WritableTenant) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WritableTenant) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComments

`func (o *WritableTenant) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *WritableTenant) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *WritableTenant) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *WritableTenant) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *WritableTenant) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WritableTenant) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WritableTenant) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WritableTenant) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *WritableTenant) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *WritableTenant) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *WritableTenant) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *WritableTenant) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *WritableTenant) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *WritableTenant) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *WritableTenant) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *WritableTenant) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *WritableTenant) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *WritableTenant) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *WritableTenant) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *WritableTenant) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetCircuitCount

`func (o *WritableTenant) GetCircuitCount() int32`

GetCircuitCount returns the CircuitCount field if non-nil, zero value otherwise.

### GetCircuitCountOk

`func (o *WritableTenant) GetCircuitCountOk() (*int32, bool)`

GetCircuitCountOk returns a tuple with the CircuitCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitCount

`func (o *WritableTenant) SetCircuitCount(v int32)`

SetCircuitCount sets CircuitCount field to given value.

### HasCircuitCount

`func (o *WritableTenant) HasCircuitCount() bool`

HasCircuitCount returns a boolean if a field has been set.

### GetDeviceCount

`func (o *WritableTenant) GetDeviceCount() int32`

GetDeviceCount returns the DeviceCount field if non-nil, zero value otherwise.

### GetDeviceCountOk

`func (o *WritableTenant) GetDeviceCountOk() (*int32, bool)`

GetDeviceCountOk returns a tuple with the DeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceCount

`func (o *WritableTenant) SetDeviceCount(v int32)`

SetDeviceCount sets DeviceCount field to given value.

### HasDeviceCount

`func (o *WritableTenant) HasDeviceCount() bool`

HasDeviceCount returns a boolean if a field has been set.

### GetIpaddressCount

`func (o *WritableTenant) GetIpaddressCount() int32`

GetIpaddressCount returns the IpaddressCount field if non-nil, zero value otherwise.

### GetIpaddressCountOk

`func (o *WritableTenant) GetIpaddressCountOk() (*int32, bool)`

GetIpaddressCountOk returns a tuple with the IpaddressCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpaddressCount

`func (o *WritableTenant) SetIpaddressCount(v int32)`

SetIpaddressCount sets IpaddressCount field to given value.

### HasIpaddressCount

`func (o *WritableTenant) HasIpaddressCount() bool`

HasIpaddressCount returns a boolean if a field has been set.

### GetPrefixCount

`func (o *WritableTenant) GetPrefixCount() int32`

GetPrefixCount returns the PrefixCount field if non-nil, zero value otherwise.

### GetPrefixCountOk

`func (o *WritableTenant) GetPrefixCountOk() (*int32, bool)`

GetPrefixCountOk returns a tuple with the PrefixCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixCount

`func (o *WritableTenant) SetPrefixCount(v int32)`

SetPrefixCount sets PrefixCount field to given value.

### HasPrefixCount

`func (o *WritableTenant) HasPrefixCount() bool`

HasPrefixCount returns a boolean if a field has been set.

### GetRackCount

`func (o *WritableTenant) GetRackCount() int32`

GetRackCount returns the RackCount field if non-nil, zero value otherwise.

### GetRackCountOk

`func (o *WritableTenant) GetRackCountOk() (*int32, bool)`

GetRackCountOk returns a tuple with the RackCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRackCount

`func (o *WritableTenant) SetRackCount(v int32)`

SetRackCount sets RackCount field to given value.

### HasRackCount

`func (o *WritableTenant) HasRackCount() bool`

HasRackCount returns a boolean if a field has been set.

### GetSiteCount

`func (o *WritableTenant) GetSiteCount() int32`

GetSiteCount returns the SiteCount field if non-nil, zero value otherwise.

### GetSiteCountOk

`func (o *WritableTenant) GetSiteCountOk() (*int32, bool)`

GetSiteCountOk returns a tuple with the SiteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteCount

`func (o *WritableTenant) SetSiteCount(v int32)`

SetSiteCount sets SiteCount field to given value.

### HasSiteCount

`func (o *WritableTenant) HasSiteCount() bool`

HasSiteCount returns a boolean if a field has been set.

### GetVirtualmachineCount

`func (o *WritableTenant) GetVirtualmachineCount() int32`

GetVirtualmachineCount returns the VirtualmachineCount field if non-nil, zero value otherwise.

### GetVirtualmachineCountOk

`func (o *WritableTenant) GetVirtualmachineCountOk() (*int32, bool)`

GetVirtualmachineCountOk returns a tuple with the VirtualmachineCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualmachineCount

`func (o *WritableTenant) SetVirtualmachineCount(v int32)`

SetVirtualmachineCount sets VirtualmachineCount field to given value.

### HasVirtualmachineCount

`func (o *WritableTenant) HasVirtualmachineCount() bool`

HasVirtualmachineCount returns a boolean if a field has been set.

### GetVlanCount

`func (o *WritableTenant) GetVlanCount() int32`

GetVlanCount returns the VlanCount field if non-nil, zero value otherwise.

### GetVlanCountOk

`func (o *WritableTenant) GetVlanCountOk() (*int32, bool)`

GetVlanCountOk returns a tuple with the VlanCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanCount

`func (o *WritableTenant) SetVlanCount(v int32)`

SetVlanCount sets VlanCount field to given value.

### HasVlanCount

`func (o *WritableTenant) HasVlanCount() bool`

HasVlanCount returns a boolean if a field has been set.

### GetVrfCount

`func (o *WritableTenant) GetVrfCount() int32`

GetVrfCount returns the VrfCount field if non-nil, zero value otherwise.

### GetVrfCountOk

`func (o *WritableTenant) GetVrfCountOk() (*int32, bool)`

GetVrfCountOk returns a tuple with the VrfCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfCount

`func (o *WritableTenant) SetVrfCount(v int32)`

SetVrfCount sets VrfCount field to given value.

### HasVrfCount

`func (o *WritableTenant) HasVrfCount() bool`

HasVrfCount returns a boolean if a field has been set.

### GetClusterCount

`func (o *WritableTenant) GetClusterCount() int32`

GetClusterCount returns the ClusterCount field if non-nil, zero value otherwise.

### GetClusterCountOk

`func (o *WritableTenant) GetClusterCountOk() (*int32, bool)`

GetClusterCountOk returns a tuple with the ClusterCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterCount

`func (o *WritableTenant) SetClusterCount(v int32)`

SetClusterCount sets ClusterCount field to given value.

### HasClusterCount

`func (o *WritableTenant) HasClusterCount() bool`

HasClusterCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


