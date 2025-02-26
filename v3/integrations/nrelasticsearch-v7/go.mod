module github.com/Easypay/go-agent/v3/integrations/nrelasticsearch-v7

// As of Jan 2020, the v7 elasticsearch go.mod uses 1.11:
// https://github.com/elastic/go-elasticsearch/blob/7.x/go.mod
go 1.11

require (
	github.com/elastic/go-elasticsearch/v7 v7.5.0
	github.com/Easypay/go-agent/v3 v3.0.0
)
