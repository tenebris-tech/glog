//
// Copyright (c) 2020-2021 Tenebris Technologies Inc.
//

package glog

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

// All log messages with a format string pass through this function
func gLogf(level string, format string, data ...interface{}) {
	tmp := fmt.Sprintf(format, data...)
	gLog(level, tmp)
}

// All unstructured messages pass through this function
func gLog(level string, msg string) {
	e := New()
	e["level"] = level
	e["message"] = msg
	event(e)
}

// All log messages pass through this function
func event(e LogEvent) {

	// Clean up the message
	if val, ok := e["message"]; ok {
		msg := fmt.Sprintf("%v", val)
		msg = strings.Trim(msg, "\n ")
		msg = strings.TrimSpace(msg)
		e["message"] = msg
	}

	switch outputFormat {

	case OutputJSON:
		txt, err := json.Marshal(e)
		if err != nil {
			log.Printf("error creating json: %s", err.Error())
			log.Printf("raw message: %v", e)
		}
		log.Println(string(txt))

	case OutputNVPairs:
		tmp := ""
		// Iterate over e
		for n, v := range e {
			if tmp != "" {
				tmp = tmp + ", "
			}
			tmp = tmp + fmt.Sprintf("%s=\"%v\"", n, v)
		}
		log.Println(tmp)

	case OutputMessage:
		var level = ""

		if val, ok := e["level"]; ok {
			level = fmt.Sprintf("%v", val)
		} else {
			level = "none"
		}

		if val, ok := e["message"]; ok {
			log.Printf("%s %v", level, val)
		} else {
			log.Printf("logger error, message element is missing")
			log.Printf("raw message: %v", e)
		}
	}
}
