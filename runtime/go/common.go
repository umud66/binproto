package buffertool

type BinBase interface {
	Serialize() []byte
	DeSerializeByByte(v []byte)
	DeSerialize(reader *ByteBufferReader)
}
