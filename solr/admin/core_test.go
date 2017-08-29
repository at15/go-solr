package admin

import (
	"context"
	"testing"

	"github.com/at15/go-solr/pkg/common"
	asst "github.com/stretchr/testify/assert"
)

func TestService_CoreStatus(t *testing.T) {
	assert := asst.New(t)

	// TODO: need to create core before the status test
	_, err := tSvc.CoreStatus(context.Background(), false, "")
	assert.Nil(err)
}

func TestService_DeleteCore(t *testing.T) {
	assert := asst.New(t)

	coreName := "demotobedeleted"
	err := tSvc.CreateCore(context.Background(), common.NewCore(coreName))
	assert.Nil(err)
	err = tSvc.DeleteCore(context.Background(), coreName)
	assert.Nil(err)
}
