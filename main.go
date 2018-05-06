package main

import (
	"./toml"
)

var blob = `
[stanza1]
xxx = 111
yyy = 222

[stanza2]
aaa = 111
bbb = true

[stanza3]
list = xxx
`

func main()  {
	toml.Decode(blob)
}