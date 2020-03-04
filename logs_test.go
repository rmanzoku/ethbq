package ethbq

import (
	"fmt"
	"testing"
)

func TestUnmarshalLogs(t *testing.T) {
	is := initTest(t)
	query := `select log_index,transaction_hash,transaction_index,address,data,topics,block_timestamp,block_number,block_hash from ` + LogsTable + ` where block_timestamp > "` + today + `" limit 10`
	it, err := c.Query(query)
	is.Nil(err)

	ret := []*Log{}
	err = UnmarshalLogs(it, &ret)
	is.Nil(err)

	fmt.Println(*ret[0])
}
