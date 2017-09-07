package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/at15/go-solr/solr"
)

const coreName = "job"

func main() {
	fmt.Println("example for storing job log in solr")
	fmt.Println("creating solr client")
	c := solr.Config{}
	if addr := os.Getenv("GO_SOLR_ADDR"); addr != "" {
		log.Printf("solr addr %s set via env", addr)
		c.Addr = addr
	}
	solrClient, err := solr.NewClient(c)
	if err != nil {
		log.Fatal(err)
		return
	}
	if err := solrClient.IsUp(context.Background()); err != nil {
		log.Fatalf("Solr is not up %v", err)
		return
	}
	log.Println("Solr is up")

	core := solr.NewCore(coreName)
	core.ConfigSet = coreName
	exists, err := solrClient.CreateCoreIfNotExists(context.Background(), core);
	if err != nil {
		log.Fatalf("Create core %s failed %v", coreName, err)
		return
	}
	if exists {
		log.Printf("Core %s already exists\n", coreName)
	} else {
		log.Printf("Created core %s", coreName)
	}
	solrClient.UseCore(coreName)
	if status, err := solrClient.DefaultCore.Status(context.Background(), false); err != nil {
		log.Fatalf("Check core status failed %v", err)
		return
	} else {
		log.Printf("Got status for core %s %v\n", coreName, status)
	}
}
