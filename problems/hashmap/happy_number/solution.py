# todo rewrite to Floyd Cycle Detection
# https://leetcode.com/problems/happy-number
def is_happy(n: int) -> bool:
    n = str(n)
    nums = set()

    return _is_happy(n, nums)

def _is_happy(n: str, nums: set[str]):
    res = 0
    for i in n:
        res += int(i)**2
    res = str(res)
    if res in nums:
        return False
    nums.add(res)
    if res != '1':
        return _is_happy(res, nums)
    return True


if __name__ == '__main__':
    num = 19
    result = is_happy(num)

    assert result, f'{num} is happy'

    num = 2
    result = is_happy(num)

    assert not result, f'{num} is not happy'
