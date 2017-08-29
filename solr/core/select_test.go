package core

import (
	"context"
	"fmt"
	"testing"

	"github.com/at15/go-solr/pkg/common"
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

func TestService_SelectFacet(t *testing.T) {
	t.Skip("only works locally")

	assert := asst.New(t)

	jobSvc := New(tSvc.client, common.NewCore("jobs"), tSvc.admin)
	sq := search.StdQuery{}
	sq.And("*", "*")
	sq.IncludeField("raw_document")
	sq.DefaultField("namespace")
	sq.FacetField("status")
	sq.FacetField("submittedby")
	sq.FacetField("namespace")
	res, err := jobSvc.Select(context.Background(), &sq)
	fmt.Printf("%+v", err)
	assert.Nil(err)
	//t.Log(res.Response.Docs)
	t.Log(len(res.FacetCounts.FacetFields))
	t.Log(len(res.FacetCounts.FacetFields["status"].Values))
	t.Log(len(res.FacetCounts.FacetFields["submittedby"].Values))
}
