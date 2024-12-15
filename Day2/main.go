package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Level struct {
	row                     []int
	max_Difference_in_Level int
	isSorted                bool
	diff                    []int
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func remove_an_element_from_map(levels *map[int]Level, element_to_remove *int) *map[int]Level {
	delete(*levels, *element_to_remove)
	return levels
}

func calculate_the_diff_of_adjacent_elements_row(levels *map[int]Level) *map[int]Level {
	var diff_array []int
	sum := 0
	for _, val := range *levels {
		for idx := range val.row[:len(val.row)-1] {
			diff_array = append(diff_array, absInt(val.row[idx+1]-val.row[idx]))
		}
		// fmt.Println(diff_array)
		sort.Ints(diff_array)
		if diff_array[len(diff_array)-1] <= 3 {
			sum += 1
		}
		diff_array = []int{}
	}
	// Part-1 Ans :- sum
	// fmt.Println(sum)
	return levels
}

func is_Slices_are_strictly_sorted(level []int) (bool, bool) {
	is_strictly_increasing := true
	is_strictly_decreasing := true
	for idx := 0; idx <= len(level)-2; idx++ {
		if level[idx] >= level[idx+1] {
			is_strictly_increasing = false
		}
		if level[idx] <= level[idx+1] {
			is_strictly_decreasing = false
		}
	}
	return is_strictly_increasing, is_strictly_decreasing
}

func main() {
	filePath := "./inputs/input.txt"
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var data [][]int
	for scanner.Scan() {
		word := scanner.Text()
		splitted := strings.Split(word, " ")

		var row []int
		for _, numStr := range splitted {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Println(err)
				continue
			}
			row = append(row, num)
		}
		data = append(data, row)
	}

	levels := make(map[int]Level, 1)

	for idx, ele := range data {
		is_sorted := false
		is_strcly_inc, is_strcly_dec := is_Slices_are_strictly_sorted(ele)
		if is_strcly_dec {
			is_sorted = true
		}
		if is_strcly_inc {
			is_sorted = true
		}
		levels[idx] = Level{
			row:                     ele,
			max_Difference_in_Level: 0,
			isSorted:                is_sorted,
		}
	}
	for key, value := range levels {
		fmt.Printf("Level %d: %v\n", key+1, value.row)
	}
	// removing the undorted levels
	for key, value := range levels {
		if !value.isSorted {
			remove_an_element_from_map(&levels, &key)
		}
	}

	fmt.Println("After Deleteting")
	// print the value of the levels
	for key, value := range levels {
		fmt.Printf("Level %d: %v\n", key+1, value.row)
	}

	// calculate_the_diff_of_adjacent_elements_row()
	calculate_the_diff_of_adjacent_elements_row(&levels)
}
