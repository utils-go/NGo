package directory

import (
	"os"
)

func CreateDirectory(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}
func Delete(path string) error {
	return os.RemoveAll(path)
}
func GetDirectories(path string) ([]string, error) {
	dirs, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	result := make([]string, 0)
	for _, dir := range dirs {
		result = append(result, dir.Name())
	}
	return result, nil
}
func GetFiles(path string) ([]string, error) {
	return nil, nil
}
func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return true
	} else {
		return false
	}
}
