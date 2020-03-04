package ethbq

import (
	"context"
	"testing"
	"time"

	"github.com/cheekybits/is"
)

var (
	c         = &Client{}
	projectID = "mch-prod"
	today     = time.Now().Format("2006-01-02")
)

func initTest(t *testing.T) is.I {
	var err error
	is := is.New(t)
	ctx := context.TODO()
	c, err = NewClient(ctx, projectID)
	is.Nil(err)
	return is
}
