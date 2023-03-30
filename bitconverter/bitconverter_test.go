package bitconverter

import (
	"testing"
)

func TestByteConverterInt16(t *testing.T) {
	converter := BitConverter{}
	value := int16(-12)
	v, err := converter.GetBytesFromInt16E(value)
	if err != nil {
		t.Error(err)
	}
	t.Logf("长度：%d\n", len(v))
	for _, b := range v {
		t.Log(b)
	}
	value1, err := converter.ToInt16E(v, 0)
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
	v, err := converter.GetBytesFromInt32E(value)
	if err != nil {
		t.Error(err)
	}
	t.Logf("长度：%d\n", len(v))
	for _, b := range v {
		t.Log(b)
	}
	value1, err := converter.ToInt32E(v, 0)
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
	v, err := converter.GetBytesFromInt64E(value)
	if err != nil {
		t.Error(err)
	}
	t.Logf("长度：%d\n", len(v))
	for _, b := range v {
		t.Log(b)
	}
	value1, err := converter.ToInt64E(v, 0)
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
	v, err := converter.GetBytesFromUInt16E(value)
	if err != nil {
		t.Error(err)
	}
	t.Logf("长度：%d\n", len(v))
	for _, b := range v {
		t.Log(b)
	}
	value1, err := converter.ToUInt16E(v, 0)
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
	v, err := converter.GetBytesFromUInt32E(value)
	if err != nil {
		t.Error(err)
	}
	t.Logf("长度：%d\n", len(v))
	for _, b := range v {
		t.Log(b)
	}
	value1, err := converter.ToUInt32E(v, 0)
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
	v, err := converter.GetBytesFromUInt64E(value)
	if err != nil {
		t.Error(err)
	}
	t.Logf("长度：%d\n", len(v))
	for _, b := range v {
		t.Log(b)
	}
	value1, err := converter.ToUInt64E(v, 0)
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
	v, err := converter.GetBytesFromBoolE(value)
	if err != nil {
		t.Error(err)
	}
	t.Logf("长度：%d\n", len(v))
	for _, b := range v {
		t.Log(b)
	}
	value1, err := converter.ToBooleanE(v, 0)
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
	v, err := converter.GetBytesFromDoubleE(value)
	if err != nil {
		t.Error(err)
	}
	t.Logf("长度：%d\n", len(v))
	for _, b := range v {
		t.Log(b)
	}
	value1, err := converter.ToDoubleE(v, 0)
	if err != nil {
		t.Error(err)
	}
	if value != value1 {
		t.Errorf("%v is not equals to %v", value, value1)
	}
}
