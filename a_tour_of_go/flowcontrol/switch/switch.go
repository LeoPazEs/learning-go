package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Running on ")
	os := runtime.GOOS
	switch os {
	case "darwin":
		fmt.Println("macOS")
	case "linux":
		fmt.Println("the right one :D.")
	case "windows":
		fmt.Println("the wrong one :(.")
	default:
		fmt.Printf("%s. \n", os)
	}
}
