package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type LevelType struct {
	row       []int
	max_diff  int
	is_sorted bool
}

func get_diff(a, b *int) int {
	if *a < *b {
		return (*b) - (*a)
	}
	return (*a) - (*b)
}

func max_difference(arr *[]int) *int {
	var temp []int
	for i := 0; i < len(*arr)-1; i++ {
		diff := get_diff(&(*arr)[i+1], &(*arr)[i])
		temp = append(temp, diff)
	}
	sort.Ints(temp)
	slices.Reverse(temp)
	return &temp[0]
}

func calculateDifference(data *map[int]LevelType) {
	for key, val := range *data {
		// fmt.Println(val.row)
		diff := max_difference(&val.row)
		asc, des := chk_for_sorted(&val.row)
		if (*diff <= 3 && *diff >= 0) && (asc || des) {
			(*data)[key] = LevelType{
				row:       val.row,
				is_sorted: true,
				max_diff:  *diff,
			}
		}
		if (*diff <= 3 && *diff >= 0) && (asc == false && des == false) {
			(*data)[key] = LevelType{
				row:       val.row,
				is_sorted: false,
				max_diff:  *diff,
			}
		}
		if *diff > 3 {
			delete(*data, key)
		}
	}
}

// find the defective element in an array ad then simply remove it from the array and then again chk for sort if it is sorted, tehn it is { safe } || { sorted }
func find_defective_element_and_check_for_sort(arr *[]int) bool {
	// check for increasing
	// fmt.Println(*arr)
	for idx := 0; idx < len(*arr); idx++ {
		newArr := append((*arr)[:idx], (*arr)[idx+1:]...)
		asc, dec := chk_for_sorted(&newArr)
		if asc || dec {
			return true
		}
		newArr = []int{}
	}
	// fmt.Println(*arr)
	return false
}

func chk_for_sorted(arr *[]int) (bool, bool) {
	is_ascending := true
	is_descending := true
	for i := 0; i < len(*arr)-1; i++ {
		if (*arr)[i] > (*arr)[i+1] {
			is_ascending = false
		}
		if (*arr)[i] < (*arr)[i+1] {
			is_descending = false
		}
	}
	return is_ascending, is_descending
}

func main() {
	filePath := "./inputs/input.txt"
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()
	var arr [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		splitted := strings.Split(word, " ")
		row := make([]int, 0)
		for _, value := range splitted {

			ele_converted_to_int, err := strconv.Atoi(value)
			if err != nil {
				fmt.Println("Error converting to integer: ", err)
				return
			}
			row = append(row, ele_converted_to_int)
		}
		arr = append(arr, row)
	}

	data := make(map[int]LevelType, 1)
	for idx, row := range arr {
		data[idx] = LevelType{
			row:       row,
			max_diff:  0,
			is_sorted: true,
		}
	}
	// Calculate the max difference in between the level of each levels
	calculateDifference(&data)
	// arr1 := []int{1, 4, 3, 2, 5}
	safe_sum := 0
	for _, value := range data {
		if value.is_sorted || find_defective_element_and_check_for_sort(&value.row) {
			safe_sum += 1
		}
	}
	fmt.Println(safe_sum)
	// Print the data
	// for key, value := range data {
	// 	fmt.Printf("Level %d: %v, max_diff : %d, sorted : %t\n", key, value.row, value.max_diff, value.is_sorted)
	// }
}
