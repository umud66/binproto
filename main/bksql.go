package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"umud.online/bin/core"
	"umud.online/bin/generate"
)

var username string
var password string
var host string
var db string
var outFile string
var filetype string
var gopackage string

func getLink() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		username,
		password,
		host,
		"3306",
		db,
		"utf8")
}

func doTables(codes []core.CodeClass) {
	if filetype == "go" {
		os.WriteFile(outFile, []byte("package "+gopackage+"\n"+generate.GenerateGOFile(codes, false)), os.ModePerm)
	} else if filetype == "cs" {
		os.WriteFile(outFile, []byte(generate.GenerateCSFile(codes)), os.ModePerm)
	} else if filetype == "godb" {
		os.WriteFile(outFile, []byte("package "+gopackage+"\n"+generate.GenerateGOFile(codes, true)), os.ModePerm)
	} else if filetype == "ts" {
		os.WriteFile(outFile, []byte((&generate.TSFileGenerate{
			Class: codes,
		}).WriteAll()), os.ModePerm)
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
	flag.StringVar(&gopackage, "gp", "binproto", "go package default:binproto")
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
				code.Put(filedName, parseSqlType(dataType))
			} else {
				code.Put(filedName+"#"+comment, parseSqlType(dataType))
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
	doTables(structs)
}
