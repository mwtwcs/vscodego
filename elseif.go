package main

import "fmt"

func elseif() {
	var year int
	fmt.Scan(&year)
	if year <= 13 {
		fmt.Print("детство")
	} else if year >= 14 && year <=24{
		fmt.Print("молодость")
	} else if year >= 25 && year <=59{
		fmt.Print("зрелость")
	} else if year >= 60{
		fmt.Print("старость")
	}
}
