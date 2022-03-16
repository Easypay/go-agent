module github.com/Easypay/go-agent/v3/integrations/nrmysql

// 1.10 is the Go version in mysql's go.mod
go 1.10

require (
	// v3.3.9 includes the easypay changes
	github.com/Easypay/go-agent/v3 v3.3.9
	// v1.5.0 is the first mysql version to support gomod
	github.com/go-sql-driver/mysql v1.5.0
)
