# DeviceWithConfigContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Display** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **NullableString** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] [readonly] 
**DeviceType** | [**NestedDeviceType**](NestedDeviceType.md) |  | 
**DeviceRole** | [**NestedDeviceRole**](NestedDeviceRole.md) |  | 
**Tenant** | Pointer to [**NullableNestedTenant**](NestedTenant.md) |  | [optional] 
**Platform** | Pointer to [**NullableNestedPlatform**](NestedPlatform.md) |  | [optional] 
**Serial** | Pointer to **string** |  | [optional] 
**AssetTag** | Pointer to **NullableString** | A unique tag used to identify this device | [optional] 
**Site** | [**NestedSite**](NestedSite.md) |  | 
**Location** | Pointer to [**NullableNestedLocation**](NestedLocation.md) |  | [optional] 
**Rack** | Pointer to [**NullableNestedRack**](NestedRack.md) |  | [optional] 
**Position** | Pointer to **NullableInt32** | The lowest-numbered unit occupied by the device | [optional] 
**Face** | Pointer to [**Face**](Face.md) |  | [optional] 
**ParentDevice** | Pointer to [**NestedDevice**](NestedDevice.md) |  | [optional] 
**Status** | Pointer to [**Status2**](Status2.md) |  | [optional] 
**PrimaryIp** | Pointer to [**NestedIPAddress**](NestedIPAddress.md) |  | [optional] 
**PrimaryIp4** | Pointer to [**NestedIPAddress**](NestedIPAddress.md) |  | [optional] 
**PrimaryIp6** | Pointer to [**NestedIPAddress**](NestedIPAddress.md) |  | [optional] 
**Cluster** | Pointer to [**NullableNestedCluster**](NestedCluster.md) |  | [optional] 
**VirtualChassis** | Pointer to [**NullableNestedVirtualChassis**](NestedVirtualChassis.md) |  | [optional] 
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

### NewDeviceWithConfigContext

`func NewDeviceWithConfigContext(deviceType NestedDeviceType, deviceRole NestedDeviceRole, site NestedSite, ) *DeviceWithConfigContext`

NewDeviceWithConfigContext instantiates a new DeviceWithConfigContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceWithConfigContextWithDefaults

`func NewDeviceWithConfigContextWithDefaults() *DeviceWithConfigContext`

NewDeviceWithConfigContextWithDefaults instantiates a new DeviceWithConfigContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeviceWithConfigContext) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceWithConfigContext) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceWithConfigContext) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *DeviceWithConfigContext) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *DeviceWithConfigContext) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DeviceWithConfigContext) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DeviceWithConfigContext) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *DeviceWithConfigContext) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDisplay

`func (o *DeviceWithConfigContext) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *DeviceWithConfigContext) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *DeviceWithConfigContext) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *DeviceWithConfigContext) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetName

`func (o *DeviceWithConfigContext) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceWithConfigContext) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceWithConfigContext) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeviceWithConfigContext) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *DeviceWithConfigContext) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DeviceWithConfigContext) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDisplayName

`func (o *DeviceWithConfigContext) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DeviceWithConfigContext) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *DeviceWithConfigContext) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *DeviceWithConfigContext) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDeviceType

`func (o *DeviceWithConfigContext) GetDeviceType() NestedDeviceType`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *DeviceWithConfigContext) GetDeviceTypeOk() (*NestedDeviceType, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *DeviceWithConfigContext) SetDeviceType(v NestedDeviceType)`

SetDeviceType sets DeviceType field to given value.


### GetDeviceRole

`func (o *DeviceWithConfigContext) GetDeviceRole() NestedDeviceRole`

GetDeviceRole returns the DeviceRole field if non-nil, zero value otherwise.

### GetDeviceRoleOk

`func (o *DeviceWithConfigContext) GetDeviceRoleOk() (*NestedDeviceRole, bool)`

GetDeviceRoleOk returns a tuple with the DeviceRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceRole

`func (o *DeviceWithConfigContext) SetDeviceRole(v NestedDeviceRole)`

SetDeviceRole sets DeviceRole field to given value.


### GetTenant

`func (o *DeviceWithConfigContext) GetTenant() NestedTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *DeviceWithConfigContext) GetTenantOk() (*NestedTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *DeviceWithConfigContext) SetTenant(v NestedTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *DeviceWithConfigContext) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *DeviceWithConfigContext) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *DeviceWithConfigContext) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetPlatform

`func (o *DeviceWithConfigContext) GetPlatform() NestedPlatform`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *DeviceWithConfigContext) GetPlatformOk() (*NestedPlatform, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *DeviceWithConfigContext) SetPlatform(v NestedPlatform)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *DeviceWithConfigContext) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatformNil

`func (o *DeviceWithConfigContext) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *DeviceWithConfigContext) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetSerial

`func (o *DeviceWithConfigContext) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *DeviceWithConfigContext) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *DeviceWithConfigContext) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *DeviceWithConfigContext) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetAssetTag

`func (o *DeviceWithConfigContext) GetAssetTag() string`

GetAssetTag returns the AssetTag field if non-nil, zero value otherwise.

### GetAssetTagOk

`func (o *DeviceWithConfigContext) GetAssetTagOk() (*string, bool)`

GetAssetTagOk returns a tuple with the AssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTag

`func (o *DeviceWithConfigContext) SetAssetTag(v string)`

SetAssetTag sets AssetTag field to given value.

### HasAssetTag

`func (o *DeviceWithConfigContext) HasAssetTag() bool`

HasAssetTag returns a boolean if a field has been set.

### SetAssetTagNil

`func (o *DeviceWithConfigContext) SetAssetTagNil(b bool)`

 SetAssetTagNil sets the value for AssetTag to be an explicit nil

### UnsetAssetTag
`func (o *DeviceWithConfigContext) UnsetAssetTag()`

UnsetAssetTag ensures that no value is present for AssetTag, not even an explicit nil
### GetSite

`func (o *DeviceWithConfigContext) GetSite() NestedSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *DeviceWithConfigContext) GetSiteOk() (*NestedSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *DeviceWithConfigContext) SetSite(v NestedSite)`

SetSite sets Site field to given value.


### GetLocation

`func (o *DeviceWithConfigContext) GetLocation() NestedLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *DeviceWithConfigContext) GetLocationOk() (*NestedLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *DeviceWithConfigContext) SetLocation(v NestedLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *DeviceWithConfigContext) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *DeviceWithConfigContext) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *DeviceWithConfigContext) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetRack

`func (o *DeviceWithConfigContext) GetRack() NestedRack`

GetRack returns the Rack field if non-nil, zero value otherwise.

### GetRackOk

`func (o *DeviceWithConfigContext) GetRackOk() (*NestedRack, bool)`

GetRackOk returns a tuple with the Rack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRack

`func (o *DeviceWithConfigContext) SetRack(v NestedRack)`

SetRack sets Rack field to given value.

### HasRack

`func (o *DeviceWithConfigContext) HasRack() bool`

HasRack returns a boolean if a field has been set.

### SetRackNil

`func (o *DeviceWithConfigContext) SetRackNil(b bool)`

 SetRackNil sets the value for Rack to be an explicit nil

### UnsetRack
`func (o *DeviceWithConfigContext) UnsetRack()`

UnsetRack ensures that no value is present for Rack, not even an explicit nil
### GetPosition

`func (o *DeviceWithConfigContext) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *DeviceWithConfigContext) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *DeviceWithConfigContext) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *DeviceWithConfigContext) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### SetPositionNil

`func (o *DeviceWithConfigContext) SetPositionNil(b bool)`

 SetPositionNil sets the value for Position to be an explicit nil

### UnsetPosition
`func (o *DeviceWithConfigContext) UnsetPosition()`

UnsetPosition ensures that no value is present for Position, not even an explicit nil
### GetFace

`func (o *DeviceWithConfigContext) GetFace() Face`

GetFace returns the Face field if non-nil, zero value otherwise.

### GetFaceOk

`func (o *DeviceWithConfigContext) GetFaceOk() (*Face, bool)`

GetFaceOk returns a tuple with the Face field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFace

`func (o *DeviceWithConfigContext) SetFace(v Face)`

SetFace sets Face field to given value.

### HasFace

`func (o *DeviceWithConfigContext) HasFace() bool`

HasFace returns a boolean if a field has been set.

### GetParentDevice

`func (o *DeviceWithConfigContext) GetParentDevice() NestedDevice`

GetParentDevice returns the ParentDevice field if non-nil, zero value otherwise.

### GetParentDeviceOk

`func (o *DeviceWithConfigContext) GetParentDeviceOk() (*NestedDevice, bool)`

GetParentDeviceOk returns a tuple with the ParentDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentDevice

`func (o *DeviceWithConfigContext) SetParentDevice(v NestedDevice)`

SetParentDevice sets ParentDevice field to given value.

### HasParentDevice

`func (o *DeviceWithConfigContext) HasParentDevice() bool`

HasParentDevice returns a boolean if a field has been set.

### GetStatus

`func (o *DeviceWithConfigContext) GetStatus() Status2`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeviceWithConfigContext) GetStatusOk() (*Status2, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeviceWithConfigContext) SetStatus(v Status2)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeviceWithConfigContext) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPrimaryIp

`func (o *DeviceWithConfigContext) GetPrimaryIp() NestedIPAddress`

GetPrimaryIp returns the PrimaryIp field if non-nil, zero value otherwise.

### GetPrimaryIpOk

`func (o *DeviceWithConfigContext) GetPrimaryIpOk() (*NestedIPAddress, bool)`

GetPrimaryIpOk returns a tuple with the PrimaryIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIp

`func (o *DeviceWithConfigContext) SetPrimaryIp(v NestedIPAddress)`

SetPrimaryIp sets PrimaryIp field to given value.

### HasPrimaryIp

`func (o *DeviceWithConfigContext) HasPrimaryIp() bool`

HasPrimaryIp returns a boolean if a field has been set.

### GetPrimaryIp4

`func (o *DeviceWithConfigContext) GetPrimaryIp4() NestedIPAddress`

GetPrimaryIp4 returns the PrimaryIp4 field if non-nil, zero value otherwise.

### GetPrimaryIp4Ok

`func (o *DeviceWithConfigContext) GetPrimaryIp4Ok() (*NestedIPAddress, bool)`

GetPrimaryIp4Ok returns a tuple with the PrimaryIp4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIp4

`func (o *DeviceWithConfigContext) SetPrimaryIp4(v NestedIPAddress)`

SetPrimaryIp4 sets PrimaryIp4 field to given value.

### HasPrimaryIp4

`func (o *DeviceWithConfigContext) HasPrimaryIp4() bool`

HasPrimaryIp4 returns a boolean if a field has been set.

### GetPrimaryIp6

`func (o *DeviceWithConfigContext) GetPrimaryIp6() NestedIPAddress`

GetPrimaryIp6 returns the PrimaryIp6 field if non-nil, zero value otherwise.

### GetPrimaryIp6Ok

`func (o *DeviceWithConfigContext) GetPrimaryIp6Ok() (*NestedIPAddress, bool)`

GetPrimaryIp6Ok returns a tuple with the PrimaryIp6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIp6

`func (o *DeviceWithConfigContext) SetPrimaryIp6(v NestedIPAddress)`

SetPrimaryIp6 sets PrimaryIp6 field to given value.

### HasPrimaryIp6

`func (o *DeviceWithConfigContext) HasPrimaryIp6() bool`

HasPrimaryIp6 returns a boolean if a field has been set.

### GetCluster

`func (o *DeviceWithConfigContext) GetCluster() NestedCluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DeviceWithConfigContext) GetClusterOk() (*NestedCluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DeviceWithConfigContext) SetCluster(v NestedCluster)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DeviceWithConfigContext) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### SetClusterNil

`func (o *DeviceWithConfigContext) SetClusterNil(b bool)`

 SetClusterNil sets the value for Cluster to be an explicit nil

### UnsetCluster
`func (o *DeviceWithConfigContext) UnsetCluster()`

UnsetCluster ensures that no value is present for Cluster, not even an explicit nil
### GetVirtualChassis

`func (o *DeviceWithConfigContext) GetVirtualChassis() NestedVirtualChassis`

GetVirtualChassis returns the VirtualChassis field if non-nil, zero value otherwise.

### GetVirtualChassisOk

`func (o *DeviceWithConfigContext) GetVirtualChassisOk() (*NestedVirtualChassis, bool)`

GetVirtualChassisOk returns a tuple with the VirtualChassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualChassis

`func (o *DeviceWithConfigContext) SetVirtualChassis(v NestedVirtualChassis)`

SetVirtualChassis sets VirtualChassis field to given value.

### HasVirtualChassis

`func (o *DeviceWithConfigContext) HasVirtualChassis() bool`

HasVirtualChassis returns a boolean if a field has been set.

### SetVirtualChassisNil

`func (o *DeviceWithConfigContext) SetVirtualChassisNil(b bool)`

 SetVirtualChassisNil sets the value for VirtualChassis to be an explicit nil

### UnsetVirtualChassis
`func (o *DeviceWithConfigContext) UnsetVirtualChassis()`

UnsetVirtualChassis ensures that no value is present for VirtualChassis, not even an explicit nil
### GetVcPosition

`func (o *DeviceWithConfigContext) GetVcPosition() int32`

GetVcPosition returns the VcPosition field if non-nil, zero value otherwise.

### GetVcPositionOk

`func (o *DeviceWithConfigContext) GetVcPositionOk() (*int32, bool)`

GetVcPositionOk returns a tuple with the VcPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcPosition

`func (o *DeviceWithConfigContext) SetVcPosition(v int32)`

SetVcPosition sets VcPosition field to given value.

### HasVcPosition

`func (o *DeviceWithConfigContext) HasVcPosition() bool`

HasVcPosition returns a boolean if a field has been set.

### SetVcPositionNil

`func (o *DeviceWithConfigContext) SetVcPositionNil(b bool)`

 SetVcPositionNil sets the value for VcPosition to be an explicit nil

### UnsetVcPosition
`func (o *DeviceWithConfigContext) UnsetVcPosition()`

UnsetVcPosition ensures that no value is present for VcPosition, not even an explicit nil
### GetVcPriority

`func (o *DeviceWithConfigContext) GetVcPriority() int32`

GetVcPriority returns the VcPriority field if non-nil, zero value otherwise.

### GetVcPriorityOk

`func (o *DeviceWithConfigContext) GetVcPriorityOk() (*int32, bool)`

GetVcPriorityOk returns a tuple with the VcPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcPriority

`func (o *DeviceWithConfigContext) SetVcPriority(v int32)`

SetVcPriority sets VcPriority field to given value.

### HasVcPriority

`func (o *DeviceWithConfigContext) HasVcPriority() bool`

HasVcPriority returns a boolean if a field has been set.

### SetVcPriorityNil

`func (o *DeviceWithConfigContext) SetVcPriorityNil(b bool)`

 SetVcPriorityNil sets the value for VcPriority to be an explicit nil

### UnsetVcPriority
`func (o *DeviceWithConfigContext) UnsetVcPriority()`

UnsetVcPriority ensures that no value is present for VcPriority, not even an explicit nil
### GetComments

`func (o *DeviceWithConfigContext) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *DeviceWithConfigContext) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *DeviceWithConfigContext) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *DeviceWithConfigContext) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetLocalContextData

`func (o *DeviceWithConfigContext) GetLocalContextData() string`

GetLocalContextData returns the LocalContextData field if non-nil, zero value otherwise.

### GetLocalContextDataOk

`func (o *DeviceWithConfigContext) GetLocalContextDataOk() (*string, bool)`

GetLocalContextDataOk returns a tuple with the LocalContextData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalContextData

`func (o *DeviceWithConfigContext) SetLocalContextData(v string)`

SetLocalContextData sets LocalContextData field to given value.

### HasLocalContextData

`func (o *DeviceWithConfigContext) HasLocalContextData() bool`

HasLocalContextData returns a boolean if a field has been set.

### SetLocalContextDataNil

`func (o *DeviceWithConfigContext) SetLocalContextDataNil(b bool)`

 SetLocalContextDataNil sets the value for LocalContextData to be an explicit nil

### UnsetLocalContextData
`func (o *DeviceWithConfigContext) UnsetLocalContextData()`

UnsetLocalContextData ensures that no value is present for LocalContextData, not even an explicit nil
### GetTags

`func (o *DeviceWithConfigContext) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DeviceWithConfigContext) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DeviceWithConfigContext) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DeviceWithConfigContext) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *DeviceWithConfigContext) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *DeviceWithConfigContext) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *DeviceWithConfigContext) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *DeviceWithConfigContext) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetConfigContext

`func (o *DeviceWithConfigContext) GetConfigContext() map[string]string`

GetConfigContext returns the ConfigContext field if non-nil, zero value otherwise.

### GetConfigContextOk

`func (o *DeviceWithConfigContext) GetConfigContextOk() (*map[string]string, bool)`

GetConfigContextOk returns a tuple with the ConfigContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigContext

`func (o *DeviceWithConfigContext) SetConfigContext(v map[string]string)`

SetConfigContext sets ConfigContext field to given value.

### HasConfigContext

`func (o *DeviceWithConfigContext) HasConfigContext() bool`

HasConfigContext returns a boolean if a field has been set.

### GetCreated

`func (o *DeviceWithConfigContext) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *DeviceWithConfigContext) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *DeviceWithConfigContext) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *DeviceWithConfigContext) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *DeviceWithConfigContext) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *DeviceWithConfigContext) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *DeviceWithConfigContext) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *DeviceWithConfigContext) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


