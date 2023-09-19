package main

import (
	"fmt"
	"time"

	_ "ent-orm-test/ent/runtime"
)

type MySet map[any]struct{}

func duplicateRemoving[T any](s []T) []T {
	res := make([]T, 0, len(s))
	mySet := make(MySet)
	for _, t := range s {
		if _, ok := mySet[t]; !ok {
			res = append(res, t)
			mySet[t] = struct{}{}
		}
	}
	return res
}

func main() {
	fmt.Println(time.Now())
}
