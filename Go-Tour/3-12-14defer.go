package main

import("fmt")

func main() {
	defer fmt.Println("world")
	fmt.Println("hell0")
}
