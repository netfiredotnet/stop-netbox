# WritableCircuit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Display** | Pointer to **string** |  | [optional] [readonly] 
**Cid** | **string** |  | 
**Provider** | **int32** |  | 
**Type** | **int32** |  | 
**Status** | Pointer to **string** |  | [optional] 
**Tenant** | Pointer to **NullableInt32** |  | [optional] 
**InstallDate** | Pointer to **NullableString** |  | [optional] 
**CommitRate** | Pointer to **NullableInt32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**TerminationA** | Pointer to **int32** |  | [optional] [readonly] 
**TerminationZ** | Pointer to **int32** |  | [optional] [readonly] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTag**](NestedTag.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewWritableCircuit

`func NewWritableCircuit(cid string, provider int32, type_ int32, ) *WritableCircuit`

NewWritableCircuit instantiates a new WritableCircuit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritableCircuitWithDefaults

`func NewWritableCircuitWithDefaults() *WritableCircuit`

NewWritableCircuitWithDefaults instantiates a new WritableCircuit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WritableCircuit) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WritableCircuit) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WritableCircuit) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *WritableCircuit) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *WritableCircuit) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WritableCircuit) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WritableCircuit) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *WritableCircuit) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDisplay

`func (o *WritableCircuit) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *WritableCircuit) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *WritableCircuit) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *WritableCircuit) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetCid

`func (o *WritableCircuit) GetCid() string`

GetCid returns the Cid field if non-nil, zero value otherwise.

### GetCidOk

`func (o *WritableCircuit) GetCidOk() (*string, bool)`

GetCidOk returns a tuple with the Cid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCid

`func (o *WritableCircuit) SetCid(v string)`

SetCid sets Cid field to given value.


### GetProvider

`func (o *WritableCircuit) GetProvider() int32`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *WritableCircuit) GetProviderOk() (*int32, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *WritableCircuit) SetProvider(v int32)`

SetProvider sets Provider field to given value.


### GetType

`func (o *WritableCircuit) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WritableCircuit) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WritableCircuit) SetType(v int32)`

SetType sets Type field to given value.


### GetStatus

`func (o *WritableCircuit) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WritableCircuit) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WritableCircuit) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WritableCircuit) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTenant

`func (o *WritableCircuit) GetTenant() int32`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *WritableCircuit) GetTenantOk() (*int32, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *WritableCircuit) SetTenant(v int32)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *WritableCircuit) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *WritableCircuit) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *WritableCircuit) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetInstallDate

`func (o *WritableCircuit) GetInstallDate() string`

GetInstallDate returns the InstallDate field if non-nil, zero value otherwise.

### GetInstallDateOk

`func (o *WritableCircuit) GetInstallDateOk() (*string, bool)`

GetInstallDateOk returns a tuple with the InstallDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallDate

`func (o *WritableCircuit) SetInstallDate(v string)`

SetInstallDate sets InstallDate field to given value.

### HasInstallDate

`func (o *WritableCircuit) HasInstallDate() bool`

HasInstallDate returns a boolean if a field has been set.

### SetInstallDateNil

`func (o *WritableCircuit) SetInstallDateNil(b bool)`

 SetInstallDateNil sets the value for InstallDate to be an explicit nil

### UnsetInstallDate
`func (o *WritableCircuit) UnsetInstallDate()`

UnsetInstallDate ensures that no value is present for InstallDate, not even an explicit nil
### GetCommitRate

`func (o *WritableCircuit) GetCommitRate() int32`

GetCommitRate returns the CommitRate field if non-nil, zero value otherwise.

### GetCommitRateOk

`func (o *WritableCircuit) GetCommitRateOk() (*int32, bool)`

GetCommitRateOk returns a tuple with the CommitRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitRate

`func (o *WritableCircuit) SetCommitRate(v int32)`

SetCommitRate sets CommitRate field to given value.

### HasCommitRate

`func (o *WritableCircuit) HasCommitRate() bool`

HasCommitRate returns a boolean if a field has been set.

### SetCommitRateNil

`func (o *WritableCircuit) SetCommitRateNil(b bool)`

 SetCommitRateNil sets the value for CommitRate to be an explicit nil

### UnsetCommitRate
`func (o *WritableCircuit) UnsetCommitRate()`

UnsetCommitRate ensures that no value is present for CommitRate, not even an explicit nil
### GetDescription

`func (o *WritableCircuit) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WritableCircuit) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WritableCircuit) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WritableCircuit) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTerminationA

`func (o *WritableCircuit) GetTerminationA() int32`

GetTerminationA returns the TerminationA field if non-nil, zero value otherwise.

### GetTerminationAOk

`func (o *WritableCircuit) GetTerminationAOk() (*int32, bool)`

GetTerminationAOk returns a tuple with the TerminationA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationA

`func (o *WritableCircuit) SetTerminationA(v int32)`

SetTerminationA sets TerminationA field to given value.

### HasTerminationA

`func (o *WritableCircuit) HasTerminationA() bool`

HasTerminationA returns a boolean if a field has been set.

### GetTerminationZ

`func (o *WritableCircuit) GetTerminationZ() int32`

GetTerminationZ returns the TerminationZ field if non-nil, zero value otherwise.

### GetTerminationZOk

`func (o *WritableCircuit) GetTerminationZOk() (*int32, bool)`

GetTerminationZOk returns a tuple with the TerminationZ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationZ

`func (o *WritableCircuit) SetTerminationZ(v int32)`

SetTerminationZ sets TerminationZ field to given value.

### HasTerminationZ

`func (o *WritableCircuit) HasTerminationZ() bool`

HasTerminationZ returns a boolean if a field has been set.

### GetComments

`func (o *WritableCircuit) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *WritableCircuit) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *WritableCircuit) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *WritableCircuit) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *WritableCircuit) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WritableCircuit) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WritableCircuit) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WritableCircuit) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *WritableCircuit) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *WritableCircuit) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *WritableCircuit) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *WritableCircuit) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *WritableCircuit) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *WritableCircuit) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *WritableCircuit) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *WritableCircuit) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *WritableCircuit) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *WritableCircuit) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *WritableCircuit) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *WritableCircuit) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


