package cat

import "fmt"

type Cat struct {
	name string
	age  int
}

// (cat Cat) 이라면 cat을 복사받아서 넘겨받는다.
func (cat *Cat) SetDetails(name string, age int) {
	cat.name = name
	cat.age = age
	fmt.Println("See Details", cat)
}

func (cat Cat) Name() string {
	return cat.name
}

func (cat Cat) Age() int {
	return cat.age
}
