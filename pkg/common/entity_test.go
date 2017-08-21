package common

import (
	"testing"

	"encoding/json"
	asst "github.com/stretchr/testify/assert"
	"github.com/at15/go-solr/pkg/common/fieldtype"
)

func TestNewField(t *testing.T) {
	assert := asst.New(t)
	// NOTE: we use *bool instead of bool, so when docValues is not specified, it is not in the payload,
	// if we use omitempty, we can't pass false, if we don't use omitempty, we may pass false for flags that are default
	// as true in the server side
	f := NewField("name", fieldtype.String)
	b, err := json.Marshal(f)
	assert.Nil(err)
	assert.NotContains(string(b), "docValues")
	f.DocValues = &False
	b, err = json.Marshal(f)
	assert.Contains(string(b), "docValues")
}
