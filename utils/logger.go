package utils

import (
	"github.com/forsam-education/simplelogger"
	"os"
)

// Logger is the shared logger for the application.
var Logger simplelogger.Logger

// VerboseFlag describes if Kerberos is in verbose mode.
var VerboseFlag bool

func init() {
	Logger = simplelogger.NewDefaultLogger()
}

// LogAndForceExit is a quick helper to force exit on unrecoverable errors.
func LogAndForceExit(err error) {
	Logger.StdError(err, nil)
	os.Exit(1)
}

// LogVerbose is used to log only if verbose mode is enabled.
func LogVerbose(message string, data map[string]interface{}) {
	if VerboseFlag {
		Logger.Debug(message, data)
	}
}
