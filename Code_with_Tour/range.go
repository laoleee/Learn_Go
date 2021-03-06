// for 循环的 range 形式可遍历切片或映射
// 当使用 for 循环遍历切片时，每次迭代都会返回两个值。
// 第一个值为当前元素的下标，第二个值为该下标所对应元素的一份副本

package main

import "fmt"

func main() {
	var values = []int{1, 2, 3, 4, 5, 6, 7}

	for i, j := range values { // range关键字迭代
		fmt.Println(i, j)
	}
}
