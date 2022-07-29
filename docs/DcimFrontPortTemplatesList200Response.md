# DcimFrontPortTemplatesList200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** |  | 
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 
**Results** | [**[]FrontPortTemplate**](FrontPortTemplate.md) |  | 

## Methods

### NewDcimFrontPortTemplatesList200Response

`func NewDcimFrontPortTemplatesList200Response(count int32, results []FrontPortTemplate, ) *DcimFrontPortTemplatesList200Response`

NewDcimFrontPortTemplatesList200Response instantiates a new DcimFrontPortTemplatesList200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDcimFrontPortTemplatesList200ResponseWithDefaults

`func NewDcimFrontPortTemplatesList200ResponseWithDefaults() *DcimFrontPortTemplatesList200Response`

NewDcimFrontPortTemplatesList200ResponseWithDefaults instantiates a new DcimFrontPortTemplatesList200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *DcimFrontPortTemplatesList200Response) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *DcimFrontPortTemplatesList200Response) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *DcimFrontPortTemplatesList200Response) SetCount(v int32)`

SetCount sets Count field to given value.


### GetNext

`func (o *DcimFrontPortTemplatesList200Response) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *DcimFrontPortTemplatesList200Response) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *DcimFrontPortTemplatesList200Response) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *DcimFrontPortTemplatesList200Response) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *DcimFrontPortTemplatesList200Response) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *DcimFrontPortTemplatesList200Response) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *DcimFrontPortTemplatesList200Response) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *DcimFrontPortTemplatesList200Response) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *DcimFrontPortTemplatesList200Response) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *DcimFrontPortTemplatesList200Response) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *DcimFrontPortTemplatesList200Response) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *DcimFrontPortTemplatesList200Response) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *DcimFrontPortTemplatesList200Response) GetResults() []FrontPortTemplate`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *DcimFrontPortTemplatesList200Response) GetResultsOk() (*[]FrontPortTemplate, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *DcimFrontPortTemplatesList200Response) SetResults(v []FrontPortTemplate)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


