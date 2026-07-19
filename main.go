package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"sync"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		nums[i], _ = strconv.Atoi(sc.Text())
	}

	var mu sync.Mutex
	var wg sync.WaitGroup
	total := 0

	// TODO: split `nums` into 4 chunks and launch a goroutine per chunk.
	// Each goroutine should sum its chunk and add the result into `total`,
	// guarded by `mu`. Use `wg` (Add/Done/Wait) to wait for all goroutines
	// to finish before printing.

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

	fmt.Println(total)
}
