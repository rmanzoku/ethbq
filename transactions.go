package ethbq

import (
	"math/big"
	"time"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/iterator"
)

const (
	TransactionsTable = "`bigquery-public-data.crypto_ethereum.transactions`"
)

type Transaction struct {
	Hash                     string        `bigquery:"hash"`
	Nonce                    int64         `bigquery:"nonce"`
	TransactionIndex         int64         `bigquery:"transaction_index"`
	FromAddress              string        `bigquery:"from_address"`
	ToAddress                string        `bigquery:"to_address"`
	Value                    *big.Rat      `bigquery:"value"`
	Gas                      int64         `bigquery:"gas"`
	GasPrice                 int64         `bigquery:"gas_price"`
	Input                    string        `bigquery:"input"`
	ReceiptCumulativeGasUsed int64         `bigquery:"receipt_cumulative_gas_used"`
	ReceiptGasUsed           int64         `bigquery:"receipt_gas_used"`
	ReceiptContractAddress   string        `bigquery:"receipt_contract_address"`
	ReceiptRoot              string        `bigquery:"receipt_root"`
	ReceiptStatus            ReceiptStatus `bigquery:"receipt_status"`
	BlockTimestamp           time.Time     `bigquery:"block_timestamp"`
	BlockNumber              int64         `bigquery:"block_number"`
	BlockHash                string        `bigquery:"block_hash"`
}

type ReceiptStatus int8

const (
	Failure ReceiptStatus = iota
	Success
)

func (t Transaction) Success() bool {
	return t.ReceiptStatus == Success
}

func UnmarshalTransactions(it *bigquery.RowIterator, dst *[]*Transaction) (err error) {
	tmp := make([]*Transaction, it.TotalRows)
	i := 0
	for {
		values := new(Transaction)
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
