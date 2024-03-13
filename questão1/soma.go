package main

import "fmt"

func main() {
	const indice = 13
	soma := 0
	var k = 0

	for k < indice {
		k = k + 1
		soma = soma + k
	}

	fmt.Println(soma)
}
