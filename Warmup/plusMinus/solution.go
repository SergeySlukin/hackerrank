package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"fmt"
	"io"
)

func plusMinus(arr []int32)  {
	denominator := len(arr)
	positive := 0
	negative := 0
	zero := 0
	for _, a := range arr {
		switch {
		case a > 0:
			positive++
			break
		case a < 0:
			negative++
			break
		default:
			zero++
			break
		}
	}

	printResult(positive, denominator)
	printResult(negative, denominator)
	printResult(zero, denominator)
}

func printResult(number int, denominator int) {
	r := float32(number) / float32(denominator)
	fmt.Println(r)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	plusMinus(arr)
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
