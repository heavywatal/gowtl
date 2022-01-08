package main

import (
	"flag"
	"os"

	"github.com/heavywatal/gowtl/pkg/latex"
)

func main() {
	outfile := flag.String("o", "-", "outfile")
	flag.Parse()
	bibfile := flag.Arg(0)
	auxfile := flag.Arg(1)
	content := latex.BibFilter(bibfile, auxfile)
	if *outfile != "-" {
		os.WriteFile(*outfile, content, 0644)
	} else {
		os.Stdout.Write(content)
	}
}
