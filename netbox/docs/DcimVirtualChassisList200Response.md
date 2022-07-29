# DcimVirtualChassisList200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** |  | 
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 
**Results** | [**[]VirtualChassis**](VirtualChassis.md) |  | 

## Methods

### NewDcimVirtualChassisList200Response

`func NewDcimVirtualChassisList200Response(count int32, results []VirtualChassis, ) *DcimVirtualChassisList200Response`

NewDcimVirtualChassisList200Response instantiates a new DcimVirtualChassisList200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDcimVirtualChassisList200ResponseWithDefaults

`func NewDcimVirtualChassisList200ResponseWithDefaults() *DcimVirtualChassisList200Response`

NewDcimVirtualChassisList200ResponseWithDefaults instantiates a new DcimVirtualChassisList200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *DcimVirtualChassisList200Response) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *DcimVirtualChassisList200Response) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *DcimVirtualChassisList200Response) SetCount(v int32)`

SetCount sets Count field to given value.


### GetNext

`func (o *DcimVirtualChassisList200Response) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *DcimVirtualChassisList200Response) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *DcimVirtualChassisList200Response) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *DcimVirtualChassisList200Response) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *DcimVirtualChassisList200Response) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *DcimVirtualChassisList200Response) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *DcimVirtualChassisList200Response) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *DcimVirtualChassisList200Response) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *DcimVirtualChassisList200Response) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *DcimVirtualChassisList200Response) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *DcimVirtualChassisList200Response) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *DcimVirtualChassisList200Response) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *DcimVirtualChassisList200Response) GetResults() []VirtualChassis`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *DcimVirtualChassisList200Response) GetResultsOk() (*[]VirtualChassis, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *DcimVirtualChassisList200Response) SetResults(v []VirtualChassis)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


