package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	fields := strings.Fields(sc.Text())
	nums := make([]int, n)
	for i, f := range fields {
		nums[i], _ = strconv.Atoi(f)
	}
	// split into 4 chunks, goroutine each, sum total
	var wg sync.WaitGroup
	var total int
	var mu sync.Mutex

	size := len(nums)
	sizeChuck := int(math.Ceil(float64(size) / 4.0))
	chuck := 0
	for i := 0; i < 4; i++ {

		wg.Add(1)
		go func(left, right int) {
			defer wg.Done()
			for left < right && left < size {
				mu.Lock()
				total += nums[left]
				mu.Unlock()
				left++
			}
		}(chuck, chuck+sizeChuck)
		chuck += sizeChuck
	}
	wg.Wait()
	fmt.Println(total) // replace with the real total
}
