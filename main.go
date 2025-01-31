package main

import "fmt"

type keyFmt struct {
	key string
}

func main() {
	mp := make(map[keyFmt]int)

	key1 := keyFmt {
		key: "1243",
	}

	mp[key1] = 1000

	for key, val := range mp {
		fmt.Println("key: ", key, " and value: ", val)
	}
}