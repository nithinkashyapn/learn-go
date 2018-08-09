package main

import (
	"fmt"
)

func main() {
	numbers := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}

	for i := range numbers {
		fmt.Println("i is ", i, "value is ", numbers[i])
	}

	antonyms := map[string]string{"small": "big", "open": "close", "comfort": "hurt"}

	for word := range antonyms {
		fmt.Printf("Antonym of %v is %v\n", word, antonyms[word])
	}

	isPresent, ok := antonyms["small"]
	isNotPresent, notok := antonyms["sit"]

	fmt.Printf("isNotPresent is '%v' and notok is %v\n", isNotPresent, notok)

	if ok {

		fmt.Printf("isPresent is '%v' and ok is %v\n", isPresent, ok)

		for synonym, antonym := range antonyms {
			fmt.Printf("Antonym of %v is %v\n", synonym, antonym)
		}
	}

}
