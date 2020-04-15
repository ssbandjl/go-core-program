package main
import (
	"fmt"
)

// var arr = [5]int{20, 11, -1, 44, 9}

func Test(a *[5]int){
	tmp := 0
	fmt.Println("冒泡排序法")
	for j:=0; j < len(*a)-1;j++{
		for i:=0; i<len(*a)-1-j; i++{
			if a[i] > (*a)[i+1]{
				tmp=(*a)[i]
				(*a)[i]=(*a)[i+1]
				(*a)[i+1]=tmp
			}
		}
	}	
}


