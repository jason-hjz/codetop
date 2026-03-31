package __LongestSubString

//func main() {
//	fmt.Println(lengthOfLongestSubstring("abcabcc"))
//	return
//}

func lengthOfLongestSubstring(s string) int {
	m := map[byte]int{}
	n := len(s)

	i, rk, ans := 0, 0, 0
	for ; i < n; i++ {
		for rk < n && m[s[rk]] == 0 {
			m[s[rk]]++
			rk++
		}
		ans = max(ans, rk-i)
		delete(m, s[i])
	}
	return ans
}
