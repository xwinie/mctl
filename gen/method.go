package gen

import (
	"github.com/wenj91/mctl/go-zero/tools/goctl/util"
	"github.com/wenj91/mctl/template"
)

func genMethod(table Table, withCache bool) (string, error) {
	text, err := util.LoadTemplate(category, modelMethodTemplateFile, template.Method)
	if err != nil {
		return "", err
	}

	output, err := util.With("method").
		Parse(text).
		Execute(map[string]interface{}{
			"table":                 wrapWithRawString(table.Name.Source()),
			"withCache":             withCache,
			"upperStartCamelObject": table.Name.ToCamel(),
		})
	if err != nil {
		return "", err
	}

	return output.String(), nil
}
