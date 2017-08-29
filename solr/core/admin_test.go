package core

import (
	"context"
	"testing"

	asst "github.com/stretchr/testify/assert"
)

func TestService_Ping(t *testing.T) {
	assert := asst.New(t)

	err := tSvc.Ping(context.Background())
	assert.Nil(err)
}

func TestService_Status(t *testing.T) {
	assert := asst.New(t)

	status, err := tSvc.Status(context.Background(), false)
	assert.Nil(err)
	assert.Equal("demo", status.Name)
}
