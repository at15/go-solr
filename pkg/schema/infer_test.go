package schema

import (
	"reflect"
	"testing"

	"github.com/at15/go-solr/pkg/common/fieldtype"
	"github.com/at15/go-solr/pkg/schema/fixture"
	asst "github.com/stretchr/testify/assert"
)

func TestInferSchema(t *testing.T) {
	// TODO: refactor to table driven test
	assert := asst.New(t)

	job := fixture.Job{}
	// support struct and pointer to struct
	_, err := InferSchema(job)
	assert.Nil(err)
	_, err = InferSchema(&job)
	assert.Nil(err)
	_, err = InferSchema("haha")
	assert.NotNil(err)

	// error when there is no exported field
	private := fixture.AllPrivate{}
	_, err = InferSchema(private)
	assert.NotNil(err)

	// []byte is treated as text_general by default
	sma, err := InferSchema(fixture.ByteSlice{})
	assert.Nil(err)
	assert.Equal(fieldtype.TextGeneral, sma.Fields[0].Type)

}

func TestStd_Time(t *testing.T) {
	tp := reflect.TypeOf(fixture.ByteSlice{})
	f := tp.Field(0)
	t.Log(typeOfTime.Kind()) // struct
	t.Log(typeOfByteSlice)   // []uint8
	t.Log(f.Type)            // uint8
}
