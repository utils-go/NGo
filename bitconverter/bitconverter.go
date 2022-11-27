package bitconverter

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
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
func (b *BitConverter) getValue(value []byte) (v any, err error) {
	buffer := bytes.NewBuffer(value)

	if b.IsLittleEndian {
		err = binary.Read(buffer, binary.LittleEndian, v)
	} else {
		err = binary.Read(buffer, binary.BigEndian, v)
	}
	return v, err
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
	v, err := b.getValue(value[startIndex:])
	if err != nil {
		return false, err
	}

	result, ok := v.(bool)
	if !ok {
		return false, errors.New(fmt.Sprintf("%v can not convert to bool", v))
	}

	return result, err
}

func (b *BitConverter) ToDouble(value []byte, startIndex int) (float64, error) {
	v, err := b.getValue(value[startIndex:])
	if err != nil {
		return 0, err
	}

	result, ok := v.(float64)
	if !ok {
		return 0, errors.New(fmt.Sprintf("%v can not convert to float64", v))
	}

	return result, err
}
func (b *BitConverter) ToInt16(value []byte, startIndex int) (int16, error) {
	v, err := b.getValue(value[startIndex:])
	if err != nil {
		return 0, err
	}

	result, ok := v.(int16)
	if !ok {
		return 0, errors.New(fmt.Sprintf("%v can not convert to int16", v))
	}

	return result, err
}

func (b *BitConverter) ToInt32(value []byte, startIndex int) (int32, error) {
	v, err := b.getValue(value[startIndex:])
	if err != nil {
		return 0, err
	}

	result, ok := v.(int32)
	if !ok {
		return 0, errors.New(fmt.Sprintf("%v can not convert to int32", v))
	}

	return result, err
}

func (b *BitConverter) ToInt64(value []byte, startIndex int) (int64, error) {
	v, err := b.getValue(value[startIndex:])
	if err != nil {
		return 0, err
	}

	result, ok := v.(int64)
	if !ok {
		return 0, errors.New(fmt.Sprintf("%v can not convert to int64", v))
	}

	return result, err
}

func (b *BitConverter) ToUInt16(value []byte, startIndex int) (uint16, error) {
	v, err := b.getValue(value[startIndex:])
	if err != nil {
		return 0, err
	}

	result, ok := v.(uint16)
	if !ok {
		return 0, errors.New(fmt.Sprintf("%v can not convert to uint16", v))
	}

	return result, err
}

func (b *BitConverter) ToUInt32(value []byte, startIndex int) (uint32, error) {
	v, err := b.getValue(value[startIndex:])
	if err != nil {
		return 0, err
	}

	result, ok := v.(uint32)
	if !ok {
		return 0, errors.New(fmt.Sprintf("%v can not convert to uint32", v))
	}

	return result, err
}

func (b *BitConverter) ToUInt64(value []byte, startIndex int) (uint64, error) {
	v, err := b.getValue(value[startIndex:])
	if err != nil {
		return 0, err
	}

	result, ok := v.(uint64)
	if !ok {
		return 0, errors.New(fmt.Sprintf("%v can not convert to uint64", v))
	}

	return result, err
}
