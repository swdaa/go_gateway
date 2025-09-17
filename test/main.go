package test

import "time"

func main() {
	// This is just a placeholder to make the package "test" valid.
	var userMap = make(map[int]string, 10)
	var aaa map[int]string = map[int]string{

		1: "a",
		2: "b",
	}
	userMap[1] = "hello"
	userMap[2] = "world"
	println(userMap[1], userMap[2])
	println(aaa[1], aaa[2])
	awaitAdd(3)(1, 2, 3, 4, 5)
	add1 := awaitAdd(2)
	add1(1, 2, 3, 4, 5)
}

func awaitAdd(awaitSecond int) func(...int) int {
	time.Sleep(time.Duration(awaitSecond) * time.Second)
	return func(number ...int) int {
		var sum int
		for _, v := range number {
			sum += v
		}
		println("sum:", sum)
		println("await second:", awaitSecond)
		return sum
	}
}

func add(number ...int) {
	sum := 0
	for _, v := range number {
		sum += v
	}
}
