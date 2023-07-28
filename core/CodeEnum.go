package core

import "strings"

type CodeEnum struct {
	Name   string
	Values []int
	Names  []string
}

func (this *CodeEnum) WriteEnum(v string) {
	tmp := strings.Split(v, "#")
	name := tmp[0]
	this.Names = append(this.Names, name)
}
