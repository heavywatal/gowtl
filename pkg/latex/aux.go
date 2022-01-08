package latex

import (
	"regexp"
)

func CollectLabels(content []byte) map[string]string {
	pattern := regexp.MustCompile(`\\newlabel{([^}]+)}{{([^}]+)`)
	labels := make(map[string]string)
	for _, x := range pattern.FindAllSubmatch(content, -1) {
		labels[string(x[1])] = string(x[2])
	}
	return labels
}
