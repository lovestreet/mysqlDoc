package database

//TableDefine 表定义
type TableDefine struct {
	Schema  string
	Name    string
	Comment string
}

//ColumnDefine 表列定义
type ColumnDefine struct {
	Schema     string
	Table      string
	ColumnName string
	ColumnType string
	Nullable   string
	Default    string
	Comment    string
}
