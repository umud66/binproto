package buffertool

type ByteBufferReader struct {
	B      []byte
	point  int
	offset byte
}

func (this *ByteBufferReader) ReadInt8() int8 {
	return int8(this.ReadUInt8())
}

func (this *ByteBufferReader) ReadUInt8() byte {
	// if this.offset > 0 {
	// 	b := this.B[this.point-1]
	// 	this.offset--
	// 	return b
	// }
	// if this.point > 0 {
	// 	this.point++
	// }
	b := this.B[this.point]
	this.point++
	// this.offset = this.B[this.point]
	return b
}

func (this *ByteBufferReader) ReadUInt16() uint16 {
	return uint16(this.ReadInt16())
}

func (this *ByteBufferReader) ReadInt16() int16 {
	return int16(this.ReadUInt8()) | int16(this.ReadUInt8())<<8
}

func (this *ByteBufferReader) ReadInt32() int {
	// return int(int32(this.ReadInt16()) | int32(this.ReadInt16())<<16)
	return int(int32(this.ReadUInt8())<<24 | int32(this.ReadUInt8())<<16 | int32(this.ReadUInt8())<<8 | int32(this.ReadUInt8()))
}

func (this *ByteBufferReader) ReadUInt32() uint {
	return uint(this.ReadInt32())
}

func (this *ByteBufferReader) ReadBool() bool {
	return this.ReadInt8() == 1
}

func (this *ByteBufferReader) ReadInt64() int64 {
	// return int64(this.ReadInt32()) | int64(this.ReadInt32())<<32
	return int64(this.ReadUInt8())<<56 | int64(this.ReadUInt8())<<48 | int64(this.ReadUInt8())<<40 | int64(this.ReadUInt8())<<32 | int64(this.ReadUInt8())<<24 | int64(this.ReadUInt8())<<16 | int64(this.ReadUInt8())<<8 | int64(this.ReadUInt8())
}

func (this *ByteBufferReader) ReadUInt64() uint64 {
	return uint64(this.ReadInt64())
}

func (this *ByteBufferReader) ReadInt8Arr() []int8 {
	len := this.ReadInt32()
	r := make([]int8, len)
	for i := 0; i < len; i++ {
		r[i] = this.ReadInt8()
	}
	return r
}

func (this *ByteBufferReader) ReadUInt8Arr() []byte {
	len := this.ReadInt32()
	r := make([]uint8, len)
	for i := 0; i < len; i++ {
		r[i] = this.ReadUInt8()
	}
	return r
}

func (this *ByteBufferReader) ReadInt16Arr() []int16 {
	len := this.ReadInt32()
	r := make([]int16, len)
	for i := 0; i < len; i++ {
		r[i] = this.ReadInt16()
	}
	return r
}

func (this *ByteBufferReader) ReadUInt16Arr() []uint16 {
	len := this.ReadInt32()
	r := make([]uint16, len)
	for i := 0; i < len; i++ {
		r[i] = this.ReadUInt16()
	}
	return r
}

func (this *ByteBufferReader) ReadInt32Arr() []int {
	len := this.ReadInt32()
	r := make([]int, len)
	for i := 0; i < len; i++ {
		r[i] = this.ReadInt32()
	}
	return r
}

func (this *ByteBufferReader) ReadUInt32Arr() []uint {
	len := this.ReadInt32()
	r := make([]uint, len)
	for i := 0; i < len; i++ {
		r[i] = this.ReadUInt32()
	}
	return r
}

func (this *ByteBufferReader) ReadInt64Arr() []int64 {
	len := this.ReadInt32()
	r := make([]int64, len)
	for i := 0; i < len; i++ {
		r[i] = this.ReadInt64()
	}
	return r
}

func (this *ByteBufferReader) ReadUInt64Arr() []uint64 {
	len := this.ReadInt32()
	r := make([]uint64, len)
	for i := 0; i < len; i++ {
		r[i] = this.ReadUInt64()
	}
	return r
}
func (this *ByteBufferReader) ReadUInt64ArrArr() [][]uint64 {
	len := this.ReadInt32()
	r := make([][]uint64, len)
	for i := 0; i < len; i++ {
		r[i] = this.ReadUInt64Arr()
	}
	return r
}
func (this *ByteBufferReader) ReadInt32ArrArr() [][]int {
	len := this.ReadInt32()
	r := make([][]int, len)
	for i := 0; i < len; i++ {
		r[i] = this.ReadInt32Arr()
	}
	return r
}
func (this *ByteBufferReader) ReadUInt32ArrArr() [][]uint {
	len := this.ReadInt32()
	r := make([][]uint, len)
	for i := 0; i < len; i++ {
		r[i] = this.ReadUInt32Arr()
	}
	return r
}
func (this *ByteBufferReader) ReadUInt8ArrArr() [][]byte {
	len := this.ReadInt32()
	r := make([][]byte, len)
	for i := 0; i < len; i++ {
		r[i] = this.ReadUInt8Arr()
	}
	return r
}
func (this *ByteBufferReader) ReadStringArrArr() [][]string {
	len := this.ReadInt32()
	r := make([][]string, len)
	for i := 0; i < len; i++ {
		r[i] = this.ReadStringArr()
	}
	return r
}
func (this *ByteBufferReader) ReadBoolArrArr() [][]bool {
	len := this.ReadInt32()
	r := make([][]bool, len)
	for i := 0; i < len; i++ {
		r[i] = this.ReadBoolArr()
	}
	return r
}

func (this *ByteBufferReader) ReadBoolArr() []bool {
	len := this.ReadInt32()
	r := make([]bool, len)
	for i := 0; i < len; i++ {
		r[i] = this.ReadBool()
	}
	return r
}

func (this *ByteBufferReader) ReadString() string {
	arr := this.ReadUInt8Arr()
	return string(arr)
}

func (this *ByteBufferReader) ReadStringArr() []string {
	len := this.ReadInt32()
	r := make([]string, len)
	for i := 0; i < len; i++ {
		r[i] = this.ReadString()
	}
	return r
}
