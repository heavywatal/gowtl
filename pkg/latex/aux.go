package latex

import (
	"bytes"
	"regexp"
)

func CollectCitekeys(content []byte) map[string]struct{} {
	pattern := regexp.MustCompile(`\\(?:citation|abx@aux@cite\{\d\})\{(.+?)\}`)
	set := make(map[string]struct{}) // Go lacks "set"
	for _, subm := range pattern.FindAllSubmatch(content, -1) {
		for _, b := range bytes.Split(subm[1], []byte(",")) {
			set[string(b)] = struct{}{}
		}
	}
	return set
}

func CollectLabels(content []byte) map[string]string {
	pattern := regexp.MustCompile(`\\newlabel{(.+?)}{{(.+?)}`)
	labels := make(map[string]string)
	for _, x := range pattern.FindAllSubmatch(content, -1) {
		labels[string(x[1])] = string(x[2])
	}
	return labels
}
