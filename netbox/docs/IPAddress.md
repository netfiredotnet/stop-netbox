# IPAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Display** | Pointer to **string** |  | [optional] [readonly] 
**Family** | Pointer to [**Family**](Family.md) |  | [optional] 
**Address** | **string** | IPv4 or IPv6 address (with mask) | 
**Vrf** | Pointer to [**NullableNestedVRF**](NestedVRF.md) |  | [optional] 
**Tenant** | Pointer to [**NullableNestedTenant**](NestedTenant.md) |  | [optional] 
**Status** | Pointer to [**Status7**](Status7.md) |  | [optional] 
**Role** | Pointer to [**Role1**](Role1.md) |  | [optional] 
**AssignedObjectType** | Pointer to **NullableString** |  | [optional] 
**AssignedObjectId** | Pointer to **NullableInt32** |  | [optional] 
**AssignedObject** | Pointer to **map[string]string** |  | [optional] [readonly] 
**NatInside** | Pointer to [**NestedIPAddress**](NestedIPAddress.md) |  | [optional] 
**NatOutside** | Pointer to [**NestedIPAddress**](NestedIPAddress.md) |  | [optional] 
**DnsName** | Pointer to **string** | Hostname or FQDN (not case-sensitive) | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTag**](NestedTag.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewIPAddress

`func NewIPAddress(address string, ) *IPAddress`

NewIPAddress instantiates a new IPAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPAddressWithDefaults

`func NewIPAddressWithDefaults() *IPAddress`

NewIPAddressWithDefaults instantiates a new IPAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IPAddress) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IPAddress) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IPAddress) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *IPAddress) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *IPAddress) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *IPAddress) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *IPAddress) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *IPAddress) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDisplay

`func (o *IPAddress) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *IPAddress) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *IPAddress) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *IPAddress) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetFamily

`func (o *IPAddress) GetFamily() Family`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *IPAddress) GetFamilyOk() (*Family, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *IPAddress) SetFamily(v Family)`

SetFamily sets Family field to given value.

### HasFamily

`func (o *IPAddress) HasFamily() bool`

HasFamily returns a boolean if a field has been set.

### GetAddress

`func (o *IPAddress) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *IPAddress) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *IPAddress) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetVrf

`func (o *IPAddress) GetVrf() NestedVRF`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *IPAddress) GetVrfOk() (*NestedVRF, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *IPAddress) SetVrf(v NestedVRF)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *IPAddress) HasVrf() bool`

HasVrf returns a boolean if a field has been set.

### SetVrfNil

`func (o *IPAddress) SetVrfNil(b bool)`

 SetVrfNil sets the value for Vrf to be an explicit nil

### UnsetVrf
`func (o *IPAddress) UnsetVrf()`

UnsetVrf ensures that no value is present for Vrf, not even an explicit nil
### GetTenant

`func (o *IPAddress) GetTenant() NestedTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *IPAddress) GetTenantOk() (*NestedTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *IPAddress) SetTenant(v NestedTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *IPAddress) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *IPAddress) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *IPAddress) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetStatus

`func (o *IPAddress) GetStatus() Status7`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IPAddress) GetStatusOk() (*Status7, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IPAddress) SetStatus(v Status7)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IPAddress) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRole

`func (o *IPAddress) GetRole() Role1`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *IPAddress) GetRoleOk() (*Role1, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *IPAddress) SetRole(v Role1)`

SetRole sets Role field to given value.

### HasRole

`func (o *IPAddress) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetAssignedObjectType

`func (o *IPAddress) GetAssignedObjectType() string`

GetAssignedObjectType returns the AssignedObjectType field if non-nil, zero value otherwise.

### GetAssignedObjectTypeOk

`func (o *IPAddress) GetAssignedObjectTypeOk() (*string, bool)`

GetAssignedObjectTypeOk returns a tuple with the AssignedObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedObjectType

`func (o *IPAddress) SetAssignedObjectType(v string)`

SetAssignedObjectType sets AssignedObjectType field to given value.

### HasAssignedObjectType

`func (o *IPAddress) HasAssignedObjectType() bool`

HasAssignedObjectType returns a boolean if a field has been set.

### SetAssignedObjectTypeNil

`func (o *IPAddress) SetAssignedObjectTypeNil(b bool)`

 SetAssignedObjectTypeNil sets the value for AssignedObjectType to be an explicit nil

### UnsetAssignedObjectType
`func (o *IPAddress) UnsetAssignedObjectType()`

UnsetAssignedObjectType ensures that no value is present for AssignedObjectType, not even an explicit nil
### GetAssignedObjectId

`func (o *IPAddress) GetAssignedObjectId() int32`

GetAssignedObjectId returns the AssignedObjectId field if non-nil, zero value otherwise.

### GetAssignedObjectIdOk

`func (o *IPAddress) GetAssignedObjectIdOk() (*int32, bool)`

GetAssignedObjectIdOk returns a tuple with the AssignedObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedObjectId

`func (o *IPAddress) SetAssignedObjectId(v int32)`

SetAssignedObjectId sets AssignedObjectId field to given value.

### HasAssignedObjectId

`func (o *IPAddress) HasAssignedObjectId() bool`

HasAssignedObjectId returns a boolean if a field has been set.

### SetAssignedObjectIdNil

`func (o *IPAddress) SetAssignedObjectIdNil(b bool)`

 SetAssignedObjectIdNil sets the value for AssignedObjectId to be an explicit nil

### UnsetAssignedObjectId
`func (o *IPAddress) UnsetAssignedObjectId()`

UnsetAssignedObjectId ensures that no value is present for AssignedObjectId, not even an explicit nil
### GetAssignedObject

`func (o *IPAddress) GetAssignedObject() map[string]string`

GetAssignedObject returns the AssignedObject field if non-nil, zero value otherwise.

### GetAssignedObjectOk

`func (o *IPAddress) GetAssignedObjectOk() (*map[string]string, bool)`

GetAssignedObjectOk returns a tuple with the AssignedObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedObject

`func (o *IPAddress) SetAssignedObject(v map[string]string)`

SetAssignedObject sets AssignedObject field to given value.

### HasAssignedObject

`func (o *IPAddress) HasAssignedObject() bool`

HasAssignedObject returns a boolean if a field has been set.

### GetNatInside

`func (o *IPAddress) GetNatInside() NestedIPAddress`

GetNatInside returns the NatInside field if non-nil, zero value otherwise.

### GetNatInsideOk

`func (o *IPAddress) GetNatInsideOk() (*NestedIPAddress, bool)`

GetNatInsideOk returns a tuple with the NatInside field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatInside

`func (o *IPAddress) SetNatInside(v NestedIPAddress)`

SetNatInside sets NatInside field to given value.

### HasNatInside

`func (o *IPAddress) HasNatInside() bool`

HasNatInside returns a boolean if a field has been set.

### GetNatOutside

`func (o *IPAddress) GetNatOutside() NestedIPAddress`

GetNatOutside returns the NatOutside field if non-nil, zero value otherwise.

### GetNatOutsideOk

`func (o *IPAddress) GetNatOutsideOk() (*NestedIPAddress, bool)`

GetNatOutsideOk returns a tuple with the NatOutside field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatOutside

`func (o *IPAddress) SetNatOutside(v NestedIPAddress)`

SetNatOutside sets NatOutside field to given value.

### HasNatOutside

`func (o *IPAddress) HasNatOutside() bool`

HasNatOutside returns a boolean if a field has been set.

### GetDnsName

`func (o *IPAddress) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *IPAddress) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *IPAddress) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *IPAddress) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### GetDescription

`func (o *IPAddress) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IPAddress) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IPAddress) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IPAddress) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *IPAddress) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IPAddress) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IPAddress) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *IPAddress) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *IPAddress) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *IPAddress) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *IPAddress) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *IPAddress) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *IPAddress) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *IPAddress) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *IPAddress) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *IPAddress) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *IPAddress) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *IPAddress) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *IPAddress) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *IPAddress) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


