package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	people := []Person{
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 30},
		{Name: "Charlie", Age: 20},
	}

	fmt.Println("Before sorting:")
	fmt.Println(people)

	// 使用 sort.Slice 对切片进行排序，按 Age 字段降序排序
	sort.Slice(people, func(i, j int) bool {
		fmt.Println("i ", i, people[i].Age)
		fmt.Println("j ", j, people[j].Age)
		return people[i].Age > people[j].Age
	})

	fmt.Println("After sorting:")
	fmt.Println(people)

}
