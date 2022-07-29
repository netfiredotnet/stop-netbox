# ConfigContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Display** | Pointer to **string** |  | [optional] [readonly] 
**Name** | **string** |  | 
**Weight** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**Regions** | Pointer to [**[]NestedRegion**](NestedRegion.md) |  | [optional] 
**SiteGroups** | Pointer to [**[]NestedSiteGroup**](NestedSiteGroup.md) |  | [optional] 
**Sites** | Pointer to [**[]NestedSite**](NestedSite.md) |  | [optional] 
**DeviceTypes** | Pointer to [**[]NestedDeviceType**](NestedDeviceType.md) |  | [optional] 
**Roles** | Pointer to [**[]NestedDeviceRole**](NestedDeviceRole.md) |  | [optional] 
**Platforms** | Pointer to [**[]NestedPlatform**](NestedPlatform.md) |  | [optional] 
**ClusterGroups** | Pointer to [**[]NestedClusterGroup**](NestedClusterGroup.md) |  | [optional] 
**Clusters** | Pointer to [**[]NestedCluster**](NestedCluster.md) |  | [optional] 
**TenantGroups** | Pointer to [**[]NestedTenantGroup**](NestedTenantGroup.md) |  | [optional] 
**Tenants** | Pointer to [**[]NestedTenant**](NestedTenant.md) |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Data** | **string** |  | 
**Created** | Pointer to **string** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewConfigContext

`func NewConfigContext(name string, data string, ) *ConfigContext`

NewConfigContext instantiates a new ConfigContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigContextWithDefaults

`func NewConfigContextWithDefaults() *ConfigContext`

NewConfigContextWithDefaults instantiates a new ConfigContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConfigContext) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConfigContext) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConfigContext) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ConfigContext) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *ConfigContext) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ConfigContext) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ConfigContext) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ConfigContext) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDisplay

`func (o *ConfigContext) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *ConfigContext) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *ConfigContext) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *ConfigContext) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetName

`func (o *ConfigContext) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigContext) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigContext) SetName(v string)`

SetName sets Name field to given value.


### GetWeight

`func (o *ConfigContext) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *ConfigContext) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *ConfigContext) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *ConfigContext) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetDescription

`func (o *ConfigContext) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConfigContext) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConfigContext) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConfigContext) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsActive

`func (o *ConfigContext) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *ConfigContext) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *ConfigContext) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *ConfigContext) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetRegions

`func (o *ConfigContext) GetRegions() []NestedRegion`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *ConfigContext) GetRegionsOk() (*[]NestedRegion, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *ConfigContext) SetRegions(v []NestedRegion)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *ConfigContext) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetSiteGroups

`func (o *ConfigContext) GetSiteGroups() []NestedSiteGroup`

GetSiteGroups returns the SiteGroups field if non-nil, zero value otherwise.

### GetSiteGroupsOk

`func (o *ConfigContext) GetSiteGroupsOk() (*[]NestedSiteGroup, bool)`

GetSiteGroupsOk returns a tuple with the SiteGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteGroups

`func (o *ConfigContext) SetSiteGroups(v []NestedSiteGroup)`

SetSiteGroups sets SiteGroups field to given value.

### HasSiteGroups

`func (o *ConfigContext) HasSiteGroups() bool`

HasSiteGroups returns a boolean if a field has been set.

### GetSites

`func (o *ConfigContext) GetSites() []NestedSite`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *ConfigContext) GetSitesOk() (*[]NestedSite, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *ConfigContext) SetSites(v []NestedSite)`

SetSites sets Sites field to given value.

### HasSites

`func (o *ConfigContext) HasSites() bool`

HasSites returns a boolean if a field has been set.

### GetDeviceTypes

`func (o *ConfigContext) GetDeviceTypes() []NestedDeviceType`

GetDeviceTypes returns the DeviceTypes field if non-nil, zero value otherwise.

### GetDeviceTypesOk

`func (o *ConfigContext) GetDeviceTypesOk() (*[]NestedDeviceType, bool)`

GetDeviceTypesOk returns a tuple with the DeviceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceTypes

`func (o *ConfigContext) SetDeviceTypes(v []NestedDeviceType)`

SetDeviceTypes sets DeviceTypes field to given value.

### HasDeviceTypes

`func (o *ConfigContext) HasDeviceTypes() bool`

HasDeviceTypes returns a boolean if a field has been set.

### GetRoles

`func (o *ConfigContext) GetRoles() []NestedDeviceRole`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *ConfigContext) GetRolesOk() (*[]NestedDeviceRole, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *ConfigContext) SetRoles(v []NestedDeviceRole)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *ConfigContext) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetPlatforms

`func (o *ConfigContext) GetPlatforms() []NestedPlatform`

GetPlatforms returns the Platforms field if non-nil, zero value otherwise.

### GetPlatformsOk

`func (o *ConfigContext) GetPlatformsOk() (*[]NestedPlatform, bool)`

GetPlatformsOk returns a tuple with the Platforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatforms

`func (o *ConfigContext) SetPlatforms(v []NestedPlatform)`

SetPlatforms sets Platforms field to given value.

### HasPlatforms

`func (o *ConfigContext) HasPlatforms() bool`

HasPlatforms returns a boolean if a field has been set.

### GetClusterGroups

`func (o *ConfigContext) GetClusterGroups() []NestedClusterGroup`

GetClusterGroups returns the ClusterGroups field if non-nil, zero value otherwise.

### GetClusterGroupsOk

`func (o *ConfigContext) GetClusterGroupsOk() (*[]NestedClusterGroup, bool)`

GetClusterGroupsOk returns a tuple with the ClusterGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterGroups

`func (o *ConfigContext) SetClusterGroups(v []NestedClusterGroup)`

SetClusterGroups sets ClusterGroups field to given value.

### HasClusterGroups

`func (o *ConfigContext) HasClusterGroups() bool`

HasClusterGroups returns a boolean if a field has been set.

### GetClusters

`func (o *ConfigContext) GetClusters() []NestedCluster`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *ConfigContext) GetClustersOk() (*[]NestedCluster, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *ConfigContext) SetClusters(v []NestedCluster)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *ConfigContext) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

### GetTenantGroups

`func (o *ConfigContext) GetTenantGroups() []NestedTenantGroup`

GetTenantGroups returns the TenantGroups field if non-nil, zero value otherwise.

### GetTenantGroupsOk

`func (o *ConfigContext) GetTenantGroupsOk() (*[]NestedTenantGroup, bool)`

GetTenantGroupsOk returns a tuple with the TenantGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantGroups

`func (o *ConfigContext) SetTenantGroups(v []NestedTenantGroup)`

SetTenantGroups sets TenantGroups field to given value.

### HasTenantGroups

`func (o *ConfigContext) HasTenantGroups() bool`

HasTenantGroups returns a boolean if a field has been set.

### GetTenants

`func (o *ConfigContext) GetTenants() []NestedTenant`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *ConfigContext) GetTenantsOk() (*[]NestedTenant, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *ConfigContext) SetTenants(v []NestedTenant)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *ConfigContext) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetTags

`func (o *ConfigContext) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ConfigContext) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ConfigContext) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ConfigContext) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetData

`func (o *ConfigContext) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ConfigContext) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ConfigContext) SetData(v string)`

SetData sets Data field to given value.


### GetCreated

`func (o *ConfigContext) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ConfigContext) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ConfigContext) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ConfigContext) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ConfigContext) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ConfigContext) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ConfigContext) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ConfigContext) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


