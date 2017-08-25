package core

import (
	"context"
	"testing"

	"github.com/dyweb/gommon/util"
	"github.com/at15/go-solr/pkg/search"
	asst "github.com/stretchr/testify/assert"
	"encoding/json"
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
