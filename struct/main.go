package main

import "fmt"

type profile struct {
	phoneNumber int
	name string
}

type messageToSend struct {
	profile
	message string
}

func main() {
	mesa := messageToSend{
		profile: profile{
			phoneNumber: 123456789,
			name: "naol",
		},
		message: "Hello",
	}

	fmt.Println(mesa.phoneNumber)

}