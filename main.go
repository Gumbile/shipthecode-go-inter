package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func safeDivide(a, b int) (q int, err error) {
	// TODO: use defer + recover to catch a panic from a/b (when b == 0)
	// and set err = fmt.Errorf("divide by zero") instead of crashing.
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("divide by zero")
		}
	}()
	return a / b, nil
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	a, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	b, _ := strconv.Atoi(sc.Text())
	q, err := safeDivide(a, b)
	if err != nil {
		fmt.Printf("error: %s\n", err)
	} else {
		fmt.Printf("result: %d\n", q)
	}
}
