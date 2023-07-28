package core

import (
	"strconv"
	"strings"
)

type CodeEnum struct {
	Name   string
	Values []int
	Names  []string
}

func (this *CodeEnum) WriteEnum(v string) {
	tmp := strings.Split(v, "#")
	v1 := strings.Split(tmp[0], "=")
	if len(tmp) == 2 {
		v1[0] += "#" + tmp[1]
	}
	this.Names = append(this.Names, v1[0])
	if len(v1) == 2 {
		v2, _ := strconv.Atoi(v1[1])
		this.Values = append(this.Values, v2)
	} else {
		this.Values = append(this.Values, -1)
	}
}
