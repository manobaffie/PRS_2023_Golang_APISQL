package test

type Mytest1 struct {
	a string
}

type Mytest2 struct {
	a string
}

var A string = "test a"

func init() {
	print(A)
	print("test\n")
}

func MyTest() {
	print(A)
	print("mytest\n")
}