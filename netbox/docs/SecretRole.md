# SecretRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Display** | Pointer to **string** |  | [optional] [readonly] 
**Name** | **string** |  | 
**Slug** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] [readonly] 
**SecretCount** | Pointer to **int32** |  | [optional] [readonly] 

## Methods

### NewSecretRole

`func NewSecretRole(name string, slug string, ) *SecretRole`

NewSecretRole instantiates a new SecretRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretRoleWithDefaults

`func NewSecretRoleWithDefaults() *SecretRole`

NewSecretRoleWithDefaults instantiates a new SecretRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SecretRole) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecretRole) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecretRole) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SecretRole) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *SecretRole) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SecretRole) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SecretRole) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SecretRole) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDisplay

`func (o *SecretRole) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *SecretRole) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *SecretRole) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *SecretRole) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetName

`func (o *SecretRole) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecretRole) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecretRole) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *SecretRole) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *SecretRole) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *SecretRole) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetDescription

`func (o *SecretRole) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SecretRole) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SecretRole) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SecretRole) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCustomFields

`func (o *SecretRole) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *SecretRole) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *SecretRole) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *SecretRole) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *SecretRole) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *SecretRole) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *SecretRole) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *SecretRole) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *SecretRole) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *SecretRole) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *SecretRole) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *SecretRole) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetSecretCount

`func (o *SecretRole) GetSecretCount() int32`

GetSecretCount returns the SecretCount field if non-nil, zero value otherwise.

### GetSecretCountOk

`func (o *SecretRole) GetSecretCountOk() (*int32, bool)`

GetSecretCountOk returns a tuple with the SecretCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretCount

`func (o *SecretRole) SetSecretCount(v int32)`

SetSecretCount sets SecretCount field to given value.

### HasSecretCount

`func (o *SecretRole) HasSecretCount() bool`

HasSecretCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


