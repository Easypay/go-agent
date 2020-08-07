module github.com/newrelic/go-agent/v4/integrations/nrnats

// As of Dec 2019, 1.11 is the earliest version of Go tested by nats:
// https://github.com/nats-io/nats.go/blob/master/.travis.yml
go 1.11

require (
	// v1.8.0 is the first nats version with a go.mod.
	github.com/nats-io/nats.go v1.8.0
	github.com/newrelic/go-agent/v4 v4.0.0
)
