package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * HackerRank URL: https://www.hackerrank.com/challenges/mini-max-sum
 *
/

/*
 * Complete the 'miniMaxSum' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
*/

// miniMaxSum returns the minimum and maximum values that can be calculated by summing exactly 4 of the integers
func miniMaxSum(arr []int32) (int64, int64) {

	var min int64
	var max int64

	sortedArr := sortAsc(arr)

	for _, v := range sortedArr[:4] {
		min += int64(v)
	}

	for _, v := range sortedArr[1:] {
		max += int64(v)
	}

	fmt.Println(min, max)
	return min, max
}

func sortAsc(arr []int32) []int32 {
	changeReq := true

	for changeReq {
		changeReq = false
		for x := 0; x < len(arr)-1; x++ {
			if arr[x] > arr[x+1] {
				arr[x], arr[x+1] = arr[x+1], arr[x]
				changeReq = true
			}
		}
	}

	return arr
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

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
