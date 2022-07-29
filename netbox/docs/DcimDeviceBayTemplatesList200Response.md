# DcimDeviceBayTemplatesList200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** |  | 
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 
**Results** | [**[]DeviceBayTemplate**](DeviceBayTemplate.md) |  | 

## Methods

### NewDcimDeviceBayTemplatesList200Response

`func NewDcimDeviceBayTemplatesList200Response(count int32, results []DeviceBayTemplate, ) *DcimDeviceBayTemplatesList200Response`

NewDcimDeviceBayTemplatesList200Response instantiates a new DcimDeviceBayTemplatesList200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDcimDeviceBayTemplatesList200ResponseWithDefaults

`func NewDcimDeviceBayTemplatesList200ResponseWithDefaults() *DcimDeviceBayTemplatesList200Response`

NewDcimDeviceBayTemplatesList200ResponseWithDefaults instantiates a new DcimDeviceBayTemplatesList200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *DcimDeviceBayTemplatesList200Response) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *DcimDeviceBayTemplatesList200Response) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *DcimDeviceBayTemplatesList200Response) SetCount(v int32)`

SetCount sets Count field to given value.


### GetNext

`func (o *DcimDeviceBayTemplatesList200Response) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *DcimDeviceBayTemplatesList200Response) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *DcimDeviceBayTemplatesList200Response) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *DcimDeviceBayTemplatesList200Response) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *DcimDeviceBayTemplatesList200Response) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *DcimDeviceBayTemplatesList200Response) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *DcimDeviceBayTemplatesList200Response) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *DcimDeviceBayTemplatesList200Response) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *DcimDeviceBayTemplatesList200Response) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *DcimDeviceBayTemplatesList200Response) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *DcimDeviceBayTemplatesList200Response) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *DcimDeviceBayTemplatesList200Response) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *DcimDeviceBayTemplatesList200Response) GetResults() []DeviceBayTemplate`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *DcimDeviceBayTemplatesList200Response) GetResultsOk() (*[]DeviceBayTemplate, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *DcimDeviceBayTemplatesList200Response) SetResults(v []DeviceBayTemplate)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


