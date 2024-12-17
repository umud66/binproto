package core

import (
	"os"
	"strings"

	"umud.online/bin/utils"
)

type CodeStruct struct {
	Codes  []CodeClass
	Enums  []CodeEnum
	Consts []CodeConst
}

func doGenerateCodes(contArr []string) CodeStruct {
	codeArr := make([]CodeClass, 0)
	codeEnum := make([]CodeEnum, 0)
	codeConst := make([]CodeConst, 0)
	tmpCode := &CodeClass{}
	tmpEnum := &CodeEnum{}
	tmpConst := &CodeConst{}
	findEnd := false
	findEnum := false
	findConst := false
	for _, v := range contArr {
		if strings.HasPrefix(v, "struct") {
			findEnd = true
			tmpCode = &CodeClass{
				Name: strings.Replace(utils.TrimStr(strings.Split(v, ":")[1]), "{", "", -1),
			}
			continue
		}
		if strings.HasPrefix(v, "const") {
			findEnd = true
			findConst = true
			tmpConst = &CodeConst{
				Name:      utils.TrimStr(strings.Split(v, ":")[1]),
				ValueType: strings.Replace(utils.TrimStr(strings.Split(v, ":")[2]), "{", "", -1),
			}
			continue
		}
		if strings.HasPrefix(v, "enum") {
			findEnd = true
			findEnum = true
			tmpEnum = &CodeEnum{
				Name: strings.Replace(utils.TrimStr(strings.Split(v, ":")[1]), "{", "", -1),
			}
			continue
		}
		if strings.HasPrefix(v, "}") {
			if findEnum {
				codeEnum = append(codeEnum, *tmpEnum)
			} else if findConst {
				codeConst = append(codeConst, *tmpConst)
			} else {
				codeArr = append(codeArr, *tmpCode)
			}
			findEnd = false
			findEnum = false
			findConst = false
			continue
		}
		if findEnd {
			if findEnum {
				// tmpEnum
				tmpEnum.WriteEnum(utils.TrimStr(v))
			} else if findConst {
				tmpConst.WriteConst(utils.TrimStr(v))
			} else {
				tmpCode.PutField(utils.TrimStr(v))
			}
		}
	}
	return CodeStruct{
		Codes:  codeArr,
		Enums:  codeEnum,
		Consts: codeConst,
	}
}

func ReadFile(filePath string) CodeStruct {
	file, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	content := string(file)
	contentArr := strings.Split(content, "\n")
	return doGenerateCodes(contentArr)
}
