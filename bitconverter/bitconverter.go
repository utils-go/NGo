package bitconverter

import (
	"bytes"
	"encoding/binary"
)

type BitConverter struct {
	//是否是小端在前
	IsLittleEndian bool
}

func (b *BitConverter) toBytes(value any) ([]byte, error) {
	buffer := bytes.NewBuffer([]byte{})
	var err error = nil
	if b.IsLittleEndian {
		err = binary.Write(buffer, binary.LittleEndian, value)
	} else {
		err = binary.Write(buffer, binary.BigEndian, value)
	}
	return buffer.Bytes(), err
}
func (b *BitConverter) getValue(value []byte, v any) (err error) {
	buffer := bytes.NewBuffer(value)

	if b.IsLittleEndian {
		err = binary.Read(buffer, binary.LittleEndian, v)
	} else {
		err = binary.Read(buffer, binary.BigEndian, v)
	}
	return err
}

func (b *BitConverter) GetBytesFromBool(value bool) ([]byte, error) {
	return b.toBytes(value)
}
func (b *BitConverter) GetBytesFromInt16(value int16) ([]byte, error) {
	return b.toBytes(value)
}
func (b *BitConverter) GetBytesFromInt32(value int32) ([]byte, error) {
	return b.toBytes(value)
}
func (b *BitConverter) GetBytesFromInt64(value int64) ([]byte, error) {
	return b.toBytes(value)
}
func (b *BitConverter) GetBytesFromUInt16(value uint16) ([]byte, error) {
	return b.toBytes(value)
}
func (b *BitConverter) GetBytesFromUInt32(value uint32) ([]byte, error) {
	return b.toBytes(value)
}
func (b *BitConverter) GetBytesFromUInt64(value uint64) ([]byte, error) {
	return b.toBytes(value)
}
func (b *BitConverter) ToBoolean(value []byte, startIndex int) (bool, error) {
	var result bool
	err := b.getValue(value[startIndex:], &result)
	if err != nil {
		return false, err
	}
	return result, nil
}

func (b *BitConverter) ToDouble(value []byte, startIndex int) (float64, error) {
	var result float64
	err := b.getValue(value[startIndex:], &result)
	if err != nil {
		return 0, err
	}

	return result, err
}
func (b *BitConverter) ToInt16(value []byte, startIndex int) (int16, error) {
	var result int16
	err := b.getValue(value[startIndex:], &result)
	if err != nil {
		return 0, err
	}

	return result, err
}

func (b *BitConverter) ToInt32(value []byte, startIndex int) (int32, error) {
	var result int32
	err := b.getValue(value[startIndex:], &result)
	if err != nil {
		return 0, err
	}

	return result, err
}

func (b *BitConverter) ToInt64(value []byte, startIndex int) (int64, error) {
	var result int64
	err := b.getValue(value[startIndex:], &result)
	if err != nil {
		return 0, err
	}

	return result, err
}

func (b *BitConverter) ToUInt16(value []byte, startIndex int) (uint16, error) {
	var result uint16
	err := b.getValue(value[startIndex:], &result)
	if err != nil {
		return 0, err
	}
	return result, err
}

func (b *BitConverter) ToUInt32(value []byte, startIndex int) (uint32, error) {
	var result uint32
	err := b.getValue(value[startIndex:], &result)
	if err != nil {
		return 0, err
	}

	return result, err
}

func (b *BitConverter) ToUInt64(value []byte, startIndex int) (uint64, error) {
	var result uint64
	err := b.getValue(value[startIndex:], &result)
	if err != nil {
		return 0, err
	}
	return result, err
}
