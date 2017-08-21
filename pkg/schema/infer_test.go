package schema

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/at15/go-solr/pkg/common/fieldtype"
	"github.com/at15/go-solr/pkg/schema/fixture"
	asst "github.com/stretchr/testify/assert"
)

func TestInferSchema(t *testing.T) {
	// TODO: refactor to table driven test
	// TODO: use MustInferSchema in test

	t.Run("only supports struct and pointer to struct", func(t *testing.T) {
		assert := asst.New(t)
		job := fixture.Job{}
		sma, err := InferSchema(job)
		assert.Nil(err)
		sma, err = InferSchema(&job)
		assert.Nil(err)
		assert.NotEmpty(sma.Fields)
		// http://changelog.ca/log/2015/03/09/golang, fancier printing
		fmt.Printf("%#v\n", sma.Fields[0])

		sma, err = InferSchema("haha")
		assert.Nil(sma)
		assert.NotNil(err)
	})

	t.Run("uses json tag for name", func(t *testing.T) {
		assert := asst.New(t)
		sma, err := InferSchema(&fixture.JsonTag{})
		assert.Nil(err)
		assert.Equal("foo", sma.Fields[0].Name)
	})

	t.Run("returns error when no exported field found", func(t *testing.T) {
		assert := asst.New(t)
		// error when there is no exported field
		private := fixture.AllPrivate{}
		_, err := InferSchema(private)
		assert.NotNil(err)
	})

	t.Run("use text_general for []byte by default", func(t *testing.T) {
		assert := asst.New(t)
		sma, err := InferSchema(fixture.ByteSlice{})
		assert.Nil(err)
		assert.Equal(fieldtype.TextGeneral, sma.Fields[0].Type)
	})
}

func TestStd_Types(t *testing.T) {
	tp := reflect.TypeOf(fixture.ByteSlice{})
	f := tp.Field(0)
	t.Log(typeOfTime.Kind()) // struct
	t.Log(typeOfByteSlice)   // []uint8
	t.Log(f.Type)            // uint8
}
