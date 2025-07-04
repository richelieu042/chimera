package excelKit

import "github.com/xuri/excelize/v2"

// SetCellValue 设置单元格的值.
/*
@param sheetName 	工作表名
@param col 			[1, 16384]
@param row 			[1, 1048576]
*/
func SetCellValue(f *excelize.File, sheetName string, col, row int, value interface{}) error {
	cellName, err := CoordinatesToCellName(col, row)
	if err != nil {
		return err
	}
	return f.SetCellValue(sheetName, cellName, value)
}

func SetCellBool(f *excelize.File, sheetName string, col, row int, value bool) error {
	cellName, err := CoordinatesToCellName(col, row)
	if err != nil {
		return err
	}
	return f.SetCellBool(sheetName, cellName, value)
}

func SetCellInt(f *excelize.File, sheetName string, col, row int, value int64) error {
	cellName, err := CoordinatesToCellName(col, row)
	if err != nil {
		return err
	}
	return f.SetCellInt(sheetName, cellName, value)
}

func SetCellUint(f *excelize.File, sheetName string, col, row int, value uint64) error {
	cellName, err := CoordinatesToCellName(col, row)
	if err != nil {
		return err
	}
	return f.SetCellUint(sheetName, cellName, value)
}

func SetCellFloat(f *excelize.File, sheetName string, col, row int, value float64, precision, bitSize int) error {
	cellName, err := CoordinatesToCellName(col, row)
	if err != nil {
		return err
	}
	return f.SetCellFloat(sheetName, cellName, value, precision, bitSize)
}

func SetCellStr(f *excelize.File, sheetName string, col, row int, value string) error {
	cellName, err := CoordinatesToCellName(col, row)
	if err != nil {
		return err
	}
	return f.SetCellStr(sheetName, cellName, value)
}
