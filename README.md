# NGo
the .net api implement by go.you can write go as like .net

this is based on .net4.7.2 api

refrence:https://learn.microsoft.com/en-us/dotnet/api/?view=netframework-4.7.2&preserve-view=true

# Usage
- import ngo
 ```
 go get -u github.com/lishuangquan1987/ngo
 ```
- use ngo
file:
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
bitconverter:
```go
package main

import (
	"fmt"
	"github.com/lishuangquan1987/ngo/bitconverter"
)

func main() {
	converter := bitconverter.BitConverter{IsLittleEndian: true} //littleEndian at the front
	value := int64(11234567489)
	bytes, err := converter.GetBytesFromInt64(value)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bytes)
	valueNew, err := converter.ToInt64(bytes, 0)
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

# InProgress:
- Diretory+
- DateTime
- TimeSpan
- ...