package main

import (
	"fmt"
	"os"
)

const (
	SUCCESS = iota
	WRONG_NUM_DEPS_INPUT
	WRONG_NUM_PEOPLE_INPUT
	WRONG_SIGN_INPUT
)

const (
	NOT_LESS string = ">="
	NOT_MORE string = "<="
)

type range_t struct {
	min uint16
	max uint16
}

func main() {
	var n, k, query uint16
	var sign string

	_, err := fmt.Scanln(&n)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Wrong input for departments number")
		os.Exit(WRONG_NUM_DEPS_INPUT)
	}

	for i := uint16(0); i < n; i++ {
		temperature := range_t{15, 30}

		_, err = fmt.Scanln(&k)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Wrong input for people number")
			os.Exit(WRONG_NUM_PEOPLE_INPUT)
		}

		for j := uint16(0); j < k; j++ {
			_, err = fmt.Scan(&sign)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Wrong input for sign of query")
				os.Exit(WRONG_SIGN_INPUT)
			}

			_, err = fmt.Scan(&query)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Wrong input for temperature query")
				os.Exit(WRONG_SIGN_INPUT)
			}

			if sign == NOT_LESS {
				if temperature.min < query {
					temperature.min = query
				}
			} else if sign == NOT_MORE {
				if temperature.max > query {
					temperature.max = query
				}
			} else {
				fmt.Fprintln(os.Stderr, "Wrong sign")
				os.Exit(WRONG_SIGN_INPUT)
			}

			if temperature.min <= temperature.max && temperature.min >= 15 && temperature.max <= 30 {
				fmt.Println(temperature.min)
			} else {
				fmt.Println(-1)
			}
		}
	}
}
