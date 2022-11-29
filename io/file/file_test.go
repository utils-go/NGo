package file

import (
	"testing"
	"time"
)

func TestAppedAllText(t *testing.T) {
	path := "test.txt"
	AppedAllText(path, time.Now().Format("2006-"))
}
