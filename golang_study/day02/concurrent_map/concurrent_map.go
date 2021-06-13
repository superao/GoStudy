package main

// GitHub地址: https://github.com/satori/go.uuid.git
// 在本地引入远程依赖的包:
// go get github.com/satori/go.uuid         // 默认获取的为最新版
// go get github.com/satori/go.uuid@v1.1.0  // 获取tag为v1.1.0的版本
// go get github.com/satori/go.uuid@master  // 获取分支为master上的版本
// go get github.com/satori/go.uuid@fb121vd // 获取提交git哈希为fb121vd的版本

import (
	"fmt"
	cm "github.com/easierway/concurrent_map"
)

func main() {
	m := cm.CreateConcurrentMap(10)
	m.Set(cm.StrKey("hello"), 10)

	fmt.Println(m.Get(cm.StrKey("hello")))
}