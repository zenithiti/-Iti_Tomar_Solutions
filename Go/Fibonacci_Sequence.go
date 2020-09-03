package main

import (
	"fmt"
	"strconv"
)

func fibonacciNumber(n uint64) uint64 {
	f := make([]uint64, n+1, n+2)
	if n < 2 {
		f = f[0:2]
	}
	f[0] = 1
	f[1] = 1
	for i := uint64(2); i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}
func sumOfEvenFibonacciNumber(n uint64)  uint64{
	var prev, cur, sum uint64 = 0, 1, 0
	for cur+prev <= n {
		next := prev + cur
		if next%2 == 0 {
			sum += next
		}
		prev = cur
		cur = next
	}
	return sum
}

func main()  {
	var i,n uint64=0,100
	fmt.Println("Fibonacci Series: ")
	for i = 0; i <n; i++ {
		fmt.Print(strconv.FormatUint(fibonacciNumber(i),10) + " ")
	}
	fmt.Println("")
	fmt.Println("Sum of Even Fibonacci Number ", sumOfEvenFibonacciNumber(8))
}