package helpers

import (
	"bytes"
	"encoding/binary"
	"io"
)

func GetBytes(value interface{}, params ...bool) []byte {
	writer := new(bytes.Buffer)

	switch value.(type) {
	case uint8:
		binary.Write(writer, binary.LittleEndian, value.(uint8))
	case int8:
		binary.Write(writer, binary.LittleEndian, value.(int8))
	case uint16:
		binary.Write(writer, binary.LittleEndian, value.(uint16))
	case int16:
		binary.Write(writer, binary.LittleEndian, value.(int16))
	case uint32:
		binary.Write(writer, binary.LittleEndian, value.(uint32))
	case int32:
		binary.Write(writer, binary.LittleEndian, value.(int32))
	case uint: // uint is always 64 bit.
	case uint64:
		binary.Write(writer, binary.LittleEndian, value.(uint64))
	case int: // int is always 64 bit.
	case int64:
		binary.Write(writer, binary.LittleEndian, value.(int64))
	case float32:
		binary.Write(writer, binary.LittleEndian, value.(float32))
	case float64:
		binary.Write(writer, binary.LittleEndian, value.(float64))
	case string:
		if len(value.(string)) == 0 && len(params) > 0 && params[0] == true {
			writer.WriteByte(byte(0))
		} else {
			writer.WriteByte(byte(11))
			writer.WriteByte(byte(len(value.(string))))
			writer.WriteString(value.(string))
		}
	case []int32:
		writer.Write(GetBytes(uint16(len(value.([]int32)))))

		for _, item := range value.([]int32) {
			writer.Write(GetBytes(item))
		}

	}

	return writer.Bytes()
}

func WriteBytes(w io.Writer, value interface{}, params ...bool) (err error) {
	switch value.(type) {
	case uint8:
		err = binary.Write(w, binary.LittleEndian, value.(uint8))
	case int8:
		err = binary.Write(w, binary.LittleEndian, value.(int8))
	case uint16:
		err = binary.Write(w, binary.LittleEndian, value.(uint16))
	case int16:
		err = binary.Write(w, binary.LittleEndian, value.(int16))
	case uint32:
		err = binary.Write(w, binary.LittleEndian, value.(uint32))
	case int32:
		err = binary.Write(w, binary.LittleEndian, value.(int32))
	case uint: // uint is always 64 bit.
	case uint64:
		err = binary.Write(w, binary.LittleEndian, value.(uint64))
	case int: // int is always 64 bit.
	case int64:
		err = binary.Write(w, binary.LittleEndian, value.(int64))
	case float32:
		err = binary.Write(w, binary.LittleEndian, value.(float32))
	case float64:
		err = binary.Write(w, binary.LittleEndian, value.(float64))
	case string:
		if len(value.(string)) == 0 && len(params) > 0 && params[0] == true {
			_, err = w.Write([]byte{0})
		} else {
			_, err = w.Write([]byte{11})
			_, err = w.Write([]byte{byte(len(value.(string)))})
			_, err = w.Write([]byte(value.(string)))
		}
	case []int32:
		err = WriteBytes(w, uint16(len(value.([]int32))))

		for _, item := range value.([]int32) {
			err = WriteBytes(w, item)
		}
	}

	return
}

