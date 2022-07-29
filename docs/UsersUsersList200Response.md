# UsersUsersList200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** |  | 
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 
**Results** | [**[]User**](User.md) |  | 

## Methods

### NewUsersUsersList200Response

`func NewUsersUsersList200Response(count int32, results []User, ) *UsersUsersList200Response`

NewUsersUsersList200Response instantiates a new UsersUsersList200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersUsersList200ResponseWithDefaults

`func NewUsersUsersList200ResponseWithDefaults() *UsersUsersList200Response`

NewUsersUsersList200ResponseWithDefaults instantiates a new UsersUsersList200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *UsersUsersList200Response) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *UsersUsersList200Response) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *UsersUsersList200Response) SetCount(v int32)`

SetCount sets Count field to given value.


### GetNext

`func (o *UsersUsersList200Response) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *UsersUsersList200Response) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *UsersUsersList200Response) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *UsersUsersList200Response) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *UsersUsersList200Response) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *UsersUsersList200Response) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *UsersUsersList200Response) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *UsersUsersList200Response) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *UsersUsersList200Response) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *UsersUsersList200Response) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *UsersUsersList200Response) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *UsersUsersList200Response) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *UsersUsersList200Response) GetResults() []User`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *UsersUsersList200Response) GetResultsOk() (*[]User, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *UsersUsersList200Response) SetResults(v []User)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


