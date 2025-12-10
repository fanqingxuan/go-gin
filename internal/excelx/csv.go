package excelx

import (
	"bytes"
	"encoding/csv"
	"os"
)

// ExportCSVToFile 导出 CSV 并保存到文件
// filename: 文件路径（含扩展名 .csv）
// headers: 标题行
// rows: 数据行
func ExportCSVToFile(filename string, headers []string, rows [][]string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// 写入 UTF-8 BOM，确保 Excel 正确识别中文
	if _, err := file.Write([]byte{0xEF, 0xBB, 0xBF}); err != nil {
		return err
	}

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// 写入标题行
	if err := writer.Write(headers); err != nil {
		return err
	}

	// 写入数据行
	for _, row := range rows {
		if err := writer.Write(row); err != nil {
			return err
		}
	}

	return nil
}

// ExportCSVToBytes 导出 CSV 并返回字节数据
// headers: 标题行
// rows: 数据行
func ExportCSVToBytes(headers []string, rows [][]string) ([]byte, error) {
	buf := &bytes.Buffer{}

	// 写入 UTF-8 BOM
	buf.Write([]byte{0xEF, 0xBB, 0xBF})

	writer := csv.NewWriter(buf)

	// 写入标题行
	if err := writer.Write(headers); err != nil {
		return nil, err
	}

	// 写入数据行
	for _, row := range rows {
		if err := writer.Write(row); err != nil {
			return nil, err
		}
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
