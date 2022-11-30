package path

import (
	"path/filepath"
	"strings"
)

// Refrence:https://learn.microsoft.com/en-us/dotnet/api/system.io.path?view=netframework-4.7.2

func ChangeExtension(path, extension string) string {
	ext := filepath.Ext(path)
	fullFileNameWithoutExtension := strings.TrimSuffix(path, ext)

	extensionNew := extension
	if !strings.HasPrefix(extensionNew, ".") {
		extensionNew = "." + extensionNew
	}
	return filepath.Join(fullFileNameWithoutExtension, extensionNew)
}
func Combine(path string, paths ...string) string {
	pathList := []string{path}
	pathList = append(pathList, paths...)
	return filepath.Join(pathList...)
}

func GetDirectoryName(path string) string {
	return filepath.Dir(path)
}
func GetExtension(path string) string {
	return filepath.Ext(path)
}
func GetFileName(path string) string {
	return filepath.Base(path)
}
func GetFileNameWithoutExtension(path string) string {
	fileName := GetFileName(path)
	extension := GetExtension(path)
	return strings.TrimSuffix(fileName, extension)
}
func GetFullPath(path string) (string, error) {
	return filepath.Abs(path)
}

func HasExtension(path string) bool {
	return filepath.Ext(path) != ""
}

// judge wether the path is same
func Equals(path1, path2 string) bool {
	newPath1 := strings.ReplaceAll(path1, "\\", "/")
	newPath2 := strings.ReplaceAll(path2, "\\", "/")
	return newPath1 == newPath2
}
