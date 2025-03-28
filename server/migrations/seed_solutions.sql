-- Seed file for problem solutions
-- Assuming the problems from seed_test_problems.sql are already loaded

-- Two Sum solutions
INSERT INTO problem_solutions (problem_id, language, solution_code, created_at)
VALUES 
(1, 'Go', 'func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
    for i, num := range nums {
        if j, exists := m[target-num]; exists {
            return []int{j, i}
        }
        m[num] = i
    }
    return nil
}', NOW()),

(1, 'Python', 'def twoSum(self, nums, target):
    seen = {}
    for i, num in enumerate(nums):
        if target - num in seen:
            return [seen[target - num], i]
        seen[num] = i', NOW()),

(1, 'Javascript', 'var twoSum = function(nums, target) {
    const map = new Map();
    for (let i = 0; i < nums.length; i++) {
        const complement = target - nums[i];
        if (map.has(complement)) {
            return [map.get(complement), i];
        }
        map.set(nums[i], i);
    }
};', NOW()),

-- Add Solution to Problem ID 9 (Palindrome Number)
(9, 'Go', 'func isPalindrome(x int) bool {
    // Negative numbers are not palindromes
    if x < 0 {
        return false
    }
    
    // Single digit numbers are palindromes
    if x < 10 {
        return true
    }
    
    // Find the reverse of the number
    original := x
    reversed := 0
    
    for x > 0 {
        reversed = reversed*10 + x%10
        x /= 10
    }
    
    return original == reversed
}', NOW()),

(9, 'Python', 'def isPalindrome(self, x: int) -> bool:
    # Negative numbers are not palindromes
    if x < 0:
        return False
    
    # Convert to string and check if it equals its reverse
    return str(x) == str(x)[::-1]', NOW());