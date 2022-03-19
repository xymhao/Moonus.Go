package main

// https://leetcode-cn.com/problems/reverse-string/
func reverseString(s []byte) {
	for i, j := 0, len(s)-1; i < j; {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}

func reverseWords(s string) string {
	var ret []byte
	for i := 0; i < len(s); {
		start := i
		for i < len(s) && s[i] != ' ' {
			i++
		}

		for x := start; x < i; x++ {
			ret = append(ret, s[i+start-1-x])
		}

		for ; i < len(s) && s[i] == ' '; i++ {
			ret = append(ret, ' ')
		}
	}
	return string(ret)
}

func main() {
	words := reverseWords("Let's take LeetCode contest")
	if words == "s'teL ekat edoCteeL tsetnoc" {
		println(true)
	}
}
