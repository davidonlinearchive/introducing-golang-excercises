// GOAL: Write an app that converts feet to meters. 1 ft = 0.3048 meters

package exercises

import "fmt"

func ConvertToMeters() {
	var ft float32
	var meters float32

	fmt.Print("Enter height in feet: ")
	fmt.Scanln(&ft)

	meters = ft * 0.3048

	fmt.Printf("Your height in meters is %.4f\n", meters)
}
