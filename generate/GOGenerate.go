package generate

import (
	"strconv"
	"strings"

	"umud.online/bin/core"
)

func isBaseType(t string) bool {
	if strings.HasPrefix(t, "int") {
		return true
	}
	if strings.HasPrefix(t, "uint") {
		return true
	}
	if strings.HasPrefix(t, "byte") {
		return true
	}
	if strings.HasPrefix(t, "bool") {
		return true
	}
	if strings.HasPrefix(t, "string") {
		return true
	}
	return false
}

func createGORead(fieldName string, fieldType string) string {
	r := "data." + strings.Title(fieldName) + " = "
	if fieldType == "int" {
		return r + "reader.ReadInt32()\n"
	} else if fieldType == "uint" {
		return r + "reader.ReadUInt32()\n"
	} else if fieldType == "byte" || fieldType == "uin8" {
		return r + "reader.ReadUInt8()\n"
	} else if fieldType == "int8" {
		return r + "reader.ReadInt8()\n"
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
	} else if fieldType == "int8[]" {
		return r + "reader.ReadInt8Arr()\n"
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
		r = fieldName + "ArrSize := reader.ReadInt32()\n\t\t" + r + "make([]" + strings.Title(baseType) + "," + fieldName + "ArrSize);\n"
		r += "\t\tfor i := 0;i < " + fieldName + "ArrSize;i++ {\n"
		r += "\t\t\t_tmp:= " + strings.Title(baseType) + "{}\n"
		r += "\t\t\tdata." + strings.Title(fieldName) + "[i] = _tmp\n"
		r += "\t\t\t_tmp.DeSerialize(reader)\n\t\t}\n"
		return r
	}

	return fieldName + " := " + strings.Title(fieldType) + "{}\n" + fieldName + ".DeSerialize(reader)\n" + r + fieldName + "\n"
}
func createGOWrite(fieldName string, fieldType string) string {
	if fieldType == "int" {
		return "writer.WriteInt32(data." + strings.Title(fieldName) + ")\n"
	} else if fieldType == "uint" {
		return "writer.WriteUInt32(data." + strings.Title(fieldName) + ")\n"
	} else if fieldType == "byte" || fieldType == "uint8" {
		return "writer.WriteUInt8(data." + strings.Title(fieldName) + ")\n"
	} else if fieldType == "int8" {
		return "writer.WriteInt8(data." + strings.Title(fieldName) + ")\n"
	} else if fieldType == "int64" {
		return "writer.WriteInt64(data." + strings.Title(fieldName) + ")\n"
	} else if fieldType == "uint64" {
		return "writer.WriteUInt64(data." + strings.Title(fieldName) + ")\n"
	} else if fieldType == "bool" {
		return "writer.WriteBool(data." + strings.Title(fieldName) + ")\n"
	} else if fieldType == "int16" {
		return "writer.WriteInt16(data." + strings.Title(fieldName) + ")\n"
	} else if fieldType == "uint16" {
		return "writer.WriteUInt16(data." + strings.Title(fieldName) + ")\n"
	} else if fieldType == "string" {
		return "writer.WriteString(data." + strings.Title(fieldName) + ")\n"
	} else if fieldType == "int[]" {
		return "writer.WriteInt32Arr(data." + strings.Title(fieldName) + ")\n"
	} else if fieldType == "uint[]" {
		return "writer.WriteUInt32Arr(data." + strings.Title(fieldName) + ")\n"
	} else if fieldType == "byte[]" || fieldType == "uin8[]" {
		return "writer.WriteUInt8Arr(data." + strings.Title(fieldName) + ")\n"
	} else if fieldType == "int8[]" {
		return "writer.WriteInt8Arr(data." + strings.Title(fieldName) + ")\n"
	} else if fieldType == "int16[]" {
		return "writer.WriteInt16Arr(data." + strings.Title(fieldName) + ")\n"
	} else if fieldType == "uint16[]" {
		return "writer.WriteUInt16Arr(data." + strings.Title(fieldName) + ")\n"
	} else if fieldType == "int64[]" {
		return "writer.WriteInt64Arr(data." + strings.Title(fieldName) + ")\n"
	} else if fieldType == "uint64[]" {
		return "writer.WriteUInt64Arr(data." + strings.Title(fieldName) + ")\n"
	} else if fieldType == "string[]" {
		return "writer.WriteStringArr(data." + strings.Title(fieldName) + ")\n"
	} else if fieldType == "bool[]" {
		return "writer.WriteBoolArr(data." + strings.Title(fieldName) + ")\n"
	} else if strings.HasSuffix(fieldType, "[]") {
		r := fieldName + "ArrSize := len(data." + strings.Title(fieldName) + ")\n"
		r += "\twriter.WriteInt32(" + fieldName + "ArrSize)\n"
		r += "\t\tfor i := 0;i < " + fieldName + "ArrSize;i++ {\n"
		r += "\t\t\twriter.WriteBytes(data." + strings.Title(fieldName) + "[i].Serialize())\n\t\t}\n"
		return r
	}
	return "writer.WriteBytes(data." + strings.Title(fieldName) + ".Serialize())\n"
}
func GenerateGOFile(codes []core.CodeClass, godb bool) string {
	sb := &strings.Builder{}
	if !godb {
		// sb.WriteString("type BinBase interface {\n\tSerialize() []byte\n\tDeSerializeByByte(v []byte)\n\tDeSerialize(reader *ByteBufferReader)\n}\n")
	}
	for _, v := range codes {
		baseName := v.Name
		v.Name = strings.Title(v.Name)
		sb.WriteString("type " + v.Name + " struct{")
		sb.WriteString("\n")
		tableFuncStr := &strings.Builder{}
		writeFuncStr := &strings.Builder{}
		writeFuncStr.WriteString("func (data *" + v.Name + ") Serialize()[]byte{\n")
		writeFuncStr.WriteString("\twriter:= &ByteBufferWriter{}\n")
		readfuncStr := &strings.Builder{}
		readfuncStr.WriteString("func (data *" + v.Name + ") DeSerializeByByte(v []byte) {\n")
		readfuncStr.WriteString("\t data.DeSerialize(&ByteBufferReader{B:v,})")
		readfuncStr.WriteString("\n}\n")
		readfuncStr.WriteString("func (data *" + v.Name + ") DeSerialize(reader *ByteBufferReader){\n")
		for i, typename := range v.Types {
			sb.WriteString("\t")
			name := strings.Split(v.Names[i], "#")
			sb.WriteString(strings.Title(name[0]))
			sb.WriteString("\t")
			if strings.HasSuffix(typename, "[]") {
				if isBaseType(typename) {
					sb.WriteString("[]" + strings.Replace(typename, "[]", "", 1))
				} else {
					sb.WriteString("[]" + strings.Replace(strings.Title(typename), "[]", "", 1))
				}
			} else {
				if isBaseType(typename) {
					sb.WriteString(typename)
				} else {
					sb.WriteString(strings.Title(typename))
				}
			}
			if godb {
				sb.WriteString("\t" + v.Tags[i])
				sb.WriteString("\t")
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
		readfuncStr.WriteString("}\n")
		sb.WriteString("}\n")
		if !godb {
			sb.WriteString(readfuncStr.String())
			sb.WriteString(writeFuncStr.String())
		} else {
			tableFuncStr.WriteString("func (data *" + v.Name + ") TableName() string {\n")
			tableFuncStr.WriteString("\treturn \"" + baseName + "\"\n}\n")
			sb.WriteString(tableFuncStr.String())
		}
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

func GenerateGOConstFile(enums []core.CodeConst) string {
	sb := &strings.Builder{}
	for _, e := range enums {
		typeName := strings.Title(e.Name)
		sb.WriteString("const (\n")
		for i, basename := range e.Names {
			names := strings.Split(basename, "#")
			sb.WriteString("\t" + typeName + "_" + strings.Title(names[0]) + " = ")
			tmpValue := e.Values[i]
			if strings.Index(tmpValue, "=") > 0 {
				tmpValue = strings.Split(tmpValue, "=")[1]
			}
			if e.ValueType == "string" {
				sb.WriteString("\"" + tmpValue + "\"")
			} else {
				sb.WriteString(tmpValue)
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
