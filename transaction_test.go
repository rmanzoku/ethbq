package ethbq

import (
	"fmt"
	"testing"
)

func TestUnmarshalTransactions(t *testing.T) {
	is := initTest(t)
	query := `select ` + "`hash`" + `,from_address,to_address,value,receipt_status from ` + TransactionsTable + ` where block_timestamp > "2020-02-19" limit 10`
	it, err := c.Query(query)
	is.Nil(err)

	tx := []*Transaction{}
	err = UnmarshalTransactions(it, &tx)
	is.Nil(err)

	fmt.Println(*tx[0])
}
