package xlsx

import (
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"mime"
	"os"
	"path/filepath"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/xuri/excelize/v2"
)

type Writer interface {
	Sheet(name string) Writer
	Index(ind int) Writer
	WriteRow(row []any) error
	Save() error
}

type writer struct {
	xlsx  *excelize.File
	path  string
	sheet string
	index int
}

func (w *writer) Sheet(name string) Writer {
	w.sheet = name
	w.index = 0
	return w
}

func (w *writer) Index(ind int) Writer {
	w.index = ind
	return w
}

func (w *writer) WriteRow(row []any) error {
	w.index++
	cell, err := excelize.CoordinatesToCellName(1, w.index)
	if err != nil {
		return err
	}

	var files = make(map[string]*os.File)
	for ind, col := range row {
		// 判断是否为文件
		file, is := col.(*os.File)
		if is {
			row[ind] = nil

			es, err := excelize.CoordinatesToCellName(ind+1, w.index)
			if err != nil {
				return err
			}
			files[es] = file
		}
	}
	if err := w.xlsx.SetSheetRow(w.sheet, cell, &row); err != nil {
		return err
	}
	for cell, file := range files {
		fileMime := mime.TypeByExtension(filepath.Ext(file.Name()))
		if strings.Contains(fileMime, "image") {
			b, err := io.ReadAll(file)
			if err != nil {
				return err
			}
			row, col, _ := w.getCellIndex(cell)
			if err := w.xlsx.SetRowHeight(w.sheet, row+1, 100); err != nil {
				return err
			}
			if err := w.xlsx.SetColWidth(w.sheet, w.getCellName(col+1), w.getCellName(col+1), 16); err != nil {
				return err
			}
			if err := w.xlsx.AddPictureFromBytes(w.sheet, cell, &excelize.Picture{
				Extension: filepath.Ext(file.Name()),
				File:      b,
				Format: &excelize.GraphicOptions{
					AltText:         file.Name(),
					PrintObject:     proto.Bool(true),
					LockAspectRatio: true,
					AutoFit:         true,
					Locked:          proto.Bool(false),
				},
			}); err != nil {
				return err
			}
		}
	}
	return nil
}

func (w *writer) Save() error {
	return w.xlsx.SaveAs(w.path)
}

func (w *writer) getCellIndex(cell string) (int, int, error) {
	colName, rowName, err := excelize.SplitCellName(cell)
	if err != nil {
		return 0, 0, err
	}
	colNum, err := excelize.ColumnNameToNumber(colName)
	if err != nil {
		return 0, 0, err
	}
	return rowName - 1, colNum - 1, nil
}

func (w *writer) getCellName(num int) string {
	name, _ := excelize.ColumnNumberToName(num)
	return name
}
