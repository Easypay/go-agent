// Copyright 2020 New Relic Corporation. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package newrelic

import (
	"io"

	"github.com/Easypay/go-agent/internal/logger"
)

// Logger is the interface that is used for logging in the go-agent.  Assign the
// Config.Logger field to the Logger you wish to use.  Loggers must be safe for
// use in multiple goroutines.  Two Logger implementations are included:
// NewLogger, which logs at info level, and NewDebugLogger which logs at debug
// level.  logrus and logxi are supported by the integration packages
// https://godoc.org/github.com/Easypay/go-agent/_integrations/nrlogrus and
// https://godoc.org/github.com/Easypay/go-agent/_integrations/nrlogxi/v1.
type Logger interface {
	Error(msg string, context map[string]interface{})
	Warn(msg string, context map[string]interface{})
	Info(msg string, context map[string]interface{})
	Debug(msg string, context map[string]interface{})
	DebugEnabled() bool
}

// NewLogger creates a basic Logger at info level.
func NewLogger(w io.Writer) Logger {
	return logger.New(w, false)
}

// NewDebugLogger creates a basic Logger at debug level.
func NewDebugLogger(w io.Writer) Logger {
	return logger.New(w, true)
}
