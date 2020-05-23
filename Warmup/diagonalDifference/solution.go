package main

import (
	"math"
	"bufio"
	"os"
	"strconv"
	"strings"
	"fmt"
	"io"
)

/**
$a = 0;
    $b = 0;
    $j = 0;
    $rowElementsCount = count($arr[0]);
    $jj = $rowElementsCount - 1;
    for ($i = 0; $i < $rowElementsCount; $i++) {
        $a += $arr[$i][$j];
        $b += $arr[$i][$jj];
        $j++;
        $jj--;
    }
 */
func diagonalDifference(arr [][]int32) int32 {

	var a int32 = 0
	var b int32 = 0
	var rowElementsCount = len(arr[0])
	var j = 0
	var jj = rowElementsCount - 1
	for i := 0; i < rowElementsCount; i++ {
		fmt.Println(arr[i][jj])
		a += arr[i][j]
		b += arr[i][jj]
		j++
		jj--
	}

	return int32(math.Abs(float64(a - b)))
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)


	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	var arr [][]int32
	for i := 0; i < int(n); i++ {
		arrRowTemp := strings.Split(strings.TrimRight(readLine(reader)," \t\r\n"), " ")

		var arrRow []int32
		for _, arrRowItem := range arrRowTemp {
			arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
			checkError(err)
			arrItem := int32(arrItemTemp)
			arrRow = append(arrRow, arrItem)
		}

		if len(arrRow) != int(n) {
			panic("Bad input")
		}

		arr = append(arr, arrRow)
	}

	result := diagonalDifference(arr)

	fmt.Println(result)
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