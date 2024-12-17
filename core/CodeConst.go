package core

import (
	"strings"
)

type CodeConst struct {
	Name      string
	Values    []string
	Names     []string
	ValueType string
}

func (this *CodeConst) WriteConst(v string) {
	tmp := strings.Split(v, "#")
	v1 := strings.Split(tmp[0], "=")
	if len(tmp) == 2 {
		v1[0] += "#" + tmp[1]
	}
	this.Names = append(this.Names, v1[0])
	this.Values = append(this.Values, v)
}
