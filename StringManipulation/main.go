package main

import (
	"fmt"
	// "strings"
)

func main() {
	// s := ""

	// not efficient approach. This results in O(n²) time complexity due to repeated memory allocations.
	// for i := 0; i < 100; i++ {
	// 	s += "a"
	// }
	// fmt.Println(s)

	// var sb strings.Builder
	// for i := 0; i < 100; i++ {
	// 	sb.WriteString("a")
	// }
	// fmt.Println(sb.String())
	// strings.Builder reduces memory allocations.
	// It appends to an internal buffer, improving performance.
	// Time Complexity: O(n) instead of O(n²).

	// s:="hello"
	// s=s+s	//Inefficient: Creating new strings repeatedly
	// fmt.Println(s)

	// s:=[]byte("hello")
	// s[2]='o'
	// fmt.Println(string(s))
	// []byte allows direct modification without new allocations.

	//joining the strings
	// words:= []string {"utkarsh","kesharwani"}
	// result:=strings.Join(words, " - ")
	// fmt.Println(result)

	// Convert Between string and []rune for Unicode Safety

	//not correct way
	// s := "नमस्ते"
	// b := []byte(s)
	// b[0] = 's'
	// fmt.Println(string(b))

	//correct way is by rune
	// s := "नमस्ते" //rune handles multi-byte characters properly. Avoids corruption of UTF-8 encoded text.
	// b := []rune(s)
	// b[1] = 's'
	// fmt.Println(string(b))


// Operation	Inefficient Approach	Efficient Alternative
// Concatenation	s += str in a loop	strings.Builder
// Large Modifications	Creating new strings	bytes.Buffer
// Unicode Manipulation	[]byte (corrupts)	[]rune (safe)
// Multiple String Joins	Loop with +	strings.Join

}
