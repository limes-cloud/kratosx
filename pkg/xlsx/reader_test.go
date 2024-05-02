package xlsx

import (
	"os"
	"testing"

	"github.com/xuri/excelize/v2"
	"github.com/zeebo/assert"
)

func TestReader_GetRows(t *testing.T) {
	f, err := excelize.OpenFile("xlsx/test.xlsx")
	assert.NoError(t, err)

	fb, err := os.ReadFile("xlsx/file.png")
	assert.NoError(t, err)

	tests := [][]string{
		{
			"1", "2", "3", string(fb),
		},
		{
			"1", "2", "3", string(fb),
		},
	}

	rd := reader{
		xlsx: f,
	}
	rows, err := rd.GetRows("Sheet1")
	assert.NoError(t, err)

	for ind, row := range rows {
		assert.DeepEqual(t, row, tests[ind])
	}
}
