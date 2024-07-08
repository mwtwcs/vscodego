package main

import "fmt"

func matches() {
	results := []string{"w", "l", "w", "d", "w", "l", "l", "l", "d", "d", "w", "l", "w", "d"}
	
	var a string
	fmt.Scan(&a)
	switch a {
	case "w":
		results = append(results, "w")
	case "l":
		results = append(results, "l")
	case "d":
		results = append(results, "d")
	default:
		fmt.Println("Invalid input")
		return
	}

	// Пересчитываем общий счет
	points := 0
	for _, result := range results {
		switch result {
		case "w":
			points += 3
		case "l":
			points += 0
		case "d":
			points += 1
		}
	}

	// Вывод всех результатов для проверки
	for _, result := range results {
		fmt.Println(result)
	}

	// Вывод общего счета
	fmt.Println("Total points:", points)
}