package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// Потенциально неэффективная версия
func echoInefficient() string {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	return s
}

// Эффективная версия с использованием strings.Join
func echoEfficient() string {
	return strings.Join(os.Args[1:], " ")
}

func echotime() {
	start := time.Now()
	inefficientResult := echoInefficient()
	elapsedInefficient := time.Since(start)
	fmt.Printf("Inefficient version result: %s\n", inefficientResult)
	fmt.Printf("Inefficient version took: %s\n", elapsedInefficient)

	start = time.Now()
	efficientResult := echoEfficient()
	elapsedEfficient := time.Since(start)
	fmt.Printf("Efficient version result: %s\n", efficientResult)
	fmt.Printf("Efficient version took: %s\n", elapsedEfficient)
}
