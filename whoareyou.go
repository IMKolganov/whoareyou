package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	startLog()

	var birthYear int
	fmt.Println("Who are you?")
	fmt.Print("Enter your birth year: ")
	_, err := fmt.Scanln(&birthYear)
	if err != nil {
		log.Fatalf("Failed to read birth year: %v", err)
	}

	fmt.Println(whichYear(birthYear))
}

func whichYear(birthOfYear int) string {
	switch {
	case birthOfYear >= 1946 && birthOfYear <= 1964:
		return "Привет, бумер!"

	case birthOfYear >= 1965 && birthOfYear <= 1980:
		return "«Привет, представитель X!»"
	case birthOfYear >= 1981 && birthOfYear <= 1996:
		return "«Привет, миллениал!»"
	case birthOfYear >= 1997 && birthOfYear <= 2012:
		return "«Привет, зумер!»"
	case birthOfYear >= 2013 && birthOfYear <= 3000:
		return "«Привет, альфа!»"
	}
	return "I don't know who are you..."
}

func startLog() {
	file, err := os.OpenFile("error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer file.Close()
}
