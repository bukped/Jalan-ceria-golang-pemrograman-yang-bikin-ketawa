package main

import (
	"fmt"
	"runtime"
)

func print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Print1n((i + 1), message)
	}
}
func main() {
	runtime.GDHAXPROCS(2)
	go print(5, "halo")
	print(5, "apa kabar")
	var input string
	fmt.Scan1n(&input)
}
