INSERT INTO problems
(id, frontend_id, title, title_slug, difficulty, is_paid_only, "content", topic_tags, example_testcases, similar_questions, created_at)
VALUES(2111, 1966, 'Binary Searchable Numbers in an Unsorted Array', 'binary-searchable-numbers-in-an-unsorted-array', 'Medium', true, '', '[{"name": "Array", "slug": "array", "translatedName": null}, {"name": "Binary Search", "slug": "binary-search", "translatedName": null}]'::jsonb, '[7]
[-1,5,2]', '[]'::jsonb, '2025-03-05 12:47:56.037');
INSERT INTO problems
(id, frontend_id, title, title_slug, difficulty, is_paid_only, "content", topic_tags, example_testcases, similar_questions, created_at)
VALUES(3529, 3221, 'Maximum Array Hopping Score II', 'maximum-array-hopping-score-ii', 'Medium', true, '', '[{"name": "Array", "slug": "array", "translatedName": null}, {"name": "Stack", "slug": "stack", "translatedName": null}, {"name": "Greedy", "slug": "greedy", "translatedName": null}, {"name": "Monotonic Stack", "slug": "monotonic-stack", "translatedName": null}]'::jsonb, '[1,5,8]
[4,5,2,8,9,1,3]', '[]'::jsonb, '2025-03-05 08:09:05.241');
INSERT INTO problems
(id, frontend_id, title, title_slug, difficulty, is_paid_only, "content", topic_tags, example_testcases, similar_questions, created_at)
VALUES(1810, 1666, 'Change the Root of a Binary Tree', 'change-the-root-of-a-binary-tree', 'Medium', true, '', '[{"name": "Tree", "slug": "tree", "translatedName": null}, {"name": "Depth-First Search", "slug": "depth-first-search", "translatedName": null}, {"name": "Binary Tree", "slug": "binary-tree", "translatedName": null}]'::jsonb, '[3,5,1,6,2,0,8,null,null,7,4]
7
[3,5,1,6,2,0,8,null,null,7,4]
0', '[]'::jsonb, '2025-03-05 13:33:32.601');
INSERT INTO problems
(id, frontend_id, title, title_slug, difficulty, is_paid_only, "content", topic_tags, example_testcases, similar_questions, created_at)
VALUES(3623, 3299, 'Sum of Consecutive Subsequences', 'sum-of-consecutive-subsequences', 'Hard', true, '', '[{"name": "Array", "slug": "array", "translatedName": null}, {"name": "Hash Table", "slug": "hash-table", "translatedName": null}, {"name": "Dynamic Programming", "slug": "dynamic-programming", "translatedName": null}]'::jsonb, '[1,2]
[1,4,2,3]', '[]'::jsonb, '2025-03-05 07:54:12.760');
INSERT INTO problems
(id, frontend_id, title, title_slug, difficulty, is_paid_only, "content", topic_tags, example_testcases, similar_questions, created_at)
VALUES(3615, 3294, 'Convert Doubly Linked List to Array II', 'convert-doubly-linked-list-to-array-ii', 'Medium', true, '', '[{"name": "Array", "slug": "array", "translatedName": null}, {"name": "Linked List", "slug": "linked-list", "translatedName": null}, {"name": "Doubly-Linked List", "slug": "doubly-linked-list", "translatedName": null}]'::jsonb, '[1,2,3,4,5]
5
[4,5,6,7,8]
8', '[{"title": "Remove Linked List Elements", "titleSlug": "remove-linked-list-elements", "difficulty": "Easy", "translatedTitle": null}]'::jsonb, '2025-03-05 07:55:43.714');
INSERT INTO problems
(id, frontend_id, title, title_slug, difficulty, is_paid_only, "content", topic_tags, example_testcases, similar_questions, created_at)
VALUES(2535, 2393, 'Count Strictly Increasing Subarrays', 'count-strictly-increasing-subarrays', 'Medium', true, '', '[{"name": "Array", "slug": "array", "translatedName": null}, {"name": "Math", "slug": "math", "translatedName": null}, {"name": "Dynamic Programming", "slug": "dynamic-programming", "translatedName": null}]'::jsonb, '[1,3,5,4,4,6]
[1,2,3,4,5]', '[{"title": "Maximum Ascending Subarray Sum", "titleSlug": "maximum-ascending-subarray-sum", "difficulty": "Easy", "translatedTitle": null}]'::jsonb, '2025-03-05 10:11:31.346');
INSERT INTO problems
(id, frontend_id, title, title_slug, difficulty, is_paid_only, "content", topic_tags, example_testcases, similar_questions, created_at)
VALUES(3545, 3231, 'Minimum Number of Increasing Subsequence to Be Removed', 'minimum-number-of-increasing-subsequence-to-be-removed', 'Hard', true, '', '[{"name": "Array", "slug": "array", "translatedName": null}, {"name": "Binary Search", "slug": "binary-search", "translatedName": null}]'::jsonb, '[5,3,1,4,2]
[1,2,3,4,5]
[5,4,3,2,1]', '[]'::jsonb, '2025-03-05 08:06:28.610');
INSERT INTO problems
(id, frontend_id, title, title_slug, difficulty, is_paid_only, "content", topic_tags, example_testcases, similar_questions, created_at)
VALUES(3533, 3248, 'Snake in Matrix', 'snake-in-matrix', 'Easy', false, '<p>There is a snake in an <code>n x n</code> matrix <code>grid</code> and can move in <strong>four possible directions</strong>. Each cell in the <code>grid</code> is identified by the position: <code>grid[i][j] = (i * n) + j</code>.</p>

<p>The snake starts at cell 0 and follows a sequence of commands.</p>

<p>You are given an integer <code>n</code> representing the size of the <code>grid</code> and an array of strings <code>commands</code> where each <code>command[i]</code> is either <code>&quot;UP&quot;</code>, <code>&quot;RIGHT&quot;</code>, <code>&quot;DOWN&quot;</code>, and <code>&quot;LEFT&quot;</code>. It&#39;s guaranteed that the snake will remain within the <code>grid</code> boundaries throughout its movement.</p>

<p>Return the position of the final cell where the snake ends up after executing <code>commands</code>.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<div class="example-block">
<p><strong>Input:</strong> <span class="example-io">n = 2, commands = [&quot;RIGHT&quot;,&quot;DOWN&quot;]</span></p>

<p><strong>Output:</strong> <span class="example-io">3</span></p>

<p><strong>Explanation:</strong></p>

<div style="display:flex; gap: 12px;">
<table border="1" cellspacing="3" style="border-collapse: separate; text-align: center;">
	<tbody>
		<tr>
			<td data-darkreader-inline-border-bottom="" data-darkreader-inline-border-left="" data-darkreader-inline-border-right="" data-darkreader-inline-border-top="" style="padding: 5px 10px; border: 1px solid red; --darkreader-inline-border-top: #8c8273; --darkreader-inline-border-right: #8c8273; --darkreader-inline-border-bottom: #8c8273; --darkreader-inline-border-left: #8c8273;">0</td>
			<td data-darkreader-inline-border-bottom="" data-darkreader-inline-border-left="" data-darkreader-inline-border-right="" data-darkreader-inline-border-top="" style="padding: 5px 10px; border: 1px solid black; --darkreader-inline-border-top: #b30000; --darkreader-inline-border-right: #b30000; --darkreader-inline-border-bottom: #b30000; --darkreader-inline-border-left: #b30000;">1</td>
		</tr>
		<tr>
			<td data-darkreader-inline-border-bottom="" data-darkreader-inline-border-left="" data-darkreader-inline-border-right="" data-darkreader-inline-border-top="" style="padding: 5px 10px; border: 1px solid black; --darkreader-inline-border-top: #8c8273; --darkreader-inline-border-right: #8c8273; --darkreader-inline-border-bottom: #8c8273; --darkreader-inline-border-left: #8c8273;">2</td>
			<td data-darkreader-inline-border-bottom="" data-darkreader-inline-border-left="" data-darkreader-inline-border-right="" data-darkreader-inline-border-top="" style="padding: 5px 10px; border: 1px solid black; --darkreader-inline-border-top: #b30000; --darkreader-inline-border-right: #b30000; --darkreader-inline-border-bottom: #b30000; --darkreader-inline-border-left: #b30000;">3</td>
		</tr>
	</tbody>
</table>

<table border="1" cellspacing="3" style="border-collapse: separate; text-align: center;">
	<tbody>
		<tr>
			<td data-darkreader-inline-border-bottom="" data-darkreader-inline-border-left="" data-darkreader-inline-border-right="" data-darkreader-inline-border-top="" style="padding: 5px 10px; border: 1px solid black; --darkreader-inline-border-top: #8c8273; --darkreader-inline-border-right: #8c8273; --darkreader-inline-border-bottom: #8c8273; --darkreader-inline-border-left: #8c8273;">0</td>
			<td data-darkreader-inline-border-bottom="" data-darkreader-inline-border-left="" data-darkreader-inline-border-right="" data-darkreader-inline-border-top="" style="padding: 5px 10px; border: 1px solid red; --darkreader-inline-border-top: #b30000; --darkreader-inline-border-right: #b30000; --darkreader-inline-border-bottom: #b30000; --darkreader-inline-border-left: #b30000;">1</td>
		</tr>
		<tr>
			<td data-darkreader-inline-border-bottom="" data-darkreader-inline-border-left="" data-darkreader-inline-border-right="" data-darkreader-inline-border-top="" style="padding: 5px 10px; border: 1px solid black; --darkreader-inline-border-top: #8c8273; --darkreader-inline-border-right: #8c8273; --darkreader-inline-border-bottom: #8c8273; --darkreader-inline-border-left: #8c8273;">2</td>
			<td data-darkreader-inline-border-bottom="" data-darkreader-inline-border-left="" data-darkreader-inline-border-right="" data-darkreader-inline-border-top="" style="padding: 5px 10px; border: 1px solid black; --darkreader-inline-border-top: #b30000; --darkreader-inline-border-right: #b30000; --darkreader-inline-border-bottom: #b30000; --darkreader-inline-border-left: #b30000;">3</td>
		</tr>
	</tbody>
</table>

<table border="1" cellspacing="3" style="border-collapse: separate; text-align: center;">
	<tbody>
		<tr>
			<td data-darkreader-inline-border-bottom="" data-darkreader-inline-border-left="" data-darkreader-inline-border-right="" data-darkreader-inline-border-top="" style="padding: 5px 10px; border: 1px solid black; --darkreader-inline-border-top: #8c8273; --darkreader-inline-border-right: #8c8273; --darkreader-inline-border-bottom: #8c8273; --darkreader-inline-border-left: #8c8273;">0</td>
			<td data-darkreader-inline-border-bottom="" data-darkreader-inline-border-left="" data-darkreader-inline-border-right="" data-darkreader-inline-border-top="" style="padding: 5px 10px; border: 1px solid black; --darkreader-inline-border-top: #b30000; --darkreader-inline-border-right: #b30000; --darkreader-inline-border-bottom: #b30000; --darkreader-inline-border-left: #b30000;">1</td>
		</tr>
		<tr>
			<td data-darkreader-inline-border-bottom="" data-darkreader-inline-border-left="" data-darkreader-inline-border-right="" data-darkreader-inline-border-top="" style="padding: 5px 10px; border: 1px solid black; --darkreader-inline-border-top: #8c8273; --darkreader-inline-border-right: #8c8273; --darkreader-inline-border-bottom: #8c8273; --darkreader-inline-border-left: #8c8273;">2</td>
			<td data-darkreader-inline-border-bottom="" data-darkreader-inline-border-left="" data-darkreader-inline-border-right="" data-darkreader-inline-border-top="" style="padding: 5px 10px; border: 1px solid red; --darkreader-inline-border-top: #b30000; --darkreader-inline-border-right: #b30000; --darkreader-inline-border-bottom: #b30000; --darkreader-inline-border-left: #b30000;">3</td>
		</tr>
	</tbody>
</table>
</div>
</div>

<p><strong class="example">Example 2:</strong></p>

<div class="example-block">
<p><strong>Input:</strong> <span class="example-io">n = 3, commands = [&quot;DOWN&quot;,&quot;RIGHT&quot;,&quot;UP&quot;]</span></p>

<p><strong>Output:</strong> <span class="example-io">1</span></p>

<p><strong>Explanation:</strong></p>

<div style="display:flex; gap: 12px;">
<table border="1" cellspacing="3" style="border-collapse: separate; text-align: center;">
	<tbody>
		<tr>
			<td data-darkreader-inline-border-bottom="" data-darkreader-inline-border-left="" data-darkreader-inline-border-right="" data-darkreader-inline-border-top="" style="padding: 5px 10px; border: 1px solid red; --darkreader-inline-border-top: #8c8273; --darkreader-inline-border-right: #8c8273; --darkreader-inline-border-bottom: #8c8273; --darkreader-inline-border-left: #8c8273;">0</td>
			<td data-darkreader-inline-border-bottom="" data-darkreader-inline-border-left="" data-darkreader-inline-border-right="" data-darkreader-inline-border-top="" style="padding: 5px 10px; border: 1px solid black; --darkreader-inline-border-top: #b30000; --darkreader-inline-border-right: #b30000; --darkreader-inline-border-bottom: #b30000; --darkreader-inline-border-left: #b30000;">1</td>
			<td data-darkreader-inline-border-bottom="" data-darkreader-inline-border-left="" data-darkreader-inline-border-right="" data-darkreader-inline-border-top="" style="padding: 5px 10px; border: 1px solid black; --darkreader-inline-border-top: #8c8273; --darkreader-inline-border-right: #8c8273; --darkreader-inline-border-bottom: #8c8273; --darkreader-inline-border-left: #8c8273;">2</td>
		</tr>
		<tr>
			<td data-darkreader-inline-border-bottom="" data-darkreader-inline-border-left="" data-darkreader-inline-border-right="" data-darkreader-inline-border-top="" style="padding: 5px 10px; border: 1px solid black; --darkreader-inline-border-top: #8c8273; --darkreader-inline-border-right: #8c8273; --darkreader-inline-border-bottom: #8c8273; --darkreader-inline-border-left: #8c8273;">3</td>
			<td data-darkreader-inline-border-bottom="" data-darkreader-inline-border-left="" data-darkreader-inline-border-right="" data-darkreader-inline-border-top="" style="padding: 5px 10px; border: 1px solid black; --darkreader-inline-border-top: #b30000; --darkreader-inline-border-right: #b30000; --darkreader-inline-border-bottom: #b30000; --darkreader-inline-border-left: #b30000;">4</td>
			<td data-darkreader-inline-border-bottom="" data-darkreader-inline-border-left="" data-darkreader-inline-border-right="" data-darkreader-inline-border-top="" style="padding: 5px 10px; border: 1px solid black; --darkreader-inline-border-top: #b30000; --darkreader-inline-border-right: #b30000; --darkreader-inline-border-bottom: #b30000; --darkreader-inline-border-left: #b30000;">5</td>
		</tr>
		<tr>
			<td data-darkreader-inline-border-bottom="" data-darkreader-inline-border-left="" data-darkreader-inline-border-right="" data-darkreader-inline-border-top="" style="padding: 5px 10px; border: 1px solid black; --darkreader-inline-border-top: #8c8273; --darkreader-inline-border-right: #8c8273; --darkreader-inline-border-bottom: #8c8273; --darkreader-inline-border-left: #8c8273;">6</td>
			<td data-darkreader-inline-border-bottom="" data-darkreader-inline-border-left="" data-darkreader-inline-border-right="" data-darkreader-inline-border-top="" style="padding: 5px 10px; border: 1px solid black; --darkreader-inline-border-top: #8c8273; --darkreader-inline-border-right: #8c8273; --darkreader-inline-border-bottom: #8c8273; --darkreader-inline-border-left: #8c8273;">7</td>
			<td data-darkreader-inline-border-bottom="" data-darkreader-inline-border-left="" data-darkreader-inline-border-right="" data-darkreader-inline-border-top="" style="padding: 5px 10px; border: 1px solid black; --darkreader-inline-border-top: #8c8273; --darkreader-inline-border-right: #8c8273; --darkreader-inline-border-bottom: #8c8273; --darkreader-inline-border-left: #8c8273;">8</td>
		</tr>
	</tbody>
</table>

<table border="1" cellspacing="3" style="border-collapse: separate; text-align: center;">
	<tbody>
		<tr>
			<td data-darkreader-inline-border-bottom="" data-darkreader-inline-border-left="" data-darkreader-inline-border-right="" data-darkreader-inline-border-top="" style="padding: 5px 10px; border: 1px solid black; --darkreader-inline-border-top: #8c8273; --darkreader-inline-border-right: #8c8273; --darkreader-inline-border-bottom: #8c8273; --darkreader-inline-border-left: #8c8273;">0</td>
			<td data-darkreader-inline-border-bottom="" data-darkreader-inline-border-left="" data-darkreader-inline-border-right="" data-darkreader-inline-border-top="" style="padding: 5px 10px; border: 1px solid black; --darkreader-inline-border-top: #b30000; --darkreader-inline-border-right: #b30000; --darkreader-inline-border-bottom: #b30000; --darkreader-inline-border-left: #b30000;">1</td>
			<td data-darkreader-inline-border-bottom="" data-darkreader-inline-border-left="" data-darkreader-inline-border-right="" data-darkreader-inline-border-top="" style="padding: 5px 10px; border: 1px solid black; --darkreader-inline-border-top: #8c8273; --darkreader-inline-border-right: #8c8273; --darkreader-inline-border-bottom: #8c8273; --darkreader-inline-border-left: #8c8273;">2</td>
		</tr>
		<tr>
			<td data-darkreader-inline-border-bottom="" data-darkreader-inline-border-left="" data-darkreader-inline-border-right="" data-darkreader-inline-border-top="" style="padding: 5px 10px; border: 1px solid red; --darkreader-inline-border-top: #8c8273; --darkreader-inline-border-right: #8c8273; --darkreader-inline-border-bottom: #8c8273; --darkreader-inline-border-left: #8c8273;">3</td>
			<td data-darkreader-inline-border-bottom="" data-darkreader-inline-border-left="" data-darkreader-inline-border-right="" data-darkreader-inline-border-top="" style="padding: 5px 10px; border: 1px solid black; --darkreader-inline-border-top: #b30000; --darkreader-inline-border-right: #b30000; --darkreader-inline-border-bottom: #b30000; --darkreader-inline-border-left: #b30000;">4</td>
			<td data-darkreader-inline-border-bottom="" data-darkreader-inline-border-left="" data-darkreader-inline-border-right="" data-darkreader-inline-border-top="" style="padding: 5px 10px; border: 1px solid black; --darkreader-inline-border-top: #b30000; --darkreader-inline-border-right: #b30000; --darkreader-inline-border-bottom: #b30000; --darkreader-inline-border-left: #b30000;">5</td>
		</tr>
		<tr>
			<td data-darkreader-inline-border-bottom="" data-darkreader-inline-border-left="" data-darkreader-inline-border-right="" data-darkreader-inline-border-top="" style="padding: 5px 10px; border: 1px solid black; --darkreader-inline-border-top: #8c8273; --darkreader-inline-border-right: #8c8273; --darkreader-inline-border-bottom: #8c8273; --darkreader-inline-border-left: #8c8273;">6</td>
			<td data-darkreader-inline-border-bottom="" data-darkreader-inline-border-left="" data-darkreader-inline-border-right="" data-darkreader-inline-border-top="" style="padding: 5px 10px; border: 1px solid black; --darkreader-inline-border-top: #8c8273; --darkreader-inline-border-right: #8c8273; --darkreader-inline-border-bottom: #8c8273; --darkreader-inline-border-left: #8c8273;">7</td>
			<td data-darkreader-inline-border-bottom="" data-darkreader-inline-border-left="" data-darkreader-inline-border-right="" data-darkreader-inline-border-top="" style="padding: 5px 10px; border: 1px solid black; --darkreader-inline-border-top: #8c8273; --darkreader-inline-border-right: #8c8273; --darkreader-inline-border-bottom: #8c8273; --darkreader-inline-border-left: #8c8273;">8</td>
		</tr>
	</tbody>
</table>

<table border="1" cellspacing="3" style="border-collapse: separate; text-align: center;">
	<tbody>
		<tr>
			<td data-darkreader-inline-border-bottom="" data-darkreader-inline-border-left="" data-darkreader-inline-border-right="" data-darkreader-inline-border-top="" style="padding: 5px 10px; border: 1px solid black; --darkreader-inline-border-top: #8c8273; --darkreader-inline-border-right: #8c8273; --darkreader-inline-border-bottom: #8c8273; --darkreader-inline-border-left: #8c8273;">0</td>
			<td data-darkreader-inline-border-bottom="" data-darkreader-inline-border-left="" data-darkreader-inline-border-right="" data-darkreader-inline-border-top="" style="padding: 5px 10px; border: 1px solid black; --darkreader-inline-border-top: #8c8273; --darkreader-inline-border-right: #8c8273; --darkreader-inline-border-bottom: #8c8273; --darkreader-inline-border-left: #8c8273;">1</td>
			<td data-darkreader-inline-border-bottom="" data-darkreader-inline-border-left="" data-darkreader-inline-border-right="" data-darkreader-inline-border-top="" style="padding: 5px 10px; border: 1px solid black; --darkreader-inline-border-top: #8c8273; --darkreader-inline-border-right: #8c8273; --darkreader-inline-border-bottom: #8c8273; --darkreader-inline-border-left: #8c8273;">2</td>
		</tr>
		<tr>
			<td data-darkreader-inline-border-bottom="" data-darkreader-inline-border-left="" data-darkreader-inline-border-right="" data-darkreader-inline-border-top="" style="padding: 5px 10px; border: 1px solid black; --darkreader-inline-border-top: #8c8273; --darkreader-inline-border-right: #8c8273; --darkreader-inline-border-bottom: #8c8273; --darkreader-inline-border-left: #8c8273;">3</td>
			<td data-darkreader-inline-border-bottom="" data-darkreader-inline-border-left="" data-darkreader-inline-border-right="" data-darkreader-inline-border-top="" style="padding: 5px 10px; border: 1px solid red; --darkreader-inline-border-top: #b30000; --darkreader-inline-border-right: #b30000; --darkreader-inline-border-bottom: #b30000; --darkreader-inline-border-left: #b30000;">4</td>
			<td data-darkreader-inline-border-bottom="" data-darkreader-inline-border-left="" data-darkreader-inline-border-right="" data-darkreader-inline-border-top="" style="padding: 5px 10px; border: 1px solid black; --darkreader-inline-border-top: #b30000; --darkreader-inline-border-right: #b30000; --darkreader-inline-border-bottom: #b30000; --darkreader-inline-border-left: #b30000;">5</td>
		</tr>
		<tr>
			<td data-darkreader-inline-border-bottom="" data-darkreader-inline-border-left="" data-darkreader-inline-border-right="" data-darkreader-inline-border-top="" style="padding: 5px 10px; border: 1px solid black; --darkreader-inline-border-top: #8c8273; --darkreader-inline-border-right: #8c8273; --darkreader-inline-border-bottom: #8c8273; --darkreader-inline-border-left: #8c8273;">6</td>
			<td data-darkreader-inline-border-bottom="" data-darkreader-inline-border-left="" data-darkreader-inline-border-right="" data-darkreader-inline-border-top="" style="padding: 5px 10px; border: 1px solid black; --darkreader-inline-border-top: #b30000; --darkreader-inline-border-right: #b30000; --darkreader-inline-border-bottom: #b30000; --darkreader-inline-border-left: #b30000;">7</td>
			<td data-darkreader-inline-border-bottom="" data-darkreader-inline-border-left="" data-darkreader-inline-border-right="" data-darkreader-inline-border-top="" style="padding: 5px 10px; border: 1px solid black; --darkreader-inline-border-top: #8c8273; --darkreader-inline-border-right: #8c8273; --darkreader-inline-border-bottom: #8c8273; --darkreader-inline-border-left: #8c8273;">8</td>
		</tr>
	</tbody>
</table>

<table border="1" cellspacing="3" style="border-collapse: separate; text-align: center;">
	<tbody>
		<tr>
			<td data-darkreader-inline-border-bottom="" data-darkreader-inline-border-left="" data-darkreader-inline-border-right="" data-darkreader-inline-border-top="" style="padding: 5px 10px; border: 1px solid black; --darkreader-inline-border-top: #8c8273; --darkreader-inline-border-right: #8c8273; --darkreader-inline-border-bottom: #8c8273; --darkreader-inline-border-left: #8c8273;">0</td>
			<td data-darkreader-inline-border-bottom="" data-darkreader-inline-border-left="" data-darkreader-inline-border-right="" data-darkreader-inline-border-top="" style="padding: 5px 10px; border: 1px solid red; --darkreader-inline-border-top: #b30000; --darkreader-inline-border-right: #b30000; --darkreader-inline-border-bottom: #b30000; --darkreader-inline-border-left: #b30000;">1</td>
			<td data-darkreader-inline-border-bottom="" data-darkreader-inline-border-left="" data-darkreader-inline-border-right="" data-darkreader-inline-border-top="" style="padding: 5px 10px; border: 1px solid black; --darkreader-inline-border-top: #8c8273; --darkreader-inline-border-right: #8c8273; --darkreader-inline-border-bottom: #8c8273; --darkreader-inline-border-left: #8c8273;">2</td>
		</tr>
		<tr>
			<td data-darkreader-inline-border-bottom="" data-darkreader-inline-border-left="" data-darkreader-inline-border-right="" data-darkreader-inline-border-top="" style="padding: 5px 10px; border: 1px solid black; --darkreader-inline-border-top: #8c8273; --darkreader-inline-border-right: #8c8273; --darkreader-inline-border-bottom: #8c8273; --darkreader-inline-border-left: #8c8273;">3</td>
			<td data-darkreader-inline-border-bottom="" data-darkreader-inline-border-left="" data-darkreader-inline-border-right="" data-darkreader-inline-border-top="" style="padding: 5px 10px; border: 1px solid black; --darkreader-inline-border-top: #b30000; --darkreader-inline-border-right: #b30000; --darkreader-inline-border-bottom: #b30000; --darkreader-inline-border-left: #b30000;">4</td>
			<td data-darkreader-inline-border-bottom="" data-darkreader-inline-border-left="" data-darkreader-inline-border-right="" data-darkreader-inline-border-top="" style="padding: 5px 10px; border: 1px solid black; --darkreader-inline-border-top: #b30000; --darkreader-inline-border-right: #b30000; --darkreader-inline-border-bottom: #b30000; --darkreader-inline-border-left: #b30000;">5</td>
		</tr>
		<tr>
			<td data-darkreader-inline-border-bottom="" data-darkreader-inline-border-left="" data-darkreader-inline-border-right="" data-darkreader-inline-border-top="" style="padding: 5px 10px; border: 1px solid black; --darkreader-inline-border-top: #8c8273; --darkreader-inline-border-right: #8c8273; --darkreader-inline-border-bottom: #8c8273; --darkreader-inline-border-left: #8c8273;">6</td>
			<td data-darkreader-inline-border-bottom="" data-darkreader-inline-border-left="" data-darkreader-inline-border-right="" data-darkreader-inline-border-top="" style="padding: 5px 10px; border: 1px solid black; --darkreader-inline-border-top: #8c8273; --darkreader-inline-border-right: #8c8273; --darkreader-inline-border-bottom: #8c8273; --darkreader-inline-border-left: #8c8273;">7</td>
			<td data-darkreader-inline-border-bottom="" data-darkreader-inline-border-left="" data-darkreader-inline-border-right="" data-darkreader-inline-border-top="" style="padding: 5px 10px; border: 1px solid black; --darkreader-inline-border-top: #8c8273; --darkreader-inline-border-right: #8c8273; --darkreader-inline-border-bottom: #8c8273; --darkreader-inline-border-left: #8c8273;">8</td>
		</tr>
	</tbody>
</table>
</div>
</div>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>2 &lt;= n &lt;= 10</code></li>
	<li><code>1 &lt;= commands.length &lt;= 100</code></li>
	<li><code>commands</code> consists only of <code>&quot;UP&quot;</code>, <code>&quot;RIGHT&quot;</code>, <code>&quot;DOWN&quot;</code>, and <code>&quot;LEFT&quot;</code>.</li>
	<li>The input is generated such the snake will not move outside of the boundaries.</li>
</ul>
', '[{"name": "Array", "slug": "array", "translatedName": null}, {"name": "String", "slug": "string", "translatedName": null}, {"name": "Simulation", "slug": "simulation", "translatedName": null}]'::jsonb, '2
["RIGHT","DOWN"]
3
["DOWN","RIGHT","UP"]', '[]'::jsonb, '2025-03-05 08:08:30.538');
INSERT INTO problems
(id, frontend_id, title, title_slug, difficulty, is_paid_only, "content", topic_tags, example_testcases, similar_questions, created_at)
VALUES(3499, 3329, 'Count Substrings With K-Frequency Characters II', 'count-substrings-with-k-frequency-characters-ii', 'Hard', true, '', '[{"name": "Hash Table", "slug": "hash-table", "translatedName": null}, {"name": "String", "slug": "string", "translatedName": null}, {"name": "Sliding Window", "slug": "sliding-window", "translatedName": null}]'::jsonb, '"abacb"
2
"abcde"
1', '[]'::jsonb, '2025-03-05 08:13:53.759');
INSERT INTO problems
(id, frontend_id, title, title_slug, difficulty, is_paid_only, "content", topic_tags, example_testcases, similar_questions, created_at)
VALUES(1384, 1618, 'Maximum Font to Fit a Sentence in a Screen', 'maximum-font-to-fit-a-sentence-in-a-screen', 'Medium', true, '', '[{"name": "Array", "slug": "array", "translatedName": null}, {"name": "String", "slug": "string", "translatedName": null}, {"name": "Binary Search", "slug": "binary-search", "translatedName": null}, {"name": "Interactive", "slug": "interactive", "translatedName": null}]'::jsonb, '"helloworld"
80
20
[6,8,10,12,14,16,18,24,36]
"leetcode"
1000
50
[1,2,4]
"easyquestion"
100
100
[10,15,20,25]', '[]'::jsonb, '2025-03-05 14:35:36.882');
INSERT INTO problems
(id, frontend_id, title, title_slug, difficulty, is_paid_only, "content", topic_tags, example_testcases, similar_questions, created_at)
VALUES(3474, 3167, 'Better Compression of String', 'better-compression-of-string', 'Medium', true, '', '[{"name": "Hash Table", "slug": "hash-table", "translatedName": null}, {"name": "String", "slug": "string", "translatedName": null}, {"name": "Sorting", "slug": "sorting", "translatedName": null}, {"name": "Counting", "slug": "counting", "translatedName": null}]'::jsonb, '"a3c9b2c1"
"c2b3a1"
"a2b4c1"', '[{"title": "String Compression", "titleSlug": "string-compression", "difficulty": "Medium", "translatedTitle": null}]'::jsonb, '2025-03-05 08:18:04.299');
INSERT INTO problems
(id, frontend_id, title, title_slug, difficulty, is_paid_only, "content", topic_tags, example_testcases, similar_questions, created_at)
VALUES(3449, 3141, 'Maximum Hamming Distances', 'maximum-hamming-distances', 'Hard', true, '', '[{"name": "Array", "slug": "array", "translatedName": null}, {"name": "Bit Manipulation", "slug": "bit-manipulation", "translatedName": null}, {"name": "Breadth-First Search", "slug": "breadth-first-search", "translatedName": null}]'::jsonb, '[9,12,9,11]
4
[3,4,6,10]
4', '[]'::jsonb, '2025-03-05 08:22:06.693');
INSERT INTO problems
(id, frontend_id, title, title_slug, difficulty, is_paid_only, "content", topic_tags, example_testcases, similar_questions, created_at)
VALUES(3393, 3088, 'Make String Anti-palindrome', 'make-string-anti-palindrome', 'Hard', true, '', '[{"name": "String", "slug": "string", "translatedName": null}, {"name": "Greedy", "slug": "greedy", "translatedName": null}, {"name": "Sorting", "slug": "sorting", "translatedName": null}, {"name": "Counting Sort", "slug": "counting-sort", "translatedName": null}]'::jsonb, '"abca"
"abba"
"cccd"', '[]'::jsonb, '2025-03-05 08:31:17.190');
INSERT INTO problems
(id, frontend_id, title, title_slug, difficulty, is_paid_only, "content", topic_tags, example_testcases, similar_questions, created_at)
VALUES(3378, 3073, 'Maximum Increasing Triplet Value', 'maximum-increasing-triplet-value', 'Medium', true, '', '[{"name": "Array", "slug": "array", "translatedName": null}, {"name": "Ordered Set", "slug": "ordered-set", "translatedName": null}]'::jsonb, '[5,6,9]', '[]'::jsonb, '2025-03-05 08:33:53.047');
INSERT INTO problems
(id, frontend_id, title, title_slug, difficulty, is_paid_only, "content", topic_tags, example_testcases, similar_questions, created_at)
VALUES(1383, 2198, 'Number of Single Divisor Triplets', 'number-of-single-divisor-triplets', 'Medium', true, '', '[{"name": "Math", "slug": "math", "translatedName": null}]'::jsonb, '[4,6,7,3,2]
[1,2,2]
[1,1,1]', '[{"title": "Count Array Pairs Divisible by K", "titleSlug": "count-array-pairs-divisible-by-k", "difficulty": "Hard", "translatedTitle": null}]'::jsonb, '2025-03-05 14:35:47.739');
INSERT INTO problems
(id, frontend_id, title, title_slug, difficulty, is_paid_only, "content", topic_tags, example_testcases, similar_questions, created_at)
VALUES(3368, 3062, 'Winner of the Linked List Game', 'winner-of-the-linked-list-game', 'Easy', true, '', '[{"name": "Linked List", "slug": "linked-list", "translatedName": null}]'::jsonb, '[2,1]', '[]'::jsonb, '2025-03-05 08:35:19.880');
INSERT INTO problems
(id, frontend_id, title, title_slug, difficulty, is_paid_only, "content", topic_tags, example_testcases, similar_questions, created_at)
VALUES(3333, 3023, 'Find Pattern in Infinite Stream I', 'find-pattern-in-infinite-stream-i', 'Medium', true, '', '[{"name": "Array", "slug": "array", "translatedName": null}, {"name": "Sliding Window", "slug": "sliding-window", "translatedName": null}, {"name": "Rolling Hash", "slug": "rolling-hash", "translatedName": null}, {"name": "String Matching", "slug": "string-matching", "translatedName": null}, {"name": "Hash Function", "slug": "hash-function", "translatedName": null}]'::jsonb, '[1,1,1,0,1]
[0,1]
[0]
[0]
[1,0,1,1,0,1]
[1,1,0,1]', '[]'::jsonb, '2025-03-05 08:40:17.314');
INSERT INTO problems
(id, frontend_id, title, title_slug, difficulty, is_paid_only, "content", topic_tags, example_testcases, similar_questions, created_at)
VALUES(3327, 3086, 'Minimum Moves to Pick K Ones', 'minimum-moves-to-pick-k-ones', 'Hard', false, '<p>You are given a binary array <code>nums</code> of length <code>n</code>, a <strong>positive</strong> integer <code>k</code> and a <strong>non-negative</strong> integer <code>maxChanges</code>.</p>

<p>Alice plays a game, where the goal is for Alice to pick up <code>k</code> ones from <code>nums</code> using the <strong>minimum</strong> number of <strong>moves</strong>. When the game starts, Alice picks up any index <code>aliceIndex</code> in the range <code>[0, n - 1]</code> and stands there. If <code>nums[aliceIndex] == 1</code> , Alice picks up the one and <code>nums[aliceIndex]</code> becomes <code>0</code>(this <strong>does not</strong> count as a move). After this, Alice can make <strong>any</strong> number of <strong>moves</strong> (<strong>including</strong> <strong>zero</strong>) where in each move Alice must perform <strong>exactly</strong> one of the following actions:</p>

<ul>
	<li>Select any index <code>j != aliceIndex</code> such that <code>nums[j] == 0</code> and set <code>nums[j] = 1</code>. This action can be performed <strong>at</strong> <strong>most</strong> <code>maxChanges</code> times.</li>
	<li>Select any two adjacent indices <code>x</code> and <code>y</code> (<code>|x - y| == 1</code>) such that <code>nums[x] == 1</code>, <code>nums[y] == 0</code>, then swap their values (set <code>nums[y] = 1</code> and <code>nums[x] = 0</code>). If <code>y == aliceIndex</code>, Alice picks up the one after this move and <code>nums[y]</code> becomes <code>0</code>.</li>
</ul>

<p>Return <em>the <strong>minimum</strong> number of moves required by Alice to pick <strong>exactly </strong></em><code>k</code> <em>ones</em>.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<div class="example-block" style="border-color: var(--border-tertiary); border-left-width: 2px; color: var(--text-secondary); font-size: .875rem; margin-bottom: 1rem; margin-top: 1rem; overflow: visible; padding-left: 1rem;">
<p><strong>Input: </strong><span class="example-io" style="font-family: Menlo,sans-serif; font-size: 0.85rem;">nums = [1,1,0,0,0,1,1,0,0,1], k = 3, maxChanges = 1</span></p>

<p><strong>Output: </strong><span class="example-io" style="font-family: Menlo,sans-serif; font-size: 0.85rem;">3</span></p>

<p><strong>Explanation:</strong> Alice can pick up <code>3</code> ones in <code>3</code> moves, if Alice performs the following actions in each move when standing at <code>aliceIndex == 1</code>:</p>

<ul>
	<li>At the start of the game Alice picks up the one and <code>nums[1]</code> becomes <code>0</code>. <code>nums</code> becomes <code>[1,<strong><u>0</u></strong>,0,0,0,1,1,0,0,1]</code>.</li>
	<li>Select <code>j == 2</code> and perform an action of the first type. <code>nums</code> becomes <code>[1,<strong><u>0</u></strong>,1,0,0,1,1,0,0,1]</code></li>
	<li>Select <code>x == 2</code> and <code>y == 1</code>, and perform an action of the second type. <code>nums</code> becomes <code>[1,<strong><u>1</u></strong>,0,0,0,1,1,0,0,1]</code>. As <code>y == aliceIndex</code>, Alice picks up the one and <code>nums</code> becomes <code>[1,<strong><u>0</u></strong>,0,0,0,1,1,0,0,1]</code>.</li>
	<li>Select <code>x == 0</code> and <code>y == 1</code>, and perform an action of the second type. <code>nums</code> becomes <code>[0,<strong><u>1</u></strong>,0,0,0,1,1,0,0,1]</code>. As <code>y == aliceIndex</code>, Alice picks up the one and <code>nums</code> becomes <code>[0,<strong><u>0</u></strong>,0,0,0,1,1,0,0,1]</code>.</li>
</ul>

<p>Note that it may be possible for Alice to pick up <code>3</code> ones using some other sequence of <code>3</code> moves.</p>
</div>

<p><strong class="example">Example 2:</strong></p>

<div class="example-block" style="border-color: var(--border-tertiary); border-left-width: 2px; color: var(--text-secondary); font-size: .875rem; margin-bottom: 1rem; margin-top: 1rem; overflow: visible; padding-left: 1rem;">
<p><strong>Input: </strong><span class="example-io" style="font-family: Menlo,sans-serif; font-size: 0.85rem;">nums = [0,0,0,0], k = 2, maxChanges = 3</span></p>

<p><strong>Output: </strong><span class="example-io" style="font-family: Menlo,sans-serif; font-size: 0.85rem;">4</span></p>

<p><strong>Explanation:</strong> Alice can pick up <code>2</code> ones in <code>4</code> moves, if Alice performs the following actions in each move when standing at <code>aliceIndex == 0</code>:</p>

<ul>
	<li>Select <code>j == 1</code> and perform an action of the first type. <code>nums</code> becomes <code>[<strong><u>0</u></strong>,1,0,0]</code>.</li>
	<li>Select <code>x == 1</code> and <code>y == 0</code>, and perform an action of the second type. <code>nums</code> becomes <code>[<strong><u>1</u></strong>,0,0,0]</code>. As <code>y == aliceIndex</code>, Alice picks up the one and <code>nums</code> becomes <code>[<strong><u>0</u></strong>,0,0,0]</code>.</li>
	<li>Select <code>j == 1</code> again and perform an action of the first type. <code>nums</code> becomes <code>[<strong><u>0</u></strong>,1,0,0]</code>.</li>
	<li>Select <code>x == 1</code> and <code>y == 0</code> again, and perform an action of the second type. <code>nums</code> becomes <code>[<strong><u>1</u></strong>,0,0,0]</code>. As <code>y == aliceIndex</code>, Alice picks up the one and <code>nums</code> becomes <code>[<strong><u>0</u></strong>,0,0,0]</code>.</li>
</ul>
</div>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>2 &lt;= n &lt;= 10<sup>5</sup></code></li>
	<li><code>0 &lt;= nums[i] &lt;= 1</code></li>
	<li><code>1 &lt;= k &lt;= 10<sup>5</sup></code></li>
	<li><code>0 &lt;= maxChanges &lt;= 10<sup>5</sup></code></li>
	<li><code>maxChanges + sum(nums) &gt;= k</code></li>
</ul>
', '[{"name": "Array", "slug": "array", "translatedName": null}, {"name": "Greedy", "slug": "greedy", "translatedName": null}, {"name": "Sliding Window", "slug": "sliding-window", "translatedName": null}, {"name": "Prefix Sum", "slug": "prefix-sum", "translatedName": null}]'::jsonb, '[1,1,0,0,0,1,1,0,0,1]
3
1
[0,0,0,0]
2
3', '[{"title": "Minimum Swaps to Group All 1''s Together", "titleSlug": "minimum-swaps-to-group-all-1s-together", "difficulty": "Medium", "translatedTitle": null}]'::jsonb, '2025-03-05 08:41:22.222');
INSERT INTO problems
(id, frontend_id, title, title_slug, difficulty, is_paid_only, "content", topic_tags, example_testcases, similar_questions, created_at)
VALUES(3323, 3018, 'Maximum Number of Removal Queries That Can Be Processed I', 'maximum-number-of-removal-queries-that-can-be-processed-i', 'Hard', true, '', '[{"name": "Array", "slug": "array", "translatedName": null}, {"name": "Dynamic Programming", "slug": "dynamic-programming", "translatedName": null}]'::jsonb, '[1,2,3,4,5]
[1,2,3,4,6]
[2,3,2]
[2,2,3]
[3,4,3]
[4,3,2]', '[]'::jsonb, '2025-03-05 08:42:05.196');
INSERT INTO problems
(id, frontend_id, title, title_slug, difficulty, is_paid_only, "content", topic_tags, example_testcases, similar_questions, created_at)
VALUES(158, 158, 'Read N Characters Given read4 II - Call Multiple Times', 'read-n-characters-given-read4-ii-call-multiple-times', 'Hard', true, '', '[{"name": "Array", "slug": "array", "translatedName": null}, {"name": "Simulation", "slug": "simulation", "translatedName": null}, {"name": "Interactive", "slug": "interactive", "translatedName": null}]'::jsonb, '"abc"
[1,2,1]
"abc"
[4,1]', '[{"title": "Read N Characters Given Read4", "titleSlug": "read-n-characters-given-read4", "difficulty": "Easy", "translatedTitle": null}]'::jsonb, '2025-03-05 17:48:57.993');
INSERT INTO problems
(id, frontend_id, title, title_slug, difficulty, is_paid_only, "content", topic_tags, example_testcases, similar_questions, created_at)
VALUES(3254, 2964, 'Number of Divisible Triplet Sums', 'number-of-divisible-triplet-sums', 'Medium', true, '', '[{"name": "Array", "slug": "array", "translatedName": null}, {"name": "Hash Table", "slug": "hash-table", "translatedName": null}]'::jsonb, '[3,3,4,7,8]
5
[3,3,3,3]
3
[3,3,3,3]
6', '[]'::jsonb, '2025-03-05 08:50:34.331');
INSERT INTO problems
(id, frontend_id, title, title_slug, difficulty, is_paid_only, "content", topic_tags, example_testcases, similar_questions, created_at)
VALUES(157, 157, 'Read N Characters Given Read4', 'read-n-characters-given-read4', 'Easy', true, '', '[{"name": "Array", "slug": "array", "translatedName": null}, {"name": "Simulation", "slug": "simulation", "translatedName": null}, {"name": "Interactive", "slug": "interactive", "translatedName": null}]'::jsonb, '"abc"
4
"abcde"
5
"abcdABCD1234"
12', '[{"title": "Read N Characters Given read4 II - Call Multiple Times", "titleSlug": "read-n-characters-given-read4-ii-call-multiple-times", "difficulty": "Hard", "translatedTitle": null}]'::jsonb, '2025-03-05 17:49:09.032');
INSERT INTO problems
(id, frontend_id, title, title_slug, difficulty, is_paid_only, "content", topic_tags, example_testcases, similar_questions, created_at)
VALUES(3216, 2927, 'Distribute Candies Among Children III', 'distribute-candies-among-children-iii', 'Hard', true, '', '[{"name": "Math", "slug": "math", "translatedName": null}, {"name": "Combinatorics", "slug": "combinatorics", "translatedName": null}]'::jsonb, '5
2
3
3', '[]'::jsonb, '2025-03-05 08:57:18.518');
INSERT INTO problems
(id, frontend_id, title, title_slug, difficulty, is_paid_only, "content", topic_tags, example_testcases, similar_questions, created_at)
VALUES(3204, 2921, 'Maximum Profitable Triplets With Increasing Prices II', 'maximum-profitable-triplets-with-increasing-prices-ii', 'Hard', true, '', '[{"name": "Array", "slug": "array", "translatedName": null}, {"name": "Binary Indexed Tree", "slug": "binary-indexed-tree", "translatedName": null}, {"name": "Segment Tree", "slug": "segment-tree", "translatedName": null}]'::jsonb, '[10,2,3,4]
[100,2,7,10]
[1,2,3,4,5]
[1,5,3,4,6]
[4,3,2,1]
[33,20,19,87]', '[]'::jsonb, '2025-03-05 08:59:16.432');
INSERT INTO problems
(id, frontend_id, title, title_slug, difficulty, is_paid_only, "content", topic_tags, example_testcases, similar_questions, created_at)
VALUES(3198, 2912, 'Number of Ways to Reach Destination in the Grid', 'number-of-ways-to-reach-destination-in-the-grid', 'Hard', true, '', '[{"name": "Math", "slug": "math", "translatedName": null}, {"name": "Dynamic Programming", "slug": "dynamic-programming", "translatedName": null}, {"name": "Combinatorics", "slug": "combinatorics", "translatedName": null}]'::jsonb, '3
2
2
[1,1]
[2,2]
3
4
3
[1,2]
[2,3]', '[]'::jsonb, '2025-03-05 09:00:23.066');
INSERT INTO problems
(id, frontend_id, title, title_slug, difficulty, is_paid_only, "content", topic_tags, example_testcases, similar_questions, created_at)
VALUES(356, 356, 'Line Reflection', 'line-reflection', 'Medium', true, '', '[{"name": "Array", "slug": "array", "translatedName": null}, {"name": "Hash Table", "slug": "hash-table", "translatedName": null}, {"name": "Math", "slug": "math", "translatedName": null}]'::jsonb, '[[1,1],[-1,1]]
[[1,1],[-1,-1]]', '[{"title": "Max Points on a Line", "titleSlug": "max-points-on-a-line", "difficulty": "Hard", "translatedTitle": null}, {"title": "Number of Boomerangs", "titleSlug": "number-of-boomerangs", "difficulty": "Medium", "translatedTitle": null}]'::jsonb, '2025-03-05 17:16:26.318');
INSERT INTO problems
(id, frontend_id, title, title_slug, difficulty, is_paid_only, "content", topic_tags, example_testcases, similar_questions, created_at)
VALUES(3158, 2863, 'Maximum Length of Semi-Decreasing Subarrays', 'maximum-length-of-semi-decreasing-subarrays', 'Medium', true, '', '[{"name": "Array", "slug": "array", "translatedName": null}, {"name": "Stack", "slug": "stack", "translatedName": null}, {"name": "Sorting", "slug": "sorting", "translatedName": null}, {"name": "Monotonic Stack", "slug": "monotonic-stack", "translatedName": null}]'::jsonb, '[7,6,5,4,3,2,1,6,10,11]
[57,55,50,60,61,58,63,59,64,60,63]
[1,2,3,4]', '[]'::jsonb, '2025-03-05 09:05:54.607');
INSERT INTO problems
(id, frontend_id, title, title_slug, difficulty, is_paid_only, "content", topic_tags, example_testcases, similar_questions, created_at)
VALUES(3016, 2792, 'Count Nodes That Are Great Enough', 'count-nodes-that-are-great-enough', 'Hard', true, '', '[{"name": "Divide and Conquer", "slug": "divide-and-conquer", "translatedName": null}, {"name": "Tree", "slug": "tree", "translatedName": null}, {"name": "Depth-First Search", "slug": "depth-first-search", "translatedName": null}, {"name": "Binary Tree", "slug": "binary-tree", "translatedName": null}]'::jsonb, '[7,6,5,4,3,2,1]
2
[1,2,3]
1
[3,1,2]
2', '[]'::jsonb, '2025-03-05 09:15:07.565');
INSERT INTO problems
(id, frontend_id, title, title_slug, difficulty, is_paid_only, "content", topic_tags, example_testcases, similar_questions, created_at)
VALUES(2995, 2782, 'Number of Unique Categories', 'number-of-unique-categories', 'Medium', true, '', '[{"name": "Union Find", "slug": "union-find", "translatedName": null}, {"name": "Interactive", "slug": "interactive", "translatedName": null}, {"name": "Counting", "slug": "counting", "translatedName": null}]'::jsonb, '6
[1,1,2,2,3,3]
5
[1,2,3,4,5]
3
[1,1,1]', '[]'::jsonb, '2025-03-05 09:16:01.422');
INSERT INTO problems
(id, frontend_id, title, title_slug, difficulty, is_paid_only, "content", topic_tags, example_testcases, similar_questions, created_at)
VALUES(2897, 2753, 'Count Houses in a Circular Street II', 'count-houses-in-a-circular-street-ii', 'Hard', true, '', '[]'::jsonb, '[1,1,1,1]
10
[1,0,1,1,0]
5', '[]'::jsonb, '2025-03-05 09:19:49.439');