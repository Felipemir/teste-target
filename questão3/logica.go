//a) 1, 3, 5, 7, ___
//b) 2, 4, 8, 16, 32, 64, ____
//c) 0, 1, 4, 9, 16, 25, 36, ____
//d) 4, 16, 36, 64, ____
//e) 1, 1, 2, 3, 5, 8, ____
//f) 2,10, 12, 16, 17, 18, 19, ____

package main

import (
	"fmt"
	"math"
)

func main() {
	//a
	fmt.Println(" 1, 3, 5, 7,... o proximo numero é :", 7+2)
	// a sequencia é de numeros impares.

	//b
	fmt.Println("2, 4, 8, 16, 32, 64... o proximo numero é :", 64*2)
	// a sequencia é o dobro do numero anterior.

	//c
	fmt.Println(" 0, 1, 4, 9, 16, 25, 36... o proximo numero é :", int(math.Pow(7, 2)))
	// a sequencia é o quadrado dos numeros inteiros.

	//d
	fmt.Println(" 4, 16, 36, 64... o proximo numero é :", int(math.Pow(10, 2)))
	// a sequencia é o quadrado dos numeros pares.

	//e
	fib := []int{1, 1, 2, 3, 5, 8}

	fmt.Println(" 1, 1, 2, 3, 5, 8... o proximo numero é :", fib[len(fib)-1]+fib[len(fib)-2])
	// a sequencia é a soma dos dois numeros anteriores, uma sequencia de fibonacci.

	//f
	fmt.Println(" 2,10, 12, 16, 17, 18, 19... o proximo numero é :", 19+1)
	// a sequencia aqui, é que todos os numeros começam com a letra D,dois, dez, doze, dezesseis, dezessete, dezoito, dezenove.. logo o proximo numero é vinte.
}
