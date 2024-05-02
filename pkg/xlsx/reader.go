package xlsx

import (
	"github.com/xuri/excelize/v2"
)

type Iterator struct {
}

type Reader interface {
	GetSheets() []string
	GetRows(sheet string) ([][]string, error)
}

type reader struct {
	xlsx *excelize.File
}

func (r reader) GetSheets() []string {
	return r.xlsx.GetSheetList()
}

func (r reader) GetRows(sheet string) ([][]string, error) {
	var (
		res [][]string
		err error
	)
	res, err = r.xlsx.GetRows(sheet)
	if err != nil {
		return nil, err
	}

	// 获取图片
	cells, err := r.xlsx.GetPictureCells(sheet)
	if err != nil {
		return nil, err
	}

	for _, cell := range cells {
		row, col, _ := r.getCellIndex(cell)
		if len(res) < row+1 {
			continue
		}
		if len(res[row]) <= col {
			res[row] = append(res[row], "")
			col = len(res[row]) - 1
		}

		pics, err := r.xlsx.GetPictures(sheet, cell)
		if err != nil {
			return nil, err
		}
		if len(pics) == 0 {
			continue
		}
		res[row][col] = string(pics[0].File)
	}
	return res, nil
}

func (r reader) getCellIndex(cell string) (int, int, error) {
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
