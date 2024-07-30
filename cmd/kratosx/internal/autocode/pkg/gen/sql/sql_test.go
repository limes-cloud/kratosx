package gen

import (
	"encoding/json"
	"fmt"
	"github.com/limes-cloud/kratosx/cmd/kratosx/internal/autocode/pkg/gen/types"
	"os"
	"testing"
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

	sql := NewSQLBuilder(&builder{
		db:    nil,
		table: initTable(),
	})
	tableSql := sql.GenTableSQL()
	fmt.Println(tableSql)
}
