package gen

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/limes-cloud/kratosx/cmd/kratosx/internal/webutil/autocode/pkg/gen"
	"github.com/limes-cloud/kratosx/cmd/kratosx/internal/webutil/autocode/pkg/gen/types"
)

func TestSQL(t *testing.T) {
	initTable := func() *types.Table {
		var table types.Table
		content, err := os.ReadFile("./auto.json")
		if err != nil {
			panic(err)
		}

		if err := json.Unmarshal(content, &table); err != nil {
			panic(err)
		}
		return &table
	}

	sql := NewSQLBuilder(gen.NewBuilder(nil, initTable()))
	tableSql := sql.GenTableSQL()
	fmt.Println(tableSql)
}
