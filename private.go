//
// Copyright (c) 2020-2021 Tenebris Technologies Inc.
//
package glog

import (
	"fmt"
	"log"
	"strings"
)

// All log messages with a format string pass through here
func gLogf(level string, format string, event ...interface{}) {
	tmp := fmt.Sprintf(format, event...)
	gLog(level, tmp)
}

// All log messages pass through here
func gLog(level string, event string) {
	log.Print("["+strings.ToUpper(level)+"]", " ", event)
}
