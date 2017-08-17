package main

import (
	"fmt"
	"github.com/at15/go-solr/pkg"
	"log"
	"context"
	"github.com/at15/go-solr/pkg/core"
	"os"
)

func main() {
	fmt.Println("example for storing job log in solr")
	fmt.Println("creating solr client")
	c := pkg.Config{}
	if addr := os.Getenv("GO_SOLR_ADDR"); addr != "" {
		log.Printf("solr addr %s set via env", addr)
		c.Addr = addr
	}
	solr, err := pkg.New(c)
	if err != nil {
		log.Fatal(err)
		return
	}
	if err := solr.IsUp(context.Background()); err != nil {
		log.Fatalf("Solr is not up %v", err)
		return
	} else {
		log.Println("Solr is up")
	}
	if err := solr.DefaultCore.Create(context.Background(), core.NewCore("demo")); err != nil {
		log.Fatalf("Create core demo failed %v", err)
		return
	} else {
		log.Println("Created core demo ")
	}
}
