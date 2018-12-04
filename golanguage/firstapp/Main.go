package main

import "fmt"

func main() {
	statePopulation := map[string]int{
		"california": 12343545,
		"texas":      2232345,
		"florida":    23456787,
		"newyork":    234569876,
	}
	fmt.Printf("%v", statePopulation)
}
