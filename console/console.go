package console

import "fmt"

func WriteLine(a ...any) {
	fmt.Println(a)
}
func Write(a ...any) {
	fmt.Print(a)
}
func ReadLine() string {
	var str string
	fmt.Scan(&str)
	return str
}
