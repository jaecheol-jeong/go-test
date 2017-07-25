package controllers

import "fmt"

func Hab2Html(a int, b int) string {
	c := a + b
	str := fmt.Sprintf("<p>%d + %d = %d</p>", a, b, c)
	return str
}

func DivHtml(a int, b int) string {
	c := float32(a) / float32(b)

	str := fmt.Sprintf("%d / %d = %d", a, b, c)
	return str
}
