Sure, I would be happy to help you improve the format of your README.MD file. Here are some formatting suggestions for your consideration:
 ## Project Name
Provide a brief introduction to your project here, including its purpose and key features.
### Installation
Run the following command to install the project:
go get -u github.com/lishuangquan1987/ngo
### Usage
#### File
Provide instructions on how to use the File package here.
package main
import (
	"fmt"
	"github.com/lishuangquan1987/ngo/io/file"
)
func main() {
	content := "This is an ngo example"
	file.WriteAllText("./test.txt", content)
	result, err := file.ReadAllText("./test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
#### BitConverter
Provide instructions on how to use the BitConverter package here.
package main
import (
	"fmt"
	"github.com/lishuangquan1987/ngo/bitconverter"
)
func main() {
	converter := bitconverter.BitConverter{IsLittleEndian: true} //littleEndian at the front
	value := int64(11234567489)
	bytes, err := converter.GetBytesFromInt64E(value)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bytes)
	valueNew, err := converter.ToInt64E(bytes, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(value == valueNew)
}
### Progress
- [x] File
- [x] Path
- [x] BitConverter
- [x] Diretory
- [x] DateTime
- [x] TimeSpan
- [x] Console
- [ ] ...
### Contact
Provide a way for users to contact you here for reporting issues or providing feedback.
I hope these formatting suggestions help you improve your README.MD file.