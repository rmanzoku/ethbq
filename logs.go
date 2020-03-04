package ethbq

import (
	"strings"
	"time"

	"cloud.google.com/go/bigquery"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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

func (l Log) GoEthereumLog() *types.Log {
	ret := new(types.Log)
	ret.Address = common.HexToAddress(l.Address)
	ret.Topics = make([]common.Hash, len(l.Topics))
	for i, t := range l.Topics {
		ret.Topics[i] = common.HexToHash(t)
	}
	ret.Data = common.FromHex(l.Data)
	ret.BlockNumber = uint64(l.BlockNumber)
	ret.TxHash = common.HexToHash(l.TransactionHash)
	ret.TxIndex = uint(l.TransactionIndex)

	ret.Index = uint(l.LogIndex)
	ret.Removed = false
	return ret
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

func NewBoundContractForUnpack(abiStr string) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(abiStr))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(common.HexToAddress("0x0000000000000000000000000000000000000000"), parsed, nil, nil, nil), nil
}
