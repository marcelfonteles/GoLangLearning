package baseTwo

import (
	"strconv"
)

func ToBaseTwo(number int) (int) {
	var quotient, remainder int
	var sliceNumber []int

	quotient = 1
	for quotient > 0 {
		remainder = number % 2
		quotient  = number / 2
		sliceNumber = append(sliceNumber, remainder)
		number = quotient
	}
	var stringNumber string
	for index := len(sliceNumber)-1; index >= 0; index--{
		stringNumber += strconv.Itoa(sliceNumber[index])
		//Problema na conversão de int para string
	}
	finalNumber, _ := strconv.Atoi(stringNumber)

	return finalNumber
}


/*
10 / 2 = 5 (0)
05 / 2 = 2 (1)
02 / 2 = 1 (0)
01 / 2 = 0 (1)
*/