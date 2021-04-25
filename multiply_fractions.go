package main
import "fmt"

func main() {
	num1 := 10
	num2 := 16
	result1, result2 := reduce(num1, num2)
	fmt.Println(result1, result2)
}

func reduce(x, y int) (rest1,rest2 int ) {
	  gcd := calculateGCD(x, y);
	  rest1 = x/gcd 
	  rest2 = y/gcd 
	    return 
}

func calculateGCD( x ,y int) int {
	if (y == 0) {
		return x
	}
	return calculateGCD(y, x % y)
}