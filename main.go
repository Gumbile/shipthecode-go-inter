package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func validateAge(s string) (int, error) {
	// implement
	age, err := strconv.Atoi(s)
	if err != nil {
		err = fmt.Errorf("parse: %w", err)
		return 0, err
	}
	if age < 0 {
		negativeValue := errors.New("negative")
		return age, negativeValue

	}
	return age, nil
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	age, err := validateAge(sc.Text())
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
	} else {
		fmt.Printf("age: %d\n", age)
	}
}
