package main

import "fmt"

func main() {
	// Decalre map and add kv
	emails := map[string]string{"Xyz": "xyz@gmail.com", "Xzsda": "xzsda@gmail.com"}

	emails["Lok"] = "lok@gmail.com"
	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Xyz"])

	// Delete from map
	delete(emails, "Xyz")
	fmt.Println(emails)
}
