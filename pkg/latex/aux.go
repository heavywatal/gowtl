package latex

import (
	"bytes"
	"regexp"
)

func CollectCitekeys(content []byte) []string {
	pattern := regexp.MustCompile(`\\citation{(.+?)}`)
	set := make(map[string]struct{}) // Go lacks "set"
	for _, subm := range pattern.FindAllSubmatch(content, -1) {
		for _, b := range bytes.Split(subm[1], []byte(",")) {
			set[string(b)] = struct{}{}
		}
	}
	return keys(set)
}

func CollectLabels(content []byte) map[string]string {
	pattern := regexp.MustCompile(`\\newlabel{(.+?)}{{(.+?)}`)
	labels := make(map[string]string)
	for _, x := range pattern.FindAllSubmatch(content, -1) {
		labels[string(x[1])] = string(x[2])
	}
	return labels
}

func keys(m map[string]struct{}) []string {
	v := make([]string, 0, len(m))
	for k := range m {
		v = append(v, k)
	}
	return v
}
