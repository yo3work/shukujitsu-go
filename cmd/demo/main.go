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
		fmt.Printf("%d年%2d月%2d日 (%s) = %s\n", e.Year, e.Month, e.Day, e.YMD, e.Name)
	}
}
