package main

import (
	"os"

	"umud.online/bin/generate"
)

func main() {
	codes := ReadFile("./test.bk")
	os.WriteFile("./out.go", []byte(generate.GenerateGOEnumFile(codes.Enums)+generate.GenerateGOFile(codes.Codes)), os.ModePerm)
	os.WriteFile("./out.cs", []byte(generate.GenerateCSEnumFile(codes.Enums)+generate.GenerateCSFile(codes.Codes)), os.ModePerm)
}
