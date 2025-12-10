package excelx

import (
	"github.com/xuri/excelize/v2"
)

// Sheet 定义单个 Sheet 的数据
type Sheet struct {
	Name    string   // Sheet 名称
	Headers []string // 标题行
	Rows    [][]any  // 数据行
}

// writeSheet 写入单个 Sheet 数据
func writeSheet(f *excelize.File, sheet string, headers []string, rows [][]any, boldStyle int) error {
	// 写入标题行
	for col, header := range headers {
		cell, err := excelize.CoordinatesToCellName(col+1, 1)
		if err != nil {
			return err
		}
		if err := f.SetCellValue(sheet, cell, header); err != nil {
			return err
		}
		if err := f.SetCellStyle(sheet, cell, cell, boldStyle); err != nil {
			return err
		}
	}

	// 写入数据行
	for rowIdx, row := range rows {
		for colIdx, value := range row {
			cell, err := excelize.CoordinatesToCellName(colIdx+1, rowIdx+2)
			if err != nil {
				return err
			}
			if err := f.SetCellValue(sheet, cell, value); err != nil {
				return err
			}
		}
	}

	return nil
}

// export 导出单 Sheet Excel 文件（内部方法）
func export(headers []string, rows [][]any) (*excelize.File, error) {
	f := excelize.NewFile()

	boldStyle, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
		},
	})
	if err != nil {
		return nil, err
	}

	if err := writeSheet(f, "Sheet1", headers, rows, boldStyle); err != nil {
		return nil, err
	}

	return f, nil
}

// exportMultiSheet 导出多 Sheet Excel 文件（内部方法）
func exportMultiSheet(sheets []Sheet) (*excelize.File, error) {
	f := excelize.NewFile()

	boldStyle, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
		},
	})
	if err != nil {
		return nil, err
	}

	for i, sheet := range sheets {
		sheetName := sheet.Name
		if sheetName == "" {
			sheetName = "Sheet1"
		}

		if i == 0 {
			// 重命名默认 Sheet
			if err := f.SetSheetName("Sheet1", sheetName); err != nil {
				return nil, err
			}
		} else {
			// 创建新 Sheet
			if _, err := f.NewSheet(sheetName); err != nil {
				return nil, err
			}
		}

		if err := writeSheet(f, sheetName, sheet.Headers, sheet.Rows, boldStyle); err != nil {
			return nil, err
		}
	}

	return f, nil
}

// ExportToFile 导出 Excel 并保存到文件
// filename: 文件路径（含扩展名 .xlsx）
// headers: 标题行
// rows: 数据行
func ExportToFile(filename string, headers []string, rows [][]any) error {
	f, err := export(headers, rows)
	if err != nil {
		return err
	}
	defer f.Close()

	return f.SaveAs(filename)
}

// ExportToBytes 导出 Excel 并返回字节数据
// headers: 标题行
// rows: 数据行
func ExportToBytes(headers []string, rows [][]any) ([]byte, error) {
	f, err := export(headers, rows)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	buf, err := f.WriteToBuffer()
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// ExportMultiSheetToFile 导出多 Sheet Excel 并保存到文件
// filename: 文件路径（含扩展名 .xlsx）
// sheets: Sheet 数据列表
func ExportMultiSheetToFile(filename string, sheets []Sheet) error {
	f, err := exportMultiSheet(sheets)
	if err != nil {
		return err
	}
	defer f.Close()

	return f.SaveAs(filename)
}

// ExportMultiSheetToBytes 导出多 Sheet Excel 并返回字节数据
// sheets: Sheet 数据列表
func ExportMultiSheetToBytes(sheets []Sheet) ([]byte, error) {
	f, err := exportMultiSheet(sheets)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	buf, err := f.WriteToBuffer()
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
