package gen

import (
	"strings"

	"github.com/wenj91/mctl/go-zero/core/collection"
	"github.com/wenj91/mctl/go-zero/tools/goctl/util"
	"github.com/wenj91/mctl/go-zero/tools/goctl/util/stringx"
	"github.com/wenj91/mctl/template"
)

func genDelete(table Table, withCache bool) (string, string, string, error) {
	keySet := collection.NewSet()
	keyVariableSet := collection.NewSet()
	for fieldName, key := range table.CacheKey {
		if fieldName == table.PrimaryKey.Name.Source() {
			keySet.AddStr(key.KeyExpression)
		} else {
			keySet.AddStr(key.DataKeyExpression)
		}
		keyVariableSet.AddStr(key.Variable)
	}

	camel := table.Name.ToCamel()
	text, err := util.LoadTemplate(category, deleteTemplateFile, template.Delete)
	if err != nil {
		return "", "", "", err
	}

	output, err := util.With("delete").
		Parse(text).
		Execute(map[string]interface{}{
			"upperStartCamelObject":     camel,
			"withCache":                 withCache,
			"containsIndexCache":        table.ContainsUniqueKey,
			"upperStartCamelPrimaryKey": table.PrimaryKey.Name.ToCamel(),
			"lowerStartCamelPrimaryKey": stringx.From(table.PrimaryKey.Name.ToCamel()).Untitle(),
			"dataType":                  strings.ReplaceAll(table.PrimaryKey.DataType, "*", ""),
			"keys":                      strings.Join(keySet.KeysStr(), "\n"),
			"originalPrimaryKey":        wrapWithRawString(table.PrimaryKey.Name.Source()),
			"keyValues":                 strings.Join(keyVariableSet.KeysStr(), ", "),
		})
	if err != nil {
		return "", "", "", err
	}

	// interface method
	text, err = util.LoadTemplate(category, deleteMethodTemplateFile, template.DeleteMethod)
	if err != nil {
		return "", "", "", err
	}

	deleteMethodOut, err := util.With("deleteMethod").
		Parse(text).
		Execute(map[string]interface{}{
			"lowerStartCamelPrimaryKey": stringx.From(table.PrimaryKey.Name.ToCamel()).Untitle(),
			"dataType":                  strings.ReplaceAll(table.PrimaryKey.DataType, "*", ""),
		})
	if err != nil {
		return "", "", "", err
	}

	// mapper
	text, err = util.LoadTemplate(category, deleteMapperTemplateFile, template.DeleteMapper)
	if err != nil {
		return "", "", "", err
	}

	deleteMapperOutput, err := util.With("deleteMapper").
		Parse(text).
		Execute(map[string]interface{}{
			"table": table.Name.Source(),
			"field": table.PrimaryKey.Name.Source(),
			"value": table.PrimaryKey.Name.ToCamel(),
		})
	if err != nil {
		return "", "", "", err
	}

	return output.String(),
		deleteMethodOut.String(),
		strings.Trim(deleteMapperOutput.String(), "\n"),
		nil
}
