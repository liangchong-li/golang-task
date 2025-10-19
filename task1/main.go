package main

import (
	"fmt"
	"strconv"
	// "errors"
)

func main() {
	// var arr1 = []int{1,1,2,2,3,4,5,4,5}
	// var arr2 = []int{1,1,2,2,4,5,4,5}
	// target1 := singleNumber(arr1)
	// fmt.Println(target1)
	// target2 := singleNumber(arr2)
	// fmt.Println(target2)

	// isPalindrome1 := isPalindrome(12321)
	// fmt.Println(isPalindrome1)
	// isPalindrome2 := isPalindrome(1221)
	// fmt.Println(isPalindrome2)
	// isPalindrome3 := isPalindrome(122)
	// fmt.Println(isPalindrome3)

	// isValid1 := isValid("()")
	// fmt.Println(isValid1)
	// isValid2 := isValid("()[]{}")
	// fmt.Println(isValid2)
	// isValid3 := isValid("(]")
	// fmt.Println(isValid3)
	// isValid4 := isValid("([])")
	// fmt.Println(isValid4)
	// isValid5 := isValid("([)]")
	// fmt.Println(isValid5)

	// strs1 := []string{"flower","flow","flight"}
	// longestCommonPrefix1 := longestCommonPrefix(strs1)
	// fmt.Println(longestCommonPrefix1)
	// strs2 := []string{"dog","racecar","car"}
	// longestCommonPrefix2 := longestCommonPrefix(strs2)
	// fmt.Println(longestCommonPrefix2)
	// strs3 := []string{"flower","flow","floight"}
	// longestCommonPrefix3 := longestCommonPrefix(strs3)
	// fmt.Println(longestCommonPrefix3)

	// digits1 := []int{1,2,3}
	// result1 := plusOne(digits1)
	// fmt.Println(result1)
	// digits2 := []int{4,3,2,1}
	// result2 := plusOne(digits2)
	// fmt.Println(result2)
	// digits3 := []int{9}
	// result3 := plusOne(digits3)
	// fmt.Println(result3)

	// num1 := []int{1,1,2}
	// length1 := removeDuplicates(num1)
	// fmt.Println(length1)
	// num2 := []int{0,0,1,1,1,2,2,3,3,4}
	// length2 := removeDuplicates(num2)
	// fmt.Println(length2)

	// intervals1 := [][]int{
	// 	{1,3},
	// 	{2,6},
	// 	{8,10},
	// 	{15,18},
	// }
	// aflterMerge1 := merge(intervals1)
	// fmt.Println(aflterMerge1)
	// intervals2 := [][]int{
	// 	{1,4},
	// 	{4,5},
	// }
	// aflterMerge2 := merge(intervals2)
	// fmt.Println(aflterMerge2)
	// intervals3 := [][]int{
	// 	{4,7},
	// 	{1,4},
	// }
	// aflterMerge3 := merge(intervals3)
	// fmt.Println(aflterMerge3)
	// intervals4 := [][]int{
	// 	{1,4},
	// 	{2,3},
	// }
	// aflterMerge4 := merge(intervals4)
	// fmt.Println(aflterMerge4)
	// intervals5 := [][]int{
	// 	{2,3},
	// 	{4,5},
	// 	{6,7},
	// 	{8,9},
	// 	{1,10},
	// }
	// aflterMerge5 := merge(intervals5)
	// fmt.Println(aflterMerge5)

	// nums1 := []int{2,7,11,15}
	// target1 := 9
	// tartNums1 := twoSum(nums1, target1)
	// fmt.Println(tartNums1)
	// nums2 := []int{-3,4,3,90}
	// target2 := 0
	// tartNums2 := twoSum(nums2, target2)
	// fmt.Println(tartNums2)

	var a = "abc"
	b := "abc"
	fmt.Println(a)
	fmt.Println(b)
	demo()
}

func demo() {
	var a,b,c = 1,2,3
	fmt.Println("demo", a, b, c)
}

// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
// 可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，
// 例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
func singleNumber(nums []int) int {
	tempMap := make(map[int]int)
	for _, v := range nums {
		_,ok := tempMap[v]
		if ok {
			tempMap[v] += 1	
		}else {
			tempMap[v] = 1
		}
		// fmt.Println(tempMap)
	}
	for k, v := range tempMap {
		if v < 2 {
			return k
		}
	}
	return 0
	// return errors.New("no single number")
}

// 给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
// 回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
// 考察：数字操作、条件判断
func isPalindrome(x int) bool {
	s := strconv.Itoa(x)

	var left int = 0
	var right int = len(s) - 1

	for {
		// 左右指针相遇或交叉时退出
		if left >= right {
			return true
		}
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}

	return false
}


// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

// 有效字符串需满足：
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
// 每个右括号都有一个对应的相同类型的左括号。

// 考察：字符串处理、栈的使用
func isValid(s string) bool {
	contain := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	stack := make([]rune, 0)

	for _, v := range s {
		if len(stack) != 0 && stack[len(stack) - 1] == contain[v] {
			stack = stack[:len(stack) - 1]
		}else {
			stack = append(stack, v)
		}
		
	}
	
	return len(stack) == 0
}

// 编写一个函数来查找字符串数组中的最长公共前缀。
// 如果不存在公共前缀，返回空字符串 ""。
// 考察：字符串处理、循环嵌套
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	pre := strs[0]
	for i := 1; i < len(strs); i++ {
		index := 0
		var minLength int
		if len(strs[i]) > len(pre) {
			minLength = len(pre)
		}else {
			minLength = len(strs[i])
		}
		
		for index < minLength && pre[index] == strs[i][index] {
			index++
		}
		pre = pre[:index]
	}
	return pre
}

// 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
// 将大整数加 1，并返回结果的数字数组。
// 考察：数组操作、进位处理
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >=0 ; i-- {
		if digits[i] == 9 {
			digits[i] = 0
			// 如果原切片第一位已经进位了，需要再增加一位
			if i == 0 {
				digits = append([]int{1}, digits...)
				return digits 
			}
		}else {
			digits[i] += 1
			return digits
		}
	}
	return digits
}

// 给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。
// 不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
// 可以使用双指针法，一个慢指针 i 用于记录不重复元素的位置，一个快指针 j 用于遍历数组，
// 当 nums[i] 与 nums[j] 不相等时，将 nums[j] 赋值给 nums[i + 1]，并将 i 后移一位。
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0
	j := i + 1
	// 1222223445
	// 1232223445
	// 1232223445
	for j < len(nums){
		if nums[i] != nums[j] {
			nums[i+1] = nums[j]
			i++
		}else {
			j++
		}
	}
	return i + 1
}

// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
// 请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
// 可以先对区间数组按照区间的起始位置进行排序，然后使用一个切片来存储合并后的区间，遍历排序后的区间数组，
// 将当前区间与切片中最后一个区间进行比较，如果有重叠，则合并区间；如果没有重叠，则将当前区间添加到切片中。
func merge(intervals [][]int) [][]int {
	n := len(intervals)
    // 外层循环控制遍历次数
    for i := 0; i < n-1; i++ {
        // 内层循环进行相邻元素的比较和可能的交换
        for j := 0; j < n-i-1; j++ { // 每次内循环可以少比较一次，因为每次外循环会将一个最大值放到正确的位置
            if intervals[j][0] > intervals[j+1][0] {
                // 交换 arr[j] 和 arr[j+1]
                intervals[j], intervals[j+1] = intervals[j+1], intervals[j]
            }
        }
    }
	afterSort := intervals

	fmt.Println("afterSort", afterSort)
	merge := make([][]int, 0)
	merge = append(merge, afterSort[0])
	for i := 1; i < len(afterSort); i++ {
		fmt.Println("afterSort[i]：", afterSort[i])
		// 如果重叠了
		if afterSort[i][0] >= merge[len(merge) - 1][0] && afterSort[i][0] <= merge[len(merge) - 1][1] {
			merge[len(merge) - 1] = []int{min(merge[len(merge) - 1][0], afterSort[i][0]), max(merge[len(merge) - 1][1], afterSort[i][1])}
		}else if merge[len(merge) - 1][0] >= afterSort[i][0] && merge[len(merge) - 1][0] <= afterSort[i][1] {
			merge[len(merge) - 1] = []int{min(merge[len(merge) - 1][0], afterSort[i][0]), max(merge[len(merge) - 1][1], afterSort[i][1])}
		}else {
			merge = append(merge, afterSort[i])
		}
		fmt.Println(merge)
	}

	return merge
}

func min(x,y int) int{
	if x > y {
		return y
	}else {
		return x
	}
}
func max(x,y int) int{
	if x > y {
		return x
	}else {
		return y
	}
}

// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
// 考察：数组遍历、map使用
func twoSum(nums []int, target int) []int {
	contain := make(map[int]int)
	for i, v := range nums {
		contain[i] = v
	}

	for k, v := range contain {
		for k2, v2 := range contain {
			if k != k2 && v + v2 == target {
				return []int{k, k2}
			}
		}	
	}
    return nums
}