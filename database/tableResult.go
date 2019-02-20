package database

//TableResult 查询结果
type TableResult struct {
	value []*TableRow
	index int
}

//Rows 返回所有的行
func (r *TableResult) Rows() (value []*TableRow) {
	return r.value
}

//Next 数据指向下一行
func (r *TableResult) Next() bool {
	if r.index >= len(r.value) {
		return false
	}
	r.index++
	return true
}

//GetField 得到某一行，某一列的数据，配置Next()使用
func (r *TableResult) GetField(colName string) string {
	if r.index >= len(r.value) {
		return ""
	}

	return r.value[r.index].GetField(colName)
}

//newTableResult 生成TableResult
func newTableResult(value []*TableRow) *TableResult {
	var result = new(TableResult)
	result.value = value
	result.index = -1 //初始化,配合Next()使用
	return result
}

// ======================================

//TableRow 查询行结果
type TableRow struct {
	value map[string]string
}

//GetField 得到字段内容
func (r *TableRow) GetField(colName string) string {
	if len(r.value) == 0 {
		return ""
	}

	v, _ := r.value[colName]
	return v
}

//newTableRow 生成一行数据
func newTableRow(value map[string]string) *TableRow {
	var row = new(TableRow)
	row.value = value
	return row
}
