package core

import (
	"context"
	"testing"

	"github.com/at15/go-solr/pkg/search"
	asst "github.com/stretchr/testify/assert"
)

func TestService_Select(t *testing.T) {
	assert := asst.New(t)

	q := search.StdQuery{}
	q.And("*", "*")
	res, err := tSvc.Select(context.Background(), &q)
	assert.Nil(err)
	t.Log(res.Response.Docs)
	t.Log(res)
}
