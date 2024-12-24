package main

import (
	"umud.online/bin/method"
	"umud.online/bin/utils"
)

func main() {
	method.BKSqlMain(utils.ParseEnvVal())
}
