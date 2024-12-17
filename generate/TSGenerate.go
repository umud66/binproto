package generate

import (
	"strconv"
	"strings"

	"umud.online/bin/core"
)

type TSFileGenerate struct {
	Enums  []core.CodeEnum
	Class  []core.CodeClass
	Consts []core.CodeConst
}

func (this *TSFileGenerate) WriteAll() string {
	return "import { BinProtoReader, BinProtoWriter } from './BinProto';\n" + this.GenerateTSEnumFile(this.Enums) + this.GenerateTSConstFile(this.Consts) + this.GenerateTSFile(this.Class) + this.writeExport()
}
func (this *TSFileGenerate) writeExport() string {
	sb := &strings.Builder{}
	sb.WriteString("export{")
	if this.Enums != nil {
		for _, r := range this.Enums {
			sb.WriteString(r.Name + ",")
		}
	}
	if this.Consts != nil {
		for _, r := range this.Consts {
			sb.WriteString(r.Name + ",")
		}
	}
	for _, r := range this.Class {
		sb.WriteString(r.Name + ",")
	}
	sb.WriteString("}")
	return sb.String()
}

func createTSRead(fieldName string, fieldType string) string {
	r := "this." + fieldName + " = "
	if fieldType == "int" {
		return r + "__r__.ReadInt32();\n"
	} else if fieldType == "uint" {
		return r + "__r__.ReadUInt32();\n"
	} else if fieldType == "byte" {
		return r + "__r__.ReadByte();\n"
	} else if fieldType == "int8" {
		return r + "__r__.ReadSByte();\n"
	} else if fieldType == "int16" {
		return r + "__r__.ReadInt16();\n"
	} else if fieldType == "uint16" {
		return r + "__r__.ReadUInt16();\n"
	} else if fieldType == "int64" {
		return r + "__r__.ReadInt64();\n"
	} else if fieldType == "uint64" {
		return r + "__r__.ReadUInt64();\n"
	} else if fieldType == "string" {
		return r + "__r__.ReadString();\n"
	} else if fieldType == "bool" {
		return r + "__r__.ReadBool();\n"
	} else if fieldType == "string[]" {
		return r + "__r__.ReadStringArray();\n"
	} else if fieldType == "byte[]" {
		return r + "__r__.ReadByteArray();\n"
	} else if fieldType == "int8[]" {
		return r + "__r__.ReadInt8Array();\n"
	} else if fieldType == "uint8[]" {
		return r + "__r__.ReadUInt8Array();\n"
	} else if fieldType == "int16[]" {
		return r + "__r__.ReadInt16Array();\n"
	} else if fieldType == "uint16[]" {
		return r + "__r__.ReadUInt16Array();\n"
	} else if fieldType == "int[]" {
		return r + "__r__.ReadInt32Array();\n"
	} else if fieldType == "uint[]" {
		return r + "__r__.ReadUInt32Array();\n"
	} else if fieldType == "uint64[]" {
		return r + "__r__.ReadUInt64Array();\n"
	} else if fieldType == "int64[]" {
		return r + "__r__.ReadInt64Array();\n"
	} else if fieldType == "bool[]" {
		return r + "__r__.ReadBoolArray();\n"
	} else if strings.HasSuffix(fieldType, "[]") {
		baseType := strings.Replace(fieldType, "[]", "", -1)
		r = "let " + fieldName + "ArrSize = __r__.ReadInt32();\n\t\t" + r + "[];\n"
		r += "\t\tfor(let i = 0;i < " + fieldName + "ArrSize;i++){\n"
		r += "\t\t\t let " + fieldName + "_temp = new " + baseType + "()\n"
		r += "\t\t\t " + fieldName + "_temp.DeSerializeReader(__r__);\n"
		r += "\t\t\t this." + fieldName + ".push(" + fieldName + "_temp);\n\t\t}\n"
		return r
	}
	return "data." + fieldName + " = " + fieldType + ".DeSerializeReader(__r__);\n"
}

func createTSWrite(fieldName string, fieldType string) string {
	r := ""
	fieldName = "this." + fieldName
	if fieldType == "int" {
		return r + "__w__.WriteInt32(" + fieldName + ");\n"
	} else if fieldType == "uint" {
		return r + "__w__.WriteUInt32(" + fieldName + ");\n"
	} else if fieldType == "byte" {
		return r + "__w__.WriteByte(" + fieldName + ");\n"
	} else if fieldType == "int8" {
		return r + "__w__.WriteSByte(" + fieldName + ");\n"
	} else if fieldType == "int16" {
		return r + "__w__.WriteInt16(" + fieldName + ");\n"
	} else if fieldType == "uint16" {
		return r + "__w__.WriteUInt16(" + fieldName + ");\n"
	} else if fieldType == "int64" {
		return r + "__w__.WriteInt64(" + fieldName + ");\n"
	} else if fieldType == "uint64" {
		return r + "__w__.WriteUInt64(" + fieldName + ");\n"
	} else if fieldType == "string" {
		return r + "__w__.WriteString(" + fieldName + ");\n"
	} else if fieldType == "bool" {
		return r + "__w__.WriteBool(" + fieldName + ");\n"
	} else if fieldType == "string[]" {
		return r + "__w__.WriteStringArray(" + fieldName + ");\n"
	} else if fieldType == "byte[]" {
		return r + "__w__.WriteByteArray(" + fieldName + ");\n"
	} else if fieldType == "int8[]" {
		return r + "__w__.WriteInt8Array(" + fieldName + ");\n"
	} else if fieldType == "int16[]" {
		return r + "__w__.WriteInt16Array(" + fieldName + ");\n"
	} else if fieldType == "uint16[]" {
		return r + "__w__.WriteUInt16Array(" + fieldName + ");\n"
	} else if fieldType == "int[]" {
		return r + "__w__.WriteInt32Array(" + fieldName + ");\n"
	} else if fieldType == "uint[]" {
		return r + "__w__.WriteUInt32Array(" + fieldName + ");\n"
	} else if fieldType == "uint64[]" {
		return r + "__w__.WriteUInt64Array(" + fieldName + ");\n"
	} else if fieldType == "int64[]" {
		return r + "__w__.WriteInt64Array(" + fieldName + ");\n"
	} else if fieldType == "bool[]" {
		return r + "__w__.WriteBoolArray(" + fieldName + ");\n"
	} else if strings.HasSuffix(fieldType, "[]") {
		r = "__w__.WriteInt32(" + fieldName + ".length);\n"
		r += "\t\tfor(let i = 0;i < " + fieldName + ".length;i++){\n"
		r += "\t\t\t__w__.WriteBytes(" + fieldName + "[i].Serialize());\n\t\t}\n"
		return r
	}
	return r + "__w__.WriteBytes(" + fieldName + ".Serialize());\n"
}

func convTSType(typename string) string {
	isArr := false
	if strings.HasSuffix(typename, "[]") {
		isArr = true
		typename = strings.Replace(typename, "[]", "", -1)
	}
	if typename == "uint16" {
		typename = "number"
	} else if typename == "int16" {
		typename = "number"
	} else if typename == "int8" {
		typename = "number"
	} else if typename == "int64" {
		typename = "number"
	} else if typename == "uint64" {
		typename = "number"
	} else if typename == "byte" {
		typename = "number"
	} else if typename == "int" {
		typename = "number"
	} else if typename == "uint" {
		typename = "number"
	}
	if isArr {
		return typename + "[]"
	}
	return typename
}

func generateConstructor(v core.CodeClass) string {
	sb := &strings.Builder{}
	sb.WriteString("\t public constructor(")

	bodyStr := &strings.Builder{}
	for i, typeName := range v.Types {
		name := strings.Split(v.Names[i], "#")[0]
		if i == 0 {
			sb.WriteString("byteOr" + name + ":Uint8Array | BinProtoReader | " + convTSType(typeName))
			sb.WriteString(" = null")
			bodyStr.WriteString("\t\tif(byteOr" + name + " == null)\n\t\t\treturn;\n")
			bodyStr.WriteString("\t\tif(byteOr" + name + " instanceof(Uint8Array)){\n\t\t\tthis.DeSerialize(byteOr" + name + " as Uint8Array)\n\t\t\treturn;\n\t\t}\n")
			bodyStr.WriteString("\t\tif(byteOr" + name + " instanceof(BinProtoReader)){\n\t\t\tthis.DeSerializeReader(byteOr" + name + " as BinProtoReader)\n\t\t\treturn;\n\t\t}\n")
			bodyStr.WriteString("\t\tthis." + name + " = byteOr" + name + " as " + convTSType(typeName))

		} else {
			sb.WriteString(name + ":" + convTSType(typeName) + " = null")
			bodyStr.WriteString("\t\tthis." + name + " = " + name)
		}
		if i < len(v.Types)-1 {
			sb.WriteString(",")
			bodyStr.WriteString("\n")
		}
	}
	sb.WriteString("){\n")
	sb.WriteString(bodyStr.String())
	sb.WriteString("\n")
	sb.WriteString("\t}")
	return sb.String()
}

func (this *TSFileGenerate) GenerateTSFile(codes []core.CodeClass) string {
	sb := &strings.Builder{}
	for _, v := range codes {
		sb.WriteString("class ")
		sb.WriteString(v.Name)
		sb.WriteString("{")
		sb.WriteString("\n")
		sb.WriteString(generateConstructor(v))
		sb.WriteString("\n")
		writeFuncStr := &strings.Builder{}
		writeFuncStr.WriteString("\tSerialize():number[]{\n")
		writeFuncStr.WriteString("\t\tlet __w__ = new BinProtoWriter();\n")
		funcStr := &strings.Builder{}
		funcStr.WriteString("\tDeSerialize(bytes:Uint8Array|number[]){\n")
		funcStr.WriteString("\t\treturn this.DeSerializeReader(new BinProtoReader(bytes));\n")
		funcStr.WriteString("\t}\n")

		funcStr.WriteString("\tpublic static DeSerializeStatic(bytes:Uint8Array|number[]):" + v.Name + "{\n")
		funcStr.WriteString("\t\tlet r = new " + v.Name + "();\n")
		funcStr.WriteString("\t\tr.DeSerialize(bytes);\n")
		funcStr.WriteString("\t\treturn r;\n")
		funcStr.WriteString("\t}\n")
		funcStr.WriteString("\tDeSerializeReader(__r__:BinProtoReader){\n")
		for i, typename := range v.Types {
			sb.WriteString("\t")
			name := strings.Split(v.Names[i], "#")
			sb.WriteString("public ")
			sb.WriteString(name[0])
			sb.WriteString(":")
			sb.WriteString(convTSType(typename))
			sb.WriteString(";")
			if len(name) == 2 {
				sb.WriteString("//" + name[1])
			}
			sb.WriteString("\n")
			funcStr.WriteString("\t\t")
			funcStr.WriteString(createTSRead(name[0], typename))
			writeFuncStr.WriteString("\t\t")
			writeFuncStr.WriteString(createTSWrite(name[0], typename))
		}
		writeFuncStr.WriteString("\t\treturn __w__.GetNumberArray();\n\t}\n")
		writeFuncStr.WriteString("\tSerializeUint8Array():Uint8Array{\n")
		writeFuncStr.WriteString("\t\treturn new Uint8Array(this.Serialize());\n\t}\n")
		funcStr.WriteString("\t}\n")
		sb.WriteString(funcStr.String())
		sb.WriteString(writeFuncStr.String())
		sb.WriteString("}\n")
	}
	return sb.String()
}

func (this *TSFileGenerate) GenerateTSEnumFile(enums []core.CodeEnum) string {
	sb := &strings.Builder{}
	for _, e := range enums {
		lastValue := 0
		sb.WriteString("enum " + e.Name + " {\n")
		for i, basename := range e.Names {
			names := strings.Split(basename, "#")
			sb.WriteString("\t" + names[0] + " = ")
			tmpValue := e.Values[i]
			if tmpValue == -1 {
				lastValue += 1
				sb.WriteString(strconv.Itoa(lastValue))
			} else {
				sb.WriteString(strconv.Itoa(tmpValue))
				lastValue = tmpValue
			}
			sb.WriteString(",")
			if len(names) == 2 {
				sb.WriteString(" //" + names[1])
			}
			sb.WriteString("\n")
		}
		sb.WriteString("}\n")

	}
	return sb.String()
}

func (this *TSFileGenerate) GenerateTSConstFile(consts []core.CodeConst) string {
	sb := &strings.Builder{}
	for _, e := range consts {
		sb.WriteString("class " + e.Name + " {\n")
		for i, basename := range e.Names {
			names := strings.Split(basename, "#")
			sb.WriteString("\tpublic static " + names[0] + " = ")
			tmpValue := e.Values[i]
			if strings.Index(tmpValue, "=") > 0 {
				tmpValue = strings.Split(tmpValue, "=")[1]
			}
			if e.ValueType == "string" {
				sb.WriteString("\"" + tmpValue + "\"")
			} else {
				sb.WriteString(tmpValue)
			}
			sb.WriteString(";")
			if len(names) == 2 {
				sb.WriteString(" //" + names[1])
			}
			sb.WriteString("\n")
		}
		sb.WriteString("}\n")

	}
	return sb.String()
}
