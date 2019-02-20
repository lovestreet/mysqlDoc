package database

import (
	"database/sql"
	"fmt"
	"strings"
)

//GetTables 查询表信息
func GetTables(schema []string) []TableDefine {
	if len(schema) == 0 {
		return nil
	}

	var strSchemas string
	for _, item := range schema {
		strSchemas += fmt.Sprintf("'%v',", item)
	}
	strSchemas = strSchemas[:len(strSchemas)-1]

	var strSQL = fmt.Sprintf(`
select table_schema,table_name,table_comment
from information_schema.tables 
where lower( table_schema) in (%v) 
	and table_type = 'BASE TABLE'
order by table_schema,table_name`, strings.ToLower(strSchemas))

	var db = gMySQL.Database()
	if nil == db {
		return nil
	}

	rows, err := db.Query(strSQL)
	if err != nil {
		fmt.Printf("query database error :[%v]", err)
		return nil
	}

	var tables = make([]TableDefine, 0, 100)
	var result = getTableResult(rows)
	for result.Next() {
		var table TableDefine
		table.Schema = result.GetField("table_schema")
		table.Name = result.GetField("table_name")
		table.Comment = result.GetField("table_comment")
		tables = append(tables, table)
	}
	return tables
}

//GetColumns 查询表信息
func GetColumns(schema []string) []ColumnDefine {
	if len(schema) == 0 {
		return nil
	}

	var strSchemas string
	for _, item := range schema {
		strSchemas += fmt.Sprintf("'%v',", item)
	}
	strSchemas = strSchemas[:len(strSchemas)-1]

	var strSQL = fmt.Sprintf(`
	select  table_schema,table_name, column_name,column_type,is_nullable,column_default,column_comment
	from information_schema.columns c
	where lower( table_schema) in (%v) 
	order by table_schema,table_name, ordinal_position`, strings.ToLower(strSchemas))

	var db = gMySQL.Database()
	if nil == db {
		return nil
	}

	rows, err := db.Query(strSQL)
	if err != nil {
		fmt.Printf("query database error :[%v]", err)
		return nil
	}

	var columns = make([]ColumnDefine, 0, 20)
	var result = getTableResult(rows)
	for result.Next() {
		var column ColumnDefine
		column.Schema = result.GetField("table_schema")
		column.Table = result.GetField("table_name")
		column.ColumnName = result.GetField("column_name")
		column.ColumnType = result.GetField("column_type")
		column.Nullable = result.GetField("is_nullable")
		column.Default = result.GetField("column_default")
		column.Comment = result.GetField("column_comment")
		columns = append(columns, column)
	}
	return columns
}

//遍历查询结果
func getTableResult(rows *sql.Rows) (result *TableResult) {

	//构造列名数组
	var columns = make(map[string]int, 100)
	col, err := rows.Columns()
	if err != nil {
		fmt.Println("error is ", err.Error())
		return nil
	}
	for k, v := range col {
		columns[strings.ToLower(v)] = k
	}

	//查询数据
	var dbResult = make([]*TableRow, 0, 100)

	values := make([]sql.RawBytes, len(col))
	scanArgs := make([]interface{}, len(values))

	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		var row []string
		err := rows.Scan(scanArgs...)
		if err != nil {
			return nil
		}
		for _, v := range values {
			if v != nil {
				row = append(row, string(v))
			} else {
				row = append(row, "")
			}
		}

		//将row []string 转化为map[string]string
		var rowItem = make(map[string]string, 100)
		for col, idx := range columns {
			rowItem[col] = row[idx]
		}
		var rowValue = newTableRow(rowItem)

		dbResult = append(dbResult, rowValue)
	}

	return newTableResult(dbResult)
}
