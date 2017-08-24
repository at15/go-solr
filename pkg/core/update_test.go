package core

import (
	"context"
	"testing"

	asst "github.com/stretchr/testify/assert"
)

type foo struct {
	Bar string `json:"bar"`
}

func TestService_Update(t *testing.T) {
	assert := asst.New(t)

	b := []foo{
		{"a"}, {"b"}, {"c"},
	}
	assert.Nil(tSvc.Update(context.Background(), b))
}
