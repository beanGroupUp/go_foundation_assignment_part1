package main

import (
	"fmt"
)

/*
给你一个 非空 整数数组 nums ，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
你必须设计并实现线性时间复杂度的算法来解决此问题，且该算法只使用常量额外空间。
*/

func OnceNumber(nums []int) int {
	single := 0
	for _, num := range nums {
		/*房间里有：[4, 1, 2, 1, 2]

		第一个人是 4：你没见过他，和他绑定。

		你手里现在：4

		第二个人是 1：你没见过他，和他绑定。

		你手里现在：4 和 1 -> 但计算机用一个小技巧把它们合在一起变成了 5（这个细节不用管，知道是“绑定”了就行）。

		第三个人是 2：你没见过他，和他绑定。

		你手里现在：4, 1, 2 -> 合起来变成了 7。

		第四个人是 1：你见过他！（因为之前和1绑定过）。根据规则，要取消绑定。

		你把之前绑定的1取消了。

		你手里现在剩下了：4 和 2 -> 合起来是 6。

		第五个人是 2：你见过他！（因为之前和2绑定过）。根据规则，要取消绑定。

		你把之前绑定的2也取消了。

		你手里最后只剩下：4。*/
		single ^= num
	}
	return single
}

/**
回文数
给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
例如，121 是回文，而 123 不是。
*/

func isPalindrome(x int) bool {
	// 检查特殊情况：负数或以0结尾但不是0的数
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	// 初始化反转数字变量
	revertedNumber := 0

	// 当原始数字大于反转后的数字时继续循环
	for x > revertedNumber {
		// 将x的最后一位加到revertedNumber的末尾
		revertedNumber = revertedNumber*10 + x%10
		// 移除x的最后一位
		x /= 10
	}

	// 检查数字是否相等（偶数位）或除去中间位后是否相等（奇数位）
	return x == revertedNumber || x == revertedNumber/10
}

/**
有效括号
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
每个右括号都有一个对应的相同类型的左括号。
*/

func validParentheses(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}
	paris := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	//定义一个数组，作为栈
	stack := []byte{}
	for i := 0; i < n; i++ {
		if paris[s[i]] > 0 {
			//len(stack) == 0 - 栈为空
			//stack[len(stack)-1] != pairs[s[i]] - 栈顶元素不匹配
			if len(stack) == 0 || stack[len(stack)-1] != paris[s[i]] {
				return false
			}
			// stack[:len(stack)-1] - 创建原切片的一个新切片，包含从开始到倒数第二个元素
			//赋值给 stack - 用新切片替换原切片，实现弹出效果
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}

/*
*
最长公共前缀
编写一个函数来查找字符串数组中的最长公共前缀。
如果不存在公共前缀，返回空字符串 ""。
*/
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		//空数组直接返回
		return ""
	}
	//定义一个初始值
	prefix := strs[0]
	count := len(strs)
	//因为初始值的下标为0，所以循环是从1开始的
	for i := 1; i < count; i++ {
		prefix = lcp(prefix, strs[i])
		if len(prefix) == 0 {
			break
		}
	}
	return prefix
}

func lcp(prefix string, s string) string {
	//选取入参中最小的字符串长度
	length := min(len(s), len(prefix))
	//定义下标初始值
	index := 0
	for index < length && prefix[index] == s[index] {
		index++
	}
	//找出适合的下标截取
	return prefix[:index]
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

/**
删除有序数组中的重复项

给你一个 非严格递增排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。然后返回 nums 中唯一元素的个数。

考虑 nums 的唯一元素的数量为 k ，你需要做以下事情确保你的题解可以被通过：

更改数组 nums ，使 nums 的前 k 个元素包含唯一元素，并按照它们最初在 nums 中出现的顺序排列。nums 的其余元素与 nums 的大小不重要。
返回 k 。
判题标准:

系统会用下面的代码来测试你的题解:

int[] nums = [...]; // 输入数组
int[] expectedNums = [...]; // 长度正确的期望答案

int k = removeDuplicates(nums); // 调用

assert k == expectedNums.length;
for (int i = 0; i < k; i++) {
    assert nums[i] == expectedNums[i];
}
如果所有断言都通过，那么您的题解将被 通过。
*/

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	slow := 1
	for fast := 1; fast < n; fast++ {
		if nums[fast] != nums[slow] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

/**
加一

给定一个表示 大整数 的整数数组 digits，其中 digits[i] 是整数的第 i 位数字。这些数字按从左到右，从最高位到最低位排列。这个大整数不包含任何前导 0。

将大整数加 1，并返回结果的数字数组。



示例 1：

输入：digits = [1,2,3]
输出：[1,2,4]
解释：输入数组表示数字 123。
加 1 后得到 123 + 1 = 124。
因此，结果应该是 [1,2,4]。
示例 2：

输入：digits = [4,3,2,1]
输出：[4,3,2,2]
解释：输入数组表示数字 4321。
加 1 后得到 4321 + 1 = 4322。
因此，结果应该是 [4,3,2,2]。
示例 3：

输入：digits = [9]
输出：[1,0]
解释：输入数组表示数字 9。
加 1 得到了 9 + 1 = 10。
因此，结果应该是 [1,0]。


提示：

1 <= digits.length <= 100
0 <= digits[i] <= 9
digits 不包含任何前导 0。
*/

func plusOne(digits []int) []int {
	//获取数组长度，即数字的位数
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		//如果当前为不是9，加一不会产生进位
		if digits[i] != 9 {
			//当前位+1
			digits[i]++
			//将当前位之后的所有位都重置为0
			//这是因为在加一过程中，这些原本都是9并产生了进位
			for j := i + 1; j < n; j++ {
				digits[j] = 0
			}
			return digits //返回结果集
		}
	}
	//如果所有为都是9（如999），需要增加一位（如1000）
	digits = make([]int, n+1) //创建长度增加1的新数组
	digits[0] = 1             //最高位设为1，其余默认为0
	return digits             //返回新数组
}

/**
两数之和

给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案，并且你不能使用两次相同的元素。

你可以按任意顺序返回答案。



示例 1：

输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
示例 2：

输入：nums = [3,2,4], target = 6
输出：[1,2]
示例 3：

输入：nums = [3,3], target = 6
输出：[0,1]


提示：

2 <= nums.length <= 104
-109 <= nums[i] <= 109
-109 <= target <= 109
只会存在一个有效答案
*/

func twoSum01(nums []int, target int) []int {
	for i, x := range nums {
		for j := i + 1; j < len(nums); j++ {
			if x+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSum02(nums []int, target int) []int {
	mp := make(map[int]int, len(nums))
	for i, v := range nums {
		if idx, ok := mp[target-v]; ok {
			return []int{idx, i}
		}
		mp[v] = i
	}
	return nil
}

func main() {
	//非空整数数组,只出现一次的元素
	nums := []int{2, 2, 1}
	result1 := OnceNumber(nums)
	fmt.Println("非空整数数组，只出现一次的元素为：", result1)
	//回文数
	x := 121
	flag := isPalindrome(x)
	fmt.Printf("数字 %d 是回文子串：%t\n", x, flag)
	//有效括号
	s := "[](){}"
	flagV := validParentheses(s)
	fmt.Println("入参为 %d 的有效括号为%t\n", s, flagV)
	//最长公共前缀
	strs := []string{"flower", "flow", "flight"}
	pre := longestCommonPrefix(strs)
	fmt.Println("入参为 %d 的最长公共前缀%t\n", strs, pre)

	//删除有序数组中的重复项
	param01 := []int{1, 2, 1, 1}
	result2 := removeDuplicates(nums)
	fmt.Println("入参为 %d 的删除有序数组中的重复项%t\n", param01, result2)
	//加一
	parma02 := []int{1, 2, 3}
	result3 := plusOne(parma02)
	fmt.Println("入参为 %d 的加一%t\n", parma02, result3)
	//两数之和
	param03 := []int{2, 7, 11, 15}
	target := 9
	result4 := twoSum02(param03, target)
	fmt.Println("入参为 %d 的两数之和%t\n", param03, result4)
}
