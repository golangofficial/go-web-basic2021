package main

import ( "fmt"; "runtime"; )

func main() {
	fmt.Print("Go runs on : ")
	switch os := runtime.GOOS; os {
	case "dorwin": fmt.Println("OS X.")
	case "linux": fmt.Println("Linux")
	case "android": fmt.Println("Android")
	default: fmt.Printf("%s.\n", os)
	}
}
