package generate

import (
	"strconv"
	"strings"

	"umud.online/bin/core"
)

func createGORead(fieldName string, fieldType string) string {
	r := "r." + strings.Title(fieldName) + " = "
	if fieldType == "int" {
		return r + "reader.ReadInt32()\n"
	} else if fieldType == "uint" {
		return r + "reader.ReadUInt32()\n"
	} else if fieldType == "byte" || fieldType == "uin8" {
		return r + "reader.ReadUInt8()\n"
	} else if fieldType == "int16" {
		return r + "reader.ReadInt16()\n"
	} else if fieldType == "bool" {
		return r + "reader.ReadBool()\n"
	} else if fieldType == "uint16" {
		return r + "reader.ReadUInt16()\n"
	} else if fieldType == "string" {
		return r + "reader.ReadString()\n"
	} else if fieldType == "uint64" {
		return r + "reader.ReadUInt64()\n"
	} else if fieldType == "int64" {
		return r + "reader.ReadInt64()\n"
	} else if fieldType == "bool" {
		return r + "reader.ReadBool()\n"
	} else if fieldType == "string[]" {
		return r + "reader.ReadStringArr()\n"
	} else if fieldType == "byte[]" || fieldType == "uint8[]" {
		return r + "reader.ReadUInt8Arr()\n"
	} else if fieldType == "int[]" {
		return r + "reader.ReadInt32Arr()\n"
	} else if fieldType == "bool[]" {
		return r + "reader.ReadBoolArr()\n"
	} else if fieldType == "uint[]" {
		return r + "reader.ReadUInt32Arr();\n"
	} else if fieldType == "uint64[]" {
		return r + "reader.ReadUInt64Arr();\n"
	} else if fieldType == "int64[]" {
		return r + "reader.ReadInt64Arr();\n"
	} else if fieldType == "int16[]" {
		return r + "reader.ReadInt16Arr();\n"
	} else if fieldType == "uint16[]" {
		return r + "reader.ReadUInt16Arr();\n"
	} else if strings.HasSuffix(fieldType, "[]") {
		baseType := strings.Replace(fieldType, "[]", "", -1)
		r = fieldName + "ArrSize := reader.ReadInt32()\n\t\t" + r + "make([]" + baseType + "," + fieldName + "ArrSize);\n"
		r += "\t\tfor i := 0;i < " + fieldName + "ArrSize;i++ {\n"
		r += "\t\t\tr." + strings.Title(fieldName) + "[i] = DeSerialize" + strings.Title(baseType) + "(reader)\n\t\t}\n"
		return r
	}
	return r + " DeSerialize" + strings.Title(fieldType) + "(reader)\n"
}
func createGOWrite(fieldName string, fieldType string) string {
	if fieldType == "int" {
		return "writer.WriteInt32(this." + strings.Title(fieldName) + ")\n"
	} else if fieldType == "uint" {
		return "writer.WriteUInt32(this." + strings.Title(fieldName) + ")\n"
	} else if fieldType == "byte" || fieldType == "uint8" {
		return "writer.WriteUInt8(this." + strings.Title(fieldName) + ")\n"
	} else if fieldType == "int64" {
		return "writer.WriteInt64(this." + strings.Title(fieldName) + ")\n"
	} else if fieldType == "uint64" {
		return "writer.WriteUInt64(this." + strings.Title(fieldName) + ")\n"
	} else if fieldType == "bool" {
		return "writer.WriteBool(this." + strings.Title(fieldName) + ")\n"
	} else if fieldType == "int16" {
		return "writer.WriteInt16(this." + strings.Title(fieldName) + ")\n"
	} else if fieldType == "uint16" {
		return "writer.WriteUInt16(this." + strings.Title(fieldName) + ")\n"
	} else if fieldType == "string" {
		return "writer.WriteString(this." + strings.Title(fieldName) + ")\n"
	} else if fieldType == "int[]" {
		return "writer.WriteInt32Arr(this." + strings.Title(fieldName) + ")\n"
	} else if fieldType == "uint[]" {
		return "writer.WriteUInt32Arr(this." + strings.Title(fieldName) + ")\n"
	} else if fieldType == "byte[]" || fieldType == "uin8[]" {
		return "writer.WriteUInt8Arr(this." + strings.Title(fieldName) + ")\n"
	} else if fieldType == "int16[]" {
		return "writer.WriteInt16Arr(this." + strings.Title(fieldName) + ")\n"
	} else if fieldType == "uint16[]" {
		return "writer.WriteUInt16Arr(this." + strings.Title(fieldName) + ")\n"
	} else if fieldType == "int64[]" {
		return "writer.WriteInt64Arr(this." + strings.Title(fieldName) + ")\n"
	} else if fieldType == "uint64[]" {
		return "writer.WriteUInt64Arr(this." + strings.Title(fieldName) + ")\n"
	} else if fieldType == "string[]" {
		return "writer.WriteStringArr(this." + strings.Title(fieldName) + ")\n"
	} else if fieldType == "bool[]" {
		return "writer.WriteBoolArr(this." + strings.Title(fieldName) + ")\n"
	} else if strings.HasSuffix(fieldType, "[]") {
		r := fieldName + "ArrSize := len(this." + strings.Title(fieldName) + ")\n"
		r += "\twriter.WriteInt32(" + fieldName + "ArrSize)\n"
		r += "\t\tfor i := 0;i < " + fieldName + "ArrSize;i++ {\n"
		r += "\t\t\twriter.WriteBytes(this." + strings.Title(fieldName) + "[i].Serialize())\n\t\t}\n"
		return r
	}
	return "writer.WriteBytes(this." + strings.Title(fieldName) + ".Serialize())\n"
}
func GenerateGOFile(codes []core.CodeClass) string {
	sb := &strings.Builder{}
	for _, v := range codes {
		sb.WriteString("type " + v.Name + " struct{")
		sb.WriteString("\n")

		writeFuncStr := &strings.Builder{}
		writeFuncStr.WriteString("func (this *" + v.Name + ") Serialize()[]byte{\n")
		writeFuncStr.WriteString("\twriter:= &buffertool.ByteBufferWriter{}\n")
		readfuncStr := &strings.Builder{}
		readfuncStr.WriteString("func DeSerialize" + strings.Title(v.Name) + "ByByte(v []byte) " + v.Name + "{\n")
		readfuncStr.WriteString("\treturn DeSerialize" + strings.Title(v.Name) + "(&buffertool.ByteBufferReader{B:v,})")
		readfuncStr.WriteString("\n}\n")
		readfuncStr.WriteString("func DeSerialize" + strings.Title(v.Name) + "(reader *buffertool.ByteBufferReader)" + v.Name + "{\n\tr:= &" + v.Name + "{}")
		readfuncStr.WriteString("\n")
		for i, typename := range v.Types {
			sb.WriteString("\t")
			name := strings.Split(v.Names[i], "#")
			sb.WriteString(strings.Title(name[0]))

			sb.WriteString("\t")
			if strings.HasSuffix(typename, "[]") {
				sb.WriteString("[]" + strings.Replace(typename, "[]", "", 1))
			} else {
				sb.WriteString(typename)
			}

			if len(name) == 2 {
				sb.WriteString(" ")
				sb.WriteString("//" + name[1])
			}
			sb.WriteString("\n")
			readfuncStr.WriteString("\t")
			readfuncStr.WriteString(createGORead(name[0], typename))
			writeFuncStr.WriteString("\t")
			writeFuncStr.WriteString(createGOWrite(name[0], typename))
		}
		writeFuncStr.WriteString("\treturn writer.GetBytes()\n")
		writeFuncStr.WriteString("}\n")
		readfuncStr.WriteString("\treturn *r\n")
		readfuncStr.WriteString("}\n")
		sb.WriteString("}\n")
		sb.WriteString(readfuncStr.String())
		sb.WriteString(writeFuncStr.String())
	}
	return sb.String()
}
func GenerateGOEnumFile(enums []core.CodeEnum) string {
	sb := &strings.Builder{}
	for _, e := range enums {
		lastValue := 0
		typeName := strings.Title(e.Name)
		sb.WriteString("const (\n")
		for i, basename := range e.Names {
			names := strings.Split(basename, "#")
			sb.WriteString("\t" + typeName + "_" + strings.Title(names[0]) + " = ")
			tmpValue := e.Values[i]
			if tmpValue == -1 {
				lastValue += 1
				sb.WriteString(strconv.Itoa(lastValue))
			} else {
				sb.WriteString(strconv.Itoa(tmpValue))
				lastValue = tmpValue
			}
			if len(names) == 2 {
				sb.WriteString(" //" + names[1])
			}
			sb.WriteString("\n")
		}
		sb.WriteString(")\n")
	}
	return sb.String()
}
