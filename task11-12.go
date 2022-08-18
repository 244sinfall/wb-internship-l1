package main

import (
	"fmt"
)

type Set[T comparable] map[T]bool

func createSet[T comparable](from []T) Set[T] {
	outputMap := make(Set[T])
	for _, val := range from {
		outputMap[val] = true
	}
	return outputMap
}

func (s *Set[T]) intersect(with Set[T]) {
	for val := range with {
		if _, ok := (*s)[val]; !ok {
			(*s)[val] = true
		}
	}
}

func (s *Set[T]) values() []T {
	slice := make([]T, 0, len(*s))
	for val := range *s {
		slice = append(slice, val)
	}
	return slice
}

func (s *Set[T]) add(value T) {
	(*s)[value] = true
}

func (s *Set[T]) has(value T) bool {
	_, ok := (*s)[value]
	return ok
}

func (s *Set[T]) remove(value T) bool {
	exist := s.has(value)
	if exist {
		delete(*s, value)
	}
	return exist
}

func task11(seq1 []string, seq2 []string) {
	set1 := createSet(seq1)
	fmt.Println("Первое множество:")
	fmt.Println(set1.values())
	set2 := createSet(seq2)
	fmt.Println("Второе множество:")
	fmt.Println(set2.values())
	set1.intersect(set2)
	fmt.Println("Пересечение 1 и 2 множества:")
	fmt.Println(set1.values())
}
