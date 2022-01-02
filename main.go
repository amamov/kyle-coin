package main

import (
	"fmt"

	"github.com/amamov/kyle-coin/cat"
)

func plus(a, b int, name string) (int, string) {
	return a + b, name
}

func hello(x ...int) int {
	var total int
	for idx, item := range x {
		fmt.Println(idx)
		total += item
	}
	return total
}

//* structs
type person struct {
	name string
	age  int
}

// struct method
func (p person) sayHello() {
	fmt.Printf("Hello! my name is %s and I'm %d", p.name, p.age)
}

func main() {
	fmt.Println("Hello Kyle Coin!")

	var name string = "yoon sang seok"
	name = "amamov"
	age := 24
	age = 25

	const job string = "programmer"
	fmt.Println(name, age, job)

	result, name := plus(1, 2, "kyle")

	fmt.Println(result)

	hello(1, 2, 3, 4, 5, 6)

	for index, letter := range job {
		// fmt.Println(index, letter)
		fmt.Println(index, string(letter))
		fmt.Printf("%x\n", letter)
	}

	//* array
	// foods := [3]string{"pasta", "pizza", "pasta"}
	foods := []string{"pasta", "pizza", "pasta"}

	for _, food := range foods {
		fmt.Println(food)
	}

	for i := 0; i < len(foods); i++ {
		fmt.Println(foods[i])

	}
	foods = append(foods, "tomato")
	fmt.Printf("%v\n", foods)

	//* pointer
	a := 17
	b := &a
	fmt.Println(&a, b) // 0x14000124040 0x14000124040
	// "*" : 메모리에 해당하는 값을 꺼낸다.
	fmt.Println(*&a, *b) // 17 17
	a = 18
	fmt.Println(*&a, *b) // 18 18
	*b = 19
	fmt.Println(*&a, *b) // 19 19

	kyle := person{name: "yoon", age: 24} // person{"yoon", 24}
	fmt.Println(kyle.name)
	fmt.Println(kyle.age)
	kyle.sayHello()

	yss := cat.Cat{}
	yss.SetDetails("kyle", 12)
	fmt.Println(yss)
}
