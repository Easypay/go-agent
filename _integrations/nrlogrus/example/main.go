// Copyright 2020 New Relic Corporation. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	newrelic "github.com/Easypay/go-agent"
	"github.com/Easypay/go-agent/_integrations/nrlogrus"
	"github.com/sirupsen/logrus"
)

func mustGetEnv(key string) string {
	if val := os.Getenv(key); "" != val {
		return val
	}
	panic(fmt.Sprintf("environment variable %s unset", key))
}

func main() {
	cfg := newrelic.NewConfig("Logrus App", mustGetEnv("NEW_RELIC_LICENSE_KEY"))
	logrus.SetLevel(logrus.DebugLevel)
	cfg.Logger = nrlogrus.StandardLogger()

	app, err := newrelic.NewApplication(cfg)
	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}

	http.HandleFunc(newrelic.WrapHandleFunc(app, "/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello world")
	}))

	http.ListenAndServe(":8000", nil)
}
