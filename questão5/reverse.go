package main

import "fmt"

func main() {
	fmt.Println("Digite uma palavra: ")
	var s string
	fmt.Scan(&s)
	fmt.Println(reverterString(s))
}

func reverterString(s string) string {
	conversor := []rune(s)
	for i, j := 0, len(conversor)-1; i < j; i, j = i+1, j-1 {
		conversor[i], conversor[j] = conversor[j], conversor[i]
	}

	return string(conversor)
}
