# ObjectChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Display** | Pointer to **string** |  | [optional] [readonly] 
**Time** | Pointer to **time.Time** |  | [optional] [readonly] 
**User** | Pointer to [**NestedUser**](NestedUser.md) |  | [optional] 
**UserName** | Pointer to **string** |  | [optional] [readonly] 
**RequestId** | Pointer to **string** |  | [optional] [readonly] 
**Action** | Pointer to [**Action**](Action.md) |  | [optional] 
**ChangedObjectType** | Pointer to **string** |  | [optional] [readonly] 
**ChangedObjectId** | **int32** |  | 
**ChangedObject** | Pointer to **map[string]string** |  Serialize a nested representation of the changed object.  | [optional] [readonly] 
**PrechangeData** | Pointer to **string** |  | [optional] [readonly] 
**PostchangeData** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewObjectChange

`func NewObjectChange(changedObjectId int32, ) *ObjectChange`

NewObjectChange instantiates a new ObjectChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectChangeWithDefaults

`func NewObjectChangeWithDefaults() *ObjectChange`

NewObjectChangeWithDefaults instantiates a new ObjectChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ObjectChange) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ObjectChange) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ObjectChange) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ObjectChange) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *ObjectChange) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ObjectChange) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ObjectChange) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ObjectChange) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDisplay

`func (o *ObjectChange) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *ObjectChange) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *ObjectChange) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *ObjectChange) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetTime

`func (o *ObjectChange) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ObjectChange) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ObjectChange) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *ObjectChange) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetUser

`func (o *ObjectChange) GetUser() NestedUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ObjectChange) GetUserOk() (*NestedUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ObjectChange) SetUser(v NestedUser)`

SetUser sets User field to given value.

### HasUser

`func (o *ObjectChange) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetUserName

`func (o *ObjectChange) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *ObjectChange) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *ObjectChange) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *ObjectChange) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetRequestId

`func (o *ObjectChange) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ObjectChange) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ObjectChange) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *ObjectChange) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetAction

`func (o *ObjectChange) GetAction() Action`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ObjectChange) GetActionOk() (*Action, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ObjectChange) SetAction(v Action)`

SetAction sets Action field to given value.

### HasAction

`func (o *ObjectChange) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetChangedObjectType

`func (o *ObjectChange) GetChangedObjectType() string`

GetChangedObjectType returns the ChangedObjectType field if non-nil, zero value otherwise.

### GetChangedObjectTypeOk

`func (o *ObjectChange) GetChangedObjectTypeOk() (*string, bool)`

GetChangedObjectTypeOk returns a tuple with the ChangedObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedObjectType

`func (o *ObjectChange) SetChangedObjectType(v string)`

SetChangedObjectType sets ChangedObjectType field to given value.

### HasChangedObjectType

`func (o *ObjectChange) HasChangedObjectType() bool`

HasChangedObjectType returns a boolean if a field has been set.

### GetChangedObjectId

`func (o *ObjectChange) GetChangedObjectId() int32`

GetChangedObjectId returns the ChangedObjectId field if non-nil, zero value otherwise.

### GetChangedObjectIdOk

`func (o *ObjectChange) GetChangedObjectIdOk() (*int32, bool)`

GetChangedObjectIdOk returns a tuple with the ChangedObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedObjectId

`func (o *ObjectChange) SetChangedObjectId(v int32)`

SetChangedObjectId sets ChangedObjectId field to given value.


### GetChangedObject

`func (o *ObjectChange) GetChangedObject() map[string]string`

GetChangedObject returns the ChangedObject field if non-nil, zero value otherwise.

### GetChangedObjectOk

`func (o *ObjectChange) GetChangedObjectOk() (*map[string]string, bool)`

GetChangedObjectOk returns a tuple with the ChangedObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedObject

`func (o *ObjectChange) SetChangedObject(v map[string]string)`

SetChangedObject sets ChangedObject field to given value.

### HasChangedObject

`func (o *ObjectChange) HasChangedObject() bool`

HasChangedObject returns a boolean if a field has been set.

### GetPrechangeData

`func (o *ObjectChange) GetPrechangeData() string`

GetPrechangeData returns the PrechangeData field if non-nil, zero value otherwise.

### GetPrechangeDataOk

`func (o *ObjectChange) GetPrechangeDataOk() (*string, bool)`

GetPrechangeDataOk returns a tuple with the PrechangeData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrechangeData

`func (o *ObjectChange) SetPrechangeData(v string)`

SetPrechangeData sets PrechangeData field to given value.

### HasPrechangeData

`func (o *ObjectChange) HasPrechangeData() bool`

HasPrechangeData returns a boolean if a field has been set.

### GetPostchangeData

`func (o *ObjectChange) GetPostchangeData() string`

GetPostchangeData returns the PostchangeData field if non-nil, zero value otherwise.

### GetPostchangeDataOk

`func (o *ObjectChange) GetPostchangeDataOk() (*string, bool)`

GetPostchangeDataOk returns a tuple with the PostchangeData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostchangeData

`func (o *ObjectChange) SetPostchangeData(v string)`

SetPostchangeData sets PostchangeData field to given value.

### HasPostchangeData

`func (o *ObjectChange) HasPostchangeData() bool`

HasPostchangeData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


