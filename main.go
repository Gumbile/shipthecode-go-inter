package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * 3.14
}

type Square struct {
	Side float64
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}
func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	kind := sc.Text()
	sc.Scan()
	dim, _ := strconv.ParseFloat(sc.Text(), 64)
	var s interface{ Area() float64 }
	switch kind {
	case "circle":
		s = Circle{Radius: dim}
	case "square":
		s = Square{Side: dim}
	default:
		os.Exit(-1)
	}
	if s != nil {
		fmt.Printf("%.2f\n", s.Area())
	}
}
