package ethbq

import (
	"context"

	"cloud.google.com/go/bigquery"
)

type Client struct {
	ProjectID string
	bq        *bigquery.Client
	ctx       context.Context
}

func NewClient(ctx context.Context, projectID string) (*Client, error) {
	var err error
	c := new(Client)
	c.ProjectID = projectID
	c.ctx = ctx
	c.bq, err = bigquery.NewClient(ctx, c.ProjectID)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c Client) Query(query string) (it *bigquery.RowIterator, err error) {
	return c.bq.Query(query).Read(c.ctx)
}
