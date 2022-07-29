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

// WritableVirtualMachineWithConfigContext struct for WritableVirtualMachineWithConfigContext
type WritableVirtualMachineWithConfigContext struct {
	Id *int32 `json:"id,omitempty"`
	Url *string `json:"url,omitempty"`
	Display *string `json:"display,omitempty"`
	Name string `json:"name"`
	Status *string `json:"status,omitempty"`
	Site *string `json:"site,omitempty"`
	Cluster int32 `json:"cluster"`
	Role NullableInt32 `json:"role,omitempty"`
	Tenant NullableInt32 `json:"tenant,omitempty"`
	Platform NullableInt32 `json:"platform,omitempty"`
	PrimaryIp *string `json:"primary_ip,omitempty"`
	PrimaryIp4 NullableInt32 `json:"primary_ip4,omitempty"`
	PrimaryIp6 NullableInt32 `json:"primary_ip6,omitempty"`
	Vcpus NullableFloat64 `json:"vcpus,omitempty"`
	Memory NullableInt32 `json:"memory,omitempty"`
	Disk NullableInt32 `json:"disk,omitempty"`
	Comments *string `json:"comments,omitempty"`
	LocalContextData NullableString `json:"local_context_data,omitempty"`
	Tags []NestedTag `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	ConfigContext *map[string]string `json:"config_context,omitempty"`
	Created *string `json:"created,omitempty"`
	LastUpdated *time.Time `json:"last_updated,omitempty"`
}

// NewWritableVirtualMachineWithConfigContext instantiates a new WritableVirtualMachineWithConfigContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWritableVirtualMachineWithConfigContext(name string, cluster int32) *WritableVirtualMachineWithConfigContext {
	this := WritableVirtualMachineWithConfigContext{}
	this.Name = name
	this.Cluster = cluster
	return &this
}

// NewWritableVirtualMachineWithConfigContextWithDefaults instantiates a new WritableVirtualMachineWithConfigContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWritableVirtualMachineWithConfigContextWithDefaults() *WritableVirtualMachineWithConfigContext {
	this := WritableVirtualMachineWithConfigContext{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WritableVirtualMachineWithConfigContext) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableVirtualMachineWithConfigContext) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WritableVirtualMachineWithConfigContext) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *WritableVirtualMachineWithConfigContext) SetId(v int32) {
	o.Id = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *WritableVirtualMachineWithConfigContext) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableVirtualMachineWithConfigContext) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *WritableVirtualMachineWithConfigContext) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *WritableVirtualMachineWithConfigContext) SetUrl(v string) {
	o.Url = &v
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *WritableVirtualMachineWithConfigContext) GetDisplay() string {
	if o == nil || o.Display == nil {
		var ret string
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableVirtualMachineWithConfigContext) GetDisplayOk() (*string, bool) {
	if o == nil || o.Display == nil {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *WritableVirtualMachineWithConfigContext) HasDisplay() bool {
	if o != nil && o.Display != nil {
		return true
	}

	return false
}

// SetDisplay gets a reference to the given string and assigns it to the Display field.
func (o *WritableVirtualMachineWithConfigContext) SetDisplay(v string) {
	o.Display = &v
}

// GetName returns the Name field value
func (o *WritableVirtualMachineWithConfigContext) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WritableVirtualMachineWithConfigContext) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WritableVirtualMachineWithConfigContext) SetName(v string) {
	o.Name = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *WritableVirtualMachineWithConfigContext) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableVirtualMachineWithConfigContext) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *WritableVirtualMachineWithConfigContext) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *WritableVirtualMachineWithConfigContext) SetStatus(v string) {
	o.Status = &v
}

// GetSite returns the Site field value if set, zero value otherwise.
func (o *WritableVirtualMachineWithConfigContext) GetSite() string {
	if o == nil || o.Site == nil {
		var ret string
		return ret
	}
	return *o.Site
}

// GetSiteOk returns a tuple with the Site field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableVirtualMachineWithConfigContext) GetSiteOk() (*string, bool) {
	if o == nil || o.Site == nil {
		return nil, false
	}
	return o.Site, true
}

// HasSite returns a boolean if a field has been set.
func (o *WritableVirtualMachineWithConfigContext) HasSite() bool {
	if o != nil && o.Site != nil {
		return true
	}

	return false
}

// SetSite gets a reference to the given string and assigns it to the Site field.
func (o *WritableVirtualMachineWithConfigContext) SetSite(v string) {
	o.Site = &v
}

// GetCluster returns the Cluster field value
func (o *WritableVirtualMachineWithConfigContext) GetCluster() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value
// and a boolean to check if the value has been set.
func (o *WritableVirtualMachineWithConfigContext) GetClusterOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cluster, true
}

// SetCluster sets field value
func (o *WritableVirtualMachineWithConfigContext) SetCluster(v int32) {
	o.Cluster = v
}

// GetRole returns the Role field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableVirtualMachineWithConfigContext) GetRole() int32 {
	if o == nil || o.Role.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Role.Get()
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableVirtualMachineWithConfigContext) GetRoleOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Role.Get(), o.Role.IsSet()
}

// HasRole returns a boolean if a field has been set.
func (o *WritableVirtualMachineWithConfigContext) HasRole() bool {
	if o != nil && o.Role.IsSet() {
		return true
	}

	return false
}

// SetRole gets a reference to the given NullableInt32 and assigns it to the Role field.
func (o *WritableVirtualMachineWithConfigContext) SetRole(v int32) {
	o.Role.Set(&v)
}
// SetRoleNil sets the value for Role to be an explicit nil
func (o *WritableVirtualMachineWithConfigContext) SetRoleNil() {
	o.Role.Set(nil)
}

// UnsetRole ensures that no value is present for Role, not even an explicit nil
func (o *WritableVirtualMachineWithConfigContext) UnsetRole() {
	o.Role.Unset()
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableVirtualMachineWithConfigContext) GetTenant() int32 {
	if o == nil || o.Tenant.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableVirtualMachineWithConfigContext) GetTenantOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *WritableVirtualMachineWithConfigContext) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableInt32 and assigns it to the Tenant field.
func (o *WritableVirtualMachineWithConfigContext) SetTenant(v int32) {
	o.Tenant.Set(&v)
}
// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *WritableVirtualMachineWithConfigContext) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *WritableVirtualMachineWithConfigContext) UnsetTenant() {
	o.Tenant.Unset()
}

// GetPlatform returns the Platform field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableVirtualMachineWithConfigContext) GetPlatform() int32 {
	if o == nil || o.Platform.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Platform.Get()
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableVirtualMachineWithConfigContext) GetPlatformOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Platform.Get(), o.Platform.IsSet()
}

// HasPlatform returns a boolean if a field has been set.
func (o *WritableVirtualMachineWithConfigContext) HasPlatform() bool {
	if o != nil && o.Platform.IsSet() {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given NullableInt32 and assigns it to the Platform field.
func (o *WritableVirtualMachineWithConfigContext) SetPlatform(v int32) {
	o.Platform.Set(&v)
}
// SetPlatformNil sets the value for Platform to be an explicit nil
func (o *WritableVirtualMachineWithConfigContext) SetPlatformNil() {
	o.Platform.Set(nil)
}

// UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
func (o *WritableVirtualMachineWithConfigContext) UnsetPlatform() {
	o.Platform.Unset()
}

// GetPrimaryIp returns the PrimaryIp field value if set, zero value otherwise.
func (o *WritableVirtualMachineWithConfigContext) GetPrimaryIp() string {
	if o == nil || o.PrimaryIp == nil {
		var ret string
		return ret
	}
	return *o.PrimaryIp
}

// GetPrimaryIpOk returns a tuple with the PrimaryIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableVirtualMachineWithConfigContext) GetPrimaryIpOk() (*string, bool) {
	if o == nil || o.PrimaryIp == nil {
		return nil, false
	}
	return o.PrimaryIp, true
}

// HasPrimaryIp returns a boolean if a field has been set.
func (o *WritableVirtualMachineWithConfigContext) HasPrimaryIp() bool {
	if o != nil && o.PrimaryIp != nil {
		return true
	}

	return false
}

// SetPrimaryIp gets a reference to the given string and assigns it to the PrimaryIp field.
func (o *WritableVirtualMachineWithConfigContext) SetPrimaryIp(v string) {
	o.PrimaryIp = &v
}

// GetPrimaryIp4 returns the PrimaryIp4 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableVirtualMachineWithConfigContext) GetPrimaryIp4() int32 {
	if o == nil || o.PrimaryIp4.Get() == nil {
		var ret int32
		return ret
	}
	return *o.PrimaryIp4.Get()
}

// GetPrimaryIp4Ok returns a tuple with the PrimaryIp4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableVirtualMachineWithConfigContext) GetPrimaryIp4Ok() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrimaryIp4.Get(), o.PrimaryIp4.IsSet()
}

// HasPrimaryIp4 returns a boolean if a field has been set.
func (o *WritableVirtualMachineWithConfigContext) HasPrimaryIp4() bool {
	if o != nil && o.PrimaryIp4.IsSet() {
		return true
	}

	return false
}

// SetPrimaryIp4 gets a reference to the given NullableInt32 and assigns it to the PrimaryIp4 field.
func (o *WritableVirtualMachineWithConfigContext) SetPrimaryIp4(v int32) {
	o.PrimaryIp4.Set(&v)
}
// SetPrimaryIp4Nil sets the value for PrimaryIp4 to be an explicit nil
func (o *WritableVirtualMachineWithConfigContext) SetPrimaryIp4Nil() {
	o.PrimaryIp4.Set(nil)
}

// UnsetPrimaryIp4 ensures that no value is present for PrimaryIp4, not even an explicit nil
func (o *WritableVirtualMachineWithConfigContext) UnsetPrimaryIp4() {
	o.PrimaryIp4.Unset()
}

// GetPrimaryIp6 returns the PrimaryIp6 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableVirtualMachineWithConfigContext) GetPrimaryIp6() int32 {
	if o == nil || o.PrimaryIp6.Get() == nil {
		var ret int32
		return ret
	}
	return *o.PrimaryIp6.Get()
}

// GetPrimaryIp6Ok returns a tuple with the PrimaryIp6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableVirtualMachineWithConfigContext) GetPrimaryIp6Ok() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrimaryIp6.Get(), o.PrimaryIp6.IsSet()
}

// HasPrimaryIp6 returns a boolean if a field has been set.
func (o *WritableVirtualMachineWithConfigContext) HasPrimaryIp6() bool {
	if o != nil && o.PrimaryIp6.IsSet() {
		return true
	}

	return false
}

// SetPrimaryIp6 gets a reference to the given NullableInt32 and assigns it to the PrimaryIp6 field.
func (o *WritableVirtualMachineWithConfigContext) SetPrimaryIp6(v int32) {
	o.PrimaryIp6.Set(&v)
}
// SetPrimaryIp6Nil sets the value for PrimaryIp6 to be an explicit nil
func (o *WritableVirtualMachineWithConfigContext) SetPrimaryIp6Nil() {
	o.PrimaryIp6.Set(nil)
}

// UnsetPrimaryIp6 ensures that no value is present for PrimaryIp6, not even an explicit nil
func (o *WritableVirtualMachineWithConfigContext) UnsetPrimaryIp6() {
	o.PrimaryIp6.Unset()
}

// GetVcpus returns the Vcpus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableVirtualMachineWithConfigContext) GetVcpus() float64 {
	if o == nil || o.Vcpus.Get() == nil {
		var ret float64
		return ret
	}
	return *o.Vcpus.Get()
}

// GetVcpusOk returns a tuple with the Vcpus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableVirtualMachineWithConfigContext) GetVcpusOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vcpus.Get(), o.Vcpus.IsSet()
}

// HasVcpus returns a boolean if a field has been set.
func (o *WritableVirtualMachineWithConfigContext) HasVcpus() bool {
	if o != nil && o.Vcpus.IsSet() {
		return true
	}

	return false
}

// SetVcpus gets a reference to the given NullableFloat64 and assigns it to the Vcpus field.
func (o *WritableVirtualMachineWithConfigContext) SetVcpus(v float64) {
	o.Vcpus.Set(&v)
}
// SetVcpusNil sets the value for Vcpus to be an explicit nil
func (o *WritableVirtualMachineWithConfigContext) SetVcpusNil() {
	o.Vcpus.Set(nil)
}

// UnsetVcpus ensures that no value is present for Vcpus, not even an explicit nil
func (o *WritableVirtualMachineWithConfigContext) UnsetVcpus() {
	o.Vcpus.Unset()
}

// GetMemory returns the Memory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableVirtualMachineWithConfigContext) GetMemory() int32 {
	if o == nil || o.Memory.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Memory.Get()
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableVirtualMachineWithConfigContext) GetMemoryOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Memory.Get(), o.Memory.IsSet()
}

// HasMemory returns a boolean if a field has been set.
func (o *WritableVirtualMachineWithConfigContext) HasMemory() bool {
	if o != nil && o.Memory.IsSet() {
		return true
	}

	return false
}

// SetMemory gets a reference to the given NullableInt32 and assigns it to the Memory field.
func (o *WritableVirtualMachineWithConfigContext) SetMemory(v int32) {
	o.Memory.Set(&v)
}
// SetMemoryNil sets the value for Memory to be an explicit nil
func (o *WritableVirtualMachineWithConfigContext) SetMemoryNil() {
	o.Memory.Set(nil)
}

// UnsetMemory ensures that no value is present for Memory, not even an explicit nil
func (o *WritableVirtualMachineWithConfigContext) UnsetMemory() {
	o.Memory.Unset()
}

// GetDisk returns the Disk field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableVirtualMachineWithConfigContext) GetDisk() int32 {
	if o == nil || o.Disk.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Disk.Get()
}

// GetDiskOk returns a tuple with the Disk field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableVirtualMachineWithConfigContext) GetDiskOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Disk.Get(), o.Disk.IsSet()
}

// HasDisk returns a boolean if a field has been set.
func (o *WritableVirtualMachineWithConfigContext) HasDisk() bool {
	if o != nil && o.Disk.IsSet() {
		return true
	}

	return false
}

// SetDisk gets a reference to the given NullableInt32 and assigns it to the Disk field.
func (o *WritableVirtualMachineWithConfigContext) SetDisk(v int32) {
	o.Disk.Set(&v)
}
// SetDiskNil sets the value for Disk to be an explicit nil
func (o *WritableVirtualMachineWithConfigContext) SetDiskNil() {
	o.Disk.Set(nil)
}

// UnsetDisk ensures that no value is present for Disk, not even an explicit nil
func (o *WritableVirtualMachineWithConfigContext) UnsetDisk() {
	o.Disk.Unset()
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *WritableVirtualMachineWithConfigContext) GetComments() string {
	if o == nil || o.Comments == nil {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableVirtualMachineWithConfigContext) GetCommentsOk() (*string, bool) {
	if o == nil || o.Comments == nil {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *WritableVirtualMachineWithConfigContext) HasComments() bool {
	if o != nil && o.Comments != nil {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *WritableVirtualMachineWithConfigContext) SetComments(v string) {
	o.Comments = &v
}

// GetLocalContextData returns the LocalContextData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableVirtualMachineWithConfigContext) GetLocalContextData() string {
	if o == nil || o.LocalContextData.Get() == nil {
		var ret string
		return ret
	}
	return *o.LocalContextData.Get()
}

// GetLocalContextDataOk returns a tuple with the LocalContextData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableVirtualMachineWithConfigContext) GetLocalContextDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LocalContextData.Get(), o.LocalContextData.IsSet()
}

// HasLocalContextData returns a boolean if a field has been set.
func (o *WritableVirtualMachineWithConfigContext) HasLocalContextData() bool {
	if o != nil && o.LocalContextData.IsSet() {
		return true
	}

	return false
}

// SetLocalContextData gets a reference to the given NullableString and assigns it to the LocalContextData field.
func (o *WritableVirtualMachineWithConfigContext) SetLocalContextData(v string) {
	o.LocalContextData.Set(&v)
}
// SetLocalContextDataNil sets the value for LocalContextData to be an explicit nil
func (o *WritableVirtualMachineWithConfigContext) SetLocalContextDataNil() {
	o.LocalContextData.Set(nil)
}

// UnsetLocalContextData ensures that no value is present for LocalContextData, not even an explicit nil
func (o *WritableVirtualMachineWithConfigContext) UnsetLocalContextData() {
	o.LocalContextData.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *WritableVirtualMachineWithConfigContext) GetTags() []NestedTag {
	if o == nil || o.Tags == nil {
		var ret []NestedTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableVirtualMachineWithConfigContext) GetTagsOk() ([]NestedTag, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *WritableVirtualMachineWithConfigContext) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTag and assigns it to the Tags field.
func (o *WritableVirtualMachineWithConfigContext) SetTags(v []NestedTag) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *WritableVirtualMachineWithConfigContext) GetCustomFields() map[string]interface{} {
	if o == nil || o.CustomFields == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableVirtualMachineWithConfigContext) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || o.CustomFields == nil {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *WritableVirtualMachineWithConfigContext) HasCustomFields() bool {
	if o != nil && o.CustomFields != nil {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *WritableVirtualMachineWithConfigContext) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetConfigContext returns the ConfigContext field value if set, zero value otherwise.
func (o *WritableVirtualMachineWithConfigContext) GetConfigContext() map[string]string {
	if o == nil || o.ConfigContext == nil {
		var ret map[string]string
		return ret
	}
	return *o.ConfigContext
}

// GetConfigContextOk returns a tuple with the ConfigContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableVirtualMachineWithConfigContext) GetConfigContextOk() (*map[string]string, bool) {
	if o == nil || o.ConfigContext == nil {
		return nil, false
	}
	return o.ConfigContext, true
}

// HasConfigContext returns a boolean if a field has been set.
func (o *WritableVirtualMachineWithConfigContext) HasConfigContext() bool {
	if o != nil && o.ConfigContext != nil {
		return true
	}

	return false
}

// SetConfigContext gets a reference to the given map[string]string and assigns it to the ConfigContext field.
func (o *WritableVirtualMachineWithConfigContext) SetConfigContext(v map[string]string) {
	o.ConfigContext = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *WritableVirtualMachineWithConfigContext) GetCreated() string {
	if o == nil || o.Created == nil {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableVirtualMachineWithConfigContext) GetCreatedOk() (*string, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *WritableVirtualMachineWithConfigContext) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *WritableVirtualMachineWithConfigContext) SetCreated(v string) {
	o.Created = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *WritableVirtualMachineWithConfigContext) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableVirtualMachineWithConfigContext) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *WritableVirtualMachineWithConfigContext) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *WritableVirtualMachineWithConfigContext) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

func (o WritableVirtualMachineWithConfigContext) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Site != nil {
		toSerialize["site"] = o.Site
	}
	if true {
		toSerialize["cluster"] = o.Cluster
	}
	if o.Role.IsSet() {
		toSerialize["role"] = o.Role.Get()
	}
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
	if o.Platform.IsSet() {
		toSerialize["platform"] = o.Platform.Get()
	}
	if o.PrimaryIp != nil {
		toSerialize["primary_ip"] = o.PrimaryIp
	}
	if o.PrimaryIp4.IsSet() {
		toSerialize["primary_ip4"] = o.PrimaryIp4.Get()
	}
	if o.PrimaryIp6.IsSet() {
		toSerialize["primary_ip6"] = o.PrimaryIp6.Get()
	}
	if o.Vcpus.IsSet() {
		toSerialize["vcpus"] = o.Vcpus.Get()
	}
	if o.Memory.IsSet() {
		toSerialize["memory"] = o.Memory.Get()
	}
	if o.Disk.IsSet() {
		toSerialize["disk"] = o.Disk.Get()
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
	if o.ConfigContext != nil {
		toSerialize["config_context"] = o.ConfigContext
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.LastUpdated != nil {
		toSerialize["last_updated"] = o.LastUpdated
	}
	return json.Marshal(toSerialize)
}

type NullableWritableVirtualMachineWithConfigContext struct {
	value *WritableVirtualMachineWithConfigContext
	isSet bool
}

func (v NullableWritableVirtualMachineWithConfigContext) Get() *WritableVirtualMachineWithConfigContext {
	return v.value
}

func (v *NullableWritableVirtualMachineWithConfigContext) Set(val *WritableVirtualMachineWithConfigContext) {
	v.value = val
	v.isSet = true
}

func (v NullableWritableVirtualMachineWithConfigContext) IsSet() bool {
	return v.isSet
}

func (v *NullableWritableVirtualMachineWithConfigContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWritableVirtualMachineWithConfigContext(val *WritableVirtualMachineWithConfigContext) *NullableWritableVirtualMachineWithConfigContext {
	return &NullableWritableVirtualMachineWithConfigContext{value: val, isSet: true}
}

func (v NullableWritableVirtualMachineWithConfigContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWritableVirtualMachineWithConfigContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

