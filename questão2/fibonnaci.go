package main

import (
	"fmt"
	"math/rand"
	"time"
)

func fibonacci(n int) []int {
	fib := make([]int, n)
	fib[0], fib[1] = 0, 1

	for i := 2; i < n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}

	return fib
}

func main() {
	//n := 11
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(15) + 1
	fib := fibonacci(n)
	fmt.Println(fib[:n-1], "...")

	var aposta int

	for {
		fmt.Println("qual o ultimo numero da sequencia de fibonacci?")
		_, erro := fmt.Scan(&aposta)
		if erro == nil {
			break
		}

		fmt.Println("por favor, insira um numero valido.!")
		fmt.Scanln()
	}

	if aposta == fib[n-1] {
		fmt.Println("Parabens, você acertou!")
	} else {
		fmt.Println("Você errou, o numero correto é", fib[n-1])
	}
}
