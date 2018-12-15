package helpers

import (
	"bytes"
	"encoding/binary"
)

func GetUintBytes(value interface{}) []byte {
	writer := new(bytes.Buffer)

	switch value.(type) {
	case uint8:
		binary.Write(writer, binary.LittleEndian, value.(uint8))
	case uint16:
		binary.Write(writer, binary.LittleEndian, value.(uint16))
	case uint32:
		binary.Write(writer, binary.LittleEndian, value.(uint32))
	case uint64:
		binary.Write(writer, binary.LittleEndian, value.(uint64))

	}

	return writer.Bytes()
}

func GetIntBytes(value interface{}) []byte {
	writer := new(bytes.Buffer)

	switch value.(type) {
	case int8:
		binary.Write(writer, binary.LittleEndian, value.(int8))
	case int16:
		binary.Write(writer, binary.LittleEndian, value.(int16))
	case int32:
		binary.Write(writer, binary.LittleEndian, value.(int32))
	case int64:
		binary.Write(writer, binary.LittleEndian, value.(int64))

	}

	return writer.Bytes()
}

func GetFloatByte(value interface{}) []byte {
	writer := new(bytes.Buffer)

	switch value.(type) {
	case float32:
		binary.Write(writer, binary.LittleEndian, value.(float32))
	case float64:
		binary.Write(writer, binary.LittleEndian, value.(float64))

	}

	return writer.Bytes()
}

func GetStringBytes(value string, params ...bool) []byte {
	writer := new(bytes.Buffer)

	if len(value) == 0 && len(params) > 0 && params[0] == true {
		writer.Write(GetUintBytes(0))
	} else {
		writer.WriteByte(byte(11))
		writer.WriteByte(byte(len(value)))
		writer.WriteString(value)
	}

	return writer.Bytes()
}

func GetArrayBytes(value []int32) []byte {
	writer := new(bytes.Buffer)

	writer.Write(GetUintBytes(uint16(len(value))))

	for _, item := range value {
		writer.Write(GetIntBytes(item))
	}

	return writer.Bytes()
}
