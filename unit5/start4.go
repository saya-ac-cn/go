package main

func fa(a int) func(i int) int{
	return func(i int) int {
		println(&a,a)
		a = a + 1
		return a
	}
}

func main() {
	f := fa(1)
	g := fa(10)

	println(f(1))
	println(f(1))

	println(g(1))
	println(g(1))
}
