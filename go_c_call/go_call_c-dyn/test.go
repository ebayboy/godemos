package main

/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L. -lfoo
#include <stdlib.h>
#include <stdio.h>
#include "foo.h"
*/
import "C"
import "fmt"

func main() {
	C.foo()
	fmt.Println("vim-go")
}
