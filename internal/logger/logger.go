// Copyright 2021 The webserver Authors. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package logger

import (
	"log"

	"github.com/thalesfsp/sypl"
	"github.com/thalesfsp/sypl/level"
	"github.com/thalesfsp/sypl/output"
	"github.com/thalesfsp/sypl/processor"
	"github.com/thalesfsp/sypl/status"
)

// Global, singleton, cached logger. It's safe to be retrieved via `Get`.
var l *sypl.Sypl

// Get safely returns the global application logger.
func Get() *sypl.Sypl {
	if l != nil {
		return l
	}

	log.Fatalln("Logger isn't setup")

	return nil
}

// Setup logger.
func Setup(name, logLevel, requestLogLevel, logFilePath string) *sypl.Sypl {
	logLevelAsLevel := level.MustFromString(logLevel)
	requesLogLevelAsLevel := level.MustFromString(requestLogLevel)

	l = sypl.NewDefault(
		name,
		logLevelAsLevel,
		processor.ChangeFirstCharCase(processor.Lowercase),
	)

	l.SetDefaultIoWriterLevel(requesLogLevelAsLevel)

	// Should only enable File output if path is set.
	if logFilePath != "" {
		l.AddOutputs(output.File(
			logFilePath,
			logLevelAsLevel,
			processor.ChangeFirstCharCase(processor.Lowercase),
		))

		// "-" special case makes the `File` Output behave as `Console`.
		// To avoid duplication, it disables the `Console` output.
		if logFilePath == "-" {
			l.GetOutput("Console").SetStatus(status.Disabled)
		}
	}

	return l
}
