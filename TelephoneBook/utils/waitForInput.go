package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func WaitForInput(convertToInt bool) (string, int) {
	for {
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil || strings.ReplaceAll(input, " ", "") == "\n" {
			fmt.Println("Не валидные данные. Повторите ввод:")
			continue
		}
		input = input[:len(input)-1]
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
