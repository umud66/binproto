package core

type DataTable struct {
	Name     string
	Size     uint
	Rows     []DataTableRow
	ColsName []string
	ColsType []string
}
type DataTableRow struct {
	Data []interface{}
}
