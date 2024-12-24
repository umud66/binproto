package utils

import (
	"flag"
)

type EnvVal struct {
	InFile    string
	OutFile   string
	Filetype  string
	Gopackage string
	Username  string
	Password  string
	Host      string
	Db        string
	EndName   string
}

func ParseEnvVal() *EnvVal {
	envVal := &EnvVal{}
	flag.StringVar(&envVal.InFile, "i", "", "file Or Dir")
	flag.StringVar(&envVal.Filetype, "t", "", "Type: cs go ts")
	flag.StringVar(&envVal.OutFile, "o", "", "Out File(bk file input) Or Dir")
	flag.StringVar(&envVal.Gopackage, "gp", "binproto", "go package default:binproto")
	flag.StringVar(&envVal.EndName, "e", "bytes", "end Name")
	flag.StringVar(&envVal.Username, "u", "", "sql username")
	flag.StringVar(&envVal.Password, "p", "", "sql password")
	flag.StringVar(&envVal.Host, "h", "", "sql host")
	flag.StringVar(&envVal.Db, "d", "", "sql database")
	flag.Parse()
	return envVal
}
