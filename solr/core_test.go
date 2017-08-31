package solr

import (
	"context"
	"testing"

	"bytes"
	"github.com/dyweb/gommon/util"
	asst "github.com/stretchr/testify/assert"
)

func TestCoreClient_Ping(t *testing.T) {
	assert := asst.New(t)

	_, err := tDemoCoreClient.Ping(context.Background())
	assert.Nil(err)
}

func TestCoreClient_Status(t *testing.T) {
	assert := asst.New(t)

	status, err := tDemoCoreClient.Status(context.Background(), false)
	assert.Nil(err)
	assert.Equal("demo", status.Name)
}

func TestCoreClient_Select(t *testing.T) {
	t.Run("return everything using *:*", func(t *testing.T) {
		assert := asst.New(t)
		q := StdQuery{}
		q.And("*", "*")
		res, err := tFilmClient.Select(context.Background(), &q)
		assert.Nil(err)
		t.Log(res.Response.NumFound)
		t.Log(len(res.Response.Docs))
	})
	// TODO: test facet
	//jobSvc := New(tSvc.client, common.NewCore("jobs"), tSvc.admin)
	//sq := search.StdQuery{}
	//sq.And("*", "*")
	//sq.IncludeField("raw_document")
	//sq.DefaultField("namespace")
	//sq.FacetField("status")
	//sq.FacetField("submittedby")
	//sq.FacetField("namespace")
	//res, err := jobSvc.Select(context.Background(), &sq)
	//fmt.Printf("%+v", err)
	//assert.Nil(err)
	////t.Log(res.Response.Docs)
	//t.Log(len(res.FacetCounts.FacetFields))
	//t.Log(len(res.FacetCounts.FacetFields["status"].Values))
	//t.Log(len(res.FacetCounts.FacetFields["submittedby"].Values))
}

func TestCoreClient_Update(t *testing.T) {
	t.Run("update JSON file using example/film/films.json", func(t *testing.T) {
		assert := asst.New(t)

		b := util.ReadFixture(t, "../example/film/films.json")
		err := tFilmClient.Update(context.Background(), bytes.NewReader(b))
		assert.Nil(err)
	})
}
