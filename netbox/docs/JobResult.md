# JobResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Display** | Pointer to **string** |  | [optional] [readonly] 
**Created** | Pointer to **time.Time** |  | [optional] [readonly] 
**Completed** | Pointer to **NullableTime** |  | [optional] 
**Name** | **string** |  | 
**ObjType** | Pointer to **string** |  | [optional] [readonly] 
**Status** | Pointer to [**Status6**](Status6.md) |  | [optional] 
**User** | Pointer to [**NestedUser**](NestedUser.md) |  | [optional] 
**Data** | Pointer to **NullableString** |  | [optional] 
**JobId** | **string** |  | 

## Methods

### NewJobResult

`func NewJobResult(name string, jobId string, ) *JobResult`

NewJobResult instantiates a new JobResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobResultWithDefaults

`func NewJobResultWithDefaults() *JobResult`

NewJobResultWithDefaults instantiates a new JobResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JobResult) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JobResult) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JobResult) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *JobResult) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *JobResult) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *JobResult) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *JobResult) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *JobResult) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDisplay

`func (o *JobResult) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *JobResult) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *JobResult) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *JobResult) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetCreated

`func (o *JobResult) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *JobResult) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *JobResult) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *JobResult) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCompleted

`func (o *JobResult) GetCompleted() time.Time`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *JobResult) GetCompletedOk() (*time.Time, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *JobResult) SetCompleted(v time.Time)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *JobResult) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.

### SetCompletedNil

`func (o *JobResult) SetCompletedNil(b bool)`

 SetCompletedNil sets the value for Completed to be an explicit nil

### UnsetCompleted
`func (o *JobResult) UnsetCompleted()`

UnsetCompleted ensures that no value is present for Completed, not even an explicit nil
### GetName

`func (o *JobResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JobResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JobResult) SetName(v string)`

SetName sets Name field to given value.


### GetObjType

`func (o *JobResult) GetObjType() string`

GetObjType returns the ObjType field if non-nil, zero value otherwise.

### GetObjTypeOk

`func (o *JobResult) GetObjTypeOk() (*string, bool)`

GetObjTypeOk returns a tuple with the ObjType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjType

`func (o *JobResult) SetObjType(v string)`

SetObjType sets ObjType field to given value.

### HasObjType

`func (o *JobResult) HasObjType() bool`

HasObjType returns a boolean if a field has been set.

### GetStatus

`func (o *JobResult) GetStatus() Status6`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *JobResult) GetStatusOk() (*Status6, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *JobResult) SetStatus(v Status6)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *JobResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUser

`func (o *JobResult) GetUser() NestedUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *JobResult) GetUserOk() (*NestedUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *JobResult) SetUser(v NestedUser)`

SetUser sets User field to given value.

### HasUser

`func (o *JobResult) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetData

`func (o *JobResult) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *JobResult) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *JobResult) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *JobResult) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *JobResult) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *JobResult) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetJobId

`func (o *JobResult) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *JobResult) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *JobResult) SetJobId(v string)`

SetJobId sets JobId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


