package main

/*
#include <stdlib.h>
#include "echo.h"
*/
import "C"

import (
	"flag"
	"unsafe"
	"strings"
)

func init() {
	flag.Parse()
}

func main() {
	cs := C.CString(strings.Join(flag.Args(), " "))
	C.echo(cs)
	C.free(unsafe.Pointer(cs))
}