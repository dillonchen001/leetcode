//给你一个字符串 s，找到 s 中最长的回文子串。 
//
// 
//
// 示例 1： 
//
// 
//输入：s = "babad"
//输出："bab"
//解释："aba" 同样是符合题意的答案。
// 
//
// 示例 2： 
//
// 
//输入：s = "cbbd"
//输出："bb"
// 
//
// 示例 3： 
//
// 
//输入：s = "a"
//输出："a"
// 
//
// 示例 4： 
//
// 
//输入：s = "ac"
//输出："a"
// 
//
// 
//
// 提示： 
//
// 
// 1 <= s.length <= 1000 
// s 仅由数字和英文字母（大写和/或小写）组成 
// 
// Related Topics 字符串 动态规划 
// 👍 3770 👎 0


package test

import (
    "fmt"
    "testing"
)

func TestLongestPalindromic(t *testing.T) {
    str := "babad"
    res := longestPalindrome(str)
    fmt.Println(res)
}

//leetcode submit region begin(Prohibit modification and deletion)
func longestPalindrome(s string) string {
    ls := len(s)
    if ls <= 1 {
        return s
    }

    tmpStr := "#"
    for i := 0; i < ls; i++ {
        tmpStr += string(s[i]) + "#"
    }

    lt := len(tmpStr)
    location := make(map[int]int)
    location[0] = 1
    mid := 0
    max := 1

    resl := 1
    resi := 0


    for i := 0; i < lt - 1; i++ {
        if max > i {
            if location[2*mid-i] > max - i {
                location[i] = max - i
            } else {
                location[i] = location[2*mid-i]
            }
        } else {
            location[i] = 1
        }

        for {
            if i < location[i] || i + location[i] > lt - 1 ||  tmpStr[i-location[i]] != tmpStr[i+location[i]] {
                break
            }
            location[i]++
        }

        if location[i] + i > max {
            max = location[i] + i
            mid = i
        }

        if location[i] > resl {
            resl = location[i]
            resi = i
        }
    }
    start := (resi + 1 - resl)/2

    return s[start:start+resl-1]
}
//leetcode submit region end(Prohibit modification and deletion)
