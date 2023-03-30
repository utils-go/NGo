# NGo
`ngo` means .net api implement by go

the .net api implement by go. you can write go as like .net

this is based on .net4.7.2 api

api reference:https://learn.microsoft.com/en-us/dotnet/api/?view=netframework-4.7.2&preserve-view=true

source code reference:https://referencesource.microsoft.com/
# Usage
 ```
 go get -u github.com/lishuangquan1987/ngo
 ```
file: same as `Sytem.IO.File`
```go
package main

import (
	"fmt"
	"github.com/lishuangquan1987/ngo/io/file"
)

func main() {
	content := "this is ngo example"
	file.WriteAllText("./test.txt", content)
	result, err := file.ReadAllText("./test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}

```
bitconverter: same as `System.BitConverter`,but it is not static methods,so you can change `LittleEndian` or `BigEndian` any where
```go
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

```


# Finish:
- File
- Path
- BitConverter
- Diretory
- DateTime
- TimeSpan
- Console

# InProgress:
- ...
