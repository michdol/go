package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"time"
	"unicode"
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

	sort_birthdates()

	/*
	arr := _sorted_slice_of_100_ints()
	left := 0
	right := len(arr) - 1
	for _, i := range arr {
		fmt.Println("binarySearch", binarySearch(arr, left, right, i))
	}
	*/

	Permutation([]rune("abc"), printCallback)

	runFiniteStateMachine()

	// TODO: Exercises page 64
}

// ========================================================================

func read_string() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
}

// ========================================================================

func print_pi_up_to_n_digits(n int) {
	pi := fmt.Sprintf("%.*f", n, math.Pi)
	fmt.Println(pi)
}

// ========================================================================

func get_weekday_from_date(d time.Time) time.Weekday {
	return d.Weekday()
}

// ========================================================================

func competitive_programming_publish_date() (time.Time, error) {
	layout := "2006-01-02T15:04:05.000Z"
	str := "2010-08-09T10:00:00.000Z"
	return time.Parse(layout, str) 
}

// ========================================================================

func get_unique_int_slice(input []int) []int {
	unique := make([]int, 0, len(input))
	unique_map := make(map[int]bool)
	for _, val := range input {
		if _, ok := unique_map[val]; !ok {
			unique = append(unique, val)
			unique_map[val] = true
		}
	}
	sort.Ints(unique)
	return unique
}

// ========================================================================

func sort_birthdates() {
	ret := _inputBirthDates()
	sort.Sort(byBirthdate(ret))
	fmt.Println(ret)
	for _, d := range ret {
		fmt.Println(d)
	}
}

func _inputBirthDates() []*Birthdate {
	return []*Birthdate{
		&Birthdate{3, 1, 2010},
		&Birthdate{2, 1, 2010},
		&Birthdate{2, 6, 2010},
		&Birthdate{2, 2, 2010},
		&Birthdate{2, 2, 2009},
	}
}

type Birthdate struct {
	Day 		int
	Month 	int
	Year 		int
}

func (b Birthdate) String() string {
	return fmt.Sprintf("{%d/%d/%d}", b.Day, b.Month, b.Year)
}

// Sort by Month ascending > Day ascending > Age ascending
type byBirthdate []*Birthdate

func (b byBirthdate) Len() int { return len(b) }

func (b byBirthdate) Less(i, j int) bool {
	yearEqual := b[i].Year == b[j].Year
	monthEqual := b[i].Month == b[j].Month
	if yearEqual && monthEqual {
		return b[i].Day < b[j].Day
	} else if yearEqual {
		return b[i].Month < b[j].Month
	}
	return b[i].Year > b[j].Year
}

func (b byBirthdate) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

// ========================================================================

func binarySearch(arr []int, left, right, number int) int {
	var mid int
	cnt := 0
	for left <= right {
		mid = (left + right) / 2
		cnt += 1
		if arr[mid] == number {
			fmt.Println("Binary search comparisons count", cnt)
			return mid
		}
		if arr[mid] > number {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	fmt.Println("Binary search comparisons count", cnt)
	return -1
}

func _sorted_slice_of_100_ints() []int {
	arr := make([]int, 100, 100)
	for i := 0; i < len(arr); i++ {
		arr[i] = i + 1
	}
	return arr
}

// ========================================================================

func Permutation(a []rune, f func([]rune)) {
	permute(a, f, 0)
}

func permute(a []rune, f func([]rune), i int) {
	if i > len(a) {
		f(a)
		return
	}

	permute(a, f, i + 1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		permute(a, f, i + 1)
		a[i], a[j] = a[j], a[i]
	}
}

func printCallback(a []rune) {
	fmt.Println(string(a))
}

// ========================================================================

func runFiniteStateMachine() {
	//input := []rune("a70 ab a b10")
	input := []rune(" a70 a21 asdbe a55 a b10")
	originalInput := []rune(" a70 a21 asdbe a55 a b10")
	fsm := NewFiniteStateMachine()
	length := len(input)
	var isLast bool
	ret := make([][3]int, 0)
	for i, a := range input {
		isLast = i == length - 1
		fsm.apply(a, i, isLast)
		if fsm.finished {
			ret = append(ret, fsm.indexes)
			fsm.reset()
		}
	}
	for _, indexes := range ret {
		for _, index := range indexes {
			input[index] = '*'
		}
	}
	fmt.Println(string(originalInput))
	fmt.Println(string(input))
}

type FiniteStateMachine struct {
	step int
	indexes [3]int
	finished bool
}

func NewFiniteStateMachine() FiniteStateMachine {
	return FiniteStateMachine{
		step: 0,
		indexes: [3]int{-1, -1, -1},
		finished: false,
	}
}

func (s *FiniteStateMachine) apply(a rune, index int, isLast bool) {
	switch s.step {
	case 0:
		s.stateFirstCharacter(a, index, isLast)
	case 1:
		s.stateSecondCharacter(a, index, isLast)
	case 2:
		s.stateThirdCharacter(a, index, isLast)
	case 3:
		s.stateFourthCharacter(a, index, isLast)
		
	}
}

func (s *FiniteStateMachine) stateFirstCharacter(a rune, index int, isLast bool) {
	if isLast {
		s.reset()
	}
	if unicode.IsLetter(a) {
		s.indexes[s.step] = index
		s.step += 1
	} else {
		s.reset()
	}
}

func (s *FiniteStateMachine) stateSecondCharacter(a rune, index int, isLast bool) {
	if isLast {
		s.reset()
	}
	if unicode.IsDigit(a) {
		s.indexes[s.step] = index
		s.step += 1
	} else {
		s.reset()
	}
}

func (s *FiniteStateMachine) stateThirdCharacter(a rune, index int, isLast bool) {
	if unicode.IsDigit(a) {
		s.indexes[s.step] = index
		s.step += 1
		if isLast {
			s.finished = true
		}
	} else {
		s.reset()
	}
}

func (s *FiniteStateMachine) stateFourthCharacter(a rune, index int, isLast bool) {
	if a == ' ' {
		s.finished = true
	}
}

func (s *FiniteStateMachine) reset() {
	s.step = 0
	s.indexes = [3]int{-1, -1, -1}
	s.finished = false
}

// ========================================================================
