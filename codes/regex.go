package main

import (
	"fmt"
	"regexp"
)

func main() {
	text := "Contact us at support@google.com or 0919626703"
	emailRegex := regexp.MustCompile(`[a-zA-Z0-0]+@[a-zA-Z]+\.[a-zA-Z]{2,}`)
	//phoneRegex := regexp.MustCompile(`^[0-9]$`)
	email := emailRegex.FindString(text)
	//phone := phoneRegex.FindString(text)
	fmt.Println("email is:", email)
	//fmt.Println("phone is :", phone)
}
