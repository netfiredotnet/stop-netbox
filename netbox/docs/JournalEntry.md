# JournalEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Display** | Pointer to **string** |  | [optional] [readonly] 
**AssignedObjectType** | **string** |  | 
**AssignedObjectId** | **int32** |  | 
**AssignedObject** | Pointer to **map[string]string** |  | [optional] [readonly] 
**Created** | Pointer to **time.Time** |  | [optional] [readonly] 
**CreatedBy** | Pointer to **NullableInt32** |  | [optional] 
**Kind** | Pointer to [**Kind**](Kind.md) |  | [optional] 
**Comments** | **string** |  | 

## Methods

### NewJournalEntry

`func NewJournalEntry(assignedObjectType string, assignedObjectId int32, comments string, ) *JournalEntry`

NewJournalEntry instantiates a new JournalEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJournalEntryWithDefaults

`func NewJournalEntryWithDefaults() *JournalEntry`

NewJournalEntryWithDefaults instantiates a new JournalEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JournalEntry) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JournalEntry) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JournalEntry) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *JournalEntry) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *JournalEntry) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *JournalEntry) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *JournalEntry) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *JournalEntry) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDisplay

`func (o *JournalEntry) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *JournalEntry) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *JournalEntry) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *JournalEntry) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetAssignedObjectType

`func (o *JournalEntry) GetAssignedObjectType() string`

GetAssignedObjectType returns the AssignedObjectType field if non-nil, zero value otherwise.

### GetAssignedObjectTypeOk

`func (o *JournalEntry) GetAssignedObjectTypeOk() (*string, bool)`

GetAssignedObjectTypeOk returns a tuple with the AssignedObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedObjectType

`func (o *JournalEntry) SetAssignedObjectType(v string)`

SetAssignedObjectType sets AssignedObjectType field to given value.


### GetAssignedObjectId

`func (o *JournalEntry) GetAssignedObjectId() int32`

GetAssignedObjectId returns the AssignedObjectId field if non-nil, zero value otherwise.

### GetAssignedObjectIdOk

`func (o *JournalEntry) GetAssignedObjectIdOk() (*int32, bool)`

GetAssignedObjectIdOk returns a tuple with the AssignedObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedObjectId

`func (o *JournalEntry) SetAssignedObjectId(v int32)`

SetAssignedObjectId sets AssignedObjectId field to given value.


### GetAssignedObject

`func (o *JournalEntry) GetAssignedObject() map[string]string`

GetAssignedObject returns the AssignedObject field if non-nil, zero value otherwise.

### GetAssignedObjectOk

`func (o *JournalEntry) GetAssignedObjectOk() (*map[string]string, bool)`

GetAssignedObjectOk returns a tuple with the AssignedObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedObject

`func (o *JournalEntry) SetAssignedObject(v map[string]string)`

SetAssignedObject sets AssignedObject field to given value.

### HasAssignedObject

`func (o *JournalEntry) HasAssignedObject() bool`

HasAssignedObject returns a boolean if a field has been set.

### GetCreated

`func (o *JournalEntry) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *JournalEntry) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *JournalEntry) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *JournalEntry) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCreatedBy

`func (o *JournalEntry) GetCreatedBy() int32`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *JournalEntry) GetCreatedByOk() (*int32, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *JournalEntry) SetCreatedBy(v int32)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *JournalEntry) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *JournalEntry) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *JournalEntry) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetKind

`func (o *JournalEntry) GetKind() Kind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *JournalEntry) GetKindOk() (*Kind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *JournalEntry) SetKind(v Kind)`

SetKind sets Kind field to given value.

### HasKind

`func (o *JournalEntry) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetComments

`func (o *JournalEntry) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *JournalEntry) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *JournalEntry) SetComments(v string)`

SetComments sets Comments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


