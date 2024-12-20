package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/xuri/excelize/v2"
	"umud.online/bin/core"
	"umud.online/bin/generate"
)

var inFile string
var outFile string
var filetype string
var gopackage string

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

	flag.StringVar(&inFile, "i", "", "excelfile")
	flag.StringVar(&outFile, "o", "", "Out File")
	flag.StringVar(&filetype, "t", "", "Type: cs go")
	flag.StringVar(&gopackage, "gp", "binproto", "go package default:binproto")
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
	structs := make([]core.CodeClass, 0)
	for _, v := range tables {
		//query table struct

		tsr, _ := excel.GetRows(v)
		var cols = len(tsr[0])
		code := core.CodeClass{
			Name: v,
		}
		for i := 0; i < cols; i++ {
			if len(tsr[2]) <= i {
				code.Put(tsr[0][i], tsr[1][i])
			} else {
				code.Put(tsr[0][i]+"#"+tsr[2][i], tsr[1][i])
			}
		}
		structs = append(structs, code)
	}
	doTables(structs)
}
