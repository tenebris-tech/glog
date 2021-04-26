//
// Copyright (c) 2020-2021 Tenebris Technologies Inc.
//

package glog

import "log"

// Initialization
func init() {
	log.SetFlags(0)
	log.SetOutput(new(logWriter))
}
