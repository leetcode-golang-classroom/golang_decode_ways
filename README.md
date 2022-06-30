# golang_decode_ways

A message containing letters from `A-Z` can be **encoded** into numbers using the following mapping:

```
'A' -> "1"
'B' -> "2"
...
'Z' -> "26"

```

To **decode** an encoded message, all the digits must be grouped then mapped back into letters using the reverse of the mapping above (there may be multiple ways). For example, `"11106"` can be mapped into:

- `"AAJF"` with the grouping `(1 1 10 6)`
- `"KJF"` with the grouping `(11 10 6)`

Note that the grouping `(1 11 06)` is invalid because `"06"` cannot be mapped into `'F'` since `"6"` is different from `"06"`.

Given a string `s` containing only digits, return *the **number** of ways to **decode** it*.

The test cases are generated so that the answer fits in a **32-bit** integer.

## Examples

**Example 1:**

```
Input: s = "12"
Output: 2
Explanation: "12" could be decoded as "AB" (1 2) or "L" (12).

```

**Example 2:**

```
Input: s = "226"
Output: 3
Explanation: "226" could be decoded as "BZ" (2 26), "VF" (22 6), or "BBF" (2 2 6).

```

**Example 3:**

```
Input: s = "06"
Output: 0
Explanation: "06" cannot be mapped to "F" because of the leading zero ("6" is different from "06").

```

**Constraints:**

- `1 <= s.length <= 100`
- `s` contains only digits and may contain leading zero(s).

## 解析

題目給定一個字串 s , 每個字元都是用 ‘0’ 到 ‘9’ 的 digit 來組成

給定一個規則 把 ‘1’ 到 ‘26’ 對應到 ‘A’ 到 ‘Z’

要求實作一個演算法計算有多少種 把 s 轉譯成字母組合的方式

一個 digit 只有一個對應

困難度在於當兩個 digit 時有兩種方式

順著轉譯不是很好看出問題結構

試著反向來解譯

初始化走到最末端也就是只剩空字元，剛好不用解 方法數設定為 1

假設從最後一個字元倒過來來解

如果這個字元是 ‘0’ 單個字元沒有這種對應所以是 0

如果這個字元非 ‘0’ 代表可以用單個字元來解 解這個字元跟從下一個字元開始一樣

如果這個字元與前一個字元可以用兩個字元來做對應 則代表可以從下二個字元開始一樣

舉一個例子來看

s = “121”

其解碼決策樹如下圖：

![](https://i.imgur.com/XZsxtGD.png)

從上圖發現 該解碼問題

根據其相對位置字元的解碼情況會有以下對應關係

f(i) = 0 , if s[i]==’0’ 且 s[i]+s[i+1] 無法被解碼

f(i) = f(i+1) if s[i] ≠ ‘0’ 且 s[i]+s[i+1] 無法被解碼

f(i) = f(i+2) if s[i] == ‘0’ 且 s[i]+s[i+1] 可以被解碼

f(i) = f(i+1) + f(i+2) if s[i] ≠ ‘0’ 且 s[i]+s[i+1] 可以被解碼

## 程式碼
```go
package sol

func numDecodings(s string) int {
	digitMap := map[byte]struct{}{
		'0': {},
		'1': {},
		'2': {},
		'3': {},
		'4': {},
		'5': {},
		'6': {},
	}
	sLen := len(s)
	prevTwo := 1
	prevOne := 0
	if s[sLen-1] != '0' {
		prevOne = prevTwo
	}
	ways := prevOne
	for start := sLen - 2; start >= 0; start-- {
		ways = 0
		if s[start] != '0' {
			ways += prevOne
		}
		if start+1 < sLen {
			if _, ok := digitMap[s[start+1]]; s[start] == '1' || (s[start] == '2' && ok) {
				ways += prevTwo
			}
		}
		prevTwo = prevOne
		prevOne = ways
	}
	return ways
}

```
## 困難點

1. 遇到兩個 digits 需要處理不同種的 decode

## Solve Point

- [x]  初始化 最後一個 位置的解碼方式 = 1
- [x]  透過 遞迴關係推算目前位置的方法數