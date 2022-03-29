package main

import "fmt"

func main() {
	var a [7]string
	a[0] = "Operating System List"
	a[1] = "Windows XP"
	a[2] = "Linux 1.0"
	a[3] = "Raspbian Teensy"
	a[4] = "MacOS"
	a[5] = "iOS"
	a[6] = "Google Android"

	for _, v := range a {
		fmt.Println(len(v))
	}
}
