/*
NetBox API

API to access NetBox

API version: 2.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// Device struct for Device
type Device struct {
	Id *int32 `json:"id,omitempty"`
	Url *string `json:"url,omitempty"`
	Display *string `json:"display,omitempty"`
	Name NullableString `json:"name,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
	DeviceType NestedDeviceType `json:"device_type"`
	DeviceRole NestedDeviceRole `json:"device_role"`
	Tenant NullableNestedTenant `json:"tenant,omitempty"`
	Platform NullableNestedPlatform `json:"platform,omitempty"`
	Serial *string `json:"serial,omitempty"`
	// A unique tag used to identify this device
	AssetTag NullableString `json:"asset_tag,omitempty"`
	Site NestedSite `json:"site"`
	Location NullableNestedLocation `json:"location,omitempty"`
	Rack NullableNestedRack `json:"rack,omitempty"`
	// The lowest-numbered unit occupied by the device
	Position NullableInt32 `json:"position,omitempty"`
	Face *Face `json:"face,omitempty"`
	ParentDevice *NestedDevice `json:"parent_device,omitempty"`
	Status *Status2 `json:"status,omitempty"`
	PrimaryIp *NestedIPAddress `json:"primary_ip,omitempty"`
	PrimaryIp4 *NestedIPAddress `json:"primary_ip4,omitempty"`
	PrimaryIp6 *NestedIPAddress `json:"primary_ip6,omitempty"`
	Cluster NullableNestedCluster `json:"cluster,omitempty"`
	VirtualChassis NullableNestedVirtualChassis `json:"virtual_chassis,omitempty"`
	VcPosition NullableInt32 `json:"vc_position,omitempty"`
	VcPriority NullableInt32 `json:"vc_priority,omitempty"`
	Comments *string `json:"comments,omitempty"`
	LocalContextData NullableString `json:"local_context_data,omitempty"`
	Tags []NestedTag `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Created *string `json:"created,omitempty"`
	LastUpdated *time.Time `json:"last_updated,omitempty"`
}

// NewDevice instantiates a new Device object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevice(deviceType NestedDeviceType, deviceRole NestedDeviceRole, site NestedSite) *Device {
	this := Device{}
	this.DeviceType = deviceType
	this.DeviceRole = deviceRole
	this.Site = site
	return &this
}

// NewDeviceWithDefaults instantiates a new Device object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceWithDefaults() *Device {
	this := Device{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Device) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Device) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Device) SetId(v int32) {
	o.Id = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *Device) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *Device) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *Device) SetUrl(v string) {
	o.Url = &v
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *Device) GetDisplay() string {
	if o == nil || o.Display == nil {
		var ret string
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetDisplayOk() (*string, bool) {
	if o == nil || o.Display == nil {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *Device) HasDisplay() bool {
	if o != nil && o.Display != nil {
		return true
	}

	return false
}

// SetDisplay gets a reference to the given string and assigns it to the Display field.
func (o *Device) SetDisplay(v string) {
	o.Display = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Device) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Device) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *Device) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *Device) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *Device) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *Device) UnsetName() {
	o.Name.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *Device) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *Device) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *Device) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDeviceType returns the DeviceType field value
func (o *Device) GetDeviceType() NestedDeviceType {
	if o == nil {
		var ret NestedDeviceType
		return ret
	}

	return o.DeviceType
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value
// and a boolean to check if the value has been set.
func (o *Device) GetDeviceTypeOk() (*NestedDeviceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceType, true
}

// SetDeviceType sets field value
func (o *Device) SetDeviceType(v NestedDeviceType) {
	o.DeviceType = v
}

// GetDeviceRole returns the DeviceRole field value
func (o *Device) GetDeviceRole() NestedDeviceRole {
	if o == nil {
		var ret NestedDeviceRole
		return ret
	}

	return o.DeviceRole
}

// GetDeviceRoleOk returns a tuple with the DeviceRole field value
// and a boolean to check if the value has been set.
func (o *Device) GetDeviceRoleOk() (*NestedDeviceRole, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceRole, true
}

// SetDeviceRole sets field value
func (o *Device) SetDeviceRole(v NestedDeviceRole) {
	o.DeviceRole = v
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Device) GetTenant() NestedTenant {
	if o == nil || o.Tenant.Get() == nil {
		var ret NestedTenant
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Device) GetTenantOk() (*NestedTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *Device) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableNestedTenant and assigns it to the Tenant field.
func (o *Device) SetTenant(v NestedTenant) {
	o.Tenant.Set(&v)
}
// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *Device) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *Device) UnsetTenant() {
	o.Tenant.Unset()
}

// GetPlatform returns the Platform field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Device) GetPlatform() NestedPlatform {
	if o == nil || o.Platform.Get() == nil {
		var ret NestedPlatform
		return ret
	}
	return *o.Platform.Get()
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Device) GetPlatformOk() (*NestedPlatform, bool) {
	if o == nil {
		return nil, false
	}
	return o.Platform.Get(), o.Platform.IsSet()
}

// HasPlatform returns a boolean if a field has been set.
func (o *Device) HasPlatform() bool {
	if o != nil && o.Platform.IsSet() {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given NullableNestedPlatform and assigns it to the Platform field.
func (o *Device) SetPlatform(v NestedPlatform) {
	o.Platform.Set(&v)
}
// SetPlatformNil sets the value for Platform to be an explicit nil
func (o *Device) SetPlatformNil() {
	o.Platform.Set(nil)
}

// UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
func (o *Device) UnsetPlatform() {
	o.Platform.Unset()
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *Device) GetSerial() string {
	if o == nil || o.Serial == nil {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetSerialOk() (*string, bool) {
	if o == nil || o.Serial == nil {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *Device) HasSerial() bool {
	if o != nil && o.Serial != nil {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *Device) SetSerial(v string) {
	o.Serial = &v
}

// GetAssetTag returns the AssetTag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Device) GetAssetTag() string {
	if o == nil || o.AssetTag.Get() == nil {
		var ret string
		return ret
	}
	return *o.AssetTag.Get()
}

// GetAssetTagOk returns a tuple with the AssetTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Device) GetAssetTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssetTag.Get(), o.AssetTag.IsSet()
}

// HasAssetTag returns a boolean if a field has been set.
func (o *Device) HasAssetTag() bool {
	if o != nil && o.AssetTag.IsSet() {
		return true
	}

	return false
}

// SetAssetTag gets a reference to the given NullableString and assigns it to the AssetTag field.
func (o *Device) SetAssetTag(v string) {
	o.AssetTag.Set(&v)
}
// SetAssetTagNil sets the value for AssetTag to be an explicit nil
func (o *Device) SetAssetTagNil() {
	o.AssetTag.Set(nil)
}

// UnsetAssetTag ensures that no value is present for AssetTag, not even an explicit nil
func (o *Device) UnsetAssetTag() {
	o.AssetTag.Unset()
}

// GetSite returns the Site field value
func (o *Device) GetSite() NestedSite {
	if o == nil {
		var ret NestedSite
		return ret
	}

	return o.Site
}

// GetSiteOk returns a tuple with the Site field value
// and a boolean to check if the value has been set.
func (o *Device) GetSiteOk() (*NestedSite, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Site, true
}

// SetSite sets field value
func (o *Device) SetSite(v NestedSite) {
	o.Site = v
}

// GetLocation returns the Location field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Device) GetLocation() NestedLocation {
	if o == nil || o.Location.Get() == nil {
		var ret NestedLocation
		return ret
	}
	return *o.Location.Get()
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Device) GetLocationOk() (*NestedLocation, bool) {
	if o == nil {
		return nil, false
	}
	return o.Location.Get(), o.Location.IsSet()
}

// HasLocation returns a boolean if a field has been set.
func (o *Device) HasLocation() bool {
	if o != nil && o.Location.IsSet() {
		return true
	}

	return false
}

// SetLocation gets a reference to the given NullableNestedLocation and assigns it to the Location field.
func (o *Device) SetLocation(v NestedLocation) {
	o.Location.Set(&v)
}
// SetLocationNil sets the value for Location to be an explicit nil
func (o *Device) SetLocationNil() {
	o.Location.Set(nil)
}

// UnsetLocation ensures that no value is present for Location, not even an explicit nil
func (o *Device) UnsetLocation() {
	o.Location.Unset()
}

// GetRack returns the Rack field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Device) GetRack() NestedRack {
	if o == nil || o.Rack.Get() == nil {
		var ret NestedRack
		return ret
	}
	return *o.Rack.Get()
}

// GetRackOk returns a tuple with the Rack field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Device) GetRackOk() (*NestedRack, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rack.Get(), o.Rack.IsSet()
}

// HasRack returns a boolean if a field has been set.
func (o *Device) HasRack() bool {
	if o != nil && o.Rack.IsSet() {
		return true
	}

	return false
}

// SetRack gets a reference to the given NullableNestedRack and assigns it to the Rack field.
func (o *Device) SetRack(v NestedRack) {
	o.Rack.Set(&v)
}
// SetRackNil sets the value for Rack to be an explicit nil
func (o *Device) SetRackNil() {
	o.Rack.Set(nil)
}

// UnsetRack ensures that no value is present for Rack, not even an explicit nil
func (o *Device) UnsetRack() {
	o.Rack.Unset()
}

// GetPosition returns the Position field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Device) GetPosition() int32 {
	if o == nil || o.Position.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Position.Get()
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Device) GetPositionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Position.Get(), o.Position.IsSet()
}

// HasPosition returns a boolean if a field has been set.
func (o *Device) HasPosition() bool {
	if o != nil && o.Position.IsSet() {
		return true
	}

	return false
}

// SetPosition gets a reference to the given NullableInt32 and assigns it to the Position field.
func (o *Device) SetPosition(v int32) {
	o.Position.Set(&v)
}
// SetPositionNil sets the value for Position to be an explicit nil
func (o *Device) SetPositionNil() {
	o.Position.Set(nil)
}

// UnsetPosition ensures that no value is present for Position, not even an explicit nil
func (o *Device) UnsetPosition() {
	o.Position.Unset()
}

// GetFace returns the Face field value if set, zero value otherwise.
func (o *Device) GetFace() Face {
	if o == nil || o.Face == nil {
		var ret Face
		return ret
	}
	return *o.Face
}

// GetFaceOk returns a tuple with the Face field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetFaceOk() (*Face, bool) {
	if o == nil || o.Face == nil {
		return nil, false
	}
	return o.Face, true
}

// HasFace returns a boolean if a field has been set.
func (o *Device) HasFace() bool {
	if o != nil && o.Face != nil {
		return true
	}

	return false
}

// SetFace gets a reference to the given Face and assigns it to the Face field.
func (o *Device) SetFace(v Face) {
	o.Face = &v
}

// GetParentDevice returns the ParentDevice field value if set, zero value otherwise.
func (o *Device) GetParentDevice() NestedDevice {
	if o == nil || o.ParentDevice == nil {
		var ret NestedDevice
		return ret
	}
	return *o.ParentDevice
}

// GetParentDeviceOk returns a tuple with the ParentDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetParentDeviceOk() (*NestedDevice, bool) {
	if o == nil || o.ParentDevice == nil {
		return nil, false
	}
	return o.ParentDevice, true
}

// HasParentDevice returns a boolean if a field has been set.
func (o *Device) HasParentDevice() bool {
	if o != nil && o.ParentDevice != nil {
		return true
	}

	return false
}

// SetParentDevice gets a reference to the given NestedDevice and assigns it to the ParentDevice field.
func (o *Device) SetParentDevice(v NestedDevice) {
	o.ParentDevice = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Device) GetStatus() Status2 {
	if o == nil || o.Status == nil {
		var ret Status2
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetStatusOk() (*Status2, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Device) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given Status2 and assigns it to the Status field.
func (o *Device) SetStatus(v Status2) {
	o.Status = &v
}

// GetPrimaryIp returns the PrimaryIp field value if set, zero value otherwise.
func (o *Device) GetPrimaryIp() NestedIPAddress {
	if o == nil || o.PrimaryIp == nil {
		var ret NestedIPAddress
		return ret
	}
	return *o.PrimaryIp
}

// GetPrimaryIpOk returns a tuple with the PrimaryIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetPrimaryIpOk() (*NestedIPAddress, bool) {
	if o == nil || o.PrimaryIp == nil {
		return nil, false
	}
	return o.PrimaryIp, true
}

// HasPrimaryIp returns a boolean if a field has been set.
func (o *Device) HasPrimaryIp() bool {
	if o != nil && o.PrimaryIp != nil {
		return true
	}

	return false
}

// SetPrimaryIp gets a reference to the given NestedIPAddress and assigns it to the PrimaryIp field.
func (o *Device) SetPrimaryIp(v NestedIPAddress) {
	o.PrimaryIp = &v
}

// GetPrimaryIp4 returns the PrimaryIp4 field value if set, zero value otherwise.
func (o *Device) GetPrimaryIp4() NestedIPAddress {
	if o == nil || o.PrimaryIp4 == nil {
		var ret NestedIPAddress
		return ret
	}
	return *o.PrimaryIp4
}

// GetPrimaryIp4Ok returns a tuple with the PrimaryIp4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetPrimaryIp4Ok() (*NestedIPAddress, bool) {
	if o == nil || o.PrimaryIp4 == nil {
		return nil, false
	}
	return o.PrimaryIp4, true
}

// HasPrimaryIp4 returns a boolean if a field has been set.
func (o *Device) HasPrimaryIp4() bool {
	if o != nil && o.PrimaryIp4 != nil {
		return true
	}

	return false
}

// SetPrimaryIp4 gets a reference to the given NestedIPAddress and assigns it to the PrimaryIp4 field.
func (o *Device) SetPrimaryIp4(v NestedIPAddress) {
	o.PrimaryIp4 = &v
}

// GetPrimaryIp6 returns the PrimaryIp6 field value if set, zero value otherwise.
func (o *Device) GetPrimaryIp6() NestedIPAddress {
	if o == nil || o.PrimaryIp6 == nil {
		var ret NestedIPAddress
		return ret
	}
	return *o.PrimaryIp6
}

// GetPrimaryIp6Ok returns a tuple with the PrimaryIp6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetPrimaryIp6Ok() (*NestedIPAddress, bool) {
	if o == nil || o.PrimaryIp6 == nil {
		return nil, false
	}
	return o.PrimaryIp6, true
}

// HasPrimaryIp6 returns a boolean if a field has been set.
func (o *Device) HasPrimaryIp6() bool {
	if o != nil && o.PrimaryIp6 != nil {
		return true
	}

	return false
}

// SetPrimaryIp6 gets a reference to the given NestedIPAddress and assigns it to the PrimaryIp6 field.
func (o *Device) SetPrimaryIp6(v NestedIPAddress) {
	o.PrimaryIp6 = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Device) GetCluster() NestedCluster {
	if o == nil || o.Cluster.Get() == nil {
		var ret NestedCluster
		return ret
	}
	return *o.Cluster.Get()
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Device) GetClusterOk() (*NestedCluster, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cluster.Get(), o.Cluster.IsSet()
}

// HasCluster returns a boolean if a field has been set.
func (o *Device) HasCluster() bool {
	if o != nil && o.Cluster.IsSet() {
		return true
	}

	return false
}

// SetCluster gets a reference to the given NullableNestedCluster and assigns it to the Cluster field.
func (o *Device) SetCluster(v NestedCluster) {
	o.Cluster.Set(&v)
}
// SetClusterNil sets the value for Cluster to be an explicit nil
func (o *Device) SetClusterNil() {
	o.Cluster.Set(nil)
}

// UnsetCluster ensures that no value is present for Cluster, not even an explicit nil
func (o *Device) UnsetCluster() {
	o.Cluster.Unset()
}

// GetVirtualChassis returns the VirtualChassis field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Device) GetVirtualChassis() NestedVirtualChassis {
	if o == nil || o.VirtualChassis.Get() == nil {
		var ret NestedVirtualChassis
		return ret
	}
	return *o.VirtualChassis.Get()
}

// GetVirtualChassisOk returns a tuple with the VirtualChassis field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Device) GetVirtualChassisOk() (*NestedVirtualChassis, bool) {
	if o == nil {
		return nil, false
	}
	return o.VirtualChassis.Get(), o.VirtualChassis.IsSet()
}

// HasVirtualChassis returns a boolean if a field has been set.
func (o *Device) HasVirtualChassis() bool {
	if o != nil && o.VirtualChassis.IsSet() {
		return true
	}

	return false
}

// SetVirtualChassis gets a reference to the given NullableNestedVirtualChassis and assigns it to the VirtualChassis field.
func (o *Device) SetVirtualChassis(v NestedVirtualChassis) {
	o.VirtualChassis.Set(&v)
}
// SetVirtualChassisNil sets the value for VirtualChassis to be an explicit nil
func (o *Device) SetVirtualChassisNil() {
	o.VirtualChassis.Set(nil)
}

// UnsetVirtualChassis ensures that no value is present for VirtualChassis, not even an explicit nil
func (o *Device) UnsetVirtualChassis() {
	o.VirtualChassis.Unset()
}

// GetVcPosition returns the VcPosition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Device) GetVcPosition() int32 {
	if o == nil || o.VcPosition.Get() == nil {
		var ret int32
		return ret
	}
	return *o.VcPosition.Get()
}

// GetVcPositionOk returns a tuple with the VcPosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Device) GetVcPositionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.VcPosition.Get(), o.VcPosition.IsSet()
}

// HasVcPosition returns a boolean if a field has been set.
func (o *Device) HasVcPosition() bool {
	if o != nil && o.VcPosition.IsSet() {
		return true
	}

	return false
}

// SetVcPosition gets a reference to the given NullableInt32 and assigns it to the VcPosition field.
func (o *Device) SetVcPosition(v int32) {
	o.VcPosition.Set(&v)
}
// SetVcPositionNil sets the value for VcPosition to be an explicit nil
func (o *Device) SetVcPositionNil() {
	o.VcPosition.Set(nil)
}

// UnsetVcPosition ensures that no value is present for VcPosition, not even an explicit nil
func (o *Device) UnsetVcPosition() {
	o.VcPosition.Unset()
}

// GetVcPriority returns the VcPriority field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Device) GetVcPriority() int32 {
	if o == nil || o.VcPriority.Get() == nil {
		var ret int32
		return ret
	}
	return *o.VcPriority.Get()
}

// GetVcPriorityOk returns a tuple with the VcPriority field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Device) GetVcPriorityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.VcPriority.Get(), o.VcPriority.IsSet()
}

// HasVcPriority returns a boolean if a field has been set.
func (o *Device) HasVcPriority() bool {
	if o != nil && o.VcPriority.IsSet() {
		return true
	}

	return false
}

// SetVcPriority gets a reference to the given NullableInt32 and assigns it to the VcPriority field.
func (o *Device) SetVcPriority(v int32) {
	o.VcPriority.Set(&v)
}
// SetVcPriorityNil sets the value for VcPriority to be an explicit nil
func (o *Device) SetVcPriorityNil() {
	o.VcPriority.Set(nil)
}

// UnsetVcPriority ensures that no value is present for VcPriority, not even an explicit nil
func (o *Device) UnsetVcPriority() {
	o.VcPriority.Unset()
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *Device) GetComments() string {
	if o == nil || o.Comments == nil {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetCommentsOk() (*string, bool) {
	if o == nil || o.Comments == nil {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *Device) HasComments() bool {
	if o != nil && o.Comments != nil {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *Device) SetComments(v string) {
	o.Comments = &v
}

// GetLocalContextData returns the LocalContextData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Device) GetLocalContextData() string {
	if o == nil || o.LocalContextData.Get() == nil {
		var ret string
		return ret
	}
	return *o.LocalContextData.Get()
}

// GetLocalContextDataOk returns a tuple with the LocalContextData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Device) GetLocalContextDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LocalContextData.Get(), o.LocalContextData.IsSet()
}

// HasLocalContextData returns a boolean if a field has been set.
func (o *Device) HasLocalContextData() bool {
	if o != nil && o.LocalContextData.IsSet() {
		return true
	}

	return false
}

// SetLocalContextData gets a reference to the given NullableString and assigns it to the LocalContextData field.
func (o *Device) SetLocalContextData(v string) {
	o.LocalContextData.Set(&v)
}
// SetLocalContextDataNil sets the value for LocalContextData to be an explicit nil
func (o *Device) SetLocalContextDataNil() {
	o.LocalContextData.Set(nil)
}

// UnsetLocalContextData ensures that no value is present for LocalContextData, not even an explicit nil
func (o *Device) UnsetLocalContextData() {
	o.LocalContextData.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Device) GetTags() []NestedTag {
	if o == nil || o.Tags == nil {
		var ret []NestedTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetTagsOk() ([]NestedTag, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Device) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTag and assigns it to the Tags field.
func (o *Device) SetTags(v []NestedTag) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *Device) GetCustomFields() map[string]interface{} {
	if o == nil || o.CustomFields == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || o.CustomFields == nil {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *Device) HasCustomFields() bool {
	if o != nil && o.CustomFields != nil {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *Device) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *Device) GetCreated() string {
	if o == nil || o.Created == nil {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetCreatedOk() (*string, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *Device) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *Device) SetCreated(v string) {
	o.Created = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *Device) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *Device) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *Device) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

func (o Device) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.Display != nil {
		toSerialize["display"] = o.Display
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	if true {
		toSerialize["device_type"] = o.DeviceType
	}
	if true {
		toSerialize["device_role"] = o.DeviceRole
	}
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
	if o.Platform.IsSet() {
		toSerialize["platform"] = o.Platform.Get()
	}
	if o.Serial != nil {
		toSerialize["serial"] = o.Serial
	}
	if o.AssetTag.IsSet() {
		toSerialize["asset_tag"] = o.AssetTag.Get()
	}
	if true {
		toSerialize["site"] = o.Site
	}
	if o.Location.IsSet() {
		toSerialize["location"] = o.Location.Get()
	}
	if o.Rack.IsSet() {
		toSerialize["rack"] = o.Rack.Get()
	}
	if o.Position.IsSet() {
		toSerialize["position"] = o.Position.Get()
	}
	if o.Face != nil {
		toSerialize["face"] = o.Face
	}
	if o.ParentDevice != nil {
		toSerialize["parent_device"] = o.ParentDevice
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.PrimaryIp != nil {
		toSerialize["primary_ip"] = o.PrimaryIp
	}
	if o.PrimaryIp4 != nil {
		toSerialize["primary_ip4"] = o.PrimaryIp4
	}
	if o.PrimaryIp6 != nil {
		toSerialize["primary_ip6"] = o.PrimaryIp6
	}
	if o.Cluster.IsSet() {
		toSerialize["cluster"] = o.Cluster.Get()
	}
	if o.VirtualChassis.IsSet() {
		toSerialize["virtual_chassis"] = o.VirtualChassis.Get()
	}
	if o.VcPosition.IsSet() {
		toSerialize["vc_position"] = o.VcPosition.Get()
	}
	if o.VcPriority.IsSet() {
		toSerialize["vc_priority"] = o.VcPriority.Get()
	}
	if o.Comments != nil {
		toSerialize["comments"] = o.Comments
	}
	if o.LocalContextData.IsSet() {
		toSerialize["local_context_data"] = o.LocalContextData.Get()
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.CustomFields != nil {
		toSerialize["custom_fields"] = o.CustomFields
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.LastUpdated != nil {
		toSerialize["last_updated"] = o.LastUpdated
	}
	return json.Marshal(toSerialize)
}

type NullableDevice struct {
	value *Device
	isSet bool
}

func (v NullableDevice) Get() *Device {
	return v.value
}

func (v *NullableDevice) Set(val *Device) {
	v.value = val
	v.isSet = true
}

func (v NullableDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevice(val *Device) *NullableDevice {
	return &NullableDevice{value: val, isSet: true}
}

func (v NullableDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


