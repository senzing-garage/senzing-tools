package cmdhelper

import (
	"os"
	"strconv"
	"strings"
)

func OsLookupEnvBool(envar string, aDefault bool) bool {
	resultString, isSet := os.LookupEnv(envar)
	if !isSet {
		return aDefault
	}
	result, err := strconv.ParseBool(strings.ToUpper(resultString))
	if err != nil {
		panic(err)
	}
	return result
}

func OsLookupEnvInt(envar string, aDefault int) int {
	resultString, isSet := os.LookupEnv(envar)
	if !isSet {
		return aDefault
	}
	result, err := strconv.Atoi(resultString)
	if err != nil {
		panic(err)
	}
	return result
}

func OsLookupEnvString(envar string, aDefault string) string {
	result, isSet := os.LookupEnv(envar)
	if !isSet {
		return aDefault
	}
	return result
}
