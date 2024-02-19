const UINT64MAX = Math.pow(2,64) -1;
const UINT32MAX = Math.pow(2,32) -1;
const UINT16MAX = Math.pow(2,16) - 1;
const UINT8MAX = Math.pow(2,8) - 1;
class BinProtoReader{
    constructor(data:Uint8Array){
        this.data = data.entries()
        this.readPoint = 0;
        
    }
    data:IterableIterator<[number, number]>;
    readPoint:number;
    ReadByte():number {
        return this.data.next().value[1];
    }
    ReadInt16():number{
        return this.ReadByte() | (this.ReadByte() << 8);
    }
    ReadUInt16():number{
        return this.ReadInt16()
    }
    ReadInt32():number{
        return this.ReadInt16() | (this.ReadInt16() << 16);
    }
    ReadUInt32():number{
        return this.ReadInt32() >>> 0;
    }
    ReadInt64():number{
        let big = (BigInt)(this.ReadUInt32()) | (BigInt) (this.ReadUInt32()) << 32n
        return Number(big);
    }
    ReadUInt64():number{
        return this.ReadInt64() >>> 0;
    }
    ReadString():string{
        let arr = this.ReadByteArray();
        return new TextDecoder().decode(new Uint8Array(arr));
    }
    ReadBool():boolean{
        return this.ReadByte() == 1;
    }
    ReadByteArray():number[]{
        let size = this.ReadUInt32()
        let r:number[] = [];
        for (let i = 0; i < size; i++) {
            r.push(this.ReadByte())
        }
        return r;
    }
    ReadInt16Array():number[]{
        let size = this.ReadUInt32()
        let r:number[] = [];
        for (let i = 0; i < size; i++) {
            r.push(this.ReadInt16())
        }
        return r;
    }
    ReadUInt16Array():number[]{
        let size = this.ReadUInt32()
        let r:number[] = [];
        for (let i = 0; i < size; i++) {
            r.push(this.ReadUInt16())
        }
        return r;
    }
    ReadInt32Array():number[]{
        let size = this.ReadUInt32()
        let r:number[] = [];
        for (let i = 0; i < size; i++) {
            r.push(this.ReadInt32())
        }
        return r;
    }
    ReadUInt32Array():number[]{
        let size = this.ReadUInt32()
        let r:number[] = [];
        for (let i = 0; i < size; i++) {
            r.push(this.ReadUInt32())
        }
        return r;
    }
    ReadInt64Array():number[]{
        let size = this.ReadUInt32()
        let r:number[] = [];
        for (let i = 0; i < size; i++) {
            r.push(this.ReadInt64())
        }
        return r;
    }
    ReadUInt64Array():number[]{
        let size = this.ReadUInt32()
        let r:number[] = [];
        for (let i = 0; i < size; i++) {
            r.push(this.ReadUInt64())
        }
        return r;
    }
    ReadStringArray():string[]{
        let size = this.ReadUInt32()
        let r:string[] = [];
        for (let i = 0; i < size; i++) {
            r.push(this.ReadString())
        }
        return r;
    }
}
class BinProtoWriter{
    data:number[];
    writePoint:number;
    constructor(){
        this.data = [];
        this.writePoint = 0;
    }

    WriteByte(val:number){
        // if(val >= UINT8MAX)
        //     val = UINT8MAX
        this.data.push(val & 0xFF)
    }

    WriteUInt16(val:number){
        // if(val >= UINT16MAX)
        //     val = UINT16MAX;
        val &= 0xFFFF
        this.WriteByte(val)
        this.WriteByte(val >>> 8);
    }

    WriteInt32(val:number){
        val &= 0xFFFFFFFF
        this.WriteUInt16(val);
        this.WriteUInt16(val >>> 16);
    }

    WriteUInt32(val:number){
        this.WriteInt32(val)
    }

    WriteInt64(val:number){
        let v = (BigInt)(val)
        this.WriteUInt32(Number(v))
        this.WriteUInt32(Number(v >> 32n))
        // if(val > UINT32MAX){
        //     this.WriteUInt32(UINT32MAX);
        //     this.WriteUInt32(val / UINT32MAX);
        // }else{
        //     this.WriteUInt32(val);
        //     this.WriteUInt32(0)
        // }
    }

    WriteUInt64(val:number){
        this.WriteInt64(val)
    }
    WriteString(val:string){
        if(!val){
            this.WriteByteArray([]);
            return
        }
        let arr = new TextEncoder().encode(val);
        this.WriteInt32(arr.length);
        this.WriteUint8ArrayBytes(arr)
    }
    WriteBool(val:boolean){
        this.WriteByte(val ? 1 : 0);
    }
    WriteByteArray(val:number[]){
        this.WriteUInt32(val.length)
        for (let i = 0; i < val.length; i++) {
            this.WriteByte(val[i]);
        }
    }
    WriteInt16Array(val:number[]){
        this.WriteUInt32(val.length)
        val.forEach(it=>{
            this.WriteUInt16(it);
        })
    }
    WriteInt32Array(val:number[]){
        this.WriteUInt32(val.length)
        val.forEach(it=>{
            this.WriteUInt32(it);
        })
    }
    WriteInt64Array(val:number[]){
        this.WriteUInt32(val.length)
        val.forEach(it=>{
            this.WriteUInt64(it);
        })
    }
    WriteStringArray(val:string[]){
        this.WriteUInt32(val.values.length)
        val.forEach(it=>{
            this.WriteString(it);
        })
    }
    WriteBytes(val:number[]){
        val.forEach(it => {
            this.WriteByte(it)
        });
    }
    WriteUint8ArrayBytes(val:Uint8Array){
        for (let i = 0; i < val.byteLength; i++) {
            this.WriteByte(val[i]);
        }
    }
    GetBytes():Uint8Array{
        return new Uint8Array(this.data)
    }
    GetNumberArray():number[]{
        return this.data
    }
}
