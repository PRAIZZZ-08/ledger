package main

import  (
	"fmt"
	"errors"
)

func validateUser(username string) (bool, error) {
	if username == "" {
		return false, errors.New("username cannot be empty")
	}
	return true, nil
}


func main() {
	namecheck, err := validateUser("")

	if err != nil {
		fmt.Println("Error found:", err)
		return
	}

	fmt.Println("Status:", namecheck)
}
