# Aggregate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Display** | Pointer to **string** |  | [optional] [readonly] 
**Family** | Pointer to [**Family**](Family.md) |  | [optional] 
**Prefix** | **string** |  | 
**Rir** | [**NestedRIR**](NestedRIR.md) |  | 
**Tenant** | Pointer to [**NullableNestedTenant**](NestedTenant.md) |  | [optional] 
**DateAdded** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTag**](NestedTag.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewAggregate

`func NewAggregate(prefix string, rir NestedRIR, ) *Aggregate`

NewAggregate instantiates a new Aggregate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregateWithDefaults

`func NewAggregateWithDefaults() *Aggregate`

NewAggregateWithDefaults instantiates a new Aggregate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Aggregate) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Aggregate) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Aggregate) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Aggregate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *Aggregate) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Aggregate) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Aggregate) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Aggregate) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDisplay

`func (o *Aggregate) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *Aggregate) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *Aggregate) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *Aggregate) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetFamily

`func (o *Aggregate) GetFamily() Family`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *Aggregate) GetFamilyOk() (*Family, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *Aggregate) SetFamily(v Family)`

SetFamily sets Family field to given value.

### HasFamily

`func (o *Aggregate) HasFamily() bool`

HasFamily returns a boolean if a field has been set.

### GetPrefix

`func (o *Aggregate) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *Aggregate) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *Aggregate) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.


### GetRir

`func (o *Aggregate) GetRir() NestedRIR`

GetRir returns the Rir field if non-nil, zero value otherwise.

### GetRirOk

`func (o *Aggregate) GetRirOk() (*NestedRIR, bool)`

GetRirOk returns a tuple with the Rir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRir

`func (o *Aggregate) SetRir(v NestedRIR)`

SetRir sets Rir field to given value.


### GetTenant

`func (o *Aggregate) GetTenant() NestedTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *Aggregate) GetTenantOk() (*NestedTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *Aggregate) SetTenant(v NestedTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *Aggregate) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *Aggregate) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *Aggregate) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetDateAdded

`func (o *Aggregate) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *Aggregate) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *Aggregate) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *Aggregate) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### SetDateAddedNil

`func (o *Aggregate) SetDateAddedNil(b bool)`

 SetDateAddedNil sets the value for DateAdded to be an explicit nil

### UnsetDateAdded
`func (o *Aggregate) UnsetDateAdded()`

UnsetDateAdded ensures that no value is present for DateAdded, not even an explicit nil
### GetDescription

`func (o *Aggregate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Aggregate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Aggregate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Aggregate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *Aggregate) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Aggregate) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Aggregate) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Aggregate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *Aggregate) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Aggregate) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Aggregate) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Aggregate) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *Aggregate) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Aggregate) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Aggregate) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Aggregate) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *Aggregate) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Aggregate) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Aggregate) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *Aggregate) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


