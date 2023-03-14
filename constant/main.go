package constant

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
	LogLevel                = "log-level"
	SetEnvPrefix            = "SENZING_TOOLS"
	VersionTemplate         = `{{printf "%s: %s - version %s\n" .Name .Short .Version}}`
)
