package main

import (
	"fmt"
)

type stack struct {
	data []interface{}
}

func (s *stack) push(item interface{}) {
	s.data = append(s.data, item)

}

func (s *stack) pop() interface{} {
	if len(s.data) == 0 {
		return nil
	}
	data := s.data[len(s.data)-1]
	s.data = s.data[0 : len(s.data)-1]
	return data
}

func main() {
	s1 := stack{}
	for i := 10; i <= 100; i = i + 10 {
		s1.push(i)
	}

	fmt.Println(s1)
	fmt.Println(s1.pop())
	fmt.Println(s1)

	s2 := stack{}
	s2.push("Hello")
	s2.push("World")
	s2.push("Pranab")
	fmt.Println(s2)
	fmt.Println(s2.pop())
	fmt.Println(s2)

}
