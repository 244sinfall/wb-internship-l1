package main

import "fmt"

func task12(args ...string) {
	set := createSet(args)
	fmt.Println("Получено множество:")
	fmt.Println(set.values())
}
