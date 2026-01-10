package auth

import "fmt"

// func loginWithCredentials(username, password string) {
// 	fmt.Println("Logging user using", username, password)
// }

//to access this function outside this package, we need to capitalize the first letter of the function name
func LoginWithCredentials(username, password string) {
	fmt.Println("Logging user using", username, password)
}