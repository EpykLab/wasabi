# LastScan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | Pointer to **NullableTime** | Date when the latest scan finished. If it is in progress, or no scans have been run, null will be returned | [optional] 
**Start** | Pointer to **NullableTime** | Date when the latest scan started. If no scans have been run, null will be returned | [optional] 

## Methods

### NewLastScan

`func NewLastScan() *LastScan`

NewLastScan instantiates a new LastScan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLastScanWithDefaults

`func NewLastScanWithDefaults() *LastScan`

NewLastScanWithDefaults instantiates a new LastScan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *LastScan) GetEnd() time.Time`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *LastScan) GetEndOk() (*time.Time, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *LastScan) SetEnd(v time.Time)`

SetEnd sets End field to given value.

### HasEnd

`func (o *LastScan) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### SetEndNil

`func (o *LastScan) SetEndNil(b bool)`

 SetEndNil sets the value for End to be an explicit nil

### UnsetEnd
`func (o *LastScan) UnsetEnd()`

UnsetEnd ensures that no value is present for End, not even an explicit nil
### GetStart

`func (o *LastScan) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *LastScan) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *LastScan) SetStart(v time.Time)`

SetStart sets Start field to given value.

### HasStart

`func (o *LastScan) HasStart() bool`

HasStart returns a boolean if a field has been set.

### SetStartNil

`func (o *LastScan) SetStartNil(b bool)`

 SetStartNil sets the value for Start to be an explicit nil

### UnsetStart
`func (o *LastScan) UnsetStart()`

UnsetStart ensures that no value is present for Start, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


