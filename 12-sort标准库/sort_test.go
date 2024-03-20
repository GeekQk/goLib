package main

import (
	"fmt"
	"sort"
	"testing"
)

type NewInts []uint

func (n NewInts) Len() int {
	return len(n)
}

func (n NewInts) Less(i, j int) bool {
	// fmt.Println(i, j, n[i] < n[j], n)
	return n[i] < n[j]
}

func (n NewInts) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func TestSortUint(t *testing.T) {
	//uint排序
	n := []uint{1, 3, 2, 5}
	sort.Sort(NewInts(n))
	fmt.Println(n)
	//倒序排行
	sort.Sort(sort.Reverse(NewInts(n)))
	fmt.Println(n)

	// 使用 sort.Slice 对切片进行排序
	nums := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	fmt.Println(nums) // 输出: [1 1 2 3 3 4 5 5 5 6 9]
}

func TestSortInt(t *testing.T) {
	nums := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}

	// 正序排序
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	fmt.Println("正序排序:", nums) // 输出: [1 1 2 3 3 4 5 5 5 6 9]

	// 倒序排序
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	fmt.Println("倒序排序:", nums) // 输出: [9 6 5 5 5 4 3 3 2 1 1]
}

// 字符串是不可变的，因此不能直接对其进行排序。
// 但如果你有一个字符串切片，你可以对切片中的字符串进行排序。
func TestSortString(t *testing.T) {
	strs := []string{"banana", "apple", "cherry"}

	// 正序排序
	sort.Strings(strs)
	fmt.Println("正序排序:", strs) // 输出: [apple banana cherry]

	// 倒序排序
	sort.Sort(sort.Reverse(sort.StringSlice(strs)))
	fmt.Println("倒序排序:", strs) // 输出: [cherry banana apple]
}

type Person struct {
	Name string
	Age  int
}

func TestSortStruct(t *testing.T) {
	people := []Person{
		{"Alice", 30},
		{"Bob", 20},
		{"Charlie", 25},
	}

	// 按 Age 正序排序
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println("按 Age 正序排序:", people)

	// 按 Age 倒序排序
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age > people[j].Age
	})
	fmt.Println("按 Age 倒序排序:", people)
}

// 哈希表（在Go语言中通常指map）是无序的，不能直接排序。
// 但是，你可以通过创建一个切片来存储哈希表的键或值，然后对切片进行排序。
// 排序之后，你可以通过排序后的键来重新排列哈希表。
func TestSortHash(t *testing.T) {
	// 创建一个哈希表
	m := map[string]int{
		"apple":  3,
		"banana": 2,
		"cherry": 1,
	}

	// 正序排列
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys) // 对键进行排序

	fmt.Println("正序排列:")
	for _, k := range keys {
		fmt.Printf("%s: %d\n", k, m[k])
	}

	// 倒序排列
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] > keys[j] // 逆序比较函数
	})

	fmt.Println("倒序排列:")
	for _, k := range keys {
		fmt.Printf("%s: %d\n", k, m[k])
	}
}
