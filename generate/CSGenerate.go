package generate

import (
	"strings"

	"umud.online/bin/core"
)

func createCSRead(fieldName string, fieldType string) string {
	r := "data." + fieldName + " = "
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
		r = "int " + fieldName + "ArrSize = __r__.ReadInt32();\n\t\t" + r + "new " + baseType + "[" + fieldName + "ArrSize];\n"
		r += "\t\tfor(int i = 0;i < " + fieldName + "ArrSize;i++){\n"
		r += "\t\t\tdata." + fieldName + "[i] = " + baseType + ".DeSerializeReader(__r__);\n\t\t}\n"
		return r
	}
	return "data." + fieldName + " = " + fieldType + ".DeSerializeReader(__r__);\n"
}

func createCSWrite(fieldName string, fieldType string) string {
	r := ""
	if fieldType == "int" {
		return r + "__w__.WriteInt32(" + fieldName + ");\n"
	} else if fieldType == "uint" {
		return r + "__w__.WriteUInt32(" + fieldName + ");\n"
	} else if fieldType == "byte" {
		return r + "__w__.WriteByte(" + fieldName + ");\n"
	} else if fieldType == "int8" {
		return r + "__w__.Writeint8(" + fieldName + ");\n"
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
		return r + "__w__.WriteUInt8Array(" + fieldName + ");\n"
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
		r = "__w__.WriteInt32(" + fieldName + ".Length);\n"
		r += "\t\tfor(int i = 0;i < " + fieldName + ".Length;i++){\n"
		r += "\t\t\t__w__.WriteBytes(" + fieldName + "[i].Serialize());\n\t\t}\n"
		return r
	}
	return r + "__w__.WriteBytes(" + fieldName + ".Serialize());\n"
}

func convCSType(typename string) string {
	isArr := false
	if strings.HasSuffix(typename, "[]") {
		isArr = true
		typename = strings.Replace(typename, "[]", "", -1)
	}
	if typename == "uint16" {
		typename = "ushort"
	} else if typename == "int16" {
		typename = "short"
	} else if typename == "int8" {
		typename = "sbyte"
	} else if typename == "int64" {
		typename = "long"
	} else if typename == "uint64" {
		typename = "ulong"
	}
	if isArr {
		return typename + "[]"
	}
	return typename
}

func GenerateCSFile(codes []core.CodeClass) string {
	sb := &strings.Builder{}
	for _, v := range codes {
		sb.WriteString("public class ")
		sb.WriteString(v.Name)
		sb.WriteString("{")
		sb.WriteString("\n")
		writeFuncStr := &strings.Builder{}
		writeFuncStr.WriteString("\tpublic byte[] Serialize(){\n")
		writeFuncStr.WriteString("\t\tBinProto.BinWriter __w__ = new BinProto.BinWriter();\n")
		funcStr := &strings.Builder{}
		funcStr.WriteString("\tpublic static " + v.Name + " DeSerialize(byte[] bytes){\n")
		funcStr.WriteString("\t\treturn DeSerializeReader(new BinProto.BinReader(bytes));\n")
		funcStr.WriteString("\t}\n")
		funcStr.WriteString("\tpublic static " + v.Name + " DeSerializeReader(BinProto.BinReader __r__){\n")
		funcStr.WriteString("\t\t" + v.Name + " data = new " + v.Name + "();\n")
		for i, typename := range v.Types {
			sb.WriteString("\tpublic ")
			sb.WriteString(convCSType(typename))
			sb.WriteString(" ")
			name := strings.Split(v.Names[i], "#")
			sb.WriteString(name[0] + ";")
			if len(name) == 2 {
				sb.WriteString("//" + name[1])
			}
			sb.WriteString("\n")
			funcStr.WriteString("\t\t")
			funcStr.WriteString(createCSRead(name[0], typename))
			writeFuncStr.WriteString("\t\t")
			writeFuncStr.WriteString(createCSWrite(name[0], typename))
		}
		writeFuncStr.WriteString("\t\treturn __w__.GetBytes();\n\t}\n")
		funcStr.WriteString("\t\treturn data;\n")
		funcStr.WriteString("\t}\n")
		// sb.WriteString("}\n")
		// sb.WriteString("partial class BinProto {\n")
		sb.WriteString(funcStr.String())
		sb.WriteString(writeFuncStr.String())
		sb.WriteString("}\n")
	}
	return sb.String()
}
