package latex

import (
	"bytes"
	"regexp"
	"strings"

	"github.com/heavywatal/gowtl/pkg/debug"
	"github.com/heavywatal/gowtl/pkg/pathlib"
)

func BibFilter(bibfile string, auxfile string) []byte {
	if !strings.HasSuffix(bibfile, ".bib") {
		panic(bibfile + " is not .bib file.")
	}
	if !strings.HasSuffix(auxfile, ".aux") {
		panic(auxfile + " is not .aux file.")
	}
	citekeys := CollectCitekeys(pathlib.ReadFile(auxfile))
	buf := bytes.Buffer{}
	content := pathlib.ReadFile(bibfile)
	pattern := regexp.MustCompile(`(?ms)@\w+?{(\S+?),.+?}$`)
	for _, subm := range pattern.FindAllSubmatch(content, -1) {
		kbib := string(subm[1])
		_, ok := citekeys[kbib]
		if ok {
			delete(citekeys, kbib)
			buf.Write(selectField(subm[0]))
			buf.Write([]byte("\n\n"))
		}
	}
	if len(citekeys) > 0 {
		debug.Println(citekeys)
		panic("some citekeys not found.")
	}
	return buf.Bytes()
}

func selectField(entry []byte) []byte {
	buf := bytes.Buffer{}
	lines := bytes.Split(entry, []byte("\n"))
	buf.Write(lines[0])
	for _, line := range lines[1:] {
		l := bytes.TrimLeft(line, " \t")
		for _, prefix := range fieldsPrefix {
			if bytes.HasPrefix(l, prefix) {
				buf.Write([]byte("\n"))
				buf.Write(line)
			}
		}
	}
	b := buf.Bytes()
	b[len(b)-1] = '}' // Replace trailing comma
	return b
}

var (
	fields = []string{
		"author", "title", "journal", "year", "volume", "number", "pages",
		"publisher", "address", "editor"}
	fieldsPrefix = appendEq(fields)
)

func appendEq(vs []string) [][]byte {
	vb := make([][]byte, 0, len(vs))
	for _, s := range vs {
		vb = append(vb, []byte(s+" = "))
	}
	return vb
}
