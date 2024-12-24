package method

import (
	"fmt"
	"os"
	"strings"

	"github.com/xuri/excelize/v2"
	"umud.online/bin/core"
	"umud.online/bin/generate"
	"umud.online/bin/utils"
)

type bkExcelVal struct {
	envData *utils.EnvVal
}

func (this *bkExcelVal) doTables(codes []core.CodeClass) {
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

func (this *bkExcelVal) parseSqlType(typeName string) string {
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

func BKExcelMain(data *utils.EnvVal) {
	envVal := &bkExcelVal{}
	envVal.envData = data
	// flag.StringVar(&envVal.inFile, "i", "", "excelfile")
	// flag.StringVar(&envVal.outFile, "o", "", "Out File")
	// flag.StringVar(&envVal.filetype, "t", "", "Type: cs go")
	// flag.StringVar(&envVal.gopackage, "gp", "binproto", "go package default:binproto")
	// flag.Parse()

	excel, err := excelize.OpenFile(envVal.envData.InFile)
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
	envVal.doTables(structs)
}
