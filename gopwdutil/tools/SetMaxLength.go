package tools

import (
	"fmt"
	"strconv"
	"strings"
)

func SetMaxLength(minLength int, capacity int) int {
	var isValidInput = false
	maxLength := 0
	
	for !isValidInput {
		fmt.Printf("Set maximum length (%d - %d): ", minLength + 1, capacity)
		line, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Error: reader failure.")
			break
		}

		maxLength, err = strconv.Atoi(strings.TrimSpace(line))
		if err != nil {
			fmt.Println("Error: conversion error.")
			break
		}

		if maxLength < minLength + 1|| maxLength > capacity {
			fmt.Println("Error: invalid number.")
			break
		}

		isValidInput = true
	}

	return maxLength
}
