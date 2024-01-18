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
 * HackerRank URL: https://www.hackerrank.com/challenges/leonardo-and-prime
 *
/

/*
 * Complete the 'primeCount' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts LONG_INTEGER n as parameter.
*/

var primes = []int64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47}

func primeCount(n int64) int64 {

	if n == 1 {
		return 0
	}

	var count int64 = 0
	var res int64 = 1
	var i int64

	for i = 2; i <= n; i++ {
		if isPrime(i) {
			res *= i
			count++
			if res > int64(n) {
				return count - 1
			}
		}
	}

	return 1
}

func primeCount1(n int64) int64 {

	if n == 1 {
		return 0
	}

	var count int64 = 0
	var res int64 = 1

	for (res < n) && (count < 15) {
		res *= primes[count]
		count++
	}

	if n < 2 {
		count = 0
	}

	if res > n {
		count = count - 1
	}

	return count
}

func primeCount2(n int64) int64 {

	if n == 1 {
		return 0
	}

	var count int64 = 0
	var res int64 = 1
	var i int64

	for i = 0; res < n; i++ {
		count = i
		res *= primes[i]
	}

	if res == n {
		count = count + 1
	}
	return count
}

var pprods = make([]int64, 16)

func primeCount3(n int64) int64 {
	var count int32 = 0
	var i int32

	if n == 1 {
		return 0
	}

	if pprods[1] == 0 {

		// calc prime prods so ith element is product of 1st i primes

		pprods[1] = 2
		for i = 2; i <= 15; i++ {
			pprods[i] = pprods[i-1] * primes[i-1]
		}
	}

	for (pprods[count] < n) && (count < 15) {
		count++
	}

	if n < 2 {
		count = 0
	}

	if pprods[count] > n {
		count = count - 1
	}

	return int64(count)
}

func isPrime(i int64) bool {
	for _, prime := range primes {
		if prime == i {
			return true
		}
	}

	return false
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create("output.txt")
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		n, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)

		result := primeCount(n)

		fmt.Fprintf(writer, "%d\n", result)
	}

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
