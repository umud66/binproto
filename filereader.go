package main

import (
	"os"
	"strings"

	"umud.online/bin/core"
	"umud.online/bin/utils"
)

type CodeStruct struct {
	Codes []core.CodeClass
	Enums []core.CodeEnum
}

func doGenerateCodes(contArr []string) CodeStruct {
	codeArr := make([]core.CodeClass, 0)
	codeEnum := make([]core.CodeEnum, 0)
	tmpCode := &core.CodeClass{}
	tmpEnum := &core.CodeEnum{}
	findEnd := false
	findEnum := false
	for _, v := range contArr {
		if strings.HasPrefix(v, "struct") {
			findEnd = true
			tmpCode = &core.CodeClass{
				Name: strings.Replace(utils.TrimStr(strings.Split(v, ":")[1]), "{", "", -1),
			}
			continue
		}
		if strings.HasPrefix(v, "enum") {
			findEnd = true
			findEnum = true
			tmpEnum = &core.CodeEnum{
				Name: strings.Replace(utils.TrimStr(strings.Split(v, ":")[1]), "{", "", -1),
			}
			continue
		}
		if strings.HasPrefix(v, "}") {
			if findEnum {
				codeEnum = append(codeEnum, *tmpEnum)
			} else {
				codeArr = append(codeArr, *tmpCode)
			}
			findEnd = false
			findEnum = false
			continue
		}
		if findEnd {
			if findEnum {
				// tmpEnum
				tmpEnum.WriteEnum(utils.TrimStr(v))
			} else {
				tmpCode.PutField(utils.TrimStr(v))
			}
		}
	}
	return CodeStruct{
		Codes: codeArr,
		Enums: codeEnum,
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
