package main

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"flag"
	"fmt"
	"golang.org/x/tools/go/packages"
	"log"
)

func main() {
	var (
		schemaPath = flag.String("path", "", "path to schema directory")
	)
	flag.Parse()
	if *schemaPath == "" {
		log.Fatal("entproto: must specify schema path. use entproto -path ./ent/schema")
	}
	graph, err := entc.LoadGraph(*schemaPath, &gen.Config{})

	pkgs, err := packages.Load(&packages.Config{
		Mode: packages.NeedName | packages.NeedTypes | packages.NeedTypesInfo | packages.NeedModule,
	}, *schemaPath)
	fmt.Println(pkgs)
	if err != nil {
		log.Fatalf("entproto: failed loading ent graph: %v", err)
	}
	if err := entproto.Generate(graph); err != nil {
		log.Fatalf("entproto: failed generating protos: %s", err)
	}
}
