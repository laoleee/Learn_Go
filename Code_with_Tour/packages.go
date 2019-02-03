package main

import (
	"fmt"
	"math/rand"  // 按照惯例，包名与导入路径的最后一个目录一致。例如，`"math/rand"` 包由 package rand 语句开始
)

// 这样导入也可以，但推荐上面的写法
// import "ftm"
// import "math/rand"


func main() {
	fmt.Print("输出个随机数:", rand.Intn(10))
}