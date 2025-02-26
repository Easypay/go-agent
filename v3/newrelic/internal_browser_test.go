// Copyright 2020 New Relic Corporation. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package newrelic

import (
	"reflect"
	"testing"

	"github.com/Easypay/go-agent/v3/internal"
)

func browserReplyFields(reply *internal.ConnectReply) {
	reply.AgentLoader = "loader"
	reply.Beacon = "beacon"
	reply.BrowserKey = "key"
	reply.AppID = "app"
	reply.ErrorBeacon = "error"
	reply.JSAgentFile = "agent"
}

func TestBrowserTimingHeaderSuccess(t *testing.T) {
	includeAttributes := func(cfg *Config) {
		cfg.BrowserMonitoring.Attributes.Enabled = true
		cfg.BrowserMonitoring.Attributes.Include = []string{AttributeResponseCode}
	}
	app := testApp(browserReplyFields, includeAttributes, t)
	txn := app.StartTransaction("hello")
	rw := txn.SetWebResponse(nil)
	rw.WriteHeader(200)
	txn.AddAttribute("zip", "zap")
	hdr := txn.BrowserTimingHeader()
	app.expectNoLoggedErrors(t)

	encodingKey := browserEncodingKey(testLicenseKey)
	obfuscatedTxnName, _ := obfuscate([]byte("OtherTransaction/Go/hello"), encodingKey)
	obfuscatedAttributes, _ := obfuscate([]byte(`{"u":{"zip":"zap"},"a":{"http.statusCode":200}}`), encodingKey)

	// This is a cheat: we can't deterministically set this, but DeepEqual
	// doesn't have any ability to say "equal everything except these
	// fields".
	hdr.info.QueueTimeMillis = 12
	hdr.info.ApplicationTimeMillis = 34
	expected := &BrowserTimingHeader{
		agentLoader: "loader",
		info: browserInfo{
			Beacon:                "beacon",
			LicenseKey:            "key",
			ApplicationID:         "app",
			TransactionName:       obfuscatedTxnName,
			QueueTimeMillis:       12,
			ApplicationTimeMillis: 34,
			ObfuscatedAttributes:  obfuscatedAttributes,
			ErrorBeacon:           "error",
			Agent:                 "agent",
		},
	}
	if !reflect.DeepEqual(hdr, expected) {
		txnName, _ := deobfuscate(hdr.info.TransactionName, encodingKey)
		attr, _ := deobfuscate(hdr.info.ObfuscatedAttributes, encodingKey)
		t.Errorf("header did not match: expected %#v; got %#v txnName=%s attr=%s",
			expected, hdr, string(txnName), string(attr))
	}
}

func TestBrowserTimingHeaderSuccessWithoutAttributes(t *testing.T) {
	// Test that attributes do not get put in the browser footer by default
	// configuration.

	app := testApp(browserReplyFields, nil, t)
	txn := app.StartTransaction("hello")
	rw := txn.SetWebResponse(nil)
	rw.WriteHeader(200)
	txn.AddAttribute("zip", "zap")
	hdr := txn.BrowserTimingHeader()
	app.expectNoLoggedErrors(t)

	encodingKey := browserEncodingKey(testLicenseKey)
	obfuscatedTxnName, _ := obfuscate([]byte("OtherTransaction/Go/hello"), encodingKey)
	obfuscatedAttributes, _ := obfuscate([]byte(`{"u":{},"a":{}}`), encodingKey)

	// This is a cheat: we can't deterministically set this, but DeepEqual
	// doesn't have any ability to say "equal everything except these
	// fields".
	hdr.info.QueueTimeMillis = 12
	hdr.info.ApplicationTimeMillis = 34
	expected := &BrowserTimingHeader{
		agentLoader: "loader",
		info: browserInfo{
			Beacon:                "beacon",
			LicenseKey:            "key",
			ApplicationID:         "app",
			TransactionName:       obfuscatedTxnName,
			QueueTimeMillis:       12,
			ApplicationTimeMillis: 34,
			ObfuscatedAttributes:  obfuscatedAttributes,
			ErrorBeacon:           "error",
			Agent:                 "agent",
		},
	}
	if !reflect.DeepEqual(hdr, expected) {
		txnName, _ := deobfuscate(hdr.info.TransactionName, encodingKey)
		attr, _ := deobfuscate(hdr.info.ObfuscatedAttributes, encodingKey)
		t.Errorf("header did not match: expected %#v; got %#v txnName=%s attr=%s",
			expected, hdr, string(txnName), string(attr))
	}
}

func TestBrowserTimingHeaderDisabled(t *testing.T) {
	disableBrowser := func(cfg *Config) {
		cfg.BrowserMonitoring.Enabled = false
	}
	app := testApp(browserReplyFields, disableBrowser, t)
	txn := app.StartTransaction("hello")
	hdr := txn.BrowserTimingHeader()
	app.expectSingleLoggedError(t, "unable to create browser timing header", map[string]interface{}{
		"reason": errBrowserDisabled.Error(),
	})
	if hdr.WithTags() != nil {
		t.Error(hdr.WithTags())
	}
}

func TestBrowserTimingHeaderNotConnected(t *testing.T) {
	app := testApp(nil, nil, t)
	txn := app.StartTransaction("hello")
	hdr := txn.BrowserTimingHeader()
	// No error expected if the app is not yet connected.
	app.expectNoLoggedErrors(t)
	if hdr.WithTags() != nil {
		t.Error(hdr.WithTags())
	}
}

func TestBrowserTimingHeaderAlreadyFinished(t *testing.T) {
	app := testApp(browserReplyFields, nil, t)
	txn := app.StartTransaction("hello")
	txn.End()
	hdr := txn.BrowserTimingHeader()
	app.expectSingleLoggedError(t, "unable to create browser timing header", map[string]interface{}{
		"reason": errAlreadyEnded.Error(),
	})
	if hdr.WithTags() != nil {
		t.Error(hdr.WithTags())
	}
}

func TestBrowserTimingHeaderTxnIgnored(t *testing.T) {
	app := testApp(browserReplyFields, nil, t)
	txn := app.StartTransaction("hello")
	txn.Ignore()
	hdr := txn.BrowserTimingHeader()
	app.expectSingleLoggedError(t, "unable to create browser timing header", map[string]interface{}{
		"reason": errTransactionIgnored.Error(),
	})
	if hdr.WithTags() != nil {
		t.Error(hdr.WithTags())
	}
}

func BenchmarkBrowserTimingHeaderSuccess(b *testing.B) {
	app := testApp(browserReplyFields, nil, b)
	txn := app.StartTransaction("hello")

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		hdr := txn.BrowserTimingHeader()
		if nil == hdr {
			b.Fatal(hdr)
		}
		app.expectNoLoggedErrors(b)
		hdr.WithTags()
	}
}
