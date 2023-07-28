package main

import (
	"os"
	"strings"

	"umud.online/bin/core"
	"umud.online/bin/utils"
)

func doGenerateCodes(contArr []string) []core.CodeClass {
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
			} else {
				tmpCode.PutField(v)
			}
		}
	}
	return codeArr
}

func ReadFile(filePath string) []core.CodeClass {
	file, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	content := string(file)
	contentArr := strings.Split(content, "\n")
	return doGenerateCodes(contentArr)
}
