package solr

import (
	"encoding/json"
	"testing"

	"github.com/dyweb/gommon/util"
	asst "github.com/stretchr/testify/assert"
)

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
