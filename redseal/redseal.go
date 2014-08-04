package main


import "fmt"
import "jf/httpconnect"

func main () {
	fmt.Printf("Test:\n")
	fmt.Printf("%s\n", httpconnect.ConnectTo("www.google.com",true))
}
