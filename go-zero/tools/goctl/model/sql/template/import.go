package template

var (
	Imports = `import (
	"database/sql"
	"fmt"
	"strings"
	{{if .time}}"time"{{end}}

	"github.com/wenj91/mctl/go-zero/core/stores/cache"
	"github.com/wenj91/mctl/go-zero/core/stores/sqlc"
	"github.com/wenj91/mctl/go-zero/core/stores/sqlx"
	"github.com/wenj91/mctl/go-zero/core/stringx"
	"github.com/wenj91/mctl/go-zero/tools/goctl/model/sql/builderx"
)
`
	ImportsNoCache = `import (
	"database/sql"
	"fmt"
	"strings"
	{{if .time}}"time"{{end}}

	"github.com/wenj91/mctl/go-zero/core/stores/sqlc"
	"github.com/wenj91/mctl/go-zero/core/stores/sqlx"
	"github.com/wenj91/mctl/go-zero/core/stringx"
	"github.com/wenj91/mctl/go-zero/tools/goctl/model/sql/builderx"
)
`
)
