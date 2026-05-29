package main

import (
	"fmt"
	"math"
)

func CelsiusToFahrenheit(celsius float64) float64 {
	f := (celsius * (float64(9) / float64(5))) + 32
	return Round(f, 2)

}

func FahrenheitToCelsius(fahrenheit float64) float64 {
	c := (fahrenheit - 32) * (float64(5) / float64(9))
	return Round(c, 2)
}

func Round(value float64, decimals int) float64 {
	precision := math.Pow10(decimals)
	return math.Round(value*precision) / precision
}

func main() {
	celsius := 25.0
	fahrenheit := CelsiusToFahrenheit(celsius)
	fmt.Printf("%.2f°C is equal to %.2f°F\n", celsius, fahrenheit)

	fahrenheit = 68.0
	celsius = FahrenheitToCelsius(fahrenheit)
	fmt.Printf("%.2f°F is equal to %.2f°C\n", fahrenheit, celsius)

}
