# https://leetcode.com/problems/greatest-common-divisor-of-strings
def greatest_common_divisor_of_strings(str1: str, str2: str) -> str:
    if str1 + str2 != str2 + str1:
        return ''

    ln = gcd(len(str1), len(str2))
    return str1[:ln]

def gcd(a: int, b: int) -> int:
    while b > 0:
        a, b = b, a%b
    return a


if __name__ == '__main__':
    str1 = 'ABCABC'
    str2 = 'ABC'

    result = greatest_common_divisor_of_strings(str1, str2)
    target = 'ABC'
    assert result == target, f'greatest_common_divisor_of_strings FAILED:{result} != {target}'

    str1 = 'ABCABC'
    str2 = 'AB'

    result = greatest_common_divisor_of_strings(str1, str2)
    target = ''
    assert result == target, f'greatest_common_divisor_of_strings FAILED:{result} != {target}'
