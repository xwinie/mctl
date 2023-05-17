package gen

import (
	"strings"

	"github.com/wenj91/mctl/go-zero/tools/goctl/util"
	"github.com/wenj91/mctl/go-zero/tools/goctl/util/stringx"
	"github.com/wenj91/mctl/template"
)

func genFindOne(table Table, withCache bool) (string, string, string, error) {
	camel := table.Name.ToCamel()
	text, err := util.LoadTemplate(category, findOneTemplateFile, template.FindOne)
	if err != nil {
		return "", "", "", err
	}

	output, err := util.With("findOne").
		Parse(text).
		Execute(map[string]interface{}{
			"withCache":                 withCache,
			"upperStartCamelObject":     camel,
			"lowerStartCamelObject":     stringx.From(camel).Untitle(),
			"originalPrimaryKey":        wrapWithRawString(table.PrimaryKey.Name.Source()),
			"upperStartCamelPrimaryKey": table.PrimaryKey.Name.ToCamel(),
			"lowerStartCamelPrimaryKey": stringx.From(table.PrimaryKey.Name.ToCamel()).Untitle(),
			"dataType":                  strings.ReplaceAll(table.PrimaryKey.DataType, "*", ""),
			"cacheKey":                  table.CacheKey[table.PrimaryKey.Name.Source()].KeyExpression,
			"cacheKeyVariable":          table.CacheKey[table.PrimaryKey.Name.Source()].Variable,
		})
	if err != nil {
		return "", "", "", err
	}

	text, err = util.LoadTemplate(category, findOneMethodTemplateFile, template.FindOneMethod)
	if err != nil {
		return "", "", "", err
	}

	findOneMethod, err := util.With("findOneMethod").
		Parse(text).
		Execute(map[string]interface{}{
			"upperStartCamelObject":     camel,
			"lowerStartCamelPrimaryKey": stringx.From(table.PrimaryKey.Name.ToCamel()).Untitle(),
			"dataType":                  strings.ReplaceAll(table.PrimaryKey.DataType, "*", ""),
		})
	if err != nil {
		return "", "", "", err
	}

	// mapper
	text, err = util.LoadTemplate(category, findOneMapperTemplateFile, template.FindOneMapper)
	if err != nil {
		return "", "", "", err
	}

	findOneMapperOutput, err := util.With("findOneMapper").
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
		findOneMethod.String(),
		strings.Trim(findOneMapperOutput.String(), "\n"),
		nil
}
