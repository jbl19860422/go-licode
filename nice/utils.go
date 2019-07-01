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