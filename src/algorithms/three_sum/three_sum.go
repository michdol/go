package main


import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"time"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		panic("Need file path")
	}
	path := os.Args[1]
	arr := ReadIntsFromFile(path)
	start := time.Now()
	fmt.Printf("ThreeSum > %d\n", ThreeSum(arr))
	fmt.Println("Duration >", time.Since(start))
}

func ReadIntsFromFile(path string) []int {
	fileReader := GetFileReader(path)
	scanner := bufio.NewScanner(fileReader)
	scanner.Split(bufio.ScanWords)
	var result []int
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Printf("Error converting %s, %s", n, err.Error())
		}
		result = append(result, n)
	}
	return result
}

func GetFileReader(path string) io.Reader {
	reader, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return reader
}

func ThreeSum(arr []int) int {
	length := len(arr)
	count := 0
	for i := 0; i < length; i ++ {
		for j := i + 1; j < length; j ++ {
			for k := j + 1; k < length; k ++ {
				if arr[i] + arr[j] + arr[k] == 0 {
					count += 1
				}
			}
		}
	}
	return count
}
