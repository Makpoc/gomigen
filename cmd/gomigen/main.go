package main

import (
	"flag"
	"log"
	"os"

	"github.com/Makpoc/gomigen/internal/app"
)

func main() {
	outputDir := flag.String(
		"out",
		"./", // default is current dir
		"optional path to output directory where the generated package with middleware files will be placed",
	)
	flag.Parse()

	if flag.NArg() < 2 {
		log.Fatalf("Usage: %s [-out output_directory] <directory> <interface>", os.Args[0])
	}

	pkg, iface := flag.Arg(0), flag.Arg(1)
	a := app.New(
		pkg, iface,
		app.WithOutputDirectory(*outputDir), app.WithLogger(log.Default()),
	)
	if err := a.Run(); err != nil {
		log.Fatalf("Generating middleware failed: %v", err)
	}
}
