package core

import (
	"strings"

	"umud.online/bin/utils"
)

type CodeClass struct {
	Name  string
	Types []string
	Names []string
}

func (this *CodeClass) PutField(field string) {
	tmp := strings.Split(field, ":")
	this.Names = append(this.Names, utils.TrimStr(tmp[1]))
	this.Types = append(this.Types, utils.TrimStr(tmp[0]))
}
func (this *CodeClass) Put(field string, fieldType string) {
	this.Names = append(this.Names, field)
	this.Types = append(this.Types, fieldType)
}
