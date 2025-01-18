package util

import "fmt"

func CheckError(message string, err error) {
	fmt.Printf("%s: %v\n", message, err)
}
