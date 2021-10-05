//
// Copyright (c) 2020-2021 Tenebris Technologies Inc.
//

// Package glog implements a basic generic logger
package glog

import (
	"strings"
)

type LogEvent map[string]interface{}

func New() LogEvent {
	return LogEvent{}
}

//goland:noinspection GoUnusedExportedFunction
func ConsoleOutput(b bool) {
	consoleOutput = b
}

//goland:noinspection GoUnusedExportedFunction
func SetFile(s string) {
	logFileName = s
}

//goland:noinspection GoUnusedExportedFunction
func SetFormat(f string) {
	switch strings.ToLower(f) {
	case "json":
		outputFormat = OutputJSON
	case "nv":
		outputFormat = OutputNVPairs
	case "pairs":
		outputFormat = OutputNVPairs
	case "message":
		outputFormat = OutputMessage
	case "text":
		outputFormat = OutputMessage
	default:
		outputFormat = OutputJSON
	}
}

//goland:noinspection GoUnusedExportedFunction
func Debug(e string) {
	gLog("debug", e)
}

//goland:noinspection GoUnusedExportedFunction
func Info(e string) {
	gLog("info", e)
}

//goland:noinspection GoUnusedExportedFunction
func Notice(e string) {
	gLog("notice", e)
}

//goland:noinspection GoUnusedExportedFunction
func Warn(e string) {
	gLog("warn", e)
}

//goland:noinspection GoUnusedExportedFunction
func Error(e string) {
	gLog("error", e)
}

//goland:noinspection GoUnusedExportedFunction
func Debugf(format string, e ...interface{}) {
	gLogf("debug", format, e...)
}

//goland:noinspection GoUnusedExportedFunction
func Infof(format string, e ...interface{}) {
	gLogf("info", format, e...)
}

//goland:noinspection GoUnusedExportedFunction
func Noticef(format string, e ...interface{}) {
	gLogf("notice", format, e...)
}

//goland:noinspection GoUnusedExportedFunction
func Warnf(format string, e ...interface{}) {
	gLogf("warn", format, e...)
}

//goland:noinspection GoUnusedExportedFunction
func Errorf(format string, e ...interface{}) {
	gLogf("error", format, e...)
}

//goland:noinspection GoUnusedExportedFunction
func Event(e LogEvent) {

	// If log level not supplied or empty, default to info
	val, ok := e["level"]
	if ok == false || val == "" {
		e["level"] = "info"
	}

	// There should always be a log message, but just in case...
	val, ok = e["message"]
	if ok == false || val == "" {
		e["message"] = "none"
	}

	// Send event for processing
	event(e)
}
