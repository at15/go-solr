package schema

import (
	"context"
	"testing"

	"github.com/at15/go-solr/pkg/common"
	"github.com/at15/go-solr/pkg/common/fieldtype"
	asst "github.com/stretchr/testify/assert"
	"github.com/at15/go-solr/pkg/fixture"
)

func TestService_AddField(t *testing.T) {
	assert := asst.New(t)

	name := "addfieldtest"
	err := tSvc.AddField(context.Background(), common.NewField(name, fieldtype.String))
	assert.Nil(err)
	err = tSvc.DeleteField(context.Background(), name)
	assert.Nil(err)
}

func TestService_AddFields(t *testing.T) {
	assert := asst.New(t)
	// FIXME: this will create the schema from genesis in demo core while it should not ....
	err := tSvc.AddFields(context.Background(), fixture.JobFieldsSchema...)
	assert.Nil(err)
}
