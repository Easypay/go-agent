// Copyright 2020 New Relic Corporation. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"fmt"
	"os"
	"time"

	redis "github.com/go-redis/redis/v8"
	nrredis "github.com/Easypay/go-agent/v3/integrations/nrredis-v8"
	newrelic "github.com/Easypay/go-agent/v3/newrelic"
)

func main() {
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("Redis App"),
		newrelic.ConfigLicense(os.Getenv("NEW_RELIC_LICENSE_KEY")),
		newrelic.ConfigDebugLogger(os.Stdout),
	)
	if nil != err {
		panic(err)
	}
	app.WaitForConnection(10 * time.Second)
	txn := app.StartTransaction("ping txn")

	opts := &redis.Options{
		Addr: "localhost:6379",
	}
	client := redis.NewClient(opts)

	//
	// Step 1:  Add a nrredis.NewHook() to your redis client.
	//
	client.AddHook(nrredis.NewHook(opts))

	//
	// Step 2: Ensure that all client calls contain a context which includes
	// the transaction.
	//
	ctx := newrelic.NewContext(context.Background(), txn)
	pipe := client.WithContext(ctx).Pipeline()
	incr := pipe.Incr(ctx, "pipeline_counter")
	pipe.Expire(ctx, "pipeline_counter", time.Hour)
	_, err = pipe.Exec(ctx)
	fmt.Println(incr.Val(), err)

	txn.End()
	app.Shutdown(5 * time.Second)
}
