package main

import (
	"fmt"
	"github.com/at15/go-solr/pkg"
	"log"
	"context"
)

func main() {
	fmt.Println("example for storing job log in solr")
	fmt.Println("creating solr client")
	c := pkg.Config{}
	solr, err := pkg.New(c)
	if err != nil {
		log.Fatal(err)
		return
	}
	if err := solr.IsUp(context.Background()); err != nil {
		log.Fatalf("Solr is not up %v", err)
		return
	}
}
