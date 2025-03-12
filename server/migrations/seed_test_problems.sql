INSERT INTO problems
(id, frontend_id, title, title_slug, difficulty, is_paid_only, "content", topic_tags, example_testcases, similar_questions, created_at)
VALUES(1, 1, 'Two Sum', 'two-sum', 'Easy', false, '<p>Given an array of integers <code>nums</code>&nbsp;and an integer <code>target</code>, return <em>indices of the two numbers such that they add up to <code>target</code></em>.</p>

<p>You may assume that each input would have <strong><em>exactly</em> one solution</strong>, and you may not use the <em>same</em> element twice.</p>

<p>You can return the answer in any order.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<pre>
<strong>Input:</strong> nums = [2,7,11,15], target = 9
<strong>Output:</strong> [0,1]
<strong>Explanation:</strong> Because nums[0] + nums[1] == 9, we return [0, 1].
</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input:</strong> nums = [3,2,4], target = 6
<strong>Output:</strong> [1,2]
</pre>

<p><strong class="example">Example 3:</strong></p>

<pre>
<strong>Input:</strong> nums = [3,3], target = 6
<strong>Output:</strong> [0,1]
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>2 &lt;= nums.length &lt;= 10<sup>4</sup></code></li>
	<li><code>-10<sup>9</sup> &lt;= nums[i] &lt;= 10<sup>9</sup></code></li>
	<li><code>-10<sup>9</sup> &lt;= target &lt;= 10<sup>9</sup></code></li>
	<li><strong>Only one valid answer exists.</strong></li>
</ul>

<p>&nbsp;</p>
<strong>Follow-up:&nbsp;</strong>Can you come up with an algorithm that is less than <code>O(n<sup>2</sup>)</code><font face="monospace">&nbsp;</font>time complexity?', '[{"name": "Array", "slug": "array", "translatedName": null}, {"name": "Hash Table", "slug": "hash-table", "translatedName": null}]'::jsonb, '[2,7,11,15]
9
[3,2,4]
6
[3,3]
6', '[{"title": "3Sum", "titleSlug": "3sum", "difficulty": "Medium", "translatedTitle": null}, {"title": "4Sum", "titleSlug": "4sum", "difficulty": "Medium", "translatedTitle": null}, {"title": "Two Sum II - Input Array Is Sorted", "titleSlug": "two-sum-ii-input-array-is-sorted", "difficulty": "Medium", "translatedTitle": null}, {"title": "Two Sum III - Data structure design", "titleSlug": "two-sum-iii-data-structure-design", "difficulty": "Easy", "translatedTitle": null}, {"title": "Subarray Sum Equals K", "titleSlug": "subarray-sum-equals-k", "difficulty": "Medium", "translatedTitle": null}, {"title": "Two Sum IV - Input is a BST", "titleSlug": "two-sum-iv-input-is-a-bst", "difficulty": "Easy", "translatedTitle": null}, {"title": "Two Sum Less Than K", "titleSlug": "two-sum-less-than-k", "difficulty": "Easy", "translatedTitle": null}, {"title": "Max Number of K-Sum Pairs", "titleSlug": "max-number-of-k-sum-pairs", "difficulty": "Medium", "translatedTitle": null}, {"title": "Count Good Meals", "titleSlug": "count-good-meals", "difficulty": "Medium", "translatedTitle": null}, {"title": "Count Number of Pairs With Absolute Difference K", "titleSlug": "count-number-of-pairs-with-absolute-difference-k", "difficulty": "Easy", "translatedTitle": null}, {"title": "Number of Pairs of Strings With Concatenation Equal to Target", "titleSlug": "number-of-pairs-of-strings-with-concatenation-equal-to-target", "difficulty": "Medium", "translatedTitle": null}, {"title": "Find All K-Distant Indices in an Array", "titleSlug": "find-all-k-distant-indices-in-an-array", "difficulty": "Easy", "translatedTitle": null}, {"title": "First Letter to Appear Twice", "titleSlug": "first-letter-to-appear-twice", "difficulty": "Easy", "translatedTitle": null}, {"title": "Number of Excellent Pairs", "titleSlug": "number-of-excellent-pairs", "difficulty": "Hard", "translatedTitle": null}, {"title": "Number of Arithmetic Triplets", "titleSlug": "number-of-arithmetic-triplets", "difficulty": "Easy", "translatedTitle": null}, {"title": "Node With Highest Edge Score", "titleSlug": "node-with-highest-edge-score", "difficulty": "Medium", "translatedTitle": null}, {"title": "Check Distances Between Same Letters", "titleSlug": "check-distances-between-same-letters", "difficulty": "Easy", "translatedTitle": null}, {"title": "Find Subarrays With Equal Sum", "titleSlug": "find-subarrays-with-equal-sum", "difficulty": "Easy", "translatedTitle": null}, {"title": "Largest Positive Integer That Exists With Its Negative", "titleSlug": "largest-positive-integer-that-exists-with-its-negative", "difficulty": "Easy", "translatedTitle": null}, {"title": "Number of Distinct Averages", "titleSlug": "number-of-distinct-averages", "difficulty": "Easy", "translatedTitle": null}, {"title": "Count Pairs Whose Sum is Less than Target", "titleSlug": "count-pairs-whose-sum-is-less-than-target", "difficulty": "Easy", "translatedTitle": null}]'::jsonb, '2025-03-05 18:17:02.129');
INSERT INTO problems
(id, frontend_id, title, title_slug, difficulty, is_paid_only, "content", topic_tags, example_testcases, similar_questions, created_at)
VALUES(2, 2, 'Add Two Numbers', 'add-two-numbers', 'Medium', false, '<p>You are given two <strong>non-empty</strong> linked lists representing two non-negative integers. The digits are stored in <strong>reverse order</strong>, and each of their nodes contains a single digit. Add the two numbers and return the sum&nbsp;as a linked list.</p>

<p>You may assume the two numbers do not contain any leading zero, except the number 0 itself.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>
<img alt="" src="https://assets.leetcode.com/uploads/2020/10/02/addtwonumber1.jpg" style="width: 483px; height: 342px;" />
<pre>
<strong>Input:</strong> l1 = [2,4,3], l2 = [5,6,4]
<strong>Output:</strong> [7,0,8]
<strong>Explanation:</strong> 342 + 465 = 807.
</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input:</strong> l1 = [0], l2 = [0]
<strong>Output:</strong> [0]
</pre>

<p><strong class="example">Example 3:</strong></p>

<pre>
<strong>Input:</strong> l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
<strong>Output:</strong> [8,9,9,9,0,0,0,1]
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li>The number of nodes in each linked list is in the range <code>[1, 100]</code>.</li>
	<li><code>0 &lt;= Node.val &lt;= 9</code></li>
	<li>It is guaranteed that the list represents a number that does not have leading zeros.</li>
</ul>
', '[{"name": "Linked List", "slug": "linked-list", "translatedName": null}, {"name": "Math", "slug": "math", "translatedName": null}, {"name": "Recursion", "slug": "recursion", "translatedName": null}]'::jsonb, '[2,4,3]
[5,6,4]
[0]
[0]
[9,9,9,9,9,9,9]
[9,9,9,9]', '[{"title": "Multiply Strings", "titleSlug": "multiply-strings", "difficulty": "Medium", "translatedTitle": null}, {"title": "Add Binary", "titleSlug": "add-binary", "difficulty": "Easy", "translatedTitle": null}, {"title": "Sum of Two Integers", "titleSlug": "sum-of-two-integers", "difficulty": "Medium", "translatedTitle": null}, {"title": "Add Strings", "titleSlug": "add-strings", "difficulty": "Easy", "translatedTitle": null}, {"title": "Add Two Numbers II", "titleSlug": "add-two-numbers-ii", "difficulty": "Medium", "translatedTitle": null}, {"title": "Add to Array-Form of Integer", "titleSlug": "add-to-array-form-of-integer", "difficulty": "Easy", "translatedTitle": null}, {"title": "Add Two Polynomials Represented as Linked Lists", "titleSlug": "add-two-polynomials-represented-as-linked-lists", "difficulty": "Medium", "translatedTitle": null}, {"title": "Double a Number Represented as a Linked List", "titleSlug": "double-a-number-represented-as-a-linked-list", "difficulty": "Medium", "translatedTitle": null}]'::jsonb, '2025-03-05 18:16:51.178');
INSERT INTO problems
(id, frontend_id, title, title_slug, difficulty, is_paid_only, "content", topic_tags, example_testcases, similar_questions, created_at)
VALUES(3, 3, 'Longest Substring Without Repeating Characters', 'longest-substring-without-repeating-characters', 'Medium', false, '<p>Given a string <code>s</code>, find the length of the <strong>longest</strong> <span data-keyword="substring-nonempty"><strong>substring</strong></span> without duplicate characters.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<pre>
<strong>Input:</strong> s = &quot;abcabcbb&quot;
<strong>Output:</strong> 3
<strong>Explanation:</strong> The answer is &quot;abc&quot;, with the length of 3.
</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input:</strong> s = &quot;bbbbb&quot;
<strong>Output:</strong> 1
<strong>Explanation:</strong> The answer is &quot;b&quot;, with the length of 1.
</pre>

<p><strong class="example">Example 3:</strong></p>

<pre>
<strong>Input:</strong> s = &quot;pwwkew&quot;
<strong>Output:</strong> 3
<strong>Explanation:</strong> The answer is &quot;wke&quot;, with the length of 3.
Notice that the answer must be a substring, &quot;pwke&quot; is a subsequence and not a substring.
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>0 &lt;= s.length &lt;= 5 * 10<sup>4</sup></code></li>
	<li><code>s</code> consists of English letters, digits, symbols and spaces.</li>
</ul>
', '[{"name": "Hash Table", "slug": "hash-table", "translatedName": null}, {"name": "String", "slug": "string", "translatedName": null}, {"name": "Sliding Window", "slug": "sliding-window", "translatedName": null}]'::jsonb, '"abcabcbb"
"bbbbb"
"pwwkew"', '[{"title": "Longest Substring with At Most Two Distinct Characters", "titleSlug": "longest-substring-with-at-most-two-distinct-characters", "difficulty": "Medium", "translatedTitle": null}, {"title": "Longest Substring with At Most K Distinct Characters", "titleSlug": "longest-substring-with-at-most-k-distinct-characters", "difficulty": "Medium", "translatedTitle": null}, {"title": "Subarrays with K Different Integers", "titleSlug": "subarrays-with-k-different-integers", "difficulty": "Hard", "translatedTitle": null}, {"title": "Maximum Erasure Value", "titleSlug": "maximum-erasure-value", "difficulty": "Medium", "translatedTitle": null}, {"title": "Number of Equal Count Substrings", "titleSlug": "number-of-equal-count-substrings", "difficulty": "Medium", "translatedTitle": null}, {"title": "Minimum Consecutive Cards to Pick Up", "titleSlug": "minimum-consecutive-cards-to-pick-up", "difficulty": "Medium", "translatedTitle": null}, {"title": "Longest Nice Subarray", "titleSlug": "longest-nice-subarray", "difficulty": "Medium", "translatedTitle": null}, {"title": "Optimal Partition of String", "titleSlug": "optimal-partition-of-string", "difficulty": "Medium", "translatedTitle": null}, {"title": "Count Complete Subarrays in an Array", "titleSlug": "count-complete-subarrays-in-an-array", "difficulty": "Medium", "translatedTitle": null}, {"title": "Find Longest Special Substring That Occurs Thrice II", "titleSlug": "find-longest-special-substring-that-occurs-thrice-ii", "difficulty": "Medium", "translatedTitle": null}, {"title": "Find Longest Special Substring That Occurs Thrice I", "titleSlug": "find-longest-special-substring-that-occurs-thrice-i", "difficulty": "Medium", "translatedTitle": null}]'::jsonb, '2025-03-05 18:16:40.547');
INSERT INTO problems
(id, frontend_id, title, title_slug, difficulty, is_paid_only, "content", topic_tags, example_testcases, similar_questions, created_at)
VALUES(4, 4, 'Median of Two Sorted Arrays', 'median-of-two-sorted-arrays', 'Hard', false, '<p>Given two sorted arrays <code>nums1</code> and <code>nums2</code> of size <code>m</code> and <code>n</code> respectively, return <strong>the median</strong> of the two sorted arrays.</p>

<p>The overall run time complexity should be <code>O(log (m+n))</code>.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<pre>
<strong>Input:</strong> nums1 = [1,3], nums2 = [2]
<strong>Output:</strong> 2.00000
<strong>Explanation:</strong> merged array = [1,2,3] and median is 2.
</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input:</strong> nums1 = [1,2], nums2 = [3,4]
<strong>Output:</strong> 2.50000
<strong>Explanation:</strong> merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>nums1.length == m</code></li>
	<li><code>nums2.length == n</code></li>
	<li><code>0 &lt;= m &lt;= 1000</code></li>
	<li><code>0 &lt;= n &lt;= 1000</code></li>
	<li><code>1 &lt;= m + n &lt;= 2000</code></li>
	<li><code>-10<sup>6</sup> &lt;= nums1[i], nums2[i] &lt;= 10<sup>6</sup></code></li>
</ul>
', '[{"name": "Array", "slug": "array", "translatedName": null}, {"name": "Binary Search", "slug": "binary-search", "translatedName": null}, {"name": "Divide and Conquer", "slug": "divide-and-conquer", "translatedName": null}]'::jsonb, '[1,3]
[2]
[1,2]
[3,4]', '[{"title": "Median of a Row Wise Sorted Matrix", "titleSlug": "median-of-a-row-wise-sorted-matrix", "difficulty": "Medium", "translatedTitle": null}]'::jsonb, '2025-03-05 18:16:29.996');
INSERT INTO problems
(id, frontend_id, title, title_slug, difficulty, is_paid_only, "content", topic_tags, example_testcases, similar_questions, created_at)
VALUES(5, 5, 'Longest Palindromic Substring', 'longest-palindromic-substring', 'Medium', false, '<p>Given a string <code>s</code>, return <em>the longest</em> <span data-keyword="palindromic-string"><em>palindromic</em></span> <span data-keyword="substring-nonempty"><em>substring</em></span> in <code>s</code>.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<pre>
<strong>Input:</strong> s = &quot;babad&quot;
<strong>Output:</strong> &quot;bab&quot;
<strong>Explanation:</strong> &quot;aba&quot; is also a valid answer.
</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input:</strong> s = &quot;cbbd&quot;
<strong>Output:</strong> &quot;bb&quot;
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= s.length &lt;= 1000</code></li>
	<li><code>s</code> consist of only digits and English letters.</li>
</ul>
', '[{"name": "Two Pointers", "slug": "two-pointers", "translatedName": null}, {"name": "String", "slug": "string", "translatedName": null}, {"name": "Dynamic Programming", "slug": "dynamic-programming", "translatedName": null}]'::jsonb, '"babad"
"cbbd"', '[{"title": "Shortest Palindrome", "titleSlug": "shortest-palindrome", "difficulty": "Hard", "translatedTitle": null}, {"title": "Palindrome Permutation", "titleSlug": "palindrome-permutation", "difficulty": "Easy", "translatedTitle": null}, {"title": "Palindrome Pairs", "titleSlug": "palindrome-pairs", "difficulty": "Hard", "translatedTitle": null}, {"title": "Longest Palindromic Subsequence", "titleSlug": "longest-palindromic-subsequence", "difficulty": "Medium", "translatedTitle": null}, {"title": "Palindromic Substrings", "titleSlug": "palindromic-substrings", "difficulty": "Medium", "translatedTitle": null}, {"title": "Maximum Number of Non-overlapping Palindrome Substrings", "titleSlug": "maximum-number-of-non-overlapping-palindrome-substrings", "difficulty": "Hard", "translatedTitle": null}]'::jsonb, '2025-03-05 18:16:18.718');
INSERT INTO problems
(id, frontend_id, title, title_slug, difficulty, is_paid_only, "content", topic_tags, example_testcases, similar_questions, created_at)
VALUES(6, 6, 'Zigzag Conversion', 'zigzag-conversion', 'Medium', false, '<p>The string <code>&quot;PAYPALISHIRING&quot;</code> is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)</p>

<pre>
P   A   H   N
A P L S I I G
Y   I   R
</pre>

<p>And then read line by line: <code>&quot;PAHNAPLSIIGYIR&quot;</code></p>

<p>Write the code that will take a string and make this conversion given a number of rows:</p>

<pre>
string convert(string s, int numRows);
</pre>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<pre>
<strong>Input:</strong> s = &quot;PAYPALISHIRING&quot;, numRows = 3
<strong>Output:</strong> &quot;PAHNAPLSIIGYIR&quot;
</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input:</strong> s = &quot;PAYPALISHIRING&quot;, numRows = 4
<strong>Output:</strong> &quot;PINALSIGYAHRPI&quot;
<strong>Explanation:</strong>
P     I    N
A   L S  I G
Y A   H R
P     I
</pre>

<p><strong class="example">Example 3:</strong></p>

<pre>
<strong>Input:</strong> s = &quot;A&quot;, numRows = 1
<strong>Output:</strong> &quot;A&quot;
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= s.length &lt;= 1000</code></li>
	<li><code>s</code> consists of English letters (lower-case and upper-case), <code>&#39;,&#39;</code> and <code>&#39;.&#39;</code>.</li>
	<li><code>1 &lt;= numRows &lt;= 1000</code></li>
</ul>
', '[{"name": "String", "slug": "string", "translatedName": null}]'::jsonb, '"PAYPALISHIRING"
3
"PAYPALISHIRING"
4
"A"
1', '[]'::jsonb, '2025-03-05 18:16:08.119');
INSERT INTO problems
(id, frontend_id, title, title_slug, difficulty, is_paid_only, "content", topic_tags, example_testcases, similar_questions, created_at)
VALUES(7, 7, 'Reverse Integer', 'reverse-integer', 'Medium', false, '<p>Given a signed 32-bit integer <code>x</code>, return <code>x</code><em> with its digits reversed</em>. If reversing <code>x</code> causes the value to go outside the signed 32-bit integer range <code>[-2<sup>31</sup>, 2<sup>31</sup> - 1]</code>, then return <code>0</code>.</p>

<p><strong>Assume the environment does not allow you to store 64-bit integers (signed or unsigned).</strong></p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<pre>
<strong>Input:</strong> x = 123
<strong>Output:</strong> 321
</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input:</strong> x = -123
<strong>Output:</strong> -321
</pre>

<p><strong class="example">Example 3:</strong></p>

<pre>
<strong>Input:</strong> x = 120
<strong>Output:</strong> 21
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>-2<sup>31</sup> &lt;= x &lt;= 2<sup>31</sup> - 1</code></li>
</ul>
', '[{"name": "Math", "slug": "math", "translatedName": null}]'::jsonb, '123
-123
120', '[{"title": "String to Integer (atoi)", "titleSlug": "string-to-integer-atoi", "difficulty": "Medium", "translatedTitle": null}, {"title": "Reverse Bits", "titleSlug": "reverse-bits", "difficulty": "Easy", "translatedTitle": null}, {"title": "A Number After a Double Reversal", "titleSlug": "a-number-after-a-double-reversal", "difficulty": "Easy", "translatedTitle": null}, {"title": "Count Number of Distinct Integers After Reverse Operations", "titleSlug": "count-number-of-distinct-integers-after-reverse-operations", "difficulty": "Medium", "translatedTitle": null}]'::jsonb, '2025-03-05 18:15:57.637');
INSERT INTO problems
(id, frontend_id, title, title_slug, difficulty, is_paid_only, "content", topic_tags, example_testcases, similar_questions, created_at)
VALUES(8, 8, 'String to Integer (atoi)', 'string-to-integer-atoi', 'Medium', false, '<p>Implement the <code>myAtoi(string s)</code> function, which converts a string to a 32-bit signed integer.</p>

<p>The algorithm for <code>myAtoi(string s)</code> is as follows:</p>

<ol>
	<li><strong>Whitespace</strong>: Ignore any leading whitespace (<code>&quot; &quot;</code>).</li>
	<li><strong>Signedness</strong>: Determine the sign by checking if the next character is <code>&#39;-&#39;</code> or <code>&#39;+&#39;</code>, assuming positivity if neither present.</li>
	<li><strong>Conversion</strong>: Read the integer by skipping leading zeros&nbsp;until a non-digit character is encountered or the end of the string is reached. If no digits were read, then the result is 0.</li>
	<li><strong>Rounding</strong>: If the integer is out of the 32-bit signed integer range <code>[-2<sup>31</sup>, 2<sup>31</sup> - 1]</code>, then round the integer to remain in the range. Specifically, integers less than <code>-2<sup>31</sup></code> should be rounded to <code>-2<sup>31</sup></code>, and integers greater than <code>2<sup>31</sup> - 1</code> should be rounded to <code>2<sup>31</sup> - 1</code>.</li>
</ol>

<p>Return the integer as the final result.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<div class="example-block">
<p><strong>Input:</strong> <span class="example-io">s = &quot;42&quot;</span></p>

<p><strong>Output:</strong> <span class="example-io">42</span></p>

<p><strong>Explanation:</strong></p>

<pre>
The underlined characters are what is read in and the caret is the current reader position.
Step 1: &quot;42&quot; (no characters read because there is no leading whitespace)
         ^
Step 2: &quot;42&quot; (no characters read because there is neither a &#39;-&#39; nor &#39;+&#39;)
         ^
Step 3: &quot;<u>42</u>&quot; (&quot;42&quot; is read in)
           ^
</pre>
</div>

<p><strong class="example">Example 2:</strong></p>

<div class="example-block">
<p><strong>Input:</strong> <span class="example-io">s = &quot; -042&quot;</span></p>

<p><strong>Output:</strong> <span class="example-io">-42</span></p>

<p><strong>Explanation:</strong></p>

<pre>
Step 1: &quot;<u>   </u>-042&quot; (leading whitespace is read and ignored)
            ^
Step 2: &quot;   <u>-</u>042&quot; (&#39;-&#39; is read, so the result should be negative)
             ^
Step 3: &quot;   -<u>042</u>&quot; (&quot;042&quot; is read in, leading zeros ignored in the result)
               ^
</pre>
</div>

<p><strong class="example">Example 3:</strong></p>

<div class="example-block">
<p><strong>Input:</strong> <span class="example-io">s = &quot;1337c0d3&quot;</span></p>

<p><strong>Output:</strong> <span class="example-io">1337</span></p>

<p><strong>Explanation:</strong></p>

<pre>
Step 1: &quot;1337c0d3&quot; (no characters read because there is no leading whitespace)
         ^
Step 2: &quot;1337c0d3&quot; (no characters read because there is neither a &#39;-&#39; nor &#39;+&#39;)
         ^
Step 3: &quot;<u>1337</u>c0d3&quot; (&quot;1337&quot; is read in; reading stops because the next character is a non-digit)
             ^
</pre>
</div>

<p><strong class="example">Example 4:</strong></p>

<div class="example-block">
<p><strong>Input:</strong> <span class="example-io">s = &quot;0-1&quot;</span></p>

<p><strong>Output:</strong> <span class="example-io">0</span></p>

<p><strong>Explanation:</strong></p>

<pre>
Step 1: &quot;0-1&quot; (no characters read because there is no leading whitespace)
         ^
Step 2: &quot;0-1&quot; (no characters read because there is neither a &#39;-&#39; nor &#39;+&#39;)
         ^
Step 3: &quot;<u>0</u>-1&quot; (&quot;0&quot; is read in; reading stops because the next character is a non-digit)
          ^
</pre>
</div>

<p><strong class="example">Example 5:</strong></p>

<div class="example-block">
<p><strong>Input:</strong> <span class="example-io">s = &quot;words and 987&quot;</span></p>

<p><strong>Output:</strong> <span class="example-io">0</span></p>

<p><strong>Explanation:</strong></p>

<p>Reading stops at the first non-digit character &#39;w&#39;.</p>
</div>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>0 &lt;= s.length &lt;= 200</code></li>
	<li><code>s</code> consists of English letters (lower-case and upper-case), digits (<code>0-9</code>), <code>&#39; &#39;</code>, <code>&#39;+&#39;</code>, <code>&#39;-&#39;</code>, and <code>&#39;.&#39;</code>.</li>
</ul>
', '[{"name": "String", "slug": "string", "translatedName": null}]'::jsonb, '"42"
"   -042"
"1337c0d3"
"0-1"
"words and 987"', '[{"title": "Reverse Integer", "titleSlug": "reverse-integer", "difficulty": "Medium", "translatedTitle": null}, {"title": "Valid Number", "titleSlug": "valid-number", "difficulty": "Hard", "translatedTitle": null}, {"title": "Check if Numbers Are Ascending in a Sentence", "titleSlug": "check-if-numbers-are-ascending-in-a-sentence", "difficulty": "Easy", "translatedTitle": null}]'::jsonb, '2025-03-05 18:15:47.150');
INSERT INTO problems
(id, frontend_id, title, title_slug, difficulty, is_paid_only, "content", topic_tags, example_testcases, similar_questions, created_at)
VALUES(9, 9, 'Palindrome Number', 'palindrome-number', 'Easy', false, '<p>Given an integer <code>x</code>, return <code>true</code><em> if </em><code>x</code><em> is a </em><span data-keyword="palindrome-integer"><em><strong>palindrome</strong></em></span><em>, and </em><code>false</code><em> otherwise</em>.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<pre>
<strong>Input:</strong> x = 121
<strong>Output:</strong> true
<strong>Explanation:</strong> 121 reads as 121 from left to right and from right to left.
</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input:</strong> x = -121
<strong>Output:</strong> false
<strong>Explanation:</strong> From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
</pre>

<p><strong class="example">Example 3:</strong></p>

<pre>
<strong>Input:</strong> x = 10
<strong>Output:</strong> false
<strong>Explanation:</strong> Reads 01 from right to left. Therefore it is not a palindrome.
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>-2<sup>31</sup>&nbsp;&lt;= x &lt;= 2<sup>31</sup>&nbsp;- 1</code></li>
</ul>

<p>&nbsp;</p>
<strong>Follow up:</strong> Could you solve it without converting the integer to a string?', '[{"name": "Math", "slug": "math", "translatedName": null}]'::jsonb, '121
-121
10', '[{"title": "Palindrome Linked List", "titleSlug": "palindrome-linked-list", "difficulty": "Easy", "translatedTitle": null}, {"title": "Find Palindrome With Fixed Length", "titleSlug": "find-palindrome-with-fixed-length", "difficulty": "Medium", "translatedTitle": null}, {"title": "Strictly Palindromic Number", "titleSlug": "strictly-palindromic-number", "difficulty": "Medium", "translatedTitle": null}, {"title": "  Count Symmetric Integers", "titleSlug": "count-symmetric-integers", "difficulty": "Easy", "translatedTitle": null}, {"title": "Find the Count of Good Integers", "titleSlug": "find-the-count-of-good-integers", "difficulty": "Hard", "translatedTitle": null}, {"title": "Find the Largest Palindrome Divisible by K", "titleSlug": "find-the-largest-palindrome-divisible-by-k", "difficulty": "Hard", "translatedTitle": null}]'::jsonb, '2025-03-05 18:15:36.673');
INSERT INTO problems
(id, frontend_id, title, title_slug, difficulty, is_paid_only, "content", topic_tags, example_testcases, similar_questions, created_at)
VALUES(10, 10, 'Regular Expression Matching', 'regular-expression-matching', 'Hard', false, '<p>Given an input string <code>s</code>&nbsp;and a pattern <code>p</code>, implement regular expression matching with support for <code>&#39;.&#39;</code> and <code>&#39;*&#39;</code> where:</p>

<ul>
	<li><code>&#39;.&#39;</code> Matches any single character.​​​​</li>
	<li><code>&#39;*&#39;</code> Matches zero or more of the preceding element.</li>
</ul>

<p>The matching should cover the <strong>entire</strong> input string (not partial).</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<pre>
<strong>Input:</strong> s = &quot;aa&quot;, p = &quot;a&quot;
<strong>Output:</strong> false
<strong>Explanation:</strong> &quot;a&quot; does not match the entire string &quot;aa&quot;.
</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input:</strong> s = &quot;aa&quot;, p = &quot;a*&quot;
<strong>Output:</strong> true
<strong>Explanation:</strong> &#39;*&#39; means zero or more of the preceding element, &#39;a&#39;. Therefore, by repeating &#39;a&#39; once, it becomes &quot;aa&quot;.
</pre>

<p><strong class="example">Example 3:</strong></p>

<pre>
<strong>Input:</strong> s = &quot;ab&quot;, p = &quot;.*&quot;
<strong>Output:</strong> true
<strong>Explanation:</strong> &quot;.*&quot; means &quot;zero or more (*) of any character (.)&quot;.
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= s.length&nbsp;&lt;= 20</code></li>
	<li><code>1 &lt;= p.length&nbsp;&lt;= 20</code></li>
	<li><code>s</code> contains only lowercase English letters.</li>
	<li><code>p</code> contains only lowercase English letters, <code>&#39;.&#39;</code>, and&nbsp;<code>&#39;*&#39;</code>.</li>
	<li>It is guaranteed for each appearance of the character <code>&#39;*&#39;</code>, there will be a previous valid character to match.</li>
</ul>
', '[{"name": "String", "slug": "string", "translatedName": null}, {"name": "Dynamic Programming", "slug": "dynamic-programming", "translatedName": null}, {"name": "Recursion", "slug": "recursion", "translatedName": null}]'::jsonb, '"aa"
"a"
"aa"
"a*"
"ab"
".*"', '[{"title": "Wildcard Matching", "titleSlug": "wildcard-matching", "difficulty": "Hard", "translatedTitle": null}]'::jsonb, '2025-03-05 18:15:26.195');
INSERT INTO problems
(id, frontend_id, title, title_slug, difficulty, is_paid_only, "content", topic_tags, example_testcases, similar_questions, created_at)
VALUES(11, 11, 'Container With Most Water', 'container-with-most-water', 'Medium', false, '<p>You are given an integer array <code>height</code> of length <code>n</code>. There are <code>n</code> vertical lines drawn such that the two endpoints of the <code>i<sup>th</sup></code> line are <code>(i, 0)</code> and <code>(i, height[i])</code>.</p>

<p>Find two lines that together with the x-axis form a container, such that the container contains the most water.</p>

<p>Return <em>the maximum amount of water a container can store</em>.</p>

<p><strong>Notice</strong> that you may not slant the container.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>
<img alt="" src="https://s3-lc-upload.s3.amazonaws.com/uploads/2018/07/17/question_11.jpg" style="width: 600px; height: 287px;" />
<pre>
<strong>Input:</strong> height = [1,8,6,2,5,4,8,3,7]
<strong>Output:</strong> 49
<strong>Explanation:</strong> The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.
</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input:</strong> height = [1,1]
<strong>Output:</strong> 1
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>n == height.length</code></li>
	<li><code>2 &lt;= n &lt;= 10<sup>5</sup></code></li>
	<li><code>0 &lt;= height[i] &lt;= 10<sup>4</sup></code></li>
</ul>
', '[{"name": "Array", "slug": "array", "translatedName": null}, {"name": "Two Pointers", "slug": "two-pointers", "translatedName": null}, {"name": "Greedy", "slug": "greedy", "translatedName": null}]'::jsonb, '[1,8,6,2,5,4,8,3,7]
[1,1]', '[{"title": "Trapping Rain Water", "titleSlug": "trapping-rain-water", "difficulty": "Hard", "translatedTitle": null}, {"title": "Maximum Tastiness of Candy Basket", "titleSlug": "maximum-tastiness-of-candy-basket", "difficulty": "Medium", "translatedTitle": null}, {"title": "House Robber IV", "titleSlug": "house-robber-iv", "difficulty": "Medium", "translatedTitle": null}]'::jsonb, '2025-03-05 18:15:15.751');
INSERT INTO problems
(id, frontend_id, title, title_slug, difficulty, is_paid_only, "content", topic_tags, example_testcases, similar_questions, created_at)
VALUES(12, 12, 'Integer to Roman', 'integer-to-roman', 'Medium', false, '<p>Seven different symbols represent Roman numerals with the following values:</p>

<table>
	<thead>
		<tr>
			<th>Symbol</th>
			<th>Value</th>
		</tr>
	</thead>
	<tbody>
		<tr>
			<td>I</td>
			<td>1</td>
		</tr>
		<tr>
			<td>V</td>
			<td>5</td>
		</tr>
		<tr>
			<td>X</td>
			<td>10</td>
		</tr>
		<tr>
			<td>L</td>
			<td>50</td>
		</tr>
		<tr>
			<td>C</td>
			<td>100</td>
		</tr>
		<tr>
			<td>D</td>
			<td>500</td>
		</tr>
		<tr>
			<td>M</td>
			<td>1000</td>
		</tr>
	</tbody>
</table>

<p>Roman numerals are formed by appending&nbsp;the conversions of&nbsp;decimal place values&nbsp;from highest to lowest. Converting a decimal place value into a Roman numeral has the following rules:</p>

<ul>
	<li>If the value does not start with 4 or&nbsp;9, select the symbol of the maximal value that can be subtracted from the input, append that symbol to the result, subtract its value, and convert the remainder to a Roman numeral.</li>
	<li>If the value starts with 4 or 9 use the&nbsp;<strong>subtractive form</strong>&nbsp;representing&nbsp;one symbol subtracted from the following symbol, for example,&nbsp;4 is 1 (<code>I</code>) less than 5 (<code>V</code>): <code>IV</code>&nbsp;and 9 is 1 (<code>I</code>) less than 10 (<code>X</code>): <code>IX</code>.&nbsp;Only the following subtractive forms are used: 4 (<code>IV</code>), 9 (<code>IX</code>),&nbsp;40 (<code>XL</code>), 90 (<code>XC</code>), 400 (<code>CD</code>) and 900 (<code>CM</code>).</li>
	<li>Only powers of 10 (<code>I</code>, <code>X</code>, <code>C</code>, <code>M</code>) can be appended consecutively at most 3 times to represent multiples of 10. You cannot append 5&nbsp;(<code>V</code>), 50 (<code>L</code>), or 500 (<code>D</code>) multiple times. If you need to append a symbol&nbsp;4 times&nbsp;use the <strong>subtractive form</strong>.</li>
</ul>

<p>Given an integer, convert it to a Roman numeral.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<div class="example-block">
<p><strong>Input:</strong> <span class="example-io">num = 3749</span></p>

<p><strong>Output:</strong> <span class="example-io">&quot;MMMDCCXLIX&quot;</span></p>

<p><strong>Explanation:</strong></p>

<pre>
3000 = MMM as 1000 (M) + 1000 (M) + 1000 (M)
 700 = DCC as 500 (D) + 100 (C) + 100 (C)
  40 = XL as 10 (X) less of 50 (L)
   9 = IX as 1 (I) less of 10 (X)
Note: 49 is not 1 (I) less of 50 (L) because the conversion is based on decimal places
</pre>
</div>

<p><strong class="example">Example 2:</strong></p>

<div class="example-block">
<p><strong>Input:</strong> <span class="example-io">num = 58</span></p>

<p><strong>Output:</strong> <span class="example-io">&quot;LVIII&quot;</span></p>

<p><strong>Explanation:</strong></p>

<pre>
50 = L
 8 = VIII
</pre>
</div>

<p><strong class="example">Example 3:</strong></p>

<div class="example-block">
<p><strong>Input:</strong> <span class="example-io">num = 1994</span></p>

<p><strong>Output:</strong> <span class="example-io">&quot;MCMXCIV&quot;</span></p>

<p><strong>Explanation:</strong></p>

<pre>
1000 = M
 900 = CM
  90 = XC
   4 = IV
</pre>
</div>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= num &lt;= 3999</code></li>
</ul>
', '[{"name": "Hash Table", "slug": "hash-table", "translatedName": null}, {"name": "Math", "slug": "math", "translatedName": null}, {"name": "String", "slug": "string", "translatedName": null}]'::jsonb, '3749
58
1994', '[{"title": "Roman to Integer", "titleSlug": "roman-to-integer", "difficulty": "Easy", "translatedTitle": null}, {"title": "Integer to English Words", "titleSlug": "integer-to-english-words", "difficulty": "Hard", "translatedTitle": null}]'::jsonb, '2025-03-05 18:15:05.204');
INSERT INTO problems
(id, frontend_id, title, title_slug, difficulty, is_paid_only, "content", topic_tags, example_testcases, similar_questions, created_at)
VALUES(13, 13, 'Roman to Integer', 'roman-to-integer', 'Easy', false, '<p>Roman numerals are represented by seven different symbols:&nbsp;<code>I</code>, <code>V</code>, <code>X</code>, <code>L</code>, <code>C</code>, <code>D</code> and <code>M</code>.</p>

<pre>
<strong>Symbol</strong>       <strong>Value</strong>
I             1
V             5
X             10
L             50
C             100
D             500
M             1000</pre>

<p>For example,&nbsp;<code>2</code> is written as <code>II</code>&nbsp;in Roman numeral, just two ones added together. <code>12</code> is written as&nbsp;<code>XII</code>, which is simply <code>X + II</code>. The number <code>27</code> is written as <code>XXVII</code>, which is <code>XX + V + II</code>.</p>

<p>Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not <code>IIII</code>. Instead, the number four is written as <code>IV</code>. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as <code>IX</code>. There are six instances where subtraction is used:</p>

<ul>
	<li><code>I</code> can be placed before <code>V</code> (5) and <code>X</code> (10) to make 4 and 9.&nbsp;</li>
	<li><code>X</code> can be placed before <code>L</code> (50) and <code>C</code> (100) to make 40 and 90.&nbsp;</li>
	<li><code>C</code> can be placed before <code>D</code> (500) and <code>M</code> (1000) to make 400 and 900.</li>
</ul>

<p>Given a roman numeral, convert it to an integer.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<pre>
<strong>Input:</strong> s = &quot;III&quot;
<strong>Output:</strong> 3
<strong>Explanation:</strong> III = 3.
</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input:</strong> s = &quot;LVIII&quot;
<strong>Output:</strong> 58
<strong>Explanation:</strong> L = 50, V= 5, III = 3.
</pre>

<p><strong class="example">Example 3:</strong></p>

<pre>
<strong>Input:</strong> s = &quot;MCMXCIV&quot;
<strong>Output:</strong> 1994
<strong>Explanation:</strong> M = 1000, CM = 900, XC = 90 and IV = 4.
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= s.length &lt;= 15</code></li>
	<li><code>s</code> contains only&nbsp;the characters <code>(&#39;I&#39;, &#39;V&#39;, &#39;X&#39;, &#39;L&#39;, &#39;C&#39;, &#39;D&#39;, &#39;M&#39;)</code>.</li>
	<li>It is <strong>guaranteed</strong>&nbsp;that <code>s</code> is a valid roman numeral in the range <code>[1, 3999]</code>.</li>
</ul>
', '[{"name": "Hash Table", "slug": "hash-table", "translatedName": null}, {"name": "Math", "slug": "math", "translatedName": null}, {"name": "String", "slug": "string", "translatedName": null}]'::jsonb, '"III"
"LVIII"
"MCMXCIV"', '[{"title": "Integer to Roman", "titleSlug": "integer-to-roman", "difficulty": "Medium", "translatedTitle": null}]'::jsonb, '2025-03-05 18:14:54.751');
INSERT INTO problems
(id, frontend_id, title, title_slug, difficulty, is_paid_only, "content", topic_tags, example_testcases, similar_questions, created_at)
VALUES(14, 14, 'Longest Common Prefix', 'longest-common-prefix', 'Easy', false, '<p>Write a function to find the longest common prefix string amongst an array of strings.</p>

<p>If there is no common prefix, return an empty string <code>&quot;&quot;</code>.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<pre>
<strong>Input:</strong> strs = [&quot;flower&quot;,&quot;flow&quot;,&quot;flight&quot;]
<strong>Output:</strong> &quot;fl&quot;
</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input:</strong> strs = [&quot;dog&quot;,&quot;racecar&quot;,&quot;car&quot;]
<strong>Output:</strong> &quot;&quot;
<strong>Explanation:</strong> There is no common prefix among the input strings.
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= strs.length &lt;= 200</code></li>
	<li><code>0 &lt;= strs[i].length &lt;= 200</code></li>
	<li><code>strs[i]</code> consists of only lowercase English letters if it is non-empty.</li>
</ul>
', '[{"name": "String", "slug": "string", "translatedName": null}, {"name": "Trie", "slug": "trie", "translatedName": null}]'::jsonb, '["flower","flow","flight"]
["dog","racecar","car"]', '[{"title": "Smallest Missing Integer Greater Than Sequential Prefix Sum", "titleSlug": "smallest-missing-integer-greater-than-sequential-prefix-sum", "difficulty": "Easy", "translatedTitle": null}, {"title": "Find the Length of the Longest Common Prefix", "titleSlug": "find-the-length-of-the-longest-common-prefix", "difficulty": "Medium", "translatedTitle": null}, {"title": "Longest Common Suffix Queries", "titleSlug": "longest-common-suffix-queries", "difficulty": "Hard", "translatedTitle": null}, {"title": "Longest Common Prefix After at Most One Removal", "titleSlug": "longest-common-prefix-after-at-most-one-removal", "difficulty": "Medium", "translatedTitle": null}]'::jsonb, '2025-03-05 18:14:44.230');
INSERT INTO problems
(id, frontend_id, title, title_slug, difficulty, is_paid_only, "content", topic_tags, example_testcases, similar_questions, created_at)
VALUES(15, 15, '3Sum', '3sum', 'Medium', false, '<p>Given an integer array nums, return all the triplets <code>[nums[i], nums[j], nums[k]]</code> such that <code>i != j</code>, <code>i != k</code>, and <code>j != k</code>, and <code>nums[i] + nums[j] + nums[k] == 0</code>.</p>

<p>Notice that the solution set must not contain duplicate triplets.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<pre>
<strong>Input:</strong> nums = [-1,0,1,2,-1,-4]
<strong>Output:</strong> [[-1,-1,2],[-1,0,1]]
<strong>Explanation:</strong> 
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
The distinct triplets are [-1,0,1] and [-1,-1,2].
Notice that the order of the output and the order of the triplets does not matter.
</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input:</strong> nums = [0,1,1]
<strong>Output:</strong> []
<strong>Explanation:</strong> The only possible triplet does not sum up to 0.
</pre>

<p><strong class="example">Example 3:</strong></p>

<pre>
<strong>Input:</strong> nums = [0,0,0]
<strong>Output:</strong> [[0,0,0]]
<strong>Explanation:</strong> The only possible triplet sums up to 0.
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>3 &lt;= nums.length &lt;= 3000</code></li>
	<li><code>-10<sup>5</sup> &lt;= nums[i] &lt;= 10<sup>5</sup></code></li>
</ul>
', '[{"name": "Array", "slug": "array", "translatedName": null}, {"name": "Two Pointers", "slug": "two-pointers", "translatedName": null}, {"name": "Sorting", "slug": "sorting", "translatedName": null}]'::jsonb, '[-1,0,1,2,-1,-4]
[0,1,1]
[0,0,0]', '[{"title": "Two Sum", "titleSlug": "two-sum", "difficulty": "Easy", "translatedTitle": null}, {"title": "3Sum Closest", "titleSlug": "3sum-closest", "difficulty": "Medium", "translatedTitle": null}, {"title": "4Sum", "titleSlug": "4sum", "difficulty": "Medium", "translatedTitle": null}, {"title": "3Sum Smaller", "titleSlug": "3sum-smaller", "difficulty": "Medium", "translatedTitle": null}, {"title": "Number of Arithmetic Triplets", "titleSlug": "number-of-arithmetic-triplets", "difficulty": "Easy", "translatedTitle": null}, {"title": "Minimum Sum of Mountain Triplets I", "titleSlug": "minimum-sum-of-mountain-triplets-i", "difficulty": "Easy", "translatedTitle": null}, {"title": "Minimum Sum of Mountain Triplets II", "titleSlug": "minimum-sum-of-mountain-triplets-ii", "difficulty": "Medium", "translatedTitle": null}]'::jsonb, '2025-03-05 18:14:33.743');
INSERT INTO problems
(id, frontend_id, title, title_slug, difficulty, is_paid_only, "content", topic_tags, example_testcases, similar_questions, created_at)
VALUES(16, 16, '3Sum Closest', '3sum-closest', 'Medium', false, '<p>Given an integer array <code>nums</code> of length <code>n</code> and an integer <code>target</code>, find three integers in <code>nums</code> such that the sum is closest to <code>target</code>.</p>

<p>Return <em>the sum of the three integers</em>.</p>

<p>You may assume that each input would have exactly one solution.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<pre>
<strong>Input:</strong> nums = [-1,2,1,-4], target = 1
<strong>Output:</strong> 2
<strong>Explanation:</strong> The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input:</strong> nums = [0,0,0], target = 1
<strong>Output:</strong> 0
<strong>Explanation:</strong> The sum that is closest to the target is 0. (0 + 0 + 0 = 0).
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>3 &lt;= nums.length &lt;= 500</code></li>
	<li><code>-1000 &lt;= nums[i] &lt;= 1000</code></li>
	<li><code>-10<sup>4</sup> &lt;= target &lt;= 10<sup>4</sup></code></li>
</ul>
', '[{"name": "Array", "slug": "array", "translatedName": null}, {"name": "Two Pointers", "slug": "two-pointers", "translatedName": null}, {"name": "Sorting", "slug": "sorting", "translatedName": null}]'::jsonb, '[-1,2,1,-4]
1
[0,0,0]
1', '[{"title": "3Sum", "titleSlug": "3sum", "difficulty": "Medium", "translatedTitle": null}, {"title": "3Sum Smaller", "titleSlug": "3sum-smaller", "difficulty": "Medium", "translatedTitle": null}]'::jsonb, '2025-03-05 18:14:23.263');
INSERT INTO problems
(id, frontend_id, title, title_slug, difficulty, is_paid_only, "content", topic_tags, example_testcases, similar_questions, created_at)
VALUES(17, 17, 'Letter Combinations of a Phone Number', 'letter-combinations-of-a-phone-number', 'Medium', false, '<p>Given a string containing digits from <code>2-9</code> inclusive, return all possible letter combinations that the number could represent. Return the answer in <strong>any order</strong>.</p>

<p>A mapping of digits to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.</p>
<img alt="" src="https://assets.leetcode.com/uploads/2022/03/15/1200px-telephone-keypad2svg.png" style="width: 300px; height: 243px;" />
<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<pre>
<strong>Input:</strong> digits = &quot;23&quot;
<strong>Output:</strong> [&quot;ad&quot;,&quot;ae&quot;,&quot;af&quot;,&quot;bd&quot;,&quot;be&quot;,&quot;bf&quot;,&quot;cd&quot;,&quot;ce&quot;,&quot;cf&quot;]
</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input:</strong> digits = &quot;&quot;
<strong>Output:</strong> []
</pre>

<p><strong class="example">Example 3:</strong></p>

<pre>
<strong>Input:</strong> digits = &quot;2&quot;
<strong>Output:</strong> [&quot;a&quot;,&quot;b&quot;,&quot;c&quot;]
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>0 &lt;= digits.length &lt;= 4</code></li>
	<li><code>digits[i]</code> is a digit in the range <code>[&#39;2&#39;, &#39;9&#39;]</code>.</li>
</ul>
', '[{"name": "Hash Table", "slug": "hash-table", "translatedName": null}, {"name": "String", "slug": "string", "translatedName": null}, {"name": "Backtracking", "slug": "backtracking", "translatedName": null}]'::jsonb, '"23"
""
"2"', '[{"title": "Generate Parentheses", "titleSlug": "generate-parentheses", "difficulty": "Medium", "translatedTitle": null}, {"title": "Combination Sum", "titleSlug": "combination-sum", "difficulty": "Medium", "translatedTitle": null}, {"title": "Binary Watch", "titleSlug": "binary-watch", "difficulty": "Easy", "translatedTitle": null}, {"title": "Count Number of Texts", "titleSlug": "count-number-of-texts", "difficulty": "Medium", "translatedTitle": null}, {"title": "Minimum Number of Pushes to Type Word I", "titleSlug": "minimum-number-of-pushes-to-type-word-i", "difficulty": "Easy", "translatedTitle": null}, {"title": "Minimum Number of Pushes to Type Word II", "titleSlug": "minimum-number-of-pushes-to-type-word-ii", "difficulty": "Medium", "translatedTitle": null}]'::jsonb, '2025-03-05 18:14:12.465');
INSERT INTO problems
(id, frontend_id, title, title_slug, difficulty, is_paid_only, "content", topic_tags, example_testcases, similar_questions, created_at)
VALUES(18, 18, '4Sum', '4sum', 'Medium', false, '<p>Given an array <code>nums</code> of <code>n</code> integers, return <em>an array of all the <strong>unique</strong> quadruplets</em> <code>[nums[a], nums[b], nums[c], nums[d]]</code> such that:</p>

<ul>
	<li><code>0 &lt;= a, b, c, d&nbsp;&lt; n</code></li>
	<li><code>a</code>, <code>b</code>, <code>c</code>, and <code>d</code> are <strong>distinct</strong>.</li>
	<li><code>nums[a] + nums[b] + nums[c] + nums[d] == target</code></li>
</ul>

<p>You may return the answer in <strong>any order</strong>.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<pre>
<strong>Input:</strong> nums = [1,0,-1,0,-2,2], target = 0
<strong>Output:</strong> [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input:</strong> nums = [2,2,2,2,2], target = 8
<strong>Output:</strong> [[2,2,2,2]]
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 200</code></li>
	<li><code>-10<sup>9</sup> &lt;= nums[i] &lt;= 10<sup>9</sup></code></li>
	<li><code>-10<sup>9</sup> &lt;= target &lt;= 10<sup>9</sup></code></li>
</ul>
', '[{"name": "Array", "slug": "array", "translatedName": null}, {"name": "Two Pointers", "slug": "two-pointers", "translatedName": null}, {"name": "Sorting", "slug": "sorting", "translatedName": null}]'::jsonb, '[1,0,-1,0,-2,2]
0
[2,2,2,2,2]
8', '[{"title": "Two Sum", "titleSlug": "two-sum", "difficulty": "Easy", "translatedTitle": null}, {"title": "3Sum", "titleSlug": "3sum", "difficulty": "Medium", "translatedTitle": null}, {"title": "4Sum II", "titleSlug": "4sum-ii", "difficulty": "Medium", "translatedTitle": null}, {"title": "Count Special Quadruplets", "titleSlug": "count-special-quadruplets", "difficulty": "Easy", "translatedTitle": null}]'::jsonb, '2025-03-05 18:14:01.535');
INSERT INTO problems
(id, frontend_id, title, title_slug, difficulty, is_paid_only, "content", topic_tags, example_testcases, similar_questions, created_at)
VALUES(19, 19, 'Remove Nth Node From End of List', 'remove-nth-node-from-end-of-list', 'Medium', false, '<p>Given the <code>head</code> of a linked list, remove the <code>n<sup>th</sup></code> node from the end of the list and return its head.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>
<img alt="" src="https://assets.leetcode.com/uploads/2020/10/03/remove_ex1.jpg" style="width: 542px; height: 222px;" />
<pre>
<strong>Input:</strong> head = [1,2,3,4,5], n = 2
<strong>Output:</strong> [1,2,3,5]
</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input:</strong> head = [1], n = 1
<strong>Output:</strong> []
</pre>

<p><strong class="example">Example 3:</strong></p>

<pre>
<strong>Input:</strong> head = [1,2], n = 1
<strong>Output:</strong> [1]
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li>The number of nodes in the list is <code>sz</code>.</li>
	<li><code>1 &lt;= sz &lt;= 30</code></li>
	<li><code>0 &lt;= Node.val &lt;= 100</code></li>
	<li><code>1 &lt;= n &lt;= sz</code></li>
</ul>

<p>&nbsp;</p>
<p><strong>Follow up:</strong> Could you do this in one pass?</p>
', '[{"name": "Linked List", "slug": "linked-list", "translatedName": null}, {"name": "Two Pointers", "slug": "two-pointers", "translatedName": null}]'::jsonb, '[1,2,3,4,5]
2
[1]
1
[1,2]
1', '[{"title": "Swapping Nodes in a Linked List", "titleSlug": "swapping-nodes-in-a-linked-list", "difficulty": "Medium", "translatedTitle": null}, {"title": "Delete N Nodes After M Nodes of a Linked List", "titleSlug": "delete-n-nodes-after-m-nodes-of-a-linked-list", "difficulty": "Easy", "translatedTitle": null}, {"title": "Delete the Middle Node of a Linked List", "titleSlug": "delete-the-middle-node-of-a-linked-list", "difficulty": "Medium", "translatedTitle": null}]'::jsonb, '2025-03-05 18:13:50.771');
INSERT INTO problems
(id, frontend_id, title, title_slug, difficulty, is_paid_only, "content", topic_tags, example_testcases, similar_questions, created_at)
VALUES(20, 20, 'Valid Parentheses', 'valid-parentheses', 'Easy', false, '<p>Given a string <code>s</code> containing just the characters <code>&#39;(&#39;</code>, <code>&#39;)&#39;</code>, <code>&#39;{&#39;</code>, <code>&#39;}&#39;</code>, <code>&#39;[&#39;</code> and <code>&#39;]&#39;</code>, determine if the input string is valid.</p>

<p>An input string is valid if:</p>

<ol>
	<li>Open brackets must be closed by the same type of brackets.</li>
	<li>Open brackets must be closed in the correct order.</li>
	<li>Every close bracket has a corresponding open bracket of the same type.</li>
</ol>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<div class="example-block">
<p><strong>Input:</strong> <span class="example-io">s = &quot;()&quot;</span></p>

<p><strong>Output:</strong> <span class="example-io">true</span></p>
</div>

<p><strong class="example">Example 2:</strong></p>

<div class="example-block">
<p><strong>Input:</strong> <span class="example-io">s = &quot;()[]{}&quot;</span></p>

<p><strong>Output:</strong> <span class="example-io">true</span></p>
</div>

<p><strong class="example">Example 3:</strong></p>

<div class="example-block">
<p><strong>Input:</strong> <span class="example-io">s = &quot;(]&quot;</span></p>

<p><strong>Output:</strong> <span class="example-io">false</span></p>
</div>

<p><strong class="example">Example 4:</strong></p>

<div class="example-block">
<p><strong>Input:</strong> <span class="example-io">s = &quot;([])&quot;</span></p>

<p><strong>Output:</strong> <span class="example-io">true</span></p>
</div>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= s.length &lt;= 10<sup>4</sup></code></li>
	<li><code>s</code> consists of parentheses only <code>&#39;()[]{}&#39;</code>.</li>
</ul>
', '[{"name": "String", "slug": "string", "translatedName": null}, {"name": "Stack", "slug": "stack", "translatedName": null}]'::jsonb, '"()"
"()[]{}"
"(]"
"([])"', '[{"title": "Generate Parentheses", "titleSlug": "generate-parentheses", "difficulty": "Medium", "translatedTitle": null}, {"title": "Longest Valid Parentheses", "titleSlug": "longest-valid-parentheses", "difficulty": "Hard", "translatedTitle": null}, {"title": "Remove Invalid Parentheses", "titleSlug": "remove-invalid-parentheses", "difficulty": "Hard", "translatedTitle": null}, {"title": "Check If Word Is Valid After Substitutions", "titleSlug": "check-if-word-is-valid-after-substitutions", "difficulty": "Medium", "translatedTitle": null}, {"title": "Check if a Parentheses String Can Be Valid", "titleSlug": "check-if-a-parentheses-string-can-be-valid", "difficulty": "Medium", "translatedTitle": null}, {"title": "Move Pieces to Obtain a String", "titleSlug": "move-pieces-to-obtain-a-string", "difficulty": "Medium", "translatedTitle": null}]'::jsonb, '2025-03-05 18:13:40.280');
INSERT INTO problems
(id, frontend_id, title, title_slug, difficulty, is_paid_only, "content", topic_tags, example_testcases, similar_questions, created_at)
VALUES(21, 21, 'Merge Two Sorted Lists', 'merge-two-sorted-lists', 'Easy', false, '<p>You are given the heads of two sorted linked lists <code>list1</code> and <code>list2</code>.</p>

<p>Merge the two lists into one <strong>sorted</strong> list. The list should be made by splicing together the nodes of the first two lists.</p>

<p>Return <em>the head of the merged linked list</em>.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>
<img alt="" src="https://assets.leetcode.com/uploads/2020/10/03/merge_ex1.jpg" style="width: 662px; height: 302px;" />
<pre>
<strong>Input:</strong> list1 = [1,2,4], list2 = [1,3,4]
<strong>Output:</strong> [1,1,2,3,4,4]
</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input:</strong> list1 = [], list2 = []
<strong>Output:</strong> []
</pre>

<p><strong class="example">Example 3:</strong></p>

<pre>
<strong>Input:</strong> list1 = [], list2 = [0]
<strong>Output:</strong> [0]
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li>The number of nodes in both lists is in the range <code>[0, 50]</code>.</li>
	<li><code>-100 &lt;= Node.val &lt;= 100</code></li>
	<li>Both <code>list1</code> and <code>list2</code> are sorted in <strong>non-decreasing</strong> order.</li>
</ul>
', '[{"name": "Linked List", "slug": "linked-list", "translatedName": null}, {"name": "Recursion", "slug": "recursion", "translatedName": null}]'::jsonb, '[1,2,4]
[1,3,4]
[]
[]
[]
[0]', '[{"title": "Merge k Sorted Lists", "titleSlug": "merge-k-sorted-lists", "difficulty": "Hard", "translatedTitle": null}, {"title": "Merge Sorted Array", "titleSlug": "merge-sorted-array", "difficulty": "Easy", "translatedTitle": null}, {"title": "Sort List", "titleSlug": "sort-list", "difficulty": "Medium", "translatedTitle": null}, {"title": "Shortest Word Distance II", "titleSlug": "shortest-word-distance-ii", "difficulty": "Medium", "translatedTitle": null}, {"title": "Add Two Polynomials Represented as Linked Lists", "titleSlug": "add-two-polynomials-represented-as-linked-lists", "difficulty": "Medium", "translatedTitle": null}, {"title": "Longest Common Subsequence Between Sorted Arrays", "titleSlug": "longest-common-subsequence-between-sorted-arrays", "difficulty": "Medium", "translatedTitle": null}, {"title": "Merge Two 2D Arrays by Summing Values", "titleSlug": "merge-two-2d-arrays-by-summing-values", "difficulty": "Easy", "translatedTitle": null}]'::jsonb, '2025-03-05 18:13:29.351');
INSERT INTO problems
(id, frontend_id, title, title_slug, difficulty, is_paid_only, "content", topic_tags, example_testcases, similar_questions, created_at)
VALUES(22, 22, 'Generate Parentheses', 'generate-parentheses', 'Medium', false, '<p>Given <code>n</code> pairs of parentheses, write a function to <em>generate all combinations of well-formed parentheses</em>.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>
<pre><strong>Input:</strong> n = 3
<strong>Output:</strong> ["((()))","(()())","(())()","()(())","()()()"]
</pre><p><strong class="example">Example 2:</strong></p>
<pre><strong>Input:</strong> n = 1
<strong>Output:</strong> ["()"]
</pre>
<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= n &lt;= 8</code></li>
</ul>
', '[{"name": "String", "slug": "string", "translatedName": null}, {"name": "Dynamic Programming", "slug": "dynamic-programming", "translatedName": null}, {"name": "Backtracking", "slug": "backtracking", "translatedName": null}]'::jsonb, '3
1', '[{"title": "Letter Combinations of a Phone Number", "titleSlug": "letter-combinations-of-a-phone-number", "difficulty": "Medium", "translatedTitle": null}, {"title": "Valid Parentheses", "titleSlug": "valid-parentheses", "difficulty": "Easy", "translatedTitle": null}, {"title": "Check if a Parentheses String Can Be Valid", "titleSlug": "check-if-a-parentheses-string-can-be-valid", "difficulty": "Medium", "translatedTitle": null}]'::jsonb, '2025-03-05 18:13:18.923');
INSERT INTO problems
(id, frontend_id, title, title_slug, difficulty, is_paid_only, "content", topic_tags, example_testcases, similar_questions, created_at)
VALUES(23, 23, 'Merge k Sorted Lists', 'merge-k-sorted-lists', 'Hard', false, '<p>You are given an array of <code>k</code> linked-lists <code>lists</code>, each linked-list is sorted in ascending order.</p>

<p><em>Merge all the linked-lists into one sorted linked-list and return it.</em></p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<pre>
<strong>Input:</strong> lists = [[1,4,5],[1,3,4],[2,6]]
<strong>Output:</strong> [1,1,2,3,4,4,5,6]
<strong>Explanation:</strong> The linked-lists are:
[
  1-&gt;4-&gt;5,
  1-&gt;3-&gt;4,
  2-&gt;6
]
merging them into one sorted list:
1-&gt;1-&gt;2-&gt;3-&gt;4-&gt;4-&gt;5-&gt;6
</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input:</strong> lists = []
<strong>Output:</strong> []
</pre>

<p><strong class="example">Example 3:</strong></p>

<pre>
<strong>Input:</strong> lists = [[]]
<strong>Output:</strong> []
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>k == lists.length</code></li>
	<li><code>0 &lt;= k &lt;= 10<sup>4</sup></code></li>
	<li><code>0 &lt;= lists[i].length &lt;= 500</code></li>
	<li><code>-10<sup>4</sup> &lt;= lists[i][j] &lt;= 10<sup>4</sup></code></li>
	<li><code>lists[i]</code> is sorted in <strong>ascending order</strong>.</li>
	<li>The sum of <code>lists[i].length</code> will not exceed <code>10<sup>4</sup></code>.</li>
</ul>
', '[{"name": "Linked List", "slug": "linked-list", "translatedName": null}, {"name": "Divide and Conquer", "slug": "divide-and-conquer", "translatedName": null}, {"name": "Heap (Priority Queue)", "slug": "heap-priority-queue", "translatedName": null}, {"name": "Merge Sort", "slug": "merge-sort", "translatedName": null}]'::jsonb, '[[1,4,5],[1,3,4],[2,6]]
[]
[[]]', '[{"title": "Merge Two Sorted Lists", "titleSlug": "merge-two-sorted-lists", "difficulty": "Easy", "translatedTitle": null}, {"title": "Ugly Number II", "titleSlug": "ugly-number-ii", "difficulty": "Medium", "translatedTitle": null}, {"title": "Smallest Subarrays With Maximum Bitwise OR", "titleSlug": "smallest-subarrays-with-maximum-bitwise-or", "difficulty": "Medium", "translatedTitle": null}]'::jsonb, '2025-03-05 18:13:08.172');
INSERT INTO problems
(id, frontend_id, title, title_slug, difficulty, is_paid_only, "content", topic_tags, example_testcases, similar_questions, created_at)
VALUES(24, 24, 'Swap Nodes in Pairs', 'swap-nodes-in-pairs', 'Medium', false, '<p>Given a&nbsp;linked list, swap every two adjacent nodes and return its head. You must solve the problem without&nbsp;modifying the values in the list&#39;s nodes (i.e., only nodes themselves may be changed.)</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<div class="example-block">
<p><strong>Input:</strong> <span class="example-io">head = [1,2,3,4]</span></p>

<p><strong>Output:</strong> <span class="example-io">[2,1,4,3]</span></p>

<p><strong>Explanation:</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2020/10/03/swap_ex1.jpg" style="width: 422px; height: 222px;" /></p>
</div>

<p><strong class="example">Example 2:</strong></p>

<div class="example-block">
<p><strong>Input:</strong> <span class="example-io">head = []</span></p>

<p><strong>Output:</strong> <span class="example-io">[]</span></p>
</div>

<p><strong class="example">Example 3:</strong></p>

<div class="example-block">
<p><strong>Input:</strong> <span class="example-io">head = [1]</span></p>

<p><strong>Output:</strong> <span class="example-io">[1]</span></p>
</div>

<p><strong class="example">Example 4:</strong></p>

<div class="example-block">
<p><strong>Input:</strong> <span class="example-io">head = [1,2,3]</span></p>

<p><strong>Output:</strong> <span class="example-io">[2,1,3]</span></p>
</div>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li>The number of nodes in the&nbsp;list&nbsp;is in the range <code>[0, 100]</code>.</li>
	<li><code>0 &lt;= Node.val &lt;= 100</code></li>
</ul>
', '[{"name": "Linked List", "slug": "linked-list", "translatedName": null}, {"name": "Recursion", "slug": "recursion", "translatedName": null}]'::jsonb, '[1,2,3,4]
[]
[1]
[1,2,3]', '[{"title": "Reverse Nodes in k-Group", "titleSlug": "reverse-nodes-in-k-group", "difficulty": "Hard", "translatedTitle": null}, {"title": "Swapping Nodes in a Linked List", "titleSlug": "swapping-nodes-in-a-linked-list", "difficulty": "Medium", "translatedTitle": null}]'::jsonb, '2025-03-05 18:12:57.681');
INSERT INTO problems
(id, frontend_id, title, title_slug, difficulty, is_paid_only, "content", topic_tags, example_testcases, similar_questions, created_at)
VALUES(25, 25, 'Reverse Nodes in k-Group', 'reverse-nodes-in-k-group', 'Hard', false, '<p>Given the <code>head</code> of a linked list, reverse the nodes of the list <code>k</code> at a time, and return <em>the modified list</em>.</p>

<p><code>k</code> is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of <code>k</code> then left-out nodes, in the end, should remain as it is.</p>

<p>You may not alter the values in the list&#39;s nodes, only nodes themselves may be changed.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>
<img alt="" src="https://assets.leetcode.com/uploads/2020/10/03/reverse_ex1.jpg" style="width: 542px; height: 222px;" />
<pre>
<strong>Input:</strong> head = [1,2,3,4,5], k = 2
<strong>Output:</strong> [2,1,4,3,5]
</pre>

<p><strong class="example">Example 2:</strong></p>
<img alt="" src="https://assets.leetcode.com/uploads/2020/10/03/reverse_ex2.jpg" style="width: 542px; height: 222px;" />
<pre>
<strong>Input:</strong> head = [1,2,3,4,5], k = 3
<strong>Output:</strong> [3,2,1,4,5]
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li>The number of nodes in the list is <code>n</code>.</li>
	<li><code>1 &lt;= k &lt;= n &lt;= 5000</code></li>
	<li><code>0 &lt;= Node.val &lt;= 1000</code></li>
</ul>

<p>&nbsp;</p>
<p><strong>Follow-up:</strong> Can you solve the problem in <code>O(1)</code> extra memory space?</p>
', '[{"name": "Linked List", "slug": "linked-list", "translatedName": null}, {"name": "Recursion", "slug": "recursion", "translatedName": null}]'::jsonb, '[1,2,3,4,5]
2
[1,2,3,4,5]
3', '[{"title": "Swap Nodes in Pairs", "titleSlug": "swap-nodes-in-pairs", "difficulty": "Medium", "translatedTitle": null}, {"title": "Swapping Nodes in a Linked List", "titleSlug": "swapping-nodes-in-a-linked-list", "difficulty": "Medium", "translatedTitle": null}, {"title": "Reverse Nodes in Even Length Groups", "titleSlug": "reverse-nodes-in-even-length-groups", "difficulty": "Medium", "translatedTitle": null}]'::jsonb, '2025-03-05 18:12:46.015');
INSERT INTO problems
(id, frontend_id, title, title_slug, difficulty, is_paid_only, "content", topic_tags, example_testcases, similar_questions, created_at)
VALUES(26, 26, 'Remove Duplicates from Sorted Array', 'remove-duplicates-from-sorted-array', 'Easy', false, '<p>Given an integer array <code>nums</code> sorted in <strong>non-decreasing order</strong>, remove the duplicates <a href="https://en.wikipedia.org/wiki/In-place_algorithm" target="_blank"><strong>in-place</strong></a> such that each unique element appears only <strong>once</strong>. The <strong>relative order</strong> of the elements should be kept the <strong>same</strong>. Then return <em>the number of unique elements in </em><code>nums</code>.</p>

<p>Consider the number of unique elements of <code>nums</code> to be <code>k</code>, to get accepted, you need to do the following things:</p>

<ul>
	<li>Change the array <code>nums</code> such that the first <code>k</code> elements of <code>nums</code> contain the unique elements in the order they were present in <code>nums</code> initially. The remaining elements of <code>nums</code> are not important as well as the size of <code>nums</code>.</li>
	<li>Return <code>k</code>.</li>
</ul>

<p><strong>Custom Judge:</strong></p>

<p>The judge will test your solution with the following code:</p>

<pre>
int[] nums = [...]; // Input array
int[] expectedNums = [...]; // The expected answer with correct length

int k = removeDuplicates(nums); // Calls your implementation

assert k == expectedNums.length;
for (int i = 0; i &lt; k; i++) {
    assert nums[i] == expectedNums[i];
}
</pre>

<p>If all assertions pass, then your solution will be <strong>accepted</strong>.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<pre>
<strong>Input:</strong> nums = [1,1,2]
<strong>Output:</strong> 2, nums = [1,2,_]
<strong>Explanation:</strong> Your function should return k = 2, with the first two elements of nums being 1 and 2 respectively.
It does not matter what you leave beyond the returned k (hence they are underscores).
</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input:</strong> nums = [0,0,1,1,1,2,2,3,3,4]
<strong>Output:</strong> 5, nums = [0,1,2,3,4,_,_,_,_,_]
<strong>Explanation:</strong> Your function should return k = 5, with the first five elements of nums being 0, 1, 2, 3, and 4 respectively.
It does not matter what you leave beyond the returned k (hence they are underscores).
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 3 * 10<sup>4</sup></code></li>
	<li><code>-100 &lt;= nums[i] &lt;= 100</code></li>
	<li><code>nums</code> is sorted in <strong>non-decreasing</strong> order.</li>
</ul>
', '[{"name": "Array", "slug": "array", "translatedName": null}, {"name": "Two Pointers", "slug": "two-pointers", "translatedName": null}]'::jsonb, '[1,1,2]
[0,0,1,1,1,2,2,3,3,4]', '[{"title": "Remove Element", "titleSlug": "remove-element", "difficulty": "Easy", "translatedTitle": null}, {"title": "Remove Duplicates from Sorted Array II", "titleSlug": "remove-duplicates-from-sorted-array-ii", "difficulty": "Medium", "translatedTitle": null}, {"title": "Apply Operations to an Array", "titleSlug": "apply-operations-to-an-array", "difficulty": "Easy", "translatedTitle": null}, {"title": "Sum of Distances", "titleSlug": "sum-of-distances", "difficulty": "Medium", "translatedTitle": null}]'::jsonb, '2025-03-05 18:12:35.254');
INSERT INTO problems
(id, frontend_id, title, title_slug, difficulty, is_paid_only, "content", topic_tags, example_testcases, similar_questions, created_at)
VALUES(27, 27, 'Remove Element', 'remove-element', 'Easy', false, '<p>Given an integer array <code>nums</code> and an integer <code>val</code>, remove all occurrences of <code>val</code> in <code>nums</code> <a href="https://en.wikipedia.org/wiki/In-place_algorithm" target="_blank"><strong>in-place</strong></a>. The order of the elements may be changed. Then return <em>the number of elements in </em><code>nums</code><em> which are not equal to </em><code>val</code>.</p>

<p>Consider the number of elements in <code>nums</code> which are not equal to <code>val</code> be <code>k</code>, to get accepted, you need to do the following things:</p>

<ul>
	<li>Change the array <code>nums</code> such that the first <code>k</code> elements of <code>nums</code> contain the elements which are not equal to <code>val</code>. The remaining elements of <code>nums</code> are not important as well as the size of <code>nums</code>.</li>
	<li>Return <code>k</code>.</li>
</ul>

<p><strong>Custom Judge:</strong></p>

<p>The judge will test your solution with the following code:</p>

<pre>
int[] nums = [...]; // Input array
int val = ...; // Value to remove
int[] expectedNums = [...]; // The expected answer with correct length.
                            // It is sorted with no values equaling val.

int k = removeElement(nums, val); // Calls your implementation

assert k == expectedNums.length;
sort(nums, 0, k); // Sort the first k elements of nums
for (int i = 0; i &lt; actualLength; i++) {
    assert nums[i] == expectedNums[i];
}
</pre>

<p>If all assertions pass, then your solution will be <strong>accepted</strong>.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<pre>
<strong>Input:</strong> nums = [3,2,2,3], val = 3
<strong>Output:</strong> 2, nums = [2,2,_,_]
<strong>Explanation:</strong> Your function should return k = 2, with the first two elements of nums being 2.
It does not matter what you leave beyond the returned k (hence they are underscores).
</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input:</strong> nums = [0,1,2,2,3,0,4,2], val = 2
<strong>Output:</strong> 5, nums = [0,1,4,0,3,_,_,_]
<strong>Explanation:</strong> Your function should return k = 5, with the first five elements of nums containing 0, 0, 1, 3, and 4.
Note that the five elements can be returned in any order.
It does not matter what you leave beyond the returned k (hence they are underscores).
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>0 &lt;= nums.length &lt;= 100</code></li>
	<li><code>0 &lt;= nums[i] &lt;= 50</code></li>
	<li><code>0 &lt;= val &lt;= 100</code></li>
</ul>
', '[{"name": "Array", "slug": "array", "translatedName": null}, {"name": "Two Pointers", "slug": "two-pointers", "translatedName": null}]'::jsonb, '[3,2,2,3]
3
[0,1,2,2,3,0,4,2]
2', '[{"title": "Remove Duplicates from Sorted Array", "titleSlug": "remove-duplicates-from-sorted-array", "difficulty": "Easy", "translatedTitle": null}, {"title": "Remove Linked List Elements", "titleSlug": "remove-linked-list-elements", "difficulty": "Easy", "translatedTitle": null}, {"title": "Move Zeroes", "titleSlug": "move-zeroes", "difficulty": "Easy", "translatedTitle": null}]'::jsonb, '2025-03-05 18:12:24.837');
INSERT INTO problems
(id, frontend_id, title, title_slug, difficulty, is_paid_only, "content", topic_tags, example_testcases, similar_questions, created_at)
VALUES(28, 28, 'Find the Index of the First Occurrence in a String', 'find-the-index-of-the-first-occurrence-in-a-string', 'Easy', false, '<p>Given two strings <code>needle</code> and <code>haystack</code>, return the index of the first occurrence of <code>needle</code> in <code>haystack</code>, or <code>-1</code> if <code>needle</code> is not part of <code>haystack</code>.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<pre>
<strong>Input:</strong> haystack = &quot;sadbutsad&quot;, needle = &quot;sad&quot;
<strong>Output:</strong> 0
<strong>Explanation:</strong> &quot;sad&quot; occurs at index 0 and 6.
The first occurrence is at index 0, so we return 0.
</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input:</strong> haystack = &quot;leetcode&quot;, needle = &quot;leeto&quot;
<strong>Output:</strong> -1
<strong>Explanation:</strong> &quot;leeto&quot; did not occur in &quot;leetcode&quot;, so we return -1.
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= haystack.length, needle.length &lt;= 10<sup>4</sup></code></li>
	<li><code>haystack</code> and <code>needle</code> consist of only lowercase English characters.</li>
</ul>
', '[{"name": "Two Pointers", "slug": "two-pointers", "translatedName": null}, {"name": "String", "slug": "string", "translatedName": null}, {"name": "String Matching", "slug": "string-matching", "translatedName": null}]'::jsonb, '"sadbutsad"
"sad"
"leetcode"
"leeto"', '[{"title": "Shortest Palindrome", "titleSlug": "shortest-palindrome", "difficulty": "Hard", "translatedTitle": null}, {"title": "Repeated Substring Pattern", "titleSlug": "repeated-substring-pattern", "difficulty": "Easy", "translatedTitle": null}]'::jsonb, '2025-03-05 18:12:14.302');
INSERT INTO problems
(id, frontend_id, title, title_slug, difficulty, is_paid_only, "content", topic_tags, example_testcases, similar_questions, created_at)
VALUES(29, 29, 'Divide Two Integers', 'divide-two-integers', 'Medium', false, '<p>Given two integers <code>dividend</code> and <code>divisor</code>, divide two integers <strong>without</strong> using multiplication, division, and mod operator.</p>

<p>The integer division should truncate toward zero, which means losing its fractional part. For example, <code>8.345</code> would be truncated to <code>8</code>, and <code>-2.7335</code> would be truncated to <code>-2</code>.</p>

<p>Return <em>the <strong>quotient</strong> after dividing </em><code>dividend</code><em> by </em><code>divisor</code>.</p>

<p><strong>Note: </strong>Assume we are dealing with an environment that could only store integers within the <strong>32-bit</strong> signed integer range: <code>[&minus;2<sup>31</sup>, 2<sup>31</sup> &minus; 1]</code>. For this problem, if the quotient is <strong>strictly greater than</strong> <code>2<sup>31</sup> - 1</code>, then return <code>2<sup>31</sup> - 1</code>, and if the quotient is <strong>strictly less than</strong> <code>-2<sup>31</sup></code>, then return <code>-2<sup>31</sup></code>.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<pre>
<strong>Input:</strong> dividend = 10, divisor = 3
<strong>Output:</strong> 3
<strong>Explanation:</strong> 10/3 = 3.33333.. which is truncated to 3.
</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input:</strong> dividend = 7, divisor = -3
<strong>Output:</strong> -2
<strong>Explanation:</strong> 7/-3 = -2.33333.. which is truncated to -2.
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>-2<sup>31</sup> &lt;= dividend, divisor &lt;= 2<sup>31</sup> - 1</code></li>
	<li><code>divisor != 0</code></li>
</ul>
', '[{"name": "Math", "slug": "math", "translatedName": null}, {"name": "Bit Manipulation", "slug": "bit-manipulation", "translatedName": null}]'::jsonb, '10
3
7
-3', '[]'::jsonb, '2025-03-05 18:12:03.818');
INSERT INTO problems
(id, frontend_id, title, title_slug, difficulty, is_paid_only, "content", topic_tags, example_testcases, similar_questions, created_at)
VALUES(30, 30, 'Substring with Concatenation of All Words', 'substring-with-concatenation-of-all-words', 'Hard', false, '<p>You are given a string <code>s</code> and an array of strings <code>words</code>. All the strings of <code>words</code> are of <strong>the same length</strong>.</p>

<p>A <strong>concatenated string</strong> is a string that exactly contains all the strings of any permutation of <code>words</code> concatenated.</p>

<ul>
	<li>For example, if <code>words = [&quot;ab&quot;,&quot;cd&quot;,&quot;ef&quot;]</code>, then <code>&quot;abcdef&quot;</code>, <code>&quot;abefcd&quot;</code>, <code>&quot;cdabef&quot;</code>, <code>&quot;cdefab&quot;</code>, <code>&quot;efabcd&quot;</code>, and <code>&quot;efcdab&quot;</code> are all concatenated strings. <code>&quot;acdbef&quot;</code> is not a concatenated string because it is not the concatenation of any permutation of <code>words</code>.</li>
</ul>

<p>Return an array of <em>the starting indices</em> of all the concatenated substrings in <code>s</code>. You can return the answer in <strong>any order</strong>.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<div class="example-block">
<p><strong>Input:</strong> <span class="example-io">s = &quot;barfoothefoobarman&quot;, words = [&quot;foo&quot;,&quot;bar&quot;]</span></p>

<p><strong>Output:</strong> <span class="example-io">[0,9]</span></p>

<p><strong>Explanation:</strong></p>

<p>The substring starting at 0 is <code>&quot;barfoo&quot;</code>. It is the concatenation of <code>[&quot;bar&quot;,&quot;foo&quot;]</code> which is a permutation of <code>words</code>.<br />
The substring starting at 9 is <code>&quot;foobar&quot;</code>. It is the concatenation of <code>[&quot;foo&quot;,&quot;bar&quot;]</code> which is a permutation of <code>words</code>.</p>
</div>

<p><strong class="example">Example 2:</strong></p>

<div class="example-block">
<p><strong>Input:</strong> <span class="example-io">s = &quot;wordgoodgoodgoodbestword&quot;, words = [&quot;word&quot;,&quot;good&quot;,&quot;best&quot;,&quot;word&quot;]</span></p>

<p><strong>Output:</strong> <span class="example-io">[]</span></p>

<p><strong>Explanation:</strong></p>

<p>There is no concatenated substring.</p>
</div>

<p><strong class="example">Example 3:</strong></p>

<div class="example-block">
<p><strong>Input:</strong> <span class="example-io">s = &quot;barfoofoobarthefoobarman&quot;, words = [&quot;bar&quot;,&quot;foo&quot;,&quot;the&quot;]</span></p>

<p><strong>Output:</strong> <span class="example-io">[6,9,12]</span></p>

<p><strong>Explanation:</strong></p>

<p>The substring starting at 6 is <code>&quot;foobarthe&quot;</code>. It is the concatenation of <code>[&quot;foo&quot;,&quot;bar&quot;,&quot;the&quot;]</code>.<br />
The substring starting at 9 is <code>&quot;barthefoo&quot;</code>. It is the concatenation of <code>[&quot;bar&quot;,&quot;the&quot;,&quot;foo&quot;]</code>.<br />
The substring starting at 12 is <code>&quot;thefoobar&quot;</code>. It is the concatenation of <code>[&quot;the&quot;,&quot;foo&quot;,&quot;bar&quot;]</code>.</p>
</div>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= s.length &lt;= 10<sup>4</sup></code></li>
	<li><code>1 &lt;= words.length &lt;= 5000</code></li>
	<li><code>1 &lt;= words[i].length &lt;= 30</code></li>
	<li><code>s</code> and <code>words[i]</code> consist of lowercase English letters.</li>
</ul>
', '[{"name": "Hash Table", "slug": "hash-table", "translatedName": null}, {"name": "String", "slug": "string", "translatedName": null}, {"name": "Sliding Window", "slug": "sliding-window", "translatedName": null}]'::jsonb, '"barfoothefoobarman"
["foo","bar"]
"wordgoodgoodgoodbestword"
["word","good","best","word"]
"barfoofoobarthefoobarman"
["bar","foo","the"]', '[{"title": "Minimum Window Substring", "titleSlug": "minimum-window-substring", "difficulty": "Hard", "translatedTitle": null}]'::jsonb, '2025-03-05 18:11:53.427');