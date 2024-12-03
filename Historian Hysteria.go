package main

import (
    "fmt"
    "sort"
)

func sumOfDifferences(list1, list2 []int) int {
    // Sort both lists
    sort.Ints(list1)
    sort.Ints(list2)

    sum := 0
    i, j := 0, 0

    // Compare and sum differences
    for i < len(list1) && j < len(list2) {
        diff := list1[i] - list2[j]
        if diff < 0 {
            diff = -diff
        }
        sum += diff
        i++
        j++
    }

    return sum
}

func main() {
    list1 := []int{3, 1, 4, 1, 5}
    list2 := []int{9, 2, 6, 5, 3}

    result := sumOfDifferences(list1, list2)
    fmt.Println("Sum of differences:", result)
}