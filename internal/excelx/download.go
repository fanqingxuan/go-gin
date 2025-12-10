package excelx

import (
	"net/url"

	"github.com/gin-gonic/gin"
)

// Download 下载 Excel 文件
// ctx: gin.Context 或 httpx.Context
// filename: 文件名（含扩展名 .xlsx）
// headers: 标题行
// rows: 数据行
func Download(ctx *gin.Context, filename string, headers []string, rows [][]any) error {
	data, err := ExportToBytes(headers, rows)
	if err != nil {
		return err
	}

	ctx.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	ctx.Header("Content-Disposition", "attachment; filename*=UTF-8''"+url.QueryEscape(filename))
	ctx.Data(200, "application/octet-stream", data)
	return nil
}

// DownloadMultiSheet 下载多 Sheet Excel 文件
// ctx: gin.Context 或 httpx.Context
// filename: 文件名（含扩展名 .xlsx）
// sheets: Sheet 数据列表
func DownloadMultiSheet(ctx *gin.Context, filename string, sheets []Sheet) error {
	data, err := ExportMultiSheetToBytes(sheets)
	if err != nil {
		return err
	}

	ctx.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	ctx.Header("Content-Disposition", "attachment; filename*=UTF-8''"+url.QueryEscape(filename))
	ctx.Data(200, "application/octet-stream", data)
	return nil
}

// DownloadCSV 下载 CSV 文件
// ctx: gin.Context 或 httpx.Context
// filename: 文件名（含扩展名 .csv）
// headers: 标题行
// rows: 数据行
func DownloadCSV(ctx *gin.Context, filename string, headers []string, rows [][]string) error {
	data, err := ExportCSVToBytes(headers, rows)
	if err != nil {
		return err
	}

	ctx.Header("Content-Type", "text/csv; charset=utf-8")
	ctx.Header("Content-Disposition", "attachment; filename*=UTF-8''"+url.QueryEscape(filename))
	ctx.Data(200, "text/csv", data)
	return nil
}
