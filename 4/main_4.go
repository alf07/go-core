package main

import "fmt"

func difference(first, second []string) []string {
	secondSet := make(map[string]struct{}, len(second))
	for _, value := range second {
		secondSet[value] = struct{}{}
	}

	result := make([]string, 0)
	for _, value := range first {
		if _, exists := secondSet[value]; !exists {
			result = append(result, value)
		}
	}
	return result
}

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}
	fmt.Println(difference(slice1, slice2))
}
