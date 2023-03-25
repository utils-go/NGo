package path

import (
	"testing"
)

func TestChangeExtension(t *testing.T) {
	path := "c:/a/b/c.txt"
	newExtension := ".pdf"
	newPath := "c:\\a\\b\\c.pdf"
	result := ChangeExtension(path, newExtension)
	if Equals(result, newPath) {
		t.Errorf("%s is not equals %s\n", newPath, result)
		return
	}
}
func TestCombine(t *testing.T) {
	path := "c:/a/b"
	fileName := "c.txt"
	newPath := Combine(path, fileName)
	if !Equals(newPath, "c:/a/b/c.txt") {
		t.Errorf("%s is not expected\n", newPath)
		return
	}
}
func TestGetDirectoryName(t *testing.T) {
	path := "c:/a/b/c.txt"
	dir := "c:/a/b"
	dirResult := GetDirectoryName(path)
	if !Equals(dir, dirResult) {
		t.Errorf("%s is not expected\n", dirResult)
		return
	}
}

func TestGetExtension(t *testing.T) {
	path := "c:/a/b/c.txt"
	extension := ".txt"
	result := GetExtension(path)
	if result != extension {
		t.Errorf("%s is not expected\n", result)
		return
	}
}
func TestGetFileName(t *testing.T) {
	path := "c:/a/b/c.txt"
	fileName := "c.txt"
	result := GetFileName(path)
	if result != fileName {
		t.Errorf("%s is not expected\n", result)
		return
	}
}
func TestGetFileNameWithoutExtension(t *testing.T) {
	path := "c:/a/b/c.txt"
	fileName := "c"
	result := GetFileNameWithoutExtension(path)
	if result != fileName {
		t.Errorf("%s is not expected\n", result)
		return
	}
}
func TestGetFullPath(t *testing.T) {
	path := "./c.txt"
	fileName := "E:/GoProject/ngo/io/path/c.txt"
	result, err := GetFullPath(path)
	if err != nil {
		t.Error(err)
		return
	}
	if !Equals(result, fileName) {
		t.Errorf("%s is not expected\n", result)
		return
	}
}
func TestHasExtension(t *testing.T) {
	path := "./c.txt"

	if result := HasExtension(path); !result {
		t.Errorf("%v is not expected\n", result)
		return
	}
}
func TestHasExtension2(t *testing.T) {
	path := "c:/c"

	if result := HasExtension(path); result {
		t.Errorf("%v is not expected\n", result)
		return
	}
}
