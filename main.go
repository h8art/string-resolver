package main

import (
	"fmt"
	"strconv"
)

import str "strings"

func repeatStr(s string, c int) string {
	return str.Repeat(s, c)
}
func resolveStr(s string) string {
	res := ""
	for i, run := range s {
		_, errCur := strconv.Atoi(string(run))
		if len([]rune(s)) > i+1 && errCur != nil {
			nextRun := string([]rune(s)[i+1])
			nextCount, err := strconv.Atoi(string(nextRun))
			if err == nil {
				res += repeatStr(string(run), nextCount)
			} else {
				if errCur != nil {
					res += string(run)
				}
			}
		}
		if errCur != nil && len([]rune(s))-1 == i {
			res += string(run)
		}
	}
	if res == "" {
		res = "(некорректная строка)"
	}
	return res
}
func main() {
	fmt.Println("* a4bc2d5e => aaaabccddddde", resolveStr("a4bc2d5e"), resolveStr("a4bc2d5e") == "aaaabccddddde")
	fmt.Println("* abcd => abcd", resolveStr("abcd"), resolveStr("abcd") == "abcd")
	fmt.Println("* 45 => (некорректная строка)", resolveStr("45"), resolveStr("45") == "(некорректная строка)")
	//fmt.Println("* qwe\4\5 => qwe45 (*')")
	//fmt.Println("* qwe\45 => qwe44444 (*')")
	//fmt.Println("* qwe\\5 => qwe\\\\ (*)")
}
