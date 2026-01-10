package main

import (
	"fmt"

	"github.com/HaniRawat/Go/packages/user"

	"github.com/HaniRawat/Go/packages/auth"
	"github.com/fatih/color"
)

// to make a package, we need to type "go mod init <github-repo-path>" in the terminal. which creates a go.mod file
func main() {
	auth.LoginWithCredentials("hani rawat", "1234")

	session := auth.GetSession()
	fmt.Println("Session Status:", session)

	user := user.User{
		Email: "hanirawat1122@gmail.com",
		// Name : "Hani Rawat",
	}

	// fmt.Println(user.Email, user.Name)
	color.Blue(user.Email)
}
