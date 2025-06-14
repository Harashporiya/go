package main

import (
	"fmt"
	"github.com/Harashporiya/golang/auth"
)

func main() {
	auth.LoginWithCredential("Harash", "123456")
	session := auth.GetString()
	fmt.Println("session", session)
}
