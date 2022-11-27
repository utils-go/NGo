package file

import (
	"bufio"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

//refrence:https://learn.microsoft.com/en-us/dotnet/api/system.io?view=netframework-4.7.2

func readFile(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	result := make([]byte, 0, 1024)
	for {
		buffer := make([]byte, 0, 1024)
		n, err := reader.Read(buffer)
		if err != nil {
			if err == io.EOF {
				result = append(result, buffer[:n]...)
				break
			} else {
				return nil, err
			}
		}
		result = append(result, buffer[:n]...)
	}
	return result, nil
}

func AppendAllLines(path string) ([]string, error) {
	contents, err := readFile(path)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(contents), "\n"), nil
}
func AppedAllText(path string) (string, error) {
	contents, err := readFile(path)
	if err != nil {
		return "", err
	}
	return string(contents), nil
}
func Copy(srcFile string, dstFile string) error {
	dir := filepath.Dir(dstFile)
	if err := os.MkdirAll(dir, fs.ModePerm); err != nil {
		//文件夹不存在，创建
		return err
	}

	src, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	dst, err := os.OpenFile(dstFile, os.O_EXCL, os.ModePerm)
	if err != nil {
		return err
	}
	defer src.Close()
	defer dst.Close()

	srcReader := bufio.NewReader(src)
	dstWriter := bufio.NewWriter(dst)
	for {
		tempBuf := make([]byte, 1024)
		n, err := srcReader.Read(tempBuf)
		if err != nil {
			if err == io.EOF {
				dstWriter.Write(tempBuf[:n])
				break
			} else {
				return err
			}
		}
		dstWriter.Write(tempBuf[:n])
	}
	return nil
}

func CopyWithOverwrite(srcFile string, dstFile string) error {

}

func Move(srcFile, dstFile string) error {

}
func MoveWithOverwrite(srcFile, dstFile string) error {

}
func ReadAllBytes(path string) ([]byte, error) {

}
func ReadAllLines(path string) ([]string, error) {

}
func ReadAllText(path string) (string, error) {

}
func WriteAllBytes(path string, bytes []byte) error {

}
func WriteAllLines(path string, contents []string) error {

}
func WriteAllText(path string, content string) error {

}
