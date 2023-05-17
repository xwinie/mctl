package template

var Error = `package {{.pkg}}

import "github.com/wenj91/mctl/go-zero/core/stores/sqlx"

var ErrNotFound = sqlx.ErrNotFound
`
