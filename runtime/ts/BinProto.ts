const UINT64MAX = Math.pow(2, 64) - 1;
const UINT32MAX = Math.pow(2, 32) - 1;
const UINT16MAX = Math.pow(2, 16) - 1;
const UINT8MAX = Math.pow(2, 8) - 1;
export class BinProtoReader {
    constructor(data: Uint8Array) {
        this.data = data.entries()
        this.readPoint = 0;

    }
    data: IterableIterator<[number, number]>;
    readPoint: number;
    ReadByte(): number {
        return this.data.next().value[1];
    }
    ReadInt16(): number {
        return this.ReadByte() | (this.ReadByte() << 8);
    }
    ReadUInt16(): number {
        return this.ReadInt16()
    }
    ReadInt32(): number {
        // return this.ReadInt16() | (this.ReadInt16() << 16);
        return this.ReadByte() << 24 | this.ReadByte() << 16 | this.ReadByte() << 8 | this.ReadByte();
    }
    ReadUInt32(): number {
        return this.ReadInt32() >>> 0;
    }
    ReadInt64(): number {
        let value = BigInt(0);
        for (let i = 0; i < 8; i++) {
            value = (value << BigInt(8)) | BigInt(this.ReadByte());
        }
        return Number(value);
    }
    ReadUInt64(): number {
        return this.ReadInt64() >>> 0;
    }
    ReadUInt64ArrayArr(): number[][] {
        let arr: number[][] = [];
        let size = this.ReadUInt32()
        for (let i = 0; i < size; i++) {
            arr.push(this.ReadInt64Array())
        }
        return arr
    }
    ReadInt64ArrayArr(): number[][] {
        let arr: number[][] = [];
        let size = this.ReadUInt32()
        for (let i = 0; i < size; i++) {
            arr.push(this.ReadInt64Array())
        }
        return arr
    }
    ReadInt32ArrayArr(): number[][] {
        let arr: number[][] = [];
        let size = this.ReadUInt32()
        for (let i = 0; i < size; i++) {
            arr.push(this.ReadInt32Array())
        }
        return arr
    }
    ReadUInt32ArrayArr(): number[][] {
        let arr: number[][] = [];
        let size = this.ReadUInt32()
        for (let i = 0; i < size; i++) {
            arr.push(this.ReadUInt32Array())
        }
        return arr
    }
    ReadInt16ArrayArr(): number[][] {
        let arr: number[][] = [];
        let size = this.ReadUInt32()
        for (let i = 0; i < size; i++) {
            arr.push(this.ReadInt16Array())
        }
        return arr
    }
    ReadUInt16ArrayArr(): number[][] {
        let arr: number[][] = [];
        let size = this.ReadUInt32()
        for (let i = 0; i < size; i++) {
            arr.push(this.ReadUInt16Array())
        }
        return arr
    }
    ReadUInt8ArrayArr(): number[][] {
        let arr: number[][] = [];
        let size = this.ReadUInt32()
        for (let i = 0; i < size; i++) {
            arr.push(this.ReadByteArray())
        }
        return arr
    }
    ReadStringArrayArr():string[][]{
        let arr: string[][] = [];
        let size = this.ReadUInt32()
        for (let i = 0; i < size; i++) {
            arr.push(this.ReadStringArray())
        }
        return arr
    }
    ReadString(): string {
        let arr = this.ReadByteArray();
        return new TextDecoder().decode(new Uint8Array(arr));
    }
    ReadBool(): boolean {
        return this.ReadByte() == 1;
    }
    ReadByteArray(): number[] {
        let size = this.ReadUInt32()
        let r: number[] = [];
        for (let i = 0; i < size; i++) {
            r.push(this.ReadByte())
        }
        return r;
    }
    ReadInt16Array(): number[] {
        let size = this.ReadUInt32()
        let r: number[] = [];
        for (let i = 0; i < size; i++) {
            r.push(this.ReadInt16())
        }
        return r;
    }
    ReadUInt16Array(): number[] {
        let size = this.ReadUInt32()
        let r: number[] = [];
        for (let i = 0; i < size; i++) {
            r.push(this.ReadUInt16())
        }
        return r;
    }
    ReadInt32Array(): number[] {
        let size = this.ReadUInt32()
        let r: number[] = [];
        for (let i = 0; i < size; i++) {
            r.push(this.ReadInt32())
        }
        return r;
    }
    ReadUInt32Array(): number[] {
        let size = this.ReadUInt32()
        let r: number[] = [];
        for (let i = 0; i < size; i++) {
            r.push(this.ReadUInt32())
        }
        return r;
    }
    ReadInt64Array(): number[] {
        let size = this.ReadUInt32()
        let r: number[] = [];
        for (let i = 0; i < size; i++) {
            r.push(this.ReadInt64())
        }
        return r;
    }
    ReadUInt64Array(): number[] {
        let size = this.ReadUInt32()
        let r: number[] = [];
        for (let i = 0; i < size; i++) {
            r.push(this.ReadUInt64())
        }
        return r;
    }
    ReadStringArray(): string[] {
        let size = this.ReadUInt32()
        let r: string[] = [];
        for (let i = 0; i < size; i++) {
            r.push(this.ReadString())
        }
        return r;
    }
}
export class BinProtoWriter {
    data: number[];
    writePoint: number;
    constructor() {
        this.data = [];
        this.writePoint = 0;
    }

    WriteByte(val: number) {
        // if(val >= UINT8MAX)
        //     val = UINT8MAX
        this.data.push(val & 0xFF)
    }

    WriteUInt16(val: number) {
        // if(val >= UINT16MAX)
        //     val = UINT16MAX;
        val &= 0xFFFF
        this.WriteByte(val >> 8);
        this.WriteByte(val);
        // this.WriteByte(val)
        // this.WriteByte(val >>> 8);
    }

    WriteInt32(val: number) {
        val &= 0xFFFFFFFF
        // this.WriteUInt16(val);
        // this.WriteUInt16(val >>> 16);
        this.WriteByte(val >> 24);
        this.WriteByte(val >> 16);
        this.WriteByte(val >> 8);
        this.WriteByte(val);
    }

    WriteUInt32(val: number) {
        this.WriteInt32(val)
    }

    WriteInt64(val: number) {
        let v = (BigInt)(val)
        for (let i = 7; i >= 0; i--) {
            this.WriteByte(Number((v >> BigInt(i * 8)) & BigInt(0xFF)));
        }
        // this.WriteUInt32(Number(v))
        // this.WriteUInt32(Number(v >> 32n))
        // if(val > UINT32MAX){
        //     this.WriteUInt32(UINT32MAX);
        //     this.WriteUInt32(val / UINT32MAX);
        // }else{
        //     this.WriteUInt32(val);
        //     this.WriteUInt32(0)
        // }
    }

    WriteUInt64(val: number) {
        this.WriteInt64(val)
    }
    WriteString(val: string) {
        if (!val) {
            this.WriteByteArray([]);
            return
        }
        let arr = new TextEncoder().encode(val);
        this.WriteArrayLength(arr.length);
        this.WriteUint8ArrayBytes(arr)
    }
    WriteBool(val: boolean) {
        this.WriteByte(val ? 1 : 0);
    }
    WriteByteArray(val: number[]) {
        this.WriteArrayLength(val.length)
        for (let i = 0; i < val.length; i++) {
            this.WriteByte(val[i]);
        }
    }
    WriteInt16Array(val: number[]) {
        this.WriteArrayLength(val.length)
        val.forEach(it => {
            this.WriteUInt16(it);
        })
    }
    WriteInt32Array(val: number[]) {
        this.WriteArrayLength(val.length)
        val.forEach(it => {
            this.WriteInt32(it);
        })
    }
    WriteUInt32Array(val: number[]) {
        this.WriteArrayLength(val.length)
        val.forEach(it => {
            this.WriteUInt32(it);
        })
    }
    WriteInt64Array(val: number[]) {
        this.WriteArrayLength(val.length)
        val.forEach(it => {
            this.WriteInt64(it);
        })
    }
    WriteUInt64Array(val: number[]) {
        this.WriteArrayLength(val.length)
        val.forEach(it => {
            this.WriteUInt64(it);
        })
    }
    WriteArrayLength(val:number){
        this.WriteUInt32(val)
    }
    WriteStringArray(val: string[]) {
        this.WriteArrayLength(val.values.length)
        val.forEach(it => {
            this.WriteString(it);
        })
    }
    WriteBytes(val: number[]) {
        val.forEach(it => {
            this.WriteByte(it)
        });
    }
    WriteUint8ArrayBytes(val: Uint8Array) {
        for (let i = 0; i < val.byteLength; i++) {
            this.WriteByte(val[i]);
        }
    }
    WriteUInt64ArrayArr(val: number[][]) {
        this.WriteArrayLength(val.length)
        val.forEach(it => {
            this.WriteInt64Array(it);
        })
    }
    WriteInt64ArrayArr(val: number[][]) {
        this.WriteArrayLength(val.length)
        val.forEach(it => {
            this.WriteInt64Array(it);
        })
    }
    WriteInt32ArrayArr(val: number[][]){
        this.WriteArrayLength(val.length)
        val.forEach(it => {
            this.WriteInt32Array(it);
        })
    }
    WriteUInt32ArrayArr(val: number[][]){
        this.WriteArrayLength(val.length)
        val.forEach(it => {
            this.WriteUInt32Array(it);
        })
    }
    GetBytes(): Uint8Array {
        return new Uint8Array(this.data)
    }
    GetNumberArray(): number[] {
        return this.data
    }
}
// let writer = new BinProtoWriter();
// console.log(229921394282)
// writer.WriteInt64(229921394282)
// writer.WriteUInt16(65533)
// writer.WriteString("123")
// console.log(writer.GetBytes())

// let reader = new BinProtoReader(writer.GetBytes());
// console.log(reader)
// console.log(reader.ReadInt64())
// console.log(reader.ReadUInt16())
// console.log(reader.ReadUInt16())
// console.log(reader.ReadString())

// class Test{
//     buffer:number[];
//     constructor(){
//         this.buffer = [];
//     }
//     WriteByte(val:number) {

//         if(val > 255)
//             val = 255;
//         this.buffer.push(val)    
//     }
// }

// function testInt(val:number){
//     let buffer:any = [];
//     buffer[0] = val & 255;
//     buffer[1] = val>>=8 & 255;
//     buffer[2] = val>>=8 & 255;
//     buffer[3] = val>>=8 & 255;
//     console.log(buffer)
//     let t = buffer[0]
//     t |= buffer[1] << 8
//     t |= buffer[2] << 16
//     t |= buffer[3] << 24
//     console.log(t)
// }
// // testInt(65536)
// console.log(UINT8MAX)
// console.log(256 & 256)