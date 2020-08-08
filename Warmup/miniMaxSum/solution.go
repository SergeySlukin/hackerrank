package main

import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"io"
	"fmt"
)

func miniMaxSum(arr []int32) {

	min := int64(0)
	max := int64(0)
	for k, v := range arr {
		sum := int64(0)
		for _, lv := range arr {
			if v == lv {
				continue
			}
			sum += int64(lv)
		}
		if k == 0 {
			min = sum
			max = sum
		}
		if sum < min {
			min = sum
		}
		if sum > max {
			max = sum
		}
	}

	if min == 0 && max == 0 {
		sum := int64(arr[0] * 4)
		min = sum
		max = sum
	}

	fmt.Println(min, max)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < 5; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	miniMaxSum(arr)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
