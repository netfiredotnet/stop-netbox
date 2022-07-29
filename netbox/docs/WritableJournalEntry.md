# WritableJournalEntry

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
**Kind** | Pointer to **string** |  | [optional] 
**Comments** | **string** |  | 

## Methods

### NewWritableJournalEntry

`func NewWritableJournalEntry(assignedObjectType string, assignedObjectId int32, comments string, ) *WritableJournalEntry`

NewWritableJournalEntry instantiates a new WritableJournalEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritableJournalEntryWithDefaults

`func NewWritableJournalEntryWithDefaults() *WritableJournalEntry`

NewWritableJournalEntryWithDefaults instantiates a new WritableJournalEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WritableJournalEntry) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WritableJournalEntry) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WritableJournalEntry) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *WritableJournalEntry) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *WritableJournalEntry) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WritableJournalEntry) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WritableJournalEntry) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *WritableJournalEntry) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDisplay

`func (o *WritableJournalEntry) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *WritableJournalEntry) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *WritableJournalEntry) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *WritableJournalEntry) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetAssignedObjectType

`func (o *WritableJournalEntry) GetAssignedObjectType() string`

GetAssignedObjectType returns the AssignedObjectType field if non-nil, zero value otherwise.

### GetAssignedObjectTypeOk

`func (o *WritableJournalEntry) GetAssignedObjectTypeOk() (*string, bool)`

GetAssignedObjectTypeOk returns a tuple with the AssignedObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedObjectType

`func (o *WritableJournalEntry) SetAssignedObjectType(v string)`

SetAssignedObjectType sets AssignedObjectType field to given value.


### GetAssignedObjectId

`func (o *WritableJournalEntry) GetAssignedObjectId() int32`

GetAssignedObjectId returns the AssignedObjectId field if non-nil, zero value otherwise.

### GetAssignedObjectIdOk

`func (o *WritableJournalEntry) GetAssignedObjectIdOk() (*int32, bool)`

GetAssignedObjectIdOk returns a tuple with the AssignedObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedObjectId

`func (o *WritableJournalEntry) SetAssignedObjectId(v int32)`

SetAssignedObjectId sets AssignedObjectId field to given value.


### GetAssignedObject

`func (o *WritableJournalEntry) GetAssignedObject() map[string]string`

GetAssignedObject returns the AssignedObject field if non-nil, zero value otherwise.

### GetAssignedObjectOk

`func (o *WritableJournalEntry) GetAssignedObjectOk() (*map[string]string, bool)`

GetAssignedObjectOk returns a tuple with the AssignedObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedObject

`func (o *WritableJournalEntry) SetAssignedObject(v map[string]string)`

SetAssignedObject sets AssignedObject field to given value.

### HasAssignedObject

`func (o *WritableJournalEntry) HasAssignedObject() bool`

HasAssignedObject returns a boolean if a field has been set.

### GetCreated

`func (o *WritableJournalEntry) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *WritableJournalEntry) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *WritableJournalEntry) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *WritableJournalEntry) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCreatedBy

`func (o *WritableJournalEntry) GetCreatedBy() int32`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *WritableJournalEntry) GetCreatedByOk() (*int32, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *WritableJournalEntry) SetCreatedBy(v int32)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *WritableJournalEntry) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *WritableJournalEntry) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *WritableJournalEntry) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetKind

`func (o *WritableJournalEntry) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *WritableJournalEntry) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *WritableJournalEntry) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *WritableJournalEntry) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetComments

`func (o *WritableJournalEntry) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *WritableJournalEntry) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *WritableJournalEntry) SetComments(v string)`

SetComments sets Comments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


