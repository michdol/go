package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"time"
)

func main() {
	// read_string()

	print_pi_up_to_n_digits(5)

	publish_date, _ := competitive_programming_publish_date()
	// tomorrow := time.Now().AddDate(0, 0, 1)
	weekday := get_weekday_from_date(publish_date)
	fmt.Println(weekday)

	integers := []int{5, 2, 5, 2, 1, 4, 7, 2}
	fmt.Println(get_unique_int_slice(integers))
}

func read_string() {
	/*
		
	*/
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
}

func print_pi_up_to_n_digits(n int) {
	pi := fmt.Sprintf("%.*f", n, math.Pi)
	fmt.Println(pi)
}

func get_weekday_from_date(d time.Time) time.Weekday {
	return d.Weekday()
}

func competitive_programming_publish_date() (time.Time, error) {
	layout := "2006-01-02T15:04:05.000Z"
	str := "2010-08-09T10:00:00.000Z"
	return time.Parse(layout, str) 
}

func get_unique_int_slice(input []int) []int {
	unique := make([]int, 0, len(input))
	unique_map := make(map[int]bool)
	for _, val := range input {
		if _, ok := unique_map[val]; !ok {
			unique = append(unique, val)
			unique_map[val] = true
		}
	}
	return unique
}