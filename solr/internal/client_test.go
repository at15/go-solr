package internal

import (
	"net/http"
	"testing"

	asst "github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	assert := asst.New(t)
	var (
		c   *Client
		err error
	)
	c, err = NewClient(nil)
	// FIXME: we should be creating new client with new transport and default timeout
	assert.Equal(c.http, http.DefaultClient)
	assert.Nil(err)
}
