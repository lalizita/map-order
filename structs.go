package main

import "sort"

type MyMap struct {
	Key   int
	Value string
}

type MyMapSlice []MyMap

func (a MyMapSlice) Len() int           { return len(a) }
func (a MyMapSlice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a MyMapSlice) Less(i, j int) bool { return a[i].Key > a[j].Key }

func sortPeople(names []string, heights []int) []string {
	var sorted MyMapSlice

	for k, v := range heights {
		sorted = append(sorted, MyMap{Key: v, Value: names[k]})
	}

	sort.Sort(MyMapSlice(sorted))

	descOrder := []string{}
	for _, v := range sorted {
		descOrder = append(descOrder, v.Value)
	}

	return descOrder
}
