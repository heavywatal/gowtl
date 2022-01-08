package debug

import (
	"encoding/json"
	"fmt"
	"os"
)

func Pprint(x interface{}) {
	Print(string(Json(x)))
}

func Pprintln(x interface{}) {
	Println(string(Json(x)))
}

func Print(x ...interface{}) {
	fmt.Fprint(os.Stderr, x...)
}

func Println(x ...interface{}) {
	fmt.Fprintln(os.Stderr, x...)
}

func Printf(format string, x ...interface{}) {
	fmt.Fprintf(os.Stderr, format, x...)
}

func Json(x interface{}) []byte {
	b, _ := json.MarshalIndent(x, "", "  ")
	return b
}

func Head(s string, n_opt ...int) string {
	n := 60
	if len(n_opt) > 0 {
		n = n_opt[0]
	}
	if len(s) > n {
		s = s[:n] + "…"
	}
	return s
}
