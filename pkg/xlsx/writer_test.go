package xlsx

import (
	"os"
	"testing"

	"github.com/xuri/excelize/v2"
	"github.com/zeebo/assert"
)

func TestWriter_WriteRow(t *testing.T) {
	w := writer{path: "xlsx/write.xlsx", xlsx: excelize.NewFile(), sheet: "Sheet1"}
	fb, err := os.Open("xlsx/file.png")
	assert.NoError(t, err)

	fb1, err := os.Open("xlsx/image.jpg")
	assert.NoError(t, err)
	tests := [][]any{
		{"1", "@", "#", fb},
		{"@", "1", "4", fb1},
	}

	for _, test := range tests {
		err := w.WriteRow(test)
		assert.NoError(t, err)
	}
	w.Save()
}
