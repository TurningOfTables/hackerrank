package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

/*
 * HackerRank URL: https://www.hackerrank.com/challenges/time-conversion
 *
/


/*
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
*/

// timeConversion returns a 24 hour time as HH:MM:SS from a 12 hour time string formatted as HH:MM:SSAM
func timeConversion(s string) string {

	inputTime, err := time.Parse("03:04:05PM", s) // Parse using Go example string https://pkg.go.dev/time#Parse
	if err != nil {
		panic(err)
	}

	return inputTime.Format("15:04:05")
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := timeConversion(s)

	fmt.Fprintf(writer, "%s\n", result)

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
