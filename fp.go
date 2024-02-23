package main

var sum = reduce[int](add, 0)
var product = reduce[int](multiply, 1)
var alltrue = reduce[bool](and, true)
var anytrue = reduce[bool](or, false)

var add = func(a, b int) int {
	return a + b
}

var multiply = func(a, b int) int {
	return a * b
}

var or = func(a, b bool) bool {
	return a || b
}

var and = func(a, b bool) bool {
	return a && b
}
