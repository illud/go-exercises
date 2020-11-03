package exercises

import "fmt"

func PrimeNumber(prime int) {
	for i := 2; i < prime; i++ {
		if prime%i == 0 {
			fmt.Println("Is not a prime number")
			break
		}
		fmt.Println("Is a prime number")
		break
	}
}
