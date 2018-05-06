package main

import (
	"./toml"
	"fmt"
)

var blob = `
[stanza1]
xxx = 111
yyy = 222

[stanza2]
aaa = 111
bbb = true

# comment 

[stanza3]
list = xxx
`

func main()  {
	data, err := toml.Decode(blob)
	fmt.Println(data)
	fmt.Println(err)

	enc := toml.Encode(data)
	fmt.Println(enc)
}