package main

import (
	"./toml"
	"fmt"
	"time"
	"io/ioutil"
)

var blob = `
[stanza1]
xxx = 111
yyy = 222

[stanza2]
aaa = 111
bbb = true

# comment 

[stanza-3]
listl = xxx, yy, zz
`

func main_()  {
	data, err := toml.Decode(blob)
	fmt.Println(data)
	fmt.Println(err)

	enc := toml.Encode(data)
	fmt.Println(enc)
}

func main()  {
	data := toml.Data{}
	n := 1000
	for i := 0; i <= n; i++ {
		data[fmt.Sprintf("stanza_%d", i)] = toml.Stanza{}
		for j := 0; j <= n; j++ {
			data[fmt.Sprintf("stanza_%d", i)][fmt.Sprintf("key_%d", j)] = fmt.Sprintf("value-%d-%d", i, j)
		}
	}


	fmt.Printf("encode\n")
	start := time.Now()
	enc := toml.Encode(data)
	ioutil.WriteFile("./data.dat", []byte(enc), 0644)
	fmt.Printf("--> %s\n", time.Since(start))

	fmt.Printf("decode\n")
	start = time.Now()
	bytes, _:= ioutil.ReadFile("./data.dat")
	data, err := toml.Decode(string(bytes))
	fmt.Printf("--> %s\n", time.Since(start))

	if err != nil {
		println("error")
	}
}