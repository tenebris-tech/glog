//
// Copyright (c) 2020-2021 Tenebris Technologies Inc.
//

package glog

import (
	"encoding/json"
	"fmt"
	"log"
)

// All log messages with a format string pass through here
func gLogf(level string, format string, data ...interface{}) {
	tmp := fmt.Sprintf(format, data...)
	gLog(level, tmp)
}

// All unstructured messages pass through here
func gLog(level string, msg string) {
	e := New()
	e["level"] = level
	e["message"] = msg
	event(e)
}

// All log messages pass through here
func event(e LogEvent) {

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
		if val, ok := e["message"]; ok {
			log.Println(val)
		} else {
			log.Printf("logger error, message element is missing")
			log.Printf("raw message: %v", e)
		}
	}
}
