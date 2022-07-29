# DcimPlatformsList200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** |  | 
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 
**Results** | [**[]Platform**](Platform.md) |  | 

## Methods

### NewDcimPlatformsList200Response

`func NewDcimPlatformsList200Response(count int32, results []Platform, ) *DcimPlatformsList200Response`

NewDcimPlatformsList200Response instantiates a new DcimPlatformsList200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDcimPlatformsList200ResponseWithDefaults

`func NewDcimPlatformsList200ResponseWithDefaults() *DcimPlatformsList200Response`

NewDcimPlatformsList200ResponseWithDefaults instantiates a new DcimPlatformsList200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *DcimPlatformsList200Response) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *DcimPlatformsList200Response) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *DcimPlatformsList200Response) SetCount(v int32)`

SetCount sets Count field to given value.


### GetNext

`func (o *DcimPlatformsList200Response) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *DcimPlatformsList200Response) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *DcimPlatformsList200Response) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *DcimPlatformsList200Response) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *DcimPlatformsList200Response) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *DcimPlatformsList200Response) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *DcimPlatformsList200Response) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *DcimPlatformsList200Response) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *DcimPlatformsList200Response) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *DcimPlatformsList200Response) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *DcimPlatformsList200Response) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *DcimPlatformsList200Response) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *DcimPlatformsList200Response) GetResults() []Platform`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *DcimPlatformsList200Response) GetResultsOk() (*[]Platform, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *DcimPlatformsList200Response) SetResults(v []Platform)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


