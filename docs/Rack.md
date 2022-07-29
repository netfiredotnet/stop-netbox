# Rack

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Display** | Pointer to **string** |  | [optional] [readonly] 
**Name** | **string** |  | 
**FacilityId** | Pointer to **NullableString** | Locally-assigned identifier | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] [readonly] 
**Site** | [**NestedSite**](NestedSite.md) |  | 
**Location** | Pointer to [**NullableNestedLocation**](NestedLocation.md) |  | [optional] 
**Tenant** | Pointer to [**NullableNestedTenant**](NestedTenant.md) |  | [optional] 
**Status** | Pointer to [**Status4**](Status4.md) |  | [optional] 
**Role** | Pointer to [**NullableNestedRackRole**](NestedRackRole.md) |  | [optional] 
**Serial** | Pointer to **string** |  | [optional] 
**AssetTag** | Pointer to **NullableString** | A unique tag used to identify this rack | [optional] 
**Type** | Pointer to [**Type6**](Type6.md) |  | [optional] 
**Width** | Pointer to [**Width**](Width.md) |  | [optional] 
**UHeight** | Pointer to **int32** | Height in rack units | [optional] 
**DescUnits** | Pointer to **bool** | Units are numbered top-to-bottom | [optional] 
**OuterWidth** | Pointer to **NullableInt32** | Outer dimension of rack (width) | [optional] 
**OuterDepth** | Pointer to **NullableInt32** | Outer dimension of rack (depth) | [optional] 
**OuterUnit** | Pointer to [**OuterUnit**](OuterUnit.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTag**](NestedTag.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] [readonly] 
**DeviceCount** | Pointer to **int32** |  | [optional] [readonly] 
**PowerfeedCount** | Pointer to **int32** |  | [optional] [readonly] 

## Methods

### NewRack

`func NewRack(name string, site NestedSite, ) *Rack`

NewRack instantiates a new Rack object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRackWithDefaults

`func NewRackWithDefaults() *Rack`

NewRackWithDefaults instantiates a new Rack object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Rack) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Rack) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Rack) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Rack) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *Rack) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Rack) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Rack) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Rack) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDisplay

`func (o *Rack) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *Rack) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *Rack) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *Rack) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetName

`func (o *Rack) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Rack) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Rack) SetName(v string)`

SetName sets Name field to given value.


### GetFacilityId

`func (o *Rack) GetFacilityId() string`

GetFacilityId returns the FacilityId field if non-nil, zero value otherwise.

### GetFacilityIdOk

`func (o *Rack) GetFacilityIdOk() (*string, bool)`

GetFacilityIdOk returns a tuple with the FacilityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityId

`func (o *Rack) SetFacilityId(v string)`

SetFacilityId sets FacilityId field to given value.

### HasFacilityId

`func (o *Rack) HasFacilityId() bool`

HasFacilityId returns a boolean if a field has been set.

### SetFacilityIdNil

`func (o *Rack) SetFacilityIdNil(b bool)`

 SetFacilityIdNil sets the value for FacilityId to be an explicit nil

### UnsetFacilityId
`func (o *Rack) UnsetFacilityId()`

UnsetFacilityId ensures that no value is present for FacilityId, not even an explicit nil
### GetDisplayName

`func (o *Rack) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Rack) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Rack) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Rack) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetSite

`func (o *Rack) GetSite() NestedSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *Rack) GetSiteOk() (*NestedSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *Rack) SetSite(v NestedSite)`

SetSite sets Site field to given value.


### GetLocation

`func (o *Rack) GetLocation() NestedLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Rack) GetLocationOk() (*NestedLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Rack) SetLocation(v NestedLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Rack) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *Rack) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *Rack) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetTenant

`func (o *Rack) GetTenant() NestedTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *Rack) GetTenantOk() (*NestedTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *Rack) SetTenant(v NestedTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *Rack) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *Rack) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *Rack) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetStatus

`func (o *Rack) GetStatus() Status4`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Rack) GetStatusOk() (*Status4, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Rack) SetStatus(v Status4)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Rack) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRole

`func (o *Rack) GetRole() NestedRackRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *Rack) GetRoleOk() (*NestedRackRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *Rack) SetRole(v NestedRackRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *Rack) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *Rack) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *Rack) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetSerial

`func (o *Rack) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *Rack) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *Rack) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *Rack) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetAssetTag

`func (o *Rack) GetAssetTag() string`

GetAssetTag returns the AssetTag field if non-nil, zero value otherwise.

### GetAssetTagOk

`func (o *Rack) GetAssetTagOk() (*string, bool)`

GetAssetTagOk returns a tuple with the AssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTag

`func (o *Rack) SetAssetTag(v string)`

SetAssetTag sets AssetTag field to given value.

### HasAssetTag

`func (o *Rack) HasAssetTag() bool`

HasAssetTag returns a boolean if a field has been set.

### SetAssetTagNil

`func (o *Rack) SetAssetTagNil(b bool)`

 SetAssetTagNil sets the value for AssetTag to be an explicit nil

### UnsetAssetTag
`func (o *Rack) UnsetAssetTag()`

UnsetAssetTag ensures that no value is present for AssetTag, not even an explicit nil
### GetType

`func (o *Rack) GetType() Type6`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Rack) GetTypeOk() (*Type6, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Rack) SetType(v Type6)`

SetType sets Type field to given value.

### HasType

`func (o *Rack) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWidth

`func (o *Rack) GetWidth() Width`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *Rack) GetWidthOk() (*Width, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *Rack) SetWidth(v Width)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *Rack) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetUHeight

`func (o *Rack) GetUHeight() int32`

GetUHeight returns the UHeight field if non-nil, zero value otherwise.

### GetUHeightOk

`func (o *Rack) GetUHeightOk() (*int32, bool)`

GetUHeightOk returns a tuple with the UHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUHeight

`func (o *Rack) SetUHeight(v int32)`

SetUHeight sets UHeight field to given value.

### HasUHeight

`func (o *Rack) HasUHeight() bool`

HasUHeight returns a boolean if a field has been set.

### GetDescUnits

`func (o *Rack) GetDescUnits() bool`

GetDescUnits returns the DescUnits field if non-nil, zero value otherwise.

### GetDescUnitsOk

`func (o *Rack) GetDescUnitsOk() (*bool, bool)`

GetDescUnitsOk returns a tuple with the DescUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescUnits

`func (o *Rack) SetDescUnits(v bool)`

SetDescUnits sets DescUnits field to given value.

### HasDescUnits

`func (o *Rack) HasDescUnits() bool`

HasDescUnits returns a boolean if a field has been set.

### GetOuterWidth

`func (o *Rack) GetOuterWidth() int32`

GetOuterWidth returns the OuterWidth field if non-nil, zero value otherwise.

### GetOuterWidthOk

`func (o *Rack) GetOuterWidthOk() (*int32, bool)`

GetOuterWidthOk returns a tuple with the OuterWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOuterWidth

`func (o *Rack) SetOuterWidth(v int32)`

SetOuterWidth sets OuterWidth field to given value.

### HasOuterWidth

`func (o *Rack) HasOuterWidth() bool`

HasOuterWidth returns a boolean if a field has been set.

### SetOuterWidthNil

`func (o *Rack) SetOuterWidthNil(b bool)`

 SetOuterWidthNil sets the value for OuterWidth to be an explicit nil

### UnsetOuterWidth
`func (o *Rack) UnsetOuterWidth()`

UnsetOuterWidth ensures that no value is present for OuterWidth, not even an explicit nil
### GetOuterDepth

`func (o *Rack) GetOuterDepth() int32`

GetOuterDepth returns the OuterDepth field if non-nil, zero value otherwise.

### GetOuterDepthOk

`func (o *Rack) GetOuterDepthOk() (*int32, bool)`

GetOuterDepthOk returns a tuple with the OuterDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOuterDepth

`func (o *Rack) SetOuterDepth(v int32)`

SetOuterDepth sets OuterDepth field to given value.

### HasOuterDepth

`func (o *Rack) HasOuterDepth() bool`

HasOuterDepth returns a boolean if a field has been set.

### SetOuterDepthNil

`func (o *Rack) SetOuterDepthNil(b bool)`

 SetOuterDepthNil sets the value for OuterDepth to be an explicit nil

### UnsetOuterDepth
`func (o *Rack) UnsetOuterDepth()`

UnsetOuterDepth ensures that no value is present for OuterDepth, not even an explicit nil
### GetOuterUnit

`func (o *Rack) GetOuterUnit() OuterUnit`

GetOuterUnit returns the OuterUnit field if non-nil, zero value otherwise.

### GetOuterUnitOk

`func (o *Rack) GetOuterUnitOk() (*OuterUnit, bool)`

GetOuterUnitOk returns a tuple with the OuterUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOuterUnit

`func (o *Rack) SetOuterUnit(v OuterUnit)`

SetOuterUnit sets OuterUnit field to given value.

### HasOuterUnit

`func (o *Rack) HasOuterUnit() bool`

HasOuterUnit returns a boolean if a field has been set.

### GetComments

`func (o *Rack) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *Rack) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *Rack) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *Rack) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *Rack) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Rack) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Rack) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Rack) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *Rack) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Rack) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Rack) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Rack) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *Rack) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Rack) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Rack) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Rack) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *Rack) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Rack) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Rack) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *Rack) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetDeviceCount

`func (o *Rack) GetDeviceCount() int32`

GetDeviceCount returns the DeviceCount field if non-nil, zero value otherwise.

### GetDeviceCountOk

`func (o *Rack) GetDeviceCountOk() (*int32, bool)`

GetDeviceCountOk returns a tuple with the DeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceCount

`func (o *Rack) SetDeviceCount(v int32)`

SetDeviceCount sets DeviceCount field to given value.

### HasDeviceCount

`func (o *Rack) HasDeviceCount() bool`

HasDeviceCount returns a boolean if a field has been set.

### GetPowerfeedCount

`func (o *Rack) GetPowerfeedCount() int32`

GetPowerfeedCount returns the PowerfeedCount field if non-nil, zero value otherwise.

### GetPowerfeedCountOk

`func (o *Rack) GetPowerfeedCountOk() (*int32, bool)`

GetPowerfeedCountOk returns a tuple with the PowerfeedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerfeedCount

`func (o *Rack) SetPowerfeedCount(v int32)`

SetPowerfeedCount sets PowerfeedCount field to given value.

### HasPowerfeedCount

`func (o *Rack) HasPowerfeedCount() bool`

HasPowerfeedCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


