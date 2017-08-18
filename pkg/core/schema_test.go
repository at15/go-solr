package core

import (
	"context"
	"testing"

	asst "github.com/stretchr/testify/assert"
)

func TestService_Schema(t *testing.T) {
	assert := asst.New(t)

	s, err := tSvc.Schema(context.Background())
	assert.Nil(err)
	assert.Equal("example-data-driven-schema", s.Name)
	assert.Equal("id", s.UniqueKey)
}
