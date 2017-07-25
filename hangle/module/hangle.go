package module

import "fmt"

var Engkey = []string{"Q", "W", "E", "R", "T", "Y", "U", "I", "O", "P", "A", "S", "D", "F", "G", "H", "J", "K", "L", "Z", "X", "C", "V", "B", "N", "M"}
var Hankey = []string{"ㅂ", "ㅈ", "ㄷ", "ㄱ", "ㅅ", "ㅛ", "ㅕ", "ㅑ", "ㅐ", "ㅔ", "ㅁ", "ㄴ", "ㅇ", "ㄹ", "ㅎ", "ㅗ", "ㅓ", "ㅏ", "ㅣ", "ㅋ", "ㅌ", "ㅊ", "ㅍ", "ㅠ", "ㅜ", "ㅡ"}

func init() {

}

func GetEngCharIndex(c string) int {
	return indexOf(Engkey, c)
}

func GetKorCharIndex(c string) int {
	return indexOf(Hankey, c)
}

func indexOf(s []string, t string) int {
	for i, c := range s {
		a := fmt.Sprintf("%x", c)
		b := fmt.Sprintf("%x", t)

		if a == b {
			return i
		}
	}
	return -1
}

func GetChar(s string) string {
	i := GetKorCharIndex(s)

	if i > -1 {
		return Engkey[i]
	}

	return ""
}
