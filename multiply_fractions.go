package fractions

import "fmt"

// To be treated as a class, to represent fractions
type Fraction struct {
	Numerator, Denominator int
}

//Takes 2 inputs as params to be converted to it's simplified form 
func multiplyFractions(a, b Fraction){ 
	var u , v Fraction
	u.Numerator, u.Denominator = reduce(a)
	v.Numerator, v.Denominator = reduce(b)
		
	fmt.Println("Input 1 in it's reduced form is" ,u)
	fmt.Println("Input 2 in it's reduced form is" ,v)

	multiply(u , v)
}

//Function to convert each of the fraction on to it's simplified form
func reduce(input Fraction) (res1, res2 int) {
	gcd := computeGCD(input.Numerator,input.Denominator)
	res1 = input.Numerator/ gcd
	res2 = input.Denominator/ gcd
	return
}

//Function to calculate the GCD 
func computeGCD( numerator ,denominator int) int {
	if (denominator == 0) {
		return numerator
	}
	return computeGCD(denominator, numerator % denominator)
}

//Function to multiply the simplified inputs
func multiply( input1 Fraction, input2 Fraction){
	output1 := input1.Numerator *  input2.Numerator
	output2 := input1.Denominator * input2.Denominator
	fmt.Println("Multiplication output after the numbers are brought down to it's simplified form is" , output1, "/" ,output2)
}

