package utils

import (
	"fmt"
	"strconv"
)

func ConvertStr2Num(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("ConvertStr2Num meets error: ", err.Error())
		return 0
	}
	return num
}
