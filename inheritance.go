/*
	Demo to show how inheritance works
	https://www.youtube.com/watch?v=gRpUfjTwSOo
*/
package main

import "fmt"

type user struct {
	name string
	email string
}

type admin struct {
	user
	level string
}

func main() {
	ad := admin{
		user: user{
			name: "tim",
			email: "tim@example.com",
		},
		level: "test",
	}
	fmt.Printf(ad.name)
	fmt.Printf(ad.user.name)
}