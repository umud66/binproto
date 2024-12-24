package main

import (
	"flag"

	"umud.online/bin/method"
	"umud.online/bin/utils"
)

func main() {
	var m string
	flag.StringVar(&m, "m", "", "code or excel or exceldata or sql or sqldata")
	data := utils.ParseEnvVal()
	switch m {
	case "code":
		method.BKCodeMain(data)
	case "excel":
		method.BKExcelMain(data)
	case "exceldata":
		method.BKExcelDataMain(data)
	case "sql":
		method.BKSqlMain(data)
	case "sqldata":
		method.BKSqlDataMain(data)
	}
}
