# Provider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Display** | Pointer to **string** |  | [optional] [readonly] 
**Name** | **string** |  | 
**Slug** | **string** |  | 
**Asn** | Pointer to **NullableInt32** | 32-bit autonomous system number | [optional] 
**Account** | Pointer to **string** |  | [optional] 
**PortalUrl** | Pointer to **string** |  | [optional] 
**NocContact** | Pointer to **string** |  | [optional] 
**AdminContact** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTag**](NestedTag.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] [readonly] 
**CircuitCount** | Pointer to **int32** |  | [optional] [readonly] 

## Methods

### NewProvider

`func NewProvider(name string, slug string, ) *Provider`

NewProvider instantiates a new Provider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderWithDefaults

`func NewProviderWithDefaults() *Provider`

NewProviderWithDefaults instantiates a new Provider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Provider) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Provider) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Provider) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Provider) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *Provider) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Provider) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Provider) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Provider) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDisplay

`func (o *Provider) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *Provider) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *Provider) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *Provider) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetName

`func (o *Provider) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Provider) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Provider) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *Provider) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *Provider) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *Provider) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetAsn

`func (o *Provider) GetAsn() int32`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *Provider) GetAsnOk() (*int32, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *Provider) SetAsn(v int32)`

SetAsn sets Asn field to given value.

### HasAsn

`func (o *Provider) HasAsn() bool`

HasAsn returns a boolean if a field has been set.

### SetAsnNil

`func (o *Provider) SetAsnNil(b bool)`

 SetAsnNil sets the value for Asn to be an explicit nil

### UnsetAsn
`func (o *Provider) UnsetAsn()`

UnsetAsn ensures that no value is present for Asn, not even an explicit nil
### GetAccount

`func (o *Provider) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *Provider) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *Provider) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *Provider) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetPortalUrl

`func (o *Provider) GetPortalUrl() string`

GetPortalUrl returns the PortalUrl field if non-nil, zero value otherwise.

### GetPortalUrlOk

`func (o *Provider) GetPortalUrlOk() (*string, bool)`

GetPortalUrlOk returns a tuple with the PortalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalUrl

`func (o *Provider) SetPortalUrl(v string)`

SetPortalUrl sets PortalUrl field to given value.

### HasPortalUrl

`func (o *Provider) HasPortalUrl() bool`

HasPortalUrl returns a boolean if a field has been set.

### GetNocContact

`func (o *Provider) GetNocContact() string`

GetNocContact returns the NocContact field if non-nil, zero value otherwise.

### GetNocContactOk

`func (o *Provider) GetNocContactOk() (*string, bool)`

GetNocContactOk returns a tuple with the NocContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNocContact

`func (o *Provider) SetNocContact(v string)`

SetNocContact sets NocContact field to given value.

### HasNocContact

`func (o *Provider) HasNocContact() bool`

HasNocContact returns a boolean if a field has been set.

### GetAdminContact

`func (o *Provider) GetAdminContact() string`

GetAdminContact returns the AdminContact field if non-nil, zero value otherwise.

### GetAdminContactOk

`func (o *Provider) GetAdminContactOk() (*string, bool)`

GetAdminContactOk returns a tuple with the AdminContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminContact

`func (o *Provider) SetAdminContact(v string)`

SetAdminContact sets AdminContact field to given value.

### HasAdminContact

`func (o *Provider) HasAdminContact() bool`

HasAdminContact returns a boolean if a field has been set.

### GetComments

`func (o *Provider) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *Provider) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *Provider) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *Provider) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *Provider) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Provider) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Provider) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Provider) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *Provider) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Provider) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Provider) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Provider) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *Provider) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Provider) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Provider) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Provider) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *Provider) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Provider) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Provider) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *Provider) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetCircuitCount

`func (o *Provider) GetCircuitCount() int32`

GetCircuitCount returns the CircuitCount field if non-nil, zero value otherwise.

### GetCircuitCountOk

`func (o *Provider) GetCircuitCountOk() (*int32, bool)`

GetCircuitCountOk returns a tuple with the CircuitCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitCount

`func (o *Provider) SetCircuitCount(v int32)`

SetCircuitCount sets CircuitCount field to given value.

### HasCircuitCount

`func (o *Provider) HasCircuitCount() bool`

HasCircuitCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


