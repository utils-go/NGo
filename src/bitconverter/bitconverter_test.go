package bitconverter

import (
	"testing"
)

func TestByteConverterInt16(t *testing.T) {
	converter := BitConverter{}
	value := int16(-12)
	v, err := converter.GetBytesFromInt16(value)
	if err != nil {
		t.Error(err)
	}
	t.Logf("长度：%d\n", len(v))
	for _, b := range v {
		t.Log(b)
	}
	value1, err := converter.ToInt16(v, 0)
	if err != nil {
		t.Error(err)
	}
	if value != value1 {
		t.Errorf("%d is not equals to %d", value, value1)
	}
}
func TestByteConverterInt32(t *testing.T) {
	converter := BitConverter{}
	value := int32(-1211215)
	v, err := converter.GetBytesFromInt32(value)
	if err != nil {
		t.Error(err)
	}
	t.Logf("长度：%d\n", len(v))
	for _, b := range v {
		t.Log(b)
	}
	value1, err := converter.ToInt32(v, 0)
	if err != nil {
		t.Error(err)
	}
	if value != value1 {
		t.Errorf("%d is not equals to %d", value, value1)
	}
}
func TestByteConverterInt64(t *testing.T) {
	converter := BitConverter{}
	value := int64(-1121212121212121212)
	v, err := converter.GetBytesFromInt64(value)
	if err != nil {
		t.Error(err)
	}
	t.Logf("长度：%d\n", len(v))
	for _, b := range v {
		t.Log(b)
	}
	value1, err := converter.ToInt64(v, 0)
	if err != nil {
		t.Error(err)
	}
	if value != value1 {
		t.Errorf("%d is not equals to %d", value, value1)
	}
}

func TestByteConverterUInt16(t *testing.T) {
	converter := BitConverter{}
	value := uint16(12)
	v, err := converter.GetBytesFromUInt16(value)
	if err != nil {
		t.Error(err)
	}
	t.Logf("长度：%d\n", len(v))
	for _, b := range v {
		t.Log(b)
	}
	value1, err := converter.ToUInt16(v, 0)
	if err != nil {
		t.Error(err)
	}
	if value != value1 {
		t.Errorf("%d is not equals to %d", value, value1)
	}
}
func TestByteConverterUInt32(t *testing.T) {
	converter := BitConverter{}
	value := uint32(1211215)
	v, err := converter.GetBytesFromUInt32(value)
	if err != nil {
		t.Error(err)
	}
	t.Logf("长度：%d\n", len(v))
	for _, b := range v {
		t.Log(b)
	}
	value1, err := converter.ToUInt32(v, 0)
	if err != nil {
		t.Error(err)
	}
	if value != value1 {
		t.Errorf("%d is not equals to %d", value, value1)
	}
}
func TestByteConverterUInt64(t *testing.T) {
	converter := BitConverter{}
	value := uint64(1121212121212121212)
	v, err := converter.GetBytesFromUInt64(value)
	if err != nil {
		t.Error(err)
	}
	t.Logf("长度：%d\n", len(v))
	for _, b := range v {
		t.Log(b)
	}
	value1, err := converter.ToUInt64(v, 0)
	if err != nil {
		t.Error(err)
	}
	if value != value1 {
		t.Errorf("%d is not equals to %d", value, value1)
	}
}
func TestByteConverterBool(t *testing.T) {
	converter := BitConverter{}
	value := true
	v, err := converter.GetBytesFromBool(value)
	if err != nil {
		t.Error(err)
	}
	t.Logf("长度：%d\n", len(v))
	for _, b := range v {
		t.Log(b)
	}
	value1, err := converter.ToBoolean(v, 0)
	if err != nil {
		t.Error(err)
	}
	if value != value1 {
		t.Errorf("%v is not equals to %v", value, value1)
	}
}
func TestByteConverterDouble(t *testing.T) {
	converter := BitConverter{}
	value := float64(2.336554)
	v, err := converter.GetBytesFromDouble(value)
	if err != nil {
		t.Error(err)
	}
	t.Logf("长度：%d\n", len(v))
	for _, b := range v {
		t.Log(b)
	}
	value1, err := converter.ToDouble(v, 0)
	if err != nil {
		t.Error(err)
	}
	if value != value1 {
		t.Errorf("%v is not equals to %v", value, value1)
	}
}
