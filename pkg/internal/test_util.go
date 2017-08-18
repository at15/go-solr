package internal

import (
	"net/http"
	"os"
)

const (
	defaultAddr = "http://localhost:8983"
	addrEnvName = "GO_SOLR_ADDR"
)

func NewInternalClient() (*Client, error) {
	var addr string
	if addr = os.Getenv(addrEnvName); addr == "" {
		addr = defaultAddr
	}
	log.Infof("use solr %s", addr)
	tr := &http.Transport{}
	c := &http.Client{Transport: tr}
	return NewClient(c, BaseURL(addr))
}

func MustNewInternalClient() *Client {
	c, err := NewInternalClient()
	if err != nil {
		log.Error("can't create internal client for test")
		panic(err)
	}
	return c
}
