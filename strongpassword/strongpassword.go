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
 * HackerRank URL: https://www.hackerrank.com/challenges/strong-password
 *
/

/*
 * Complete the 'minimumNumber' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. STRING password
*/

// minimumNumber returns the minimum number of characters required to make a password strong based on this definition of a strong password:
// Must be at least 6 characters long
// Must have at least one digit, lowercase, uppercase and special character.
func minimumNumber(currPassLen int32, password string) int32 {
	var minPassLen = int32(6)
	var requiredTypes = []string{
		"0123456789",
		"abcdefghijklmnopqrstuvwxyz",
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"!@#$%^&*()-+",
	}
	var charsToAdd int32

	for _, reqType := range requiredTypes {
		if !strings.ContainsAny(password, reqType) {
			charsToAdd += 1
		}
	}

	newPassLen := charsToAdd + currPassLen

	if newPassLen < minPassLen {
		newPassLen += (minPassLen - newPassLen)
	}

	return newPassLen - currPassLen

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	password := readLine(reader)

	answer := minimumNumber(n, password)

	fmt.Fprintf(writer, "%d\n", answer)

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
