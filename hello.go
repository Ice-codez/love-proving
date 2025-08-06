package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello bebi koo!")

	var bababowe string
	bababowe = "I hate "
	var who, why string = "me", "just because."
	bababowe += who + " " + why
	fmt.Println(bababowe)

	const love = "rhian"
	fmt.Println("but", "I love "+love+" 😁")

	var a, b int = 3, 4
	var c int
	c += a + b
	fmt.Println("how many days are there in a week?")
	fmt.Println("its:", c)

	const h, g, f, d, e = "and", "i'll", "love", "you", "everyday"
	fmt.Println(h, g, f, d, e, "❤️")

	var timeNow = time.Now()
	fmt.Println("and today is:", timeNow.Format("Monday, 01 January 2006"))

	const when = "never."
	fmt.Println("when will i stop loving you?", when)

	numberOne := 12
	numberTwo := 17
	var sum int
	sum = numberOne + numberTwo
	fmt.Println("Anniversary:", sum)
}
