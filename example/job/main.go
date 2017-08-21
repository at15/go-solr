package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/at15/go-solr/pkg"
	"github.com/at15/go-solr/pkg/common"
)

const coreName = "demojob"

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
	}
	log.Println("Solr is up")

	if err := solr.Admin.CreateCoreIfNotExists(context.Background(), common.NewCore(coreName)); err != nil {
		log.Fatalf("Create core %s failed %v", coreName, err)
		return
	}
	log.Printf("Created core %s (or it already exists)", coreName)

	if err := solr.UseCore(coreName); err != nil {
		log.Fatalf("can not use %s as default core %v", coreName, err)
		return
	}
	if status, err := solr.DefaultCore.Status(context.Background(), false); err != nil {
		log.Fatalf("Check core status failed %v", err)
		return
	} else {
		log.Printf("Got status for core %s %v\n", coreName, status)
	}
	if err := solr.Admin.DeleteCore(context.Background(), coreName); err != nil {
		log.Fatalf("Delete core %s failed %v", coreName, err)
		return
	}
	log.Println("core deleted, example finished")
}
