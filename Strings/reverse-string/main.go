package main

import "fmt"

func main() {
	s := "hello"

	ch := []byte(s)

	for left, right := 0, len(ch)-1; left < right; left, right = left+1, right-1 {
		ch[left], ch[right] = ch[right], ch[left]
	}

	fmt.Println(string(ch))
}
