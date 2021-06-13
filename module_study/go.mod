module module_study

go 1.16

require (
	// 导入本地依赖的包(不同项目)
	mymodule v1.0.0
)

replace (
	mymodule v1.0.0 => D:\goproject\src\mymodule
)
