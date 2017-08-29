package solr

import (
	"context"
	"testing"

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
