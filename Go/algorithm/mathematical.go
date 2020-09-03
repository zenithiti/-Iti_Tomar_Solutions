package algorithm

func CountEvenDigit(number uint64) int {
	result:=0
	for number>0{
		rem := number % 10
		if rem % 2 == 0 {
			result++
		}
		number = number / 10
	}
	return result
}
func CountOddDigit(number uint64) int {
	result:=0
	for number>0{
		rem := number % 10
		if rem % 2 != 0 {
			result++
		}
		number = number / 10
	}
	return result
}
func SumOfIntArrayElement(array []int) (result uint64) {
	for i:=0; i<len(array);i++ {
		result +=uint64(array[i])
	}
	return result
}
func SumOfFloat64ArrayElement(array []float64) (result float64) {
	for i:=0; i<len(array);i++ {
		result += array[i]
	}
	return result
}