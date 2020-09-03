package algorithm

import (
	"strconv"
)

func SameNumberSeriesArray(number int) (result []int) {
	temp:=""
	for i:=0;i<=number;i++{
		for j:=0;j<=i;j++{
			temp=temp+strconv.Itoa(number)
		}
		element,_:=strconv.Atoi(temp)
		result = append(result,element)
		temp=""
	}
	return result
}
func floatToString(input_num float64) string {
	return strconv.FormatFloat(input_num, 'f', 3, 64)
}
