package latex

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/heavywatal/gowtl/pkg/debug"
	"github.com/heavywatal/gowtl/pkg/pathlib"
)

func Pandocable(texfile string) []byte {
	auxfile := pathlib.WithSuffix(texfile, ".aux")
	labelmap := CollectLabels(pathlib.ReadFile(auxfile))
	debug.Pprintln(labelmap)
	content := pathlib.ReadFile(texfile)
	content = resolveRef(content, labelmap)
	content = removeAsterisk(content)
	content = labelCaption(content, labelmap)
	return content
}

func resolveRef(content []byte, labelmap map[string]string) []byte {
	pattern := regexp.MustCompile(`\\ref{[^}]+}`)
	repl := func(matched []byte) []byte {
		label := matched[5:(len(matched) - 1)]
		return []byte(labelmap[string(label)])
	}
	return pattern.ReplaceAllFunc(content, repl)
}

func removeAsterisk(content []byte) []byte {
	pattern := regexp.MustCompile(`{(table|figure)\*}`)
	return pattern.ReplaceAll(content, []byte("{$1}"))
}

func labelCaption(content []byte, labelmap map[string]string) []byte {
	pattern := regexp.MustCompile(`(?s)caption{(.+?)\\label{([^}]+)}`)
	repl := func(matched []byte) []byte {
		subm := pattern.FindSubmatch(matched)
		label := string(subm[2])
		class := classifyLabel(label)
		num := labelmap[label]
		s := fmt.Sprintf("caption{\\textbf{%s %s}. %s", class, num, subm[1])
		debug.Println(debug.Head(s))
		return []byte(s)
	}
	return pattern.ReplaceAllFunc(content, repl)
}

func classifyLabel(label string) string {
	label = strings.ToLower(label)
	if strings.HasPrefix(label, "fig") {
		return "Figure"
	} else if strings.HasPrefix(label, "tab") {
		return "Table"
	}
	return "Equation???"
}
