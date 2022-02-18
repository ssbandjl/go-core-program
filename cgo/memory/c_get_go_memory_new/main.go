package main

/*
#include<stdio.h>

void printString(const char* s, int n) {
    int i;
    for(i = 0; i < n; i++) {
        putchar(s[i]);
    }
    putchar('\n');
}
*/
import "C"
import (
	"reflect"
	"unsafe"
)

func printString(s string) {
	p := (*reflect.StringHeader)(unsafe.Pointer(&s))
	C.printString((*C.char)(unsafe.Pointer(p.Data)), C.int(len(s)))
}

// 现在的处理方式更加直接，且避免了分配额外的内存。完美的解决方案！
func main() {
	s := "hello"
	printString(s)
}
