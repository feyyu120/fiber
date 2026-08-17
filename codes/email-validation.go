package main

import (
	"fmt"
	"regexp"
)

func main() {
	regex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	email := "feyselfeyyu@gmail.com"
	fmt.Println(regex.MatchString(email))
}
