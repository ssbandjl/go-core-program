package main

import (
	"fmt"
)

//编写一个函数，完成老鼠找路, 右下角是出口
//myMap *[8][7]int:地图，保证是同一个地图，使用引用
//i,j 表示对地图的哪个点进行测试
func SetWay(myMap *[8][7]int, i int, j int) bool {

	//分析出什么情况下，就找到出路
	//myMap[6][5] == 2
	if myMap[6][5] == 2 { //右下角出口
		return true
	} else {
		//说明要继续找
		if myMap[i][j] == 0 { //如果这个点是可以探测

			//假设这个点是可以通, 但是需要探测 上下左右，如果上下左右走不通，则走下面的else语句设置为3
			//换一个策略 下右上左
			myMap[i][j] = 2
			if SetWay(myMap, i+1, j) { //下
				return true //这里返回True表示能走通
			} else if SetWay(myMap, i, j+1) { //右
				return true
			} else if SetWay(myMap, i-1, j) { //上
				return true
			} else if SetWay(myMap, i, j-1) { //左
				return true
			} else { // 死路
				myMap[i][j] = 3
				return false
			}

		} else { // 说明这个点不能探测，为1，是墙
			return false
		}

	}
}

func main() {
	//先创建一个二维数组，模拟迷宫
	//规则
	//1. 如果元素的值为1 ，就是墙
	//2. 如果元素的值为0, 是没有走过的点
	//3. 如果元素的值为2, 是一个通路
	//4. 如果元素的值为3， 是走过的点，但是走不通

	var myMap [8][7]int //8行7列二维数组

	//先把地图的最上和最下设置为1
	for i := 0; i < 7; i++ {
		myMap[0][i] = 1 //第一行全为1
		myMap[7][i] = 1 //最后一行全为1
	}

	//先把地图的最左和最右设置为1
	for i := 0; i < 8; i++ {
		myMap[i][0] = 1 //第一列全为1
		myMap[i][6] = 1 //最后一列全为1
	}

	//堵死了
	myMap[3][1] = 1
	myMap[3][2] = 1
	myMap[1][2] = 1
	myMap[2][2] = 1

	//输出地图
	fmt.Println("原始地图")
	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(myMap[i][j], " ")
		}
		fmt.Println()
	}

	/* 原始地图
	1 1 1 1 1 1 1
	1 0 1 0 0 0 1
	1 0 1 0 0 0 1
	1 1 1 0 0 0 1
	1 0 0 0 0 0 1
	1 0 0 0 0 0 1
	1 0 0 0 0 0 1
	1 1 1 1 1 1 1
	*/

	//使用测试
	// SetWay(&myMap, 1, 1) // 从第2行,第2列开始
	SetWay(&myMap, 2, 3) // 从第2行,第2列开始
	fmt.Println("探测完毕....")
	//输出地图
	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(myMap[i][j], " ")
		}
		fmt.Println()
	}

}
