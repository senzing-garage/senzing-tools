package cmdhelper

import (
	"fmt"
	"time"

	"github.com/senzing/senzing-tools/cmdhelper/optiontype"
	"github.com/senzing/senzing-tools/envar"
	"github.com/senzing/senzing-tools/help"
	"github.com/senzing/senzing-tools/option"
)

func (v ContextVariable) SetDefault(newDefault any) ContextVariable {
	v.Default = newDefault
	return v
}

var OptionConfigPath = ContextVariable{
	Default: OsLookupEnvString(envar.ConfigPath, ""),
	Envar:   envar.ConfigPath,
	Help:    help.ConfigPath,
	Arg:     option.ConfigPath,
	Type:    optiontype.String,
}

var OptionConfiguration = ContextVariable{
	Default: OsLookupEnvString(envar.Configuration, ""),
	Envar:   envar.Configuration,
	Help:    help.Configuration,
	Arg:     option.Configuration,
	Type:    optiontype.String,
}

var OptionDatabaseUrl = ContextVariable{
	Default: OsLookupEnvString(envar.DatabaseUrl, ""),
	Envar:   envar.DatabaseUrl,
	Help:    help.DatabaseUrl,
	Arg:     option.DatabaseUrl,
	Type:    optiontype.String,
}

var OptionDatasources = ContextVariable{
	Default: []string{},
	Envar:   envar.Datasources,
	Help:    help.Datasources,
	Arg:     option.Datasources,
	Type:    optiontype.StringSlice,
}

var OptionDelayInSeconds = ContextVariable{
	Default: 0,
	Envar:   envar.DelayInSeconds,
	Help:    help.DelayInSeconds,
	Arg:     option.DelayInSeconds,
	Type:    optiontype.Int,
}

var OptionEnableAll = ContextVariable{
	Default: OsLookupEnvBool(envar.EnableAll, false),
	Envar:   envar.EnableAll,
	Help:    help.EnableAll,
	Arg:     option.EnableAll,
	Type:    optiontype.Bool,
}

var OptionEnableG2config = ContextVariable{
	Default: OsLookupEnvBool(envar.EnableG2config, false),
	Envar:   envar.EnableG2config,
	Help:    help.EnableG2config,
	Arg:     option.EnableG2config,
	Type:    optiontype.Bool,
}

var OptionEnableG2configmgr = ContextVariable{
	Default: OsLookupEnvBool(envar.EnableG2configmgr, false),
	Envar:   envar.EnableG2configmgr,
	Help:    help.EnableG2configmgr,
	Arg:     option.EnableG2configmgr,
	Type:    optiontype.Bool,
}

var OptionEnableG2diagnostic = ContextVariable{
	Default: OsLookupEnvBool(envar.EnableG2diagnostic, false),
	Envar:   envar.EnableG2diagnostic,
	Help:    help.EnableG2diagnostic,
	Arg:     option.EnableG2diagnostic,
	Type:    optiontype.Bool,
}

var OptionEnableG2engine = ContextVariable{
	Default: OsLookupEnvBool(envar.EnableG2engine, false),
	Envar:   envar.EnableG2engine,
	Help:    help.EnableG2engine,
	Arg:     option.EnableG2engine,
	Type:    optiontype.Bool,
}

var OptionEnableG2product = ContextVariable{
	Default: OsLookupEnvBool(envar.EnableG2product, false),
	Envar:   envar.EnableG2product,
	Help:    help.EnableG2product,
	Arg:     option.EnableG2product,
	Type:    optiontype.Bool,
}

var OptionEnableSenzingChatApi = ContextVariable{
	Default: OsLookupEnvBool(envar.EnableSenzingChatApi, false),
	Envar:   envar.EnableSenzingChatApi,
	Help:    help.EnableSenzingChatApi,
	Arg:     option.EnableSenzingChatApi,
	Type:    optiontype.Bool,
}

var OptionEnableSenzingRestApi = ContextVariable{
	Default: OsLookupEnvBool(envar.EnableSenzingRestApi, false),
	Envar:   envar.EnableSenzingRestApi,
	Help:    help.EnableSenzingRestApi,
	Arg:     option.EnableSenzingRestApi,
	Type:    optiontype.Bool,
}

var OptionEnableSwaggerUi = ContextVariable{
	Default: OsLookupEnvBool(envar.EnableSwaggerUi, false),
	Envar:   envar.EnableSwaggerUi,
	Help:    help.EnableSwaggerUi,
	Arg:     option.EnableSwaggerUi,
	Type:    optiontype.Bool,
}

var OptionEnableXterm = ContextVariable{
	Default: OsLookupEnvBool(envar.EnableXterm, false),
	Envar:   envar.EnableXterm,
	Help:    help.EnableXterm,
	Arg:     option.EnableXterm,
	Type:    optiontype.Bool,
}

var OptionEngineConfigurationJson = ContextVariable{
	Default: OsLookupEnvString(envar.EngineConfigurationJson, ""),
	Envar:   envar.EngineConfigurationJson,
	Help:    help.EngineConfigurationJson,
	Arg:     option.EngineConfigurationJson,
	Type:    optiontype.String,
}

var OptionEngineLogLevel = ContextVariable{
	Default: OsLookupEnvInt(envar.EngineLogLevel, 0),
	Envar:   envar.EngineLogLevel,
	Help:    help.EngineLogLevel,
	Arg:     option.EngineLogLevel,
	Type:    optiontype.Int,
}

var OptionEngineModuleName = ContextVariable{
	Default: fmt.Sprintf("senzing-tools-%d", time.Now().Unix()),
	Envar:   envar.EngineModuleName,
	Help:    help.EngineModuleName,
	Arg:     option.EngineModuleName,
	Type:    optiontype.String,
}

var OptionGrpcPort = ContextVariable{
	Default: OsLookupEnvInt(envar.GrpcPort, 8260),
	Envar:   envar.GrpcPort,
	Help:    help.GrpcPort,
	Arg:     option.GrpcPort,
	Type:    optiontype.Int,
}

var OptionGrpcUrl = ContextVariable{
	Default: OsLookupEnvString(envar.GrpcUrl, ""),
	Envar:   envar.GrpcUrl,
	Help:    help.GrpcUrl,
	Arg:     option.GrpcUrl,
	Type:    optiontype.String,
}

var OptionHttpPort = ContextVariable{
	Default: OsLookupEnvInt(envar.HttpPort, 8261),
	Envar:   envar.HttpPort,
	Help:    help.HttpPort,
	Arg:     option.HttpPort,
	Type:    optiontype.Int,
}

var OptionInputFileType = ContextVariable{
	Default: OsLookupEnvString(envar.InputFileType, ""),
	Envar:   envar.InputFileType,
	Help:    help.InputFileType,
	Arg:     option.InputFileType,
	Type:    optiontype.String,
}

var OptionInputURL = ContextVariable{
	Default: OsLookupEnvString(envar.InputURL, ""),
	Envar:   envar.InputURL,
	Help:    help.InputURL,
	Arg:     option.InputURL,
	Type:    optiontype.String,
}

var OptionLicenseStringBase64 = ContextVariable{
	Default: OsLookupEnvString(envar.LicenseStringBase64, ""),
	Envar:   envar.LicenseStringBase64,
	Help:    help.LicenseStringBase64,
	Arg:     option.LicenseStringBase64,
	Type:    optiontype.String,
}

var OptionLogLevel = ContextVariable{
	Default: OsLookupEnvString(envar.LogLevel, "INFO"),
	Envar:   envar.LogLevel,
	Help:    help.LogLevel,
	Arg:     option.LogLevel,
	Type:    optiontype.String,
}

var OptionMonitoringPeriodInSeconds = ContextVariable{
	Default: OsLookupEnvInt(envar.MonitoringPeriodInSeconds, 60),
	Envar:   envar.MonitoringPeriodInSeconds,
	Help:    help.MonitoringPeriodInSeconds,
	Arg:     option.MonitoringPeriodInSeconds,
	Type:    optiontype.Int,
}

var OptionNumberOfWorkers = ContextVariable{
	Default: OsLookupEnvInt(envar.NumberOfWorkers, 0),
	Envar:   envar.NumberOfWorkers,
	Help:    help.NumberOfWorkers,
	Arg:     option.NumberOfWorkers,
	Type:    optiontype.Int,
}

var OptionObserverOrigin = ContextVariable{
	Default: OsLookupEnvString(envar.ObserverOrigin, ""),
	Envar:   envar.ObserverOrigin,
	Help:    help.ObserverOrigin,
	Arg:     option.ObserverOrigin,
	Type:    optiontype.String,
}

var OptionObserverUrl = ContextVariable{
	Default: OsLookupEnvString(envar.ObserverUrl, ""),
	Envar:   envar.ObserverUrl,
	Help:    help.ObserverUrl,
	Arg:     option.ObserverUrl,
	Type:    optiontype.String,
}

var OptionOutputURL = ContextVariable{
	Default: OsLookupEnvString(envar.OutputURL, ""),
	Envar:   envar.OutputURL,
	Help:    help.OutputURL,
	Arg:     option.OutputURL,
	Type:    optiontype.String,
}

var OptionRecordMax = ContextVariable{
	Default: OsLookupEnvInt(envar.RecordMax, 0),
	Envar:   envar.RecordMax,
	Help:    help.RecordMax,
	Arg:     option.RecordMax,
	Type:    optiontype.Int,
}

var OptionRecordMin = ContextVariable{
	Default: OsLookupEnvInt(envar.RecordMin, 0),
	Envar:   envar.RecordMin,
	Help:    help.RecordMin,
	Arg:     option.RecordMin,
	Type:    optiontype.Int,
}

var OptionRecordMonitor = ContextVariable{
	Default: OsLookupEnvInt(envar.RecordMonitor, 100000),
	Envar:   envar.RecordMonitor,
	Help:    help.RecordMonitor,
	Arg:     option.RecordMonitor,
	Type:    optiontype.Int,
}

var OptionResourcePath = ContextVariable{
	Default: OsLookupEnvString(envar.ResourcePath, ""),
	Envar:   envar.ResourcePath,
	Help:    help.ResourcePath,
	Arg:     option.ResourcePath,
	Type:    optiontype.String,
}

var OptionSenzingDirectory = ContextVariable{
	Default: OsLookupEnvString(envar.SenzingDirectory, ""),
	Envar:   envar.SenzingDirectory,
	Help:    help.SenzingDirectory,
	Arg:     option.SenzingDirectory,
	Type:    optiontype.String,
}

var OptionServerAddress = ContextVariable{
	Default: OsLookupEnvString(envar.ServerAddress, "0.0.0.0"),
	Envar:   envar.ServerAddress,
	Help:    help.ServerAddress,
	Arg:     option.ServerAddress,
	Type:    optiontype.String,
}

var OptionSupportPath = ContextVariable{
	Default: OsLookupEnvString(envar.SupportPath, "8261"),
	Envar:   envar.SupportPath,
	Help:    help.SupportPath,
	Arg:     option.SupportPath,
	Type:    optiontype.String,
}

var OptionVisibilityPeriodInSeconds = ContextVariable{
	Default: OsLookupEnvInt(envar.VisibilityPeriodInSeconds, 60),
	Envar:   envar.VisibilityPeriodInSeconds,
	Help:    help.VisibilityPeriodInSeconds,
	Arg:     option.VisibilityPeriodInSeconds,
	Type:    optiontype.Int,
}

var OptionXtermAllowedHostnames = ContextVariable{
	Default: []string{"localhost"},
	Envar:   envar.XtermAllowedHostnames,
	Help:    help.XtermAllowedHostnames,
	Arg:     option.XtermAllowedHostnames,
	Type:    optiontype.StringSlice,
}

var OptionXtermArguments = ContextVariable{
	Default: []string{},
	Envar:   envar.XtermArguments,
	Help:    help.XtermArguments,
	Arg:     option.XtermArguments,
	Type:    optiontype.StringSlice,
}

var OptionXtermCommand = ContextVariable{
	Default: OsLookupEnvString(envar.XtermCommand, "/bin/bash"),
	Envar:   envar.XtermCommand,
	Help:    help.XtermCommand,
	Arg:     option.XtermCommand,
	Type:    optiontype.String,
}

var OptionXtermConnectionErrorLimit = ContextVariable{
	Default: OsLookupEnvInt(envar.XtermConnectionErrorLimit, 10),
	Envar:   envar.XtermConnectionErrorLimit,
	Help:    help.XtermConnectionErrorLimit,
	Arg:     option.XtermConnectionErrorLimit,
	Type:    optiontype.Int,
}

var OptionXtermKeepalivePingTimeout = ContextVariable{
	Default: OsLookupEnvInt(envar.XtermKeepalivePingTimeout, 20),
	Envar:   envar.XtermKeepalivePingTimeout,
	Help:    help.XtermKeepalivePingTimeout,
	Arg:     option.XtermKeepalivePingTimeout,
	Type:    optiontype.Int,
}

var OptionXtermMaxBufferSizeBytes = ContextVariable{
	Default: OsLookupEnvInt(envar.XtermMaxBufferSizeBytes, 512),
	Envar:   envar.XtermMaxBufferSizeBytes,
	Help:    help.XtermMaxBufferSizeBytes,
	Arg:     option.XtermMaxBufferSizeBytes,
	Type:    optiontype.Int,
}
