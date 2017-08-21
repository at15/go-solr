package schema

import (
	"context"
	"testing"

	"github.com/at15/go-solr/pkg/common"
	"github.com/at15/go-solr/pkg/common/fieldtype"
	asst "github.com/stretchr/testify/assert"
)

func TestService_AddField(t *testing.T) {
	assert := asst.New(t)

	name := "addfieldtest"
	err := tSvc.AddField(context.Background(), common.NewField(name, fieldtype.String))
	assert.Nil(err)
	err = tSvc.DeleteField(context.Background(), name)
	assert.Nil(err)
}
