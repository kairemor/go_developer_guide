package main

import "fmt"

func main() {
	// map is like hash in go object in js and dict in python
	/*
	* map key are the same type and value also
	 */
	// var color = map[string]string

	// color := make(map[int]string)
	// color[10] = "#f0f"
	// color[11] = "#f0f"

	// delete(color, 10)

	color := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}

	fmt.Println(color)
	printMap(color)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("key :" + color + " value :" + hex)
	}
}

/*
/Difference between map and struct
	mqp key are in same type value also 
	but struct can have different value 
	map is reference type ans struct is value type 
	struct need to know all keys in compile time 