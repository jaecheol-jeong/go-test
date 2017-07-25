package main

import (
	m "hangle/module"
	"os"
)

func main() {
	//print(os.Args[1])
	println(m.GetChar("ㅈ"), m.GetChar("ㅓ"), m.GetChar("ㅇ"))

	for _, s := range os.Args[1] {
		print(m.GetChar(string(s)))
	}
}
