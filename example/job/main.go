package main

import (
	"fmt"
	"github.com/at15/go-solr/pkg"
	"log"
	"context"
	"github.com/at15/go-solr/pkg/common"
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
	if err := solr.Admin.CreateCoreIfNotExists(context.Background(), common.NewCore("demo")); err != nil {
		log.Fatalf("Create core demo failed %v", err)
		return
	} else {
		log.Println("Created core demo (or it already exists)")
	}
	// FIXME: we are not having error because default core is demo and we created demo core in previous lines
	if status, err := solr.DefaultCore.Status(context.Background(), false); err != nil {
		log.Fatalf("Check core status failed %v", err)
		return
	} else {
		log.Printf("Got status for core %v\n", status)
	}
}
