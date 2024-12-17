package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"umud.online/bin/core"
	"umud.online/bin/generate"
)

func doCreateFile(inputFile string, outDir string, outIsDir bool, isDir bool, filetype string, filename string) {
	codes := core.ReadFile(inputFile)
	outFile := outDir
	if outIsDir {
		outFile = path.Join(outFile, strings.Replace(filename, "bk", filetype, 1))
	}
	if filetype == "cs" {
		os.WriteFile(outFile, []byte(generate.GenerateCSEnumFile(codes.Enums)+generate.GenerateCSFile(codes.Codes)), os.ModePerm)
	} else if filetype == "go" {
		err := os.WriteFile(outFile, []byte("package binproto\n"+generate.GenerateGOEnumFile(codes.Enums)+generate.GenerateGOConstFile(codes.Consts)+generate.GenerateGOFile(codes.Codes, false)), os.ModePerm)
		if err != nil {
			fmt.Println(err)
		}
	} else if filetype == "ts" {
		os.WriteFile(outFile, []byte((&generate.TSFileGenerate{
			Class:  codes.Codes,
			Enums:  codes.Enums,
			Consts: codes.Consts,
		}).WriteAll()), os.ModePerm)
		// os.WriteFile(outFile, []byte("import { BinProtoReader, BinProtoWriter } from './BinProto';\n"+generate.GenerateTSEnumFile(codes.Enums)+generate.GenerateTSFile(codes.Codes)), os.ModePerm)
	}
}

func doCreate(input string, inputIsDir bool, output string, outIsDir bool, filetype string) {
	if inputIsDir {
		files, _ := ioutil.ReadDir(input)
		for _, v := range files {
			if strings.HasSuffix(v.Name(), ".bk") {
				doCreateFile(path.Join(input, v.Name()), output, outIsDir, inputIsDir, filetype, v.Name())
			}
		}
	} else {
		doCreateFile(input, output, outIsDir, inputIsDir, filetype, path.Base(input))
	}
}

func main() {

	var inputDir string
	var outDir string
	var filetype string
	flag.StringVar(&inputDir, "i", "", "bk File Or Dir")
	flag.StringVar(&outDir, "o", "", "Out File(bk file input) Or Dir")
	flag.StringVar(&filetype, "t", "", "Type: cs go")
	flag.Parse()
	idStat, _ := os.Stat(inputDir)
	if path.Ext(inputDir) != ".bk" {
		if idStat == nil || !idStat.IsDir() {
			fmt.Println("Input Dir Is NotExist")
			return
		}
	} else {
		if idStat == nil {
			fmt.Println("Input bk File Is NotExist")
			return
		}
	}
	outIsDir := path.Ext(outDir) == ""
	odStat, _ := os.Stat(outDir)
	if odStat == nil && outIsDir {
		fmt.Println("Out Dir Is NotExist")
		return
	}

	if !outIsDir && idStat.IsDir() {
		fmt.Println("out must be dir")
		return
	}
	doCreate(inputDir, idStat.IsDir(), outDir, outIsDir, filetype)
}
