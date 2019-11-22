# URL : https://leetcode.com/problems/search-insert-position/
# 32 ms	12.3 MB
class Solution(object):
    def searchInsert(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: int
        """
        for i in range(len(nums)):
            if nums[i] >= target:
                return i
        return len(nums)

arr=[1,3,5,6]
sol=Solution()
print(sol.searchInsert(arr,5))
print(sol.searchInsert(arr,7))
print(sol.searchInsert(arr,0))
