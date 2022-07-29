# DcimPowerPortTemplatesList200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** |  | 
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 
**Results** | [**[]PowerPortTemplate**](PowerPortTemplate.md) |  | 

## Methods

### NewDcimPowerPortTemplatesList200Response

`func NewDcimPowerPortTemplatesList200Response(count int32, results []PowerPortTemplate, ) *DcimPowerPortTemplatesList200Response`

NewDcimPowerPortTemplatesList200Response instantiates a new DcimPowerPortTemplatesList200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDcimPowerPortTemplatesList200ResponseWithDefaults

`func NewDcimPowerPortTemplatesList200ResponseWithDefaults() *DcimPowerPortTemplatesList200Response`

NewDcimPowerPortTemplatesList200ResponseWithDefaults instantiates a new DcimPowerPortTemplatesList200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *DcimPowerPortTemplatesList200Response) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *DcimPowerPortTemplatesList200Response) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *DcimPowerPortTemplatesList200Response) SetCount(v int32)`

SetCount sets Count field to given value.


### GetNext

`func (o *DcimPowerPortTemplatesList200Response) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *DcimPowerPortTemplatesList200Response) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *DcimPowerPortTemplatesList200Response) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *DcimPowerPortTemplatesList200Response) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *DcimPowerPortTemplatesList200Response) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *DcimPowerPortTemplatesList200Response) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *DcimPowerPortTemplatesList200Response) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *DcimPowerPortTemplatesList200Response) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *DcimPowerPortTemplatesList200Response) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *DcimPowerPortTemplatesList200Response) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *DcimPowerPortTemplatesList200Response) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *DcimPowerPortTemplatesList200Response) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *DcimPowerPortTemplatesList200Response) GetResults() []PowerPortTemplate`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *DcimPowerPortTemplatesList200Response) GetResultsOk() (*[]PowerPortTemplate, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *DcimPowerPortTemplatesList200Response) SetResults(v []PowerPortTemplate)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


