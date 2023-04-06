package datetime

import (
	"fmt"
	"testing"
	"time"
)

func TestDateTime_ToString(t *testing.T) {

	tt := NewDateTimeFromTime(time.Now())
	fmt.Println(tt.ToString("yyyy-MM-dd HH:mm:ss"))
	fmt.Println(tt.ToString("yyyy-MM-dd"))
	fmt.Println(tt.ToString("HH:mm:ss"))

	//fmt.Println(time.Now().Format("yyyyMMdd HH:mm:ss"))
}

func testdatetime() {

}
