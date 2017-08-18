package admin

import (
	"context"
	"testing"

	asst "github.com/stretchr/testify/assert"
)

func TestService_CoreStatus(t *testing.T) {
	assert := asst.New(t)

	// TODO: need to create core before the status test
	_, err := tSvc.CoreStatus(context.Background(), false, "")
	assert.Nil(err)
}
