# CiscatResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Benchmark** | Pointer to **string** | CIS-CAT benchmark where the profile is defined | [optional] 
**Error** | Pointer to **int32** | Number of checks that CIS-CAT was not able to run | [optional] 
**Fail** | Pointer to **int32** | Number of failed checks. If this number is higher than 0 the host will probably have a vulnerability | [optional] 
**Notchecked** | Pointer to **int32** | Number of not passed checks | [optional] 
**Pass** | Pointer to **int32** | Number of passed checks | [optional] 
**Profile** | Pointer to **string** | CIS-CAT profile scanned | [optional] 
**Scan** | Pointer to [**ScanIdTime**](ScanIdTime.md) |  | [optional] 
**Score** | Pointer to **int32** | Percentage of passed checks | [optional] 
**Unknown** | Pointer to **int32** | Number of checks which status CIS-CAT was not able to determine | [optional] 

## Methods

### NewCiscatResults

`func NewCiscatResults() *CiscatResults`

NewCiscatResults instantiates a new CiscatResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCiscatResultsWithDefaults

`func NewCiscatResultsWithDefaults() *CiscatResults`

NewCiscatResultsWithDefaults instantiates a new CiscatResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBenchmark

`func (o *CiscatResults) GetBenchmark() string`

GetBenchmark returns the Benchmark field if non-nil, zero value otherwise.

### GetBenchmarkOk

`func (o *CiscatResults) GetBenchmarkOk() (*string, bool)`

GetBenchmarkOk returns a tuple with the Benchmark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBenchmark

`func (o *CiscatResults) SetBenchmark(v string)`

SetBenchmark sets Benchmark field to given value.

### HasBenchmark

`func (o *CiscatResults) HasBenchmark() bool`

HasBenchmark returns a boolean if a field has been set.

### GetError

`func (o *CiscatResults) GetError() int32`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CiscatResults) GetErrorOk() (*int32, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CiscatResults) SetError(v int32)`

SetError sets Error field to given value.

### HasError

`func (o *CiscatResults) HasError() bool`

HasError returns a boolean if a field has been set.

### GetFail

`func (o *CiscatResults) GetFail() int32`

GetFail returns the Fail field if non-nil, zero value otherwise.

### GetFailOk

`func (o *CiscatResults) GetFailOk() (*int32, bool)`

GetFailOk returns a tuple with the Fail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFail

`func (o *CiscatResults) SetFail(v int32)`

SetFail sets Fail field to given value.

### HasFail

`func (o *CiscatResults) HasFail() bool`

HasFail returns a boolean if a field has been set.

### GetNotchecked

`func (o *CiscatResults) GetNotchecked() int32`

GetNotchecked returns the Notchecked field if non-nil, zero value otherwise.

### GetNotcheckedOk

`func (o *CiscatResults) GetNotcheckedOk() (*int32, bool)`

GetNotcheckedOk returns a tuple with the Notchecked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotchecked

`func (o *CiscatResults) SetNotchecked(v int32)`

SetNotchecked sets Notchecked field to given value.

### HasNotchecked

`func (o *CiscatResults) HasNotchecked() bool`

HasNotchecked returns a boolean if a field has been set.

### GetPass

`func (o *CiscatResults) GetPass() int32`

GetPass returns the Pass field if non-nil, zero value otherwise.

### GetPassOk

`func (o *CiscatResults) GetPassOk() (*int32, bool)`

GetPassOk returns a tuple with the Pass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPass

`func (o *CiscatResults) SetPass(v int32)`

SetPass sets Pass field to given value.

### HasPass

`func (o *CiscatResults) HasPass() bool`

HasPass returns a boolean if a field has been set.

### GetProfile

`func (o *CiscatResults) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *CiscatResults) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *CiscatResults) SetProfile(v string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *CiscatResults) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetScan

`func (o *CiscatResults) GetScan() ScanIdTime`

GetScan returns the Scan field if non-nil, zero value otherwise.

### GetScanOk

`func (o *CiscatResults) GetScanOk() (*ScanIdTime, bool)`

GetScanOk returns a tuple with the Scan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScan

`func (o *CiscatResults) SetScan(v ScanIdTime)`

SetScan sets Scan field to given value.

### HasScan

`func (o *CiscatResults) HasScan() bool`

HasScan returns a boolean if a field has been set.

### GetScore

`func (o *CiscatResults) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *CiscatResults) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *CiscatResults) SetScore(v int32)`

SetScore sets Score field to given value.

### HasScore

`func (o *CiscatResults) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetUnknown

`func (o *CiscatResults) GetUnknown() int32`

GetUnknown returns the Unknown field if non-nil, zero value otherwise.

### GetUnknownOk

`func (o *CiscatResults) GetUnknownOk() (*int32, bool)`

GetUnknownOk returns a tuple with the Unknown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknown

`func (o *CiscatResults) SetUnknown(v int32)`

SetUnknown sets Unknown field to given value.

### HasUnknown

`func (o *CiscatResults) HasUnknown() bool`

HasUnknown returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


