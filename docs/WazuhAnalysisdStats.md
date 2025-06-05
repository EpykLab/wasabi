# WazuhAnalysisdStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertsQueueSize** | Pointer to **float32** | Pending to write in disk alerts queue size | [optional] 
**AlertsQueueUsage** | Pointer to **float32** | If an event matches a rule, an alert is raised. The alerts are pushed to a _pending to write in disk alerts_ queue. This variable shows usage of that queue | [optional] 
**AlertsWritten** | Pointer to **float32** | Total number of alerts written in disk during the last 5 seconds | [optional] 
**ArchivesQueueSize** | Pointer to **float32** | _Events to write in the archives.log_ queue size | [optional] 
**ArchivesQueueUsage** | Pointer to **float32** | _Events to write in the archives.log_ queue usage | [optional] 
**EventQueueSize** | Pointer to **float32** | Non-catalogued events queue size | [optional] 
**EventQueueUsage** | Pointer to **float32** | Same as &#x60;syscheck_queue_usage&#x60; but for events not catalogued in any of the previously mentioned queues | [optional] 
**EventsDropped** | Pointer to **float32** | Discarded events because they didn&#39;t match any rule in the ruleset | [optional] 
**EventsProcessed** | Pointer to **float32** | Total number of events processed (i.e. matched against Wazuh ruleset) in the last 5 seconds | [optional] 
**EventsReceived** | Pointer to **float32** | Events received in &#x60;analysisd&#x60; from the rest of modules in the last 5 seconds | [optional] 
**FirewallQueueSize** | Pointer to **float32** | _Events to write in the firewall log_ queue size | [optional] 
**FirewallQueueUsage** | Pointer to **float32** | Percentage of use in the queue of events to write in the firewall log | [optional] 
**FirewallWritten** | Pointer to **float32** | Same as &#x60;alerts_written&#x60; but focusing on firewall alerts | [optional] 
**FtsWritten** | Pointer to **float32** | Same as &#x60;alerts_written&#x60; but focusing on [FTS alerts] (https://documentation.wazuh.com/4.12/user-manual/ruleset/ruleset-xml-syntax/decoders.html?highlight&#x3D;fts #fts) | [optional] 
**HostinfoQueueSize** | Pointer to **float32** | Hostinfo events queue size | [optional] 
**HostinfoQueueUsage** | Pointer to **float32** | Same as &#x60;syscheck_queue_usage&#x60; but for hostinfo events | [optional] 
**OtherEventsDecoded** | Pointer to **float32** | Same as &#x60;total_events_decoded&#x60; but for non-catalogued events | [optional] 
**RootcheckEventsDecoded** | Pointer to **float32** | Same as &#x60;total_events_decoded&#x60; but for rootcheck events | [optional] 
**RootcheckQueueSize** | Pointer to **float32** | Rootcheck events queue size | [optional] 
**RootcheckQueueUsage** | Pointer to **float32** | Same as &#x60;syscheck_queue_usage&#x60; but for rootcheck events | [optional] 
**RuleMatchingQueueSize** | Pointer to **float32** | Pending to process events queue size | [optional] 
**RuleMatchingQueueUsage** | Pointer to **float32** | After decoding, events are pushed to a _pending to process_ queue which will match the events against the Wazuh ruleset to raise alerts. This variable shows usage of that queue | [optional] 
**ScaEventsDecoded** | Pointer to **float32** | Same as &#x60;total_events_decoded&#x60; but for policy monitoring events | [optional] 
**ScaQueueSize** | Pointer to **float32** | Policy monitoring events queue size | [optional] 
**ScaQueueUsage** | Pointer to **float32** | Same as &#x60;syscheck_queue_usage&#x60; but for policy monitoring events | [optional] 
**StatisticalQueueSize** | Pointer to **float32** | Stats log queue size | [optional] 
**StatisticalQueueUsage** | Pointer to **float32** | Stats log queue usage | [optional] 
**SyscheckEventsDecoded** | Pointer to **float32** | Same as &#x60;total_events_decoded&#x60; but for FIM events | [optional] 
**SyscheckQueueSize** | Pointer to **float32** | Syscheck events queue size | [optional] 
**SyscheckQueueUsage** | Pointer to **float32** | Percentage of use in the syscheck events queue pending to be decoded. Events are discarded when the queue is full | [optional] 
**SyscollectorEventsDecoded** | Pointer to **float32** | Same as &#x60;total_events_decoded&#x60; but for system inventory events | [optional] 
**SyscollectorQueueSize** | Pointer to **float32** | System inventory events queue size | [optional] 
**SyscollectorQueueUsage** | Pointer to **float32** | Same as &#x60;syscheck_queue_usage&#x60; but for system inventory events | [optional] 
**TotalEventsDecoded** | Pointer to **float32** | Total events decoded in the last 5 seconds. This number is not accumulative, the number in the following 5 seconds can be lower than the previous one | [optional] 
**WinevtEventsDecoded** | Pointer to **float32** | Same as &#x60;total_events_decoded&#x60; but for Windows events | [optional] 
**WinevtQueueSize** | Pointer to **float32** | Windows events queue size | [optional] 
**WinevtQueueUsage** | Pointer to **float32** | Same as &#x60;syscheck_queue_usage&#x60; but for Windows events | [optional] 

## Methods

### NewWazuhAnalysisdStats

`func NewWazuhAnalysisdStats() *WazuhAnalysisdStats`

NewWazuhAnalysisdStats instantiates a new WazuhAnalysisdStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWazuhAnalysisdStatsWithDefaults

`func NewWazuhAnalysisdStatsWithDefaults() *WazuhAnalysisdStats`

NewWazuhAnalysisdStatsWithDefaults instantiates a new WazuhAnalysisdStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertsQueueSize

`func (o *WazuhAnalysisdStats) GetAlertsQueueSize() float32`

GetAlertsQueueSize returns the AlertsQueueSize field if non-nil, zero value otherwise.

### GetAlertsQueueSizeOk

`func (o *WazuhAnalysisdStats) GetAlertsQueueSizeOk() (*float32, bool)`

GetAlertsQueueSizeOk returns a tuple with the AlertsQueueSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertsQueueSize

`func (o *WazuhAnalysisdStats) SetAlertsQueueSize(v float32)`

SetAlertsQueueSize sets AlertsQueueSize field to given value.

### HasAlertsQueueSize

`func (o *WazuhAnalysisdStats) HasAlertsQueueSize() bool`

HasAlertsQueueSize returns a boolean if a field has been set.

### GetAlertsQueueUsage

`func (o *WazuhAnalysisdStats) GetAlertsQueueUsage() float32`

GetAlertsQueueUsage returns the AlertsQueueUsage field if non-nil, zero value otherwise.

### GetAlertsQueueUsageOk

`func (o *WazuhAnalysisdStats) GetAlertsQueueUsageOk() (*float32, bool)`

GetAlertsQueueUsageOk returns a tuple with the AlertsQueueUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertsQueueUsage

`func (o *WazuhAnalysisdStats) SetAlertsQueueUsage(v float32)`

SetAlertsQueueUsage sets AlertsQueueUsage field to given value.

### HasAlertsQueueUsage

`func (o *WazuhAnalysisdStats) HasAlertsQueueUsage() bool`

HasAlertsQueueUsage returns a boolean if a field has been set.

### GetAlertsWritten

`func (o *WazuhAnalysisdStats) GetAlertsWritten() float32`

GetAlertsWritten returns the AlertsWritten field if non-nil, zero value otherwise.

### GetAlertsWrittenOk

`func (o *WazuhAnalysisdStats) GetAlertsWrittenOk() (*float32, bool)`

GetAlertsWrittenOk returns a tuple with the AlertsWritten field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertsWritten

`func (o *WazuhAnalysisdStats) SetAlertsWritten(v float32)`

SetAlertsWritten sets AlertsWritten field to given value.

### HasAlertsWritten

`func (o *WazuhAnalysisdStats) HasAlertsWritten() bool`

HasAlertsWritten returns a boolean if a field has been set.

### GetArchivesQueueSize

`func (o *WazuhAnalysisdStats) GetArchivesQueueSize() float32`

GetArchivesQueueSize returns the ArchivesQueueSize field if non-nil, zero value otherwise.

### GetArchivesQueueSizeOk

`func (o *WazuhAnalysisdStats) GetArchivesQueueSizeOk() (*float32, bool)`

GetArchivesQueueSizeOk returns a tuple with the ArchivesQueueSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivesQueueSize

`func (o *WazuhAnalysisdStats) SetArchivesQueueSize(v float32)`

SetArchivesQueueSize sets ArchivesQueueSize field to given value.

### HasArchivesQueueSize

`func (o *WazuhAnalysisdStats) HasArchivesQueueSize() bool`

HasArchivesQueueSize returns a boolean if a field has been set.

### GetArchivesQueueUsage

`func (o *WazuhAnalysisdStats) GetArchivesQueueUsage() float32`

GetArchivesQueueUsage returns the ArchivesQueueUsage field if non-nil, zero value otherwise.

### GetArchivesQueueUsageOk

`func (o *WazuhAnalysisdStats) GetArchivesQueueUsageOk() (*float32, bool)`

GetArchivesQueueUsageOk returns a tuple with the ArchivesQueueUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivesQueueUsage

`func (o *WazuhAnalysisdStats) SetArchivesQueueUsage(v float32)`

SetArchivesQueueUsage sets ArchivesQueueUsage field to given value.

### HasArchivesQueueUsage

`func (o *WazuhAnalysisdStats) HasArchivesQueueUsage() bool`

HasArchivesQueueUsage returns a boolean if a field has been set.

### GetEventQueueSize

`func (o *WazuhAnalysisdStats) GetEventQueueSize() float32`

GetEventQueueSize returns the EventQueueSize field if non-nil, zero value otherwise.

### GetEventQueueSizeOk

`func (o *WazuhAnalysisdStats) GetEventQueueSizeOk() (*float32, bool)`

GetEventQueueSizeOk returns a tuple with the EventQueueSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventQueueSize

`func (o *WazuhAnalysisdStats) SetEventQueueSize(v float32)`

SetEventQueueSize sets EventQueueSize field to given value.

### HasEventQueueSize

`func (o *WazuhAnalysisdStats) HasEventQueueSize() bool`

HasEventQueueSize returns a boolean if a field has been set.

### GetEventQueueUsage

`func (o *WazuhAnalysisdStats) GetEventQueueUsage() float32`

GetEventQueueUsage returns the EventQueueUsage field if non-nil, zero value otherwise.

### GetEventQueueUsageOk

`func (o *WazuhAnalysisdStats) GetEventQueueUsageOk() (*float32, bool)`

GetEventQueueUsageOk returns a tuple with the EventQueueUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventQueueUsage

`func (o *WazuhAnalysisdStats) SetEventQueueUsage(v float32)`

SetEventQueueUsage sets EventQueueUsage field to given value.

### HasEventQueueUsage

`func (o *WazuhAnalysisdStats) HasEventQueueUsage() bool`

HasEventQueueUsage returns a boolean if a field has been set.

### GetEventsDropped

`func (o *WazuhAnalysisdStats) GetEventsDropped() float32`

GetEventsDropped returns the EventsDropped field if non-nil, zero value otherwise.

### GetEventsDroppedOk

`func (o *WazuhAnalysisdStats) GetEventsDroppedOk() (*float32, bool)`

GetEventsDroppedOk returns a tuple with the EventsDropped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsDropped

`func (o *WazuhAnalysisdStats) SetEventsDropped(v float32)`

SetEventsDropped sets EventsDropped field to given value.

### HasEventsDropped

`func (o *WazuhAnalysisdStats) HasEventsDropped() bool`

HasEventsDropped returns a boolean if a field has been set.

### GetEventsProcessed

`func (o *WazuhAnalysisdStats) GetEventsProcessed() float32`

GetEventsProcessed returns the EventsProcessed field if non-nil, zero value otherwise.

### GetEventsProcessedOk

`func (o *WazuhAnalysisdStats) GetEventsProcessedOk() (*float32, bool)`

GetEventsProcessedOk returns a tuple with the EventsProcessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsProcessed

`func (o *WazuhAnalysisdStats) SetEventsProcessed(v float32)`

SetEventsProcessed sets EventsProcessed field to given value.

### HasEventsProcessed

`func (o *WazuhAnalysisdStats) HasEventsProcessed() bool`

HasEventsProcessed returns a boolean if a field has been set.

### GetEventsReceived

`func (o *WazuhAnalysisdStats) GetEventsReceived() float32`

GetEventsReceived returns the EventsReceived field if non-nil, zero value otherwise.

### GetEventsReceivedOk

`func (o *WazuhAnalysisdStats) GetEventsReceivedOk() (*float32, bool)`

GetEventsReceivedOk returns a tuple with the EventsReceived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsReceived

`func (o *WazuhAnalysisdStats) SetEventsReceived(v float32)`

SetEventsReceived sets EventsReceived field to given value.

### HasEventsReceived

`func (o *WazuhAnalysisdStats) HasEventsReceived() bool`

HasEventsReceived returns a boolean if a field has been set.

### GetFirewallQueueSize

`func (o *WazuhAnalysisdStats) GetFirewallQueueSize() float32`

GetFirewallQueueSize returns the FirewallQueueSize field if non-nil, zero value otherwise.

### GetFirewallQueueSizeOk

`func (o *WazuhAnalysisdStats) GetFirewallQueueSizeOk() (*float32, bool)`

GetFirewallQueueSizeOk returns a tuple with the FirewallQueueSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallQueueSize

`func (o *WazuhAnalysisdStats) SetFirewallQueueSize(v float32)`

SetFirewallQueueSize sets FirewallQueueSize field to given value.

### HasFirewallQueueSize

`func (o *WazuhAnalysisdStats) HasFirewallQueueSize() bool`

HasFirewallQueueSize returns a boolean if a field has been set.

### GetFirewallQueueUsage

`func (o *WazuhAnalysisdStats) GetFirewallQueueUsage() float32`

GetFirewallQueueUsage returns the FirewallQueueUsage field if non-nil, zero value otherwise.

### GetFirewallQueueUsageOk

`func (o *WazuhAnalysisdStats) GetFirewallQueueUsageOk() (*float32, bool)`

GetFirewallQueueUsageOk returns a tuple with the FirewallQueueUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallQueueUsage

`func (o *WazuhAnalysisdStats) SetFirewallQueueUsage(v float32)`

SetFirewallQueueUsage sets FirewallQueueUsage field to given value.

### HasFirewallQueueUsage

`func (o *WazuhAnalysisdStats) HasFirewallQueueUsage() bool`

HasFirewallQueueUsage returns a boolean if a field has been set.

### GetFirewallWritten

`func (o *WazuhAnalysisdStats) GetFirewallWritten() float32`

GetFirewallWritten returns the FirewallWritten field if non-nil, zero value otherwise.

### GetFirewallWrittenOk

`func (o *WazuhAnalysisdStats) GetFirewallWrittenOk() (*float32, bool)`

GetFirewallWrittenOk returns a tuple with the FirewallWritten field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallWritten

`func (o *WazuhAnalysisdStats) SetFirewallWritten(v float32)`

SetFirewallWritten sets FirewallWritten field to given value.

### HasFirewallWritten

`func (o *WazuhAnalysisdStats) HasFirewallWritten() bool`

HasFirewallWritten returns a boolean if a field has been set.

### GetFtsWritten

`func (o *WazuhAnalysisdStats) GetFtsWritten() float32`

GetFtsWritten returns the FtsWritten field if non-nil, zero value otherwise.

### GetFtsWrittenOk

`func (o *WazuhAnalysisdStats) GetFtsWrittenOk() (*float32, bool)`

GetFtsWrittenOk returns a tuple with the FtsWritten field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFtsWritten

`func (o *WazuhAnalysisdStats) SetFtsWritten(v float32)`

SetFtsWritten sets FtsWritten field to given value.

### HasFtsWritten

`func (o *WazuhAnalysisdStats) HasFtsWritten() bool`

HasFtsWritten returns a boolean if a field has been set.

### GetHostinfoQueueSize

`func (o *WazuhAnalysisdStats) GetHostinfoQueueSize() float32`

GetHostinfoQueueSize returns the HostinfoQueueSize field if non-nil, zero value otherwise.

### GetHostinfoQueueSizeOk

`func (o *WazuhAnalysisdStats) GetHostinfoQueueSizeOk() (*float32, bool)`

GetHostinfoQueueSizeOk returns a tuple with the HostinfoQueueSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostinfoQueueSize

`func (o *WazuhAnalysisdStats) SetHostinfoQueueSize(v float32)`

SetHostinfoQueueSize sets HostinfoQueueSize field to given value.

### HasHostinfoQueueSize

`func (o *WazuhAnalysisdStats) HasHostinfoQueueSize() bool`

HasHostinfoQueueSize returns a boolean if a field has been set.

### GetHostinfoQueueUsage

`func (o *WazuhAnalysisdStats) GetHostinfoQueueUsage() float32`

GetHostinfoQueueUsage returns the HostinfoQueueUsage field if non-nil, zero value otherwise.

### GetHostinfoQueueUsageOk

`func (o *WazuhAnalysisdStats) GetHostinfoQueueUsageOk() (*float32, bool)`

GetHostinfoQueueUsageOk returns a tuple with the HostinfoQueueUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostinfoQueueUsage

`func (o *WazuhAnalysisdStats) SetHostinfoQueueUsage(v float32)`

SetHostinfoQueueUsage sets HostinfoQueueUsage field to given value.

### HasHostinfoQueueUsage

`func (o *WazuhAnalysisdStats) HasHostinfoQueueUsage() bool`

HasHostinfoQueueUsage returns a boolean if a field has been set.

### GetOtherEventsDecoded

`func (o *WazuhAnalysisdStats) GetOtherEventsDecoded() float32`

GetOtherEventsDecoded returns the OtherEventsDecoded field if non-nil, zero value otherwise.

### GetOtherEventsDecodedOk

`func (o *WazuhAnalysisdStats) GetOtherEventsDecodedOk() (*float32, bool)`

GetOtherEventsDecodedOk returns a tuple with the OtherEventsDecoded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherEventsDecoded

`func (o *WazuhAnalysisdStats) SetOtherEventsDecoded(v float32)`

SetOtherEventsDecoded sets OtherEventsDecoded field to given value.

### HasOtherEventsDecoded

`func (o *WazuhAnalysisdStats) HasOtherEventsDecoded() bool`

HasOtherEventsDecoded returns a boolean if a field has been set.

### GetRootcheckEventsDecoded

`func (o *WazuhAnalysisdStats) GetRootcheckEventsDecoded() float32`

GetRootcheckEventsDecoded returns the RootcheckEventsDecoded field if non-nil, zero value otherwise.

### GetRootcheckEventsDecodedOk

`func (o *WazuhAnalysisdStats) GetRootcheckEventsDecodedOk() (*float32, bool)`

GetRootcheckEventsDecodedOk returns a tuple with the RootcheckEventsDecoded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootcheckEventsDecoded

`func (o *WazuhAnalysisdStats) SetRootcheckEventsDecoded(v float32)`

SetRootcheckEventsDecoded sets RootcheckEventsDecoded field to given value.

### HasRootcheckEventsDecoded

`func (o *WazuhAnalysisdStats) HasRootcheckEventsDecoded() bool`

HasRootcheckEventsDecoded returns a boolean if a field has been set.

### GetRootcheckQueueSize

`func (o *WazuhAnalysisdStats) GetRootcheckQueueSize() float32`

GetRootcheckQueueSize returns the RootcheckQueueSize field if non-nil, zero value otherwise.

### GetRootcheckQueueSizeOk

`func (o *WazuhAnalysisdStats) GetRootcheckQueueSizeOk() (*float32, bool)`

GetRootcheckQueueSizeOk returns a tuple with the RootcheckQueueSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootcheckQueueSize

`func (o *WazuhAnalysisdStats) SetRootcheckQueueSize(v float32)`

SetRootcheckQueueSize sets RootcheckQueueSize field to given value.

### HasRootcheckQueueSize

`func (o *WazuhAnalysisdStats) HasRootcheckQueueSize() bool`

HasRootcheckQueueSize returns a boolean if a field has been set.

### GetRootcheckQueueUsage

`func (o *WazuhAnalysisdStats) GetRootcheckQueueUsage() float32`

GetRootcheckQueueUsage returns the RootcheckQueueUsage field if non-nil, zero value otherwise.

### GetRootcheckQueueUsageOk

`func (o *WazuhAnalysisdStats) GetRootcheckQueueUsageOk() (*float32, bool)`

GetRootcheckQueueUsageOk returns a tuple with the RootcheckQueueUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootcheckQueueUsage

`func (o *WazuhAnalysisdStats) SetRootcheckQueueUsage(v float32)`

SetRootcheckQueueUsage sets RootcheckQueueUsage field to given value.

### HasRootcheckQueueUsage

`func (o *WazuhAnalysisdStats) HasRootcheckQueueUsage() bool`

HasRootcheckQueueUsage returns a boolean if a field has been set.

### GetRuleMatchingQueueSize

`func (o *WazuhAnalysisdStats) GetRuleMatchingQueueSize() float32`

GetRuleMatchingQueueSize returns the RuleMatchingQueueSize field if non-nil, zero value otherwise.

### GetRuleMatchingQueueSizeOk

`func (o *WazuhAnalysisdStats) GetRuleMatchingQueueSizeOk() (*float32, bool)`

GetRuleMatchingQueueSizeOk returns a tuple with the RuleMatchingQueueSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleMatchingQueueSize

`func (o *WazuhAnalysisdStats) SetRuleMatchingQueueSize(v float32)`

SetRuleMatchingQueueSize sets RuleMatchingQueueSize field to given value.

### HasRuleMatchingQueueSize

`func (o *WazuhAnalysisdStats) HasRuleMatchingQueueSize() bool`

HasRuleMatchingQueueSize returns a boolean if a field has been set.

### GetRuleMatchingQueueUsage

`func (o *WazuhAnalysisdStats) GetRuleMatchingQueueUsage() float32`

GetRuleMatchingQueueUsage returns the RuleMatchingQueueUsage field if non-nil, zero value otherwise.

### GetRuleMatchingQueueUsageOk

`func (o *WazuhAnalysisdStats) GetRuleMatchingQueueUsageOk() (*float32, bool)`

GetRuleMatchingQueueUsageOk returns a tuple with the RuleMatchingQueueUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleMatchingQueueUsage

`func (o *WazuhAnalysisdStats) SetRuleMatchingQueueUsage(v float32)`

SetRuleMatchingQueueUsage sets RuleMatchingQueueUsage field to given value.

### HasRuleMatchingQueueUsage

`func (o *WazuhAnalysisdStats) HasRuleMatchingQueueUsage() bool`

HasRuleMatchingQueueUsage returns a boolean if a field has been set.

### GetScaEventsDecoded

`func (o *WazuhAnalysisdStats) GetScaEventsDecoded() float32`

GetScaEventsDecoded returns the ScaEventsDecoded field if non-nil, zero value otherwise.

### GetScaEventsDecodedOk

`func (o *WazuhAnalysisdStats) GetScaEventsDecodedOk() (*float32, bool)`

GetScaEventsDecodedOk returns a tuple with the ScaEventsDecoded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaEventsDecoded

`func (o *WazuhAnalysisdStats) SetScaEventsDecoded(v float32)`

SetScaEventsDecoded sets ScaEventsDecoded field to given value.

### HasScaEventsDecoded

`func (o *WazuhAnalysisdStats) HasScaEventsDecoded() bool`

HasScaEventsDecoded returns a boolean if a field has been set.

### GetScaQueueSize

`func (o *WazuhAnalysisdStats) GetScaQueueSize() float32`

GetScaQueueSize returns the ScaQueueSize field if non-nil, zero value otherwise.

### GetScaQueueSizeOk

`func (o *WazuhAnalysisdStats) GetScaQueueSizeOk() (*float32, bool)`

GetScaQueueSizeOk returns a tuple with the ScaQueueSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaQueueSize

`func (o *WazuhAnalysisdStats) SetScaQueueSize(v float32)`

SetScaQueueSize sets ScaQueueSize field to given value.

### HasScaQueueSize

`func (o *WazuhAnalysisdStats) HasScaQueueSize() bool`

HasScaQueueSize returns a boolean if a field has been set.

### GetScaQueueUsage

`func (o *WazuhAnalysisdStats) GetScaQueueUsage() float32`

GetScaQueueUsage returns the ScaQueueUsage field if non-nil, zero value otherwise.

### GetScaQueueUsageOk

`func (o *WazuhAnalysisdStats) GetScaQueueUsageOk() (*float32, bool)`

GetScaQueueUsageOk returns a tuple with the ScaQueueUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaQueueUsage

`func (o *WazuhAnalysisdStats) SetScaQueueUsage(v float32)`

SetScaQueueUsage sets ScaQueueUsage field to given value.

### HasScaQueueUsage

`func (o *WazuhAnalysisdStats) HasScaQueueUsage() bool`

HasScaQueueUsage returns a boolean if a field has been set.

### GetStatisticalQueueSize

`func (o *WazuhAnalysisdStats) GetStatisticalQueueSize() float32`

GetStatisticalQueueSize returns the StatisticalQueueSize field if non-nil, zero value otherwise.

### GetStatisticalQueueSizeOk

`func (o *WazuhAnalysisdStats) GetStatisticalQueueSizeOk() (*float32, bool)`

GetStatisticalQueueSizeOk returns a tuple with the StatisticalQueueSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatisticalQueueSize

`func (o *WazuhAnalysisdStats) SetStatisticalQueueSize(v float32)`

SetStatisticalQueueSize sets StatisticalQueueSize field to given value.

### HasStatisticalQueueSize

`func (o *WazuhAnalysisdStats) HasStatisticalQueueSize() bool`

HasStatisticalQueueSize returns a boolean if a field has been set.

### GetStatisticalQueueUsage

`func (o *WazuhAnalysisdStats) GetStatisticalQueueUsage() float32`

GetStatisticalQueueUsage returns the StatisticalQueueUsage field if non-nil, zero value otherwise.

### GetStatisticalQueueUsageOk

`func (o *WazuhAnalysisdStats) GetStatisticalQueueUsageOk() (*float32, bool)`

GetStatisticalQueueUsageOk returns a tuple with the StatisticalQueueUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatisticalQueueUsage

`func (o *WazuhAnalysisdStats) SetStatisticalQueueUsage(v float32)`

SetStatisticalQueueUsage sets StatisticalQueueUsage field to given value.

### HasStatisticalQueueUsage

`func (o *WazuhAnalysisdStats) HasStatisticalQueueUsage() bool`

HasStatisticalQueueUsage returns a boolean if a field has been set.

### GetSyscheckEventsDecoded

`func (o *WazuhAnalysisdStats) GetSyscheckEventsDecoded() float32`

GetSyscheckEventsDecoded returns the SyscheckEventsDecoded field if non-nil, zero value otherwise.

### GetSyscheckEventsDecodedOk

`func (o *WazuhAnalysisdStats) GetSyscheckEventsDecodedOk() (*float32, bool)`

GetSyscheckEventsDecodedOk returns a tuple with the SyscheckEventsDecoded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyscheckEventsDecoded

`func (o *WazuhAnalysisdStats) SetSyscheckEventsDecoded(v float32)`

SetSyscheckEventsDecoded sets SyscheckEventsDecoded field to given value.

### HasSyscheckEventsDecoded

`func (o *WazuhAnalysisdStats) HasSyscheckEventsDecoded() bool`

HasSyscheckEventsDecoded returns a boolean if a field has been set.

### GetSyscheckQueueSize

`func (o *WazuhAnalysisdStats) GetSyscheckQueueSize() float32`

GetSyscheckQueueSize returns the SyscheckQueueSize field if non-nil, zero value otherwise.

### GetSyscheckQueueSizeOk

`func (o *WazuhAnalysisdStats) GetSyscheckQueueSizeOk() (*float32, bool)`

GetSyscheckQueueSizeOk returns a tuple with the SyscheckQueueSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyscheckQueueSize

`func (o *WazuhAnalysisdStats) SetSyscheckQueueSize(v float32)`

SetSyscheckQueueSize sets SyscheckQueueSize field to given value.

### HasSyscheckQueueSize

`func (o *WazuhAnalysisdStats) HasSyscheckQueueSize() bool`

HasSyscheckQueueSize returns a boolean if a field has been set.

### GetSyscheckQueueUsage

`func (o *WazuhAnalysisdStats) GetSyscheckQueueUsage() float32`

GetSyscheckQueueUsage returns the SyscheckQueueUsage field if non-nil, zero value otherwise.

### GetSyscheckQueueUsageOk

`func (o *WazuhAnalysisdStats) GetSyscheckQueueUsageOk() (*float32, bool)`

GetSyscheckQueueUsageOk returns a tuple with the SyscheckQueueUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyscheckQueueUsage

`func (o *WazuhAnalysisdStats) SetSyscheckQueueUsage(v float32)`

SetSyscheckQueueUsage sets SyscheckQueueUsage field to given value.

### HasSyscheckQueueUsage

`func (o *WazuhAnalysisdStats) HasSyscheckQueueUsage() bool`

HasSyscheckQueueUsage returns a boolean if a field has been set.

### GetSyscollectorEventsDecoded

`func (o *WazuhAnalysisdStats) GetSyscollectorEventsDecoded() float32`

GetSyscollectorEventsDecoded returns the SyscollectorEventsDecoded field if non-nil, zero value otherwise.

### GetSyscollectorEventsDecodedOk

`func (o *WazuhAnalysisdStats) GetSyscollectorEventsDecodedOk() (*float32, bool)`

GetSyscollectorEventsDecodedOk returns a tuple with the SyscollectorEventsDecoded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyscollectorEventsDecoded

`func (o *WazuhAnalysisdStats) SetSyscollectorEventsDecoded(v float32)`

SetSyscollectorEventsDecoded sets SyscollectorEventsDecoded field to given value.

### HasSyscollectorEventsDecoded

`func (o *WazuhAnalysisdStats) HasSyscollectorEventsDecoded() bool`

HasSyscollectorEventsDecoded returns a boolean if a field has been set.

### GetSyscollectorQueueSize

`func (o *WazuhAnalysisdStats) GetSyscollectorQueueSize() float32`

GetSyscollectorQueueSize returns the SyscollectorQueueSize field if non-nil, zero value otherwise.

### GetSyscollectorQueueSizeOk

`func (o *WazuhAnalysisdStats) GetSyscollectorQueueSizeOk() (*float32, bool)`

GetSyscollectorQueueSizeOk returns a tuple with the SyscollectorQueueSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyscollectorQueueSize

`func (o *WazuhAnalysisdStats) SetSyscollectorQueueSize(v float32)`

SetSyscollectorQueueSize sets SyscollectorQueueSize field to given value.

### HasSyscollectorQueueSize

`func (o *WazuhAnalysisdStats) HasSyscollectorQueueSize() bool`

HasSyscollectorQueueSize returns a boolean if a field has been set.

### GetSyscollectorQueueUsage

`func (o *WazuhAnalysisdStats) GetSyscollectorQueueUsage() float32`

GetSyscollectorQueueUsage returns the SyscollectorQueueUsage field if non-nil, zero value otherwise.

### GetSyscollectorQueueUsageOk

`func (o *WazuhAnalysisdStats) GetSyscollectorQueueUsageOk() (*float32, bool)`

GetSyscollectorQueueUsageOk returns a tuple with the SyscollectorQueueUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyscollectorQueueUsage

`func (o *WazuhAnalysisdStats) SetSyscollectorQueueUsage(v float32)`

SetSyscollectorQueueUsage sets SyscollectorQueueUsage field to given value.

### HasSyscollectorQueueUsage

`func (o *WazuhAnalysisdStats) HasSyscollectorQueueUsage() bool`

HasSyscollectorQueueUsage returns a boolean if a field has been set.

### GetTotalEventsDecoded

`func (o *WazuhAnalysisdStats) GetTotalEventsDecoded() float32`

GetTotalEventsDecoded returns the TotalEventsDecoded field if non-nil, zero value otherwise.

### GetTotalEventsDecodedOk

`func (o *WazuhAnalysisdStats) GetTotalEventsDecodedOk() (*float32, bool)`

GetTotalEventsDecodedOk returns a tuple with the TotalEventsDecoded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEventsDecoded

`func (o *WazuhAnalysisdStats) SetTotalEventsDecoded(v float32)`

SetTotalEventsDecoded sets TotalEventsDecoded field to given value.

### HasTotalEventsDecoded

`func (o *WazuhAnalysisdStats) HasTotalEventsDecoded() bool`

HasTotalEventsDecoded returns a boolean if a field has been set.

### GetWinevtEventsDecoded

`func (o *WazuhAnalysisdStats) GetWinevtEventsDecoded() float32`

GetWinevtEventsDecoded returns the WinevtEventsDecoded field if non-nil, zero value otherwise.

### GetWinevtEventsDecodedOk

`func (o *WazuhAnalysisdStats) GetWinevtEventsDecodedOk() (*float32, bool)`

GetWinevtEventsDecodedOk returns a tuple with the WinevtEventsDecoded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinevtEventsDecoded

`func (o *WazuhAnalysisdStats) SetWinevtEventsDecoded(v float32)`

SetWinevtEventsDecoded sets WinevtEventsDecoded field to given value.

### HasWinevtEventsDecoded

`func (o *WazuhAnalysisdStats) HasWinevtEventsDecoded() bool`

HasWinevtEventsDecoded returns a boolean if a field has been set.

### GetWinevtQueueSize

`func (o *WazuhAnalysisdStats) GetWinevtQueueSize() float32`

GetWinevtQueueSize returns the WinevtQueueSize field if non-nil, zero value otherwise.

### GetWinevtQueueSizeOk

`func (o *WazuhAnalysisdStats) GetWinevtQueueSizeOk() (*float32, bool)`

GetWinevtQueueSizeOk returns a tuple with the WinevtQueueSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinevtQueueSize

`func (o *WazuhAnalysisdStats) SetWinevtQueueSize(v float32)`

SetWinevtQueueSize sets WinevtQueueSize field to given value.

### HasWinevtQueueSize

`func (o *WazuhAnalysisdStats) HasWinevtQueueSize() bool`

HasWinevtQueueSize returns a boolean if a field has been set.

### GetWinevtQueueUsage

`func (o *WazuhAnalysisdStats) GetWinevtQueueUsage() float32`

GetWinevtQueueUsage returns the WinevtQueueUsage field if non-nil, zero value otherwise.

### GetWinevtQueueUsageOk

`func (o *WazuhAnalysisdStats) GetWinevtQueueUsageOk() (*float32, bool)`

GetWinevtQueueUsageOk returns a tuple with the WinevtQueueUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinevtQueueUsage

`func (o *WazuhAnalysisdStats) SetWinevtQueueUsage(v float32)`

SetWinevtQueueUsage sets WinevtQueueUsage field to given value.

### HasWinevtQueueUsage

`func (o *WazuhAnalysisdStats) HasWinevtQueueUsage() bool`

HasWinevtQueueUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


