package main

import (
	"fmt"
	"time"
)

const dobStr = "21/09/1999"

func main() {
	dob, err := time.Parse("02/01/2006", dobStr)
	if err != nil {
		fmt.Println(err)
	}

	now := time.Now()

	age := now.Year() - dob.Year()
	if now.Month() < dob.Month() || (now.Month() == dob.Month() && now.Day() < dob.Day()) {
		age--
	}
	fmt.Println(age)
}
