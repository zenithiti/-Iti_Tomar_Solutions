package main

import (
	_ "./algorithm"
	"Study/-Iti_Tomar_Solutions/Go/algorithm"
	"fmt"
	"sort"
	"strconv"
	"time"
)

func sequence(number uint64)  {
	fmt.Println("Fibonacci sequence: ")
	for i := uint64(0); i <number; i++ {
		fmt.Print(strconv.FormatUint(algorithm.FibonacciNumber(i),10) + " ")
	}
}
func sumNumber(number uint64)  {
	fmt.Println("Sum of Even Fibonacci Number ", algorithm.SumOfEvenFibonacciNumber(number))
}
func fibonacciSolution(number uint64)  {
 	sequence(number)
	fmt.Println("")
 	sumNumber(number)
	fmt.Println("")
}

func shortedTwoArray(firstArray,secondArray []int) {
	result:=algorithm.IntersectionOfArray(firstArray, secondArray)
	sort.Sort(sort.IntSlice(result))
	fmt.Println(result)
}

func checkOddDigit(number uint64)  {
	evenDigit:=algorithm.CountEvenDigit(number)
	oddDigit:=algorithm.CountOddDigit(number)

	if evenDigit%2==0 && oddDigit%2!=0{
		println("Yes. It is")
	}else {
		println("No. It is not")
	}

}

func sumationOfDecimalDigit(number int)  {
	result:=algorithm.SameNumberSeriesArray(number)
	fmt.Println("Sum of Number::: " +strconv.FormatUint(algorithm.SumOfIntArrayElement(result),10))
}

func main()  {

	var fibonacciNumber uint64=100
	var firstArray,secondeArray =[]int{2, 3, 5,5,18,20,78,99},[]int{1,2,18,3,3,5,5}
	var number uint64
	var decimalNumber int

	/*Start Fibonacci Operation */
	fmt.Println("Fibonacci Operations:::")
	fmt.Println("")
	fibonacciSolution(fibonacciNumber)

	/*End*/

	/*Start Sorted Array*/
	fmt.Println("Sorted Array:::")
	shortedTwoArray(firstArray,secondeArray)

	/*End*/

	/*Start Number Inspection*/
	fmt.Println("")
	fmt.Println("Inspection of The Input Number :::")
	fmt.Println("Enter Big Intger Number :")
	fmt.Scan(&number)
	checkOddDigit(number)
	/*End*/

	time.Sleep(time.Duration(1) * time.Microsecond)

	/*Start Summation of Decimal Digit*/
	fmt.Println("")
	fmt.Println("Summation of Decimal Digit:::")
	fmt.Println("Enter Big Intger Number :")
	fmt.Scan(&decimalNumber)
	sumationOfDecimalDigit(decimalNumber)
	/*End*/

}