package greatest_common_divisor_of_strings

// GreatestCommonDivisorOfStrings https://leetcode.com/problems/greatest-common-divisor-of-strings
func GreatestCommonDivisorOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}
	ln := gcd(len(str1), len(str2))
	return str1[:ln]
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
