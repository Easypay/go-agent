// Copyright 2020 New Relic Corporation. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package newrelic

// This file contains the names of the automatically captured attributes.
// Attributes are key value pairs attached to transaction events, error events,
// and traced errors.  You may add your own attributes using the
// Transaction.AddAttribute method (see transaction.go).
//
// These attribute names are exposed here to facilitate configuration.
//
// For more information, see:
// https://docs.newrelic.com/docs/agents/manage-apm-agents/agent-metrics/agent-attributes

// Attributes destined for Transaction Events, Errors, and Transaction Traces:
const (
	// AttributeResponseCode is the response status code for a web request.
	AttributeResponseCode = "httpResponseCode"
	// AttributeRequestMethod is the request's method.
	AttributeRequestMethod = "request.method"
	// AttributeRequestAccept is the request's "Accept" header.
	AttributeRequestAccept = "request.headers.accept"
	// AttributeRequestContentType is the request's "Content-Type" header.
	AttributeRequestContentType = "request.headers.contentType"
	// AttributeRequestContentLength is the request's "Content-Length" header.
	AttributeRequestContentLength = "request.headers.contentLength"
	// AttributeRequestHost is the request's "Host" header.
	AttributeRequestHost = "request.headers.host"
	// AttributeRequestURI is the request's URL without query parameters,
	// fragment, user, or password.
	AttributeRequestURI = "request.uri"
	// AttributeResponseContentType is the response "Content-Type" header.
	AttributeResponseContentType = "response.headers.contentType"
	// AttributeResponseContentLength is the response "Content-Length" header.
	AttributeResponseContentLength = "response.headers.contentLength"
	// AttributeHostDisplayName contains the value of Config.HostDisplayName.
	AttributeHostDisplayName = "host.displayName"
)

// Attributes destined for Errors and Transaction Traces:
const (
	// AttributeRequestUserAgent is the request's "User-Agent" header.
	AttributeRequestUserAgent = "request.headers.User-Agent"
	// AttributeRequestReferer is the request's "Referer" header.  Query
	// string parameters are removed.
	AttributeRequestReferer = "request.headers.referer"
)

// AWS Lambda specific attributes:
const (
	AttributeAWSRequestID            = "aws.requestId"
	AttributeAWSLambdaARN            = "aws.lambda.arn"
	AttributeAWSLambdaColdStart      = "aws.lambda.coldStart"
	AttributeAWSLambdaEventSourceARN = "aws.lambda.eventSource.arn"
)

// Attributes for consumed message transactions:
//
// When a message is consumed (for example from Kafka or RabbitMQ), supported
// instrumentation packages -- i.e. those found in the _integrations
// (https://godoc.org/github.com/Easypay/go-agent/_integrations) directory --
// will add these attributes automatically.  `AttributeMessageExchangeType`,
// `AttributeMessageReplyTo`, and `AttributeMessageCorrelationID` are disabled
// by default.  To see these attributes added to all destinations, you must add
// include them in your config settings:
//
//	cfg.Attributes.Include = append(cfg.Attributes.Include,
//		AttributeMessageExchangeType, AttributeMessageReplyTo,
//		AttributeMessageCorrelationID)
//
// When not using a supported instrumentation package, you can add these
// attributes manually using the `Transaction.AddAttribute`
// (https://godoc.org/github.com/Easypay/go-agent#Transaction) API.  In this
// case, these attributes will be included on all destintations by default.
//
//	txn := app.StartTransaction("Message/RabbitMQ/Exchange/Named/MyExchange", nil, nil)
//	txn.AddAttribute(AttributeMessageRoutingKey, "myRoutingKey")
//	txn.AddAttribute(AttributeMessageQueueName, "myQueueName")
//	txn.AddAttribute(AttributeMessageExchangeType, "myExchangeType")
//	txn.AddAttribute(AttributeMessageReplyTo, "myReplyTo")
//	txn.AddAttribute(AttributeMessageCorrelationID, "myCorrelationID")
//	// ... consume a message ...
//	txn.End()
//
// It is recommended that at most one message is consumed per transaction.
const (
	// The routing key of the consumed message.
	AttributeMessageRoutingKey = "message.routingKey"
	// The name of the queue the message was consumed from.
	AttributeMessageQueueName = "message.queueName"
	// The type of exchange used for the consumed message (direct, fanout,
	// topic, or headers).
	AttributeMessageExchangeType = "message.exchangeType"
	// The callback queue used in RPC configurations.
	AttributeMessageReplyTo = "message.replyTo"
	// The application-generated identifier used in RPC configurations.
	AttributeMessageCorrelationID = "message.correlationId"
)

// Attributes destined for Span Events:
//
// To disable the capture of one of these span event attributes, db.statement
// for example, modify your Config like this:
//
//	cfg.SpanEvents.Attributes.Exclude = append(cfg.SpanEvents.Attributes.Exclude,
//		newrelic.SpanAttributeDBStatement)
const (
	SpanAttributeDBStatement  = "db.statement"
	SpanAttributeDBInstance   = "db.instance"
	SpanAttributeDBCollection = "db.collection"
	SpanAttributePeerAddress  = "peer.address"
	SpanAttributePeerHostname = "peer.hostname"
	SpanAttributeHTTPURL      = "http.url"
	SpanAttributeHTTPMethod   = "http.method"
	SpanAttributeAWSOperation = "aws.operation"
	SpanAttributeAWSRequestID = "aws.requestId"
	SpanAttributeAWSRegion    = "aws.region"
)
