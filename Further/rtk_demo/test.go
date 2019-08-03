package main

//#cgo CFLAGS: -I./rtklib
//#cgo LDFLAGS: -L${SRCDIR}/rtklib -lrtklib
//#cgo LDFLAGS: -lws2_32 -lwinmm
//#include "rtklib.h"
import "C"
import "fmt"

func main() {
	s1 := C.CString("")
	v := C.gpst2time(2045, 23456.200)
	s2 := C.time2str(v, *s1, 3)
	fmt.Println(C.GoString(s2))
}
