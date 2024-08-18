package main

import (
	"fmt"
	"math/rand/v2"
	"strconv"
)

func main() {
	repeats := 1000000

	results := make(map[string]int)

	for i := 0; i <= repeats; i++ {
		one := rand.IntN(5)
		two := rand.IntN(10)
		three := rand.IntN(10)
		four := rand.IntN(10)

		s := fmt.Sprintf("%d%d%d%d", one, two, three, four)

		val, ok := results[s]
		// If the key exists
		if ok {
			results[s] = val + 1
		} else {
			results[s] = 1
		}
	}

	fmt.Println("results length:", len(results))
	var avg float64
	max := 0
	min := repeats
	maxVal := ""
	minVal := ""

	for result, occ := range results {
		val, err := strconv.Atoi(result)

		if err != nil {
			fmt.Println(err)
		}

		avg += float64(val * occ)

		if max < occ {
			max = occ
			maxVal = result
		}

		if min > occ {
			min = occ
			minVal = result
		}
	}

	fmt.Println("Most occured: ", maxVal, max)
	fmt.Println("Least occured: ", minVal, min)
	fmt.Println("0000 occured:", results["0000"])
	fmt.Println("0001 occured:", results["0001"])
	fmt.Println("2500 occured:", results["2500"])
	fmt.Println("4999 occured:", results["4999"])
	fmt.Println("Avg: ", avg/float64(repeats))
}
