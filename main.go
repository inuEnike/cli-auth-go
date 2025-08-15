package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello struct")
	login()
}

type auth struct {
	name           string
	email          string
	phoneNumber    string
	password       string
	repeatPassword string
}

func login() {
	var authentication auth

	fmt.Println("Enter your name:")
	fmt.Scan(&authentication.name)

	for {
		fmt.Println("Enter your email:")
		fmt.Scan(&authentication.email)

		switch {
		case authentication.email == "":
			fmt.Println("Email cannot be empty")
			continue
		case !strings.Contains(authentication.email, "@"):
			fmt.Println("Email must contain @")
			continue
		case !strings.Contains(authentication.email, "."):
			fmt.Println("Email must contain .")
			continue
		case strings.Count(authentication.email, "@") > 1:
			fmt.Println("Email cannot contain more than one @")
			continue
		case strings.Count(authentication.email, ".") > 1:
			fmt.Println("Email cannot contain more than one .")
			continue
		}

		break // valid email, exit loop
	}

	fmt.Println("Enter your phone number:")
	fmt.Scan(&authentication.phoneNumber)

	fmt.Println("Enter your password:")
	fmt.Scan(&authentication.password)

	fmt.Println("Enter your repeat password:")
	fmt.Scan(&authentication.repeatPassword)

	for authentication.repeatPassword != authentication.password {
		fmt.Println("Passwords do not match. Try again:")
		fmt.Scan(&authentication.repeatPassword)
	}

	fmt.Println("Registration successful!", authentication.name)
}
