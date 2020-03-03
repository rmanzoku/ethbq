package ethbq

import (
	"context"
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
