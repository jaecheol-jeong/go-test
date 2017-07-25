package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	str := "가나다 "

	for _, s := range str {
		h := fmt.Sprintf("%x", s)
		println(h)
		src := []byte(h)
		d := hex.EncodeToString(src)
		println(d)
	}

	for i := 0; i < len(str); i++ {
		var x = str[i]
		bt := []byte{x}
		dd := hex.EncodeToString(bt)
		println(i, dd)
	}
}
