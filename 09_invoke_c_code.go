package main

/*
#include <stdlib.h>
#include <stdio.h>
#include <error.h>

void sayHi() {
	printf("hello world!");
}
 */
import "C"

func main() {
	C.sayHi()
}