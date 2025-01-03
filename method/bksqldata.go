package method

import (
	"database/sql"
	"fmt"
	"os"
	"strings"
	"unsafe"

	_ "github.com/go-sql-driver/mysql"
	"umud.online/bin/core"
	buffertool "umud.online/bin/runtime/go"
	"umud.online/bin/utils"
)

type bkSqlDataVal struct {
	// username string
	// password string
	// host     string
	// db       string
	// outFile  string
	// filetype string
	// endName  string
	evnData *utils.EnvVal
}

func isLittleEndian() bool {
	var i int32 = 0x01020304
	u := unsafe.Pointer(&i)
	pd := (*byte)(u)
	b := *pd
	return (b == 0x04)
}

func (this *bkSqlDataVal) getLink() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		this.evnData.Username,
		this.evnData.Password,
		this.evnData.Host,
		"3306",
		this.evnData.Db,
		"utf8")
}

func (this *bkSqlDataVal) doDataTables(data []core.DataTable) {
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
						tmpData.WriteString(*data.(*string))
					}
				} else {
					var v uint = 0
					if data != nil {
						v = *(data.(*uint))
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
		os.WriteFile(this.evnData.OutFile+"/"+table.Name+"."+this.evnData.EndName, tmpData.GetBytes(), os.ModePerm)
	}
	// if filetype == "go" {
	// 	os.WriteFile(outFile, []byte("package binproto\n"+generate.GenerateGOFile(codes, false)), os.ModePerm)
	// } else if filetype == "cs" {
	// 	os.WriteFile(outFile, []byte(generate.GenerateCSFile(codes)), os.ModePerm)
	// } else if filetype == "godb" {
	// 	os.WriteFile(outFile, []byte("package binproto\n"+generate.GenerateGOFile(codes, true)), os.ModePerm)
	// }
}
func (this *bkSqlDataVal) parseSqlType(typeName string) string {
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

func BKSqlDataMain(envData *utils.EnvVal) {
	envVal := &bkSqlDataVal{}
	// flag.StringVar(&envVal.username, "u", "", "sql username")
	// flag.StringVar(&envVal.password, "p", "", "sql password")
	// flag.StringVar(&envVal.host, "h", "", "sql host")
	// flag.StringVar(&envVal.db, "d", "", "sql database")
	// flag.StringVar(&envVal.outFile, "o", "", "Out File")
	// flag.StringVar(&envVal.filetype, "t", "", "Type: cs go")
	// flag.StringVar(&envVal.endName, "e", "bytes", "end Name")
	// flag.Parse()
	envVal.evnData = envData
	sqldb, err := sql.Open("mysql", envVal.getLink())
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
			dataTable.ColsType = append(dataTable.ColsType, envVal.parseSqlType(dataType))
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
				if dataTable.ColsType[k] == "string" {
					var v string
					dataRow.Data[k] = &v
				} else {
					var v uint
					dataRow.Data[k] = &v
				}
			}
			err = dataTsr.Scan(dataRow.Data...)
			dataTable.Rows = append(dataTable.Rows, dataRow)
		}
		dataTables = append(dataTables, *dataTable)
	}
	envVal.doDataTables(dataTables)
}
