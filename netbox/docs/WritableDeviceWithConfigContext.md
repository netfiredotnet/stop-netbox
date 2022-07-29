# WritableDeviceWithConfigContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Display** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **NullableString** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] [readonly] 
**DeviceType** | **int32** |  | 
**DeviceRole** | **int32** |  | 
**Tenant** | Pointer to **NullableInt32** |  | [optional] 
**Platform** | Pointer to **NullableInt32** |  | [optional] 
**Serial** | Pointer to **string** |  | [optional] 
**AssetTag** | Pointer to **NullableString** | A unique tag used to identify this device | [optional] 
**Site** | **int32** |  | 
**Location** | Pointer to **NullableInt32** |  | [optional] 
**Rack** | Pointer to **NullableInt32** |  | [optional] 
**Position** | Pointer to **NullableInt32** | The lowest-numbered unit occupied by the device | [optional] 
**Face** | Pointer to **string** |  | [optional] 
**ParentDevice** | Pointer to [**NestedDevice**](NestedDevice.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**PrimaryIp** | Pointer to **string** |  | [optional] [readonly] 
**PrimaryIp4** | Pointer to **NullableInt32** |  | [optional] 
**PrimaryIp6** | Pointer to **NullableInt32** |  | [optional] 
**Cluster** | Pointer to **NullableInt32** |  | [optional] 
**VirtualChassis** | Pointer to **NullableInt32** |  | [optional] 
**VcPosition** | Pointer to **NullableInt32** |  | [optional] 
**VcPriority** | Pointer to **NullableInt32** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**LocalContextData** | Pointer to **NullableString** |  | [optional] 
**Tags** | Pointer to [**[]NestedTag**](NestedTag.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**ConfigContext** | Pointer to **map[string]string** |  | [optional] [readonly] 
**Created** | Pointer to **string** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewWritableDeviceWithConfigContext

`func NewWritableDeviceWithConfigContext(deviceType int32, deviceRole int32, site int32, ) *WritableDeviceWithConfigContext`

NewWritableDeviceWithConfigContext instantiates a new WritableDeviceWithConfigContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritableDeviceWithConfigContextWithDefaults

`func NewWritableDeviceWithConfigContextWithDefaults() *WritableDeviceWithConfigContext`

NewWritableDeviceWithConfigContextWithDefaults instantiates a new WritableDeviceWithConfigContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WritableDeviceWithConfigContext) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WritableDeviceWithConfigContext) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WritableDeviceWithConfigContext) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *WritableDeviceWithConfigContext) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *WritableDeviceWithConfigContext) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WritableDeviceWithConfigContext) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WritableDeviceWithConfigContext) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *WritableDeviceWithConfigContext) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDisplay

`func (o *WritableDeviceWithConfigContext) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *WritableDeviceWithConfigContext) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *WritableDeviceWithConfigContext) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *WritableDeviceWithConfigContext) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetName

`func (o *WritableDeviceWithConfigContext) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WritableDeviceWithConfigContext) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WritableDeviceWithConfigContext) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WritableDeviceWithConfigContext) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *WritableDeviceWithConfigContext) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *WritableDeviceWithConfigContext) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDisplayName

`func (o *WritableDeviceWithConfigContext) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *WritableDeviceWithConfigContext) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *WritableDeviceWithConfigContext) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *WritableDeviceWithConfigContext) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDeviceType

`func (o *WritableDeviceWithConfigContext) GetDeviceType() int32`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *WritableDeviceWithConfigContext) GetDeviceTypeOk() (*int32, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *WritableDeviceWithConfigContext) SetDeviceType(v int32)`

SetDeviceType sets DeviceType field to given value.


### GetDeviceRole

`func (o *WritableDeviceWithConfigContext) GetDeviceRole() int32`

GetDeviceRole returns the DeviceRole field if non-nil, zero value otherwise.

### GetDeviceRoleOk

`func (o *WritableDeviceWithConfigContext) GetDeviceRoleOk() (*int32, bool)`

GetDeviceRoleOk returns a tuple with the DeviceRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceRole

`func (o *WritableDeviceWithConfigContext) SetDeviceRole(v int32)`

SetDeviceRole sets DeviceRole field to given value.


### GetTenant

`func (o *WritableDeviceWithConfigContext) GetTenant() int32`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *WritableDeviceWithConfigContext) GetTenantOk() (*int32, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *WritableDeviceWithConfigContext) SetTenant(v int32)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *WritableDeviceWithConfigContext) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *WritableDeviceWithConfigContext) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *WritableDeviceWithConfigContext) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetPlatform

`func (o *WritableDeviceWithConfigContext) GetPlatform() int32`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *WritableDeviceWithConfigContext) GetPlatformOk() (*int32, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *WritableDeviceWithConfigContext) SetPlatform(v int32)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *WritableDeviceWithConfigContext) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatformNil

`func (o *WritableDeviceWithConfigContext) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *WritableDeviceWithConfigContext) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetSerial

`func (o *WritableDeviceWithConfigContext) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *WritableDeviceWithConfigContext) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *WritableDeviceWithConfigContext) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *WritableDeviceWithConfigContext) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetAssetTag

`func (o *WritableDeviceWithConfigContext) GetAssetTag() string`

GetAssetTag returns the AssetTag field if non-nil, zero value otherwise.

### GetAssetTagOk

`func (o *WritableDeviceWithConfigContext) GetAssetTagOk() (*string, bool)`

GetAssetTagOk returns a tuple with the AssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTag

`func (o *WritableDeviceWithConfigContext) SetAssetTag(v string)`

SetAssetTag sets AssetTag field to given value.

### HasAssetTag

`func (o *WritableDeviceWithConfigContext) HasAssetTag() bool`

HasAssetTag returns a boolean if a field has been set.

### SetAssetTagNil

`func (o *WritableDeviceWithConfigContext) SetAssetTagNil(b bool)`

 SetAssetTagNil sets the value for AssetTag to be an explicit nil

### UnsetAssetTag
`func (o *WritableDeviceWithConfigContext) UnsetAssetTag()`

UnsetAssetTag ensures that no value is present for AssetTag, not even an explicit nil
### GetSite

`func (o *WritableDeviceWithConfigContext) GetSite() int32`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *WritableDeviceWithConfigContext) GetSiteOk() (*int32, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *WritableDeviceWithConfigContext) SetSite(v int32)`

SetSite sets Site field to given value.


### GetLocation

`func (o *WritableDeviceWithConfigContext) GetLocation() int32`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *WritableDeviceWithConfigContext) GetLocationOk() (*int32, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *WritableDeviceWithConfigContext) SetLocation(v int32)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *WritableDeviceWithConfigContext) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *WritableDeviceWithConfigContext) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *WritableDeviceWithConfigContext) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetRack

`func (o *WritableDeviceWithConfigContext) GetRack() int32`

GetRack returns the Rack field if non-nil, zero value otherwise.

### GetRackOk

`func (o *WritableDeviceWithConfigContext) GetRackOk() (*int32, bool)`

GetRackOk returns a tuple with the Rack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRack

`func (o *WritableDeviceWithConfigContext) SetRack(v int32)`

SetRack sets Rack field to given value.

### HasRack

`func (o *WritableDeviceWithConfigContext) HasRack() bool`

HasRack returns a boolean if a field has been set.

### SetRackNil

`func (o *WritableDeviceWithConfigContext) SetRackNil(b bool)`

 SetRackNil sets the value for Rack to be an explicit nil

### UnsetRack
`func (o *WritableDeviceWithConfigContext) UnsetRack()`

UnsetRack ensures that no value is present for Rack, not even an explicit nil
### GetPosition

`func (o *WritableDeviceWithConfigContext) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *WritableDeviceWithConfigContext) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *WritableDeviceWithConfigContext) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *WritableDeviceWithConfigContext) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### SetPositionNil

`func (o *WritableDeviceWithConfigContext) SetPositionNil(b bool)`

 SetPositionNil sets the value for Position to be an explicit nil

### UnsetPosition
`func (o *WritableDeviceWithConfigContext) UnsetPosition()`

UnsetPosition ensures that no value is present for Position, not even an explicit nil
### GetFace

`func (o *WritableDeviceWithConfigContext) GetFace() string`

GetFace returns the Face field if non-nil, zero value otherwise.

### GetFaceOk

`func (o *WritableDeviceWithConfigContext) GetFaceOk() (*string, bool)`

GetFaceOk returns a tuple with the Face field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFace

`func (o *WritableDeviceWithConfigContext) SetFace(v string)`

SetFace sets Face field to given value.

### HasFace

`func (o *WritableDeviceWithConfigContext) HasFace() bool`

HasFace returns a boolean if a field has been set.

### GetParentDevice

`func (o *WritableDeviceWithConfigContext) GetParentDevice() NestedDevice`

GetParentDevice returns the ParentDevice field if non-nil, zero value otherwise.

### GetParentDeviceOk

`func (o *WritableDeviceWithConfigContext) GetParentDeviceOk() (*NestedDevice, bool)`

GetParentDeviceOk returns a tuple with the ParentDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentDevice

`func (o *WritableDeviceWithConfigContext) SetParentDevice(v NestedDevice)`

SetParentDevice sets ParentDevice field to given value.

### HasParentDevice

`func (o *WritableDeviceWithConfigContext) HasParentDevice() bool`

HasParentDevice returns a boolean if a field has been set.

### GetStatus

`func (o *WritableDeviceWithConfigContext) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WritableDeviceWithConfigContext) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WritableDeviceWithConfigContext) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WritableDeviceWithConfigContext) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPrimaryIp

`func (o *WritableDeviceWithConfigContext) GetPrimaryIp() string`

GetPrimaryIp returns the PrimaryIp field if non-nil, zero value otherwise.

### GetPrimaryIpOk

`func (o *WritableDeviceWithConfigContext) GetPrimaryIpOk() (*string, bool)`

GetPrimaryIpOk returns a tuple with the PrimaryIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIp

`func (o *WritableDeviceWithConfigContext) SetPrimaryIp(v string)`

SetPrimaryIp sets PrimaryIp field to given value.

### HasPrimaryIp

`func (o *WritableDeviceWithConfigContext) HasPrimaryIp() bool`

HasPrimaryIp returns a boolean if a field has been set.

### GetPrimaryIp4

`func (o *WritableDeviceWithConfigContext) GetPrimaryIp4() int32`

GetPrimaryIp4 returns the PrimaryIp4 field if non-nil, zero value otherwise.

### GetPrimaryIp4Ok

`func (o *WritableDeviceWithConfigContext) GetPrimaryIp4Ok() (*int32, bool)`

GetPrimaryIp4Ok returns a tuple with the PrimaryIp4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIp4

`func (o *WritableDeviceWithConfigContext) SetPrimaryIp4(v int32)`

SetPrimaryIp4 sets PrimaryIp4 field to given value.

### HasPrimaryIp4

`func (o *WritableDeviceWithConfigContext) HasPrimaryIp4() bool`

HasPrimaryIp4 returns a boolean if a field has been set.

### SetPrimaryIp4Nil

`func (o *WritableDeviceWithConfigContext) SetPrimaryIp4Nil(b bool)`

 SetPrimaryIp4Nil sets the value for PrimaryIp4 to be an explicit nil

### UnsetPrimaryIp4
`func (o *WritableDeviceWithConfigContext) UnsetPrimaryIp4()`

UnsetPrimaryIp4 ensures that no value is present for PrimaryIp4, not even an explicit nil
### GetPrimaryIp6

`func (o *WritableDeviceWithConfigContext) GetPrimaryIp6() int32`

GetPrimaryIp6 returns the PrimaryIp6 field if non-nil, zero value otherwise.

### GetPrimaryIp6Ok

`func (o *WritableDeviceWithConfigContext) GetPrimaryIp6Ok() (*int32, bool)`

GetPrimaryIp6Ok returns a tuple with the PrimaryIp6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIp6

`func (o *WritableDeviceWithConfigContext) SetPrimaryIp6(v int32)`

SetPrimaryIp6 sets PrimaryIp6 field to given value.

### HasPrimaryIp6

`func (o *WritableDeviceWithConfigContext) HasPrimaryIp6() bool`

HasPrimaryIp6 returns a boolean if a field has been set.

### SetPrimaryIp6Nil

`func (o *WritableDeviceWithConfigContext) SetPrimaryIp6Nil(b bool)`

 SetPrimaryIp6Nil sets the value for PrimaryIp6 to be an explicit nil

### UnsetPrimaryIp6
`func (o *WritableDeviceWithConfigContext) UnsetPrimaryIp6()`

UnsetPrimaryIp6 ensures that no value is present for PrimaryIp6, not even an explicit nil
### GetCluster

`func (o *WritableDeviceWithConfigContext) GetCluster() int32`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *WritableDeviceWithConfigContext) GetClusterOk() (*int32, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *WritableDeviceWithConfigContext) SetCluster(v int32)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *WritableDeviceWithConfigContext) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### SetClusterNil

`func (o *WritableDeviceWithConfigContext) SetClusterNil(b bool)`

 SetClusterNil sets the value for Cluster to be an explicit nil

### UnsetCluster
`func (o *WritableDeviceWithConfigContext) UnsetCluster()`

UnsetCluster ensures that no value is present for Cluster, not even an explicit nil
### GetVirtualChassis

`func (o *WritableDeviceWithConfigContext) GetVirtualChassis() int32`

GetVirtualChassis returns the VirtualChassis field if non-nil, zero value otherwise.

### GetVirtualChassisOk

`func (o *WritableDeviceWithConfigContext) GetVirtualChassisOk() (*int32, bool)`

GetVirtualChassisOk returns a tuple with the VirtualChassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualChassis

`func (o *WritableDeviceWithConfigContext) SetVirtualChassis(v int32)`

SetVirtualChassis sets VirtualChassis field to given value.

### HasVirtualChassis

`func (o *WritableDeviceWithConfigContext) HasVirtualChassis() bool`

HasVirtualChassis returns a boolean if a field has been set.

### SetVirtualChassisNil

`func (o *WritableDeviceWithConfigContext) SetVirtualChassisNil(b bool)`

 SetVirtualChassisNil sets the value for VirtualChassis to be an explicit nil

### UnsetVirtualChassis
`func (o *WritableDeviceWithConfigContext) UnsetVirtualChassis()`

UnsetVirtualChassis ensures that no value is present for VirtualChassis, not even an explicit nil
### GetVcPosition

`func (o *WritableDeviceWithConfigContext) GetVcPosition() int32`

GetVcPosition returns the VcPosition field if non-nil, zero value otherwise.

### GetVcPositionOk

`func (o *WritableDeviceWithConfigContext) GetVcPositionOk() (*int32, bool)`

GetVcPositionOk returns a tuple with the VcPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcPosition

`func (o *WritableDeviceWithConfigContext) SetVcPosition(v int32)`

SetVcPosition sets VcPosition field to given value.

### HasVcPosition

`func (o *WritableDeviceWithConfigContext) HasVcPosition() bool`

HasVcPosition returns a boolean if a field has been set.

### SetVcPositionNil

`func (o *WritableDeviceWithConfigContext) SetVcPositionNil(b bool)`

 SetVcPositionNil sets the value for VcPosition to be an explicit nil

### UnsetVcPosition
`func (o *WritableDeviceWithConfigContext) UnsetVcPosition()`

UnsetVcPosition ensures that no value is present for VcPosition, not even an explicit nil
### GetVcPriority

`func (o *WritableDeviceWithConfigContext) GetVcPriority() int32`

GetVcPriority returns the VcPriority field if non-nil, zero value otherwise.

### GetVcPriorityOk

`func (o *WritableDeviceWithConfigContext) GetVcPriorityOk() (*int32, bool)`

GetVcPriorityOk returns a tuple with the VcPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcPriority

`func (o *WritableDeviceWithConfigContext) SetVcPriority(v int32)`

SetVcPriority sets VcPriority field to given value.

### HasVcPriority

`func (o *WritableDeviceWithConfigContext) HasVcPriority() bool`

HasVcPriority returns a boolean if a field has been set.

### SetVcPriorityNil

`func (o *WritableDeviceWithConfigContext) SetVcPriorityNil(b bool)`

 SetVcPriorityNil sets the value for VcPriority to be an explicit nil

### UnsetVcPriority
`func (o *WritableDeviceWithConfigContext) UnsetVcPriority()`

UnsetVcPriority ensures that no value is present for VcPriority, not even an explicit nil
### GetComments

`func (o *WritableDeviceWithConfigContext) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *WritableDeviceWithConfigContext) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *WritableDeviceWithConfigContext) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *WritableDeviceWithConfigContext) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetLocalContextData

`func (o *WritableDeviceWithConfigContext) GetLocalContextData() string`

GetLocalContextData returns the LocalContextData field if non-nil, zero value otherwise.

### GetLocalContextDataOk

`func (o *WritableDeviceWithConfigContext) GetLocalContextDataOk() (*string, bool)`

GetLocalContextDataOk returns a tuple with the LocalContextData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalContextData

`func (o *WritableDeviceWithConfigContext) SetLocalContextData(v string)`

SetLocalContextData sets LocalContextData field to given value.

### HasLocalContextData

`func (o *WritableDeviceWithConfigContext) HasLocalContextData() bool`

HasLocalContextData returns a boolean if a field has been set.

### SetLocalContextDataNil

`func (o *WritableDeviceWithConfigContext) SetLocalContextDataNil(b bool)`

 SetLocalContextDataNil sets the value for LocalContextData to be an explicit nil

### UnsetLocalContextData
`func (o *WritableDeviceWithConfigContext) UnsetLocalContextData()`

UnsetLocalContextData ensures that no value is present for LocalContextData, not even an explicit nil
### GetTags

`func (o *WritableDeviceWithConfigContext) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WritableDeviceWithConfigContext) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WritableDeviceWithConfigContext) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WritableDeviceWithConfigContext) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *WritableDeviceWithConfigContext) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *WritableDeviceWithConfigContext) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *WritableDeviceWithConfigContext) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *WritableDeviceWithConfigContext) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetConfigContext

`func (o *WritableDeviceWithConfigContext) GetConfigContext() map[string]string`

GetConfigContext returns the ConfigContext field if non-nil, zero value otherwise.

### GetConfigContextOk

`func (o *WritableDeviceWithConfigContext) GetConfigContextOk() (*map[string]string, bool)`

GetConfigContextOk returns a tuple with the ConfigContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigContext

`func (o *WritableDeviceWithConfigContext) SetConfigContext(v map[string]string)`

SetConfigContext sets ConfigContext field to given value.

### HasConfigContext

`func (o *WritableDeviceWithConfigContext) HasConfigContext() bool`

HasConfigContext returns a boolean if a field has been set.

### GetCreated

`func (o *WritableDeviceWithConfigContext) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *WritableDeviceWithConfigContext) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *WritableDeviceWithConfigContext) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *WritableDeviceWithConfigContext) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *WritableDeviceWithConfigContext) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *WritableDeviceWithConfigContext) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *WritableDeviceWithConfigContext) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *WritableDeviceWithConfigContext) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


