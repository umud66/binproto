package buffertool

type ByteBufferWriter struct {
	b      []byte
	point  int
	offset byte
}

func (this *ByteBufferWriter) check(size int) {
	if this.point+size >= len(this.b) {
		this.b = append(this.b, make([]byte, 1024)...)
	}
}

func (this *ByteBufferWriter) WriteUInt8(v byte) {
	this.check(2)
	// if this.point > 0 {
	// 	if this.b[this.point-1] == v {
	// 		if this.b[this.point] < 255 {
	// 			this.b[this.point]++
	// 			return
	// 		}
	// 	}
	// }
	// if this.point > 0 {
	// 	this.point++
	// }
	this.b[this.point] = v
	this.point++

}

func (this *ByteBufferWriter) WriteInt16(v int16) {
	this.WriteUInt8(byte(v))
	this.WriteUInt8(byte(v >> 8))
}

func (this *ByteBufferWriter) WriteUInt16(v uint16) {
	this.WriteInt16(int16(v))
}

func (this *ByteBufferWriter) WriteInt32(v int) {
	this.WriteUInt8(byte(v >> 24))
	this.WriteUInt8(byte(v >> 16))
	this.WriteUInt8(byte(v >> 8))
	this.WriteUInt8(byte(v))
	// this.WriteInt16(int16(v))
	// this.WriteInt16(int16(v >> 16))
}

func (this *ByteBufferWriter) WriteUInt32(v uint) {
	this.WriteInt32(int(v))
}

func (this *ByteBufferWriter) WriteInt64(v int64) {
	// this.WriteInt32(int(v))
	// this.WriteInt32(int(v >> 32))
	this.WriteUInt8(byte(v >> 56))
	this.WriteUInt8(byte(v >> 48))
	this.WriteUInt8(byte(v >> 40))
	this.WriteUInt8(byte(v >> 32))
	this.WriteUInt8(byte(v >> 24))
	this.WriteUInt8(byte(v >> 16))
	this.WriteUInt8(byte(v >> 8))
	this.WriteUInt8(byte(v))
}

func (this *ByteBufferWriter) WriteUInt64(v uint64) {
	this.WriteInt64(int64(v))
}

func (this *ByteBufferWriter) WriteBool(v bool) {
	this.check(1)
	if v {
		this.WriteUInt8(1)
	} else {
		this.WriteUInt8(0)
	}
}
func (this *ByteBufferWriter) WriteInt8Arr(v []int8) {
	if v == nil {
		this.WriteInt32(0)
	} else {
		len := len(v)
		this.WriteInt32(len)
		for _, _v := range v {
			this.WriteUInt8(byte(_v))
		}
	}
}
func (this *ByteBufferWriter) WriteUInt8Arr(v []byte) {
	len := len(v)
	this.WriteInt32(len)
	this.WriteBytes(v)
}

func (this *ByteBufferWriter) WriteInt16Arr(v []int16) {
	if v == nil {
		this.WriteInt32(0)
	} else {
		len := len(v)
		this.WriteInt32(len)
		for _, _v := range v {
			this.WriteInt16(_v)
		}
	}
}

func (this *ByteBufferWriter) WriteUInt16Arr(v []uint16) {
	if v == nil {
		this.WriteInt32(0)
	} else {
		len := len(v)
		this.WriteInt32(len)
		for _, _v := range v {
			this.WriteUInt16(_v)
		}
	}
}

func (this *ByteBufferWriter) WriteInt32Arr(v []int) {
	if v == nil {
		this.WriteInt32(0)
	} else {
		len := len(v)
		this.WriteInt32(len)
		for _, _v := range v {
			this.WriteInt32(_v)
		}
	}
}

func (this *ByteBufferWriter) WriteUInt32Arr(v []uint) {
	if v == nil {
		this.WriteInt32(0)
	} else {
		len := len(v)
		this.WriteInt32(len)
		for _, _v := range v {
			this.WriteUInt32(_v)
		}
	}
}

func (this *ByteBufferWriter) WriteInt64Arr(v []int64) {
	if v == nil {
		this.WriteInt32(0)
	} else {
		len := len(v)
		this.WriteInt32(len)
		for _, _v := range v {
			this.WriteInt64(_v)
		}
	}
}

func (this *ByteBufferWriter) WriteUInt64Arr(v []uint64) {
	if v == nil {
		this.WriteInt32(0)
	} else {
		len := len(v)
		this.WriteInt32(len)
		for _, _v := range v {
			this.WriteUInt64(_v)
		}
	}
}

func (this *ByteBufferWriter) WriteBytes(v []byte) {
	len := len(v)
	this.check(len)
	for _, v := range v {
		this.WriteUInt8(v)
	}
}

func (this *ByteBufferWriter) WriteString(v string) {
	b := []byte(v)
	this.WriteUInt8Arr(b)

}

func (this *ByteBufferWriter) WriteStringArr(v []string) {
	if v == nil {
		this.WriteInt32(0)
	} else {
		len := len(v)
		this.WriteInt32(len)
		for _, v := range v {
			this.WriteString(v)
		}
	}
}

func (this *ByteBufferWriter) WriteBoolArr(v []bool) {
	if v == nil {
		this.WriteInt32(0)
	} else {
		len := len(v)
		this.WriteInt32(len)
		for _, v := range v {
			this.WriteBool(v)
		}
	}
}

func (this *ByteBufferWriter) GetBytes() []byte {
	return this.b[:this.point]
}
