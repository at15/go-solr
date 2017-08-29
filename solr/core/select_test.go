package core

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/at15/go-solr/pkg/common"
	"github.com/at15/go-solr/pkg/search"
	"github.com/dyweb/gommon/util"
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

func TestFacetField_UnmarshalJSON(t *testing.T) {
	assert := asst.New(t)

	b := util.ReadFixture(t, "fixture/facet.json")
	res := SelectResponse{}
	assert.Nil(json.Unmarshal(b, &res))
	assert.Equal("a", res.FacetCounts.FacetFields["foo"].Values[0])
	assert.Equal(123, res.FacetCounts.FacetFields["foo"].Counts[0])
	assert.Equal("b", res.FacetCounts.FacetFields["foo"].Values[1])
	assert.Equal(321, res.FacetCounts.FacetFields["foo"].Counts[1])
	b, err := json.Marshal(res.FacetCounts.FacetFields)
	t.Log(string(b))
	assert.Nil(err)
}
