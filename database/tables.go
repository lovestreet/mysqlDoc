package database

import "fmt"

//查询表信息
func getTables(schema []string) {
	if len(schema) == 0 {
		return
	}

	var strSchemas string
	for _, item := range schema {
		strSchemas += fmt.Sprintf("'%v',", item)
	}
	strSchemas = strSchemas[:len(strSchemas)-1]

	var strSQL = fmt.Sprintf(`
select concat( table_schema,'.',table_name) as table_name,table_comment
from information_schema.tables 
where lower( table_schema) in (%v) 
	and table_type = 'BASE TABLE'
order by table_schema,table_name`, strSchemas)

}

func getColumns(schema []string) {
	if len(schema) == 0 {
		return
	}

	var strSchemas string
	for _, item := range schema {
		strSchemas += fmt.Sprintf("'%v',", item)
	}
	strSchemas = strSchemas[:len(strSchemas)-1]

	var strSQL = fmt.Sprintf(`
	select  table_schema,table_name, column_name,column_type,is_nullable,column_default,c.*
	from information_schema.columns c
	where lower( table_schema) in (%v) 
	order by table_schema,table_name, ordinal_position`, strSchemas)
}
