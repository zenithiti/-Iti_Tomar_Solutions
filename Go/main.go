package main

import (
	_ "./algorithm"
	"Study/-Iti_Tomar_Solutions/Go/algorithm"
	"fmt"
	"sort"
	"strconv"
	"time"
)

func sequence(number uint64) {
	fmt.Println("Fibonacci sequence: ")
	for i := uint64(0); i < number; i++ {
		result := algorithm.FibonacciNumber(i)
		if result%2 == 0 {
			fmt.Print(strconv.FormatUint(result, 10) + " ")
		}
	}
}
func sumNumber(number uint64) {
	fmt.Println("Sum of Even Fibonacci Number ", algorithm.SumOfEvenFibonacciNumber(number))
}
func fibonacciSolution(number uint64) {
	sequence(number)
	fmt.Println("")
	sumNumber(number)
	fmt.Println("")
}

func shortedTwoArray(firstArray, secondArray []int) {
	result := algorithm.IntersectionOfArray(firstArray, secondArray)
	sort.Sort(sort.IntSlice(result))
	fmt.Println(result)
}

func checkOddDigit(number uint64) {
	evenDigit := algorithm.CountEvenDigit(number)
	oddDigit := algorithm.CountOddDigit(number)

	if evenDigit > 0 && oddDigit == 0 {
		println("Yes. It is")
	} else if evenDigit == 0 && oddDigit != 0 {
		println("No. It is not")
	} else if evenDigit%2 == 0 && oddDigit%2 != 0 {
		println("Yes. It is")
	} else {
		println("No. It is not")
	}

}

func sumationOfDecimalDigit(number int) {
	result := algorithm.SameNumberSeriesArray(number)
	fmt.Println("Sum of Number::: " + strconv.FormatUint(algorithm.SumOfIntArrayElement(result), 10))
}

func main() {

	var fibonacciNumber uint64 //=100
	//var firstArray,secondeArray =[]int{2, 3, 5,5,18,20,78,99},[]int{1,2,18,3,3,5,5}
	var inputElement int
	var firstArraySize int
	var secondArraySize int
	var firstArray []int
	var secondeArray []int
	var number uint64
	var decimalNumber int

	/*Start Fibonacci Operation */
	fmt.Println("Fibonacci Operations:::")
	fmt.Println("Please Enter Number for Fibonacci Sequence: ")
	fmt.Scan(&fibonacciNumber)
	fibonacciSolution(fibonacciNumber)

	/*End*/
	time.Sleep(time.Duration(1) * time.Microsecond)

	/*Start Sorted Array*/
	fmt.Println("Sorted Array:::")
	fmt.Println("Please Enter First Array Size::")
	fmt.Scan(&firstArraySize)
	firstArray = make([]int, firstArraySize)
	for i := 0; i < firstArraySize; i++ {
		fmt.Scan(&inputElement)
		time.Sleep(time.Duration(1) * time.Microsecond)
		firstArray[i] = inputElement
	}
	time.Sleep(time.Duration(1) * time.Microsecond)
	fmt.Println(firstArray)
	fmt.Println("Please Enter Seconde Array Size::")
	time.Sleep(time.Duration(1) * time.Microsecond)
	fmt.Scan(&secondArraySize)
	secondeArray = make([]int, secondArraySize)
	for j := 0; j < secondArraySize; j++ {
		fmt.Scan(&inputElement)
		time.Sleep(time.Duration(1) * time.Microsecond)
		secondeArray[j] = inputElement
	}
	time.Sleep(time.Duration(1) * time.Microsecond)
	fmt.Println(secondeArray)
	fmt.Print("Result:: ")
	shortedTwoArray(firstArray, secondeArray)

	/*End*/

	time.Sleep(time.Duration(1) * time.Microsecond)
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
