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
func SetFile(fileName string) {
	logFileName = fileName
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
func Debug(event string) {
	gLog("debug", event)
}

//goland:noinspection GoUnusedExportedFunction
func Info(event string) {
	gLog("info", event)
}

//goland:noinspection GoUnusedExportedFunction
func Notice(event string) {
	gLog("notice", event)
}

//goland:noinspection GoUnusedExportedFunction
func Warn(event string) {
	gLog("warn", event)
}

//goland:noinspection GoUnusedExportedFunction
func Error(event string) {
	gLog("error", event)
}

//goland:noinspection GoUnusedExportedFunction
func Debugf(format string, event ...interface{}) {
	gLogf("debug", format, event...)
}

//goland:noinspection GoUnusedExportedFunction
func Infof(format string, event ...interface{}) {
	gLogf("info", format, event...)
}

//goland:noinspection GoUnusedExportedFunction
func Noticef(format string, event ...interface{}) {
	gLogf("notice", format, event...)
}

//goland:noinspection GoUnusedExportedFunction
func Warnf(format string, event ...interface{}) {
	gLogf("warn", format, event...)
}

//goland:noinspection GoUnusedExportedFunction
func Errorf(format string, event ...interface{}) {
	gLogf("error", format, event...)
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
