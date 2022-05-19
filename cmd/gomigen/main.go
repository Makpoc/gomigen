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
		"", // default is current dir
		"optional path to output directory where the generated package with middleware files will be placed",
	)
	flag.Parse()

	// TODO would it be better to make all args named?
	//   e.g. `go run mw/.../gen -pkg . -iface Service [-out output_directory]`
	if flag.NArg() < 2 {
		log.Printf("Usage: %s [-out output_directory] <directory> <interface>", os.Args[0])
		os.Exit(1)
	}
	var outDir string
	if outputDir != nil {
		outDir = *outputDir
	}
	pkg, iface := flag.Arg(0), flag.Arg(1)
	a := app.New(
		pkg, iface,
		app.WithOutputDirectory(outDir), app.WithLogger(log.Default()),
	)
	if err := a.Run(); err != nil {
		log.Printf("generating middleware failed: %v", err)
		os.Exit(1)
	}
}
