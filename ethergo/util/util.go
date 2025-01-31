package util

import "strings"

// FormatError applies custom formatting & noise reduction to error messages. Add more as needed.
func FormatError(err error) string {
	if err == nil {
		return ""
	}
	errMsg := err.Error()

	//if an error message contains embedded HTML (eg: many RPC errors), strip it out to reduce noise.
	if strings.Contains(errMsg, "<!DOCTYPE html>") {
		errMsg = strings.Split(errMsg, "<!DOCTYPE html>")[0] + "<html portion of error removed>"
	}
	return errMsg
}
