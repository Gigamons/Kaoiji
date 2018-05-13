package packets

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"reflect"

	"github.com/bnch/uleb128"
)

// this file includes functions (BString, RBString, IntArray) from https://github.com/bnch/bancho credits goes to thehowl. under MIT license.

// BString returns a binary array from a string that is encoded for client
func BString(s string) []byte {
	if s == "" {
		return []byte{0}
	}
	b := []byte{11}
	b = append(b, uleb128.Marshal(len(s))...)
	b = append(b, []byte(s)...)
	return b
}

// RBString reads a Osu! Encoded string
func RBString(value io.Reader) (s string, err error) {
	bufferSlice := make([]byte, 1)
	value.Read(bufferSlice)
	if bufferSlice[0] != 11 {
		return "", nil
	}
	length := uleb128.UnmarshalReader(value)
	bufferSlice = make([]byte, length)
	b, err := value.Read(bufferSlice)
	if b < length {
		err = errors.New("Unexpected end of string")
	}
	s = string(bufferSlice)
	return
}

// IntArray returns an Binary encoded IntArray
func IntArray(values []int32) []byte {
	b := new(bytes.Buffer)
	binary.Write(b, binary.LittleEndian, uint16(len(values)))
	binary.Write(b, binary.LittleEndian, values)
	return b.Bytes()
}

// RIntArray reads an IntArray
func RIntArray(value io.Reader) (i []int32, err error) {
	var length uint16
	err = binary.Read(value, binary.LittleEndian, &length)
	if err != nil {
		return
	}
	i = make([]int32, length)
	for y := 0; y < int(length); y++ {
		err = binary.Read(value, binary.LittleEndian, &i[y])
		if err != nil {
			return
		}
	}
	return
}

// RIntArray reads an IntArray
func ReadBeatmaps(value io.Reader) ([]int32, []string, error) {
	var beatmapFiles []string
	var beatmapIDs []int32
	var count int32
	var err error

	_ = beatmapIDs

	count, err = RInt32(value)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}

	beatmapFiles = make([]string, count)
	for i := 0; i < int(count); i++ {
		beatmapFiles[i], err = RBString(value)
		if err != nil {
			fmt.Println(err)
			return nil, nil, err
		}
	}

	count, err = RInt32(value)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}

	beatmapIDs = make([]int32, count)
	for i := 0; i < int(count); i++ {
		beatmapIDs[i], err = RInt32(value)
		if err != nil {
			fmt.Println(err)
			return nil, nil, err
		}
	}

	return beatmapIDs, beatmapFiles, nil
}

// Int8 returns a Binary Int8
func Int(value int) []byte {
	b := new(bytes.Buffer)
	binary.Write(b, binary.LittleEndian, value)
	return b.Bytes()
}

// RInt8 Reads an int8
func RInt(value io.Reader) (i int, err error) {
	err = binary.Read(value, binary.LittleEndian, &i)
	return
}

// Int8 returns a Binary Int8
func UInt(value uint) []byte {
	b := new(bytes.Buffer)
	binary.Write(b, binary.LittleEndian, value)
	return b.Bytes()
}

// RInt8 Reads an int8
func RUInt(value io.Reader) (i uint, err error) {
	err = binary.Read(value, binary.LittleEndian, &i)
	return
}

// Int8 returns a Binary Int8
func Int8(value int8) []byte {
	b := new(bytes.Buffer)
	binary.Write(b, binary.LittleEndian, value)
	return b.Bytes()
}

// RInt8 Reads an int8
func RInt8(value io.Reader) (i int8, err error) {
	err = binary.Read(value, binary.LittleEndian, &i)
	return
}

// UInt8 returns a Binary UInt8
func UInt8(value uint8) []byte {
	b := new(bytes.Buffer)
	binary.Write(b, binary.LittleEndian, value)
	return b.Bytes()
}

// RUInt8 Reads an int8
func RUInt8(value io.Reader) (i uint8, err error) {
	err = binary.Read(value, binary.LittleEndian, &i)
	return
}

// Int16 returns a Binary Int16
func Int16(value int16) []byte {
	b := new(bytes.Buffer)
	binary.Write(b, binary.LittleEndian, value)
	return b.Bytes()
}

// RInt16 Reads an int16
func RInt16(value io.Reader) (i int16, err error) {
	err = binary.Read(value, binary.LittleEndian, &i)
	return
}

// UInt16 returns a Binary UInt16
func UInt16(value uint16) []byte {
	b := new(bytes.Buffer)
	binary.Write(b, binary.LittleEndian, value)
	return b.Bytes()
}

// RUInt16 Reads an uint16
func RUInt16(value io.Reader) (i uint16, err error) {
	err = binary.Read(value, binary.LittleEndian, &i)
	return
}

// Int32 returns a Binary Int32
func Int32(value int32) []byte {
	b := new(bytes.Buffer)
	binary.Write(b, binary.LittleEndian, value)
	return b.Bytes()
}

// RInt32 eads an int32
func RInt32(value io.Reader) (i int32, err error) {
	err = binary.Read(value, binary.LittleEndian, &i)
	return
}

// UInt32 returns a Binary UInt32
func UInt32(value uint32) []byte {
	b := new(bytes.Buffer)
	binary.Write(b, binary.LittleEndian, value)
	return b.Bytes()
}

// RUInt32 eads an uint32
func RUInt32(value io.Reader) (i uint32, err error) {
	err = binary.Read(value, binary.LittleEndian, &i)
	return
}

// Int64 returns a Binary Int64
func Int64(value int64) []byte {
	b := new(bytes.Buffer)
	binary.Write(b, binary.LittleEndian, value)
	return b.Bytes()
}

// RInt64 eads an int64
func RInt64(value io.Reader) (i int64, err error) {
	err = binary.Read(value, binary.LittleEndian, &i)
	return
}

// UInt64 returns a Binary UInt64
func UInt64(value uint64) []byte {
	b := new(bytes.Buffer)
	binary.Write(b, binary.LittleEndian, value)
	return b.Bytes()
}

// RUInt64 eads an int64
func RUInt64(value io.Reader) (i uint64, err error) {
	err = binary.Read(value, binary.LittleEndian, &i)
	return
}

// UInt64 returns a Binary UInt64
func Float32(value float32) []byte {
	b := new(bytes.Buffer)
	binary.Write(b, binary.LittleEndian, value)
	return b.Bytes()
}

// RUInt64 eads an int64
func RFloat32(value io.Reader) (i float32, err error) {
	err = binary.Read(value, binary.LittleEndian, &i)
	return
}

// UInt64 returns a Binary UInt64
func Float64(value float64) []byte {
	b := new(bytes.Buffer)
	binary.Write(b, binary.LittleEndian, value)
	return b.Bytes()
}

// RUInt64 eads an int64
func RFloat64(value io.Reader) (i float64, err error) {
	err = binary.Read(value, binary.LittleEndian, &i)
	return
}

func Bool(value bool) []byte {
	b := new(bytes.Buffer)
	binary.Write(b, binary.LittleEndian, int8(func() int8 {
		if value {
			return int8(1)
		} else {
			return int8(0)
		}
	}()))
	return b.Bytes()
}

func RBool(value io.Reader) (i bool, err error) {
	var m int8
	err = binary.Read(value, binary.LittleEndian, &m)
	i = m > 0
	return
}

// MarshalBinary fast way to marshal Binary shit
func MarshalBinary(value interface{}) []byte {
	var buf = new(bytes.Buffer)
	var StructFields = reflect.ValueOf(value).Elem()
	for i := 0; i < StructFields.NumField(); i++ {
		t := StructFields.Field(i).Kind()
		vp := StructFields.Field(i)
		switch t {
		case reflect.Int:
			buf.Write(Int(int(vp.Int())))
			break
		case reflect.Uint:
			buf.Write(UInt(uint(vp.Uint())))
			break
		case reflect.Int8:
			buf.Write(Int8(int8(vp.Int())))
			break
		case reflect.Uint8:
			buf.Write(UInt8(uint8(vp.Uint())))
			break
		case reflect.Int16:
			buf.Write(Int16(int16(vp.Int())))
			break
		case reflect.Uint16:
			buf.Write(UInt16(uint16(vp.Uint())))
			break
		case reflect.Int32:
			buf.Write(Int32(int32(vp.Int())))
			break
		case reflect.Uint32:
			buf.Write(UInt32(uint32(vp.Uint())))
			break
		case reflect.Int64:
			buf.Write(Int64(int64(vp.Int())))
			break
		case reflect.Uint64:
			buf.Write(UInt64(uint64(vp.Uint())))
			break
		case reflect.String:
			buf.Write(BString(vp.String()))
			break
		case reflect.Float64:
			buf.Write(Float64(vp.Float()))
			break
		case reflect.Float32:
			buf.Write(Float32(float32(vp.Float())))
			break
		case reflect.Bool:
			buf.Write(Bool(bool(vp.Bool())))
			break
		}
	}
	return buf.Bytes()
}

// UnmarshalBinary unmarshals an OsuPacket data.
func UnmarshalBinary(value io.Reader, s interface{}) {
	var StructFields = reflect.ValueOf(s).Elem()
	for i := 0; i < StructFields.NumField(); i++ {
		t := StructFields.Field(i).Kind()
		vp := StructFields.Field(i)
		switch t {
		case reflect.Int:
			b, err := RInt(value)
			if err != nil {
				fmt.Println(err)
			}
			vp.SetInt(int64(b))
			break
		case reflect.Uint:
			b, err := RUInt(value)
			if err != nil {
				fmt.Println(err)
			}
			vp.SetUint(uint64(b))
			break
		case reflect.Int8:
			b, err := RInt8(value)
			if err != nil {
				fmt.Println(err)
			}
			vp.SetInt(int64(b))
			break
		case reflect.Uint8:
			b, err := RUInt8(value)
			if err != nil {
				fmt.Println(err)
			}
			vp.SetUint(uint64(b))
			break
		case reflect.Int16:
			b, err := RInt16(value)
			if err != nil {
				fmt.Println(err)
			}
			vp.SetInt(int64(b))
			break
		case reflect.Uint16:
			b, err := RUInt16(value)
			if err != nil {
				fmt.Println(err)
			}
			vp.SetUint(uint64(b))
			break
		case reflect.Int32:
			b, err := RInt32(value)
			if err != nil {
				fmt.Println(err)
			}
			vp.SetInt(int64(b))
			break
		case reflect.Uint32:
			b, err := RUInt32(value)
			if err != nil {
				fmt.Println(err)
			}
			vp.SetUint(uint64(b))
			break
		case reflect.Int64:
			b, err := RInt64(value)
			if err != nil {
				fmt.Println(err)
			}
			vp.SetInt(int64(b))
			break
		case reflect.Uint64:
			b, err := RUInt64(value)
			if err != nil {
				fmt.Println(err)
			}
			vp.SetUint(uint64(b))
			break
		case reflect.String:
			b, err := RBString(value)
			if err != nil {
				fmt.Println(err)
			}
			vp.SetString(b)
			break
		case reflect.Float64:
			b, err := RFloat64(value)
			if err != nil {
				fmt.Println(err)
			}
			vp.SetFloat(b)
			break
		case reflect.Float32:
			b, err := RFloat32(value)
			if err != nil {
				fmt.Println(err)
			}
			vp.SetFloat(float64(b))
			break
		case reflect.Bool:
			b, err := RBool(value)
			if err != nil {
				fmt.Println(err)
			}
			vp.SetBool(b)
			break
		}
	}
}
