package option

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

const (
	Configuration           = "configuration"
	DatabaseUrl             = "database-url"
	Datasources             = "datasources"
	EnableG2config          = "enable-g2config"
	EnableG2configmgr       = "enable-g2configmgr"
	EnableG2diagnostic      = "enable-g2diagnostic"
	EnableG2engine          = "enable-g2engine"
	EnableG2product         = "enable-g2product"
	EngineConfigurationJson = "engine-configuration-json"
	EngineLogLevel          = "engine-log-level"
	EngineModuleName        = "engine-module-name"
	GrpcPort                = "grpc-port"
	InputFileType           = "input-file-type"
	InputURL                = "input-url"
	LogLevel                = "log-level"
	OutputURL               = "output-url"
)

// ----------------------------------------------------------------------------
// Constant help text for parameters
// ----------------------------------------------------------------------------

const (
	LogLevelHelp = "Log level of TRACE, DEBUG, INFO, WARN, ERROR, FATAL, or PANIC [%s]"
)
