// Package strKit
/*
PS:  字符串间的比较可以用 "==" 或 "strings.Compare()"；
*/
package strKit

import (
	"fmt"
	"strings"
)

// Format 格式化文本，类似Java的"StrUtil.format".
/*
Deprecated: 直接用 fmt.Sprintf() 对编码更友好，比如GoLand.
*/
var Format func(format string, args ...interface{}) string = fmt.Sprintf

var ToLower func(s string) string = strings.ToLower

var ToUpper func(s string) string = strings.ToUpper

// Index
/*
PS:
(1) s中不存在substr的话，返回 -1
(2) s中存在多个substr的话，返回 第一个的下标

@param s	被查找的字符串
@param str	查找的字符串

e.g.
("abcabc", "ab")	=> 0
("bcabc", "ab")		=> 2
("23", "1")			=> -1
*/
var Index func(s, str string) int = strings.Index

// LastIndex
/*
e.g.
("", "1") => -1
*/
var LastIndex func(s, str string) int = strings.LastIndex

// Contains 是否包含（区分大小写）
/*
e.g.
("", "1") 		=> false
("abc", "Abc") 	=> false
*/
var Contains func(s, substr string) bool = strings.Contains

// ContainsIgnoreCase 是否包含（不区分大小写）
/*
e.g.
("abc", "Abc") 	=> true
("abcde", "BC")	=> true
*/
func ContainsIgnoreCase(s, substr string) bool {
	return strings.Contains(ToLower(s), ToLower(substr))
}

// Count 计数
/*
@return >= 0

e.g.
("12345", "3")	=>	1
("12345", "6")	=>	0

e.g.1 If substr is an empty string, Count returns 1 + the number of Unicode code points in s.
strKit.Count("12345", "") 	=> 6
*/
func Count(s, substr string) int {
	return strings.Count(s, substr)
}

// Replace
/*
@param s	原始字符串
@param old	要被替换的子串
@param new	替换后的子串
@param n	替换次数（-1 表示替换所有匹配项）

e.g.
	("abcdcba", "a", "0", 1)  => "0bcdcba"
	("abcdcba", "a", "0", 2)  => "0bcdcb0"
	("abcdcba", "a", "0", -1) => "0bcdcb0"
*/
func Replace(s, old, new string, n int) string {
	return strings.Replace(s, old, new, n)
}

// ReplaceAll
/*
等价于 Replace 的 n = -1 的情况

e.g.
	("12321", "2", "0") => "10301"
*/
func ReplaceAll(s, old, new string) string {
	return strings.ReplaceAll(s, old, new)
}

func Join(elements []string, sep string) string {
	return strings.Join(elements, sep)
}

// Split
/*
@param sep  可以为""，此种情况比较特殊，详见下例；
			如果s中不存在sep，那么返回切片的长度为1
@return 必定不为nil && len >= 1

e.g.
("", "-") 		=> []（长度为1；唯一的元素为""）
("-", "-")		=> []string{"", ""}
("123-", "-")	=> []string{"123", ""}
e.g.1
("hello world!", "") => [h e l l o   w o r l d !]
("1-2-3", "-") 	=> [1 2 3]（长度为3）
("1-2-3", "+") 	=> [1-2-3]（长度为1）
("172.18.21.50;8095;", ";") => [172.18.21.50 8095 ]（长度为3，第三个元素为""）
*/
var Split func(s, sep string) []string = strings.Split

// Compare 字符串比较
/*
@return (1) -1: if a < b
		(2)  0: if a == b
		(3) +1: if a > b
*/
var Compare func(a, b string) int = strings.Compare

// Equals 字符串是否相等？（区分大小写）
/*
PS: 也可以用 == 进行判断字符串是否相等（但不建议这么干）.

e.g.
("abc", "Abc") => false
*/
func Equals(str, str1 string) bool {
	return strings.Compare(str, str1) == 0
}

// EqualsIgnoreCase 字符串是否相等？（不区分大小写）
/*
e.g.
("abc", "Abc") => true
*/
var EqualsIgnoreCase func(str, str1 string) bool = strings.EqualFold
