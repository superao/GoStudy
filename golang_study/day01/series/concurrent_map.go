package series

// 导入包的路径是从 src 后开始写的，由于之前设置项目路径的时候有问题就多了一层 src 目录。
import "src/github.com/easierway/concurrent_map"

// TestConcurrentMap 本示例演示 vendor 包管理工具的使用
func TestConcurrentMap() {
	// 创建一个并发的 map
	concurrent_map.CreateConcurrentMap(10)
}
