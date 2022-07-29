# Secret

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Display** | Pointer to **string** |  | [optional] [readonly] 
**AssignedObjectType** | **string** |  | 
**AssignedObjectId** | **int32** |  | 
**AssignedObject** | Pointer to **map[string]string** |  | [optional] [readonly] 
**Role** | [**NestedSecretRole**](NestedSecretRole.md) |  | 
**Name** | Pointer to **string** |  | [optional] 
**Plaintext** | **string** |  | 
**Hash** | Pointer to **string** |  | [optional] [readonly] 
**Tags** | Pointer to [**[]NestedTag**](NestedTag.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewSecret

`func NewSecret(assignedObjectType string, assignedObjectId int32, role NestedSecretRole, plaintext string, ) *Secret`

NewSecret instantiates a new Secret object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretWithDefaults

`func NewSecretWithDefaults() *Secret`

NewSecretWithDefaults instantiates a new Secret object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Secret) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Secret) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Secret) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Secret) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *Secret) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Secret) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Secret) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Secret) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDisplay

`func (o *Secret) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *Secret) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *Secret) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *Secret) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetAssignedObjectType

`func (o *Secret) GetAssignedObjectType() string`

GetAssignedObjectType returns the AssignedObjectType field if non-nil, zero value otherwise.

### GetAssignedObjectTypeOk

`func (o *Secret) GetAssignedObjectTypeOk() (*string, bool)`

GetAssignedObjectTypeOk returns a tuple with the AssignedObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedObjectType

`func (o *Secret) SetAssignedObjectType(v string)`

SetAssignedObjectType sets AssignedObjectType field to given value.


### GetAssignedObjectId

`func (o *Secret) GetAssignedObjectId() int32`

GetAssignedObjectId returns the AssignedObjectId field if non-nil, zero value otherwise.

### GetAssignedObjectIdOk

`func (o *Secret) GetAssignedObjectIdOk() (*int32, bool)`

GetAssignedObjectIdOk returns a tuple with the AssignedObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedObjectId

`func (o *Secret) SetAssignedObjectId(v int32)`

SetAssignedObjectId sets AssignedObjectId field to given value.


### GetAssignedObject

`func (o *Secret) GetAssignedObject() map[string]string`

GetAssignedObject returns the AssignedObject field if non-nil, zero value otherwise.

### GetAssignedObjectOk

`func (o *Secret) GetAssignedObjectOk() (*map[string]string, bool)`

GetAssignedObjectOk returns a tuple with the AssignedObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedObject

`func (o *Secret) SetAssignedObject(v map[string]string)`

SetAssignedObject sets AssignedObject field to given value.

### HasAssignedObject

`func (o *Secret) HasAssignedObject() bool`

HasAssignedObject returns a boolean if a field has been set.

### GetRole

`func (o *Secret) GetRole() NestedSecretRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *Secret) GetRoleOk() (*NestedSecretRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *Secret) SetRole(v NestedSecretRole)`

SetRole sets Role field to given value.


### GetName

`func (o *Secret) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Secret) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Secret) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Secret) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlaintext

`func (o *Secret) GetPlaintext() string`

GetPlaintext returns the Plaintext field if non-nil, zero value otherwise.

### GetPlaintextOk

`func (o *Secret) GetPlaintextOk() (*string, bool)`

GetPlaintextOk returns a tuple with the Plaintext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaintext

`func (o *Secret) SetPlaintext(v string)`

SetPlaintext sets Plaintext field to given value.


### GetHash

`func (o *Secret) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *Secret) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *Secret) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *Secret) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetTags

`func (o *Secret) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Secret) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Secret) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Secret) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *Secret) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Secret) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Secret) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Secret) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *Secret) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Secret) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Secret) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Secret) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *Secret) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Secret) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Secret) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *Secret) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


