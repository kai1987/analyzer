package main

import "sort"

type IntMap map[int]int

func (im IntMap) ValueSum() int {
	c := 0
	for _, v := range im {
		c += v
	}
	return c
}

func (im IntMap) SortedKeys() []int {
	keys := make([]int, 0, len(im))
	for k, _ := range im {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}
