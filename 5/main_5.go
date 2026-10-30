package main

import "fmt"

func intersection(first, second []int) (bool, []int) {
	firstSet := make(map[int]struct{}, len(first))
	for _, value := range first {
		firstSet[value] = struct{}{}
	}

	result := make([]int, 0)
	seen := make(map[int]struct{})
	for _, value := range second {
		if _, exists := firstSet[value]; !exists {
			continue
		}
		if _, alreadyAdded := seen[value]; alreadyAdded {
			continue
		}
		seen[value] = struct{}{}
		result = append(result, value)
	}

	return len(result) > 0, result
}

func main() {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}

	exists, values := intersection(a, b)
	fmt.Println(exists, values)
}
