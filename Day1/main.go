package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func absolute_difference(left, right int) int {
	if left < right {
		return right - left
	}
	return left - right
}

func is_the_number_repeated_in_list(num *int, r_list *[]int, repeated_nums *map[int]int) {
	// to find the number , that how many times it going to repeatedhtat
	occurred := 0
	for _, l_ele := range *r_list {
		if *num != l_ele {
			continue
		}
		occurred++
	}
	(*repeated_nums)[*num] = occurred
}

func update_the_left_list(left *[]int, right *[]int, repeated_nums *map[int]int) {
	for _, num := range *left {
		is_the_number_repeated_in_list(&num, right, repeated_nums)
	}

	sum := 0
	product := 1
	for _, ele := range *left {
		product = ele * (*repeated_nums)[ele]
		fmt.Printf("%v - %v\n", ele, (*repeated_nums)[ele])
		sum += product
	}

	fmt.Printf("\n%v", sum)
}

func main() {

	// Read the file
	filePath := "./inputs/input.txt"

	// file streaming
	source_file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error in opening a input file: ")
		return
	}
	defer source_file.Close()

	scanner := bufio.NewScanner(source_file)
	var data [][]int
	var left, right []int

	// Read the file line by line
	for scanner.Scan() {
		word := scanner.Text()
		words := strings.Fields(word)
		var row []int
		for _, word := range words {
			// convert each string into integer
			to_int, err := strconv.Atoi(word)
			if err != nil {
				fmt.Println("Error converting to integer: ", err)
				break
			}
			row = append(row, to_int)
		}
		data = append(data, row)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input file: ", err)
	}

	for _, left_side := range data {
		left = append(left, left_side[0])
		right = append(right, left_side[1])
	}

	sort.Ints(left)
	sort.Ints(right)
	data = [][]int{}

	var min_sum int

	// finding the difference and then calculated differnce is going to sum up with the min_sum
	for i := 0; i < len(left); i++ {
		diff := absolute_difference(left[i], right[i])
		min_sum += diff
	}
	// part 1 Ans
	// fmt.Println(min_sum)

	// In Part2 we've to find that from right list number x in left list number y,  how many times x repeated in y
	number_repeated_no_of_times := make(map[int]int)
	update_the_left_list(&left, &right, &number_repeated_no_of_times)
	// fmt.Print(sum)

	// part 2 Ans
	// fmt.Println(product)

}
