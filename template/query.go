package template

var Query = `// Code generated by [mctl-zzinfo](https://github.com/wenj91/mctl/tree/zzinfo), DO NOT EDIT.
package {{.pkg}};

{{.imports}}
public class {{.upperStartCamelObject}}Query extends AbstractQuery {

    private {{.upperStartCamelObject}} Query() {}

{{.table}}
{{.fields}}

{{.queryMethod}}

{{.fieldMethods}}
    
}
`
