package utils

import (
	"fmt"
	"strconv"
)

func WaitForInput(convertToInt bool) (string, int) {
	for {
		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println("Не валидные данные. Повторите ввод:")
			continue
		}
		if convertToInt {
			t, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("Не валидные данные. Повторите ввод:")
				continue
			}
			return input, t
		}
		return input, -1
	}
}
