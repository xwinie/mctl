package util

import (
	"strings"

	"github.com/wenj91/mctl/go-zero/tools/goctl/api/spec"
)

func GetAnnotationValue(annos []spec.Annotation, key, field string) (string, bool) {
	for _, anno := range annos {
		if anno.Name == field && len(anno.Value) > 0 {
			return anno.Value, true
		}
		if anno.Name == key {
			value, ok := anno.Properties[field]
			return strings.TrimSpace(value), ok
		}
	}
	return "", false
}
