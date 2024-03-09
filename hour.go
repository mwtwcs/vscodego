package main

import "fmt"

func hour() {
	var time int
	fmt.Scan(&time) // количество минут hours := time/60 + time%60
	hour := time / 60
	minut := time % 60
	fmt.Printf("%d минут - это %d час и %d минут.\n", time, hour, minut)
}
