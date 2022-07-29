# Device

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
**Created** | Pointer to **string** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewDevice

`func NewDevice(deviceType NestedDeviceType, deviceRole NestedDeviceRole, site NestedSite, ) *Device`

NewDevice instantiates a new Device object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceWithDefaults

`func NewDeviceWithDefaults() *Device`

NewDeviceWithDefaults instantiates a new Device object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Device) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Device) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Device) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Device) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *Device) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Device) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Device) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Device) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDisplay

`func (o *Device) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *Device) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *Device) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *Device) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetName

`func (o *Device) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Device) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Device) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Device) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Device) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Device) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDisplayName

`func (o *Device) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Device) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Device) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Device) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDeviceType

`func (o *Device) GetDeviceType() NestedDeviceType`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *Device) GetDeviceTypeOk() (*NestedDeviceType, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *Device) SetDeviceType(v NestedDeviceType)`

SetDeviceType sets DeviceType field to given value.


### GetDeviceRole

`func (o *Device) GetDeviceRole() NestedDeviceRole`

GetDeviceRole returns the DeviceRole field if non-nil, zero value otherwise.

### GetDeviceRoleOk

`func (o *Device) GetDeviceRoleOk() (*NestedDeviceRole, bool)`

GetDeviceRoleOk returns a tuple with the DeviceRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceRole

`func (o *Device) SetDeviceRole(v NestedDeviceRole)`

SetDeviceRole sets DeviceRole field to given value.


### GetTenant

`func (o *Device) GetTenant() NestedTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *Device) GetTenantOk() (*NestedTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *Device) SetTenant(v NestedTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *Device) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *Device) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *Device) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetPlatform

`func (o *Device) GetPlatform() NestedPlatform`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *Device) GetPlatformOk() (*NestedPlatform, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *Device) SetPlatform(v NestedPlatform)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *Device) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatformNil

`func (o *Device) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *Device) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetSerial

`func (o *Device) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *Device) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *Device) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *Device) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetAssetTag

`func (o *Device) GetAssetTag() string`

GetAssetTag returns the AssetTag field if non-nil, zero value otherwise.

### GetAssetTagOk

`func (o *Device) GetAssetTagOk() (*string, bool)`

GetAssetTagOk returns a tuple with the AssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTag

`func (o *Device) SetAssetTag(v string)`

SetAssetTag sets AssetTag field to given value.

### HasAssetTag

`func (o *Device) HasAssetTag() bool`

HasAssetTag returns a boolean if a field has been set.

### SetAssetTagNil

`func (o *Device) SetAssetTagNil(b bool)`

 SetAssetTagNil sets the value for AssetTag to be an explicit nil

### UnsetAssetTag
`func (o *Device) UnsetAssetTag()`

UnsetAssetTag ensures that no value is present for AssetTag, not even an explicit nil
### GetSite

`func (o *Device) GetSite() NestedSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *Device) GetSiteOk() (*NestedSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *Device) SetSite(v NestedSite)`

SetSite sets Site field to given value.


### GetLocation

`func (o *Device) GetLocation() NestedLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Device) GetLocationOk() (*NestedLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Device) SetLocation(v NestedLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Device) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *Device) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *Device) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetRack

`func (o *Device) GetRack() NestedRack`

GetRack returns the Rack field if non-nil, zero value otherwise.

### GetRackOk

`func (o *Device) GetRackOk() (*NestedRack, bool)`

GetRackOk returns a tuple with the Rack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRack

`func (o *Device) SetRack(v NestedRack)`

SetRack sets Rack field to given value.

### HasRack

`func (o *Device) HasRack() bool`

HasRack returns a boolean if a field has been set.

### SetRackNil

`func (o *Device) SetRackNil(b bool)`

 SetRackNil sets the value for Rack to be an explicit nil

### UnsetRack
`func (o *Device) UnsetRack()`

UnsetRack ensures that no value is present for Rack, not even an explicit nil
### GetPosition

`func (o *Device) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *Device) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *Device) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *Device) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### SetPositionNil

`func (o *Device) SetPositionNil(b bool)`

 SetPositionNil sets the value for Position to be an explicit nil

### UnsetPosition
`func (o *Device) UnsetPosition()`

UnsetPosition ensures that no value is present for Position, not even an explicit nil
### GetFace

`func (o *Device) GetFace() Face`

GetFace returns the Face field if non-nil, zero value otherwise.

### GetFaceOk

`func (o *Device) GetFaceOk() (*Face, bool)`

GetFaceOk returns a tuple with the Face field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFace

`func (o *Device) SetFace(v Face)`

SetFace sets Face field to given value.

### HasFace

`func (o *Device) HasFace() bool`

HasFace returns a boolean if a field has been set.

### GetParentDevice

`func (o *Device) GetParentDevice() NestedDevice`

GetParentDevice returns the ParentDevice field if non-nil, zero value otherwise.

### GetParentDeviceOk

`func (o *Device) GetParentDeviceOk() (*NestedDevice, bool)`

GetParentDeviceOk returns a tuple with the ParentDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentDevice

`func (o *Device) SetParentDevice(v NestedDevice)`

SetParentDevice sets ParentDevice field to given value.

### HasParentDevice

`func (o *Device) HasParentDevice() bool`

HasParentDevice returns a boolean if a field has been set.

### GetStatus

`func (o *Device) GetStatus() Status2`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Device) GetStatusOk() (*Status2, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Device) SetStatus(v Status2)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Device) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPrimaryIp

`func (o *Device) GetPrimaryIp() NestedIPAddress`

GetPrimaryIp returns the PrimaryIp field if non-nil, zero value otherwise.

### GetPrimaryIpOk

`func (o *Device) GetPrimaryIpOk() (*NestedIPAddress, bool)`

GetPrimaryIpOk returns a tuple with the PrimaryIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIp

`func (o *Device) SetPrimaryIp(v NestedIPAddress)`

SetPrimaryIp sets PrimaryIp field to given value.

### HasPrimaryIp

`func (o *Device) HasPrimaryIp() bool`

HasPrimaryIp returns a boolean if a field has been set.

### GetPrimaryIp4

`func (o *Device) GetPrimaryIp4() NestedIPAddress`

GetPrimaryIp4 returns the PrimaryIp4 field if non-nil, zero value otherwise.

### GetPrimaryIp4Ok

`func (o *Device) GetPrimaryIp4Ok() (*NestedIPAddress, bool)`

GetPrimaryIp4Ok returns a tuple with the PrimaryIp4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIp4

`func (o *Device) SetPrimaryIp4(v NestedIPAddress)`

SetPrimaryIp4 sets PrimaryIp4 field to given value.

### HasPrimaryIp4

`func (o *Device) HasPrimaryIp4() bool`

HasPrimaryIp4 returns a boolean if a field has been set.

### GetPrimaryIp6

`func (o *Device) GetPrimaryIp6() NestedIPAddress`

GetPrimaryIp6 returns the PrimaryIp6 field if non-nil, zero value otherwise.

### GetPrimaryIp6Ok

`func (o *Device) GetPrimaryIp6Ok() (*NestedIPAddress, bool)`

GetPrimaryIp6Ok returns a tuple with the PrimaryIp6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIp6

`func (o *Device) SetPrimaryIp6(v NestedIPAddress)`

SetPrimaryIp6 sets PrimaryIp6 field to given value.

### HasPrimaryIp6

`func (o *Device) HasPrimaryIp6() bool`

HasPrimaryIp6 returns a boolean if a field has been set.

### GetCluster

`func (o *Device) GetCluster() NestedCluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *Device) GetClusterOk() (*NestedCluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *Device) SetCluster(v NestedCluster)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *Device) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### SetClusterNil

`func (o *Device) SetClusterNil(b bool)`

 SetClusterNil sets the value for Cluster to be an explicit nil

### UnsetCluster
`func (o *Device) UnsetCluster()`

UnsetCluster ensures that no value is present for Cluster, not even an explicit nil
### GetVirtualChassis

`func (o *Device) GetVirtualChassis() NestedVirtualChassis`

GetVirtualChassis returns the VirtualChassis field if non-nil, zero value otherwise.

### GetVirtualChassisOk

`func (o *Device) GetVirtualChassisOk() (*NestedVirtualChassis, bool)`

GetVirtualChassisOk returns a tuple with the VirtualChassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualChassis

`func (o *Device) SetVirtualChassis(v NestedVirtualChassis)`

SetVirtualChassis sets VirtualChassis field to given value.

### HasVirtualChassis

`func (o *Device) HasVirtualChassis() bool`

HasVirtualChassis returns a boolean if a field has been set.

### SetVirtualChassisNil

`func (o *Device) SetVirtualChassisNil(b bool)`

 SetVirtualChassisNil sets the value for VirtualChassis to be an explicit nil

### UnsetVirtualChassis
`func (o *Device) UnsetVirtualChassis()`

UnsetVirtualChassis ensures that no value is present for VirtualChassis, not even an explicit nil
### GetVcPosition

`func (o *Device) GetVcPosition() int32`

GetVcPosition returns the VcPosition field if non-nil, zero value otherwise.

### GetVcPositionOk

`func (o *Device) GetVcPositionOk() (*int32, bool)`

GetVcPositionOk returns a tuple with the VcPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcPosition

`func (o *Device) SetVcPosition(v int32)`

SetVcPosition sets VcPosition field to given value.

### HasVcPosition

`func (o *Device) HasVcPosition() bool`

HasVcPosition returns a boolean if a field has been set.

### SetVcPositionNil

`func (o *Device) SetVcPositionNil(b bool)`

 SetVcPositionNil sets the value for VcPosition to be an explicit nil

### UnsetVcPosition
`func (o *Device) UnsetVcPosition()`

UnsetVcPosition ensures that no value is present for VcPosition, not even an explicit nil
### GetVcPriority

`func (o *Device) GetVcPriority() int32`

GetVcPriority returns the VcPriority field if non-nil, zero value otherwise.

### GetVcPriorityOk

`func (o *Device) GetVcPriorityOk() (*int32, bool)`

GetVcPriorityOk returns a tuple with the VcPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcPriority

`func (o *Device) SetVcPriority(v int32)`

SetVcPriority sets VcPriority field to given value.

### HasVcPriority

`func (o *Device) HasVcPriority() bool`

HasVcPriority returns a boolean if a field has been set.

### SetVcPriorityNil

`func (o *Device) SetVcPriorityNil(b bool)`

 SetVcPriorityNil sets the value for VcPriority to be an explicit nil

### UnsetVcPriority
`func (o *Device) UnsetVcPriority()`

UnsetVcPriority ensures that no value is present for VcPriority, not even an explicit nil
### GetComments

`func (o *Device) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *Device) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *Device) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *Device) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetLocalContextData

`func (o *Device) GetLocalContextData() string`

GetLocalContextData returns the LocalContextData field if non-nil, zero value otherwise.

### GetLocalContextDataOk

`func (o *Device) GetLocalContextDataOk() (*string, bool)`

GetLocalContextDataOk returns a tuple with the LocalContextData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalContextData

`func (o *Device) SetLocalContextData(v string)`

SetLocalContextData sets LocalContextData field to given value.

### HasLocalContextData

`func (o *Device) HasLocalContextData() bool`

HasLocalContextData returns a boolean if a field has been set.

### SetLocalContextDataNil

`func (o *Device) SetLocalContextDataNil(b bool)`

 SetLocalContextDataNil sets the value for LocalContextData to be an explicit nil

### UnsetLocalContextData
`func (o *Device) UnsetLocalContextData()`

UnsetLocalContextData ensures that no value is present for LocalContextData, not even an explicit nil
### GetTags

`func (o *Device) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Device) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Device) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Device) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *Device) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Device) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Device) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Device) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *Device) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Device) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Device) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Device) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *Device) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Device) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Device) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *Device) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


