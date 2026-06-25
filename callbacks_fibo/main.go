package main

import "fmt"

// factorial -  функція, яка розраховує факторіал цілого числа
func factorial(n int) int {
    if n == 1 {  return 1  }
    return n * factorial(n-1)
}

// fibo -  функція, яка “видає” послідовність Фібоначчі через колбек yield.
func fibo(yield func(x int) bool) {
	f0, f1 := 0, 1
	for yield(f0) {
		f0, f1 = f1, f0+f1
	}
}

func main() {
	
	fibo(func(x int) bool {
		if x >= 1000 {
			return false
		}
		fmt.Printf("%d ", x)
		return true
	})

	fmt.Println("\n\n", factorial(5))
}