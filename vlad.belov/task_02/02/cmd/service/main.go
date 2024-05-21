package main

import (
	"fmt"
	"os"
	"sort"
)

const (
	SUCCESS = iota
	INVALID_NUM_INPUT
	OUT_OF_RANGE
)

func checkInputErr(err error) {
	if err != nil {
		panic("invalid input")
	}
}

func main() {
	var n uint16
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Wrong number of meals input")
		os.Exit(INVALID_NUM_INPUT)
	}

	meals := make([]int16, n)
	for i := uint16(0); i < n; i++ {
		_, err = fmt.Scanf("%d", &meals[i])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Wrong meals input")
			os.Exit(INVALID_NUM_INPUT)
		}
	}

	var k uint16
	_, err = fmt.Scanln(&k)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Wrong priority input")
		os.Exit(INVALID_NUM_INPUT)
	}

	if k > uint16(len(meals)) || k == 0 {
		if err != nil {
			fmt.Fprintln(os.Stderr, "Out of range priority")
			os.Exit(OUT_OF_RANGE)
		}
	}

	sort.SliceStable(meals, func(i, j int) bool {
		return meals[i] > meals[j]
	})
	fmt.Println(meals[k-1])
}
