package main

import "fmt"

const kelvinBoilingTemp = 373

func main() {
	kelvinTemp := kelvinBoilingTemp
	celsiusTemp := kelvinTemp - 273

	fmt.Printf("A temperatura de ebulição dá água em Kelvin é %dK e em Celsius é %d°C\n", kelvinTemp, celsiusTemp)
}

