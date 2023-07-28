package utils

import "strings"

func TrimStr(str string) string {
	return strings.Replace(strings.Replace(strings.Trim(str, " "), "\r", "", -1), "\n", "", -1)
}
