package main

import (
	"fmt"

	Basic "github.com/torikallcode/module-matematika"
)

func main() {
	tambah := Basic.Addition(2, 3)
	fmt.Println(tambah)

	kali := Basic.Multiply(2, 3)
	fmt.Println(kali)

	kurang := Basic.Subtraction(2, 3)
	fmt.Println(kurang)

	bagi := Basic.Division(2, 3)
	fmt.Println(bagi)
}
