package test

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

}
