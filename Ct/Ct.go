package main

/*

#include <stdio.h>

int doSomething(void) {
	printf("I am a C function in GO!\n");
	return 1;
}
*/
import "C"

import "fmt"

func main() {
	fmt.Println("I will call some C code!")
	myint := C.doSomething()
	fmt.Printf("I have called some C code and got %d!\n",myint)
}
