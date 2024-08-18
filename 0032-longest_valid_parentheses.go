// 32. Longest Valid Parentheses

// Given a string containing just the characters '(' and ')', return the length of the longest valid (well-formed)
// parentheses substring.

// Example 1:
// Input: s = "(()"
// Output: 2
// Explanation: The longest valid parentheses substring is "()".

// Example 2:
// Input: s = ")()())"
// Output: 4
// Explanation: The longest valid parentheses substring is "()()".

// Example 3:
// Input: s = ""
// Output: 0

// Constraints:
// 0 <= s.length <= 3 * 104
// s[i] is '(', or ')'.

func longestValidParentheses(s string) int {
	maxValidParentheses := 0
	var total int

	for x := 0; x < len(s); x++ {
		if string(s[x]) == "(" {
			total = 1

			for n := x + 1; n < len(s); n++ {

				if string(s[n]) == "(" {
					total += 1
				} else if string(s[n]) == ")" {
					total -= 1
				}

				fmt.Printf("x: %d, s[x]: %c, n: %d, s[n]: %c, total: %d\n", x, s[x], n, s[n], total)

				if total == 0 {
					if maxValidParentheses < n-x+1 {
						maxValidParentheses = n - x + 1
					}
				}

				if total < 0 {
					break
				}

			}
		}
	}
	return maxValidParentheses

}
