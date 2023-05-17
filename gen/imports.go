package gen

import (
	"github.com/wenj91/mctl/go-zero/tools/goctl/util"
	"github.com/wenj91/mctl/template"
)

func genImports(withCache, timeImport bool) (string, error) {
	if withCache {
		text, err := util.LoadTemplate(category, importsTemplateFile, template.Imports)
		if err != nil {
			return "", err
		}

		buffer, err := util.With("import").Parse(text).Execute(map[string]interface{}{
			"time": timeImport,
		})
		if err != nil {
			return "", err
		}

		return buffer.String(), nil
	} else {
		text, err := util.LoadTemplate(category, importsWithNoCacheTemplateFile, template.ImportsNoCache)
		if err != nil {
			return "", err
		}

		buffer, err := util.With("import").Parse(text).Execute(map[string]interface{}{
			"time": timeImport,
		})
		if err != nil {
			return "", err
		}

		return buffer.String(), nil
	}
}
