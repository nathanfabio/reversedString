package main

import (
	"fmt"
)

// h e l l o
// 0 1 2 3 4
// i       j

func main() {

	fmt.Println(reversedString("hello"))
	fmt.Println("-------------")
	fmt.Println(reversedStringPro("hello"))
	fmt.Println("-------------")
	fmt.Println(reversedStringPro2("hello"))

}

func reversedString(s string) string {
	r := []rune(s)
	rReversed:= make([]rune, len(s))

	j:= len(r) - 1
	for i:= 0; i < len(r); i++ {
		rReversed[i] = r[j]
		j--
	}

	return string(rReversed)
}

func reversedStringPro(s string) string {
	r:= []rune(s)

	j:= len(r) - 1
	for i:= 0; i < j; i++ {
		r[i], r[j] = r[j], r[i]
		j--
	}

	return string(r)
}

func reversedStringPro2(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}