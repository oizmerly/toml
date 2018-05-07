package main

import (
	"./toml"
	"fmt"
	"time"
)

func main()  {
	data := toml.Data{}
	n := 10
	for i := 0; i <= n; i++ {
		for j := 0; j <= n; j++ {
			data.SetValue(
				toml.Title(fmt.Sprintf("stanza_%d", i)),
				toml.Key(fmt.Sprintf("key_%d", j)),
				toml.Value(fmt.Sprintf("value-%d-%d", i, j)))
		}
	}

	fmt.Printf("encode\n")
	start := time.Now()
	data.Write("./data.dat")
	fmt.Printf("--> %s\n", time.Since(start))

	fmt.Printf("decode\n")
	start = time.Now()
	data = toml.Data{}
	data.Read("./data.dat")
	fmt.Printf("--> %s\n", time.Since(start))
}
