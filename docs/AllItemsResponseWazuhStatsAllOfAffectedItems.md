# AllItemsResponseWazuhStatsAllOfAffectedItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Averages** | Pointer to **[]int32** | Array containing the number of alerts for every hour | [optional] 
**Interactions** | Pointer to **int32** |  | [optional] 
**Sun** | Pointer to [**WazuhHourlyStats**](WazuhHourlyStats.md) |  | [optional] 
**Mon** | Pointer to [**WazuhHourlyStats**](WazuhHourlyStats.md) |  | [optional] 
**Tue** | Pointer to [**WazuhHourlyStats**](WazuhHourlyStats.md) |  | [optional] 
**Wed** | Pointer to [**WazuhHourlyStats**](WazuhHourlyStats.md) |  | [optional] 
**Thu** | Pointer to [**WazuhHourlyStats**](WazuhHourlyStats.md) |  | [optional] 
**Fri** | Pointer to [**WazuhHourlyStats**](WazuhHourlyStats.md) |  | [optional] 
**Sat** | Pointer to [**WazuhHourlyStats**](WazuhHourlyStats.md) |  | [optional] 
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
**CtrlMsgCount** | Pointer to **float32** | Number of control messages received from all agents during the last five seconds | [optional] 
**DiscardedCount** | Pointer to **float32** | Number of discarded events received from agents during the last five seconds | [optional] 
**EvtCount** | Pointer to **float32** | Number of events sent to analysisd during the last five seconds | [optional] 
**SentBytes** | Pointer to **float32** | Number of sent bytes to the agents during the last five seconds | [optional] 
**QueueSize** | Pointer to **float32** | Usage of the queue to storage events from agents | [optional] 
**RecvBytes** | Pointer to **float32** | Number of received bytes from all agents during the last five seconds | [optional] 
**TcpSessions** | Pointer to **float32** | Number of TCP active sessions during the last five seconds | [optional] 
**TotalQueueSize** | Pointer to **float32** | Total queue size to store events from agents | [optional] 

## Methods

### NewAllItemsResponseWazuhStatsAllOfAffectedItems

`func NewAllItemsResponseWazuhStatsAllOfAffectedItems() *AllItemsResponseWazuhStatsAllOfAffectedItems`

NewAllItemsResponseWazuhStatsAllOfAffectedItems instantiates a new AllItemsResponseWazuhStatsAllOfAffectedItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllItemsResponseWazuhStatsAllOfAffectedItemsWithDefaults

`func NewAllItemsResponseWazuhStatsAllOfAffectedItemsWithDefaults() *AllItemsResponseWazuhStatsAllOfAffectedItems`

NewAllItemsResponseWazuhStatsAllOfAffectedItemsWithDefaults instantiates a new AllItemsResponseWazuhStatsAllOfAffectedItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAverages

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetAverages() []int32`

GetAverages returns the Averages field if non-nil, zero value otherwise.

### GetAveragesOk

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetAveragesOk() (*[]int32, bool)`

GetAveragesOk returns a tuple with the Averages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverages

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) SetAverages(v []int32)`

SetAverages sets Averages field to given value.

### HasAverages

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) HasAverages() bool`

HasAverages returns a boolean if a field has been set.

### GetInteractions

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetInteractions() int32`

GetInteractions returns the Interactions field if non-nil, zero value otherwise.

### GetInteractionsOk

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetInteractionsOk() (*int32, bool)`

GetInteractionsOk returns a tuple with the Interactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteractions

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) SetInteractions(v int32)`

SetInteractions sets Interactions field to given value.

### HasInteractions

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) HasInteractions() bool`

HasInteractions returns a boolean if a field has been set.

### GetSun

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetSun() WazuhHourlyStats`

GetSun returns the Sun field if non-nil, zero value otherwise.

### GetSunOk

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetSunOk() (*WazuhHourlyStats, bool)`

GetSunOk returns a tuple with the Sun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSun

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) SetSun(v WazuhHourlyStats)`

SetSun sets Sun field to given value.

### HasSun

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) HasSun() bool`

HasSun returns a boolean if a field has been set.

### GetMon

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetMon() WazuhHourlyStats`

GetMon returns the Mon field if non-nil, zero value otherwise.

### GetMonOk

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetMonOk() (*WazuhHourlyStats, bool)`

GetMonOk returns a tuple with the Mon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMon

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) SetMon(v WazuhHourlyStats)`

SetMon sets Mon field to given value.

### HasMon

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) HasMon() bool`

HasMon returns a boolean if a field has been set.

### GetTue

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetTue() WazuhHourlyStats`

GetTue returns the Tue field if non-nil, zero value otherwise.

### GetTueOk

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetTueOk() (*WazuhHourlyStats, bool)`

GetTueOk returns a tuple with the Tue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTue

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) SetTue(v WazuhHourlyStats)`

SetTue sets Tue field to given value.

### HasTue

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) HasTue() bool`

HasTue returns a boolean if a field has been set.

### GetWed

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetWed() WazuhHourlyStats`

GetWed returns the Wed field if non-nil, zero value otherwise.

### GetWedOk

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetWedOk() (*WazuhHourlyStats, bool)`

GetWedOk returns a tuple with the Wed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWed

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) SetWed(v WazuhHourlyStats)`

SetWed sets Wed field to given value.

### HasWed

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) HasWed() bool`

HasWed returns a boolean if a field has been set.

### GetThu

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetThu() WazuhHourlyStats`

GetThu returns the Thu field if non-nil, zero value otherwise.

### GetThuOk

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetThuOk() (*WazuhHourlyStats, bool)`

GetThuOk returns a tuple with the Thu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThu

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) SetThu(v WazuhHourlyStats)`

SetThu sets Thu field to given value.

### HasThu

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) HasThu() bool`

HasThu returns a boolean if a field has been set.

### GetFri

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetFri() WazuhHourlyStats`

GetFri returns the Fri field if non-nil, zero value otherwise.

### GetFriOk

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetFriOk() (*WazuhHourlyStats, bool)`

GetFriOk returns a tuple with the Fri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFri

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) SetFri(v WazuhHourlyStats)`

SetFri sets Fri field to given value.

### HasFri

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) HasFri() bool`

HasFri returns a boolean if a field has been set.

### GetSat

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetSat() WazuhHourlyStats`

GetSat returns the Sat field if non-nil, zero value otherwise.

### GetSatOk

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetSatOk() (*WazuhHourlyStats, bool)`

GetSatOk returns a tuple with the Sat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSat

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) SetSat(v WazuhHourlyStats)`

SetSat sets Sat field to given value.

### HasSat

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) HasSat() bool`

HasSat returns a boolean if a field has been set.

### GetAlertsQueueSize

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetAlertsQueueSize() float32`

GetAlertsQueueSize returns the AlertsQueueSize field if non-nil, zero value otherwise.

### GetAlertsQueueSizeOk

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetAlertsQueueSizeOk() (*float32, bool)`

GetAlertsQueueSizeOk returns a tuple with the AlertsQueueSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertsQueueSize

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) SetAlertsQueueSize(v float32)`

SetAlertsQueueSize sets AlertsQueueSize field to given value.

### HasAlertsQueueSize

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) HasAlertsQueueSize() bool`

HasAlertsQueueSize returns a boolean if a field has been set.

### GetAlertsQueueUsage

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetAlertsQueueUsage() float32`

GetAlertsQueueUsage returns the AlertsQueueUsage field if non-nil, zero value otherwise.

### GetAlertsQueueUsageOk

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetAlertsQueueUsageOk() (*float32, bool)`

GetAlertsQueueUsageOk returns a tuple with the AlertsQueueUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertsQueueUsage

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) SetAlertsQueueUsage(v float32)`

SetAlertsQueueUsage sets AlertsQueueUsage field to given value.

### HasAlertsQueueUsage

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) HasAlertsQueueUsage() bool`

HasAlertsQueueUsage returns a boolean if a field has been set.

### GetAlertsWritten

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetAlertsWritten() float32`

GetAlertsWritten returns the AlertsWritten field if non-nil, zero value otherwise.

### GetAlertsWrittenOk

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetAlertsWrittenOk() (*float32, bool)`

GetAlertsWrittenOk returns a tuple with the AlertsWritten field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertsWritten

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) SetAlertsWritten(v float32)`

SetAlertsWritten sets AlertsWritten field to given value.

### HasAlertsWritten

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) HasAlertsWritten() bool`

HasAlertsWritten returns a boolean if a field has been set.

### GetArchivesQueueSize

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetArchivesQueueSize() float32`

GetArchivesQueueSize returns the ArchivesQueueSize field if non-nil, zero value otherwise.

### GetArchivesQueueSizeOk

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetArchivesQueueSizeOk() (*float32, bool)`

GetArchivesQueueSizeOk returns a tuple with the ArchivesQueueSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivesQueueSize

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) SetArchivesQueueSize(v float32)`

SetArchivesQueueSize sets ArchivesQueueSize field to given value.

### HasArchivesQueueSize

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) HasArchivesQueueSize() bool`

HasArchivesQueueSize returns a boolean if a field has been set.

### GetArchivesQueueUsage

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetArchivesQueueUsage() float32`

GetArchivesQueueUsage returns the ArchivesQueueUsage field if non-nil, zero value otherwise.

### GetArchivesQueueUsageOk

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetArchivesQueueUsageOk() (*float32, bool)`

GetArchivesQueueUsageOk returns a tuple with the ArchivesQueueUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivesQueueUsage

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) SetArchivesQueueUsage(v float32)`

SetArchivesQueueUsage sets ArchivesQueueUsage field to given value.

### HasArchivesQueueUsage

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) HasArchivesQueueUsage() bool`

HasArchivesQueueUsage returns a boolean if a field has been set.

### GetEventQueueSize

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetEventQueueSize() float32`

GetEventQueueSize returns the EventQueueSize field if non-nil, zero value otherwise.

### GetEventQueueSizeOk

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetEventQueueSizeOk() (*float32, bool)`

GetEventQueueSizeOk returns a tuple with the EventQueueSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventQueueSize

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) SetEventQueueSize(v float32)`

SetEventQueueSize sets EventQueueSize field to given value.

### HasEventQueueSize

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) HasEventQueueSize() bool`

HasEventQueueSize returns a boolean if a field has been set.

### GetEventQueueUsage

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetEventQueueUsage() float32`

GetEventQueueUsage returns the EventQueueUsage field if non-nil, zero value otherwise.

### GetEventQueueUsageOk

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetEventQueueUsageOk() (*float32, bool)`

GetEventQueueUsageOk returns a tuple with the EventQueueUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventQueueUsage

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) SetEventQueueUsage(v float32)`

SetEventQueueUsage sets EventQueueUsage field to given value.

### HasEventQueueUsage

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) HasEventQueueUsage() bool`

HasEventQueueUsage returns a boolean if a field has been set.

### GetEventsDropped

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetEventsDropped() float32`

GetEventsDropped returns the EventsDropped field if non-nil, zero value otherwise.

### GetEventsDroppedOk

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetEventsDroppedOk() (*float32, bool)`

GetEventsDroppedOk returns a tuple with the EventsDropped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsDropped

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) SetEventsDropped(v float32)`

SetEventsDropped sets EventsDropped field to given value.

### HasEventsDropped

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) HasEventsDropped() bool`

HasEventsDropped returns a boolean if a field has been set.

### GetEventsProcessed

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetEventsProcessed() float32`

GetEventsProcessed returns the EventsProcessed field if non-nil, zero value otherwise.

### GetEventsProcessedOk

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetEventsProcessedOk() (*float32, bool)`

GetEventsProcessedOk returns a tuple with the EventsProcessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsProcessed

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) SetEventsProcessed(v float32)`

SetEventsProcessed sets EventsProcessed field to given value.

### HasEventsProcessed

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) HasEventsProcessed() bool`

HasEventsProcessed returns a boolean if a field has been set.

### GetEventsReceived

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetEventsReceived() float32`

GetEventsReceived returns the EventsReceived field if non-nil, zero value otherwise.

### GetEventsReceivedOk

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetEventsReceivedOk() (*float32, bool)`

GetEventsReceivedOk returns a tuple with the EventsReceived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsReceived

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) SetEventsReceived(v float32)`

SetEventsReceived sets EventsReceived field to given value.

### HasEventsReceived

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) HasEventsReceived() bool`

HasEventsReceived returns a boolean if a field has been set.

### GetFirewallQueueSize

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetFirewallQueueSize() float32`

GetFirewallQueueSize returns the FirewallQueueSize field if non-nil, zero value otherwise.

### GetFirewallQueueSizeOk

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetFirewallQueueSizeOk() (*float32, bool)`

GetFirewallQueueSizeOk returns a tuple with the FirewallQueueSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallQueueSize

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) SetFirewallQueueSize(v float32)`

SetFirewallQueueSize sets FirewallQueueSize field to given value.

### HasFirewallQueueSize

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) HasFirewallQueueSize() bool`

HasFirewallQueueSize returns a boolean if a field has been set.

### GetFirewallQueueUsage

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetFirewallQueueUsage() float32`

GetFirewallQueueUsage returns the FirewallQueueUsage field if non-nil, zero value otherwise.

### GetFirewallQueueUsageOk

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetFirewallQueueUsageOk() (*float32, bool)`

GetFirewallQueueUsageOk returns a tuple with the FirewallQueueUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallQueueUsage

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) SetFirewallQueueUsage(v float32)`

SetFirewallQueueUsage sets FirewallQueueUsage field to given value.

### HasFirewallQueueUsage

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) HasFirewallQueueUsage() bool`

HasFirewallQueueUsage returns a boolean if a field has been set.

### GetFirewallWritten

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetFirewallWritten() float32`

GetFirewallWritten returns the FirewallWritten field if non-nil, zero value otherwise.

### GetFirewallWrittenOk

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetFirewallWrittenOk() (*float32, bool)`

GetFirewallWrittenOk returns a tuple with the FirewallWritten field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallWritten

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) SetFirewallWritten(v float32)`

SetFirewallWritten sets FirewallWritten field to given value.

### HasFirewallWritten

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) HasFirewallWritten() bool`

HasFirewallWritten returns a boolean if a field has been set.

### GetFtsWritten

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetFtsWritten() float32`

GetFtsWritten returns the FtsWritten field if non-nil, zero value otherwise.

### GetFtsWrittenOk

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetFtsWrittenOk() (*float32, bool)`

GetFtsWrittenOk returns a tuple with the FtsWritten field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFtsWritten

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) SetFtsWritten(v float32)`

SetFtsWritten sets FtsWritten field to given value.

### HasFtsWritten

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) HasFtsWritten() bool`

HasFtsWritten returns a boolean if a field has been set.

### GetHostinfoQueueSize

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetHostinfoQueueSize() float32`

GetHostinfoQueueSize returns the HostinfoQueueSize field if non-nil, zero value otherwise.

### GetHostinfoQueueSizeOk

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetHostinfoQueueSizeOk() (*float32, bool)`

GetHostinfoQueueSizeOk returns a tuple with the HostinfoQueueSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostinfoQueueSize

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) SetHostinfoQueueSize(v float32)`

SetHostinfoQueueSize sets HostinfoQueueSize field to given value.

### HasHostinfoQueueSize

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) HasHostinfoQueueSize() bool`

HasHostinfoQueueSize returns a boolean if a field has been set.

### GetHostinfoQueueUsage

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetHostinfoQueueUsage() float32`

GetHostinfoQueueUsage returns the HostinfoQueueUsage field if non-nil, zero value otherwise.

### GetHostinfoQueueUsageOk

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetHostinfoQueueUsageOk() (*float32, bool)`

GetHostinfoQueueUsageOk returns a tuple with the HostinfoQueueUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostinfoQueueUsage

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) SetHostinfoQueueUsage(v float32)`

SetHostinfoQueueUsage sets HostinfoQueueUsage field to given value.

### HasHostinfoQueueUsage

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) HasHostinfoQueueUsage() bool`

HasHostinfoQueueUsage returns a boolean if a field has been set.

### GetOtherEventsDecoded

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetOtherEventsDecoded() float32`

GetOtherEventsDecoded returns the OtherEventsDecoded field if non-nil, zero value otherwise.

### GetOtherEventsDecodedOk

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetOtherEventsDecodedOk() (*float32, bool)`

GetOtherEventsDecodedOk returns a tuple with the OtherEventsDecoded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherEventsDecoded

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) SetOtherEventsDecoded(v float32)`

SetOtherEventsDecoded sets OtherEventsDecoded field to given value.

### HasOtherEventsDecoded

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) HasOtherEventsDecoded() bool`

HasOtherEventsDecoded returns a boolean if a field has been set.

### GetRootcheckEventsDecoded

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetRootcheckEventsDecoded() float32`

GetRootcheckEventsDecoded returns the RootcheckEventsDecoded field if non-nil, zero value otherwise.

### GetRootcheckEventsDecodedOk

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetRootcheckEventsDecodedOk() (*float32, bool)`

GetRootcheckEventsDecodedOk returns a tuple with the RootcheckEventsDecoded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootcheckEventsDecoded

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) SetRootcheckEventsDecoded(v float32)`

SetRootcheckEventsDecoded sets RootcheckEventsDecoded field to given value.

### HasRootcheckEventsDecoded

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) HasRootcheckEventsDecoded() bool`

HasRootcheckEventsDecoded returns a boolean if a field has been set.

### GetRootcheckQueueSize

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetRootcheckQueueSize() float32`

GetRootcheckQueueSize returns the RootcheckQueueSize field if non-nil, zero value otherwise.

### GetRootcheckQueueSizeOk

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetRootcheckQueueSizeOk() (*float32, bool)`

GetRootcheckQueueSizeOk returns a tuple with the RootcheckQueueSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootcheckQueueSize

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) SetRootcheckQueueSize(v float32)`

SetRootcheckQueueSize sets RootcheckQueueSize field to given value.

### HasRootcheckQueueSize

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) HasRootcheckQueueSize() bool`

HasRootcheckQueueSize returns a boolean if a field has been set.

### GetRootcheckQueueUsage

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetRootcheckQueueUsage() float32`

GetRootcheckQueueUsage returns the RootcheckQueueUsage field if non-nil, zero value otherwise.

### GetRootcheckQueueUsageOk

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetRootcheckQueueUsageOk() (*float32, bool)`

GetRootcheckQueueUsageOk returns a tuple with the RootcheckQueueUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootcheckQueueUsage

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) SetRootcheckQueueUsage(v float32)`

SetRootcheckQueueUsage sets RootcheckQueueUsage field to given value.

### HasRootcheckQueueUsage

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) HasRootcheckQueueUsage() bool`

HasRootcheckQueueUsage returns a boolean if a field has been set.

### GetRuleMatchingQueueSize

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetRuleMatchingQueueSize() float32`

GetRuleMatchingQueueSize returns the RuleMatchingQueueSize field if non-nil, zero value otherwise.

### GetRuleMatchingQueueSizeOk

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetRuleMatchingQueueSizeOk() (*float32, bool)`

GetRuleMatchingQueueSizeOk returns a tuple with the RuleMatchingQueueSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleMatchingQueueSize

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) SetRuleMatchingQueueSize(v float32)`

SetRuleMatchingQueueSize sets RuleMatchingQueueSize field to given value.

### HasRuleMatchingQueueSize

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) HasRuleMatchingQueueSize() bool`

HasRuleMatchingQueueSize returns a boolean if a field has been set.

### GetRuleMatchingQueueUsage

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetRuleMatchingQueueUsage() float32`

GetRuleMatchingQueueUsage returns the RuleMatchingQueueUsage field if non-nil, zero value otherwise.

### GetRuleMatchingQueueUsageOk

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetRuleMatchingQueueUsageOk() (*float32, bool)`

GetRuleMatchingQueueUsageOk returns a tuple with the RuleMatchingQueueUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleMatchingQueueUsage

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) SetRuleMatchingQueueUsage(v float32)`

SetRuleMatchingQueueUsage sets RuleMatchingQueueUsage field to given value.

### HasRuleMatchingQueueUsage

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) HasRuleMatchingQueueUsage() bool`

HasRuleMatchingQueueUsage returns a boolean if a field has been set.

### GetScaEventsDecoded

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetScaEventsDecoded() float32`

GetScaEventsDecoded returns the ScaEventsDecoded field if non-nil, zero value otherwise.

### GetScaEventsDecodedOk

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetScaEventsDecodedOk() (*float32, bool)`

GetScaEventsDecodedOk returns a tuple with the ScaEventsDecoded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaEventsDecoded

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) SetScaEventsDecoded(v float32)`

SetScaEventsDecoded sets ScaEventsDecoded field to given value.

### HasScaEventsDecoded

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) HasScaEventsDecoded() bool`

HasScaEventsDecoded returns a boolean if a field has been set.

### GetScaQueueSize

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetScaQueueSize() float32`

GetScaQueueSize returns the ScaQueueSize field if non-nil, zero value otherwise.

### GetScaQueueSizeOk

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetScaQueueSizeOk() (*float32, bool)`

GetScaQueueSizeOk returns a tuple with the ScaQueueSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaQueueSize

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) SetScaQueueSize(v float32)`

SetScaQueueSize sets ScaQueueSize field to given value.

### HasScaQueueSize

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) HasScaQueueSize() bool`

HasScaQueueSize returns a boolean if a field has been set.

### GetScaQueueUsage

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetScaQueueUsage() float32`

GetScaQueueUsage returns the ScaQueueUsage field if non-nil, zero value otherwise.

### GetScaQueueUsageOk

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetScaQueueUsageOk() (*float32, bool)`

GetScaQueueUsageOk returns a tuple with the ScaQueueUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaQueueUsage

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) SetScaQueueUsage(v float32)`

SetScaQueueUsage sets ScaQueueUsage field to given value.

### HasScaQueueUsage

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) HasScaQueueUsage() bool`

HasScaQueueUsage returns a boolean if a field has been set.

### GetStatisticalQueueSize

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetStatisticalQueueSize() float32`

GetStatisticalQueueSize returns the StatisticalQueueSize field if non-nil, zero value otherwise.

### GetStatisticalQueueSizeOk

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetStatisticalQueueSizeOk() (*float32, bool)`

GetStatisticalQueueSizeOk returns a tuple with the StatisticalQueueSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatisticalQueueSize

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) SetStatisticalQueueSize(v float32)`

SetStatisticalQueueSize sets StatisticalQueueSize field to given value.

### HasStatisticalQueueSize

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) HasStatisticalQueueSize() bool`

HasStatisticalQueueSize returns a boolean if a field has been set.

### GetStatisticalQueueUsage

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetStatisticalQueueUsage() float32`

GetStatisticalQueueUsage returns the StatisticalQueueUsage field if non-nil, zero value otherwise.

### GetStatisticalQueueUsageOk

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetStatisticalQueueUsageOk() (*float32, bool)`

GetStatisticalQueueUsageOk returns a tuple with the StatisticalQueueUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatisticalQueueUsage

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) SetStatisticalQueueUsage(v float32)`

SetStatisticalQueueUsage sets StatisticalQueueUsage field to given value.

### HasStatisticalQueueUsage

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) HasStatisticalQueueUsage() bool`

HasStatisticalQueueUsage returns a boolean if a field has been set.

### GetSyscheckEventsDecoded

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetSyscheckEventsDecoded() float32`

GetSyscheckEventsDecoded returns the SyscheckEventsDecoded field if non-nil, zero value otherwise.

### GetSyscheckEventsDecodedOk

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetSyscheckEventsDecodedOk() (*float32, bool)`

GetSyscheckEventsDecodedOk returns a tuple with the SyscheckEventsDecoded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyscheckEventsDecoded

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) SetSyscheckEventsDecoded(v float32)`

SetSyscheckEventsDecoded sets SyscheckEventsDecoded field to given value.

### HasSyscheckEventsDecoded

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) HasSyscheckEventsDecoded() bool`

HasSyscheckEventsDecoded returns a boolean if a field has been set.

### GetSyscheckQueueSize

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetSyscheckQueueSize() float32`

GetSyscheckQueueSize returns the SyscheckQueueSize field if non-nil, zero value otherwise.

### GetSyscheckQueueSizeOk

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetSyscheckQueueSizeOk() (*float32, bool)`

GetSyscheckQueueSizeOk returns a tuple with the SyscheckQueueSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyscheckQueueSize

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) SetSyscheckQueueSize(v float32)`

SetSyscheckQueueSize sets SyscheckQueueSize field to given value.

### HasSyscheckQueueSize

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) HasSyscheckQueueSize() bool`

HasSyscheckQueueSize returns a boolean if a field has been set.

### GetSyscheckQueueUsage

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetSyscheckQueueUsage() float32`

GetSyscheckQueueUsage returns the SyscheckQueueUsage field if non-nil, zero value otherwise.

### GetSyscheckQueueUsageOk

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetSyscheckQueueUsageOk() (*float32, bool)`

GetSyscheckQueueUsageOk returns a tuple with the SyscheckQueueUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyscheckQueueUsage

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) SetSyscheckQueueUsage(v float32)`

SetSyscheckQueueUsage sets SyscheckQueueUsage field to given value.

### HasSyscheckQueueUsage

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) HasSyscheckQueueUsage() bool`

HasSyscheckQueueUsage returns a boolean if a field has been set.

### GetSyscollectorEventsDecoded

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetSyscollectorEventsDecoded() float32`

GetSyscollectorEventsDecoded returns the SyscollectorEventsDecoded field if non-nil, zero value otherwise.

### GetSyscollectorEventsDecodedOk

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetSyscollectorEventsDecodedOk() (*float32, bool)`

GetSyscollectorEventsDecodedOk returns a tuple with the SyscollectorEventsDecoded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyscollectorEventsDecoded

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) SetSyscollectorEventsDecoded(v float32)`

SetSyscollectorEventsDecoded sets SyscollectorEventsDecoded field to given value.

### HasSyscollectorEventsDecoded

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) HasSyscollectorEventsDecoded() bool`

HasSyscollectorEventsDecoded returns a boolean if a field has been set.

### GetSyscollectorQueueSize

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetSyscollectorQueueSize() float32`

GetSyscollectorQueueSize returns the SyscollectorQueueSize field if non-nil, zero value otherwise.

### GetSyscollectorQueueSizeOk

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetSyscollectorQueueSizeOk() (*float32, bool)`

GetSyscollectorQueueSizeOk returns a tuple with the SyscollectorQueueSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyscollectorQueueSize

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) SetSyscollectorQueueSize(v float32)`

SetSyscollectorQueueSize sets SyscollectorQueueSize field to given value.

### HasSyscollectorQueueSize

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) HasSyscollectorQueueSize() bool`

HasSyscollectorQueueSize returns a boolean if a field has been set.

### GetSyscollectorQueueUsage

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetSyscollectorQueueUsage() float32`

GetSyscollectorQueueUsage returns the SyscollectorQueueUsage field if non-nil, zero value otherwise.

### GetSyscollectorQueueUsageOk

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetSyscollectorQueueUsageOk() (*float32, bool)`

GetSyscollectorQueueUsageOk returns a tuple with the SyscollectorQueueUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyscollectorQueueUsage

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) SetSyscollectorQueueUsage(v float32)`

SetSyscollectorQueueUsage sets SyscollectorQueueUsage field to given value.

### HasSyscollectorQueueUsage

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) HasSyscollectorQueueUsage() bool`

HasSyscollectorQueueUsage returns a boolean if a field has been set.

### GetTotalEventsDecoded

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetTotalEventsDecoded() float32`

GetTotalEventsDecoded returns the TotalEventsDecoded field if non-nil, zero value otherwise.

### GetTotalEventsDecodedOk

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetTotalEventsDecodedOk() (*float32, bool)`

GetTotalEventsDecodedOk returns a tuple with the TotalEventsDecoded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEventsDecoded

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) SetTotalEventsDecoded(v float32)`

SetTotalEventsDecoded sets TotalEventsDecoded field to given value.

### HasTotalEventsDecoded

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) HasTotalEventsDecoded() bool`

HasTotalEventsDecoded returns a boolean if a field has been set.

### GetWinevtEventsDecoded

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetWinevtEventsDecoded() float32`

GetWinevtEventsDecoded returns the WinevtEventsDecoded field if non-nil, zero value otherwise.

### GetWinevtEventsDecodedOk

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetWinevtEventsDecodedOk() (*float32, bool)`

GetWinevtEventsDecodedOk returns a tuple with the WinevtEventsDecoded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinevtEventsDecoded

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) SetWinevtEventsDecoded(v float32)`

SetWinevtEventsDecoded sets WinevtEventsDecoded field to given value.

### HasWinevtEventsDecoded

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) HasWinevtEventsDecoded() bool`

HasWinevtEventsDecoded returns a boolean if a field has been set.

### GetWinevtQueueSize

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetWinevtQueueSize() float32`

GetWinevtQueueSize returns the WinevtQueueSize field if non-nil, zero value otherwise.

### GetWinevtQueueSizeOk

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetWinevtQueueSizeOk() (*float32, bool)`

GetWinevtQueueSizeOk returns a tuple with the WinevtQueueSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinevtQueueSize

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) SetWinevtQueueSize(v float32)`

SetWinevtQueueSize sets WinevtQueueSize field to given value.

### HasWinevtQueueSize

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) HasWinevtQueueSize() bool`

HasWinevtQueueSize returns a boolean if a field has been set.

### GetWinevtQueueUsage

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetWinevtQueueUsage() float32`

GetWinevtQueueUsage returns the WinevtQueueUsage field if non-nil, zero value otherwise.

### GetWinevtQueueUsageOk

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetWinevtQueueUsageOk() (*float32, bool)`

GetWinevtQueueUsageOk returns a tuple with the WinevtQueueUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinevtQueueUsage

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) SetWinevtQueueUsage(v float32)`

SetWinevtQueueUsage sets WinevtQueueUsage field to given value.

### HasWinevtQueueUsage

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) HasWinevtQueueUsage() bool`

HasWinevtQueueUsage returns a boolean if a field has been set.

### GetCtrlMsgCount

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetCtrlMsgCount() float32`

GetCtrlMsgCount returns the CtrlMsgCount field if non-nil, zero value otherwise.

### GetCtrlMsgCountOk

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetCtrlMsgCountOk() (*float32, bool)`

GetCtrlMsgCountOk returns a tuple with the CtrlMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCtrlMsgCount

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) SetCtrlMsgCount(v float32)`

SetCtrlMsgCount sets CtrlMsgCount field to given value.

### HasCtrlMsgCount

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) HasCtrlMsgCount() bool`

HasCtrlMsgCount returns a boolean if a field has been set.

### GetDiscardedCount

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetDiscardedCount() float32`

GetDiscardedCount returns the DiscardedCount field if non-nil, zero value otherwise.

### GetDiscardedCountOk

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetDiscardedCountOk() (*float32, bool)`

GetDiscardedCountOk returns a tuple with the DiscardedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscardedCount

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) SetDiscardedCount(v float32)`

SetDiscardedCount sets DiscardedCount field to given value.

### HasDiscardedCount

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) HasDiscardedCount() bool`

HasDiscardedCount returns a boolean if a field has been set.

### GetEvtCount

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetEvtCount() float32`

GetEvtCount returns the EvtCount field if non-nil, zero value otherwise.

### GetEvtCountOk

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetEvtCountOk() (*float32, bool)`

GetEvtCountOk returns a tuple with the EvtCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvtCount

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) SetEvtCount(v float32)`

SetEvtCount sets EvtCount field to given value.

### HasEvtCount

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) HasEvtCount() bool`

HasEvtCount returns a boolean if a field has been set.

### GetSentBytes

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetSentBytes() float32`

GetSentBytes returns the SentBytes field if non-nil, zero value otherwise.

### GetSentBytesOk

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetSentBytesOk() (*float32, bool)`

GetSentBytesOk returns a tuple with the SentBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentBytes

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) SetSentBytes(v float32)`

SetSentBytes sets SentBytes field to given value.

### HasSentBytes

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) HasSentBytes() bool`

HasSentBytes returns a boolean if a field has been set.

### GetQueueSize

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetQueueSize() float32`

GetQueueSize returns the QueueSize field if non-nil, zero value otherwise.

### GetQueueSizeOk

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetQueueSizeOk() (*float32, bool)`

GetQueueSizeOk returns a tuple with the QueueSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueSize

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) SetQueueSize(v float32)`

SetQueueSize sets QueueSize field to given value.

### HasQueueSize

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) HasQueueSize() bool`

HasQueueSize returns a boolean if a field has been set.

### GetRecvBytes

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetRecvBytes() float32`

GetRecvBytes returns the RecvBytes field if non-nil, zero value otherwise.

### GetRecvBytesOk

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetRecvBytesOk() (*float32, bool)`

GetRecvBytesOk returns a tuple with the RecvBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecvBytes

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) SetRecvBytes(v float32)`

SetRecvBytes sets RecvBytes field to given value.

### HasRecvBytes

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) HasRecvBytes() bool`

HasRecvBytes returns a boolean if a field has been set.

### GetTcpSessions

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetTcpSessions() float32`

GetTcpSessions returns the TcpSessions field if non-nil, zero value otherwise.

### GetTcpSessionsOk

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetTcpSessionsOk() (*float32, bool)`

GetTcpSessionsOk returns a tuple with the TcpSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpSessions

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) SetTcpSessions(v float32)`

SetTcpSessions sets TcpSessions field to given value.

### HasTcpSessions

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) HasTcpSessions() bool`

HasTcpSessions returns a boolean if a field has been set.

### GetTotalQueueSize

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetTotalQueueSize() float32`

GetTotalQueueSize returns the TotalQueueSize field if non-nil, zero value otherwise.

### GetTotalQueueSizeOk

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) GetTotalQueueSizeOk() (*float32, bool)`

GetTotalQueueSizeOk returns a tuple with the TotalQueueSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalQueueSize

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) SetTotalQueueSize(v float32)`

SetTotalQueueSize sets TotalQueueSize field to given value.

### HasTotalQueueSize

`func (o *AllItemsResponseWazuhStatsAllOfAffectedItems) HasTotalQueueSize() bool`

HasTotalQueueSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


