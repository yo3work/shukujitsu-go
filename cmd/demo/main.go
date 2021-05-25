package main

import (
	"fmt"

	"github.com/yo3work/shukujitsu-go"
)

// import 省略...
func main() {
	entries, err := shukujitsu.AllEntries()
	if err != nil {
		panic(err)
	}
	for _, e := range entries {
		fmt.Printf("%s = %s\n", e.YMD, e.Name)
	}
}
