package main

import "fmt"

func main() {
	// initialize variables here
	var username string
	var smsSendingLimit int
	var costPerSMS float64
	var hasPermission bool

	fmt.Printf("%v %f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)
}
