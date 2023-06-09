package main

import "fmt"

const (
	tempK = 373.0 // The boiling point of water on kelvin scale
)

func main() {
	tempC := tempK - 273 // Convert from Kelvin scale to Celsius scale

	fmt.Printf("In Celsius scale, the boiling point is %v°C.", tempC)
}
