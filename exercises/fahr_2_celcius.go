// GOAL: write a program that converts from Fahrenheit into Celsius (C = (F − 32) * 5/9).
package exercises

import "fmt"

func ConvertToFahr() {
	var fahr float32
	var celcius float32

	fmt.Print("Enter your degree in fahrenheit: ")
	fmt.Scanln(&fahr)

	celcius = (fahr - 32) * 5 / 9

	fmt.Printf("Your degree is %.2f in celcius\n", celcius)
}
