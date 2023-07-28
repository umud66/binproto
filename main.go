package main

import (
	"os"

	"umud.online/bin/generate"
)

func main() {
	codes := ReadFile("./test.bk")
	// os.WriteFile("./out.go", []byte(generate.GenerateGOFile(codes)), os.ModePerm)
	os.WriteFile("./out.cs", []byte(generate.GenerateCSFile(codes)), os.ModePerm)
}
