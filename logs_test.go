package ethbq

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var zelGatewayV1SoldABI = `[{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"user","type":"address"},{"indexed":true,"internalType":"address","name":"referrer","type":"address"},{"indexed":false,"internalType":"uint8","name":"purchaseType","type":"uint8"},{"indexed":false,"internalType":"uint256","name":"grossValue","type":"uint256"},{"indexed":false,"internalType":"uint256","name":"referralValue","type":"uint256"},{"indexed":false,"internalType":"uint256","name":"netValue","type":"uint256"}],"name":"Sold","type":"event"}]`

type zelGatewayV1Sold struct {
	User          common.Address
	Referrer      common.Address
	PurchaseType  uint8
	GrossValue    *big.Int
	ReferralValue *big.Int
	NetValue      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

func TestUnmarshalLogs(t *testing.T) {
	is := initTest(t)
	query := `select log_index,transaction_hash,transaction_index,address,data,topics,block_timestamp,block_number,block_hash from ` + LogsTable + ` where block_timestamp > "` + today + `" AND address = "0x97390a8fa25b2e69a34d50d289e7c0eb87b1c5c8" limit 10`
	it, err := c.Query(query)
	is.Nil(err)

	ret := []*Log{}
	err = UnmarshalLogs(it, &ret)
	is.Nil(err)

	e := new(zelGatewayV1Sold)

	contract, err := NewBoundContractForUnpack(zelGatewayV1SoldABI)
	is.Nil(err)
	err = contract.UnpackLog(e, "Sold", *ret[0].GoEthereumLog())
	is.Nil(err)
	fmt.Println(e.User.String(), e.Referrer.String(), e.PurchaseType, e.GrossValue.String(), e.NetValue.String())
}
