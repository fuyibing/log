// author: wsfuyibing <websearch@163.com>
// date: 2023-02-18

package conf

func (o *configuration) GetAdapter() string           { return o.Adapter }
func (o *configuration) GetAutoStart() bool           { return o.AutoStart }
func (o *configuration) GetFile() FileConfiguration   { return o.File }
func (o *configuration) GetKafka() KafkaConfiguration { return o.Kafka }
func (o *configuration) GetLevel() Level              { return o.Level }
func (o *configuration) GetPid() int                  { return o.pid }
func (o *configuration) GetPrefix() string            { return o.Prefix }
func (o *configuration) GetServiceHost() string       { return o.ServiceHost }
func (o *configuration) GetServiceName() string       { return o.ServiceName }
func (o *configuration) GetServicePort() int          { return o.ServicePort }
func (o *configuration) GetTerm() TermConfiguration   { return o.Term }
func (o *configuration) GetTimeFormat() string        { return o.TimeFormat }

// Batch mode.

func (o *configuration) GetBatchConcurrency() int32 { return o.BatchConcurrency }
func (o *configuration) GetBatchLimit() int         { return o.BatchLimit }

// Open tracing.

func (o *configuration) GetSpanId() string       { return o.SpanId }
func (o *configuration) GetParentSpanId() string { return o.ParentSpanId }
func (o *configuration) GetTraceId() string      { return o.TraceId }
func (o *configuration) GetTraceVersion() string { return o.TraceVersion }

// Switch status.

func (o *configuration) DebugOn() bool { return o.debugOn }
func (o *configuration) InfoOn() bool  { return o.infoOn }
func (o *configuration) WarnOn() bool  { return o.warnOn }
func (o *configuration) ErrorOn() bool { return o.errorOn }
func (o *configuration) FatalOn() bool { return o.fatalOn }
