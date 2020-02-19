package ethbq

import (
	"context"
	"fmt"
	"testing"

	"github.com/cheekybits/is"
)

var (
	c         = &Client{}
	projectID = "mch-prod"
)

func initTest(t *testing.T) is.I {
	var err error
	is := is.New(t)
	ctx := context.TODO()
	c, err = NewClient(ctx, projectID)
	is.Nil(err)
	return is
}

func TestQuery(t *testing.T) {
	is := initTest(t)
	query := `select ` + "`hash`" + `,from_address,to_address,value,receipt_status from ` + TransactionsTable + ` where block_timestamp > "2020-02-19" limit 10`
	it, err := c.Query(query)
	is.Nil(err)

	tx := []*Transaction{}
	err = UnmarhalTransactions(it, &tx)
	is.Nil(err)

	fmt.Println(*tx[0])
}
