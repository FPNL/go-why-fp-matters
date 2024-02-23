package main

import "reflect"

func reduce[T any](a any, i any) func(*_cons[T]) reflect.Value {
	if reflect.TypeOf(a).Kind() != reflect.Func {
		panic("protected item is not a function")
	}
	return func(b *_cons[T]) reflect.Value {
		if b == nil {
			return reflect.ValueOf(i)
		}

		g := []reflect.Value{
			reflect.ValueOf(b.a),
			reduce[T](a, i)(b.b), // 我不知道他會返回哪些型別
		}
		return reflect.ValueOf(a).Call(g)[0]
	}
}
