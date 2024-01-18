package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'lowestTriangle' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER trianglebase
 *  2. INTEGER area
 */

func lowestTriangle(trianglebase int32, area int32) int32 {
	return int32(math.Ceil((float64(area) / float64(trianglebase) * 2)))
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	trianglebaseTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	trianglebase := int32(trianglebaseTemp)

	areaTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	area := int32(areaTemp)

	height := lowestTriangle(trianglebase, area)

	fmt.Fprintf(writer, "%d\n", height)

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
