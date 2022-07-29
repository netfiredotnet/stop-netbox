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

// Site struct for Site
type Site struct {
	Id *int32 `json:"id,omitempty"`
	Url *string `json:"url,omitempty"`
	Display *string `json:"display,omitempty"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	Status *Status5 `json:"status,omitempty"`
	Region NullableNestedRegion `json:"region,omitempty"`
	Group NullableNestedSiteGroup `json:"group,omitempty"`
	Tenant NullableNestedTenant `json:"tenant,omitempty"`
	// Local facility ID or description
	Facility *string `json:"facility,omitempty"`
	// 32-bit autonomous system number
	Asn NullableInt32 `json:"asn,omitempty"`
	TimeZone *string `json:"time_zone,omitempty"`
	Description *string `json:"description,omitempty"`
	PhysicalAddress *string `json:"physical_address,omitempty"`
	ShippingAddress *string `json:"shipping_address,omitempty"`
	// GPS coordinate (latitude)
	Latitude NullableFloat64 `json:"latitude,omitempty"`
	// GPS coordinate (longitude)
	Longitude NullableFloat64 `json:"longitude,omitempty"`
	ContactName *string `json:"contact_name,omitempty"`
	ContactPhone *string `json:"contact_phone,omitempty"`
	ContactEmail *string `json:"contact_email,omitempty"`
	Comments *string `json:"comments,omitempty"`
	Tags []NestedTag `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Created *string `json:"created,omitempty"`
	LastUpdated *time.Time `json:"last_updated,omitempty"`
	CircuitCount *int32 `json:"circuit_count,omitempty"`
	DeviceCount *int32 `json:"device_count,omitempty"`
	PrefixCount *int32 `json:"prefix_count,omitempty"`
	RackCount *int32 `json:"rack_count,omitempty"`
	VirtualmachineCount *int32 `json:"virtualmachine_count,omitempty"`
	VlanCount *int32 `json:"vlan_count,omitempty"`
}

// NewSite instantiates a new Site object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSite(name string, slug string) *Site {
	this := Site{}
	this.Name = name
	this.Slug = slug
	return &this
}

// NewSiteWithDefaults instantiates a new Site object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteWithDefaults() *Site {
	this := Site{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Site) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Site) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Site) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Site) SetId(v int32) {
	o.Id = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *Site) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Site) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *Site) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *Site) SetUrl(v string) {
	o.Url = &v
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *Site) GetDisplay() string {
	if o == nil || o.Display == nil {
		var ret string
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Site) GetDisplayOk() (*string, bool) {
	if o == nil || o.Display == nil {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *Site) HasDisplay() bool {
	if o != nil && o.Display != nil {
		return true
	}

	return false
}

// SetDisplay gets a reference to the given string and assigns it to the Display field.
func (o *Site) SetDisplay(v string) {
	o.Display = &v
}

// GetName returns the Name field value
func (o *Site) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Site) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Site) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *Site) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *Site) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *Site) SetSlug(v string) {
	o.Slug = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Site) GetStatus() Status5 {
	if o == nil || o.Status == nil {
		var ret Status5
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Site) GetStatusOk() (*Status5, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Site) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given Status5 and assigns it to the Status field.
func (o *Site) SetStatus(v Status5) {
	o.Status = &v
}

// GetRegion returns the Region field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Site) GetRegion() NestedRegion {
	if o == nil || o.Region.Get() == nil {
		var ret NestedRegion
		return ret
	}
	return *o.Region.Get()
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Site) GetRegionOk() (*NestedRegion, bool) {
	if o == nil {
		return nil, false
	}
	return o.Region.Get(), o.Region.IsSet()
}

// HasRegion returns a boolean if a field has been set.
func (o *Site) HasRegion() bool {
	if o != nil && o.Region.IsSet() {
		return true
	}

	return false
}

// SetRegion gets a reference to the given NullableNestedRegion and assigns it to the Region field.
func (o *Site) SetRegion(v NestedRegion) {
	o.Region.Set(&v)
}
// SetRegionNil sets the value for Region to be an explicit nil
func (o *Site) SetRegionNil() {
	o.Region.Set(nil)
}

// UnsetRegion ensures that no value is present for Region, not even an explicit nil
func (o *Site) UnsetRegion() {
	o.Region.Unset()
}

// GetGroup returns the Group field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Site) GetGroup() NestedSiteGroup {
	if o == nil || o.Group.Get() == nil {
		var ret NestedSiteGroup
		return ret
	}
	return *o.Group.Get()
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Site) GetGroupOk() (*NestedSiteGroup, bool) {
	if o == nil {
		return nil, false
	}
	return o.Group.Get(), o.Group.IsSet()
}

// HasGroup returns a boolean if a field has been set.
func (o *Site) HasGroup() bool {
	if o != nil && o.Group.IsSet() {
		return true
	}

	return false
}

// SetGroup gets a reference to the given NullableNestedSiteGroup and assigns it to the Group field.
func (o *Site) SetGroup(v NestedSiteGroup) {
	o.Group.Set(&v)
}
// SetGroupNil sets the value for Group to be an explicit nil
func (o *Site) SetGroupNil() {
	o.Group.Set(nil)
}

// UnsetGroup ensures that no value is present for Group, not even an explicit nil
func (o *Site) UnsetGroup() {
	o.Group.Unset()
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Site) GetTenant() NestedTenant {
	if o == nil || o.Tenant.Get() == nil {
		var ret NestedTenant
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Site) GetTenantOk() (*NestedTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *Site) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableNestedTenant and assigns it to the Tenant field.
func (o *Site) SetTenant(v NestedTenant) {
	o.Tenant.Set(&v)
}
// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *Site) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *Site) UnsetTenant() {
	o.Tenant.Unset()
}

// GetFacility returns the Facility field value if set, zero value otherwise.
func (o *Site) GetFacility() string {
	if o == nil || o.Facility == nil {
		var ret string
		return ret
	}
	return *o.Facility
}

// GetFacilityOk returns a tuple with the Facility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Site) GetFacilityOk() (*string, bool) {
	if o == nil || o.Facility == nil {
		return nil, false
	}
	return o.Facility, true
}

// HasFacility returns a boolean if a field has been set.
func (o *Site) HasFacility() bool {
	if o != nil && o.Facility != nil {
		return true
	}

	return false
}

// SetFacility gets a reference to the given string and assigns it to the Facility field.
func (o *Site) SetFacility(v string) {
	o.Facility = &v
}

// GetAsn returns the Asn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Site) GetAsn() int32 {
	if o == nil || o.Asn.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Asn.Get()
}

// GetAsnOk returns a tuple with the Asn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Site) GetAsnOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Asn.Get(), o.Asn.IsSet()
}

// HasAsn returns a boolean if a field has been set.
func (o *Site) HasAsn() bool {
	if o != nil && o.Asn.IsSet() {
		return true
	}

	return false
}

// SetAsn gets a reference to the given NullableInt32 and assigns it to the Asn field.
func (o *Site) SetAsn(v int32) {
	o.Asn.Set(&v)
}
// SetAsnNil sets the value for Asn to be an explicit nil
func (o *Site) SetAsnNil() {
	o.Asn.Set(nil)
}

// UnsetAsn ensures that no value is present for Asn, not even an explicit nil
func (o *Site) UnsetAsn() {
	o.Asn.Unset()
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *Site) GetTimeZone() string {
	if o == nil || o.TimeZone == nil {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Site) GetTimeZoneOk() (*string, bool) {
	if o == nil || o.TimeZone == nil {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *Site) HasTimeZone() bool {
	if o != nil && o.TimeZone != nil {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *Site) SetTimeZone(v string) {
	o.TimeZone = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Site) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Site) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Site) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Site) SetDescription(v string) {
	o.Description = &v
}

// GetPhysicalAddress returns the PhysicalAddress field value if set, zero value otherwise.
func (o *Site) GetPhysicalAddress() string {
	if o == nil || o.PhysicalAddress == nil {
		var ret string
		return ret
	}
	return *o.PhysicalAddress
}

// GetPhysicalAddressOk returns a tuple with the PhysicalAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Site) GetPhysicalAddressOk() (*string, bool) {
	if o == nil || o.PhysicalAddress == nil {
		return nil, false
	}
	return o.PhysicalAddress, true
}

// HasPhysicalAddress returns a boolean if a field has been set.
func (o *Site) HasPhysicalAddress() bool {
	if o != nil && o.PhysicalAddress != nil {
		return true
	}

	return false
}

// SetPhysicalAddress gets a reference to the given string and assigns it to the PhysicalAddress field.
func (o *Site) SetPhysicalAddress(v string) {
	o.PhysicalAddress = &v
}

// GetShippingAddress returns the ShippingAddress field value if set, zero value otherwise.
func (o *Site) GetShippingAddress() string {
	if o == nil || o.ShippingAddress == nil {
		var ret string
		return ret
	}
	return *o.ShippingAddress
}

// GetShippingAddressOk returns a tuple with the ShippingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Site) GetShippingAddressOk() (*string, bool) {
	if o == nil || o.ShippingAddress == nil {
		return nil, false
	}
	return o.ShippingAddress, true
}

// HasShippingAddress returns a boolean if a field has been set.
func (o *Site) HasShippingAddress() bool {
	if o != nil && o.ShippingAddress != nil {
		return true
	}

	return false
}

// SetShippingAddress gets a reference to the given string and assigns it to the ShippingAddress field.
func (o *Site) SetShippingAddress(v string) {
	o.ShippingAddress = &v
}

// GetLatitude returns the Latitude field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Site) GetLatitude() float64 {
	if o == nil || o.Latitude.Get() == nil {
		var ret float64
		return ret
	}
	return *o.Latitude.Get()
}

// GetLatitudeOk returns a tuple with the Latitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Site) GetLatitudeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Latitude.Get(), o.Latitude.IsSet()
}

// HasLatitude returns a boolean if a field has been set.
func (o *Site) HasLatitude() bool {
	if o != nil && o.Latitude.IsSet() {
		return true
	}

	return false
}

// SetLatitude gets a reference to the given NullableFloat64 and assigns it to the Latitude field.
func (o *Site) SetLatitude(v float64) {
	o.Latitude.Set(&v)
}
// SetLatitudeNil sets the value for Latitude to be an explicit nil
func (o *Site) SetLatitudeNil() {
	o.Latitude.Set(nil)
}

// UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
func (o *Site) UnsetLatitude() {
	o.Latitude.Unset()
}

// GetLongitude returns the Longitude field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Site) GetLongitude() float64 {
	if o == nil || o.Longitude.Get() == nil {
		var ret float64
		return ret
	}
	return *o.Longitude.Get()
}

// GetLongitudeOk returns a tuple with the Longitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Site) GetLongitudeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Longitude.Get(), o.Longitude.IsSet()
}

// HasLongitude returns a boolean if a field has been set.
func (o *Site) HasLongitude() bool {
	if o != nil && o.Longitude.IsSet() {
		return true
	}

	return false
}

// SetLongitude gets a reference to the given NullableFloat64 and assigns it to the Longitude field.
func (o *Site) SetLongitude(v float64) {
	o.Longitude.Set(&v)
}
// SetLongitudeNil sets the value for Longitude to be an explicit nil
func (o *Site) SetLongitudeNil() {
	o.Longitude.Set(nil)
}

// UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil
func (o *Site) UnsetLongitude() {
	o.Longitude.Unset()
}

// GetContactName returns the ContactName field value if set, zero value otherwise.
func (o *Site) GetContactName() string {
	if o == nil || o.ContactName == nil {
		var ret string
		return ret
	}
	return *o.ContactName
}

// GetContactNameOk returns a tuple with the ContactName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Site) GetContactNameOk() (*string, bool) {
	if o == nil || o.ContactName == nil {
		return nil, false
	}
	return o.ContactName, true
}

// HasContactName returns a boolean if a field has been set.
func (o *Site) HasContactName() bool {
	if o != nil && o.ContactName != nil {
		return true
	}

	return false
}

// SetContactName gets a reference to the given string and assigns it to the ContactName field.
func (o *Site) SetContactName(v string) {
	o.ContactName = &v
}

// GetContactPhone returns the ContactPhone field value if set, zero value otherwise.
func (o *Site) GetContactPhone() string {
	if o == nil || o.ContactPhone == nil {
		var ret string
		return ret
	}
	return *o.ContactPhone
}

// GetContactPhoneOk returns a tuple with the ContactPhone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Site) GetContactPhoneOk() (*string, bool) {
	if o == nil || o.ContactPhone == nil {
		return nil, false
	}
	return o.ContactPhone, true
}

// HasContactPhone returns a boolean if a field has been set.
func (o *Site) HasContactPhone() bool {
	if o != nil && o.ContactPhone != nil {
		return true
	}

	return false
}

// SetContactPhone gets a reference to the given string and assigns it to the ContactPhone field.
func (o *Site) SetContactPhone(v string) {
	o.ContactPhone = &v
}

// GetContactEmail returns the ContactEmail field value if set, zero value otherwise.
func (o *Site) GetContactEmail() string {
	if o == nil || o.ContactEmail == nil {
		var ret string
		return ret
	}
	return *o.ContactEmail
}

// GetContactEmailOk returns a tuple with the ContactEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Site) GetContactEmailOk() (*string, bool) {
	if o == nil || o.ContactEmail == nil {
		return nil, false
	}
	return o.ContactEmail, true
}

// HasContactEmail returns a boolean if a field has been set.
func (o *Site) HasContactEmail() bool {
	if o != nil && o.ContactEmail != nil {
		return true
	}

	return false
}

// SetContactEmail gets a reference to the given string and assigns it to the ContactEmail field.
func (o *Site) SetContactEmail(v string) {
	o.ContactEmail = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *Site) GetComments() string {
	if o == nil || o.Comments == nil {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Site) GetCommentsOk() (*string, bool) {
	if o == nil || o.Comments == nil {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *Site) HasComments() bool {
	if o != nil && o.Comments != nil {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *Site) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Site) GetTags() []NestedTag {
	if o == nil || o.Tags == nil {
		var ret []NestedTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Site) GetTagsOk() ([]NestedTag, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Site) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTag and assigns it to the Tags field.
func (o *Site) SetTags(v []NestedTag) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *Site) GetCustomFields() map[string]interface{} {
	if o == nil || o.CustomFields == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Site) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || o.CustomFields == nil {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *Site) HasCustomFields() bool {
	if o != nil && o.CustomFields != nil {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *Site) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *Site) GetCreated() string {
	if o == nil || o.Created == nil {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Site) GetCreatedOk() (*string, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *Site) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *Site) SetCreated(v string) {
	o.Created = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *Site) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Site) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *Site) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *Site) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetCircuitCount returns the CircuitCount field value if set, zero value otherwise.
func (o *Site) GetCircuitCount() int32 {
	if o == nil || o.CircuitCount == nil {
		var ret int32
		return ret
	}
	return *o.CircuitCount
}

// GetCircuitCountOk returns a tuple with the CircuitCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Site) GetCircuitCountOk() (*int32, bool) {
	if o == nil || o.CircuitCount == nil {
		return nil, false
	}
	return o.CircuitCount, true
}

// HasCircuitCount returns a boolean if a field has been set.
func (o *Site) HasCircuitCount() bool {
	if o != nil && o.CircuitCount != nil {
		return true
	}

	return false
}

// SetCircuitCount gets a reference to the given int32 and assigns it to the CircuitCount field.
func (o *Site) SetCircuitCount(v int32) {
	o.CircuitCount = &v
}

// GetDeviceCount returns the DeviceCount field value if set, zero value otherwise.
func (o *Site) GetDeviceCount() int32 {
	if o == nil || o.DeviceCount == nil {
		var ret int32
		return ret
	}
	return *o.DeviceCount
}

// GetDeviceCountOk returns a tuple with the DeviceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Site) GetDeviceCountOk() (*int32, bool) {
	if o == nil || o.DeviceCount == nil {
		return nil, false
	}
	return o.DeviceCount, true
}

// HasDeviceCount returns a boolean if a field has been set.
func (o *Site) HasDeviceCount() bool {
	if o != nil && o.DeviceCount != nil {
		return true
	}

	return false
}

// SetDeviceCount gets a reference to the given int32 and assigns it to the DeviceCount field.
func (o *Site) SetDeviceCount(v int32) {
	o.DeviceCount = &v
}

// GetPrefixCount returns the PrefixCount field value if set, zero value otherwise.
func (o *Site) GetPrefixCount() int32 {
	if o == nil || o.PrefixCount == nil {
		var ret int32
		return ret
	}
	return *o.PrefixCount
}

// GetPrefixCountOk returns a tuple with the PrefixCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Site) GetPrefixCountOk() (*int32, bool) {
	if o == nil || o.PrefixCount == nil {
		return nil, false
	}
	return o.PrefixCount, true
}

// HasPrefixCount returns a boolean if a field has been set.
func (o *Site) HasPrefixCount() bool {
	if o != nil && o.PrefixCount != nil {
		return true
	}

	return false
}

// SetPrefixCount gets a reference to the given int32 and assigns it to the PrefixCount field.
func (o *Site) SetPrefixCount(v int32) {
	o.PrefixCount = &v
}

// GetRackCount returns the RackCount field value if set, zero value otherwise.
func (o *Site) GetRackCount() int32 {
	if o == nil || o.RackCount == nil {
		var ret int32
		return ret
	}
	return *o.RackCount
}

// GetRackCountOk returns a tuple with the RackCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Site) GetRackCountOk() (*int32, bool) {
	if o == nil || o.RackCount == nil {
		return nil, false
	}
	return o.RackCount, true
}

// HasRackCount returns a boolean if a field has been set.
func (o *Site) HasRackCount() bool {
	if o != nil && o.RackCount != nil {
		return true
	}

	return false
}

// SetRackCount gets a reference to the given int32 and assigns it to the RackCount field.
func (o *Site) SetRackCount(v int32) {
	o.RackCount = &v
}

// GetVirtualmachineCount returns the VirtualmachineCount field value if set, zero value otherwise.
func (o *Site) GetVirtualmachineCount() int32 {
	if o == nil || o.VirtualmachineCount == nil {
		var ret int32
		return ret
	}
	return *o.VirtualmachineCount
}

// GetVirtualmachineCountOk returns a tuple with the VirtualmachineCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Site) GetVirtualmachineCountOk() (*int32, bool) {
	if o == nil || o.VirtualmachineCount == nil {
		return nil, false
	}
	return o.VirtualmachineCount, true
}

// HasVirtualmachineCount returns a boolean if a field has been set.
func (o *Site) HasVirtualmachineCount() bool {
	if o != nil && o.VirtualmachineCount != nil {
		return true
	}

	return false
}

// SetVirtualmachineCount gets a reference to the given int32 and assigns it to the VirtualmachineCount field.
func (o *Site) SetVirtualmachineCount(v int32) {
	o.VirtualmachineCount = &v
}

// GetVlanCount returns the VlanCount field value if set, zero value otherwise.
func (o *Site) GetVlanCount() int32 {
	if o == nil || o.VlanCount == nil {
		var ret int32
		return ret
	}
	return *o.VlanCount
}

// GetVlanCountOk returns a tuple with the VlanCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Site) GetVlanCountOk() (*int32, bool) {
	if o == nil || o.VlanCount == nil {
		return nil, false
	}
	return o.VlanCount, true
}

// HasVlanCount returns a boolean if a field has been set.
func (o *Site) HasVlanCount() bool {
	if o != nil && o.VlanCount != nil {
		return true
	}

	return false
}

// SetVlanCount gets a reference to the given int32 and assigns it to the VlanCount field.
func (o *Site) SetVlanCount(v int32) {
	o.VlanCount = &v
}

func (o Site) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["slug"] = o.Slug
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Region.IsSet() {
		toSerialize["region"] = o.Region.Get()
	}
	if o.Group.IsSet() {
		toSerialize["group"] = o.Group.Get()
	}
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
	if o.Facility != nil {
		toSerialize["facility"] = o.Facility
	}
	if o.Asn.IsSet() {
		toSerialize["asn"] = o.Asn.Get()
	}
	if o.TimeZone != nil {
		toSerialize["time_zone"] = o.TimeZone
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.PhysicalAddress != nil {
		toSerialize["physical_address"] = o.PhysicalAddress
	}
	if o.ShippingAddress != nil {
		toSerialize["shipping_address"] = o.ShippingAddress
	}
	if o.Latitude.IsSet() {
		toSerialize["latitude"] = o.Latitude.Get()
	}
	if o.Longitude.IsSet() {
		toSerialize["longitude"] = o.Longitude.Get()
	}
	if o.ContactName != nil {
		toSerialize["contact_name"] = o.ContactName
	}
	if o.ContactPhone != nil {
		toSerialize["contact_phone"] = o.ContactPhone
	}
	if o.ContactEmail != nil {
		toSerialize["contact_email"] = o.ContactEmail
	}
	if o.Comments != nil {
		toSerialize["comments"] = o.Comments
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
	if o.CircuitCount != nil {
		toSerialize["circuit_count"] = o.CircuitCount
	}
	if o.DeviceCount != nil {
		toSerialize["device_count"] = o.DeviceCount
	}
	if o.PrefixCount != nil {
		toSerialize["prefix_count"] = o.PrefixCount
	}
	if o.RackCount != nil {
		toSerialize["rack_count"] = o.RackCount
	}
	if o.VirtualmachineCount != nil {
		toSerialize["virtualmachine_count"] = o.VirtualmachineCount
	}
	if o.VlanCount != nil {
		toSerialize["vlan_count"] = o.VlanCount
	}
	return json.Marshal(toSerialize)
}

type NullableSite struct {
	value *Site
	isSet bool
}

func (v NullableSite) Get() *Site {
	return v.value
}

func (v *NullableSite) Set(val *Site) {
	v.value = val
	v.isSet = true
}

func (v NullableSite) IsSet() bool {
	return v.isSet
}

func (v *NullableSite) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSite(val *Site) *NullableSite {
	return &NullableSite{value: val, isSet: true}
}

func (v NullableSite) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSite) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


