package method

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"umud.online/bin/core"
	"umud.online/bin/generate"
	"umud.online/bin/utils"
)

type bkSqlVal struct {
	envData *utils.EnvVal
	// username  string
	// password  string
	// host      string
	// db        string
	// outFile   string
	// filetype  string
	// gopackage string
}

func (this *bkSqlVal) getLink() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		this.envData.Username,
		this.envData.Password,
		this.envData.Host,
		"3306",
		this.envData.Db,
		"utf8")
}

func (this *bkSqlVal) doTables(codes []core.CodeClass) {
	if this.envData.Filetype == "go" {
		os.WriteFile(this.envData.OutFile, []byte("package "+this.envData.Gopackage+"\n"+generate.GenerateGOFile(codes, false)), os.ModePerm)
	} else if this.envData.Filetype == "cs" {
		os.WriteFile(this.envData.OutFile, []byte(generate.GenerateCSFile(codes)), os.ModePerm)
	} else if this.envData.Filetype == "godb" {
		os.WriteFile(this.envData.OutFile, []byte("package "+this.envData.Gopackage+"\n"+generate.GenerateGOFile(codes, true)), os.ModePerm)
	} else if this.envData.Filetype == "ts" {
		os.WriteFile(this.envData.OutFile, []byte((&generate.TSFileGenerate{
			Class: codes,
		}).WriteAll()), os.ModePerm)
	}
}

func (this *bkSqlVal) parseSqlType(typeName string) string {
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

func BKSqlMain(envData *utils.EnvVal) {
	envVal := &bkSqlVal{}
	envVal.envData = envData
	// flag.StringVar(&envVal.username, "u", "", "sql username")
	// flag.StringVar(&envVal.password, "p", "", "sql password")
	// flag.StringVar(&envVal.host, "h", "", "sql host")
	// flag.StringVar(&envVal.db, "d", "", "sql database")
	// flag.StringVar(&envVal.outFile, "o", "", "Out File")
	// flag.StringVar(&envVal.filetype, "t", "", "Type: cs go")
	// flag.StringVar(&envVal.gopackage, "gp", "binproto", "go package default:binproto")
	// flag.Parse()

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
	structs := make([]core.CodeClass, 0)
	for _, v := range tables {
		//query table struct
		tsr, _ := sqldb.Query(`SELECT COLUMN_NAME, COLUMN_TYPE, COLUMN_COMMENT,COLUMN_KEY,EXTRA FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_NAME = "` + v + `"`)
		defer tsr.Close()
		code := core.CodeClass{
			Name: v,
		}
		var filedName string
		var dataType string
		var comment string
		var pk string
		var extra string
		for tsr.Next() {
			tsr.Scan(&filedName, &dataType, &comment, &pk, &extra)
			if comment == "" {
				code.Put(filedName, envVal.parseSqlType(dataType))
			} else {
				code.Put(filedName+"#"+comment, envVal.parseSqlType(dataType))
			}
			tag := "`xorm:\"" + filedName
			if pk == "PRI" {
				tag += " pk"
			}
			if extra == "auto_increment" {
				tag += " autoincr"

			}
			tag += "\"`"
			code.PutTag(tag)

		}
		structs = append(structs, code)
	}
	envVal.doTables(structs)
}
