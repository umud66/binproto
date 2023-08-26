package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"strings"
	"unsafe"

	_ "github.com/go-sql-driver/mysql"
	"umud.online/bin/core"
	buffertool "umud.online/bin/runtime/go"
)

var username string
var password string
var host string
var db string
var outFile string
var filetype string

func isLittleEndian() bool {
	var i int32 = 0x01020304
	u := unsafe.Pointer(&i)
	pd := (*byte)(u)
	b := *pd
	return (b == 0x04)
}

func getLink() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		username,
		password,
		host,
		"3306",
		db,
		"utf8")
}

func doDataTables(data []core.DataTable) {
	for _, table := range data {
		tmpData := &buffertool.ByteBufferWriter{}
		tmpData.WriteUInt32(table.Size)
		columnSize := len(table.ColsName)
		for _, row := range table.Rows {
			for i := 0; i < columnSize; i++ {
				data := row.Data[i]
				typeName := table.ColsType[i]
				if typeName == "uint8" || typeName == "byte" {
					if data == nil {
						tmpData.WriteUInt8(0)
					} else {
						tmpData.WriteUInt8(*data.(*uint8))
					}
				} else if typeName == "int8" {
					if data == nil {
						tmpData.WriteUInt8(0)
					} else {
						tmpData.WriteUInt8(uint8(*data.(*int8)))
					}
				} else if typeName == "string" {
					if data == nil {
						tmpData.WriteString("")
					} else {
						tmpData.WriteString(*data.(*string))
					}
				} else if typeName == "uint16" {
					if data == nil {
						tmpData.WriteUInt16(0)
					} else {
						tmpData.WriteUInt16(*data.(*uint16))
					}
				} else if typeName == "int16" {
					if data == nil {
						tmpData.WriteInt16(0)
					} else {
						tmpData.WriteInt16(*data.(*int16))
					}
				} else if typeName == "uint" {
					if data == nil {
						tmpData.WriteUInt32(0)
					} else {
						tmpData.WriteUInt32(*data.(*uint))
					}
				} else if typeName == "int" {
					if data == nil {
						tmpData.WriteInt32(0)
					} else {
						tmpData.WriteInt32(*data.(*int))
					}
				} else if typeName == "int64" {
					if data == nil {
						tmpData.WriteInt64(0)
					} else {
						tmpData.WriteInt64(*data.(*int64))
					}
				} else if typeName == "uint64" {
					if data == nil {
						tmpData.WriteUInt64(0)
					} else {
						tmpData.WriteUInt64(*data.(*uint64))
					}
				} else {

					fmt.Println("unkone data type", typeName, &data)
				}
			}
		}
		if table.Name == "data_building" {
			testReader(tmpData.GetBytes())
		}
		os.WriteFile(outFile+"/"+table.Name+".bytes", tmpData.GetBytes(), os.ModePerm)
	}
	// if filetype == "go" {
	// 	os.WriteFile(outFile, []byte("package binproto\n"+generate.GenerateGOFile(codes, false)), os.ModePerm)
	// } else if filetype == "cs" {
	// 	os.WriteFile(outFile, []byte(generate.GenerateCSFile(codes)), os.ModePerm)
	// } else if filetype == "godb" {
	// 	os.WriteFile(outFile, []byte("package binproto\n"+generate.GenerateGOFile(codes, true)), os.ModePerm)
	// }
}
func testReader(data []byte) {
	fmt.Println(data)
	reader := &buffertool.ByteBufferReader{B: data}
	size := reader.ReadUInt32()
	fmt.Println(size)
	for i := 0; i < int(size); i++ {
		fmt.Println("id:", reader.ReadUInt32())
		fmt.Println("name:", reader.ReadString())
		fmt.Println("pos:", reader.ReadUInt8())
		fmt.Println("icon:", reader.ReadString())
		fmt.Println("inhome:", reader.ReadUInt8())
		fmt.Println("model:", reader.ReadString())
	}
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

	flag.StringVar(&username, "u", "", "sql username")
	flag.StringVar(&password, "p", "", "sql password")
	flag.StringVar(&host, "h", "", "sql host")
	flag.StringVar(&db, "d", "", "sql database")
	flag.StringVar(&outFile, "o", "", "Out File")
	flag.StringVar(&filetype, "t", "", "Type: cs go")
	flag.Parse()

	sqldb, err := sql.Open("mysql", getLink())
	defer sqldb.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	sr, err := sqldb.Query("show tables")
	if err != nil {
		fmt.Println(err)
		return
	}
	tables := make([]string, 0)
	defer sr.Close()
	for sr.Next() {
		var tablesName string
		sr.Scan(&tablesName)
		tables = append(tables, tablesName)
	}
	dataTables := make([]core.DataTable, 0)
	for _, v := range tables {
		//query table struct
		tsr, _ := sqldb.Query(`SELECT COLUMN_NAME, COLUMN_TYPE FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_NAME = "` + v + `"`)
		defer tsr.Close()
		dataTable := &core.DataTable{}
		dataTable.Name = v
		var filedName string
		var dataType string
		for tsr.Next() {
			tsr.Scan(&filedName, &dataType)
			dataTable.ColsName = append(dataTable.ColsName, filedName)
			dataTable.ColsType = append(dataTable.ColsType, parseSqlType(dataType))
		}
		colNames := ""
		for i, c := range dataTable.ColsName {
			colNames += c
			if i < len(dataTable.ColsName)-1 {
				colNames += ","
			}
		}
		dataTsr, err := sqldb.Query("select " + colNames + " from " + v)
		if err != nil {
			fmt.Println(err)
		}
		defer dataTsr.Close()
		for dataTsr.Next() {
			vals := make([][]byte, len(dataTable.ColsName))
			dataTable.Size++
			dataRow := core.DataTableRow{}
			dataRow.Data = make([]interface{}, len(dataTable.ColsName))
			for k, _ := range vals {
				if dataTable.ColsType[k] == "uint8" {
					var v uint8
					dataRow.Data[k] = &v
				} else if dataTable.ColsType[k] == "int8" {
					var v int8
					dataRow.Data[k] = &v
				} else if dataTable.ColsType[k] == "int16" {
					var v int16
					dataRow.Data[k] = &v
				} else if dataTable.ColsType[k] == "uint16" {
					var v uint16
					dataRow.Data[k] = &v
				} else if dataTable.ColsType[k] == "uint" {
					var v uint
					dataRow.Data[k] = &v
				} else if dataTable.ColsType[k] == "int" {
					var v int
					dataRow.Data[k] = &v
				} else if dataTable.ColsType[k] == "int64" {
					var v int64
					dataRow.Data[k] = &v
				} else if dataTable.ColsType[k] == "uint64" {
					var v int64
					dataRow.Data[k] = &v
				} else if dataTable.ColsType[k] == "string" {
					var v string
					dataRow.Data[k] = &v
				}

				// dataRow.Data[k] = &vals[k]
			}
			err = dataTsr.Scan(dataRow.Data...)
			dataTable.Rows = append(dataTable.Rows, dataRow)
		}
		dataTables = append(dataTables, *dataTable)
	}
	doDataTables(dataTables)
}
