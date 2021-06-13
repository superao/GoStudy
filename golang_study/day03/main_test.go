package day03

import (
	"reflect"
	"testing"
)

// 演示数组元素的使用
func TestArray(t *testing.T) {
	// 声明一个数组长度为3的int类型数组
	var arr [3]int
	t.Log(arr, len(arr), cap(arr))
	arr[0] = 200
	t.Log(arr, len(arr), cap(arr))

	// 声明一个数组并进行初始化
	arr1 := [3]int {1, 2, 3}
	t.Log(arr1, len(arr1), cap(arr1))
	arr1[2] = 4
	t.Log(arr1, len(arr1), cap(arr1))

	// 声明一个数组并进行部分初始化
	arr2 := [5]int {1:2, 2:4, 3:6}
	t.Log(arr2, len(arr2), cap(arr2))
	arr2[2] = 5
	t.Log(arr2, len(arr2), cap(arr2))

	// 多维数组的初始化，第一个表示行数，第二个表示列数
	arr3 := [1][2]int{{1, 2}}
	t.Log(arr3, len(arr3), cap(arr3))    // len = 1, cap = 1
	arr3[0][1] = 5
	t.Log(arr3, len(arr3), cap(arr3))

	// 多维数组的部分初始化
	arr4 := [3][2]int{0:{1, 2}, 2:{4, 5}}
	t.Log(arr4, len(arr4), cap(arr4))    // len = 3, cap = 3
	arr4[2][1] = 100
	t.Log(arr4, len(arr4), cap(arr4))

	// 不指定个数的数组类型
	//var arr5 [...]int   // 不定个数的数组类型必须对其初始化，否则无法使用，因为编译器无法知道分配多大空间
	arr5 := [...]int{1,2,3,4,5,6,7}
	t.Log(len(arr5), cap(arr5))  // 7 7
	arr5[5] = 5
	t.Log(len(arr5), cap(arr5))  // 7 7

	// 不定个数的多维数组的初始化，
	arr6 := [...][1]int {{1}, {2}, {3}}
	t.Log(arr6, len(arr6), cap(arr6))  // 3 3
	// 注意：Go语言和其他语言是一样的，多维数组的列数必须指定，否则编译错误
}

// 数组元素之间的比较
func TestArrayDiff(t *testing.T) {
	// 声明三个完全相同的数组类型
	arr1 := [3]int{1, 2, 3}
	arr2 := [3]int{4, 5, 6}
	arr3 := [3]int{1, 2, 3}
	//arr4 := [5]int{1, 2, 3}

	if arr1 == arr3 {
		t.Log("arr1 == arr3, 它们是相同的元素，类型和数据都是一样的")
	}
	if arr1 == arr2 {
		t.Log("arr1 == arr2, 它们是相同的元素，类型和数据都是一样的")
	} else {
		t.Log("arr1 != arr2, 它们不是相同的元素，数据是不一样的")
	}
	//if arr1 != arr4 {        // 比较时类型不匹配直接编译报错
	//	t.Log("arr1 != arr4, 它们是相同的元素，类型和数据都是一样的")
	//}
}

// 数组元素的遍历
func TestArrayRange(t *testing.T) {
	// 遍历一维数组
	arr := [5]int{1,2,3,4,5}
	for idx, e := range arr {
		t.Log(idx, e)
	}
	for i := 0; i < len(arr); i++ {
		t.Log(i, arr[i])
	}

	// 遍历二维数组
	arr1 := [3][2]int{{1,2},{3,4},{5,6}}
	for idx, ee := range arr1 {  // 有三个元素，每一个元素是 [2]int 类型的数组
		for i, e := range ee {
			t.Log(idx, i, e)     // 行坐标、纵坐标、值
		}
	}
	for i:=0; i < len(arr1); i++ {
		for j:=0; j < len(arr1[i]); j++ {
			t.Log(i, j, arr1[i][j])   // 行坐标、纵坐标、值
		}
	}
}

// 数组截取
func TestArraySlice(t *testing.T) {
	arr := [5]int{1,2,3,4,5}
	t.Logf("arr 的类型为: %T", arr)
	t.Log(reflect.TypeOf(arr))

	arr1 := arr[1:]
	t.Log(arr1, reflect.TypeOf(arr1))   // 数组截取后会变为切片类型

	arr2 := arr[:]
	t.Log(arr2, reflect.TypeOf(arr2))
}

// 切片类型的相关操作
func TestSlice(t *testing.T) {
	var s []int
	t.Log(s, len(s), cap(s))  // 0 0

	s1 := []int{1,2,3}
	t.Log(s1, len(s1), cap(s1))  // 3 3

	// 向切片变量中增加元素
	for i := 1; i <= 20; i++ {
		t.Logf("切片的地址为: %p, 个数: %v, 容量: %v", s1, len(s1), cap(s1))
		s1 = append(s1, i + 10)
	}

	// 切片类型的截取
	s2 := s1[0:2]
	t.Log(s2, reflect.TypeOf(s2))

	// make 函数初始化切片
	s3 := make([]int, 2, 10)
	t.Log(s3, len(s3), cap(s3))

	// 两个切片共享同一份储存空间
	//arr := [5]int{1,2,3,4,5}    // 一样的结果
	arr := []int{1,2,3,4,5}
	ss1 := arr[0:3]
	ss2 := arr[0:2]
	ss1[0] = 100
	ss1 = append(ss1, 6)
	ss1 = append(ss1, 6)
	ss1 = append(ss1, 6)
	ss1 = append(ss1, 6)
	ss1 = append(ss1, 6)
	t.Log(ss1)
	t.Log(ss2)
}

