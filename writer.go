//
// Copyright (c) 2020-2021 Tenebris Technologies Inc.
//

package glog

import (
	"fmt"
	"os"
	"time"
)

// Custom io.Writer
type logWriter struct {
}

// Write implements a customer log writer
func (writer logWriter) Write(bytes []byte) (int, error) {
	var err error

	// Create timestamp using YYYY-MM-DD HH:MM:SS
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	// If no log file or consoleOutput is true, just print the log
	if logFileName == "" || consoleOutput {
		n, err := fmt.Print(timestamp + " " + string(bytes))

		// If no log file, we're done
		if logFileName == "" {
			return n, err
		}
	}

	// If log file is not open, attempt to open it
	if logFile == nil {
		logFile, err = os.OpenFile(logFileName, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			// Failed
			fmt.Print(timestamp + " " + "[ERROR] Unable to open log file " + logFileName + " for writing")
			// Don't try again
			logFileName = ""
			logFile = nil
			// Write the original message
			return fmt.Print(timestamp + " " + string(bytes))
		}
	}

	// Write to the log file
	return logFile.WriteString(timestamp + " " + string(bytes))
}
