package bitconverter

import (
	"testing"
)

func TestGetBytes(t *testing.T) {
	converter := BitConverter{}
	value := int32(12)
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
