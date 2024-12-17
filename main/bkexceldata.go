package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
	"umud.online/bin/core"
	buffertool "umud.online/bin/runtime/go"
)

var inFile string
var outFile string
var filetype string
var endName string

func doDataTables(data []core.DataTable) {
	for _, table := range data {
		tmpData := &buffertool.ByteBufferWriter{}
		tmpData.WriteUInt32(table.Size)
		columnSize := len(table.ColsName)
		for _, row := range table.Rows {

			for i := 0; i < columnSize; i++ {
				data := row.Data[i]
				typeName := table.ColsType[i]
				if typeName == "string" {
					if data == nil {
						tmpData.WriteString("")
					} else {
						tmpData.WriteString(data.(string))
					}
				} else {
					var v uint64 = 0
					if data != nil {
						v, _ = strconv.ParseUint(data.(string), 10, 64)
					}
					if typeName == "uint8" || typeName == "byte" || typeName == "int8" {
						tmpData.WriteUInt8(uint8(v))
					} else if typeName == "int16" {
						tmpData.WriteInt16(int16(v))
					} else if typeName == "uint16" {
						tmpData.WriteUInt16(uint16(v))
					} else if typeName == "int" {
						tmpData.WriteInt32(int(v))
					} else if typeName == "uint" {
						tmpData.WriteUInt32(uint(v))
					} else if typeName == "int64" {
						tmpData.WriteInt64(int64(v))
					} else if typeName == "uint64" {
						tmpData.WriteUInt64(uint64(v))
					}
				}
			}
		}
		os.WriteFile(outFile+"/"+table.Name+"."+endName, tmpData.GetBytes(), os.ModePerm)
	}
	// if filetype == "go" {
	// 	os.WriteFile(outFile, []byte("package binproto\n"+generate.GenerateGOFile(codes, false)), os.ModePerm)
	// } else if filetype == "cs" {
	// 	os.WriteFile(outFile, []byte(generate.GenerateCSFile(codes)), os.ModePerm)
	// } else if filetype == "godb" {
	// 	os.WriteFile(outFile, []byte("package binproto\n"+generate.GenerateGOFile(codes, true)), os.ModePerm)
	// }
}

func parseSqlType(typeName string) string {
	unsigned := strings.Index(typeName, "unsigned") > 0
	if strings.HasPrefix(typeName, "tinyint") {
		if unsigned {
			return "byte"
		} else {
			return "int8"
		}
	} else if strings.HasPrefix(typeName, "varchar") || strings.HasPrefix(typeName, "text") {
		return "string"
	} else if typeName == "timestamp" {
		return "int64"
	} else if strings.HasPrefix(typeName, "bigint") {
		if unsigned {
			return "uint64"
		} else {
			return "int64"
		}
	} else if strings.HasPrefix(typeName, "int") {
		if unsigned {
			return "uint"
		} else {
			return "int"
		}
	} else if strings.HasPrefix(typeName, "smallint") {
		if unsigned {
			return "uint16"
		} else {
			return "int16"
		}
	}
	return typeName
}

func main() {

	flag.StringVar(&inFile, "i", "", "excelfile")
	flag.StringVar(&outFile, "o", "", "Out File")
	flag.StringVar(&filetype, "t", "", "Type: cs go")
	flag.StringVar(&endName, "e", "bytes", "end Name")
	flag.Parse()

	excel, err := excelize.OpenFile(inFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		// Close the spreadsheet.
		if err := excel.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	tables := excel.GetSheetList()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(tables)
	dataTables := make([]core.DataTable, 0)
	for _, v := range tables {
		tsr, _ := excel.GetRows(v)
		var cols = len(tsr[0])
		var rows = len(tsr)
		dataTable := &core.DataTable{
			Name: v,
		}

		for i := 0; i < cols; i++ {
			dataTable.ColsName = append(dataTable.ColsName, tsr[0][i])
			dataTable.ColsType = append(dataTable.ColsType, parseSqlType(tsr[1][i]))
		}
		for i := 3; i < rows; i++ {
			dataTable.Size++

			dataRow := core.DataTableRow{}
			dataRow.Data = make([]interface{}, cols)
			for j := 0; j < cols; j++ {
				dataRow.Data[j] = tsr[i][j]
			}
			dataTable.Rows = append(dataTable.Rows, dataRow)
		}
		dataTables = append(dataTables, *dataTable)
	}
	doDataTables(dataTables)
}
