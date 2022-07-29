# Circuit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Display** | Pointer to **string** |  | [optional] [readonly] 
**Cid** | **string** |  | 
**Provider** | [**NestedProvider**](NestedProvider.md) |  | 
**Type** | [**NestedCircuitType**](NestedCircuitType.md) |  | 
**Status** | Pointer to [**Status**](Status.md) |  | [optional] 
**Tenant** | Pointer to [**NullableNestedTenant**](NestedTenant.md) |  | [optional] 
**InstallDate** | Pointer to **NullableString** |  | [optional] 
**CommitRate** | Pointer to **NullableInt32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**TerminationA** | Pointer to [**CircuitCircuitTermination**](CircuitCircuitTermination.md) |  | [optional] 
**TerminationZ** | Pointer to [**CircuitCircuitTermination**](CircuitCircuitTermination.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTag**](NestedTag.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewCircuit

`func NewCircuit(cid string, provider NestedProvider, type_ NestedCircuitType, ) *Circuit`

NewCircuit instantiates a new Circuit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCircuitWithDefaults

`func NewCircuitWithDefaults() *Circuit`

NewCircuitWithDefaults instantiates a new Circuit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Circuit) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Circuit) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Circuit) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Circuit) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *Circuit) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Circuit) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Circuit) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Circuit) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDisplay

`func (o *Circuit) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *Circuit) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *Circuit) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *Circuit) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetCid

`func (o *Circuit) GetCid() string`

GetCid returns the Cid field if non-nil, zero value otherwise.

### GetCidOk

`func (o *Circuit) GetCidOk() (*string, bool)`

GetCidOk returns a tuple with the Cid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCid

`func (o *Circuit) SetCid(v string)`

SetCid sets Cid field to given value.


### GetProvider

`func (o *Circuit) GetProvider() NestedProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *Circuit) GetProviderOk() (*NestedProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *Circuit) SetProvider(v NestedProvider)`

SetProvider sets Provider field to given value.


### GetType

`func (o *Circuit) GetType() NestedCircuitType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Circuit) GetTypeOk() (*NestedCircuitType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Circuit) SetType(v NestedCircuitType)`

SetType sets Type field to given value.


### GetStatus

`func (o *Circuit) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Circuit) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Circuit) SetStatus(v Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Circuit) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTenant

`func (o *Circuit) GetTenant() NestedTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *Circuit) GetTenantOk() (*NestedTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *Circuit) SetTenant(v NestedTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *Circuit) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *Circuit) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *Circuit) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetInstallDate

`func (o *Circuit) GetInstallDate() string`

GetInstallDate returns the InstallDate field if non-nil, zero value otherwise.

### GetInstallDateOk

`func (o *Circuit) GetInstallDateOk() (*string, bool)`

GetInstallDateOk returns a tuple with the InstallDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallDate

`func (o *Circuit) SetInstallDate(v string)`

SetInstallDate sets InstallDate field to given value.

### HasInstallDate

`func (o *Circuit) HasInstallDate() bool`

HasInstallDate returns a boolean if a field has been set.

### SetInstallDateNil

`func (o *Circuit) SetInstallDateNil(b bool)`

 SetInstallDateNil sets the value for InstallDate to be an explicit nil

### UnsetInstallDate
`func (o *Circuit) UnsetInstallDate()`

UnsetInstallDate ensures that no value is present for InstallDate, not even an explicit nil
### GetCommitRate

`func (o *Circuit) GetCommitRate() int32`

GetCommitRate returns the CommitRate field if non-nil, zero value otherwise.

### GetCommitRateOk

`func (o *Circuit) GetCommitRateOk() (*int32, bool)`

GetCommitRateOk returns a tuple with the CommitRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitRate

`func (o *Circuit) SetCommitRate(v int32)`

SetCommitRate sets CommitRate field to given value.

### HasCommitRate

`func (o *Circuit) HasCommitRate() bool`

HasCommitRate returns a boolean if a field has been set.

### SetCommitRateNil

`func (o *Circuit) SetCommitRateNil(b bool)`

 SetCommitRateNil sets the value for CommitRate to be an explicit nil

### UnsetCommitRate
`func (o *Circuit) UnsetCommitRate()`

UnsetCommitRate ensures that no value is present for CommitRate, not even an explicit nil
### GetDescription

`func (o *Circuit) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Circuit) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Circuit) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Circuit) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTerminationA

`func (o *Circuit) GetTerminationA() CircuitCircuitTermination`

GetTerminationA returns the TerminationA field if non-nil, zero value otherwise.

### GetTerminationAOk

`func (o *Circuit) GetTerminationAOk() (*CircuitCircuitTermination, bool)`

GetTerminationAOk returns a tuple with the TerminationA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationA

`func (o *Circuit) SetTerminationA(v CircuitCircuitTermination)`

SetTerminationA sets TerminationA field to given value.

### HasTerminationA

`func (o *Circuit) HasTerminationA() bool`

HasTerminationA returns a boolean if a field has been set.

### GetTerminationZ

`func (o *Circuit) GetTerminationZ() CircuitCircuitTermination`

GetTerminationZ returns the TerminationZ field if non-nil, zero value otherwise.

### GetTerminationZOk

`func (o *Circuit) GetTerminationZOk() (*CircuitCircuitTermination, bool)`

GetTerminationZOk returns a tuple with the TerminationZ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationZ

`func (o *Circuit) SetTerminationZ(v CircuitCircuitTermination)`

SetTerminationZ sets TerminationZ field to given value.

### HasTerminationZ

`func (o *Circuit) HasTerminationZ() bool`

HasTerminationZ returns a boolean if a field has been set.

### GetComments

`func (o *Circuit) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *Circuit) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *Circuit) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *Circuit) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *Circuit) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Circuit) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Circuit) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Circuit) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *Circuit) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Circuit) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Circuit) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Circuit) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *Circuit) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Circuit) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Circuit) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Circuit) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *Circuit) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Circuit) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Circuit) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *Circuit) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


