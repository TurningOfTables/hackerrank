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
 * HackerRank URL: https://www.hackerrank.com/challenges/array-left-rotation
 *
/

/*
 * Complete the 'rotateLeft' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER d
 *  2. INTEGER_ARRAY arr
*/

// rotateLeft returns an []int32 where the elements of arr []int32 are rotated left by d int32
// will never have to do more than len(arr) rotations - O(len(arr))
func rotateLeft(d int32, arr []int32) []int32 {
	var rotated = arr
	var rotationCount int = int(d) % len(arr) // minimise the number of rotations required when rotations is more than the number of elements

	for x := 1; x <= rotationCount; x++ {
		rotated = append(rotated[1:], rotated[0])
	}

	return rotated
}

// A less efficient solution for anything where d > len(arr) as it does more rotations than required - O(d)
func rotateLeftSlow(d int32, arr []int32) []int32 {
	var rotated = arr
	var rotationCount int
	var x int32

	for x = 1; x <= d; x++ {
		rotationCount++
		rotated = append(rotated[1:], rotated[0])
	}

	return rotated
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	dTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	d := int32(dTemp)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	result := rotateLeft(d, arr)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, " ")
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()
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
