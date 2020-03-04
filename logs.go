package ethbq

import (
	"time"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/iterator"
)

const (
	LogsTable = "`bigquery-public-data.crypto_ethereum.logs`"
)

type Log struct {
	LogIndex         int64     `bigquery:"log_index"`
	TransactionHash  string    `bigquery:"transaction_hash"`
	TransactionIndex int64     `bigquery:"transaction_index"`
	Address          string    `bigquery:"address"`
	Data             string    `bigquery:"data"`
	Topics           []string  `bigquery:"topics"`
	BlockTimestamp   time.Time `bigquery:"block_timestamp"`
	BlockNumber      int64     `bigquery:"block_number"`
	BlockHash        string    `bigquery:"block_hash"`
}

func UnmarshalLogs(it *bigquery.RowIterator, dst *[]*Log) (err error) {
	tmp := make([]*Log, it.TotalRows)
	i := 0
	for {
		values := new(Log)
		err := it.Next(values)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		tmp[i] = values
		i++
	}
	*dst = tmp
	return nil
}
