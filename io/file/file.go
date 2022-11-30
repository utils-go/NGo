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

const nextLine string = "\r\n"

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

func copyFileWithFlag(srcFile, dstFile string, flag int) error {
	dir := filepath.Dir(dstFile)
	if err := os.MkdirAll(dir, fs.ModePerm); err != nil {
		//文件夹不存在，创建
		return err
	}

	src, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	dst, err := os.OpenFile(dstFile, flag, os.ModePerm) //file must not exist
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
	if err := dstWriter.Flush(); err != nil {
		return err
	}
	return nil
}

func writeFile(path string, contents []byte) error {
	f, err := os.OpenFile(path, os.O_TRUNC|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()
	buf := bufio.NewWriter(f)
	if _, err := buf.Write(contents); err != nil {
		return err
	}
	if err := buf.Flush(); err != nil { //保存
		return err
	}
	return nil
}
func appendFile(path string, contents []byte) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()
	buf := bufio.NewWriter(f)
	if _, err := buf.Write(contents); err != nil {
		return err
	}
	if err := buf.Flush(); err != nil {
		return err
	}
	return nil
}

func Copy(srcFile string, dstFile string) error {
	return copyFileWithFlag(srcFile, dstFile, os.O_EXCL|os.O_CREATE)
}

func CopyWithOverwrite(srcFile string, dstFile string) error {
	return copyFileWithFlag(srcFile, dstFile, os.O_CREATE)
}

func Move(srcFile, dstFile string) error {
	if err := Copy(srcFile, dstFile); err != nil {
		return err
	}
	if err := os.Remove(srcFile); err != nil {
		return err
	}
	return nil
}
func MoveWithOverwrite(srcFile, dstFile string) error {
	if err := CopyWithOverwrite(srcFile, dstFile); err != nil {
		return err
	}
	if err := os.Remove(srcFile); err != nil {
		return err
	}
	return nil
}
func ReadAllBytes(path string) ([]byte, error) {
	return readFile(path)
}
func ReadAllLines(path string) ([]string, error) {
	contents, err := readFile(path)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(contents), "\n"), nil
}
func ReadAllText(path string) (string, error) {
	contents, err := readFile(path)
	if err != nil {
		return "", err
	}
	return string(contents), nil
}
func AppendAllLines(path string, contents []string) error {
	str := strings.Join(contents, nextLine)
	str = str + nextLine
	return appendFile(path, []byte(str))
}
func AppendAllText(path string, str string) error {
	return appendFile(path, []byte(str))
}
func WriteAllBytes(path string, bytes []byte) error {
	return writeFile(path, bytes)
}
func WriteAllLines(path string, contents []string) error {
	str := strings.Join(contents, nextLine)
	str = str + nextLine
	return writeFile(path, []byte(str))
}

func WriteAllText(path string, content string) error {
	return writeFile(path, []byte(content))
}
