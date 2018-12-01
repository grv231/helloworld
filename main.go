package main

import "fmt"

func do() {
	m := map[string]string{"aaa": "AAA", "bbb": "BBB", "ccc": "CCC"}
	for k, v := range m {
		_ = fmt.Sprintf("k: %v, v: %v", k, v)
	}
}

// Hello returns a greeting
func Hello() string {
	do()
	return "Hello, world"
}

func main() {
	do()
	fmt.Println(Hello())
}
