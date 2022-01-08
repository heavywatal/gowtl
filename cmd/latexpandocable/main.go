package main

import (
	"flag"
	"os"

	"github.com/heavywatal/gowtl/pkg/latex"
)

func main() {
	outfile := flag.String("o", "", "outfile")
	flag.Parse()
	infile := flag.Arg(0)
	content := latex.Pandocable(infile)
	if *outfile != "" {
		os.WriteFile(*outfile, content, 0644)
	} else {
		os.Stdout.Write(content)
	}
}
