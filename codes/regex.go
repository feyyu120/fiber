package main

import (
	"fmt"
	"regexp"
)

func main() {
	regex := regexp.MustCompile(`^[0-9]{9}$`)

	fmt.Println(regex.MatchString("092711861"))
}
