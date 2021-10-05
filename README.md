# glog - Golang Custom Logger

Copyright (c) 2020-2021 Tenebris Technologies Inc. All rights reserved.

This package implements a customer logger. It is a work in progress, but the interface is expected to remain stable.

All public functions are in public.go

Log events can be passed as single strings, in Printf() format, or as an LogEvent object. The latter is recommended,
especially for output in JSON format.

Several output formats are supported.

Please see the LICENSE file for licensing information.
