package solr

import (
	"os"
	"testing"
)

var tClient *Client
var tDemoCoreClient *CoreClient
var tFilmClient *CoreClient

func TestMain(m *testing.M) {
	log.Info("Setup")
	addr := os.Getenv(AddrEnvName)
	if addr == "" {
		addr = DefaultAddr
	}
	config := Config{Addr: addr, DefaultCore: "demo", Cloud: false}
	var err error
	if tClient, err = NewClient(config); err != nil {
		panic(err)
		return
	}
	tDemoCoreClient = tClient.GetCore("demo")
	tFilmClient = tClient.GetCore("film")
	v := m.Run()
	log.Info("tear down")
	os.Exit(v)
}
