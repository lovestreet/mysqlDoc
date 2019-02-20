package docx

import (
	"fmt"
	"strings"

	"baliance.com/gooxml/document"
	"baliance.com/gooxml/measurement"
	"baliance.com/gooxml/schema/soo/wml"
	"github.com/lovestreet/mysqlDoc/database"
)

const (
	fontSizeCell         = 10 * measurement.Point      //表格中字体大小
	tableWidth           = 15 * measurement.Centimeter //表格宽度
	fontSizeTableSummary = 12 * measurement.Point
	cellWidthName        = 2.2 * measurement.Centimeter //列宽
	cellWidthType        = 3.3 * measurement.Centimeter //列宽
	cellWidthNullable    = 1.5 * measurement.Centimeter //列宽
	cellWidthDefault     = 3.0 * measurement.Centimeter //列宽
	//cellWidthComment     = 5 * measurement.Centimeter   //列宽
)

//WriteTables 写表以及表字段信息
func WriteTables(file string, tables []database.TableDefine, columns []database.ColumnDefine) {
	var doc = document.New()
	for _, table := range tables {
		//根据table获取table下的column列表
		var tableColumns = make([]database.ColumnDefine, 0, 100)
		for _, column := range columns {
			if strings.ToLower(column.Schema) == strings.ToLower(table.Schema) &&
				strings.ToLower(column.Table) == strings.ToLower(table.Name) {
				tableColumns = append(tableColumns, column)
			}
		}
		addTable(doc, table, tableColumns)
	}

	doc.SaveToFile(file)
}

//增加一个表格，描述数据库中的一个表定义信息
func addTable(doc *document.Document, t database.TableDefine, c []database.ColumnDefine) {
	addTableName(doc, t.Name, t.Comment)

	table := doc.AddTable()
	addTableTitle(table)
	for _, item := range c {
		addTableRow(table, item)
	}
}

//增加一个表的摘要信息，即一行文字
func addTableName(doc *document.Document, tableName, comment string) {
	var run = doc.AddParagraph().AddRun()
	run.Properties().SetSize(fontSizeTableSummary)
	run.AddText(fmt.Sprintf("%v %v", tableName, comment))
}

//为表格增加表头
func addTableTitle(table document.Table) {
	//表格格式,
	table.Properties().SetWidth(tableWidth)
	table.Properties().SetAlignment(wml.ST_JcTableCenter)

	//增加一个标题
	title := table.AddRow()

	var cell = title.AddCell()
	var run = cell.AddParagraph().AddRun()
	run.AddText("名称")
	run.Properties().SetBold(true)
	run.Properties().SetSize(fontSizeCell)
	cell.Properties().SetWidth(cellWidthName)

	cell = title.AddCell()
	run = cell.AddParagraph().AddRun()
	run.AddText("字段类型")
	run.Properties().SetBold(true)
	run.Properties().SetSize(fontSizeCell)
	cell.Properties().SetWidth(cellWidthType)

	cell = title.AddCell()
	run = cell.AddParagraph().AddRun()
	run.AddText("必填")
	run.Properties().SetBold(true)
	run.Properties().SetSize(fontSizeCell)
	cell.Properties().SetWidth(cellWidthNullable)

	cell = title.AddCell()
	run = cell.AddParagraph().AddRun()
	run.AddText("默认值")
	run.Properties().SetBold(true)
	run.Properties().SetSize(fontSizeCell)
	cell.Properties().SetWidth(cellWidthDefault)

	cell = title.AddCell()
	run = cell.AddParagraph().AddRun()
	run.AddText("字段类型")
	run.Properties().SetBold(true)
	run.Properties().SetSize(fontSizeCell)
	cell.Properties().SetWidthAuto()
}

//为表格增加一行数据
func addTableRow(table document.Table, c database.ColumnDefine) {
	row := table.AddRow()

	var cell = row.AddCell()
	var run = cell.AddParagraph().AddRun()
	run.Properties().SetSize(fontSizeCell)
	run.AddText(c.ColumnName)
	cell.Properties().SetWidth(cellWidthName)

	cell = row.AddCell()
	run = cell.AddParagraph().AddRun()
	run.Properties().SetSize(fontSizeCell)
	run.AddText(c.ColumnType)
	cell.Properties().SetWidth(cellWidthType)

	cell = row.AddCell()
	run = cell.AddParagraph().AddRun()
	run.Properties().SetSize(fontSizeCell)
	run.AddText(c.Nullable)
	cell.Properties().SetWidth(cellWidthNullable)

	cell = row.AddCell()
	run = cell.AddParagraph().AddRun()
	run.Properties().SetSize(fontSizeCell)
	run.AddText(c.Default)
	cell.Properties().SetWidth(cellWidthDefault)

	cell = row.AddCell()
	run = cell.AddParagraph().AddRun()
	run.Properties().SetSize(fontSizeCell)
	run.AddText(c.Comment)
}
