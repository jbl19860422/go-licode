/*
The MIT License (MIT)

Copyright (c) 2013-2015 GOSRS(gosrs)

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
the Software, and to permit persons to whom the Software is furnished to do so,
subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package nice
import (
	"encoding/binary"
	"bytes"
	"net"
	"errors"
	"unsafe"
)

func numberToBytes(data interface{}, order binary.ByteOrder) []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, order, data)
	if err != nil {
		return nil
	}
	return buf.Bytes()
}

func bytesToNumber(data []byte, order binary.ByteOrder, v interface{}) error {
	buf := bytes.NewReader(data)
	err := binary.Read(buf, order, v)
	return err
}

func UInt16ToBytes(data uint16, order binary.ByteOrder) []byte {
	return numberToBytes(data, order)
}

func UInt32ToBytes(data uint32, order binary.ByteOrder) []byte {
	return numberToBytes(data, order)
}

func UInt64ToBytes(data uint64, order binary.ByteOrder) []byte {
	return numberToBytes(data, order)
}

func Int16ToBytes(data int16, order binary.ByteOrder) []byte {
	return numberToBytes(data, order)
}

func Int32ToBytes(data int32, order binary.ByteOrder) []byte {
	return numberToBytes(data, order)
}

func Int64ToBytes(data int64, order binary.ByteOrder) []byte {
	return numberToBytes(data, order)
}

func Float32ToBytes(data float32, order binary.ByteOrder) []byte {
	return numberToBytes(data, order)
}

func Float64ToBytes(data float64, order binary.ByteOrder) []byte {
	return numberToBytes(data, order)
}

func BytesToUInt16(data []byte, order binary.ByteOrder) (uint16, error) {
	var v uint16 = 0
	err := bytesToNumber(data, order, &v)
	return v, err
}

func BytesToUInt32(data []byte, order binary.ByteOrder) (uint32, error) {
	var v uint32 = 0
	err := bytesToNumber(data, order, &v)
	return v, err
}

func BytesToUInt64(data []byte, order binary.ByteOrder) (uint64, error) {
	var v uint64 = 0
	err := bytesToNumber(data, order, &v)
	return v, err
}

func BytesToInt16(data []byte, order binary.ByteOrder) (int16, error) {
	var v int16 = 0
	err := bytesToNumber(data, order, &v)
	return v, err
}

func BytesToInt32(data []byte, order binary.ByteOrder) (int32, error) {
	var v int32 = 0
	err := bytesToNumber(data, order, &v)
	return v, err
}

func BytesToInt64(data []byte, order binary.ByteOrder) (int64, error) {
	var v int64 = 0
	err := bytesToNumber(data, order, &v)
	return v, err
}

func BytesToFloat32(data []byte, order binary.ByteOrder) (float32, error) {
	var v float32 = 0
	err := bytesToNumber(data, order, &v)
	return v, err
}

func BytesToFloat64(data []byte, order binary.ByteOrder) (float64, error) {
	var v float64 = 0
	err := bytesToNumber(data, order, &v)
	return v, err
}

func MaxUInt32(a uint32, b uint32) uint32 {
	if a > b {
		return a
	}
	return b
}

func MinUInt32(a uint32, b uint32) uint32 {
	if a < b {
		return a
	}
	return b
}

func SliceInsert(s []interface{}, index int, value interface{}) []interface{} {
	rear := append([]interface{}{}, s[index:]...)
	return append(append(s[:index], value), rear...)
}

func UdpAddrPortUsable(ip string, port int) bool {
	_, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.ParseIP(ip), Port: port})
	if err != nil {
		return false
	}
	return true
}

type DataStream struct {
	// current position at bytes.
	p []byte
	// the bytes data for stream to read or write.
	bytes []byte
	// current position
	pos uint32
}

func NewDataStream(data []byte) *DataStream {
	return &DataStream{
		p:     data,
		bytes: data,
		pos:   0,
	}
}

func (this *DataStream) Data() []byte {
	return this.bytes
}

func (this *DataStream) Size() uint32 {
	return uint32(len(this.bytes))
}

func (this *DataStream) Empty() bool {
	return this.bytes == nil || len(this.p) <= 0
}

func (this *DataStream) Require(required_size uint32) bool {
	return required_size <= uint32(len(this.p))
}

func (this *DataStream) Pos() uint32 {
	return this.pos
}

func (this *DataStream) Skip(size uint32) {
	this.pos += size
	this.p = this.bytes[this.pos:]
}

func (this *DataStream) PeekByte() (byte, error) {
	if !this.Require(1) {
		err := errors.New("DataStream not have enough data")
		return 0, err
	}
	return this.p[0], nil
}

func (this *DataStream) ReadByte() (byte, error) {
	if !this.Require(1) {
		err := errors.New("DataStream not have enough data")
		return 0, err
	}

	b := this.p[0]
	this.Skip(1)
	return b, nil
}

func (this *DataStream) WriteByte(data byte) {
	this.bytes = append(this.bytes, data)
}

func (this *DataStream) ReadBytes(count uint32) ([]byte, error) {
	if !this.Require(count) {
		err := errors.New("DataStream not have enough data")
		return nil, err
	}

	b := this.p[0:count]
	this.Skip(count)
	return b, nil
}

func (this *DataStream) ReadLeftBytes() []byte {
	l := len(this.p)
	b := this.p
	this.Skip(uint32(l))
	return b
}

func (this *DataStream) CopyLeftBytes() []byte {
	b := make([]byte, len(this.p))
	copy(b, this.p)
	return b
}

func (this *DataStream) PeekLeftBytes() []byte {
	return this.p
}

func (this *DataStream) PeekBytes(count uint32) ([]byte, error) {
	if !this.Require(count) {
		err := errors.New("DataStream not have enough data")
		return nil, err
	}

	b := make([]byte, count)
	copy(b, this.p[0:count])
	return b, nil
}

func (this *DataStream) WriteBytes(data []byte) {
	this.bytes = append(this.bytes, data...)
}

func (this *DataStream) ReadBool() (bool, error) {
	b, err := this.ReadByte()
	if err != nil {
		return false, err
	}
	if b == 0x01 {
		return true, nil
	} else {
		return false, nil
	}
}

func (this *DataStream) WriteBool(data bool) {
	var d byte
	if data {
		d = 1
	} else {
		d = 0
	}
	this.WriteByte(d)
}

func (this *DataStream) ReadInt8() (int8, error) {
	var b byte
	var err error
	if b, err = this.ReadByte(); err != nil {
		return 0, err
	}

	return int8(b), nil
}

func (this *DataStream) WriteInt8(d int8) error {
	this.WriteByte(byte(d))
	return nil
}

func (this *DataStream) ReadUInt8() (uint8, error) {
	var b byte
	var err error
	if b, err = this.ReadByte(); err != nil {
		return 0, err
	}

	return uint8(b), nil
}

func (this *DataStream) WriteUInt8(d uint8) error {
	this.WriteByte(byte(d))
	return nil
}

func (this *DataStream) ReadInt16(order binary.ByteOrder) (int16, error) {
	b, err := this.ReadBytes(2)
	if err != nil {
		return 0, err
	}

	v, err := BytesToInt16(b, order)
	return v, err
}

func (this *DataStream) ReadUInt16(order binary.ByteOrder) (uint16, error) {
	b, err := this.ReadBytes(2)
	if err != nil {
		return 0, err
	}

	v, err := BytesToUInt16(b, order)
	return v, err
}


func (this *DataStream) WriteInt16(data int16, order binary.ByteOrder) {
	b := Int16ToBytes(data, order)
	this.WriteBytes(b)
}

func (this *DataStream) WriteUInt16(data uint16, order binary.ByteOrder) {
	b := UInt16ToBytes(data, order)
	this.WriteBytes(b)
}

func (this *DataStream) ReadInt32(order binary.ByteOrder) (int32, error) {
	b, err := this.ReadBytes(4)
	if err != nil {
		return 0, err
	}

	v, err := BytesToInt32(b, order)
	return v, err
}

func (this *DataStream) WriteInt32(data int32, order binary.ByteOrder) {
	b := Int32ToBytes(data, order)
	this.WriteBytes(b)
}

func (this *DataStream) WriteUInt32(data uint32, order binary.ByteOrder) {
	b := UInt32ToBytes(data, order)
	this.WriteBytes(b)
}


func (this *DataStream) ReadInt64(order binary.ByteOrder) (int64, error) {
	b, err := this.ReadBytes(8)
	if err != nil {
		return 0, err
	}

	v, err := BytesToInt64(b, order)
	return v, err
}

func (this *DataStream) WriteInt64(data int64, order binary.ByteOrder) {
	b := Int64ToBytes(data, order)
	this.WriteBytes(b)
}

func (this *DataStream) ReadFloat32(order binary.ByteOrder) (float32, error) {
	b, err := this.ReadBytes(4)
	if err != nil {
		return 0, err
	}

	v, err := BytesToFloat32(b, order)
	return v, err
}

func (this *DataStream) WriteFloat32(data float32, order binary.ByteOrder) {
	b := Float32ToBytes(data, order)
	this.WriteBytes(b)
}

func (this *DataStream) ReadFloat64(order binary.ByteOrder) (float64, error) {
	b, err := this.ReadBytes(8)
	if err != nil {
		return 0, err
	}

	v, err := BytesToFloat64(b, order)
	return v, err
}

func (this *DataStream) WriteFloat64(data float64, order binary.ByteOrder) {
	b := Float64ToBytes(data, order)
	this.WriteBytes(b)
}

func (this *DataStream) ReadString(len uint32) (string, error) {
	if !this.Require(len) {
		err := errors.New("no enough data")
		return "", err
	}

	str := string(this.p[:len])
	this.Skip(len)
	return str, nil
}

func (this *DataStream) WriteString(str string) {
	this.WriteBytes([]byte(str))
}

var HostByteOrder binary.ByteOrder
func IsLittleEndian() bool {
	var i int32 = 0x01020304
	u := unsafe.Pointer(&i)
	pb := (*byte)(u)
	b := *pb
	return (b == 0x04)
}

func XOR(a []byte, b[]byte) ([]byte, error) {
	if 	len(a) != len(b) {
		return nil, errors.New("XOR must the same len")
	}

	m := make([]byte, len(a))
	for i := 0; i < len(m); i++ {
		m[i] = a[i] ^ b[i]
	}
	return m, nil
}

func ToNetworkOrder(in []byte) []byte {
	if HostByteOrder == binary.LittleEndian {
		var i int = 0
		var j int = len(in) - 1
		for i = 0; i < j; i++ {
			b := in[i]
			in[i] = in[j]
			in[j] = b
			j--
		}
	}
	return in
}

func init() {
	if IsLittleEndian() {
		HostByteOrder = binary.LittleEndian
	} else {
		HostByteOrder = binary.BigEndian
	}
}