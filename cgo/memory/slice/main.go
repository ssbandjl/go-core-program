package main

/*
#include <stdlib.h>

void* makeslice(size_t memsize) {
    return malloc(memsize);
}
*/
import "C"
import (
	"log"
	"unsafe"
)

func makeByteSlize(n int) []byte {
	p := C.makeslice(C.size_t(n))
	return ((*[1 << 31]byte)(p))[0:n:n]
}

func freeByteSlice(p []byte) {
	C.free(unsafe.Pointer(&p[0]))
}

func main() {
	log.Printf("分配内存:%vMB", (1<<32+1)/1024/1024)
	s := makeByteSlize(1<<32 + 1)
	s[len(s)-1] = 255
	print(s[len(s)-1])
	freeByteSlice(s)
}
