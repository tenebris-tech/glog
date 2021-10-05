//
// Copyright (c) 2020-2021 Tenebris Technologies Inc.
//

package glog

import "os"

// Optional log file
var logFileName = ""
var logFile *os.File = nil

// Specify and manage output format
var outputFormat = OutputJSON
var consoleOutput = false

const OutputJSON = 0
const OutputNVPairs = 1
const OutputMessage = 2
