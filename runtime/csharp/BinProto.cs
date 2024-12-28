using System;

public partial class BinProto
{
    public class BinReader
    {
        private byte[] b;
        private int point;
        private byte offset;
        public BinReader(byte[] b)
        {
            this.b = b;
        }

        public bool CanRead
        {
            get => point < b.Length;
        }

        public byte ReadByte()
        {
            if(offset > 0){
                offset --;
                return b[point -1];
            }
            if(point > 0){
                point++;
            }
            byte r = b[point++];
            offset = b[point];
            return r;
        }

        public sbyte ReadSByte()
        {
            return (sbyte)ReadByte();
        }

        public short ReadInt16()
        {
            return (short)(ReadByte() | ReadByte() << 8);
        }
        
        public ushort ReadUInt16()
        {
            return (ushort)ReadInt16();
        }
        
        public int ReadInt32()
        {
            return (int)(ReadByte() << 24 | ReadByte() << 16 | ReadByte() << 8 | ReadByte());
        }
        public uint ReadUInt32()
        {
            return (uint)ReadInt32();
        }
        public long ReadInt64()
        {
            return (long)(ReadByte() << 56 | ReadByte() << 48 | ReadByte() << 40 | ReadByte() << 32 | ReadByte() << 24 | ReadByte() << 16 | ReadByte() << 8 | ReadByte());
        }
        
        public ulong ReadUInt64()
        {
            return (ulong)ReadInt64();
        }
        
        public bool ReadBool()
        {
            return ReadByte() == 1;
        }
        
        public string ReadString()
        {
            return System.Text.UTF8Encoding.UTF8.GetString(ReadByteArray());
        }

        public byte[] ReadByteArray()
        {
            byte[] b = new byte[ReadInt32()];
            for (int i = 0; i < b.Length; i++)
            {
                b[i] = ReadByte();
            }

            return b;
        }
        
        public sbyte[] ReadSByteArray()
        {
            sbyte[] b = new sbyte[ReadInt32()];
            for (int i = 0; i < b.Length; i++)
            {
                b[i] = ReadSByte();
            }

            return b;
        }
        public short[] ReadInt16Array()
        {
            short[] b = new short[ReadInt32()];
            for (int i = 0; i < b.Length; i++)
            {
                b[i] = ReadInt16();
            }

            return b;
        }
        public ushort[] ReadUInt16Array()
        {
            ushort[] b = new ushort[ReadInt32()];
            for (int i = 0; i < b.Length; i++)
            {
                b[i] = ReadUInt16();
            }

            return b;
        }
        
        public int[] ReadInt32Array()
        {
            int[] b = new int[ReadInt32()];
            for (int i = 0; i < b.Length; i++)
            {
                b[i] = ReadInt32();
            }

            return b;
        }
        public uint[] ReadUInt32Array()
        {
            uint[] b = new uint[ReadInt32()];
            for (int i = 0; i < b.Length; i++)
            {
                b[i] = ReadUInt32();
            }

            return b;
        }
        
        public long[] ReadInt64Array()
        {
            long[] b = new long[ReadInt32()];
            for (int i = 0; i < b.Length; i++)
            {
                b[i] = ReadInt64();
            }

            return b;
        }
        public ulong[] ReadUInt64Array()
        {
            ulong[] b = new ulong[ReadInt32()];
            for (int i = 0; i < b.Length; i++)
            {
                b[i] = ReadUInt64();
            }

            return b;
        }
        
        public bool[] ReadBoolArray()
        {
            bool[] b = new bool[ReadInt32()];
            for (int i = 0; i < b.Length; i++)
            {
                b[i] = ReadBool();
            }

            return b;
        }
        
        public string[] ReadStringArray()
        {
            string[] b = new string[ReadInt32()];
            for (int i = 0; i < b.Length; i++)
            {
                b[i] = ReadString();
            }

            return b;
        }
    }
    
    public class BinWriter
    {
        private byte[] b;
        private int point;
        private byte offset;
        public byte[] GetBytes()
        {
            Array.Resize(ref b,point + 1);
            return b;
        }
        
        private void check(int size)
        {
            if (b == null)
            {
                b = new byte[1024];
            }
            if (point + size >= b.Length)
            {
                Array.Resize(ref b,b.Length + 1024);
            }
        }
        
        public void WriteBool(bool v)
        {
            WriteByte((byte)(v ? 1 : 0));
        }
        
        public void WriteString(string v)
        {
            WriteUInt8Array(System.Text.UTF8Encoding.UTF8.GetBytes(v));
        }
        
        
        public void WriteByte(byte v)
        {
            check(2);
            if(this.point > 0){
                if(b[point - 1] == v){
                    if(b[point] < 255){
                        b[point]++;
                        return;
                    }
                }
            }
            if(point > 0){
                point++;
            }
            b[point++] = v;
        }
        public void WriteSByte(sbyte v)
        {
            check(1);
            WriteByte((byte)v);
        }
        public void WriteInt16(short v)
        {
            WriteByte((byte)v);
            WriteByte((byte)(v >> 8));
        }
        public void WriteUInt16(ushort v)
        {
            WriteInt16((short)v);
        }
        public void WriteInt32(int v)
        {
            WriteByte((byte)v >> 24);
            WriteByte((byte)v >> 16);
            WriteByte((byte)v >> 8);
            WriteByte((byte)v);
            // WriteInt16((short)v);
            // WriteInt16((short)(v >> 16));
        }
        
        public void WriteUInt32(uint v)
        {
            WriteInt32((int)v);
        }
        
        public void WriteInt64(long v)
        {
            // WriteInt32((int)v);
            // WriteInt32((int) (v >> 32));
            WriteByte((byte)v >> 56);
            WriteByte((byte)v >> 48);
            WriteByte((byte)v >> 40);
            WriteByte((byte)v >> 32);
            WriteByte((byte)v >> 24);
            WriteByte((byte)v >> 16);
            WriteByte((byte)v >> 8);
            WriteByte((byte)
        }
        
        public void WriteUInt64(ulong v)
        {
            WriteInt64((long)v);
        }
        
        public void WriteInt8Array(sbyte[] v)
        {
            WriteInt32(v.Length);
            foreach (var b in v)
            {
                WriteSByte(b);
            }
        }
        public void WriteUInt8Array(byte[] v)
        {
            WriteInt32(v.Length);
            WriteBytes(v);
        }
        public void WriteInt16Array(short[] v)
        {
            WriteInt32(v.Length);
            foreach (var b in v)
            {
                WriteInt16(b);
            }
        }
        
        public void WriteUInt16Array(ushort[] v)
        {
            WriteInt32(v.Length);
            foreach (var b in v)
            {
                WriteUInt16(b);
            }
        }
        
        public void WriteInt32Array(int[] v)
        {
            WriteInt32(v.Length);
            foreach (var b in v)
            {
                WriteInt32(b);
            }
        }
        
        public void WriteUInt32Array(uint[] v)
        {
            WriteInt32(v.Length);
            foreach (var b in v)
            {
                WriteUInt32(b);
            }
        }
        
        public void WriteInt64Array(long[] v)
        {
            WriteInt32(v.Length);
            foreach (var b in v)
            {
                WriteInt64(b);
            }
        }
        
        public void WriteUInt64Array(ulong[] v)
        {
            WriteInt32(v.Length);
            foreach (var b in v)
            {
                WriteUInt64(b);
            }
        }
        public void WriteStringArray(string[] v)
        {
            WriteInt32(v.Length);
            foreach (var b in v)
            {
                WriteString(b);
            }
        }
        
        public void WriteBoolArray(bool[] v)
        {
            WriteInt32(v.Length);
            foreach (var b in v)
            {
                WriteBool(b);
            }
        }

        public void WriteBytes(byte[] bytes)
        {
            foreach (var b in bytes)
            {
                WriteByte(b);
            }
        }
    }
}