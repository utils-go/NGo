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

func (b *BitConverter) GetBytesFromBoolE(value bool) ([]byte, error) {
	return b.toBytes(value)
}
func (b *BitConverter) GetBytesFromInt16E(value int16) ([]byte, error) {
	return b.toBytes(value)
}
func (b *BitConverter) GetBytesFromInt32E(value int32) ([]byte, error) {
	return b.toBytes(value)
}
func (b *BitConverter) GetBytesFromInt64E(value int64) ([]byte, error) {
	return b.toBytes(value)
}
func (b *BitConverter) GetBytesFromUInt16E(value uint16) ([]byte, error) {
	return b.toBytes(value)
}
func (b *BitConverter) GetBytesFromUInt32E(value uint32) ([]byte, error) {
	return b.toBytes(value)
}
func (b *BitConverter) GetBytesFromUInt64E(value uint64) ([]byte, error) {
	return b.toBytes(value)
}
func (b *BitConverter) GetBytesFromDoubleE(value float64) ([]byte, error) {
	return b.toBytes(value)
}
func (b *BitConverter) ToBooleanE(value []byte, startIndex int) (bool, error) {
	var result bool
	err := b.getValue(value[startIndex:], &result)
	if err != nil {
		return false, err
	}
	return result, nil
}

func (b *BitConverter) ToDoubleE(value []byte, startIndex int) (float64, error) {
	var result float64
	err := b.getValue(value[startIndex:], &result)
	if err != nil {
		return 0, err
	}

	return result, err
}
func (b *BitConverter) ToInt16E(value []byte, startIndex int) (int16, error) {
	var result int16
	err := b.getValue(value[startIndex:], &result)
	if err != nil {
		return 0, err
	}

	return result, err
}

func (b *BitConverter) ToInt32E(value []byte, startIndex int) (int32, error) {
	var result int32
	err := b.getValue(value[startIndex:], &result)
	if err != nil {
		return 0, err
	}

	return result, err
}

func (b *BitConverter) ToInt64E(value []byte, startIndex int) (int64, error) {
	var result int64
	err := b.getValue(value[startIndex:], &result)
	if err != nil {
		return 0, err
	}

	return result, err
}

func (b *BitConverter) ToUInt16E(value []byte, startIndex int) (uint16, error) {
	var result uint16
	err := b.getValue(value[startIndex:], &result)
	if err != nil {
		return 0, err
	}
	return result, err
}

func (b *BitConverter) ToUInt32E(value []byte, startIndex int) (uint32, error) {
	var result uint32
	err := b.getValue(value[startIndex:], &result)
	if err != nil {
		return 0, err
	}

	return result, err
}

func (b *BitConverter) ToUInt64E(value []byte, startIndex int) (uint64, error) {
	var result uint64
	err := b.getValue(value[startIndex:], &result)
	if err != nil {
		return 0, err
	}
	return result, err
}

func (b *BitConverter) GetBytesFromBool(value bool) []byte {
	v, _ := b.GetBytesFromBoolE(value)
	return v
}
func (b *BitConverter) GetBytesFromInt16(value int16) []byte {
	v, _ := b.GetBytesFromInt16E(value)
	return v
}
func (b *BitConverter) GetBytesFromInt32(value int32) []byte {
	v, _ := b.GetBytesFromInt32E(value)
	return v
}
func (b *BitConverter) GetBytesFromInt64(value int64) []byte {
	v, _ := b.toBytes(value)
	return v
}
func (b *BitConverter) GetBytesFromUInt16(value uint16) []byte {
	v, _ := b.GetBytesFromUInt16E(value)
	return v
}
func (b *BitConverter) GetBytesFromUInt32(value uint32) []byte {
	v, _ := b.GetBytesFromUInt32E(value)
	return v
}
func (b *BitConverter) GetBytesFromUInt64(value uint64) []byte {
	v, _ := b.GetBytesFromUInt64E(value)
	return v
}
func (b *BitConverter) GetBytesFromDouble(value float64) []byte {
	v, _ := b.GetBytesFromDoubleE(value)
	return v
}
func (b *BitConverter) ToBoolean(value []byte, startIndex int) bool {
	v, _ := b.ToBooleanE(value, startIndex)
	return v
}

func (b *BitConverter) ToDouble(value []byte, startIndex int) float64 {
	v, _ := b.ToDoubleE(value, startIndex)
	return v
}
func (b *BitConverter) ToInt16(value []byte, startIndex int) int16 {
	v, _ := b.ToInt16E(value, startIndex)
	return v
}

func (b *BitConverter) ToInt32(value []byte, startIndex int) int32 {
	v, _ := b.ToInt32E(value, startIndex)
	return v
}

func (b *BitConverter) ToInt64(value []byte, startIndex int) int64 {
	v, _ := b.ToInt64E(value, startIndex)
	return v
}

func (b *BitConverter) ToUInt16(value []byte, startIndex int) uint16 {
	v, _ := b.ToUInt16E(value, startIndex)
	return v
}

func (b *BitConverter) ToUInt32(value []byte, startIndex int) uint32 {
	v, _ := b.ToUInt32E(value, startIndex)
	return v
}

func (b *BitConverter) ToUInt64(value []byte, startIndex int) uint64 {
	v, _ := b.ToUInt64E(value, startIndex)
	return v
}
