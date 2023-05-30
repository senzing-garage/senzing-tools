package option

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

const (
	Configuration             = "configuration"
	DatabaseUrl               = "database-url"
	Datasources               = "datasources"
	DelayInSeconds            = "delay-in-seconds"
	EnableG2config            = "enable-g2config"
	EnableG2configmgr         = "enable-g2configmgr"
	EnableG2diagnostic        = "enable-g2diagnostic"
	EnableG2engine            = "enable-g2engine"
	EnableG2product           = "enable-g2product"
	EnableSwaggerUi           = "enable-swagger-ui"
	EngineConfigurationJson   = "engine-configuration-json"
	EngineLogLevel            = "engine-log-level"
	EngineModuleName          = "engine-module-name"
	GrpcPort                  = "grpc-port"
	GrpcUrl                   = "grpc-url"
	HttpPort                  = "http-port"
	InputFileType             = "input-file-type"
	InputURL                  = "input-url"
	LogLevel                  = "log-level"
	MonitoringPeriodInSeconds = "monitoring-period-in-seconds"
	NumberOfWorkers           = "number-of-workers"
	ObserverGrpcPort          = "observer-grpc-port"
	ObserverOrigin            = "observer-origin"
	ObserverUrl               = "observer-url"
	OutputURL                 = "output-url"
	RecordMax                 = "record-max"
	RecordMin                 = "record-min"
	RecordMonitor             = "record-monitor"
)

// ----------------------------------------------------------------------------
// Constant help text for parameters
// ----------------------------------------------------------------------------

const (
	DelayInSecondsHelp            = "Number of seconds to wait before starting process"
	InputFileTypeHelp             = "Input file type to override auto-detect based on file name"
	InputURLHelp                  = "Input URL used for processing"
	LogLevelHelp                  = "Log level of TRACE, DEBUG, INFO, WARN, ERROR, FATAL, or PANIC [%s]"
	MonitoringPeriodInSecondsHelp = "Print monitoring log messages with the period given in seconds"
	NumberOfWorkersHelp           = "Override the default number of worker routines."
	OutputURLHelp                 = "Output URL used for processing"
	RecordMaxHelp                 = "Process a maximum number of records equal to this number"
	RecordMinHelp                 = "Process records starting at this record number, discarding all before"
	RecordMonitorHelp             = "Log a monitor message after this number of records have been processed"
)
