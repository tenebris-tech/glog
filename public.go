//
// Copyright (c) 2020-2021 Tenebris Technologies Inc.
//
package glog

//goland:noinspection GoUnusedExportedFunction
func SetLogFile(fileName string) {
	logFileName = fileName
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
